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

// ActivityWithTracing implements _sourceRepos.Activity interface instrumented with opentracing spans
type ActivityWithTracing struct {
	_sourceRepos.Activity
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewActivityWithTracing returns ActivityWithTracing
func NewActivityWithTracing(base _sourceRepos.Activity, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) ActivityWithTracing {
	d := ActivityWithTracing{
		Activity:  base,
		_instance: instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements _sourceRepos.Activity
func (_d ActivityWithTracing) ByID(ctx context.Context, d models.DBTX, id models.ActivityID, opts ...models.ActivitySelectConfigOption) (ap1 *models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"id":   id,
				"opts": opts}, map[string]interface{}{
				"ap1": ap1,
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
	return _d.Activity.ByID(ctx, d, id, opts...)
}

// ByName implements _sourceRepos.Activity
func (_d ActivityWithTracing) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) (ap1 *models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.ByName")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"name":      name,
				"projectID": projectID,
				"opts":      opts}, map[string]interface{}{
				"ap1": ap1,
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
	return _d.Activity.ByName(ctx, d, name, projectID, opts...)
}

// ByProjectID implements _sourceRepos.Activity
func (_d ActivityWithTracing) ByProjectID(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) (aa1 []models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.ByProjectID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"projectID": projectID,
				"opts":      opts}, map[string]interface{}{
				"aa1": aa1,
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
	return _d.Activity.ByProjectID(ctx, d, projectID, opts...)
}

// Create implements _sourceRepos.Activity
func (_d ActivityWithTracing) Create(ctx context.Context, d models.DBTX, params *models.ActivityCreateParams) (ap1 *models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"ap1": ap1,
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
	return _d.Activity.Create(ctx, d, params)
}

// Delete implements _sourceRepos.Activity
func (_d ActivityWithTracing) Delete(ctx context.Context, d models.DBTX, id models.ActivityID) (ap1 *models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"ap1": ap1,
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
	return _d.Activity.Delete(ctx, d, id)
}

// Restore implements _sourceRepos.Activity
func (_d ActivityWithTracing) Restore(ctx context.Context, d models.DBTX, id models.ActivityID) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.Restore")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
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
	return _d.Activity.Restore(ctx, d, id)
}

// Update implements _sourceRepos.Activity
func (_d ActivityWithTracing) Update(ctx context.Context, d models.DBTX, id models.ActivityID, params *models.ActivityUpdateParams) (ap1 *models.Activity, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.Activity.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"id":     id,
				"params": params}, map[string]interface{}{
				"ap1": ap1,
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
	return _d.Activity.Update(ctx, d, id, params)
}
