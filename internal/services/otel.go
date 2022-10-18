package services

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const otelName = "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"

func newOTELSpan(ctx context.Context, name string) trace.Span {
	_, span := otel.Tracer(otelName).Start(ctx, name)

	return span
}
