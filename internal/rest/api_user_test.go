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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	srv.cleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	ff := newTestFixtureFactory(t)

	t.Run("authenticated user", func(t *testing.T) {
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

	srv, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	srv.cleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	ff := newTestFixtureFactory(t)

	// scopes and roles part of rest layer. don't test any actual logic here, done in services
	t.Run("user with valid scopes can update another user's authorization info", func(t *testing.T) {
		t.Parallel()

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

		updateAuthParams := models.UpdateUserAuthRequest{
			Role: pointers.New(models.RoleManager),
		}
		badres, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(managerWithoutScopes.APIKey.APIKey))
		require.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, badres.StatusCode())

		res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, resttestutil.ReqWithAPIKey(manager.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, res.StatusCode())

		ures, err := srv.client.GetCurrentUserWithResponse(context.Background(), resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, *updateAuthParams.Role, ures.JSON200.Role)
	})

	t.Run("user can update itself", func(t *testing.T) {
		t.Parallel()

		normalUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       models.RoleUser,
			WithAPIKey: true,
		})
		require.NoError(t, err, "ff.CreateUser: %s")

		updateParams := models.UpdateUserRequest{
			FirstName: pointers.New("new name"),
			LastName:  pointers.New("new name two"),
		}

		res, err := srv.client.UpdateUserWithResponse(context.Background(), normalUser.User.UserID, updateParams, resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, normalUser.User.UserID, res.JSON200.UserID)

		ures, err := srv.client.GetCurrentUserWithResponse(context.Background(), resttestutil.ReqWithAPIKey(normalUser.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, updateParams.FirstName, ures.JSON200.FirstName)
		assert.Equal(t, updateParams.LastName, ures.JSON200.LastName)
	})
}
