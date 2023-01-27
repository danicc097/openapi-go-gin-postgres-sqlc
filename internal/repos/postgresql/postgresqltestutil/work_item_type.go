package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func RandomWorkItemTypeCreateParams(t *testing.T, projectID int) repos.WorkItemTypeCreateParams {
	t.Helper()

	return repos.WorkItemTypeCreateParams{
		Name:        "WorkItemType " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       "#aaaaaa",
	}
}
