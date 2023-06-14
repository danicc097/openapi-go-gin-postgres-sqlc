package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemComment represents the repository used for interacting with WorkItemComment records.
type WorkItemComment struct {
	q *db.Queries
}

// NewWorkItemComment instantiates the WorkItemComment repository.
func NewWorkItemComment() *WorkItemComment {
	return &WorkItemComment{
		q: db.New(),
	}
}

var _ repos.WorkItemComment = (*WorkItemComment)(nil)

func (wit *WorkItemComment) Create(ctx context.Context, d db.DBTX, params *db.WorkItemCommentCreateParams) (*db.WorkItemComment, error) {
	workItemComment, err := db.CreateWorkItemComment(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create workItemComment: %w", parseErrorDetail(err))
	}

	return workItemComment, nil
}

func (wit *WorkItemComment) Update(ctx context.Context, d db.DBTX, id int64, params *db.WorkItemCommentUpdateParams) (*db.WorkItemComment, error) {
	workItemComment, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemComment by id %w", parseErrorDetail(err))
	}

	workItemComment, err = workItemComment.Update(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemComment: %w", parseErrorDetail(err))
	}

	return workItemComment, err
}

func (wit *WorkItemComment) ByID(ctx context.Context, d db.DBTX, id int64) (*db.WorkItemComment, error) {
	workItemComment, err := db.WorkItemCommentByWorkItemCommentID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemComment: %w", parseErrorDetail(err))
	}

	return workItemComment, nil
}

func (wit *WorkItemComment) Delete(ctx context.Context, d db.DBTX, id int64) (*db.WorkItemComment, error) {
	workItemComment, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemComment by id %w", parseErrorDetail(err))
	}

	err = workItemComment.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemComment: %w", parseErrorDetail(err))
	}

	return workItemComment, err
}
