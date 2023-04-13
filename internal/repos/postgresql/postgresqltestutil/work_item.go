package postgresqltestutil

import (
	"fmt"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: Base work items never created via WorkItem repo, always through specific project struct

func RandomWorkItemCreateParams(t *testing.T, kanbanStepID, workItemTypeID, teamID int) db.WorkItemCreateParams {
	t.Helper()

	return db.WorkItemCreateParams{
		Title:          testutil.RandomNameIdentifier(3, "-"),
		Description:    "Description",
		Metadata:       []byte(fmt.Sprintf(`{"key":"%s"}`, testutil.RandomString(10))),
		Closed:         nil,
		TargetDate:     testutil.RandomDate(),
		KanbanStepID:   kanbanStepID,
		WorkItemTypeID: workItemTypeID,
		TeamID:         teamID,
	}
}
