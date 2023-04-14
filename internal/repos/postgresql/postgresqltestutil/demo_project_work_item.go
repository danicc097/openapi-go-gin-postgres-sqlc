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

func NewRandomDemoProjectWorkItem(t *testing.T, pool *pgxpool.Pool, projectID, kanbanStepID, workItemTypeID, teamID int) (*db.DemoProjectWorkItem, error) {
	t.Helper()

	dpwiRepo := postgresql.NewDemoProjectWorkItem()

	dpwicp := RandomDemoProjectWorkItemCreateParams(t)
	wicp := RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID)

	dpwi, err := dpwiRepo.Create(context.Background(), pool, repos.DemoProjectWorkItemCreateParams{DemoProject: dpwicp, Base: wicp})
	if err != nil {
		return nil, err
	}

	return dpwi, nil
}

func RandomDemoProjectWorkItemCreateParams(t *testing.T) db.DemoProjectWorkItemCreateParams {
	t.Helper()

	return db.DemoProjectWorkItemCreateParams{
		Ref:           "ref-" + testutil.RandomString(5),
		Line:          "line-" + testutil.RandomString(5),
		Reopened:      testutil.RandomBool(),
		LastMessageAt: testutil.RandomDate(),
	}
}
