package rest

import (
	"github.com/gin-gonic/gin"
)

type tracingMiddleware struct{}

func newTracingMiddleware() *tracingMiddleware {
	return &tracingMiddleware{}
}

// WithSpan creates a span in context.
func (m *tracingMiddleware) WithSpan() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := newOTelSpan().Build(c.Request.Context())
		defer span.End()

		span.SetAttributes(userIDAttribute(c)) // if we are authenticated, it sets user-id

		ctxWithSpan(c, span)

		c.Next()
	}
}
