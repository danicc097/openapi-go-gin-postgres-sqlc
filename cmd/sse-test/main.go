package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

// go run cmd/sse-test/main.go
// curl -X 'GET' -N  'http://localhost:8085/stream'
func main() {
	router := gin.Default()

	// Initialize new streaming server
	stream := NewServer()

	// We are streaming current time to clients in the interval 10 seconds
	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)
			fmt.Printf("currentTime: %v\n", currentTime)
			// Send current time to clients message channel
			stream.Message <- currentTime
		}
	}()

	router.GET("/stream", HeadersMiddleware(), stream.serveHTTP(), func(c *gin.Context) {
		v, ok := c.Get("clientChan")
		if !ok {
			return
		}
		clientChan, ok := v.(ClientChan)
		if !ok {
			return
		}
		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if msg, ok := <-clientChan; ok {
				c.SSEvent("message", msg)
				return true
			}
			c.SSEvent("message", "STOPPED")
			return false
		})
	})

	errC, err := Run(router)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func Run(router *gin.Engine) (<-chan error, error) {
	// var err error

	httpsrv := &http.Server{
		Handler: router,
		Addr:    ":8085",
		// ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		// WriteTimeout:      1 * time.Second,
		// IdleTimeout: 1 * time.Second,
	}

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		log.Print("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		// any action on shutdown must be deferred here and not in the main block
		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		httpsrv.SetKeepAlivesEnabled(false)

		if err := httpsrv.Shutdown(ctxTimeout); err != nil { //nolint: contextcheck
			errC <- err
		}

		log.Printf("Shutdown completed")
	}()

	go func() {
		log.Printf("Listening and serving")

		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		var err error

		err = httpsrv.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

// Initialize event and Start procnteessing requests
func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go event.listen()

	return
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
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
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

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
