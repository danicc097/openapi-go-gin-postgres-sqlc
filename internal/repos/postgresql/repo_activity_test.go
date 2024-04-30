package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	models1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/stretchr/testify/assert"
)

func TestActivity_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	activityRepo := postgresql.NewActivity()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectNameDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqlrandom.ActivityCreateParams(project.ProjectID)

	activity, err := activityRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("activityRepo.Create unexpected error = %v", err)
	}

	uniqueTestCases := []filterTestCase[*models1.Activity]{
		{
			name: "name",
			filter: []any{
				activity.Name,
				internal.ProjectIDByName[models.ProjectNameDemo],
			},
			repoMethod: reflect.ValueOf(activityRepo.ByName),
			callback: func(t *testing.T, res *models1.Activity) {
				assert.Equal(t, res.Name, activity.Name)
			},
		},
		{
			name:       "id",
			filter:     activity.ActivityID,
			repoMethod: reflect.ValueOf(activityRepo.ByID),
			callback: func(t *testing.T, res *models1.Activity) {
				assert.Equal(t, res.ActivityID, activity.ActivityID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}

	nonUniqueTestCases := []filterTestCase[[]models1.Activity]{
		{
			name:       "project_id",
			filter:     internal.ProjectIDByName[models.ProjectNameDemo],
			repoMethod: reflect.ValueOf(activityRepo.ByProjectID),
			callback: func(t *testing.T, res []models1.Activity) {
				assert.Equal(t, res[0].ProjectID, internal.ProjectIDByName[models.ProjectNameDemo])
			},
		},
	}
	for _, tc := range nonUniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
