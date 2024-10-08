// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// WorkItemTypeWithTracing implements _sourceRepos.WorkItemType interface instrumented with opentracing spans
type WorkItemTypeWithTracing struct {
	_sourceRepos.WorkItemType
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewWorkItemTypeWithTracing returns WorkItemTypeWithTracing
func NewWorkItemTypeWithTracing(base _sourceRepos.WorkItemType, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) WorkItemTypeWithTracing {
	d := WorkItemTypeWithTracing{
		WorkItemType: base,
		_instance:    instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements _sourceRepos.WorkItemType
func (_d WorkItemTypeWithTracing) ByID(ctx context.Context, d models.DBTX, id models.WorkItemTypeID, opts ...models.WorkItemTypeSelectConfigOption) (wp1 *models.WorkItemType, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.WorkItemType.ByID")
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
	return _d.WorkItemType.ByID(ctx, d, id, opts...)
}

// ByName implements _sourceRepos.WorkItemType
func (_d WorkItemTypeWithTracing) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTypeSelectConfigOption) (wp1 *models.WorkItemType, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.WorkItemType.ByName")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"name":      name,
				"projectID": projectID,
				"opts":      opts}, map[string]interface{}{
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
	return _d.WorkItemType.ByName(ctx, d, name, projectID, opts...)
}
