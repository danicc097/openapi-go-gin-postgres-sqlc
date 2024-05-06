package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func DemoTwoWorkItemCreateParams(kanbanStepID models.KanbanStepID, workItemTypeID models.WorkItemTypeID, teamID models.TeamID) repos.DemoTwoWorkItemCreateParams {
	return repos.DemoTwoWorkItemCreateParams{
		DemoTwoProject: models.DemoTwoWorkItemCreateParams{
			WorkItemID:            models.WorkItemID(-1),
			CustomDateForProject2: pointers.New(testutil.RandomDate()),
		},
		Base: *RandomWorkItemCreateParams(kanbanStepID, workItemTypeID, teamID),
	}
}
