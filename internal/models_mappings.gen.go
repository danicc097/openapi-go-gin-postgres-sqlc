package internal

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

var (
	ProjectNameByID = map[db.ProjectID]models.Project{
		1: models.ProjectDemo,
		2: models.ProjectDemoTwo,
	}
	ProjectIDByName = map[models.Project]db.ProjectID{
		models.ProjectDemo:    1,
		models.ProjectDemoTwo: 2,
	}
)

var (
	DemoKanbanStepsNameByID = map[db.KanbanStepID]models.DemoKanbanSteps{
		1: models.DemoKanbanStepsDisabled,
		2: models.DemoKanbanStepsReceived,
		3: models.DemoKanbanStepsUnderReview,
		4: models.DemoKanbanStepsWorkInProgress,
	}
	DemoKanbanStepsIDByName = map[models.DemoKanbanSteps]db.KanbanStepID{
		models.DemoKanbanStepsDisabled:       1,
		models.DemoKanbanStepsReceived:       2,
		models.DemoKanbanStepsUnderReview:    3,
		models.DemoKanbanStepsWorkInProgress: 4,
	}
)

var DemoKanbanStepsStepOrderByID = map[int]int{
	1: 0,
	2: 1,
	3: 2,
	4: 3,
}

var (
	DemoTwoKanbanStepsNameByID = map[db.KanbanStepID]models.DemoTwoKanbanSteps{
		5: models.DemoTwoKanbanStepsReceived,
	}
	DemoTwoKanbanStepsIDByName = map[models.DemoTwoKanbanSteps]db.KanbanStepID{
		models.DemoTwoKanbanStepsReceived: 5,
	}
)

var DemoTwoKanbanStepsStepOrderByID = map[int]int{
	5: 1,
}

var (
	DemoWorkItemTypesNameByID = map[db.WorkItemTypeID]models.DemoWorkItemTypes{
		1: models.DemoWorkItemTypesType1,
	}
	DemoWorkItemTypesIDByName = map[models.DemoWorkItemTypes]db.WorkItemTypeID{
		models.DemoWorkItemTypesType1: 1,
	}
)

var (
	DemoTwoWorkItemTypesNameByID = map[db.WorkItemTypeID]models.DemoTwoWorkItemTypes{
		2: models.DemoTwoWorkItemTypesType1,
		3: models.DemoTwoWorkItemTypesType2,
		4: models.DemoTwoWorkItemTypesAnotherType,
	}
	DemoTwoWorkItemTypesIDByName = map[models.DemoTwoWorkItemTypes]db.WorkItemTypeID{
		models.DemoTwoWorkItemTypesType1:       2,
		models.DemoTwoWorkItemTypesType2:       3,
		models.DemoTwoWorkItemTypesAnotherType: 4,
	}
)
