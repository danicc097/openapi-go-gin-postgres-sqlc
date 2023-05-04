// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// UserWithTracing implements repos.User interface instrumented with opentracing spans
type UserWithTracing struct {
	repos.User
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewUserWithTracing returns UserWithTracing
func NewUserWithTracing(base repos.User, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) UserWithTracing {
	d := UserWithTracing{
		User:      base,
		_instance: instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByAPIKey implements repos.User
func (_d UserWithTracing) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.ByAPIKey")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"apiKey": apiKey}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.ByAPIKey(ctx, d, apiKey)
}

// ByEmail implements repos.User
func (_d UserWithTracing) ByEmail(ctx context.Context, d db.DBTX, email string) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.ByEmail")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"d":     d,
				"email": email}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.ByEmail(ctx, d, email)
}

// ByExternalID implements repos.User
func (_d UserWithTracing) ByExternalID(ctx context.Context, d db.DBTX, extID string) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.ByExternalID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"d":     d,
				"extID": extID}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.ByExternalID(ctx, d, extID)
}

// ByID implements repos.User
func (_d UserWithTracing) ByID(ctx context.Context, d db.DBTX, id uuid.UUID) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.ByID(ctx, d, id)
}

// ByUsername implements repos.User
func (_d UserWithTracing) ByUsername(ctx context.Context, d db.DBTX, username string) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.ByUsername")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"d":        d,
				"username": username}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.ByUsername(ctx, d, username)
}

// Create implements repos.User
func (_d UserWithTracing) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.Create(ctx, d, params)
}

// CreateAPIKey implements repos.User
func (_d UserWithTracing) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (up1 *db.UserAPIKey, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.CreateAPIKey")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"user": user}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.CreateAPIKey(ctx, d, user)
}

// Delete implements repos.User
func (_d UserWithTracing) Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.Delete(ctx, d, id)
}

// Update implements repos.User
func (_d UserWithTracing) Update(ctx context.Context, d db.DBTX, id uuid.UUID, params *db.UserUpdateParams) (up1 *db.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.User.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"id":     id,
				"params": params}, map[string]interface{}{
				"up1": up1,
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
	return _d.User.Update(ctx, d, id, params)
}
