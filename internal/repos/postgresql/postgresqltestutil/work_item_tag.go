package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomWorkItemTag(t *testing.T, d db.DBTX, projectID db.ProjectID) (*db.WorkItemTag, error) {
	t.Helper()

	witRepo := postgresql.NewWorkItemTag()

	ucp := RandomWorkItemTagCreateParams(t, projectID)

	wit, err := witRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return wit, nil
}

func RandomWorkItemTagCreateParams(t *testing.T, projectID db.ProjectID) *db.WorkItemTagCreateParams {
	t.Helper()

	return &db.WorkItemTagCreateParams{
		Name:        "WorkItemTag " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
