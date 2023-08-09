package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"

func newOTelSpan(opts ...trace.SpanStartOption) *tracing.OTelSpanBuilder {
	builder := tracing.NewOTelSpanBuilder(opts...).
		WithName(tracing.GetOTelSpanName(2)).
		WithTracer(OtelName)

	return builder
}
