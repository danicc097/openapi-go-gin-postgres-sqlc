package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTracingMiddleware(t *testing.T) {
	t.Parallel()

	resp := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(resp)
	tmw := newTracingMiddleware()

	engine.Use(tmw.WithSpan())
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")

		span := GetSpanFromCtx(c)
		require.NotNil(t, span)
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	engine.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
