package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

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
