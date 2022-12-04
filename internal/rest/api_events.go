package rest

import (
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcasted to all registered client connection channels.
type ClientChan chan string

// Events represents server events.
func (h *Handlers) Events(c *gin.Context) {
	c.Set(skipRequestValidation, true)
	c.Set(skipResponseValidation, true)

	fmt.Printf("c.Copy().Keys (Events): %v\n", c.Copy().Keys)

	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}
	// TODO map of channels for each Role ('global' notif.) ?
	// TODO map of channels for every connected user for 'personal' notif. ?
	// will use to alert cards moving, selected as card member, etc.
	// test with curl -X 'GET' -N 'https://localhost:8090/v2/events' 'https://localhost:8090/v2/'
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)

			return true
		}
		c.SSEvent("message", "channel closed")

		return false
	})
}

// newSSEServer initializes events and starts processing requests.
func newSSEServer() *Event {
	event := &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
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
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			fmt.Printf("eventMsg: %v\n", eventMsg)
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

// TODO see if can reproduce https://github.com/gin-gonic/gin/issues/3142
// some bugs where fixed in sse example committed 4 months later, so...
func (stream *Event) serveHTTP() gin.HandlerFunc {
	fmt.Println("serveHTTP - initializing stream event")
	return func(c *gin.Context) {
		// Initialize client channel
		fmt.Println("serveHTTP - Initialize client channel")
		clientChan := make(ClientChan)

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

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
