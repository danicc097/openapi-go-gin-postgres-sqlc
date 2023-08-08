package services

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"

// When creating a Span it is recommended to provide all known span attributes
// using the `WithAttributes()` SpanOption as samplers will only have access
// to the attributes provided when a Span is created.
func newOTELSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) trace.Span {
	_, span := otel.Tracer(OtelName).Start(ctx, "services."+name, opts...)

	return span
}
