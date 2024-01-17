package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func RandomActivityCreateParams(t *testing.T, projectID db.ProjectID) *db.ActivityCreateParams {
	t.Helper()

	return &db.ActivityCreateParams{
		Name:         "Activity " + testutil.RandomNameIdentifier(3, "-"),
		Description:  testutil.RandomString(10),
		ProjectID:    projectID,
		IsProductive: testutil.RandomBool(),
	}
}
