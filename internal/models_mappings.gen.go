// Code generated by project. DO NOT EDIT.

package internal

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

var (
	ProjectNameByID = map[models.ProjectID]models.ProjectName{
		1: models.ProjectNameDemo,
		2: models.ProjectNameDemoTwo,
	}
	ProjectIDByName = map[models.ProjectName]models.ProjectID{
		models.ProjectNameDemo:    1,
		models.ProjectNameDemoTwo: 2,
	}
)

var (
	DemoKanbanStepsNameByID = map[models.KanbanStepID]models.DemoKanbanSteps{
		1: models.DemoKanbanStepsDisabled,
		2: models.DemoKanbanStepsReceived,
		3: models.DemoKanbanStepsUnderReview,
		4: models.DemoKanbanStepsWorkInProgress,
	}
	DemoKanbanStepsIDByName = map[models.DemoKanbanSteps]models.KanbanStepID{
		models.DemoKanbanStepsDisabled:       1,
		models.DemoKanbanStepsReceived:       2,
		models.DemoKanbanStepsUnderReview:    3,
		models.DemoKanbanStepsWorkInProgress: 4,
	}
)

var DemoKanbanStepsStepOrderByID = map[models.KanbanStepID]int{
	1: 0,
	2: 1,
	3: 2,
	4: 3,
}

var (
	DemoTwoKanbanStepsNameByID = map[models.KanbanStepID]models.DemoTwoKanbanSteps{
		5: models.DemoTwoKanbanStepsReceived,
	}
	DemoTwoKanbanStepsIDByName = map[models.DemoTwoKanbanSteps]models.KanbanStepID{
		models.DemoTwoKanbanStepsReceived: 5,
	}
)

var DemoTwoKanbanStepsStepOrderByID = map[models.KanbanStepID]int{
	5: 1,
}

var (
	DemoWorkItemTypesNameByID = map[models.WorkItemTypeID]models.DemoWorkItemTypes{
		1: models.DemoWorkItemTypesType1,
	}
	DemoWorkItemTypesIDByName = map[models.DemoWorkItemTypes]models.WorkItemTypeID{
		models.DemoWorkItemTypesType1: 1,
	}
)

var (
	DemoTwoWorkItemTypesNameByID = map[models.WorkItemTypeID]models.DemoTwoWorkItemTypes{
		2: models.DemoTwoWorkItemTypesType1,
		3: models.DemoTwoWorkItemTypesType2,
		4: models.DemoTwoWorkItemTypesAnotherType,
	}
	DemoTwoWorkItemTypesIDByName = map[models.DemoTwoWorkItemTypes]models.WorkItemTypeID{
		models.DemoTwoWorkItemTypesType1:       2,
		models.DemoTwoWorkItemTypesType2:       3,
		models.DemoTwoWorkItemTypesAnotherType: 4,
	}
)
