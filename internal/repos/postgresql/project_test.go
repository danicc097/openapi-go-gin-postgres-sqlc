package postgresql_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProject_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()

	// exists already
	projectID := internal.ProjectIDByName[models.ProjectDemo]

	type argsString struct {
		filter models.Project
		fn     func(context.Context, db.DBTX, models.Project) (*db.Project, error)
	}

	testString := []struct {
		name string
		args argsString
	}{
		{
			name: "name",
			args: argsString{
				filter: models.ProjectDemo,
				fn:     (projectRepo.ByName),
			},
		},
	}
	for _, tc := range testString {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundProject, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundProject.ProjectID, projectID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := models.Project("inexistent project")

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}

	type argsInt struct {
		filter int
		fn     func(context.Context, db.DBTX, int) (*db.Project, error)
	}
	testsInt := []struct {
		name string
		args argsInt
	}{
		{
			name: "project_id",
			args: argsInt{
				filter: projectID,
				fn:     (projectRepo.ByID),
			},
		},
	}
	for _, tc := range testsInt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			foundProject, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			assert.Equal(t, foundProject.ProjectID, projectID)
		})

		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
			t.Parallel()

			errContains := errNoRows

			filter := 254364 // does not exist

			_, err := tc.args.fn(context.Background(), testPool, filter)
			if err == nil {
				t.Fatalf("expected error = '%v' but got nothing", errContains)
			}
			assert.ErrorContains(t, err, errContains)
		})
	}
}

func TestProject_BoardConfigUpdate(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()
	projectID := internal.ProjectIDByName[models.ProjectDemo]

	ctx := context.Background()

	t.Run("valid_subpath_replacement", func(t *testing.T) {
		t.Parallel()

		tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
		defer tx.Rollback(ctx)

		const path = "some_path"

		obj := map[string]any{"a": []string{"a.a", "a.b"}}
		err := projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path}, obj)
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

		tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
		defer tx.Rollback(ctx)

		const path1 = "some_path"
		const path2 = "another_path"

		obj1 := map[string]any{"a": []string{"a.a", "a.b"}}
		err := projectRepo.UpdateBoardConfig(ctx, tx, projectID, []string{"visualization", path1}, obj1)
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
