package tracing

import (
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	jaegerp "go.opentelemetry.io/contrib/propagators/jaeger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// https://www.timescale.com/blog/using-postgresql-as-a-scalable-durable-and-reliable-storage-for-jaeger-tracing/

func InitOTelTracer() *sdktrace.TracerProvider {
	jaegerEndpoint := "http://localhost:14268/api/traces"

	jaegerExporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)),
	)
	if err != nil {
		log.Fatalln("Couldn't initialize jaeger exporter", err)
	}

	stdoutExporter, err := stdouttrace.New()
	if err != nil {
		log.Fatalln("Couldn't initialize stdout exporter", err)
	}

	tracerOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(jaegerExporter),
		sdktrace.WithResource(resource.NewSchemaless(attribute.KeyValue{
			Key:   semconv.ServiceNameKey,
			Value: attribute.StringValue("openapi-go-server"),
		})),
	}

	if internal.Config.AppEnv == "prod" {
		tracerOptions = append(tracerOptions, sdktrace.WithBatcher(stdoutExporter))
	}

	tp := sdktrace.NewTracerProvider(
		tracerOptions...,
	)
	otel.SetTracerProvider(tp)

	p := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
		jaegerp.Jaeger{},
	)
	otel.SetTextMapPropagator(p)

	return tp
}
