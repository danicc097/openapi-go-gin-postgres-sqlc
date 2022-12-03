package rest

import (
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
