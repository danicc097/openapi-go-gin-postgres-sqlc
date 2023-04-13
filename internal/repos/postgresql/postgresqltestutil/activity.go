package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomActivity(t *testing.T, pool *pgxpool.Pool, projectID int) *db.Activity {
	t.Helper()

	activityRepo := postgresql.NewActivity()

	ucp := RandomActivityCreateParams(t, projectID)

	activity, err := activityRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	return activity
}

func RandomActivityCreateParams(t *testing.T, projectID int) db.ActivityCreateParams {
	t.Helper()

	return db.ActivityCreateParams{
		Name:         "Activity " + testutil.RandomNameIdentifier(3, "-"),
		Description:  testutil.RandomString(10),
		ProjectID:    projectID,
		IsProductive: testutil.RandomBool(),
	}
}
