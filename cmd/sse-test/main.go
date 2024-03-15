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

type Topics string

const (
	TopicsTopic1       Topics = "Topic1"
	TopicsAnotherTopic Topics = "AnotherTopic"
)

func AllTopicsValues() []Topics {
	return []Topics{
		TopicsTopic1,
		TopicsAnotherTopic,
	}
}

type EventServer struct {
	// Events are pushed to this channel by the main events-gathering routine
	Messages      chan ClientMessage
	NewClients    chan ClientChan
	ClosedClients chan chan string
	Subscriptions map[chan string]map[Topics]struct{} // Map of client channel to subscribed topics
}

type ClientChan struct {
	Chan   chan string
	Topics []Topics
}

type ClientMessage struct {
	Message string
	Topic   Topics
}

func main() {
	router := gin.Default()

	handlers := NewHandlers()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)
			handlers.event.Messages <- ClientMessage{Message: currentTime, Topic: TopicsTopic1}
		}
	}()

	router.GET("/stream", HeadersMiddleware(), handlers.event.serveHTTP(), func(c *gin.Context) {
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
			if msg, ok := <-clientChan.Chan; ok {
				c.SSEvent(string(clientChan.Topics[0]), msg)
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

type Handlers struct {
	event *EventServer
}

func NewHandlers() *Handlers {
	event := newSSEServer()

	return &Handlers{
		event: event,
	}
}

func Run(router *gin.Engine) (<-chan error, error) {
	addr := ":8085"
	httpsrv := &http.Server{
		Handler: router,
		Addr:    addr,
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

		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		httpsrv.SetKeepAlivesEnabled(false)

		if err := httpsrv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		log.Printf("Shutdown completed")
	}()

	go func() {
		log.Printf("Listening and serving on http://localhost" + addr)

		var err error

		err = httpsrv.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

func newSSEServer() *EventServer {
	es := &EventServer{
		Messages:      make(chan ClientMessage),
		NewClients:    make(chan ClientChan),
		ClosedClients: make(chan chan string),
		Subscriptions: make(map[chan string]map[Topics]struct{}),
	}

	go es.listen()

	return es
}

func (server *EventServer) listen() {
	for {
		select {
		case client := <-server.NewClients:
			server.Subscriptions[client.Chan] = make(map[Topics]struct{})
			for _, topic := range client.Topics {
				server.Subscriptions[client.Chan][topic] = struct{}{}
			}
			log.Printf("Client added. %d registered clients", len(server.Subscriptions))

		case client := <-server.ClosedClients:
			if _, ok := server.Subscriptions[client]; ok {
				delete(server.Subscriptions, client)
				close(client)
				log.Printf("Removed client. %d registered clients", len(server.Subscriptions))
			} else {
				log.Printf("Client already removed. %d registered clients", len(server.Subscriptions))
			}

		case eventMsg := <-server.Messages:
			for client, topics := range server.Subscriptions {
				// Only send if subscribed to topic
				if _, ok := topics[eventMsg.Topic]; ok {
					client <- eventMsg.Message
				}
			}
		}
	}
}

func (server *EventServer) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		topics := c.QueryArray("topics")
		if len(topics) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "topics query parameter is required"})
			return
		}

		clientChan := ClientChan{Chan: make(chan string, 1), Topics: make([]Topics, len(topics))}
		for i, topic := range topics {
			// Check if the topic exists
			validTopic := false
			for _, t := range AllTopicsValues() {
				if Topics(topic) == t {
					validTopic = true
					break
				}
			}
			if !validTopic {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid topic: %s", topic)})
				return
			}
			clientChan.Topics[i] = Topics(topic)
		}

		server.NewClients <- clientChan

		c.Set("clientChan", clientChan)

		// Handle client disconnection
		go func() {
			<-c.Writer.CloseNotify()
			server.ClosedClients <- clientChan.Chan
		}()

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
