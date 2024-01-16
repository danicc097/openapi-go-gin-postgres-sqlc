package postgresqltestutil

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func RandomKanbanStepID(project models.Project) db.KanbanStepID {
	switch project {
	case models.ProjectDemo:
		return internal.DemoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoKanbanStepsValues())]
	case models.ProjectDemoTwo:
		return internal.DemoTwoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoTwoKanbanStepsValues())]
	default:
		panic(fmt.Sprintf("invalid project: %s", project))
	}
}

func RandomWorkItemTypeID(project models.Project) db.WorkItemTypeID {
	switch project {
	case models.ProjectDemo:
		return internal.DemoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoWorkItemTypesValues())]
	case models.ProjectDemoTwo:
		return internal.DemoTwoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoTwoWorkItemTypesValues())]
	default:
		panic(fmt.Sprintf("invalid project: %s", project))
	}
}
