package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemTag represents the repository used for interacting with WorkItemTag records.
type WorkItemTag struct {
	q *db.Queries
}

// NewWorkItemTag instantiates the WorkItemTag repository.
func NewWorkItemTag() *WorkItemTag {
	return &WorkItemTag{
		q: db.New(),
	}
}

var _ repos.WorkItemTag = (*WorkItemTag)(nil)

func (u *WorkItemTag) Create(ctx context.Context, d db.DBTX, params repos.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	workItemTag := &db.WorkItemTag{
		Name:        params.Name,
		Description: params.Description,
		ProjectID:   params.ProjectID,
		Color:       params.Color,
	}

	if err := workItemTag.Save(ctx, d); err != nil {
		return nil, err
	}

	return workItemTag, nil
}

func (u *WorkItemTag) Update(ctx context.Context, d db.DBTX, id int, params repos.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	workItemTag, err := u.WorkItemTagByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	if params.Description != nil {
		workItemTag.Description = *params.Description
	}
	if params.Name != nil {
		workItemTag.Name = *params.Name
	}
	if params.Color != nil {
		workItemTag.Color = *params.Color
	}

	err = workItemTag.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}

func (u *WorkItemTag) WorkItemTagByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error) {
	workItemTag, err := db.WorkItemTagByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (u *WorkItemTag) WorkItemTagByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	workItemTag, err := db.WorkItemTagByWorkItemTagID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (u *WorkItemTag) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	workItemTag, err := u.WorkItemTagByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	err = workItemTag.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}
