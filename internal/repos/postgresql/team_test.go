package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/stretchr/testify/assert"
)

func TestTeam_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	teamRepo := postgresql.NewTeam()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqltestutil.RandomTeamCreateParams(t, project.ProjectID)

	team, err := teamRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("teamRepo.Create unexpected error = %v", err)
	}

	type argsString struct {
		filter    string
		projectID int
		fn        func(context.Context, db.DBTX, string, int) (*db.Team, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter:    team.Name,
				projectID: team.ProjectID,
				fn:        (teamRepo.ByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundTeam, err := tc.args.fn(context.Background(), testPool, tc.args.filter, tc.args.projectID)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundTeam.TeamID, team.TeamID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent team"

			_, err := tc.args.fn(context.Background(), testPool, filter, tc.args.projectID)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.Team, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "team_id",
			args: argsInt{
				filter: team.TeamID,
				fn:     (teamRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundTeam, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundTeam.TeamID, team.TeamID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := 254364 // does not exist

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}
}
