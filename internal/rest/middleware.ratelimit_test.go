package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestRateLimitMiddleware(t *testing.T) {
	t.Parallel()

	resp := httptest.NewRecorder()
	logger := testutil.NewLogger(t)
	_, engine := gin.CreateTestContext(resp)
	rl := 1
	bl := 3
	rlmw := newRateLimitMiddleware(logger, rate.Limit(rl), bl)

	engine.Use(rlmw.Limit())
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	for i := 0; i < bl; i++ {
		engine.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	}

	resp = httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusTooManyRequests, resp.Code)
}
