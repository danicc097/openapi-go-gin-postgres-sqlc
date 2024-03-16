/**
* entr -r bash -c 'clear; go run cmd/sse-test/main.go' <<< cmd/sse-test/main.go
*
  curl -X 'GET' -N 'http://localhost:8085/stream?topics=AnotherTopic&topics=Time'
*
*/

package main

import (
	"context"
	"encoding/json"
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

type Topic string

const (
	TopicsJSONData     Topic = "JSONData"
	TopicsTime         Topic = "Time"
	TopicsTopic1       Topic = "Topic1"
	TopicsAnotherTopic Topic = "AnotherTopic"
)

func AllTopicsValues() []Topic {
	return []Topic{
		TopicsTopic1,
		TopicsAnotherTopic,
		TopicsTime,
		TopicsJSONData,
	}
}

type Clients map[chan ClientMessage]struct{}

type EventServer struct {
	// Events are pushed to this channel by the main events-gathering routine
	Messages      chan ClientMessage
	NewClients    chan Client
	ClosedClients chan chan ClientMessage
	Subscriptions map[Topic]Clients
	// Clients represents all connected clients
	Clients Clients
}

func newSSEServer() *EventServer {
	es := &EventServer{
		Messages:      make(chan ClientMessage),
		NewClients:    make(chan Client),
		ClosedClients: make(chan chan ClientMessage),
		Subscriptions: make(map[Topic]Clients),
		Clients:       make(Clients),
	}

	go es.listen()

	return es
}

type Client struct {
	Chan   chan ClientMessage
	Topics []Topic
}

type ClientMessage struct {
	Message string
	Topic   Topic
}

func main() {
	router := gin.Default()

	handlers := NewHandlers()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)
			handlers.event.Messages <- ClientMessage{Message: currentTime, Topic: TopicsTime}
		}
	}()

	go func() {
		for {
			data := struct {
				Field1 string `json:"field1"`
				Field2 int    `json:"field2"`
			}{
				Field1: "value1",
				Field2: 42,
			}

			msgData, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshaling JSON: %v", err)
				continue
			}

			handlers.event.Messages <- ClientMessage{Message: string(msgData), Topic: TopicsJSONData}

			time.Sleep(time.Second * 1)
		}
	}()

	router.GET("/stream", HeadersMiddleware(), handlers.event.serveHTTP(), func(c *gin.Context) {
		v, ok := c.Get("clientChan")
		if !ok {
			return
		}
		clientChan, ok := v.(Client)
		if !ok {
			return
		}
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
			}
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

func (server *EventServer) listen() {
	for {
		select {
		case client := <-server.NewClients:
			server.Clients[client.Chan] = struct{}{}
			for _, topic := range client.Topics {
				if _, ok := server.Subscriptions[topic]; !ok {
					server.Subscriptions[topic] = make(Clients)
				}
				server.Subscriptions[topic][client.Chan] = struct{}{}
			}
			log.Printf("Client added. %d registered clients", len(server.Clients))

		case client := <-server.ClosedClients:
			for _, subscriptions := range server.Subscriptions {
				delete(subscriptions, client)
			}
			delete(server.Clients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(server.Clients))

		case eventMsg := <-server.Messages:
			for ch := range server.Subscriptions[eventMsg.Topic] {
				ch <- eventMsg
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

		clientChan := Client{
			Chan:   make(chan ClientMessage, 1),
			Topics: make([]Topic, len(topics)),
		}
		for i, topic := range topics {
			clientChan.Topics[i] = Topic(topic)
		}

		server.NewClients <- clientChan
		defer func() {
			server.ClosedClients <- clientChan.Chan
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
