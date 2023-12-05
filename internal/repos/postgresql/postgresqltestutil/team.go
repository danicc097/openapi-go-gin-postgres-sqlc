package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomTeam(t *testing.T, d db.DBTX, projectID db.ProjectID) (*db.Team, error) {
	t.Helper()

	teamRepo := postgresql.NewTeam()

	ucp := RandomTeamCreateParams(t, projectID)

	team, err := teamRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return team, nil
}

func RandomTeamCreateParams(t *testing.T, projectID db.ProjectID) *db.TeamCreateParams {
	t.Helper()

	return &db.TeamCreateParams{
		Name:        "Team " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
	}
}
