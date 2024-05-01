package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/stretchr/testify/require"
)

type CreateWorkItemCommentParams struct {
	UserID     models.UserID
	WorkItemID models.WorkItemID
}

type CreateWorkItemCommentFixture struct {
	*models.WorkItemComment
}

// CreateWorkItemComment creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItemComment(ctx context.Context, userID models.UserID, workItemID models.WorkItemID) *CreateWorkItemCommentFixture {
	randomRepoCreateParams := postgresqlrandom.WorkItemCommentCreateParams(userID, workItemID)
	// don't use repos for test fixtures, use service logic
	workItemComment, err := ff.svc.WorkItemComment.Create(ctx, ff.d, randomRepoCreateParams)
	require.NoError(ff.t, err)

	return &CreateWorkItemCommentFixture{
		WorkItemComment: workItemComment,
	}
}
