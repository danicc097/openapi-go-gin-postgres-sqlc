package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomActivity(t *testing.T, d db.DBTX, project models.Project) *db.Activity {
	t.Helper()

	activityRepo := postgresql.NewActivity()

	// shared between projects, will require one as params.
	ucp := RandomActivityCreateParams(t, internal.ProjectIDByName[project])

	activity, err := activityRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return activity
}

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
