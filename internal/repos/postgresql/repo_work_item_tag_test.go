package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/stretchr/testify/assert"
)

func TestWorkItemTag_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	workItemTagRepo := postgresql.NewWorkItemTag()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	tcp := postgresqltestutil.RandomWorkItemTagCreateParams(t, project.ProjectID)

	workItemTag, err := workItemTagRepo.Create(ctx, testPool, tcp)
	if err != nil {
		t.Fatalf("workItemTagRepo.Create unexpected error = %v", err)
	}

	uniqueTestCases := []filterTestCase[*db.WorkItemTag]{
		{
			name: "name",
			filter: []any{
				workItemTag.Name,
				internal.ProjectIDByName[models.ProjectDemo],
			},
			repoMethod: reflect.ValueOf(workItemTagRepo.ByName),
			callback: func(t *testing.T, res *db.WorkItemTag) {
				assert.Equal(t, res.WorkItemTagID, workItemTag.WorkItemTagID)
			},
		}, {
			name:       "id",
			filter:     workItemTag.WorkItemTagID,
			repoMethod: reflect.ValueOf(workItemTagRepo.ByID),
			callback: func(t *testing.T, res *db.WorkItemTag) {
				assert.Equal(t, res.WorkItemTagID, workItemTag.WorkItemTagID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
