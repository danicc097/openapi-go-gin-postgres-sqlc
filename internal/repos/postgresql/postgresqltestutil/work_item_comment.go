package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/google/uuid"
)

func RandomWorkItemCommentCreateParams(t *testing.T, workItemID int64, userID uuid.UUID) *db.WorkItemCommentCreateParams {
	t.Helper()

	return &db.WorkItemCommentCreateParams{
		UserID:     userID,
		WorkItemID: workItemID,
		Message:    "message:" + testutil.RandomString(10),
	}
}
