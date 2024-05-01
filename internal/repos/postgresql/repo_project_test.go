package postgresql_test

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProject_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()

	uniqueTestCases := []filterTestCase[*models.Project]{
		{
			name:       "id",
			filter:     internal.ProjectIDByName[models.ProjectNameDemo],
			repoMethod: reflect.ValueOf(projectRepo.ByID),
			callback: func(t *testing.T, res *models.Project) {
				assert.Equal(t, res.ProjectID, internal.ProjectIDByName[models.ProjectNameDemo])
			},
		},
		{
			name:       "name",
			filter:     models.ProjectNameDemo,
			repoMethod: reflect.ValueOf(projectRepo.ByName),
			callback: func(t *testing.T, res *models.Project) {
				assert.Equal(t, res.ProjectID, internal.ProjectIDByName[models.ProjectNameDemo])
			},
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}

func TestProject_BoardConfigUpdate(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	projectID := internal.ProjectIDByName[models.ProjectNameDemo]

	ctx := context.Background()

	t.Run("valid_subpath_replacement", func(t *testing.T) {
		t.Parallel()

		tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
		require.NoError(t, err)
		defer tx.Rollback(ctx) // rollback errors should be ignored

		const path = "some_path"

		obj := map[string]any{"a": []string{"a.a", "a.b"}}
		err = projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path}, obj)
		require.NoError(t, err)
		p, err := projectRepo.ByID(ctx, tx, projectID)
		require.NoError(t, err)

		got, err := json.Marshal((*p.BoardConfig.Visualization)[path])
		require.NoError(t, err)
		want, err := json.Marshal(obj)
		require.NoError(t, err)

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("board config mismatch (-want +got):\n%s", diff)
		}

		obj2 := map[string]any{"b": "1"}
		err = projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path}, obj2)
		require.NoError(t, err)
		p, err = projectRepo.ByID(ctx, tx, projectID)
		require.NoError(t, err)

		got, err = json.Marshal((*p.BoardConfig.Visualization)[path])
		require.NoError(t, err)
		want, err = json.Marshal(obj2)
		require.NoError(t, err)

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("board config mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("valid_subpath_merge", func(t *testing.T) {
		t.Parallel()

		tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
		require.NoError(t, err)
		defer tx.Rollback(ctx) // rollback errors should be ignored

		const path1 = "some_path"
		const path2 = "another_path"

		obj1 := map[string]any{"a": []string{"a.a", "a.b"}}
		err = projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path1}, obj1)
		require.NoError(t, err)
		obj2 := map[string]any{"b": "1"}
		err = projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path2}, obj2)
		require.NoError(t, err)

		p, err := projectRepo.ByID(ctx, tx, projectID)
		require.NoError(t, err)

		got, err := json.Marshal(p.BoardConfig.Visualization)
		require.NoError(t, err)
		want, err := json.Marshal(map[string]any{path1: obj1, path2: obj2})
		require.NoError(t, err)

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("board config mismatch (-want +got):\n%s", diff)
		}
	})
}
