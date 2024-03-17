//go:build !race

package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO see e.g. https://dev.lucaskatayama.com/posts/go/2020/08/sse-with-gin/
// also see sse test: https://github.com/prysmaticlabs/prysm/blob/f7cecf9f8a6dca95e021bab2fc30dd7c6d6ce760/beacon-chain/rpc/apimiddleware/custom_handlers_test.go#LL250C10-L250C10
// complete implementation and tests: https://github.com/MarinX/go/blob/06804256ef814c8381e3f54b1c89a8c95cabb300/services/horizon/internal/render/sse/main.go
func TestHandlers_Events(t *testing.T) {
	// res := httptest.NewRecorder()
	// _, engine := gin.CreateTestContext(res)

	// req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// engine.Use(SSEHeadersMiddleware(), SSEServerMiddleware())

	// engine.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "ok")
	// })
	// engine.ServeHTTP(res, req)

	// assert.Equal(t, tc.status, res.Code)
	// assert.Contains(t, res.Body.String(), tc.body)
}

type StreamRecorder struct {
	*httptest.ResponseRecorder
	closeNotify chan bool
}

func (s *StreamRecorder) CloseNotify() <-chan bool {
	return s.closeNotify
}

// deprecated in favor of context
// func (s *StreamRecorder) closeClient() {
// 	s.closeNotify <- true
// }

func NewStreamRecorder() *StreamRecorder {
	return &StreamRecorder{
		httptest.NewRecorder(),
		make(chan bool, 2),
	}
}

// TODO revisit when multiple events (personal notif., global notif., etc.) are implemented.
// / see https://github.com/search?q=%22req.Context%28%29.Done%28%29%22+sse&type=code
// would need a way to stop streaming after N messages, etc.
func TestSSEStream(t *testing.T) {
	t.Parallel()

	res := NewStreamRecorder()
	req := httptest.NewRequest(http.MethodGet, MustConstructInternalPath("/events", WithQueryParams(models.EventsParams{ProjectName: models.ProjectDemo, Topics: []models.Topic{models.TopicGlobalAlerts}})), nil)

	ctx, cancel := context.WithCancel(context.Background())
	req = req.WithContext(ctx)

	srv, err := runTestServer(t, testPool,
		func(c *gin.Context) {
			c.Next()
		},
	)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.setupCleanup(t)

	// TODO: should have generated option to use a custom recorder, e.g.
	// here must use stream recorder that implements closenotifier
	// resp, err := srv.client.Events(ctx, &rest.EventsParams{Topics: []models.Topic{models.TopicGlobalAlerts}, ProjectName: models.ProjectDemo})
	// require.NoError(t, err)
	// bd, err := io.ReadAll(resp.Body)
	// require.NoError(t, err)
	// fmt.Printf("bd: %v\n", bd)

	msg := "test-message-123"
	stopCh := make(chan bool)

	go func() {
		defer func() {
			stopCh <- true
		}()
		for {
			select {
			case <-stopCh:
				return
			default:
				srv.event.Publish(msg, models.TopicGlobalAlerts)
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				srv.server.Handler.ServeHTTP(res, req)
			}
		}
	}()

	// TODO trigger events

	// TODO all sse events tests should be done alongside handler tests that trigger them.
	// could have generic test helpers as well.
	// in this file we should just unit test with a random event
	if !assert.Eventually(t, func() bool {
		if res.Body == nil {
			return false
		}
		body := res.Body.String()
		return strings.Count(body, "event:"+string(models.TopicGlobalAlerts)) >= 1 && strings.Count(body, "data:"+msg) >= 1
	}, 10*time.Second, 100*time.Millisecond) {
		t.Fatalf("did not receive event")
	}

	cancel()
	// handler should be stopped before reading body snapshot. to not have an arbitrary time sleep
	// after events are sent before shutting handler down we're using Eventually and excluding -race flag.
	stopCh <- true

	assert.Contains(t, res.Result().Header.Get("Content-Type"), "text/event-stream")
	assert.Contains(t, res.Result().Header.Get("Cache-Control"), "no-cache")
	assert.Contains(t, res.Result().Header.Get("Connection"), "keep-alive")
	assert.Contains(t, res.Result().Header.Get("Transfer-Encoding"), "chunked")
}
