package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// WorkItemType represents the repository used for interacting with WorkItemType records.
type WorkItemType struct {
	q models.Querier
}

// NewWorkItemType instantiates the WorkItemType repository.
func NewWorkItemType() *WorkItemType {
	return &WorkItemType{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.WorkItemType = (*WorkItemType)(nil)

func (wit *WorkItemType) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTypeSelectConfigOption) (*models.WorkItemType, error) {
	workItemType, err := models.WorkItemTypeByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item type: %w", ParseDBErrorDetail(err))
	}

	return workItemType, nil
}

func (wit *WorkItemType) ByID(ctx context.Context, d models.DBTX, id models.WorkItemTypeID, opts ...models.WorkItemTypeSelectConfigOption) (*models.WorkItemType, error) {
	workItemType, err := models.WorkItemTypeByWorkItemTypeID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item type: %w", ParseDBErrorDetail(err))
	}

	return workItemType, nil
}
