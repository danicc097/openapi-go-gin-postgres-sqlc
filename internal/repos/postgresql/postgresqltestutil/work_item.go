package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: Base work items never created via WorkItem repo, always through specific project struct.
func RandomWorkItemCreateParams(t *testing.T, kanbanStepID db.KanbanStepID, workItemTypeID db.WorkItemTypeID, teamID db.TeamID) *db.WorkItemCreateParams {
	t.Helper()

	return &db.WorkItemCreateParams{
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
