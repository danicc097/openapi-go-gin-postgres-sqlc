package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomDemoWorkItem(t *testing.T, d db.DBTX, kanbanStepID db.KanbanStepID, workItemTypeID db.WorkItemTypeID, teamID db.TeamID) *db.WorkItem {
	t.Helper()

	dpwiRepo := postgresql.NewDemoWorkItem()

	cp := RandomDemoWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID)
	dpwi, err := dpwiRepo.Create(context.Background(), d, cp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return dpwi
}

// NOTE: FKs should always be passed explicitly.
func RandomDemoWorkItemCreateParams(t *testing.T, kanbanStepID db.KanbanStepID, workItemTypeID db.WorkItemTypeID, teamID db.TeamID) repos.DemoWorkItemCreateParams {
	return repos.DemoWorkItemCreateParams{
		DemoProject: db.DemoWorkItemCreateParams{
			WorkItemID:    db.WorkItemID(-1),
			Ref:           "ref-" + testutil.RandomString(5),
			Line:          "line-" + testutil.RandomString(5),
			Reopened:      testutil.RandomBool(),
			LastMessageAt: testutil.RandomDate(),
		},
		Base: *RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID),
	}
}
