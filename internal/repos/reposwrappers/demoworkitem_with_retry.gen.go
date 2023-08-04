// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// DemoWorkItemWithRetry implements repos.DemoWorkItem interface instrumented with retries
type DemoWorkItemWithRetry struct {
	repos.DemoWorkItem
	_retryCount    int
	_retryInterval time.Duration
}

// NewDemoWorkItemWithRetry returns DemoWorkItemWithRetry
func NewDemoWorkItemWithRetry(base repos.DemoWorkItem, retryCount int, retryInterval time.Duration) DemoWorkItemWithRetry {
	return DemoWorkItemWithRetry{
		DemoWorkItem:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.DemoWorkItem
func (_d DemoWorkItemWithRetry) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	wp1, err = _d.DemoWorkItem.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		wp1, err = _d.DemoWorkItem.ByID(ctx, d, id, opts...)
	}
	return
}

// Create implements repos.DemoWorkItem
func (_d DemoWorkItemWithRetry) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	wp1, err = _d.DemoWorkItem.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		wp1, err = _d.DemoWorkItem.Create(ctx, d, params)
	}
	return
}

// Update implements repos.DemoWorkItem
func (_d DemoWorkItemWithRetry) Update(ctx context.Context, d db.DBTX, id int, params repos.DemoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	wp1, err = _d.DemoWorkItem.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		wp1, err = _d.DemoWorkItem.Update(ctx, d, id, params)
	}
	return
}
