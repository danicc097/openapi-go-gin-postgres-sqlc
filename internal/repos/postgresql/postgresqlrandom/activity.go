package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func ActivityCreateParams(projectID db.ProjectID) *db.ActivityCreateParams {
	return &db.ActivityCreateParams{
		Name:         "Activity " + testutil.RandomNameIdentifier(3, "-"),
		Description:  testutil.RandomString(10),
		ProjectID:    projectID,
		IsProductive: testutil.RandomBool(),
	}
}
