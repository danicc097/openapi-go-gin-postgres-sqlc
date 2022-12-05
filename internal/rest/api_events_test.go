package rest

import (
	"net/http/httptest"
	"testing"
)

// TODO see e.g. https://dev.lucaskatayama.com/posts/go/2020/08/sse-with-gin/

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

func (s *StreamRecorder) closeClient() {
	s.closeNotify <- true
}

func NewStreamRecorder() *StreamRecorder {
	return &StreamRecorder{
		httptest.NewRecorder(),
		make(chan bool, 2),
	}
}

// TODO closenotifier deprecated. Use ctx: https://stackoverflow.com/questions/32123546/eventsource-golang-how-to-detect-client-disconnection
// TODO revisit when multiple events (personal notif., global notif., etc.) are implemented.
// would need a way to stop streaming after N messages, etc.
// func TestSSEStream(t *testing.T) {
// 	res := NewStreamRecorder()
// 	req := httptest.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/events", nil)

// 	srv, err := runTestServer(t, testpool, []gin.HandlerFunc{func(c *gin.Context) {
// 		c.Next()
// 	}})
// 	if err != nil {
// 		t.Fatalf("Couldn't run test server: %s\n", err)
// 	}
// 	defer srv.Close()

// 	go srv.Handler.ServeHTTP(res, req)

// 	// do notifications here

// 	time.Sleep(3000 * time.Millisecond)
// 	// 	for !res.Flushed {
// 	// }
// 	// body, _ := io.ReadAll(res.Body) // will just read one event
// 	// fmt.Printf("body: %v\n", body)

// 	body := res.Body.String()
// 	assert.NotEmpty(t, body)
// 	fmt.Printf("body: %v\n", body)

// 	res.closeClient()
// 	assert.Contains(t, strings.ReplaceAll("event:message\ndata:", " ", ""), strings.ReplaceAll(body, " ", ""))
// 	assert.Contains(t, "The Current Time Is", strings.ReplaceAll(body, " ", ""))
// }
