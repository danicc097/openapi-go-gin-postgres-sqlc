create_args="$(test -n "$with_project" && echo ", projectID")"

# shellcheck disable=SC2028,SC2154
echo "package services_test

import (
	\"context\"
	\"testing\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil\"
	\"github.com/jackc/pgx/v5\"
	\"github.com/stretchr/testify/assert\"
	\"github.com/stretchr/testify/require\"
	\"go.uber.org/zap/zaptest\"
)

func Test${pascal_name}_Update(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	requiredProject := models.ProjectDemo

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	team, err := svc.Team.Create(context.Background(), testPool, postgresqltestutil.RandomTeamCreateParams(t, internal.ProjectIDByName[requiredProject]))
	require.NoError(t, err)
	tagCreator, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
	})
	require.NoError(t, err)

	err = svc.User.AssignTeam(context.Background(), testPool, tagCreator.User.UserID, team.TeamID)
	require.NoError(t, err)

$(test -n "$with_project" && echo "	projectID := internal.ProjectIDByName[models.ProjectDemo]")

	${camel_name}CreateParams := postgresqltestutil.Random${pascal_name}CreateParams(t $create_args)
	${lower_name}, err := svc.${pascal_name}.Create(context.Background(), testPool, ${camel_name}CreateParams)
	require.NoError(t, err)

	type args struct {
		params            *db.${pascal_name}UpdateParams
		id                db.${pascal_name}ID
		withUserInProject bool
	}

	random${pascal_name}CreateParams := postgresqltestutil.Random${pascal_name}CreateParams(t $create_args)

	tests := []struct {
		name          string
		args          args
		want          db.${pascal_name}UpdateParams
		errorContains []string
	}{
		{
			name: \"updated correctly\",
			args: args{
				params: &db.${pascal_name}UpdateParams{
			$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: &random${pascal_name}CreateParams.$f,"
done)
				},
				withUserInProject: false, //
				id:                ${lower_name}.${pascal_name}ID,
			},
			want: db.${pascal_name}UpdateParams{
				// generating fields based on randomized createparams since it's a superset of updateparams.
			$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: &random${pascal_name}CreateParams.$f,"
done)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos()
			repos.Notification = repostesting.NewFakeNotification() // unless we want to test notification integration

			ctx := context.Background()
			tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
			defer tx.Rollback(ctx)

			user, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})
			require.NoError(t, err)

			if tc.args.withUserInProject {
				err = svc.User.AssignTeam(context.Background(), testPool, user.User.UserID, team.TeamID)
				require.NoError(t, err)
			}

			w := services.New${pascal_name}(logger, repos)
			got, err := w.Update(ctx, tx, tc.args.id, tc.args.params)

			if (err != nil) && len(tc.errorContains) == 0 {
				t.Fatalf(\"unexpected error = %v\", err)
			}

			if len(tc.errorContains) > 0 {
				for _, ve := range tc.errorContains {
					assert.ErrorContains(t, err, ve)
				}

				return
			}

			// loop all fields like in above
			// assert.Equal(t, *tc.want.<Field>, got.<Field>)
			$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		assert.Equal(t, *tc.want.$f, got.$f)"
done)
		})
	}
}
"
