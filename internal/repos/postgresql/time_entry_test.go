package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/stretchr/testify/assert"
)

func TestTimeEntry_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	timeEntryRepo := postgresql.NewTimeEntry()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	user, _ := postgresqltestutil.NewRandomUser(t, testPool)
	team, _ := postgresqltestutil.NewRandomTeam(t, testPool, project.ProjectID)
	activity, _ := postgresqltestutil.NewRandomActivity(t, testPool, project.ProjectID)

	kanbanStepID := internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1]

	workItem, _ := postgresqltestutil.NewRandomDemoWorkItem(t, testPool, kanbanStepID, workItemTypeID, team.TeamID)
	timeEntry, _ := postgresqltestutil.NewRandomTimeEntry(t, testPool, activity.ActivityID, user.UserID, &workItem.WorkItemID, nil) // time entry associated to a workItem

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int, ...db.TimeEntrySelectConfigOption) (*db.TimeEntry, error)
	}
	testsInt64 := []struct {
		name string
		args argsInt
	}{
		{
			name: "timeEntry_id",
			args: argsInt{
				filter: timeEntry.TimeEntryID,
				fn:     (timeEntryRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt64 {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundTimeEntry, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundTimeEntry.TimeEntryID, timeEntry.TimeEntryID)
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

	t.Run("bad_time_entry_creation", func(t *testing.T) {
		t.Parallel()

		// test num_nonnulls which is repo's responsibility
		ucp := postgresqltestutil.RandomTimeEntryCreateParams(t, activity.ActivityID, user.UserID, nil, nil)

		_, err := timeEntryRepo.Create(context.Background(), testPool, ucp)
		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
