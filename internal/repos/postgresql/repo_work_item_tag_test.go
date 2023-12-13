package postgresql_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestWorkItemTag_Create(t *testing.T) {
	t.Parallel()

	witRepo := reposwrappers.NewWorkItemTagWithRetry(postgresql.NewWorkItemTag(), 10, 65*time.Millisecond)

	type want struct {
		db.WorkItemTagCreateParams
	}

	type args struct {
		params db.WorkItemTagCreateParams
	}

	t.Run("unique and foreign key violations show user-friendly errors", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomWorkItemTagCreateParams(t, internal.ProjectIDByName[models.ProjectDemo])

		want := want{
			WorkItemTagCreateParams: *ucp,
		}

		args := args{
			params: *ucp,
		}

		got, err := witRepo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.Name, got.Name)
		assert.Equal(t, want.Description, got.Description)
		assert.Equal(t, want.Color, got.Color)
		assert.Equal(t, want.ProjectID, got.ProjectID)

		_, err = witRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)
		require.Error(t, err)

		assert.ErrorContains(t, err, fmt.Sprintf("combination of name=%s and projectID=%d already exists", want.Name, want.ProjectID))

		args.params.ProjectID = -999
		_, err = witRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)
		require.Error(t, err)

		assert.ErrorContains(t, err, fmt.Sprintf("projectID \"%d\" is invalid", args.params.ProjectID))
	})
}
