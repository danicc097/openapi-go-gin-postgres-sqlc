package tracing

import (
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	jaegerp "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func InitTracer() *sdktrace.TracerProvider {
	jaegerEndpoint := "http://localhost:14268/api/traces"

	jaegerExporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)),
	)
	if err != nil {
		log.Fatalln("Couldn't initialize exporter", err)
	}

	// Create stdout exporter to be able to retrieve the collected spans.
	_, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalln("Couldn't initialize exporter", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(jaegerExporter),
		sdktrace.WithResource(resource.NewSchemaless(attribute.KeyValue{
			Key:   semconv.ServiceNameKey,
			Value: attribute.StringValue(envvar.GetEnv("PROJECT_PREFIX", "project-prefix")),
		})),
	)
	otel.SetTracerProvider(tp)
	p := jaegerp.Jaeger{}
	otel.SetTextMapPropagator(p)

	return tp
}
