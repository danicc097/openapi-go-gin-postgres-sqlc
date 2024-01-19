#!/bin/bash

# shellcheck disable=SC2028,SC2154
delete_method=$(test -n "$has_deleted_at" && echo "SoftDelete" || echo "Delete")
create_args="$(test -n "$with_project" && echo "projectID")"

cat <<EOF
package postgresql_test

import (
	"context"
	"reflect"
	"testing"
	"time"

$(test -n "$with_project" && echo "	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models\"")
$(test -n "$with_project" && echo "	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal\"")
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test${pascal_name}_Update(t *testing.T) {
	t.Parallel()

$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")
	${lower_name} := newRandom${pascal_name}(t, testPool $create_args)

	type args struct {
		id     db.${pascal_name}ID
		params db.${pascal_name}UpdateParams
	}
	type params struct {
		name        string
		args        args
		want        *db.${pascal_name}
		errContains string
	}

	tests := []params{
		{
			name: "updated",
			args: args{
				id:     ${lower_name}.${pascal_name}ID,
				params: db.${pascal_name}UpdateParams{
					// TODO: set fields to update as in crud-api-tests.go.tmpl.bash
				},
			},
			want: func() *db.${pascal_name} {
				u := *${lower_name}
				// TODO: set updated fields to expected values as in crud-api-tests.go.tmpl.bash

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := postgresql.New${pascal_name}()
			got, err := r.Update(context.Background(), testPool, tc.args.id, &tc.args.params)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.ErrorContains(t, err, tc.errContains)

				return
			}

			// NOTE: ignore unwanted fields
			// got.UpdatedAt = want.UpdatedAt

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test${pascal_name}_${delete_method}(t *testing.T) {
	t.Parallel()

	$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")
	${lower_name} := newRandom${pascal_name}(t, testPool $create_args)

	type args struct {
		id db.${pascal_name}ID
	}
	type params struct {
		name        string
		args        args
		errContains string
	}

	tests := []params{
		{
			name: "deleted ${sentence_name} not found",
			args: args{
				id: ${lower_name}.${pascal_name}ID,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			${camel_name}Repo := postgresql.New${pascal_name}()
			_, err := ${camel_name}Repo.Delete(context.Background(), testPool, tc.args.id)
			require.NoError(t, err)

			_, err = ${camel_name}Repo.ByID(context.Background(), testPool, tc.args.id)
			require.ErrorContains(t, err, errNoRows)
			$([[ -z "$has_deleted_at" ]] && echo "/* row was deleted")
			${lower_name}, err = ${camel_name}Repo.ByID(context.Background(), testPool, tc.args.id, db.WithDeleted${pascal_name}Only())
			require.NoError(t, err)
			assert.Equal(t, ${lower_name}.${pascal_name}ID, tc.args.id)
			$([[ -z "$has_deleted_at" ]] && echo "*/")
		})
	}
}

func Test${pascal_name}_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")
	${lower_name} := newRandom${pascal_name}(t, testPool $create_args)

	logger := testutil.NewLogger(t)

	${camel_name}Repo := reposwrappers.New${pascal_name}WithRetry(postgresql.New${pascal_name}(), logger, 10, 65*time.Millisecond)

	uniqueCallback := func(t *testing.T, res *db.${pascal_name}) {
		assert.Equal(t, res.${pascal_name}ID, ${lower_name}.${pascal_name}ID)
	}

	uniqueTestCases := []filterTestCase[*db.${pascal_name}]{
		{
			name:       "id",
			filter:     ${lower_name}.${pascal_name}ID,
			repoMethod: reflect.ValueOf(${camel_name}Repo.ByID),
			callback:   uniqueCallback,
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}

func Test${pascal_name}_Create(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	${camel_name}Repo := reposwrappers.New${pascal_name}WithRetry(postgresql.New${pascal_name}(), logger, 10, 65*time.Millisecond)

	type want struct {
		// NOTE: include db-generated fields here to test equality as well
		db.${pascal_name}CreateParams
	}

	type args struct {
		params db.${pascal_name}CreateParams
	}

	t.Run("correct_${camel_name}", func(t *testing.T) {
		t.Parallel()

$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")
		${camel_name}CreateParams := postgresqlrandom.${pascal_name}CreateParams($create_args)

		want := want{
			${pascal_name}CreateParams: *${camel_name}CreateParams,
		}

		args := args{
			params: *${camel_name}CreateParams,
		}

		got, err := ${camel_name}Repo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

$(for f in ${db_create_params_struct_fields[@]}; do
  echo "		assert.Equal(t, want.$f, got.$f)"
done)
	})

	// implement if needed
	t.Run("check constraint raises violation error", func(t *testing.T) {
		t.Skip("not implemented")
		t.Parallel()

$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")
		${camel_name}CreateParams := postgresqlrandom.${pascal_name}CreateParams($create_args)
		// NOTE: update params to trigger check error

		args := args{
			params: *${camel_name}CreateParams,
		}

		_, err := ${camel_name}Repo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)

		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
EOF
