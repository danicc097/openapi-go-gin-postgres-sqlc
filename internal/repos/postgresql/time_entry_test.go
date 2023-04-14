package postgresql_test

import (
	"context"
	"testing"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
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

	project, err := projectRepo.ByName(ctx, testPool, internalmodels.ProjectDemoProject)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	user := postgresqltestutil.NewRandomUser(t, testPool)
	team := postgresqltestutil.NewRandomTeam(t, testPool, project.ProjectID)
	activity := postgresqltestutil.NewRandomActivity(t, testPool, project.ProjectID)
	workItemType := postgresqltestutil.NewRandomWorkItemType(t, testPool, project.ProjectID)
	kanbanStep := postgresqltestutil.NewRandomKanbanStep(t, testPool, project.ProjectID)
	workItem := postgresqltestutil.NewRandomDemoProjectWorkItem(t, testPool, project.ProjectID, kanbanStep.KanbanStepID, workItemType.WorkItemTypeID, team.TeamID)
	timeEntry := postgresqltestutil.NewRandomTimeEntry(t, testPool, activity.ActivityID, user.UserID, &workItem.WorkItemID, nil) // time entry associated to a workItem

	type argsInt64 struct {
		filter int64
		fn     func(context.Context, db.DBTX, int64) (*db.TimeEntry, error)
	}
	testsInt64 := []struct {
		name string
		args argsInt64
	}{
		{
			name: "timeEntry_id",
			args: argsInt64{
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

			filter := int64(254364) // does not exist

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.Contains(t, err.Error(), errContains)
		})
	}
}
