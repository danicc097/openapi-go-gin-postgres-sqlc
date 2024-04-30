package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	models1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/stretchr/testify/require"
)

type CreateTeamParams struct {
	Project models.ProjectName
}

type CreateTeamFixture struct {
	*models1.Team
}

// CreateTeam creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateTeam(ctx context.Context, params CreateTeamParams) *CreateTeamFixture {
	randomRepoCreateParams := postgresqlrandom.TeamCreateParams(internal.ProjectIDByName[params.Project])
	// don't use repos for test fixtures, use service logic
	team, err := ff.svc.Team.Create(ctx, ff.d, randomRepoCreateParams)
	require.NoError(ff.t, err)

	return &CreateTeamFixture{
		Team: team,
	}
}
