package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func RandomWorkItemTagCreateParams(t *testing.T, projectID db.ProjectID) *db.WorkItemTagCreateParams {
	t.Helper()

	return &db.WorkItemTagCreateParams{
		Name:        "WorkItemTag " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
