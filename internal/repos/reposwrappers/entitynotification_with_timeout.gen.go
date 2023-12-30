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

// EntityNotificationWithTimeout implements repos.EntityNotification interface instrumented with timeouts
type EntityNotificationWithTimeout struct {
	repos.EntityNotification
	config EntityNotificationWithTimeoutConfig
}

type EntityNotificationWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewEntityNotificationWithTimeout returns EntityNotificationWithTimeout
func NewEntityNotificationWithTimeout(base repos.EntityNotification, config EntityNotificationWithTimeoutConfig) EntityNotificationWithTimeout {
	return EntityNotificationWithTimeout{
		EntityNotification: base,
		config:             config,
	}
}

// ByID implements repos.EntityNotification
func (_d EntityNotificationWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.EntityNotificationID, opts ...db.EntityNotificationSelectConfigOption) (ep1 *db.EntityNotification, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.EntityNotification.ByID(ctx, d, id, opts...)
}

// Create implements repos.EntityNotification
func (_d EntityNotificationWithTimeout) Create(ctx context.Context, d db.DBTX, params *db.EntityNotificationCreateParams) (ep1 *db.EntityNotification, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.EntityNotification.Create(ctx, d, params)
}

// Delete implements repos.EntityNotification
func (_d EntityNotificationWithTimeout) Delete(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (ep1 *db.EntityNotification, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.EntityNotification.Delete(ctx, d, id)
}

// Update implements repos.EntityNotification
func (_d EntityNotificationWithTimeout) Update(ctx context.Context, d db.DBTX, id db.EntityNotificationID, params *db.EntityNotificationUpdateParams) (ep1 *db.EntityNotification, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.EntityNotification.Update(ctx, d, id, params)
}
