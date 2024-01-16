package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/stretchr/testify/require"
)

type CreateTeamParams struct {
	Project models.Project
}

type CreateTeamFixture struct {
	Team *db.Team
}

// CreateTeam creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateTeam(ctx context.Context, params CreateTeamParams) *CreateTeamFixture {
	randomRepoCreateParams := postgresqltestutil.RandomTeamCreateParams(ff.t, internal.ProjectIDByName[params.Project])
	// don't use repos for test fixtures, use service logic
	team, err := ff.svc.Team.Create(ctx, ff.d, randomRepoCreateParams)
	require.NoError(ff.t, err)

	return &CreateTeamFixture{
		Team: team,
	}
}
