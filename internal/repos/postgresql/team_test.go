package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTeam_TeamByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	teamRepo := postgresql.NewTeam()

	ctx := context.Background()

	project, err := projectRepo.ProjectByName(ctx, testpool, "dummy project")
	if err != nil {
		t.Fatalf("projectRepo.ProjectByName unexpected error = %v", err)
	}
	tcp := randomTeamCreateParams(t, project.ProjectID)

	team, err := teamRepo.Create(ctx, testpool, tcp)
	if err != nil {
		t.Fatalf("teamRepo.Create unexpected error = %v", err)
	}

	type argsString struct {
		filter string
		fn     func(context.Context, db.DBTX, string) (*db.Team, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter: team.Name,
				fn:     (teamRepo.TeamByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundTeam, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundTeam.TeamID, team.TeamID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent team"

			_, err := tc.args.fn(context.Background(), testpool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
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
				fn:     (teamRepo.TeamByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundTeam, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundTeam.TeamID, team.TeamID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := 254364 // does not exist

			_, err := tc.args.fn(context.Background(), testpool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}
}

func randomTeamCreateParams(t *testing.T, projectID int) repos.TeamCreateParams {
	t.Helper()

	return repos.TeamCreateParams{
		Name:        "Team " + testutil.RandomNameIdentifier(3, "-"),
		Description: pointers.New(testutil.RandomString(10)),
		ProjectID:   projectID,
	}
}
