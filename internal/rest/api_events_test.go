package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO see e.g. https://dev.lucaskatayama.com/posts/go/2020/08/sse-with-gin/
// also see sse test: https://github.com/prysmaticlabs/prysm/blob/f7cecf9f8a6dca95e021bab2fc30dd7c6d6ce760/beacon-chain/rpc/apimiddleware/custom_handlers_test.go#LL250C10-L250C10
// complete implementation and tests: https://github.com/MarinX/go/blob/06804256ef814c8381e3f54b1c89a8c95cabb300/services/horizon/internal/render/sse/main.go
func TestHandlers_Events(t *testing.T) {
	// resp := httptest.NewRecorder()
	// _, engine := gin.CreateTestContext(resp)

	// req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// engine.Use(SSEHeadersMiddleware(), SSEServerMiddleware())

	// engine.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "ok")
	// })
	// engine.ServeHTTP(resp, req)

	// assert.Equal(t, tc.status, resp.Code)
	// assert.Contains(t, resp.Body.String(), tc.body)
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
	req := httptest.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/events", nil)

	ctx, cancel := context.WithCancel(context.Background())
	req = req.WithContext(ctx)

	srv, err := runTestServer(t, testpool, []gin.HandlerFunc{func(c *gin.Context) {
		c.Next()
	}})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	go srv.Handler.ServeHTTP(res, req)

	// trigger events
	// TODO in reality call whatever triggers the events, then wait a bit for them to
	//  be sent if its more than one since everyone will call flush
	// TODO test 2 clients concurrently receive, one leaves, the other still receives.
	// TODO while not (contains (...) && contains (...)) sleep until timeout
	// Or use assert.eventually
	time.Sleep(200 * time.Millisecond)

	for start := time.Now(); time.Since(start) < 2*time.Second; {
		for !res.Flushed {
			time.Sleep(100 * time.Millisecond)
		}
	}

	cancel()

	// FIXME data race reading body
	body := res.Body.String()

	assert.NotEmpty(t, body)
	assert.Contains(t, strings.ReplaceAll(body, " ", ""), "event:"+models.TopicsUserNotifications)
	assert.Contains(t, strings.ReplaceAll(body, " ", ""), "event:test-event")
}
