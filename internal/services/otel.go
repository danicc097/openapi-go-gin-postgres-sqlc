package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"

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

	span.SetAttributes(semconv.DBSystemPostgreSQL)

	return span
}
