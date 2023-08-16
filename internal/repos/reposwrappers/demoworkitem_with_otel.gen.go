// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// DemoWorkItemWithTracing implements repos.DemoWorkItem interface instrumented with opentracing spans
type DemoWorkItemWithTracing struct {
	repos.DemoWorkItem
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewDemoWorkItemWithTracing returns DemoWorkItemWithTracing
func NewDemoWorkItemWithTracing(base repos.DemoWorkItem, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) DemoWorkItemWithTracing {
	d := DemoWorkItemWithTracing{
		DemoWorkItem: base,
		_instance:    instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.DemoWorkItem
func (_d DemoWorkItemWithTracing) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoWorkItem.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"id":   id,
				"opts": opts}, map[string]interface{}{
				"wp1": wp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.DemoWorkItem.ByID(ctx, d, id, opts...)
}

// Create implements repos.DemoWorkItem
func (_d DemoWorkItemWithTracing) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoWorkItem.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"wp1": wp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.DemoWorkItem.Create(ctx, d, params)
}

// Update implements repos.DemoWorkItem
func (_d DemoWorkItemWithTracing) Update(ctx context.Context, d db.DBTX, id db.WorkItemID, params repos.DemoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoWorkItem.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"id":     id,
				"params": params}, map[string]interface{}{
				"wp1": wp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.DemoWorkItem.Update(ctx, d, id, params)
}
