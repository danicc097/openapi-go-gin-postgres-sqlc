package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomWorkItemTag(t *testing.T, pool *pgxpool.Pool, projectID int) (*db.WorkItemTag, error) {
	t.Helper()

	witRepo := postgresql.NewWorkItemTag()

	ucp := RandomWorkItemTagCreateParams(t, projectID)

	wit, err := witRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		return nil, err
	}

	return wit, nil
}

func RandomWorkItemTagCreateParams(t *testing.T, projectID int) db.WorkItemTagCreateParams {
	t.Helper()

	return db.WorkItemTagCreateParams{
		Name:        "WorkItemTag " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
