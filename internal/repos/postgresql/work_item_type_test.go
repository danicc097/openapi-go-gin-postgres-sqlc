package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestWorkItemType_WorkItemTypeByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	workItemTypeRepo := postgresql.NewWorkItemType()

	ctx := context.Background()
	project, err := projectRepo.ProjectByName(ctx, testpool, demoProjectName)
	if err != nil {
		t.Fatalf("projectRepo.ProjectByName unexpected error = %v", err)
	}
	tcp := randomWorkItemTypeCreateParams(t, project.ProjectID)

	workItemType, err := workItemTypeRepo.Create(ctx, testpool, tcp)
	if err != nil {
		t.Fatalf("workItemTypeRepo.Create unexpected error = %v", err)
	}

	type argsString struct {
		filter    string
		projectID int
		fn        func(context.Context, db.DBTX, string, int) (*db.WorkItemType, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter:    workItemType.Name,
				projectID: workItemType.ProjectID,
				fn:        (workItemTypeRepo.WorkItemTypeByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemType, err := tc.args.fn(context.Background(), testpool, tc.args.filter, tc.args.projectID)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemType.WorkItemTypeID, workItemType.WorkItemTypeID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent workItemType"

			_, err := tc.args.fn(context.Background(), testpool, filter, tc.args.projectID)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.WorkItemType, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "workItemType_id",
			args: argsInt{
				filter: workItemType.WorkItemTypeID,
				fn:     (workItemTypeRepo.WorkItemTypeByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemType, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemType.WorkItemTypeID, workItemType.WorkItemTypeID)
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

func randomWorkItemTypeCreateParams(t *testing.T, projectID int) repos.WorkItemTypeCreateParams {
	t.Helper()

	return repos.WorkItemTypeCreateParams{
		Name:        "WorkItemType " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
