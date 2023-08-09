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
// Span name records a relative path to the current function and an optional suffix
// to identify multiple spans in the same function.
func newOTelSpan(ctx context.Context, suffix string, opts ...trace.SpanStartOption) trace.Span {
	if suffix != "" {
		suffix = "[" + suffix + "]"
	}
	_, span := otel.Tracer(OtelName).Start(ctx, tracing.GetOTelSpanName(2)+suffix, opts...)

	return span
}

func userIDAttribute(c *gin.Context) attribute.KeyValue {
	uid := ""
	if u := getUserFromCtx(c); u != nil {
		uid = u.UserID.String()
	}

	return tracing.UserIDAttribute.String(uid)
}

// newOTelSpanWithUser creates a new OTel span with the current user included as attribute.
// Should be called in all handlers.
func newOTelSpanWithUser(c *gin.Context, opts ...trace.SpanStartOption) trace.Span {
	opts = append(opts, trace.WithAttributes(userIDAttribute(c)))

	_, span := otel.Tracer(OtelName).Start(c.Request.Context(), tracing.GetOTelSpanName(2), opts...)

	return span
}
