package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/jackc/pgx/v5"
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

func TestTriggers_sync_user_projects(t *testing.T) {
	t.Parallel()

	projectID := internal.ProjectIDByName[models.ProjectDemo]

	t.Run("syncs user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
		defer tx.Rollback(ctx)

		var err error

		user := postgresqltestutil.NewRandomUser(t, tx)
		team := postgresqltestutil.NewRandomTeam(t, tx, projectID)

		_, err = db.CreateUserTeam(ctx, tx, &db.UserTeamCreateParams{
			Member: user.UserID,
			TeamID: team.TeamID,
		})
		require.NoError(t, err)

		_, err = db.UserProjectByMemberProjectID(ctx, tx, user.UserID, projectID) // created by trigger
		require.NoError(t, err)
	})
}

func TestTriggers_sync_user_teams(t *testing.T) {
	t.Parallel()

	projectID := internal.ProjectIDByName[models.ProjectDemo]

	type params struct {
		name      string
		withScope bool
	}

	tests := []params{
		{
			name:      "user with scope",
			withScope: true,
		},
		{
			name:      "user without scope",
			withScope: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			ctx := context.Background()

			tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
			defer tx.Rollback(ctx)

			user := postgresqltestutil.NewRandomUser(t, tx)

			if tc.withScope {
				user.Scopes = append(user.Scopes, models.ScopeProjectMember)
				user, err = user.Update(ctx, tx)
				require.NoError(t, err)
			}

			previousTeam := postgresqltestutil.NewRandomTeam(t, tx, projectID)
			_, err = db.CreateUserTeam(ctx, tx, &db.UserTeamCreateParams{
				Member: user.UserID,
				TeamID: previousTeam.TeamID,
			})
			require.NoError(t, err)

			team := postgresqltestutil.NewRandomTeam(t, tx, projectID) // may trigger user_team update for existing user that is already in project

			_, err = db.UserTeamByMemberTeamID(ctx, tx, user.UserID, previousTeam.TeamID)
			require.NoError(t, err) // was created manually first time to trigger user_project creation

			_, err = db.UserTeamByMemberTeamID(ctx, tx, user.UserID, team.TeamID)
			if tc.withScope {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, postgresql.ParseDBErrorDetail(err), errNoRows)
			}
		})
	}
}
