package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OTelSpanBuilder struct {
	tracer      string
	suffix      string
	name        string
	spanOptions []trace.SpanStartOption
	attributes  []attribute.KeyValue
}

// TODO: reusable in tracing pkg, + WithTracer + WithAttributes so we have dedicated newOTelSpan for each package that returns
// builder with predefined opts.
// e.g. repos WithTracer(otelName).WithAttributes(semconv.DBSystemPostgreSQL)
// When creating a Span it is recommended to provide all known span attributes
// using the `WithAttributes()` SpanOption as samplers will only have access
// to the attributes provided when a Span is created.
func NewOTelSpanBuilder(opts ...trace.SpanStartOption) *OTelSpanBuilder {
	return &OTelSpanBuilder{spanOptions: opts}
}

// WithSuffix appends a suffix between brackets to the span name in
// order to identify multiple spans in the same function.
func (b *OTelSpanBuilder) WithSuffix(suffix string) *OTelSpanBuilder {
	b.suffix = suffix

	return b
}

// WithTracer overrides the default tracer name.
func (b *OTelSpanBuilder) WithTracer(tracer string) *OTelSpanBuilder {
	b.tracer = tracer

	return b
}

// WithName overrides the default span name.
func (b *OTelSpanBuilder) WithName(name string) *OTelSpanBuilder {
	b.name = name

	return b
}

// WithAttributes adds span attributes.
func (b *OTelSpanBuilder) WithAttributes(kv ...attribute.KeyValue) *OTelSpanBuilder {
	b.attributes = kv

	return b
}

// Build builds and returns a span with the given options.
func (b *OTelSpanBuilder) Build(ctx context.Context) trace.Span {
	if b.suffix != "" {
		b.suffix = "[" + b.suffix + "]"
	}

	if b.name == "" {
		b.name = GetOTelSpanName(2)
	}

	if b.tracer == "" {
		b.tracer = "openapi-go-server"
	}

	_, span := otel.Tracer(b.tracer).Start(
		ctx,
		b.name+b.suffix,
		b.spanOptions...,
	)

	span.SetAttributes(b.attributes...)

	return span
}
