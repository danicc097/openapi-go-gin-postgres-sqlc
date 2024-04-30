// reading headers with -race fails
//go:build !race

package rest_test

import (
	"context"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestSSEStream(t *testing.T) {
	t.Parallel()

	res := NewStreamRecorder()

	ctx, cancel := context.WithCancel(context.Background())

	srv, err := runTestServer(t, ctx, testPool,
		func(c *gin.Context) {
			c.Next()
		},
	)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.setupCleanup(t)

	publishMsg := "test-message-123"
	stopCh := make(chan bool)

	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				srv.event.Publish(publishMsg, models.TopicGlobalAlerts)
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				res, err := srv.client.Events(ctx, res, &rest.EventsParams{Topics: []models.Topic{models.TopicGlobalAlerts}, ProjectName: models.ProjectNameDemo})
				require.NoError(t, err)
				resb, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				t.Logf("response body: %v\n", string(resb))
			}
		}
	}()

	eventReceived := func() bool {
		if res.Body == nil {
			return false
		}
		body := res.Body.String()
		return strings.Count(body, "event:"+string(models.TopicGlobalAlerts)) >= 1 && strings.Count(body, "data:"+publishMsg) >= 1
	}

	// TODO all internal sse events tests should be done alongside handler tests that trigger them.
	// could have generic test helpers as well.
	// in this file we should just unit test with a random event, adhoc handlers...
	if !assert.Eventually(t, eventReceived, 5*time.Second, 200*time.Millisecond) {
		t.Fatalf("did not receive event")
	}

	cancel()
	stopCh <- true

	assert.Contains(t, res.Result().Header.Get("Content-Type"), "text/event-stream")
	assert.Contains(t, res.Result().Header.Get("Cache-Control"), "no-cache")
	assert.Contains(t, res.Result().Header.Get("Connection"), "keep-alive")
	assert.Contains(t, res.Result().Header.Get("Transfer-Encoding"), "chunked")
}
