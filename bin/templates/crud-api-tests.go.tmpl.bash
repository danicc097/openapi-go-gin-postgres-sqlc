# shellcheck disable=SC2028,SC2154
echo "package rest_test

import (
	\"context\"
	\"encoding/json\"
	\"fmt\"
	\"net/http\"
	\"testing\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers\"
	\"github.com/stretchr/testify/assert\"
	\"github.com/stretchr/testify/require\"
	\"go.uber.org/zap/zaptest\"
)

func TestHandlers_Delete${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, \"Couldn't run test server: %s\n\")

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   \"valid ${sentence_name} deletion\",
			status: http.StatusNoContent,
			scopes: []models.Scope{models.Scope${pascal_name}Delete},
		},
		{
			name:   \"unauthorized ${sentence_name} call\",
			status: http.StatusForbidden,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})
			require.NoError(t, err, \"ff.CreateUser: %s\")

			${camel_name}, err := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{})
			require.NoError(t, err, \"ff.CreateUser: %s\")

			id := ${camel_name}.${pascal_name}.${pascal_name}ID
			res, err := srv.client.Delete${pascal_name}WithResponse(context.Background(), int(id), ReqWithAPIKey(ufixture.APIKey.APIKey))
			fmt.Printf(\"res.Body: %v\n\", string(res.Body))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode())
		})
	}
}

func TestHandlers_Get${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, \"Couldn't run test server: %s\n\")

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run(\"authenticated_user\", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{models.ScopeProjectSettingsWrite}

		ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		require.NoError(t, err, \"ff.CreateUser: %s\")

		${camel_name}, err := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{})
		require.NoError(t, err, \"ff.CreateUser: %s\")

		id := ${camel_name}.${pascal_name}.${pascal_name}ID
		res, err := srv.client.Get${pascal_name}WithResponse(context.Background(), int(id), ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())

		got, err := json.Marshal(res.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&rest.${pascal_name}{${pascal_name}: *${camel_name}.${pascal_name}})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), \"\")
	})
}

func TestHandlers_Update${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, \"Couldn't run test server: %s\n\")

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name                    string
		status                  int
		body                    rest.Update${pascal_name}Request
		validationErrorContains []string
	}{
		{
			name:   \"valid ${sentence_name} update\",
			status: http.StatusOK,
			body: func() rest.Update${pascal_name}Request {
				random${pascal_name}CreateParams := postgresqltestutil.Random${pascal_name}CreateParams(t)

				return rest.Update${pascal_name}Request{
					${pascal_name}UpdateParams: db.${pascal_name}UpdateParams{
$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: pointers.New(random${pascal_name}CreateParams.$f),"
done)
					},
				}
			}(),
		},
		// NOTE: we do need to test spec validation
		// {
		// 	name:                    \"invalid ${sentence_name} update param\",
		// 	status:                  http.StatusBadRequest,
		// 	body:                    rest.Update${pascal_name}Request{},
		// 	validationErrorContains: []string{\"[\\\" field <JSON >\\\"]\", \"<error>\"},
		// },
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			normalUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
			})
			require.NoError(t, err, \"ff.CreateUser: %s\")

			${camel_name}, err := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{})
			require.NoError(t, err, \"ff.CreateUser: %s\")

			id := ${camel_name}.${pascal_name}.${pascal_name}ID
			updateRes, err := srv.client.Update${pascal_name}WithResponse(context.Background(), int(id), tc.body, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			require.EqualValues(t, tc.status, updateRes.StatusCode())

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(updateRes.Body), ve)
				}

				return
			}

			assert.EqualValues(t, id, updateRes.JSON200.${pascal_name}ID)

			res, err := srv.client.Get${pascal_name}WithResponse(context.Background(), int(id), ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		assert.EqualValues(t, tc.body.$f, res.JSON200.$f)"
done)
		})
	}
}

"
