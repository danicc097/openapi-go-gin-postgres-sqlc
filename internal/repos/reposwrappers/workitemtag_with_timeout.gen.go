// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// WorkItemTagWithTimeout implements _sourceRepos.WorkItemTag interface instrumented with timeouts
type WorkItemTagWithTimeout struct {
	_sourceRepos.WorkItemTag
	config WorkItemTagWithTimeoutConfig
}

type WorkItemTagWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByNameTimeout time.Duration

	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewWorkItemTagWithTimeout returns WorkItemTagWithTimeout
func NewWorkItemTagWithTimeout(base _sourceRepos.WorkItemTag, config WorkItemTagWithTimeoutConfig) WorkItemTagWithTimeout {
	return WorkItemTagWithTimeout{
		WorkItemTag: base,
		config:      config,
	}
}

// ByID implements _sourceRepos.WorkItemTag
func (_d WorkItemTagWithTimeout) ByID(ctx context.Context, d models.DBTX, id models.WorkItemTagID, opts ...models.WorkItemTagSelectConfigOption) (wp1 *models.WorkItemTag, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemTag.ByID(ctx, d, id, opts...)
}

// ByName implements _sourceRepos.WorkItemTag
func (_d WorkItemTagWithTimeout) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTagSelectConfigOption) (wp1 *models.WorkItemTag, err error) {
	var cancelFunc func()
	if _d.config.ByNameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByNameTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemTag.ByName(ctx, d, name, projectID, opts...)
}

// Create implements _sourceRepos.WorkItemTag
func (_d WorkItemTagWithTimeout) Create(ctx context.Context, d models.DBTX, params *models.WorkItemTagCreateParams) (wp1 *models.WorkItemTag, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemTag.Create(ctx, d, params)
}

// Delete implements _sourceRepos.WorkItemTag
func (_d WorkItemTagWithTimeout) Delete(ctx context.Context, d models.DBTX, id models.WorkItemTagID) (wp1 *models.WorkItemTag, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemTag.Delete(ctx, d, id)
}

// Update implements _sourceRepos.WorkItemTag
func (_d WorkItemTagWithTimeout) Update(ctx context.Context, d models.DBTX, id models.WorkItemTagID, params *models.WorkItemTagUpdateParams) (wp1 *models.WorkItemTag, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemTag.Update(ctx, d, id, params)
}
