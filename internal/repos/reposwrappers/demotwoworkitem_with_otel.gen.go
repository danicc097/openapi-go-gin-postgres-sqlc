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

// DemoTwoWorkItemWithTracing implements repos.DemoTwoWorkItem interface instrumented with opentracing spans
type DemoTwoWorkItemWithTracing struct {
	repos.DemoTwoWorkItem
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewDemoTwoWorkItemWithTracing returns DemoTwoWorkItemWithTracing
func NewDemoTwoWorkItemWithTracing(base repos.DemoTwoWorkItem, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) DemoTwoWorkItemWithTracing {
	d := DemoTwoWorkItemWithTracing{
		DemoTwoWorkItem: base,
		_instance:       instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTracing) ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoTwoWorkItem.ByID")
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
	return _d.DemoTwoWorkItem.ByID(ctx, d, id, opts...)
}

// Create implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTracing) Create(ctx context.Context, d db.DBTX, params repos.DemoTwoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoTwoWorkItem.Create")
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
	return _d.DemoTwoWorkItem.Create(ctx, d, params)
}

// Delete implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTracing) Delete(ctx context.Context, d db.DBTX, id int64) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoTwoWorkItem.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
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
	return _d.DemoTwoWorkItem.Delete(ctx, d, id)
}

// Restore implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTracing) Restore(ctx context.Context, d db.DBTX, id int64) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoTwoWorkItem.Restore")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
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
	return _d.DemoTwoWorkItem.Restore(ctx, d, id)
}

// Update implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTracing) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoTwoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.DemoTwoWorkItem.Update")
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
	return _d.DemoTwoWorkItem.Update(ctx, d, id, params)
}
