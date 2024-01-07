package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItem represents the repository used for interacting with WorkItem records.
type WorkItem struct {
	q db.Querier
}

// NewWorkItem instantiates the WorkItem repository.
// NOTE: maybe we can consider work item an aggregate root, since we don't
// need distinction between projects for some tasks like assigning members, tags, generic
// functionality like Delete, Restore...
// and this simplifies everything a lot.
func NewWorkItem() *WorkItem {
	return &WorkItem{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.WorkItem = (*WorkItem)(nil)

func (w *WorkItem) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error) {
	return db.WorkItemByWorkItemID(ctx, d, id, opts...)
}

func (w *WorkItem) AssignUser(ctx context.Context, d db.DBTX, params *db.WorkItemAssignedUserCreateParams) error {
	_, err := db.CreateWorkItemAssignedUser(ctx, d, params)

	return err
}

func (w *WorkItem) RemoveAssignedUser(ctx context.Context, d db.DBTX, memberID db.UserID, workItemID db.WorkItemID) error {
	lookup := &db.WorkItemAssignedUser{
		AssignedUser: memberID,
		WorkItemID:   workItemID,
	}

	return lookup.Delete(ctx, d)
}

func (w *WorkItem) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) error {
	_, err := db.CreateWorkItemWorkItemTag(ctx, d, params)

	return err
}

func (w *WorkItem) RemoveTag(ctx context.Context, d db.DBTX, tagID db.WorkItemTagID, workItemID db.WorkItemID) error {
	lookup := &db.WorkItemWorkItemTag{
		WorkItemTagID: tagID,
		WorkItemID:    workItemID,
	}

	return lookup.Delete(ctx, d)
}

func (w *WorkItem) Delete(ctx context.Context, d db.DBTX, id db.WorkItemID) (*db.WorkItem, error) {
	workItem := &db.WorkItem{
		WorkItemID: id,
	}

	err := workItem.SoftDelete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not soft delete workItem: %w", ParseDBErrorDetail(err))
	}

	return workItem, err
}

func (w *WorkItem) Restore(ctx context.Context, d db.DBTX, id db.WorkItemID) (*db.WorkItem, error) {
	var err error
	workItem := &db.WorkItem{
		WorkItemID: id,
	}

	workItem, err = workItem.Restore(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not restore workItem: %w", ParseDBErrorDetail(err))
	}

	return workItem, err
}
