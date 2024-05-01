package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// WorkItemTag represents the repository used for interacting with WorkItemTag records.
type WorkItemTag struct {
	q models.Querier
}

// NewWorkItemTag instantiates the WorkItemTag repository.
func NewWorkItemTag() *WorkItemTag {
	return &WorkItemTag{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.WorkItemTag = (*WorkItemTag)(nil)

func (wit *WorkItemTag) Create(ctx context.Context, d models.DBTX, params *models.WorkItemTagCreateParams) (*models.WorkItemTag, error) {
	workItemTag, err := models.CreateWorkItemTag(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create work item tag: %w", ParseDBErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) Update(ctx context.Context, d models.DBTX, id models.WorkItemTagID, params *models.WorkItemTagUpdateParams) (*models.WorkItemTag, error) {
	workItemTag, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag by id %w", ParseDBErrorDetail(err))
	}

	workItemTag.SetUpdateParams(params)

	workItemTag, err = workItemTag.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update work item tag: %w", ParseDBErrorDetail(err))
	}

	return workItemTag, err
}

func (wit *WorkItemTag) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.WorkItemTagSelectConfigOption) (*models.WorkItemTag, error) {
	workItemTag, err := models.WorkItemTagByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag: %w", ParseDBErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) ByID(ctx context.Context, d models.DBTX, id models.WorkItemTagID, opts ...models.WorkItemTagSelectConfigOption) (*models.WorkItemTag, error) {
	workItemTag, err := models.WorkItemTagByWorkItemTagID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item tag: %w", ParseDBErrorDetail(err))
	}

	return workItemTag, nil
}

func (wit *WorkItemTag) Delete(ctx context.Context, d models.DBTX, id models.WorkItemTagID) (*models.WorkItemTag, error) {
	workItemTag := &models.WorkItemTag{
		WorkItemTagID: id,
	}

	err := workItemTag.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete work item tag: %w", ParseDBErrorDetail(err))
	}

	return workItemTag, err
}
