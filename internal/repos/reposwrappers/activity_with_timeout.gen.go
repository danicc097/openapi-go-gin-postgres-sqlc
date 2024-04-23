// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// ActivityWithTimeout implements repos.Activity interface instrumented with timeouts
type ActivityWithTimeout struct {
	repos.Activity
	config ActivityWithTimeoutConfig
}

type ActivityWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByNameTimeout time.Duration

	ByProjectIDTimeout time.Duration

	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	RestoreTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewActivityWithTimeout returns ActivityWithTimeout
func NewActivityWithTimeout(base repos.Activity, config ActivityWithTimeoutConfig) ActivityWithTimeout {
	return ActivityWithTimeout{
		Activity: base,
		config:   config,
	}
}

// ByID implements repos.Activity
func (_d ActivityWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.ActivityID, opts ...db.ActivitySelectConfigOption) (ap1 *db.Activity, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.Activity.ByID(ctx, d, id, opts...)
}

// ByName implements repos.Activity
func (_d ActivityWithTimeout) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.ActivitySelectConfigOption) (ap1 *db.Activity, err error) {
	var cancelFunc func()
	if _d.config.ByNameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByNameTimeout)
		defer cancelFunc()
	}
	return _d.Activity.ByName(ctx, d, name, projectID, opts...)
}

// ByProjectID implements repos.Activity
func (_d ActivityWithTimeout) ByProjectID(ctx context.Context, d db.DBTX, projectID db.ProjectID, opts ...db.ActivitySelectConfigOption) (aa1 []db.Activity, err error) {
	var cancelFunc func()
	if _d.config.ByProjectIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByProjectIDTimeout)
		defer cancelFunc()
	}
	return _d.Activity.ByProjectID(ctx, d, projectID, opts...)
}

// Create implements repos.Activity
func (_d ActivityWithTimeout) Create(ctx context.Context, d db.DBTX, params *db.ActivityCreateParams) (ap1 *db.Activity, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.Activity.Create(ctx, d, params)
}

// Delete implements repos.Activity
func (_d ActivityWithTimeout) Delete(ctx context.Context, d db.DBTX, id db.ActivityID) (ap1 *db.Activity, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.Activity.Delete(ctx, d, id)
}

// Restore implements repos.Activity
func (_d ActivityWithTimeout) Restore(ctx context.Context, d db.DBTX, id db.ActivityID) (err error) {
	var cancelFunc func()
	if _d.config.RestoreTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.RestoreTimeout)
		defer cancelFunc()
	}
	return _d.Activity.Restore(ctx, d, id)
}

// Update implements repos.Activity
func (_d ActivityWithTimeout) Update(ctx context.Context, d db.DBTX, id db.ActivityID, params *db.ActivityUpdateParams) (ap1 *db.Activity, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.Activity.Update(ctx, d, id, params)
}
