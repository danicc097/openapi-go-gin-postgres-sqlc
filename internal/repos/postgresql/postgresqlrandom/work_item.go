package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: Base work items never created via WorkItem repo, always through specific project struct.
func RandomWorkItemCreateParams(kanbanStepID models.KanbanStepID, workItemTypeID models.WorkItemTypeID, teamID models.TeamID) *models.WorkItemCreateParams {
	return &models.WorkItemCreateParams{
		Title:          testutil.RandomNameIdentifier(3, "-"),
		Description:    "Description",
		Metadata:       map[string]any{"key": testutil.RandomString(10)},
		ClosedAt:       nil,
		TargetDate:     testutil.RandomDate(),
		KanbanStepID:   kanbanStepID,
		WorkItemTypeID: workItemTypeID,
		TeamID:         teamID,
	}
}
