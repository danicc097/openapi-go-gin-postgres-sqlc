package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomTeam(t *testing.T, pool *pgxpool.Pool, projectID int) (*db.Team, error) {
	t.Helper()

	teamRepo := postgresql.NewTeam()

	ucp := RandomTeamCreateParams(t, projectID)

	team, err := teamRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		t.Logf("%s", err)
		return nil, err
	}

	return team, nil
}

func RandomTeamCreateParams(t *testing.T, projectID int) db.TeamCreateParams {
	t.Helper()

	return db.TeamCreateParams{
		Name:        "Team " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
	}
}
