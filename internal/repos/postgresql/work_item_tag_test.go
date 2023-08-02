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

func TestWorkItemTag_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	workItemTagRepo := postgresql.NewWorkItemTag()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqltestutil.RandomWorkItemTagCreateParams(t, project.ProjectID)

	workItemTag, err := workItemTagRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("workItemTagRepo.Create unexpected error = %v", err)
	}

	type argsString struct {
		filter    string
		projectID int
		fn        func(context.Context, db.DBTX, string, int, ...db.WorkItemTagSelectConfigOption) (*db.WorkItemTag, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter:    workItemTag.Name,
				projectID: workItemTag.ProjectID,
				fn:        (workItemTagRepo.ByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemTag, err := tc.args.fn(context.Background(), testPool, tc.args.filter, tc.args.projectID)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemTag.WorkItemTagID, workItemTag.WorkItemTagID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent workItemTag"

			_, err := tc.args.fn(context.Background(), testPool, filter, tc.args.projectID)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int, ...db.WorkItemTagSelectConfigOption) (*db.WorkItemTag, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "workItemTag_id",
			args: argsInt{
				filter: workItemTag.WorkItemTagID,
				fn:     (workItemTagRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemTag, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemTag.WorkItemTagID, workItemTag.WorkItemTagID)
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
