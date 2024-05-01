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

// DemoTwoWorkItemWithTimeout implements repos.DemoTwoWorkItem interface instrumented with timeouts
type DemoTwoWorkItemWithTimeout struct {
	repos.DemoTwoWorkItem
	config DemoTwoWorkItemWithTimeoutConfig
}

type DemoTwoWorkItemWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	CreateTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewDemoTwoWorkItemWithTimeout returns DemoTwoWorkItemWithTimeout
func NewDemoTwoWorkItemWithTimeout(base repos.DemoTwoWorkItem, config DemoTwoWorkItemWithTimeoutConfig) DemoTwoWorkItemWithTimeout {
	return DemoTwoWorkItemWithTimeout{
		DemoTwoWorkItem: base,
		config:          config,
	}
}

// ByID implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.DemoTwoWorkItem.ByID(ctx, d, id, opts...)
}

// Create implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTimeout) Create(ctx context.Context, d db.DBTX, params repos.DemoTwoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.DemoTwoWorkItem.Create(ctx, d, params)
}

// Update implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithTimeout) Update(ctx context.Context, d db.DBTX, id db.WorkItemID, params repos.DemoTwoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.DemoTwoWorkItem.Update(ctx, d, id, params)
}
