#!/usr/bin/env bash

create_args="$(test -n "$with_project" && echo ", projectID")"

# shellcheck disable=SC2028,SC2154
cat <<EOF
package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

$(test -n "$with_project" && echo "	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal\"")
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_Delete${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

$(test -n "$with_project" && echo "		pj := models.ProjectNameDemo
		projectID := internal.ProjectIDByName[pj]")

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   "valid ${sentence_name} deletion",
			status: http.StatusNoContent,
			scopes: []models.Scope{models.Scope${pascal_name}Delete},
		},
		{
			name:   "unauthorized ${sentence_name} call",
			status: http.StatusForbidden,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})
			require.NoError(t, err, "ff.CreateUser: %s")

			${camel_name}f := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{
        $(test -n "$with_project" && echo "		ProjectID: projectID,")
      })
			require.NoError(t, err, "ff.Create${pascal_name}: %s")

			id := ${camel_name}f.${pascal_name}.${pascal_name}ID
			res, err := srv.client.Delete${pascal_name}WithResponse(context.Background() $(test -n "$with_project" && echo ", pj"), id, ReqWithAPIKey(ufixture.APIKey.APIKey))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode(), string(res.Body))
		})
	}
}

func TestHandlers_Create${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

$(test -n "$with_project" && echo "		pj := models.ProjectNameDemo
		projectID := internal.ProjectIDByName[pj]")

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{models.Scope${pascal_name}Create}

		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})


		random${pascal_name}CreateParams := postgresqlrandom.${pascal_name}CreateParams(${create_args#,})
		body := rest.Create${pascal_name}Request{
			${pascal_name}CreateParams: *random${pascal_name}CreateParams,
		}

		res, err := srv.client.Create${pascal_name}WithResponse(context.Background() $(test -n "$with_project" && echo ", pj"), body, ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode(), string(res.Body))
$(for f in ${db_create_params_struct_fields[@]}; do
  # loop db create or update params?
  echo "		assert.EqualValues(t, random${pascal_name}CreateParams.$f, res.JSON201.$f)"
done)

	})
}

func TestHandlers_Get${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

$(test -n "$with_project" && echo "		pj := models.ProjectNameDemo
		projectID := internal.ProjectIDByName[pj]")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{} // no scope needed to read

		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})


		${camel_name}f := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{
      $(test -n "$with_project" && echo "		ProjectID: projectID,")
    })
		require.NoError(t, err, "ff.Create${pascal_name}: %s")

		id := ${camel_name}f.${pascal_name}.${pascal_name}ID
		res, err := srv.client.Get${pascal_name}WithResponse(context.Background() $(test -n "$with_project" && echo ", pj"), id, ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))

		got, err := json.Marshal(res.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&rest.${pascal_name}Response{${pascal_name}: *${camel_name}f.${pascal_name}})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "") // ignore private JSON fields
	})
}

func TestHandlers_Update${pascal_name}(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

$(test -n "$with_project" && echo "		pj := models.ProjectNameDemo
		projectID := internal.ProjectIDByName[pj]")

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name                    string
		status                  int
		body                    rest.Update${pascal_name}Request
		validationErrorContains []string
	}{
		{
			name:   "valid ${sentence_name} update",
			status: http.StatusOK,
			body: func() rest.Update${pascal_name}Request {
				random${pascal_name}CreateParams := postgresqlrandom.${pascal_name}CreateParams(${create_args#,})

				return rest.Update${pascal_name}Request{
					${pascal_name}UpdateParams: models.${pascal_name}UpdateParams{
$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		$f: pointers.New(random${pascal_name}CreateParams.$f),"
done)
					},
				}
			}(),
		},
		// NOTE: we do need to test spec validation
		// {
		// 	name:                    "invalid ${sentence_name} update param",
		// 	status:                  http.StatusBadRequest,
		// 	body:                    rest.Update${pascal_name}Request{},
		// 	validationErrorContains: []string{"[\" field <JSON >\"]", "<error>"},
		// },
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			normalUser := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
				Scopes:     []models.Scope{models.Scope${pascal_name}Edit},
			})
			require.NoError(t, err, "ff.CreateUser: %s")


			${camel_name}f := ff.Create${pascal_name}(context.Background(), servicetestutil.Create${pascal_name}Params{
        $(test -n "$with_project" && echo "		ProjectID: projectID,")
      })
			require.NoError(t, err, "ff.Create${pascal_name}: %s")

			id := ${camel_name}f.${pascal_name}.${pascal_name}ID
			updateRes, err := srv.client.Update${pascal_name}WithResponse(context.Background() $(test -n "$with_project" && echo ", pj"), id, tc.body, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			require.EqualValues(t, tc.status, updateRes.StatusCode(), string(updateRes.Body))

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(updateRes.Body), ve)
				}

				return
			}

			assert.EqualValues(t, id, updateRes.JSON200.${pascal_name}ID)

			res, err := srv.client.Get${pascal_name}WithResponse(context.Background() $(test -n "$with_project" && echo ", pj"), id, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
$(for f in ${db_update_params_struct_fields[@]}; do
  echo "		assert.EqualValues(t, *tc.body.$f, res.JSON200.$f)"
done)
		})
	}
}

EOF
