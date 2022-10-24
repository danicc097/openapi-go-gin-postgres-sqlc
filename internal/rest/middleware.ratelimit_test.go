package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

func TestRateLimitMiddleware(t *testing.T) {
	resp := httptest.NewRecorder()
	logger, _ := zap.NewDevelopment()
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
	}
	resp = httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusTooManyRequests, resp.Code)
}
