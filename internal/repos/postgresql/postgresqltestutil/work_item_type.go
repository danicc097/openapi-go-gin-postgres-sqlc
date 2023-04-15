package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomWorkItemType(t *testing.T, pool *pgxpool.Pool, projectID int) (*db.WorkItemType, error) {
	t.Helper()

	witRepo := postgresql.NewWorkItemType()

	ucp := RandomWorkItemTypeCreateParams(t, projectID)

	wit, err := witRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		return nil, err
	}

	return wit, nil
}

func RandomWorkItemTypeCreateParams(t *testing.T, projectID int) db.WorkItemTypeCreateParams {
	t.Helper()

	return db.WorkItemTypeCreateParams{
		Name:        "WorkItemType " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
