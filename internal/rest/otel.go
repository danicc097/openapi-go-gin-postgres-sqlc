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
func newOTELSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) trace.Span {
	_, span := otel.Tracer(OtelName).Start(ctx, name, opts...)

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
func newOTELSpanWithUser(c *gin.Context, name string, opts ...trace.SpanStartOption) trace.Span {
	opts = append(opts, trace.WithAttributes(userIDAttribute(c)))

	return newOTELSpan(c.Request.Context(), "handlers."+name, opts...)
}
