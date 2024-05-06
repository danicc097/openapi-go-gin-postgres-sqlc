package postgresql_test

import (
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/stretchr/testify/assert"
)

func TestWorkItemTag_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	workItemTagRepo := postgresql.NewWorkItemTag()
	workItemTag := newRandomWorkItemTag(t, testPool, internal.ProjectIDByName[models.ProjectNameDemo])

	uniqueTestCases := []filterTestCase[*models.WorkItemTag]{
		{
			name: "name",
			filter: []any{
				workItemTag.Name,
				internal.ProjectIDByName[models.ProjectNameDemo],
			},
			repoMethod: reflect.ValueOf(workItemTagRepo.ByName),
			callback: func(t *testing.T, res *models.WorkItemTag) {
				assert.Equal(t, res.WorkItemTagID, workItemTag.WorkItemTagID)
			},
		}, {
			name:       "id",
			filter:     workItemTag.WorkItemTagID,
			repoMethod: reflect.ValueOf(workItemTagRepo.ByID),
			callback: func(t *testing.T, res *models.WorkItemTag) {
				assert.Equal(t, res.WorkItemTagID, workItemTag.WorkItemTagID)
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
