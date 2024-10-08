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

// UserWithTracing implements _sourceRepos.User interface instrumented with opentracing spans
type UserWithTracing struct {
	_sourceRepos.User
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewUserWithTracing returns UserWithTracing
func NewUserWithTracing(base _sourceRepos.User, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) UserWithTracing {
	d := UserWithTracing{
		User:      base,
		_instance: instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByAPIKey implements _sourceRepos.User
func (_d UserWithTracing) ByAPIKey(ctx context.Context, d models.DBTX, apiKey string) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByAPIKey")
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

// ByEmail implements _sourceRepos.User
func (_d UserWithTracing) ByEmail(ctx context.Context, d models.DBTX, email string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByEmail")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"d":     d,
				"email": email,
				"opts":  opts}, map[string]interface{}{
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
	return _d.User.ByEmail(ctx, d, email, opts...)
}

// ByExternalID implements _sourceRepos.User
func (_d UserWithTracing) ByExternalID(ctx context.Context, d models.DBTX, extID string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByExternalID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"d":     d,
				"extID": extID,
				"opts":  opts}, map[string]interface{}{
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
	return _d.User.ByExternalID(ctx, d, extID, opts...)
}

// ByID implements _sourceRepos.User
func (_d UserWithTracing) ByID(ctx context.Context, d models.DBTX, id models.UserID, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"id":   id,
				"opts": opts}, map[string]interface{}{
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
	return _d.User.ByID(ctx, d, id, opts...)
}

// ByProject implements _sourceRepos.User
func (_d UserWithTracing) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID) (ua1 []models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByProject")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"projectID": projectID}, map[string]interface{}{
				"ua1": ua1,
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
	return _d.User.ByProject(ctx, d, projectID)
}

// ByTeam implements _sourceRepos.User
func (_d UserWithTracing) ByTeam(ctx context.Context, d models.DBTX, teamID models.TeamID) (ua1 []models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByTeam")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"teamID": teamID}, map[string]interface{}{
				"ua1": ua1,
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
	return _d.User.ByTeam(ctx, d, teamID)
}

// ByUsername implements _sourceRepos.User
func (_d UserWithTracing) ByUsername(ctx context.Context, d models.DBTX, username string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.ByUsername")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"d":        d,
				"username": username,
				"opts":     opts}, map[string]interface{}{
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
	return _d.User.ByUsername(ctx, d, username, opts...)
}

// Create implements _sourceRepos.User
func (_d UserWithTracing) Create(ctx context.Context, d models.DBTX, params *models.UserCreateParams) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.Create")
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

// CreateAPIKey implements _sourceRepos.User
func (_d UserWithTracing) CreateAPIKey(ctx context.Context, d models.DBTX, user *models.User) (up1 *models.UserAPIKey, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.CreateAPIKey")
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

// Delete implements _sourceRepos.User
func (_d UserWithTracing) Delete(ctx context.Context, d models.DBTX, id models.UserID) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.Delete")
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

// DeleteAPIKey implements _sourceRepos.User
func (_d UserWithTracing) DeleteAPIKey(ctx context.Context, d models.DBTX, apiKey string) (up1 *models.UserAPIKey, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.DeleteAPIKey")
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
	return _d.User.DeleteAPIKey(ctx, d, apiKey)
}

// Paginated implements _sourceRepos.User
func (_d UserWithTracing) Paginated(ctx context.Context, d models.DBTX, params _sourceRepos.GetPaginatedUsersParams) (ua1 []models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.Paginated")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"ua1": ua1,
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
	return _d.User.Paginated(ctx, d, params)
}

// Update implements _sourceRepos.User
func (_d UserWithTracing) Update(ctx context.Context, d models.DBTX, id models.UserID, params *models.UserUpdateParams) (up1 *models.User, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "_sourceRepos.User.Update")
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
