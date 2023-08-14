package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
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

func (u *WorkItem) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error) {
	return db.WorkItemByWorkItemID(ctx, d, id, opts...)
}

// TODO: remove from demo and demo_two.
func (u *WorkItem) AssignMember(ctx context.Context, d db.DBTX, params *db.WorkItemAssignedUserCreateParams) error {
	_, err := db.CreateWorkItemAssignedUser(ctx, d, params)

	return err
}

// TODO: remove from demo and demo_two.
func (w *WorkItem) RemoveMember(ctx context.Context, d db.DBTX, memberID uuid.UUID, workItemID int) error {
	wim := &db.WorkItemAssignedUser{
		AssignedUser: memberID,
		WorkItemID:   workItemID,
	}

	return wim.Delete(ctx, d)
}

func (u *WorkItem) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	workItem := &db.WorkItem{
		WorkItemID: id,
	}

	err := workItem.SoftDelete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not soft delete workItem: %w", parseDBErrorDetail(err))
	}

	return workItem, err
}

func (u *WorkItem) Restore(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	var err error
	workItem := &db.WorkItem{
		WorkItemID: id,
	}

	workItem, err = workItem.Restore(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not restore workItem: %w", parseDBErrorDetail(err))
	}

	return workItem, err
}
