package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/stretchr/testify/assert"
)

func TestWorkItemType_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	workItemTypeRepo := postgresql.NewWorkItemType()

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
				filter:    string(models.DemoWorkItemTypesType1), // work item types table shared by all
				projectID: internal.ProjectIDByName[models.ProjectDemo],
				fn:        (workItemTypeRepo.ByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemType, err := tc.args.fn(context.Background(), testPool, tc.args.filter, tc.args.projectID)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemType.Name, string(models.DemoWorkItemTypesType1))
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent workItemType"

			_, err := tc.args.fn(context.Background(), testPool, filter, tc.args.projectID)
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
				filter: internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1],
				fn:     (workItemTypeRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemType, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemType.WorkItemTypeID, internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1])
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := 254364 // does not exist

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}
}
