package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

	srv, _, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	require.NoError(t, err, "Couldn't run test server: %s\n")

	t.Cleanup(func() {
		srv.Close()
	})

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

		req, err := http.NewRequest(http.MethodGet, resttestutil.MustConstructInternalPath("/user/me"), &bytes.Buffer{})
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Add(apiKeyHeaderKey, ufixture.APIKey.APIKey)

		resp := httptest.NewRecorder()

		srv.Handler.ServeHTTP(resp, req)

		ures := User{User: *ufixture.User, Role: role}

		res, err := json.Marshal(ures)
		if err != nil {
			t.Fatalf("could not marshal user fixture")
		}

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, string(res), resp.Body.String())
	})
}

func TestUpdateUserRoutes(t *testing.T) {
	t.Parallel()

	srv, cl, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	require.NoError(t, err, "Couldn't run test server: %s\n")

	t.Cleanup(func() {
		srv.Close()
	})

	ff := newTestFixtureFactory(t)

	t.Run("manager updates another user authorization", func(t *testing.T) {
		t.Parallel()

		scopes := models.Scopes{models.ScopeProjectSettingsWrite}

		manager, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       models.RoleManager,
			WithAPIKey: true,
			Scopes:     scopes,
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

		res, err := cl.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.User.UserID, updateAuthParams, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add(apiKeyHeaderKey, manager.APIKey.APIKey)
			return nil
		})

		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, res.StatusCode())

		ures, err := cl.GetCurrentUserWithResponse(context.Background(), func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add(apiKeyHeaderKey, normalUser.APIKey.APIKey)
			return nil
		})

		require.NoError(t, err)
		assert.Equal(t, *updateAuthParams.Role, ures.JSON200.Role)
	})

	t.Run("user updates itself", func(t *testing.T) {
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

		res, err := cl.UpdateUserWithResponse(context.Background(), normalUser.User.UserID, updateParams, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add(apiKeyHeaderKey, normalUser.APIKey.APIKey)
			return nil
		})

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, normalUser.User.UserID, res.JSON200.UserID)

		ures, err := cl.GetCurrentUserWithResponse(context.Background(), func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add(apiKeyHeaderKey, normalUser.APIKey.APIKey)
			return nil
		})

		require.NoError(t, err)
		assert.Equal(t, updateParams.FirstName, ures.JSON200.FirstName)
		assert.Equal(t, updateParams.LastName, ures.JSON200.LastName)
	})
}
