package postgresqlrandom

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func KanbanStepID(project models.ProjectName) db.KanbanStepID {
	switch project {
	case models.ProjectNameDemo:
		return internal.DemoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoKanbanStepsValues())]
	case models.ProjectNameDemoTwo:
		return internal.DemoTwoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoTwoKanbanStepsValues())]
	default:
		panic(fmt.Sprintf("invalid project: %s", project))
	}
}

func WorkItemTypeID(project models.ProjectName) db.WorkItemTypeID {
	switch project {
	case models.ProjectNameDemo:
		return internal.DemoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoWorkItemTypesValues())]
	case models.ProjectNameDemoTwo:
		return internal.DemoTwoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoTwoWorkItemTypesValues())]
	default:
		panic(fmt.Sprintf("invalid project: %s", project))
	}
}
