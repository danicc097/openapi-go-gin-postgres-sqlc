package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// WorkItemComment represents the repository used for interacting with work item comment records.
type WorkItemComment struct {
	q models.Querier
}

// NewWorkItemComment instantiates the work item comment repository.
func NewWorkItemComment() *WorkItemComment {
	return &WorkItemComment{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.WorkItemComment = (*WorkItemComment)(nil)

func (t *WorkItemComment) Create(ctx context.Context, d models.DBTX, params *models.WorkItemCommentCreateParams) (*models.WorkItemComment, error) {
	workItemComment, err := models.CreateWorkItemComment(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create workItemComment: %w", ParseDBErrorDetail(err))
	}

	return workItemComment, nil
}

func (t *WorkItemComment) Update(ctx context.Context, d models.DBTX, id models.WorkItemCommentID, params *models.WorkItemCommentUpdateParams) (*models.WorkItemComment, error) {
	workItemComment, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get work item comment by id %w", ParseDBErrorDetail(err))
	}

	workItemComment.SetUpdateParams(params)

	workItemComment, err = workItemComment.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update work item comment: %w", ParseDBErrorDetail(err))
	}

	return workItemComment, err
}

func (t *WorkItemComment) ByID(ctx context.Context, d models.DBTX, id models.WorkItemCommentID, opts ...models.WorkItemCommentSelectConfigOption) (*models.WorkItemComment, error) {
	workItemComment, err := models.WorkItemCommentByWorkItemCommentID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get work item comment: %w", ParseDBErrorDetail(err))
	}

	return workItemComment, nil
}

func (t *WorkItemComment) Delete(ctx context.Context, d models.DBTX, id models.WorkItemCommentID) (*models.WorkItemComment, error) {
	workItemComment := &models.WorkItemComment{
		WorkItemCommentID: id,
	}

	err := workItemComment.Delete(ctx, d) // use SoftDelete if a deleted_at column exists.
	if err != nil {
		return nil, fmt.Errorf("could not delete work item comment: %w", ParseDBErrorDetail(err))
	}

	return workItemComment, err
}
