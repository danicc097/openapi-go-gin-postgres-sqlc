package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemComment represents the repository used for interacting with WorkItemComment records.
type WorkItemComment struct {
	q db.Querier
}

// NewWorkItemComment instantiates the WorkItemComment repository.
func NewWorkItemComment() *WorkItemComment {
	return &WorkItemComment{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.WorkItemComment = (*WorkItemComment)(nil)

func (wit *WorkItemComment) Create(ctx context.Context, d db.DBTX, params *db.WorkItemCommentCreateParams) (*db.WorkItemComment, error) {
	workItemComment, err := db.CreateWorkItemComment(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create workItemComment: %w", parseDBErrorDetail(err))
	}

	return workItemComment, nil
}

func (wit *WorkItemComment) Update(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, params *db.WorkItemCommentUpdateParams) (*db.WorkItemComment, error) {
	workItemComment, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemComment by id %w", parseDBErrorDetail(err))
	}

	workItemComment.SetUpdateParams(params)

	workItemComment, err = workItemComment.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemComment: %w", parseDBErrorDetail(err))
	}

	return workItemComment, err
}

func (wit *WorkItemComment) ByID(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, opts ...db.WorkItemCommentSelectConfigOption) (*db.WorkItemComment, error) {
	workItemComment, err := db.WorkItemCommentByWorkItemCommentID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemComment: %w", parseDBErrorDetail(err))
	}

	return workItemComment, nil
}

func (wit *WorkItemComment) Delete(ctx context.Context, d db.DBTX, id db.WorkItemCommentID) (*db.WorkItemComment, error) {
	workItemComment := &db.WorkItemComment{
		WorkItemCommentID: id,
	}

	err := workItemComment.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemComment: %w", parseDBErrorDetail(err))
	}

	return workItemComment, err
}
