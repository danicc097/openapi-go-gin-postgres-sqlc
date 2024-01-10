package postgresql_test

import (
	"context"
	"reflect"
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
	user := postgresqltestutil.NewRandomUser(t, testPool)
	team := postgresqltestutil.NewRandomTeam(t, testPool, project.ProjectID)
	activity := postgresqltestutil.NewRandomActivity(t, testPool, project.ProjectID)

	kanbanStepID := internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1]

	workItem := postgresqltestutil.NewRandomDemoWorkItem(t, testPool, kanbanStepID, workItemTypeID, team.TeamID)
	timeEntry := postgresqltestutil.NewRandomTimeEntry(t, testPool, activity.ActivityID, user.UserID, &workItem.WorkItemID, nil) // time entry associated to a workItem

	uniqueTestCases := []filterTestCase[*db.TimeEntry]{
		{
			name:       "id",
			filter:     timeEntry.TimeEntryID,
			repoMethod: reflect.ValueOf(timeEntryRepo.ByID),
			callback: func(t *testing.T, res *db.TimeEntry) {
				assert.Equal(t, res.TimeEntryID, timeEntry.TimeEntryID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}

	t.Run("bad_time_entry_creation", func(t *testing.T) {
		t.Parallel()

		// test num_nonnulls which is repo's responsibility
		ucp := postgresqltestutil.RandomTimeEntryCreateParams(t, activity.ActivityID, user.UserID, nil, nil)

		_, err := timeEntryRepo.Create(context.Background(), testPool, ucp)
		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
