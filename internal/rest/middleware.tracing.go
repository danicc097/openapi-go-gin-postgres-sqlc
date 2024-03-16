package rest

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// RequestIDMiddleware sets a unique X-Request-ID header
func (m *tracingMiddleware) RequestIDMiddleware(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := fmt.Sprintf("%s-%s", prefix, uuid.New())

		c.Writer.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(c.Request.Context(), requestIDCtxKey{}, requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
