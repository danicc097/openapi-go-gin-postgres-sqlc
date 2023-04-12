package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}

	t.Cleanup(func() {
		srv.Close()
	})

	ff := newTestFixtureFactory(t)

	t.Run("authenticated user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleAdvancedUser
		scopes := []models.Scope{models.ScopeProjectSettingsWrite}

		ufixture, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}

		req, err := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/user/me", &bytes.Buffer{})
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Add("x-api-key", ufixture.APIKey.APIKey)

		resp := httptest.NewRecorder()

		srv.Handler.ServeHTTP(resp, req)

		ures := UserResponse{User: ufixture.User, Role: role, Scopes: ufixture.User.Scopes}

		res, err := json.Marshal(ures)
		if err != nil {
			t.Fatalf("could not marshal user fixture")
		}

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, string(res), resp.Body.String())
	})
}

func TestUpdateUserRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}

	t.Cleanup(func() {
		srv.Close()
	})

	ff := newTestFixtureFactory(t)

	t.Run("manager updates another user authorization", func(t *testing.T) {
		t.Parallel()

		scopes := []models.Scope{models.ScopeProjectSettingsWrite}

		manager, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
			Role:       models.RoleManager,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}
		normalUser, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
			Role:       models.RoleUser,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}
		var buf bytes.Buffer

		updateAuthParams := models.UpdateUserAuthRequest{
			Role: pointers.New(models.RoleManager),
		}

		if err := json.NewEncoder(&buf).Encode(updateAuthParams); err != nil {
			t.Errorf("unexpected error %v", err)
		}

		path := os.Getenv("API_VERSION") + fmt.Sprintf("/user/%s/authorization", normalUser.User.UserID)
		req, err := http.NewRequest(http.MethodPatch, path, &buf)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("x-api-key", manager.APIKey.APIKey)

		resp := httptest.NewRecorder()

		srv.Handler.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)
	})

	t.Run("user updates itself", func(t *testing.T) {
		t.Parallel()

		normalUser, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
			Role:       models.RoleUser,
			WithAPIKey: true,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}
		var buf bytes.Buffer

		updateParams := models.UpdateUserRequest{
			FirstName: pointers.New("new name"),
			LastName:  pointers.New("new name"),
		}

		if err := json.NewEncoder(&buf).Encode(updateParams); err != nil {
			t.Errorf("unexpected error %v", err)
		}

		path := os.Getenv("API_VERSION") + fmt.Sprintf("/user/%s", normalUser.User.UserID)
		req, err := http.NewRequest(http.MethodPatch, path, &buf)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("x-api-key", normalUser.APIKey.APIKey)

		resp := httptest.NewRecorder()

		srv.Handler.ServeHTTP(resp, req)

		format.PrintJSON(resp)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
