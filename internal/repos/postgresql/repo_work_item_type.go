package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemType represents the repository used for interacting with WorkItemType records.
type WorkItemType struct {
	q db.Querier
}

// NewWorkItemType instantiates the WorkItemType repository.
func NewWorkItemType() *WorkItemType {
	return &WorkItemType{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.WorkItemType = (*WorkItemType)(nil)

func (wit *WorkItemType) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.WorkItemTypeSelectConfigOption) (*db.WorkItemType, error) {
	workItemType, err := db.WorkItemTypeByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item type: %w", parseDBErrorDetail(err))
	}

	return workItemType, nil
}

func (wit *WorkItemType) ByID(ctx context.Context, d db.DBTX, id db.WorkItemTypeID, opts ...db.WorkItemTypeSelectConfigOption) (*db.WorkItemType, error) {
	workItemType, err := db.WorkItemTypeByWorkItemTypeID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item type: %w", parseDBErrorDetail(err))
	}

	return workItemType, nil
}
