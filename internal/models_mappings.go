/**
 * TODO delete after working models_mappings.gen.go
 */

package internal

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"

var (
	ProjectNameByID = map[int]models.Project{
		1: models.ProjectDemoProject,
		2: models.ProjectDemoProject2,
	}
	ProjectIDByName = map[models.Project]int{}
)

var (
	DemoProjectKanbanStepsNameByID = map[int]models.DemoProjectKanbanSteps{
		1: models.DemoProjectKanbanStepsDisabled,
		2: models.DemoProjectKanbanStepsReceived,
	}
	DemoProjectKanbanStepsIDByName = map[models.DemoProjectKanbanSteps]int{}
)

var (
	DemoProject2KanbanStepsNameByID = map[int]models.DemoProject2KanbanSteps{}
	DemoProject2KanbanStepsIDByName = map[models.DemoProject2KanbanSteps]int{}
)
