package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// WorkItem represents the repository used for interacting with WorkItem records.
type WorkItem struct {
	q models.Querier
}

// NewWorkItem instantiates the WorkItem repository.
// NOTE: maybe we can consider work item an aggregate root, since we don't
// need distinction between projects for some tasks like assigning members, tags, generic
// functionality like Delete, Restore...
// and this simplifies everything a lot.
func NewWorkItem() *WorkItem {
	return &WorkItem{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.WorkItem = (*WorkItem)(nil)

func (w *WorkItem) ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error) {
	return models.WorkItemByWorkItemID(ctx, d, id, opts...)
}

func (w *WorkItem) AssignUser(ctx context.Context, d models.DBTX, params *models.WorkItemAssigneeCreateParams) error {
	_, err := models.CreateWorkItemAssignee(ctx, d, params)

	return err
}

func (w *WorkItem) RemoveAssignedUser(ctx context.Context, d models.DBTX, memberID models.UserID, workItemID models.WorkItemID) error {
	lookup := &models.WorkItemAssignee{
		Assignee:   memberID,
		WorkItemID: workItemID,
	}

	return lookup.Delete(ctx, d)
}

func (w *WorkItem) AssignTag(ctx context.Context, d models.DBTX, params *models.WorkItemWorkItemTagCreateParams) error {
	_, err := models.CreateWorkItemWorkItemTag(ctx, d, params)

	return err
}

func (w *WorkItem) RemoveTag(ctx context.Context, d models.DBTX, tagID models.WorkItemTagID, workItemID models.WorkItemID) error {
	lookup := &models.WorkItemWorkItemTag{
		WorkItemTagID: tagID,
		WorkItemID:    workItemID,
	}

	return lookup.Delete(ctx, d)
}

func (w *WorkItem) Delete(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error) {
	workItem := &models.WorkItem{
		WorkItemID: id,
	}

	err := workItem.SoftDelete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not soft delete workItem: %w", ParseDBErrorDetail(err))
	}

	return workItem, err
}

func (w *WorkItem) Restore(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error) {
	var err error
	workItem := &models.WorkItem{
		WorkItemID: id,
	}

	workItem, err = workItem.Restore(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not restore workItem: %w", ParseDBErrorDetail(err))
	}

	return workItem, err
}
