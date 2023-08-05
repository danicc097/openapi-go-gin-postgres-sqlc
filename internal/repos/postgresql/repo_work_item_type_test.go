package postgresql_test

import (
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/stretchr/testify/assert"
)

func TestWorkItemType_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	workItemTypeRepo := postgresql.NewWorkItemType()

	uniqueTestCases := []filterTestCase[*db.WorkItemType]{
		{
			name: "name",
			filter: []any{
				string(models.DemoWorkItemTypesType1),
				internal.ProjectIDByName[models.ProjectDemo],
			},
			repoMethod: reflect.ValueOf(workItemTypeRepo.ByName),
			callback: func(t *testing.T, res *db.WorkItemType) {
				assert.Equal(t, res.Name, string(models.DemoWorkItemTypesType1))
			},
		},
		{
			name:       "id",
			filter:     internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1],
			repoMethod: reflect.ValueOf(workItemTypeRepo.ByID),
			callback: func(t *testing.T, res *db.WorkItemType) {
				assert.Equal(t, res.WorkItemTypeID, internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1])
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}
