package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func WorkItemCommentCreateParams(userID models.UserID, workItemID models.WorkItemID) *models.WorkItemCommentCreateParams {
	return &models.WorkItemCommentCreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
		Message:    "message" + testutil.RandomString(100),
		UserID:     userID,
		WorkItemID: workItemID,
	}
}
