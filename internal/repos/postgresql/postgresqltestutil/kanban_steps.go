package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomKanbanStep(t *testing.T, pool *pgxpool.Pool, projectID int) (*db.KanbanStep, error) {
	t.Helper()

	ksRepo := postgresql.NewKanbanStep()

	ucp := RandomKanbanStepCreateParams(t, projectID)

	ks, err := ksRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		t.Logf("%s", err)
		return nil, err
	}

	return ks, nil
}

func RandomKanbanStepCreateParams(t *testing.T, projectID int) db.KanbanStepCreateParams {
	t.Helper()

	return db.KanbanStepCreateParams{
		Name:          "KanbanStep " + testutil.RandomNameIdentifier(3, "-"),
		Description:   testutil.RandomString(10),
		ProjectID:     projectID,
		Color:         "#aaaaaa",
		TimeTrackable: testutil.RandomBool(),
		StepOrder:     pointers.New(int16(testutil.RandomInt64(1, 32766))),
	}
}
