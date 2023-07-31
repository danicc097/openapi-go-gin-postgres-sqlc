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

func TestWorkItemComment_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	workItemCommentRepo := postgresql.NewWorkItemComment()

	ctx := context.Background()

	projectID := internal.ProjectIDByName[models.ProjectDemo]
	team, _ := postgresqltestutil.NewRandomTeam(t, testPool, projectID)

	kanbanStepID := internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1]
	demoWorkItem, _ := postgresqltestutil.NewRandomDemoWorkItem(t, testPool, kanbanStepID, workItemTypeID, team.TeamID)

	user, _ := postgresqltestutil.NewRandomUser(t, testPool)

	wiccp := postgresqltestutil.RandomWorkItemCommentCreateParams(t, demoWorkItem.WorkItemID, user.UserID)

	workItemComment, err := workItemCommentRepo.Create(ctx, testPool, wiccp)
	if err != nil {
		t.Fatalf("workItemCommentRepo.Create unexpected error = %v", err)
	}

	type argsInt64 struct {
		filter int64
		fn     func(context.Context, db.DBTX, int64) (*db.WorkItemComment, error)
	}
	testsInt := []struct {
		name string
		args argsInt64
	}{
		{
			name: "workItemComment_id",
			args: argsInt64{
				filter: workItemComment.WorkItemCommentID,
				fn:     (workItemCommentRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundWorkItemComment, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundWorkItemComment.WorkItemCommentID, workItemComment.WorkItemCommentID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := int64(254364) // does not exist

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}
}
