#!/usr/bin/env bash

create_args="$(test -n "$with_project" && echo "projectID")"

# shellcheck disable=SC2028,SC2154
cat <<EOF
package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test${pascal_name}_Update(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	requiredProject := models.ProjectNameDemo

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	team, err := svc.Team.Create(context.Background(), testPool, postgresqlrandom.TeamCreateParams(internal.ProjectIDByName[requiredProject]))
	require.NoError(t, err)
	creator := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
	})

	creator.User, err = svc.User.AssignTeam(context.Background(), testPool, creator.UserID, team.TeamID)
	require.NoError(t, err)

$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectNameDemo]")

	${camel_name}CreateParams := postgresqlrandom.${pascal_name}CreateParams($create_args)
	${lower_name}, err := svc.${pascal_name}.Create(context.Background(), testPool, ${camel_name}CreateParams)
	require.NoError(t, err)

	type args struct {
		params            *models.${pascal_name}UpdateParams
		id                models.${pascal_name}ID
		withUserInProject bool
	}

	wantParams := postgresqlrandom.${pascal_name}CreateParams($create_args)

	tests := []struct {
		name          string
		args          args
		want          models.${pascal_name}UpdateParams
		errorContains []string
	}{
		{
			name: "updated correctly",
			args: args{
				params: &models.${pascal_name}UpdateParams{
					$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: &wantParams.$f,"
done)
				},
				withUserInProject: false, //
				id:                ${lower_name}.${pascal_name}ID,
			},
			want: models.${pascal_name}UpdateParams{
				// generating fields based on randomized CreateParams since it's a superset of updateparams.
				$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: &wantParams.$f,"
done)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos(t)
			repos.Notification = repostesting.NewFakeNotification() // unless we want to test notification integration

			ctx := context.Background()

			tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
			require.NoError(t, err)
			defer tx.Rollback(ctx) // rollback errors should be ignored

			user := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})

			if tc.args.withUserInProject {
				user.User, err = svc.User.AssignTeam(context.Background(), testPool, user.UserID, team.TeamID)
				require.NoError(t, err)
			}

			w := services.New${pascal_name}(logger, repos)
			got, err := w.Update(ctx, tx, tc.args.id, tc.args.params)

			if (err != nil) && len(tc.errorContains) == 0 {
				t.Fatalf("unexpected error = %v", err)
			}

			if len(tc.errorContains) > 0 {
				for _, ve := range tc.errorContains {
					require.ErrorContains(t, err, ve)
				}

				return
			}

			// assert.Equal(t, *tc.want.<Field>, got.<Field>)
			$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		assert.Equal(t, *tc.want.$f, got.$f)"
done)
		})
	}
}
EOF
