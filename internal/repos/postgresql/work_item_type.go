package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemType represents the repository used for interacting with WorkItemType records.
type WorkItemType struct {
	q *db.Queries
}

// NewWorkItemType instantiates the WorkItemType repository.
func NewWorkItemType() *WorkItemType {
	return &WorkItemType{
		q: db.New(),
	}
}

var _ repos.WorkItemType = (*WorkItemType)(nil)

func (u *WorkItemType) Create(ctx context.Context, d db.DBTX, params repos.WorkItemTypeCreateParams) (*db.WorkItemType, error) {
	workItemType := &db.WorkItemType{
		Name:        params.Name,
		Description: params.Description,
		ProjectID:   params.ProjectID,
		Color:       params.Color,
	}

	if err := workItemType.Save(ctx, d); err != nil {
		return nil, err
	}

	return workItemType, nil
}

func (u *WorkItemType) Update(ctx context.Context, d db.DBTX, id int, params repos.WorkItemTypeUpdateParams) (*db.WorkItemType, error) {
	workItemType, err := u.WorkItemTypeByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemType by id %w", parseErrorDetail(err))
	}

	if params.Description != nil {
		workItemType.Description = *params.Description
	}
	if params.Name != nil {
		workItemType.Name = *params.Name
	}
	if params.Color != nil {
		workItemType.Color = *params.Color
	}

	err = workItemType.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemType: %w", parseErrorDetail(err))
	}

	return workItemType, err
}

func (u *WorkItemType) WorkItemTypeByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error) {
	workItemType, err := db.WorkItemTypeByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemType: %w", parseErrorDetail(err))
	}

	return workItemType, nil
}

func (u *WorkItemType) WorkItemTypeByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error) {
	workItemType, err := db.WorkItemTypeByWorkItemTypeID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemType: %w", parseErrorDetail(err))
	}

	return workItemType, nil
}

func (u *WorkItemType) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error) {
	workItemType, err := u.WorkItemTypeByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemType by id %w", parseErrorDetail(err))
	}

	err = workItemType.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemType: %w", parseErrorDetail(err))
	}

	return workItemType, err
}
