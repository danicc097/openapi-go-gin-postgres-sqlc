package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
// debug via curl -X 'GET' -N 'https://localhost:8090/v2/events?projectName=demo'
type EventServer struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message  chan string
	Message2 chan string

	// New client connections
	NewClients  chan ClientChan
	NewClients2 chan ClientChan

	// Closed client connections
	ClosedClients chan ClientChan

	// Total client connections
	Message1Subscribers subs
}

type subs map[ClientChan]struct{}

type PubSub struct {
	mu sync.RWMutex // preferable if mostly read

	// TODO will need to have nested map for subs per project ( TODO models.Project)
	// all in all this whole PubSub arch requires more memory usage but provides more flexible notifications.
	// need to experiment buffered vs unbuffered. see https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/
	// we can't have shared topic for every project unless it's GlobalAlerts (degraded performance, maintenance alert)
	// a user will be subscribed to the current project selected in frontend.
	// this info needs to be sent when calling /events in a query param

	subs map[models.Topics]subs
	// e.g. event card moved, assigned, etc. notifies card members by userID if clients are connected (i.e. key exists in connectedUsers)
	// TODO close and delete entries when client is gone:
	// edit defer func in serveHTTP to get user from context
	connectedUsers map[db.UserID]ClientChan
	closed         bool
}

// in reality most messages wont run endlessly in a goroutine, we will e.g.
// send a message to event.WorkItemMoved every time services.WorkItems.Move(...) is called
// we would have a map of channels opened per user id like WorkItemMovedUserChans map[string]chan string
// (in the future we'll have many of these like MemberAssignedUserChans to notify when added to workitem, etc.)
// and send a message to all members with workitem.members' userIDs.
// We need specific channels so that when a message is consumed we add a hardcoded "event" name
// ( see enum Topics) and frontend properly handles it.
func NewPubSub() *PubSub {
	ps := &PubSub{}
	ps.subs = make(map[models.Topics]subs)
	ps.connectedUsers = make(map[db.UserID]ClientChan)

	for _, t := range models.AllTopicsValues() {
		ps.subs[t] = make(subs)
	}

	return ps
}

func (ps *PubSub) Subscribe(topic models.Topics, userID db.UserID) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	sub := make(chan string, 1)
	ps.subs[topic][sub] = struct{}{}

	// create personal sub chan if it doesn't exist already
	if _, ok := ps.connectedUsers[userID]; !ok {
		c := make(chan string, 1)
		ps.connectedUsers[userID] = c
	}

	return sub
}

func (ps *PubSub) Publish(topic models.Topics, msg string) {
	ps.mu.RLock() // only reading subs
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for sub := range ps.subs[topic] {
		sub <- msg
	}
}

// PushNotification sends the given message to the specified user IDs.
func (ps *PubSub) PushNotification(userIDs []db.UserID, msg string) {
	ps.mu.RLock() // only reading subs
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, id := range userIDs {
		if ch, ok := ps.connectedUsers[id]; ok {
			ch <- msg
		}
	}
}

func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for sub := range subs {
				close(sub)
				delete(subs, sub)
			}
		}
		for uid, sub := range ps.connectedUsers {
			close(sub)
			delete(ps.connectedUsers, uid)
		}
	}
}

// New event messages are broadcasted to all registered client connection channels.
type ClientChan chan string

type UserNotificationsChan ClientChan

type timeEvent struct {
	Foo string `json:"foo"`
	Msg string `json:"msg"`
}

/**
 *
 *
let evtSource = new EventSource('https://localhost:8090/v2/events?projectName=demo');

evtSource.onmessage = (e) => {
  console.log(e);
};
evtSource.addEventListener(<Event name>, (e) => {
  console.log(e);
});


TODO for actual implementation see: https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/

channel use cases,etc:
 - https://go101.org/article/channel-use-cases.html
 - https://go101.org/article/channel.html
*/

// newSSEServer initializes events and starts processing requests.
func newSSEServer() *EventServer {
	event := &EventServer{
		Message:             make(chan string),
		Message2:            make(chan string),
		NewClients:          make(chan ClientChan),
		NewClients2:         make(chan ClientChan),
		ClosedClients:       make(chan ClientChan),
		Message1Subscribers: make(map[ClientChan]struct{}),
	}

	go event.listen()

	return event
}

// It Listens all incoming requests from clients.
// Handles addition and removal of clients and broadcast messages to clients.
func (es *EventServer) listen() {
	for {
		select {
		// Add new available client
		case client := <-es.NewClients:
			es.Message1Subscribers[client] = struct{}{}
			log.Printf("Client for message type 1 added. %d registered clients", len(es.Message1Subscribers))

		// Remove closed client
		case client := <-es.ClosedClients:
			delete(es.Message1Subscribers, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(es.Message1Subscribers))

		// Broadcast message to client
		case eventMsg := <-es.Message:
			for clientMessageChan := range es.Message1Subscribers {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *EventServer) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* TODO: see cmd/sse-test/main.go working example

		   - sse withCredentials includes cookie with auth token. not authenticated? return (test with x-api-key)
		*/
		fmt.Println("stream events - Initialize client channel")

		clientChan := make(ClientChan, 1)
		userNotificationsChan := make(UserNotificationsChan, 1)

		// Send new connection to event server
		stream.NewClients <- clientChan
		stream.NewClients2 <- ClientChan(userNotificationsChan)

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
			stream.ClosedClients <- ClientChan(userNotificationsChan)
		}()

		c.Set("clientChan", clientChan)
		c.Set("userNotificationsChan", userNotificationsChan)

		c.Next()
	}
}

func SSEHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

// Events represents server events.
// TODO requires query param projectId=...
// to subscribe to the current project's topics only.
func (h *StrictHandlers) Events(c *gin.Context, request EventsRequestObject) (EventsResponseObject, error) {
	c.Set(skipRequestValidationCtxKey, true)
	clientChan, ok := c.Value("clientChan").(ClientChan)
	if !ok {
		return nil, errors.New("clientChan missing")
	}
	userNotificationsChan, ok := c.Value("userNotificationsChan").(UserNotificationsChan)
	if !ok {
		return nil, errors.New("userNotificationsChan missing")
	}
	// TODO map of channels for each Role ('global' notif.) ?
	// TODO map of channels for every connected user for 'personal' notif. ?
	// will use to alert cards moving, selected as card member, etc.
	// test with curl -X 'GET' -N 'https://localhost:8090/v2/events?projectName=demo'
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-clientChan:
			if !ok {
				return true
			}
			// TODO this part should actually be done before sending to channel.
			// should it fail marshalling earlier, msg is an error message (literal or {"error":"..."})
			// but all sse channels receive strings
			te := &timeEvent{
				Foo: "bar\n\n",
				Msg: msg,
			}
			sseMsg, err := json.Marshal(te)
			if err != nil {
				c.SSEvent("message", "could not marshal message")

				return true // should probably continue regardless. client should handle the error and stop/continue if desired
			}
			c.Render(-1, sse.Event{
				Event: "test-event",
				Data:  string(sseMsg),
			})

			c.Writer.Flush()

			return true
		case msg, ok := <-userNotificationsChan:
			if !ok {
				return true
			}
			c.Render(-1, sse.Event{
				Event: string(models.TopicsGlobalAlerts),
				Data:  msg,
			})
			c.Writer.Flush()

			return true
		case <-c.Request.Context().Done():
			fmt.Printf("Client gone. Context cancelled: %v\n", c.Request.Context().Err())
			c.SSEvent("message", "channel closed")
			c.Writer.Flush()

			return false
		}
	})

	return Events200TexteventStreamResponse{Body: strings.NewReader("")}, nil
}
