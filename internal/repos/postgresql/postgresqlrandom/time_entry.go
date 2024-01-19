package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func TimeEntryCreateParams(activityID db.ActivityID, userID db.UserID, workItemID *db.WorkItemID, teamID *db.TeamID) *db.TimeEntryCreateParams {
	return &db.TimeEntryCreateParams{
		WorkItemID:      workItemID,
		ActivityID:      activityID,
		TeamID:          teamID,
		UserID:          userID,
		Comment:         testutil.RandomString(20),
		Start:           testutil.RandomDate(),
		DurationMinutes: pointers.New(int(testutil.RandomInt64(10, 400))),
	}
}
