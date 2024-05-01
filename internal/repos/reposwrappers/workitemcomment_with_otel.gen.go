// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// WorkItemCommentWithTracing implements repos.WorkItemComment interface instrumented with opentracing spans
type WorkItemCommentWithTracing struct {
	repos.WorkItemComment
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewWorkItemCommentWithTracing returns WorkItemCommentWithTracing
func NewWorkItemCommentWithTracing(base repos.WorkItemComment, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) WorkItemCommentWithTracing {
	d := WorkItemCommentWithTracing{
		WorkItemComment: base,
		_instance:       instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.WorkItemComment
func (_d WorkItemCommentWithTracing) ByID(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, opts ...db.WorkItemCommentSelectConfigOption) (wp1 *db.WorkItemComment, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemComment.ByID")
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
	return _d.WorkItemComment.ByID(ctx, d, id, opts...)
}

// Create implements repos.WorkItemComment
func (_d WorkItemCommentWithTracing) Create(ctx context.Context, d db.DBTX, params *db.WorkItemCommentCreateParams) (wp1 *db.WorkItemComment, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemComment.Create")
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
	return _d.WorkItemComment.Create(ctx, d, params)
}

// Delete implements repos.WorkItemComment
func (_d WorkItemCommentWithTracing) Delete(ctx context.Context, d db.DBTX, id db.WorkItemCommentID) (wp1 *db.WorkItemComment, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemComment.Delete")
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
	return _d.WorkItemComment.Delete(ctx, d, id)
}

// Update implements repos.WorkItemComment
func (_d WorkItemCommentWithTracing) Update(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, params *db.WorkItemCommentUpdateParams) (wp1 *db.WorkItemComment, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemComment.Update")
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
	return _d.WorkItemComment.Update(ctx, d, id, params)
}
