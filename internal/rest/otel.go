package rest

import (
	"bytes"
	"io"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"

func userIDAttribute(c *gin.Context) attribute.KeyValue {
	uid := ""
	if u, err := GetUserCallerFromCtx(c); err != nil && u.User != nil {
		uid = u.UserID.String()
	}

	return tracing.UserIDAttributeKey.String(uid)
}

func newOTelSpan(opts ...trace.SpanStartOption) *tracing.OTelSpanBuilder {
	builder := tracing.NewOTelSpanBuilder(opts...).
		WithName(tracing.GetOTelSpanName(2)).
		WithAttributes(attribute.String("build-version", internal.Config.BuildVersion)).
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

func addRequestBodyToSpan(c *gin.Context) {
	span := GetSpanFromCtx(c)

	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		renderErrorResponse(c, "Failed to read request body", err)
	}
	span.SetAttributes(tracing.MetadataAttribute(jsonBody))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBody))
}
