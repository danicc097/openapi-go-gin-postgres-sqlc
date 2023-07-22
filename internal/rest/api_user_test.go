package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool)
	srv.cleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	ff := newTestFixtureFactory(t)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleAdvancedUser
		scopes := models.Scopes{models.ScopeProjectSettingsWrite}

		ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		require.NoError(t, err, "ff.CreateUser: %s")

		ures, err := srv.client.GetCurrentUserWithResponse(context.Background(), resttestutil.ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, ures.StatusCode())

		got, err := json.Marshal(ures.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&User{User: *ufixture.User, Role: role})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "")
	})
}

func TestUpdateUserRoutes(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool)
	srv.cleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	ff := newTestFixtureFactory(t)

	scopes := models.Scopes{models.ScopeScopesWrite}

	manager, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleManager,
		WithAPIKey: true,
		Scopes:     scopes,
	})
	require.NoError(t, err, "ff.CreateUser: %s")

	managerWithoutScopes, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleManager,
		WithAPIKey: true,
	})
	require.NoError(t, err, "ff.CreateUser: %s")

	normalUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleUser,
		WithAPIKey: true,
		Scopes:     scopes,
	})
	require.NoError(t, err, "ff.CreateUser: %s")

	// NOTE:
	// scopes and roles part of rest layer. don't test any actual logic here, done in services
	// but we do need to test spec validation

	t.Run("user_authorization", func(t *testing.T) {
		t.Parallel()

		t.Run("valid_update", func(t *testing.T) {
			t.Parallel()

			updateAuthParams := models.UpdateUserAuthRequest{
				Role: pointers.New(models.RoleManager),
			}
			res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, res.StatusCode())

			ures, err := srv.client.GetCurrentUserWithResponse(context.Background(), resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, *updateAuthParams.Role, ures.JSON200.Role)
		})

		t.Run("insufficient_caller_scopes", func(t *testing.T) {
			t.Parallel()

			updateAuthParams := models.UpdateUserAuthRequest{
				Role: pointers.New(models.RoleManager),
			}
			badres, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(managerWithoutScopes.APIKey.APIKey))
			require.NoError(t, err)
			assert.Equal(t, http.StatusForbidden, badres.StatusCode())
		})

		t.Run("invalid_role_update", func(t *testing.T) {
			t.Parallel()
			updateAuthParams := models.UpdateUserAuthRequest{
				Role: pointers.New(models.Role("bad")),
			}
			res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode())
		})

		t.Run("invalid_scopes_update", func(t *testing.T) {
			t.Parallel()
			updateAuthParams := models.UpdateUserAuthRequest{
				Scopes: &[]models.Scope{models.Scope("bad")},
			}
			res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			t.Logf("res.Body: %v\n", string(res.Body))
			assert.Equal(t, http.StatusBadRequest, res.StatusCode())
		})
	})

	tests := []struct {
		name                    string
		status                  int
		body                    models.UpdateUserRequest
		validationErrorContains []string
	}{
		{
			name:   "valid user update",
			status: http.StatusOK,
			body: models.UpdateUserRequest{
				FirstName: pointers.New("new name"),
				LastName:  pointers.New("new name two"),
			},
		},
		{
			name:   "invalid user update param",
			status: http.StatusBadRequest,
			body: models.UpdateUserRequest{
				FirstName: pointers.New("new name"),
				LastName:  pointers.New("new name 43412"),
			},
			validationErrorContains: []string{"[\"lastName\"]", "regular expression"},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			normalUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
			})
			require.NoError(t, err, "ff.CreateUser: %s")

			res, err := srv.client.UpdateUserWithResponse(context.Background(), normalUser.User.UserID, tc.body, resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, tc.status, res.StatusCode())

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(res.Body), ve)
				}

				return
			}

			assert.Equal(t, normalUser.User.UserID, res.JSON200.UserID)

			ures, err := srv.client.GetCurrentUserWithResponse(context.Background(), resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, tc.body.FirstName, ures.JSON200.FirstName)
			assert.Equal(t, tc.body.LastName, ures.JSON200.LastName)
		})
	}
}
