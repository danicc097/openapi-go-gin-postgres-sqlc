package rest

import (
	"context"
	"io"
	"strings"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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


channel use cases,etc:
 - https://go101.org/article/channel-use-cases.html
 - https://go101.org/article/channel.html
*/

type Clients map[chan ClientMessage]struct{}

type EventServer struct {
	queueMu sync.RWMutex
	subsMu  sync.RWMutex

	logger *zap.SugaredLogger
	// queuedMessages are pending messages for a given X-Request-ID.
	queuedMessages map[string][]ClientMessage
	messages       chan ClientMessage
	newClients     chan SSEClient
	closedClients  chan chan ClientMessage
	subscriptions  map[models.Topic]Clients
	// clients represents all connected clients
	clients Clients
}

// NewEventServer returns an SSE server and starts listening for messages.
func NewEventServer(logger *zap.SugaredLogger) *EventServer {
	es := &EventServer{
		logger:         logger,
		messages:       make(chan ClientMessage),
		queuedMessages: map[string][]ClientMessage{},
		newClients:     make(chan SSEClient),
		closedClients:  make(chan chan ClientMessage),
		subscriptions:  make(map[models.Topic]Clients),
		clients:        make(Clients),
	}

	go es.listen()

	return es
}

// Queue saves an event for later publishing.
func (es *EventServer) Queue(ctx context.Context, message string, topic models.Topic) {
	es.queueMu.Lock()
	defer es.queueMu.Unlock()

	rid := GetRequestIDFromCtx(ctx)

	es.queuedMessages[rid] = append(es.queuedMessages[rid], ClientMessage{Message: message, Topic: topic})
}

// Publish publishes an event immediately.
func (es *EventServer) Publish(message string, topic models.Topic) {
	// es.logger.Debugf("topic %s: sending event %v", topic, message)
	es.messages <- ClientMessage{Message: message, Topic: topic}
}

func (es *EventServer) listen() {
	for {
		select {
		case client := <-es.newClients:
			es.clients[client.Chan] = struct{}{}
			es.subsMu.Lock()
			for _, topic := range client.Topics {
				if _, ok := es.subscriptions[topic]; !ok {
					es.subscriptions[topic] = make(Clients)
				}
				es.subscriptions[topic][client.Chan] = struct{}{}
			}
			es.subsMu.Unlock()
			es.logger.Infof("Client added. %d registered clients", len(es.clients))

		case client := <-es.closedClients:
			es.subsMu.Lock()
			for _, subscriptions := range es.subscriptions {
				delete(subscriptions, client)
			}
			es.subsMu.Unlock()
			delete(es.clients, client)
			close(client)
			es.logger.Infof("Removed client. %d registered clients", len(es.clients))

		case eventMsg := <-es.messages:
			for ch := range es.subscriptions[eventMsg.Topic] {
				select {
				case ch <- eventMsg:
					// message sent successfully
				default:
					// channel closed and still not removed for some reason
					es.subsMu.Lock()
					delete(es.subscriptions[eventMsg.Topic], ch)
					es.subsMu.Unlock()
				}
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

// EventDispatcher represents middleware whose cleanup fires pending events for a request.
func (es *EventServer) EventDispatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		es.queueMu.Lock()
		defer es.queueMu.Unlock()

		rid := GetRequestIDFromCtx(c.Request.Context())
		defer func() { delete(es.queuedMessages, rid) }()
		qm := es.queuedMessages[rid]

		if CtxRequestHasError(c) {
			es.logger.Infof("request %s marked as failed", rid)
			for _, m := range qm {
				if m.SendOnFailedRequest {
					es.Publish(m.Message, m.Topic)
				}
			}
			c.Next()

			return
		}

		for _, m := range qm {
			es.Publish(m.Message, m.Topic)
		}

		c.Next()
	}
}

type SSEClient struct {
	Chan   chan ClientMessage
	Topics Topics
}

type ClientMessage struct {
	Message string
	Topic   models.Topic
	// SendOnFailedRequest defines whether the event should be dispatched
	// even on failed requests.
	SendOnFailedRequest bool
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

	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-clientChan.Chan:
			if !ok {
				return false // channel closed
			}
			c.SSEvent(string(msg.Topic), msg.Message)
			return true
		case <-c.Request.Context().Done():
			// ensure handler can return and clean up when client disconnects and no messages have been sent
			h.logger.Debugf("Client gone: %v\n", c.Request.Context().Err())

			return false
		}
	})

	return r, nil
}
