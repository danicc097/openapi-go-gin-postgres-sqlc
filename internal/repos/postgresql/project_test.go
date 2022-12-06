package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/stretchr/testify/assert"
)

func TestProject_ProjectByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()

	// exists already
	projectName := "dummy project"
	projectID := 1

	type argsString struct {
		filter string
		fn     func(context.Context, db.DBTX, string) (*db.Project, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter: projectName,
				fn:     (projectRepo.ProjectByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundProject, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundProject.ProjectID, projectID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent project"

			_, err := tc.args.fn(context.Background(), testpool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.Project, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "project_id",
			args: argsInt{
				filter: projectID,
				fn:     (projectRepo.ProjectByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundProject, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundProject.ProjectID, projectID)
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

// projects only created in migrations on demand
// func randomProjectCreateParams(t *testing.T) repos.ProjectCreateParams {
// }
