package postgresql_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
)

func TestActivity_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	activityRepo := postgresql.NewActivity()

	ctx := t.Context()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectNameDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqlrandom.ActivityCreateParams(project.ProjectID)

	activity, err := activityRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("activityRepo.Create unexpected error = %v", err)
	}

	uniqueTestCases := []filterTestCase[*models.Activity]{
		{
			name: "name",
			filter: []any{
				activity.Name,
				internal.ProjectIDByName[models.ProjectNameDemo],
			},
			repoMethod: reflect.ValueOf(activityRepo.ByName),
			callback: func(t *testing.T, res *models.Activity) {
				assert.Equal(t, res.Name, activity.Name)
			},
		},
		{
			name:       "id",
			filter:     activity.ActivityID,
			repoMethod: reflect.ValueOf(activityRepo.ByID),
			callback: func(t *testing.T, res *models.Activity) {
				assert.Equal(t, res.ActivityID, activity.ActivityID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		runGenericFilterTests(t, tc)
	}

	nonUniqueTestCases := []filterTestCase[[]models.Activity]{
		{
			name:       "project_id",
			filter:     internal.ProjectIDByName[models.ProjectNameDemo],
			repoMethod: reflect.ValueOf(activityRepo.ByProjectID),
			callback: func(t *testing.T, res []models.Activity) {
				assert.Equal(t, res[0].ProjectID, internal.ProjectIDByName[models.ProjectNameDemo])
			},
		},
	}
	for _, tc := range nonUniqueTestCases {
		runGenericFilterTests(t, tc)
	}
}
