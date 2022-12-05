package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

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
	ClientsForMessage1 map[chan string]bool
	ClientsForMessage2 map[chan string]bool
}

// New event messages are broadcasted to all registered client connection channels.
type ClientChan chan string

type UserNotificationsChan chan string

type timeEvent struct {
	Foo string `json:"foo"`
	Msg string `json:"msg"`
}

// Events represents server events.
func (h *Handlers) Events(c *gin.Context) {
	c.Set(skipRequestValidation, true)
	c.Set(skipResponseValidation, true)

	fmt.Printf("c.Copy().Keys (Events): %v\n", c.Copy().Keys)

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
			// this part should actually be done before sending to channel.
			// should it fail marshalling earlier, msg is an error message (literal or {"error":"..."})
			te := &timeEvent{
				Foo: "bar",
				Msg: msg,
			}
			sseMsg, err := json.Marshal(te)
			if err != nil {
				c.SSEvent("message", "could not marshal message")

				return true // should probably continue regardless. client should handle the error and stop/continue if desired
			}
			c.SSEvent("message", string(sseMsg))
			c.Writer.Flush()

			return true
		case msg, ok := <-userNotificationsChan:
			if !ok {
				return true
			}
			c.SSEvent("userNotificationsChan", msg)
			c.Writer.Flush()

			return true
		case <-c.Request.Context().Done():
			fmt.Println("CLIENT GONE")

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
		ClientsForMessage1: make(map[chan string]bool),
		ClientsForMessage2: make(map[chan string]bool),
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
			stream.ClientsForMessage1[client] = true
			log.Printf("Client for message type 1 added. %d registered clients", len(stream.ClientsForMessage1))
		case client := <-stream.NewClients2:
			stream.ClientsForMessage2[client] = true
			log.Printf("Client for message type 2 added. %d registered clients", len(stream.ClientsForMessage2))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.ClientsForMessage1, client)
			delete(stream.ClientsForMessage2, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.ClientsForMessage1))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			fmt.Printf("eventMsg: %v\n", eventMsg)
			for clientMessageChan := range stream.ClientsForMessage1 {
				clientMessageChan <- eventMsg
			}

		// Broadcast message 2to client
		case eventMsg := <-stream.Message2:
			fmt.Printf("eventMsg (2): %v\n", eventMsg)
			for clientMessageChan := range stream.ClientsForMessage2 {
				clientMessageChan <- eventMsg
			}
		}
	}
}

// TODO see if can reproduce https://github.com/gin-gonic/gin/issues/3142
// some bugs where fixed in sse example committed 4 months later, so...
func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("stream events - Initialize client channel")
		clientChan := make(ClientChan)
		userNotificationsChan := make(UserNotificationsChan)

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
