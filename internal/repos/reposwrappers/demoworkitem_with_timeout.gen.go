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

// DemoWorkItemWithTimeout implements repos.DemoWorkItem interface instrumented with timeouts
type DemoWorkItemWithTimeout struct {
	repos.DemoWorkItem
	config DemoWorkItemWithTimeoutConfig
}

type DemoWorkItemWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	CreateTimeout time.Duration

	PaginatedTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewDemoWorkItemWithTimeout returns DemoWorkItemWithTimeout
func NewDemoWorkItemWithTimeout(base repos.DemoWorkItem, config DemoWorkItemWithTimeoutConfig) DemoWorkItemWithTimeout {
	return DemoWorkItemWithTimeout{
		DemoWorkItem: base,
		config:       config,
	}
}

// ByID implements repos.DemoWorkItem
func (_d DemoWorkItemWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.DemoWorkItem.ByID(ctx, d, id, opts...)
}

// Create implements repos.DemoWorkItem
func (_d DemoWorkItemWithTimeout) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.DemoWorkItem.Create(ctx, d, params)
}

// Paginated implements repos.DemoWorkItem
func (_d DemoWorkItemWithTimeout) Paginated(ctx context.Context, d db.DBTX, cursor db.WorkItemID, opts ...db.CacheDemoWorkItemSelectConfigOption) (ca1 []db.CacheDemoWorkItem, err error) {
	var cancelFunc func()
	if _d.config.PaginatedTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.PaginatedTimeout)
		defer cancelFunc()
	}
	return _d.DemoWorkItem.Paginated(ctx, d, cursor, opts...)
}

// Update implements repos.DemoWorkItem
func (_d DemoWorkItemWithTimeout) Update(ctx context.Context, d db.DBTX, id db.WorkItemID, params repos.DemoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.DemoWorkItem.Update(ctx, d, id, params)
}
