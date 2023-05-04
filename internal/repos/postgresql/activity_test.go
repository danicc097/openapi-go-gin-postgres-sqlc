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

func TestActivity_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	activityRepo := postgresql.NewActivity()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqltestutil.RandomActivityCreateParams(t, project.ProjectID)

	activity, err := activityRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("activityRepo.Create unexpected error = %v", err)
	}

	type argsString struct {
		filter    string
		projectID int
		fn        func(context.Context, db.DBTX, string, int) (*db.Activity, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter:    activity.Name,
				projectID: activity.ProjectID,
				fn:        (activityRepo.ByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundActivity, err := tc.args.fn(context.Background(), testPool, tc.args.filter, tc.args.projectID)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundActivity.ActivityID, activity.ActivityID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := "inexistent activity"

			_, err := tc.args.fn(context.Background(), testPool, filter, tc.args.projectID)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.Activity, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "activity_id",
			args: argsInt{
				filter: activity.ActivityID,
				fn:     (activityRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundActivity, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundActivity.ActivityID, activity.ActivityID)
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

	t.Run("project_id", func(t *testing.T) {
		t.Parallel()

		foundActivity, err := activityRepo.ByProjectID(context.Background(), testPool, activity.ProjectID)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}
		assert.Equal(t, foundActivity[0].ProjectID, activity.ProjectID)
	})

	t.Run("project_id"+" - no rows when record does not exist", func(t *testing.T) {
		t.Parallel()

		filter := 254364 // does not exist

		aa, err := activityRepo.ByProjectID(context.Background(), testPool, filter)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}
		assert.Len(t, aa, 0)
	})
}
