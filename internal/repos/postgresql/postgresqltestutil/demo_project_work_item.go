package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomDemoWorkItem(t *testing.T, pool *pgxpool.Pool, projectID, kanbanStepID, workItemTypeID, teamID int) (*db.DemoWorkItem, error) {
	t.Helper()

	dpwiRepo := postgresql.NewDemoWorkItem()

	dpwicp := RandomDemoWorkItemCreateParams(t)
	wicp := RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID)

	dpwi, err := dpwiRepo.Create(context.Background(), pool, repos.DemoWorkItemCreateParams{DemoProject: dpwicp, Base: wicp})
	if err != nil {
		t.Logf("%s", err)
		return nil, err
	}

	return dpwi, nil
}

func RandomDemoWorkItemCreateParams(t *testing.T) db.DemoWorkItemCreateParams {
	t.Helper()

	return db.DemoWorkItemCreateParams{
		Ref:           "ref-" + testutil.RandomString(5),
		Line:          "line-" + testutil.RandomString(5),
		Reopened:      testutil.RandomBool(),
		LastMessageAt: testutil.RandomDate(),
	}
}
