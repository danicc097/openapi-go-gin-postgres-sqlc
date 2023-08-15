package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func RandomWorkItemCommentCreateParams(t *testing.T, workItemID db.WorkItemID, userID db.UserID) *db.WorkItemCommentCreateParams {
	t.Helper()

	return &db.WorkItemCommentCreateParams{
		UserID:     userID,
		WorkItemID: workItemID,
		Message:    "message:" + testutil.RandomString(10),
	}
}
