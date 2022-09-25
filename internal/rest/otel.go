package rest

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const otelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"

func newOTELSpan(ctx context.Context, tp *sdktrace.TracerProvider, name string) trace.Span {
	_, span := tp.Tracer(otelName).Start(ctx, name)

	return span
}
