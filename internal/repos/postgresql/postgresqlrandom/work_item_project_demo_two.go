package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func DemoTwoWorkItemCreateParams(kanbanStepID db.KanbanStepID, workItemTypeID db.WorkItemTypeID, teamID db.TeamID) repos.DemoTwoWorkItemCreateParams {
	return repos.DemoTwoWorkItemCreateParams{
		DemoTwoProject: db.DemoTwoWorkItemCreateParams{
			WorkItemID:            db.WorkItemID(-1),
			CustomDateForProject2: pointers.New(testutil.RandomDate()),
		},
		Base: *RandomWorkItemCreateParams(kanbanStepID, workItemTypeID, teamID),
	}
}
