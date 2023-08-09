package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"

// When creating a Span it is recommended to provide all known span attributes
// using the `WithAttributes()` SpanOption as samplers will only have access
// to the attributes provided when a Span is created.
// Span name defaults to relative path to current function if not provided.
func newOTELSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) trace.Span {
	sn := name
	if sn == "" {
		sn = tracing.GetOTELSpanName(2)
	}
	_, span := otel.Tracer(OtelName).Start(ctx, sn, opts...)

	return span
}

func userIDAttribute(c *gin.Context) attribute.KeyValue {
	uid := ""
	if u := getUserFromCtx(c); u != nil {
		uid = u.UserID.String()
	}

	return tracing.UserIDAttribute.String(uid)
}

// newOTELSpanWithUser creates a new OTEL span with the current user included as attribute.
// Should be called in all handlers.
func newOTELSpanWithUser(c *gin.Context, opts ...trace.SpanStartOption) trace.Span {
	opts = append(opts, trace.WithAttributes(userIDAttribute(c)))

	return newOTELSpan(c.Request.Context(), tracing.GetOTELSpanName(2), opts...)
}
