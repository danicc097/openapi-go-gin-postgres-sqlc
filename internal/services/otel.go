package services

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const OtelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"

func newOTELSpan(ctx context.Context, name string) trace.Span {
	_, span := otel.Tracer(OtelName).Start(ctx, name)

	return span
}
