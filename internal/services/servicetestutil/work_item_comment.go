package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/stretchr/testify/require"
)

type CreateWorkItemCommentParams struct {
	UserID     db.UserID
	Project    models.Project
	WorkItemID db.WorkItemID
}

type CreateWorkItemCommentFixture struct {
	*db.WorkItemComment
}

// CreateWorkItemComment creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItemComment(ctx context.Context, params CreateWorkItemCommentParams) *CreateWorkItemCommentFixture {
	randomRepoCreateParams := postgresqlrandom.WorkItemCommentCreateParams(params.UserID, params.WorkItemID)
	// don't use repos for test fixtures, use service logic
	workItemComment, err := ff.svc.WorkItemComment.Create(ctx, ff.d, randomRepoCreateParams)
	require.NoError(ff.t, err)

	return &CreateWorkItemCommentFixture{
		WorkItemComment: workItemComment,
	}
}
