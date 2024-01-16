package postgresqltestutil

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomWorkItemComment(t *testing.T, d db.DBTX, project models.Project) *db.WorkItemComment {
	t.Helper()

	workItemCommentRepo := reposwrappers.NewWorkItemCommentWithRetry(postgresql.NewWorkItemComment(), testutil.NewLogger(t), 3, 200*time.Millisecond)

	var workItemID db.WorkItemID
	switch project {
	case models.ProjectDemo:
		workItemID = NewRandomDemoWorkItem(t, d).WorkItemID
	case models.ProjectDemoTwo:
		workItemID = NewRandomDemoTwoWorkItem(t, d).WorkItemID
	}

	user := NewRandomUser(t, d)
	// these are repo test utils. don't care about logic concerning
	// "is user assigned to the same team as the workitem" or anything similar defined
	// at the service level, unless it's checked at the db level for some reason
	// If we need to test logic like that, use createParams.
	// services and api should use fixture factory instead so that it uses specific service logic for creation.
	// TODO: add project script checking postgresqltestutil.NewRandom* strings are not found outside repos

	ucp := RandomWorkItemCommentCreateParams(t, user.UserID, workItemID)

	workItemComment, err := workItemCommentRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return workItemComment
}

// NOTE: FKs should always be passed explicitly.
func RandomWorkItemCommentCreateParams(t *testing.T, userID db.UserID, workItemID db.WorkItemID) *db.WorkItemCommentCreateParams {
	t.Helper()

	return &db.WorkItemCommentCreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
		Message:    "message" + testutil.RandomString(100),
		UserID:     userID,
		WorkItemID: workItemID,
	}
}
