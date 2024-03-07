package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeEntry_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	timeEntryRepo := postgresql.NewTimeEntry()

	user := newRandomUser(t, testPool)
	activity := newRandomActivity(t, testPool, models.ProjectDemo)

	workItem := newRandomDemoWorkItem(t, testPool)
	timeEntry := newRandomTimeEntry(t, testPool, activity.ActivityID, user.UserID, &workItem.WorkItemID, nil) // time entry associated to a workItem

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
		ucp := postgresqlrandom.TimeEntryCreateParams(activity.ActivityID, user.UserID, nil, nil)

		_, err := timeEntryRepo.Create(context.Background(), testPool, ucp)
		require.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
