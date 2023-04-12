package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func RandomTeamCreateParams(t *testing.T, projectID int) db.TeamCreateParams {
	t.Helper()

	return db.TeamCreateParams{
		Name:        "Team " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
	}
}
