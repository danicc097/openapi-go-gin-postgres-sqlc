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

func TestKanbanStep_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	kanbanStepRepo := postgresql.NewKanbanStep()

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.KanbanStep, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "kanbanStep_id",
			args: argsInt{
				filter: internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived],
				fn:     (kanbanStepRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundKanbanStep, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundKanbanStep.KanbanStepID, internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived])
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

	type argsIntNotUnique struct {
		filter int
		fn     func(context.Context, db.DBTX, int) ([]db.KanbanStep, error)
	}
	testsIntNotUnique := []struct {
		name string
		args argsIntNotUnique
	}{
		{
			name: "project_id",
			args: argsIntNotUnique{
				filter: internal.ProjectIDByName[models.ProjectDemoTwo],
				fn:     (kanbanStepRepo.ByProject),
			},
		},
	}
	for _, tc := range testsIntNotUnique {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundKanbanSteps, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			found := false
			for _, ks := range foundKanbanSteps {
				if ks.KanbanStepID == internal.DemoTwoKanbanStepsIDByName[models.DemoTwoKanbanStepsReceived] {
					found = true
					break
				}
			}
			assert.True(t, found)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			filter := 254364 // does not exist

			foundKanbanSteps, err := tc.args.fn(context.Background(), testPool, filter)
			if err != nil {
				t.Fatalf("unexpected error = '%v'", err)
			}
			assert.Len(t, foundKanbanSteps, 0)
		})
	}
}
