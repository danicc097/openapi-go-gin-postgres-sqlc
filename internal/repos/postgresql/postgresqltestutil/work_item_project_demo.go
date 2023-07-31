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

func NewRandomDemoWorkItem(t *testing.T, pool *pgxpool.Pool, kanbanStepID, workItemTypeID, teamID int) (*db.WorkItem, error) {
	t.Helper()

	dpwiRepo := postgresql.NewDemoWorkItem()

	dpwi, err := dpwiRepo.Create(context.Background(), pool, repos.DemoWorkItemCreateParams{
		DemoProject: db.DemoWorkItemCreateParams{
			// PK is FK. it will be set in repo method after base workitem creation which is unknown beforehand.
			WorkItemID:    -1,
			Ref:           "ref-" + testutil.RandomString(5),
			Line:          "line-" + testutil.RandomString(5),
			Reopened:      testutil.RandomBool(),
			LastMessageAt: testutil.RandomDate(),
		},
		Base: *RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID),
	})
	if err != nil {
		t.Logf("%s", err)

		return nil, err
	}

	return dpwi, nil
}
