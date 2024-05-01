package postgresql_test

import (
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/stretchr/testify/assert"
)

func TestKanbanStep_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	kanbanStepRepo := postgresql.NewKanbanStep()

	uniqueTestCases := []filterTestCase[*db.KanbanStep]{
		{
			name:       "id",
			filter:     internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived],
			repoMethod: reflect.ValueOf(kanbanStepRepo.ByID),
			callback: func(t *testing.T, res *db.KanbanStep) {
				assert.Equal(t, res.KanbanStepID, internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived])
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}

	nonUniqueTestCases := []filterTestCase[[]db.KanbanStep]{
		{
			name:       "id",
			filter:     internal.ProjectIDByName[models.ProjectNameDemoTwo],
			repoMethod: reflect.ValueOf(kanbanStepRepo.ByProject),
			callback: func(t *testing.T, res []db.KanbanStep) {
				found := false
				for _, ks := range res {
					if ks.KanbanStepID == internal.DemoTwoKanbanStepsIDByName[models.DemoTwoKanbanStepsReceived] {
						found = true

						break
					}
				}
				assert.True(t, found)
				assert.Equal(t, len(internal.DemoTwoKanbanStepsIDByName), len(res))
			},
		},
	}
	for _, tc := range nonUniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
