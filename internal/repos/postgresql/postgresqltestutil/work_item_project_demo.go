package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func RandomDemoWorkItemCreateParams(t *testing.T, kanbanStepID db.KanbanStepID, workItemTypeID db.WorkItemTypeID, teamID db.TeamID) repos.DemoWorkItemCreateParams {
	return repos.DemoWorkItemCreateParams{
		DemoProject: db.DemoWorkItemCreateParams{
			WorkItemID:    db.WorkItemID(-1),
			Ref:           "ref-" + testutil.RandomString(5),
			Line:          "line-" + testutil.RandomString(5),
			Reopened:      testutil.RandomBool(),
			LastMessageAt: testutil.RandomDate(),
		},
		Base: *RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID),
	}
}
