package rest

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

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

type Clients map[chan ClientMessage]struct{}

type EventServer struct {
	messages      chan ClientMessage
	newClients    chan SSEClient
	closedClients chan chan ClientMessage
	subscriptions map[models.Topic]Clients
	// clients represents all connected clients
	clients Clients
}

func (es *EventServer) Publish(message string, topic models.Topic) {
	/**
	 * TODO: instead of using Publish, add events to queue via Queue(message, topic) and dispatch later based on x-request-id (not user id since some events may
	 *  not have authenticated user, e.g. upon registering the first time we may require manual verification, etc.)
	 * the Queue method (and all other public ones) will accept c.Request.Context, where req id is set
	 */
	es.messages <- ClientMessage{Message: message, Topic: topic}
}

func (es *EventServer) listen() {
	for {
		select {
		case client := <-es.newClients:
			es.clients[client.Chan] = struct{}{}
			for _, topic := range client.Topics {
				if _, ok := es.subscriptions[topic]; !ok {
					es.subscriptions[topic] = make(Clients)
				}
				es.subscriptions[topic][client.Chan] = struct{}{}
			}
			log.Printf("Client added. %d registered clients", len(es.clients))

		case client := <-es.closedClients:
			for _, subscriptions := range es.subscriptions {
				delete(subscriptions, client)
			}
			delete(es.clients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(es.clients))

		case eventMsg := <-es.messages:
			for ch := range es.subscriptions[eventMsg.Topic] {
				ch <- eventMsg
			}
		}
	}
}

func (es *EventServer) Subscribe(topics models.Topics) (client SSEClient, unsubscribe func()) {
	client = SSEClient{
		Chan:   make(chan ClientMessage, 1),
		Topics: make(Topics, len(topics)),
	}
	for i, topic := range topics {
		client.Topics[i] = models.Topic(topic)
	}

	es.newClients <- client

	return client, func() {
		es.closedClients <- client.Chan
	}
}

func newSSEServer() *EventServer {
	es := &EventServer{
		messages:      make(chan ClientMessage),
		newClients:    make(chan SSEClient),
		closedClients: make(chan chan ClientMessage),
		subscriptions: make(map[models.Topic]Clients),
		clients:       make(Clients),
	}

	go es.listen()

	return es
}

type SSEClient struct {
	Chan   chan ClientMessage
	Topics Topics
}

type ClientMessage struct {
	Message string
	Topic   models.Topic
}

var r = Events200TexteventStreamResponse{Body: strings.NewReader("")}

// Events represents server sent events.
func (h *StrictHandlers) Events(c *gin.Context, request EventsRequestObject) (EventsResponseObject, error) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	// TODO: request.Params.ProjectName

	clientChan, unsubscribe := h.event.Subscribe(request.Params.Topics)
	defer unsubscribe()

	ticker := time.NewTicker(1 * time.Second)
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-clientChan.Chan:
			if !ok {
				return false // channel closed
			}
			c.SSEvent(string(msg.Topic), msg.Message)
			return true
		case <-ticker.C:
			// ensure handler can return and clean up when client disconnects and no messages have been sent
			return true
		case <-c.Request.Context().Done():
			fmt.Printf("Client gone. Context cancelled: %v\n", c.Request.Context().Err())
			c.SSEvent("message", "channel closed")

			return false
		}
	})

	return r, nil
}
