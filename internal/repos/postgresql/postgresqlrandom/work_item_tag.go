package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

// NOTE: FKs should always be passed explicitly.
func WorkItemTagCreateParams(projectID models.ProjectID) *models.WorkItemTagCreateParams {
	return &models.WorkItemTagCreateParams{
		Name:        "WorkItemTag " + testutil.RandomNameIdentifier(3, "-"),
		Description: testutil.RandomString(10),
		ProjectID:   projectID,
		Color:       testutil.RandomHEXColor(),
	}
}
