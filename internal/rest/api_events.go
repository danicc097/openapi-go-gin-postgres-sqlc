package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message  chan string
	Message2 chan string

	// New client connections
	NewClients  chan chan string
	NewClients2 chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	ClientsForMessage1 map[chan string]struct{}
	ClientsForMessage2 map[chan string]struct{}
}

type subs map[chan string]struct{}

type PubSub struct {
	mu sync.RWMutex // preferable if mostly read

	// TODO will need to have nested map for subs per project ( TODO models.Project)
	// all in all this whole PubSub arch requires more memory usage but provides more flexible notifications.
	// need to experiment buffered vs unbuffered. see https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/
	// we can't have shared topic for every project unless it's GlobalAlerts (degraded performance, maintenance alert)
	// a user will be subscribed to the current project selected in frontend.
	// this info needs to be sent when calling /events in a query param

	subs map[models.Topics]subs
	// e.g. event card moved notifies all card members' userID - attempt to send if clients are connected (i.e. key exists in personalSub)
	// TODO close and delete entries when client is gone:
	// edit defer func in serveHTTP to get user from context
	personalSub map[string]chan string
	closed      bool
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
	ps.personalSub = make(map[string]chan string)

	for _, t := range models.AllTopicsValues() {
		ps.subs[t] = make(subs)
	}

	return ps
}

func (ps *PubSub) Subscribe(topic models.Topics, userID string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	sub := make(chan string, 1)
	ps.subs[topic][sub] = struct{}{}

	// create personal sub chan if it doesn't exist already
	if _, ok := ps.personalSub[userID]; !ok {
		c := make(chan string, 1)
		ps.personalSub[userID] = c
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

// PublishTo sends a direct message to the specified user IDs.
func (ps *PubSub) PublishTo(userIDs []string, msg string) {
	ps.mu.RLock() // only reading subs
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, id := range userIDs {
		if ch, ok := ps.personalSub[id]; ok {
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
		for uid, sub := range ps.personalSub {
			close(sub)
			delete(ps.personalSub, uid)
		}
	}
}

// New event messages are broadcasted to all registered client connection channels.
type ClientChan chan string

type UserNotificationsChan chan string

type timeEvent struct {
	Foo string `json:"foo"`
	Msg string `json:"msg"`
}

/**
 *
 *
let evtSource = new EventSource('/v2/events');

evtSource.onmessage = (e) => {
  console.log(e)
}
evtSource.addEventListener(<Event name>, (e) => {
  console.log(e)
});


TODO for actual implementation see: https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/

channel use cases,etc:
 - https://go101.org/article/channel-use-cases.html
 - https://go101.org/article/channel.html
*/

// Events represents server events.
// TODO requires query param projectId=...
// to subscribe to the current project's topics only
func (h *Handlers) Events(c *gin.Context, params models.EventsParams) {
	c.Set(skipRequestValidationCtxKey, true)
	clientChan, ok := c.Value("clientChan").(ClientChan)
	if !ok {
		return
	}
	userNotificationsChan, ok := c.Value("userNotificationsChan").(UserNotificationsChan)
	if !ok {
		return
	}
	// TODO map of channels for each Role ('global' notif.) ?
	// TODO map of channels for every connected user for 'personal' notif. ?
	// will use to alert cards moving, selected as card member, etc.
	// test with curl -X 'GET' -N 'https://localhost:8090/v2/events' 'https://localhost:8090/v2/'
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
				Foo: "bar",
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
}

// newSSEServer initializes events and starts processing requests.
func newSSEServer() *Event {
	event := &Event{
		Message:            make(chan string),
		Message2:           make(chan string),
		NewClients:         make(chan chan string),
		NewClients2:        make(chan chan string),
		ClosedClients:      make(chan chan string),
		ClientsForMessage1: make(map[chan string]struct{}),
		ClientsForMessage2: make(map[chan string]struct{}),
	}

	go event.listen()

	return event
}

// It Listens all incoming requests from clients.
// Handles addition and removal of clients and broadcast messages to clients.
func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.ClientsForMessage1[client] = struct{}{}
			log.Printf("Client for message type 1 added. %d registered clients", len(stream.ClientsForMessage1))
		case client := <-stream.NewClients2:
			stream.ClientsForMessage2[client] = struct{}{}
			log.Printf("Client for message type 2 added. %d registered clients", len(stream.ClientsForMessage2))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.ClientsForMessage1, client)
			delete(stream.ClientsForMessage2, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.ClientsForMessage1))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			// fmt.Printf("eventMsg: %v\n", eventMsg)
			for clientMessageChan := range stream.ClientsForMessage1 {
				clientMessageChan <- eventMsg
			}

		// Broadcast message 2to client
		case eventMsg := <-stream.Message2:
			// fmt.Printf("eventMsg (2): %v\n", eventMsg)
			for clientMessageChan := range stream.ClientsForMessage2 {
				clientMessageChan <- eventMsg
			}
		}
	}
}

// TODO see if can reproduce https://github.com/gin-gonic/gin/issues/3142
// some bugs were fixed in sse example committed 4 months later, so...
func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("stream events - Initialize client channel")
		clientChan := make(ClientChan, 1)
		userNotificationsChan := make(UserNotificationsChan, 1)

		// Send new connection to event server
		stream.NewClients <- clientChan
		stream.NewClients2 <- userNotificationsChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
			stream.ClosedClients <- userNotificationsChan
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
