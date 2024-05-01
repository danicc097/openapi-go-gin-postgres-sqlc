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

// WorkItemWithTimeout implements repos.WorkItem interface instrumented with timeouts
type WorkItemWithTimeout struct {
	repos.WorkItem
	config WorkItemWithTimeoutConfig
}

type WorkItemWithTimeoutConfig struct {
	AssignTagTimeout time.Duration

	AssignUserTimeout time.Duration

	ByIDTimeout time.Duration

	DeleteTimeout time.Duration

	RemoveAssignedUserTimeout time.Duration

	RemoveTagTimeout time.Duration

	RestoreTimeout time.Duration
}

// NewWorkItemWithTimeout returns WorkItemWithTimeout
func NewWorkItemWithTimeout(base repos.WorkItem, config WorkItemWithTimeoutConfig) WorkItemWithTimeout {
	return WorkItemWithTimeout{
		WorkItem: base,
		config:   config,
	}
}

// AssignTag implements repos.WorkItem
func (_d WorkItemWithTimeout) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) (err error) {
	var cancelFunc func()
	if _d.config.AssignTagTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.AssignTagTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.AssignTag(ctx, d, params)
}

// AssignUser implements repos.WorkItem
func (_d WorkItemWithTimeout) AssignUser(ctx context.Context, d db.DBTX, params *db.WorkItemAssigneeCreateParams) (err error) {
	var cancelFunc func()
	if _d.config.AssignUserTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.AssignUserTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.AssignUser(ctx, d, params)
}

// ByID implements repos.WorkItem
func (_d WorkItemWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.ByID(ctx, d, id, opts...)
}

// Delete implements repos.WorkItem
func (_d WorkItemWithTimeout) Delete(ctx context.Context, d db.DBTX, id db.WorkItemID) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.Delete(ctx, d, id)
}

// RemoveAssignedUser implements repos.WorkItem
func (_d WorkItemWithTimeout) RemoveAssignedUser(ctx context.Context, d db.DBTX, memberID db.UserID, workItemID db.WorkItemID) (err error) {
	var cancelFunc func()
	if _d.config.RemoveAssignedUserTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.RemoveAssignedUserTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.RemoveAssignedUser(ctx, d, memberID, workItemID)
}

// RemoveTag implements repos.WorkItem
func (_d WorkItemWithTimeout) RemoveTag(ctx context.Context, d db.DBTX, tagID db.WorkItemTagID, workItemID db.WorkItemID) (err error) {
	var cancelFunc func()
	if _d.config.RemoveTagTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.RemoveTagTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.RemoveTag(ctx, d, tagID, workItemID)
}

// Restore implements repos.WorkItem
func (_d WorkItemWithTimeout) Restore(ctx context.Context, d db.DBTX, id db.WorkItemID) (wp1 *db.WorkItem, err error) {
	var cancelFunc func()
	if _d.config.RestoreTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.RestoreTimeout)
		defer cancelFunc()
	}
	return _d.WorkItem.Restore(ctx, d, id)
}
