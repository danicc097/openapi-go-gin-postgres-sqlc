package postgresqltestutil

import (
	"fmt"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgtype"
)

func RandomWorkItemCreateParams(t *testing.T, kanbanStepID, workItemTypeID, teamID int) db.WorkItemCreateParams {
	t.Helper()

	return db.WorkItemCreateParams{
		Title:       testutil.RandomNameIdentifier(3, "-"),
		Description: "Description",
		Metadata: pgtype.JSONB{
			Bytes: []byte(fmt.Sprintf(`{"key":"%s"}`, testutil.RandomString(10))),
		},
		Closed:         nil,
		TargetDate:     testutil.RandomDate(),
		KanbanStepID:   kanbanStepID,
		WorkItemTypeID: workItemTypeID,
		TeamID:         teamID,
	}
}
