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

func TestKanbanStep_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	kanbanStepRepo := postgresql.NewKanbanStep()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, internalmodels.ProjectDemoProject)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqltestutil.RandomKanbanStepCreateParams(t, project.ProjectID)
	kanbanStep, err := kanbanStepRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("kanbanStepRepo.Create unexpected error = %v", err)
	}

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
				filter: kanbanStep.KanbanStepID,
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
			assert.Equal(t, foundKanbanStep.KanbanStepID, kanbanStep.KanbanStepID)
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
				filter: kanbanStep.ProjectID,
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
				if ks.KanbanStepID == kanbanStep.KanbanStepID {
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
