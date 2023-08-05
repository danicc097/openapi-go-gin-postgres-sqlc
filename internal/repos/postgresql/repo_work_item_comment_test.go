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

	uniqueTestCases := []filterTestCase[*db.WorkItemComment]{
		{
			name:       "id",
			filter:     workItemComment.WorkItemCommentID,
			repoMethod: reflect.ValueOf(workItemCommentRepo.ByID),
			callback: func(t *testing.T, res *db.WorkItemComment) {
				assert.Equal(t, res.WorkItemCommentID, workItemComment.WorkItemCommentID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
