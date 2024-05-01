package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func DemoWorkItemCreateParams(kanbanStepID models.KanbanStepID, workItemTypeID models.WorkItemTypeID, teamID models.TeamID) repos.DemoWorkItemCreateParams {
	return repos.DemoWorkItemCreateParams{
		DemoProject: models.DemoWorkItemCreateParams{
			WorkItemID:    models.WorkItemID(-1),
			Ref:           "ref-" + testutil.RandomString(5),
			Line:          "line-" + testutil.RandomString(5),
			Reopened:      testutil.RandomBool(),
			LastMessageAt: testutil.RandomDate(),
		},
		Base: *RandomWorkItemCreateParams(kanbanStepID, workItemTypeID, teamID),
	}
}
