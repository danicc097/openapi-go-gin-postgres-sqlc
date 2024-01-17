package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func RandomTeamCreateParams(t *testing.T, projectID db.ProjectID) *db.TeamCreateParams {
	t.Helper()

	return &db.TeamCreateParams{
		Name:        "Team " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
	}
}
