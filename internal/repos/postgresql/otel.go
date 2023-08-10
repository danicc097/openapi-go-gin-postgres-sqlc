package postgresql

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"

func newOTelSpan(opts ...trace.SpanStartOption) *tracing.OTelSpanBuilder {
	builder := tracing.NewOTelSpanBuilder(opts...).
		WithName(tracing.GetOTelSpanName(2)).
		WithTracer(OtelName).
		WithAttributes(semconv.DBSystemPostgreSQL)

	return builder
}
