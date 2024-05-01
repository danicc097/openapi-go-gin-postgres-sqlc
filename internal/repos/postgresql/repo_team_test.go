package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeam_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	teamRepo := postgresql.NewTeam()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectNameDemo)
	require.NoError(t, err)

	tcp := postgresqlrandom.TeamCreateParams(project.ProjectID)

	team, err := teamRepo.Create(ctx, testPool, tcp)
	require.NoError(t, err)

	uniqueTestCases := []filterTestCase[*models.Team]{
		{
			name:       "name",
			filter:     []any{team.Name, project.ProjectID},
			repoMethod: reflect.ValueOf(teamRepo.ByName),
			callback: func(t *testing.T, res *models.Team) {
				assert.Equal(t, res.Name, team.Name)
			},
		},
		{
			name:       "id",
			filter:     team.TeamID,
			repoMethod: reflect.ValueOf(teamRepo.ByID),
			callback: func(t *testing.T, res *models.Team) {
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

	projectID := internal.ProjectIDByName[models.ProjectNameDemo]

	t.Run("syncs user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
		require.NoError(t, err)
		defer tx.Rollback(ctx) // rollback errors should be ignored

		user := newRandomUser(t, tx)
		team := newRandomTeam(t, tx, projectID)

		_, err = models.CreateUserTeam(ctx, tx, &models.UserTeamCreateParams{
			Member: user.UserID,
			TeamID: team.TeamID,
		})
		require.NoError(t, err)

		_, err = models.UserProjectByMemberProjectID(ctx, tx, user.UserID, projectID) // created by trigger
		require.NoError(t, err)
	})
}

/*
Cannot use current_timestamp when using transactions and column is unique.
Savepoints have no effect, transaction timestamp is the same

create table users (

	user_id serial primary key
	, email text not null unique
	-- use clock instead so it doesn't use transaction time
	--, created_at timestamp with time zone default clock_timestamp() not null unique
	, created_at timestamp with time zone default current_timestamp not null unique

);.

BEGIN;
insert into users (email) values ('mail1');
insert into users (email) values ('mail2'); --fails
COMMIT;.
*/
func TestTriggers_sync_user_teams(t *testing.T) {
	t.Parallel()

	projectID := internal.ProjectIDByName[models.ProjectNameDemo]

	type params struct {
		name      string
		withScope bool
	}

	testCases := []params{
		{
			name:      "user with scope",
			withScope: true,
		},
		{
			name:      "user without scope",
			withScope: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			ctx := context.Background()

			tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
			require.NoError(t, err)
			defer tx.Rollback(ctx) // rollback errors should be ignored

			// IMPORTANT: see note above
			user := newRandomUser(t, tx)

			if tc.withScope {
				user.Scopes = append(user.Scopes, models.ScopeProjectMember)
				user, err = user.Update(ctx, tx)
				require.NoError(t, err)
			}

			previousTeam := newRandomTeam(t, tx, projectID)
			_, err = models.CreateUserTeam(ctx, tx, &models.UserTeamCreateParams{
				Member: user.UserID,
				TeamID: previousTeam.TeamID,
			})
			// FIXME: could not create team: Process 1347 waits for RowExclusiveLock on relation 27813 of database 26116; blocked by process 1401.
			// Process 1401 waits for ShareRowExclusiveLock on relation 27841 of database 26116; blocked by process 1347. | deadlock detected
			require.NoError(t, err)

			team := newRandomTeam(t, tx, projectID) // may trigger user_team update for existing user that is already in project

			_, err = models.UserTeamByMemberTeamID(ctx, tx, user.UserID, previousTeam.TeamID)
			require.NoError(t, err) // was created manually first time to trigger user_project creation

			_, err = models.UserTeamByMemberTeamID(ctx, tx, user.UserID, team.TeamID)
			if tc.withScope {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, postgresql.ParseDBErrorDetail(err), errNoRows)
			}
		})
	}
}
