package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeam_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	teamRepo := postgresql.NewTeam()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	require.NoError(t, err)

	tcp := postgresqltestutil.RandomTeamCreateParams(t, project.ProjectID)

	team, err := teamRepo.Create(ctx, testPool, tcp)
	require.NoError(t, err)

	uniqueTestCases := []filterTestCase[*db.Team]{
		{
			name:       "name",
			filter:     []any{team.Name, project.ProjectID},
			repoMethod: reflect.ValueOf(teamRepo.ByName),
			callback: func(t *testing.T, res *db.Team) {
				assert.Equal(t, res.Name, team.Name)
			},
		},
		{
			name:       "id",
			filter:     team.TeamID,
			repoMethod: reflect.ValueOf(teamRepo.ByID),
			callback: func(t *testing.T, res *db.Team) {
				assert.Equal(t, res.TeamID, team.TeamID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
