package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/require"
)

func NewRandomDemoTwoWorkItem(t *testing.T, d db.DBTX, kanbanStepID, workItemTypeID, teamID int) (*db.WorkItem, error) {
	t.Helper()

	dpwiRepo := postgresql.NewDemoTwoWorkItem()

	dpwi, err := dpwiRepo.Create(context.Background(), d, repos.DemoTwoWorkItemCreateParams{
		DemoTwoProject: db.DemoTwoWorkItemCreateParams{
			// PK is FK. it will be set in repo method after base workitem creation which is unknown beforehand.
			WorkItemID:            -1,
			CustomDateForProject2: pointers.New(testutil.RandomDate()),
		},
		Base: *RandomWorkItemCreateParams(t, kanbanStepID, workItemTypeID, teamID),
	})
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing failures use random create params instead

	return dpwi, nil
}
