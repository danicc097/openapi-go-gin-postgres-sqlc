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

func (wit *WorkItemTag) Create(ctx context.Context, d db.DBTX, params *db.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	workItemTag, err := db.CreateWorkItemTag(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create time entry: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) Update(ctx context.Context, d db.DBTX, id int, params *db.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	workItemTag, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag by id %w", parseErrorDetail(err))
	}

	workItemTag.SetUpdateParams(params)

	workItemTag, err = workItemTag.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update work item tag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}

func (wit *WorkItemTag) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error) {
	workItemTag, err := db.WorkItemTagByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	workItemTag, err := db.WorkItemTagByWorkItemTagID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	workItemTag, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	err = workItemTag.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete work item tag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}
