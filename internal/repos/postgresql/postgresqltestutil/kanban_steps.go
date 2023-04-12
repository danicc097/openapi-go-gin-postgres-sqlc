package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

func RandomKanbanStepCreateParams(t *testing.T, projectID int) db.KanbanStepCreateParams {
	t.Helper()

	return db.KanbanStepCreateParams{
		Name:          "KanbanStep " + testutil.RandomNameIdentifier(3, "-"),
		Description:   testutil.RandomString(10),
		ProjectID:     projectID,
		Color:         "#aaaaaa",
		TimeTrackable: testutil.RandomBool(),
		StepOrder:     pointers.New(int16(testutil.RandomInt(1, 8))),
	}
}
