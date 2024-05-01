// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// TeamWithTimeout implements repos.Team interface instrumented with timeouts
type TeamWithTimeout struct {
	repos.Team
	config TeamWithTimeoutConfig
}

type TeamWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByNameTimeout time.Duration

	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewTeamWithTimeout returns TeamWithTimeout
func NewTeamWithTimeout(base repos.Team, config TeamWithTimeoutConfig) TeamWithTimeout {
	return TeamWithTimeout{
		Team:   base,
		config: config,
	}
}

// ByID implements repos.Team
func (_d TeamWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.TeamID, opts ...db.TeamSelectConfigOption) (tp1 *db.Team, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.Team.ByID(ctx, d, id, opts...)
}

// ByName implements repos.Team
func (_d TeamWithTimeout) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.TeamSelectConfigOption) (tp1 *db.Team, err error) {
	var cancelFunc func()
	if _d.config.ByNameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByNameTimeout)
		defer cancelFunc()
	}
	return _d.Team.ByName(ctx, d, name, projectID, opts...)
}

// Create implements repos.Team
func (_d TeamWithTimeout) Create(ctx context.Context, d db.DBTX, params *db.TeamCreateParams) (tp1 *db.Team, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.Team.Create(ctx, d, params)
}

// Delete implements repos.Team
func (_d TeamWithTimeout) Delete(ctx context.Context, d db.DBTX, id db.TeamID) (tp1 *db.Team, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.Team.Delete(ctx, d, id)
}

// Update implements repos.Team
func (_d TeamWithTimeout) Update(ctx context.Context, d db.DBTX, id db.TeamID, params *db.TeamUpdateParams) (tp1 *db.Team, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.Team.Update(ctx, d, id, params)
}
