package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"

func userIDAttribute(c *gin.Context) attribute.KeyValue {
	uid := ""
	if u := getUserFromCtx(c); u != nil {
		uid = uuid.UUID(u.UserID).String()
	}

	return tracing.UserIDAttribute.String(uid)
}

func newOTelSpan(opts ...trace.SpanStartOption) *tracing.OTelSpanBuilder {
	builder := tracing.NewOTelSpanBuilder(opts...).
		WithName(tracing.GetOTelSpanName(2)).
		WithTracer(OtelName)

	return builder
}

// newOTelSpanWithUser creates a new OTel span with the current user included as attribute.
// Should be called in all handlers that require authentication.
func newOTelSpanWithUser(c *gin.Context, opts ...trace.SpanStartOption) trace.Span {
	opts = append(opts, trace.WithAttributes(userIDAttribute(c)))

	builder := newOTelSpan(opts...).
		WithName(tracing.GetOTelSpanName(2))

	return builder.Build(c.Request.Context())
}
