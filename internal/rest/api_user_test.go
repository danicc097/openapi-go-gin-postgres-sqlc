package rest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_DeleteUser(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

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
			name:   "valid user deletion 1",
			status: http.StatusNoContent,
			role:   models.RoleAdmin,
		},
		{
			name:   "valid user deletion 2",
			status: http.StatusNoContent,
			scopes: []models.Scope{models.ScopeUsersDelete},
		},
		{
			name:   "unauthorized user call",
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

			res, err := srv.client.DeleteUserWithResponse(context.Background(), ufixture.UserID.UUID, ReqWithAPIKey(ufixture.APIKey.APIKey))
			fmt.Printf("res.Body: %v\n", string(res.Body))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode(), string(res.Body))
		})
	}
}

func TestHandlers_GetCurrentUser(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleAdvancedUser
		scopes := models.Scopes{models.ScopeProjectSettingsWrite}

		p := models.ProjectNameDemo
		team1f := ff.CreateTeam(context.Background(), servicetestutil.CreateTeamParams{Project: p})
		team2f := ff.CreateTeam(context.Background(), servicetestutil.CreateTeamParams{Project: p})
		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
			TeamIDs:    []models.TeamID{team1f.TeamID, team2f.TeamID},
		})

		res, err := srv.client.GetCurrentUserWithResponse(context.Background(), ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))

		got, err := json.Marshal(res.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&rest.UserResponse{
			User:     ufixture.User,
			Role:     role,
			Teams:    ufixture.MemberTeamsJoin,
			Projects: ufixture.MemberProjectsJoin,
		})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "") // ignore private fields
	})
}

func TestHandlers_UpdateUser(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	// NOTE:
	// scopes and roles part of rest layer. don't test any actual logic here, done in services
	// but we do need to test spec validation

	t.Run("user_authorization", func(t *testing.T) {
		t.Parallel()

		scopes := models.Scopes{models.ScopeScopesWrite}

		manager := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       models.RoleManager,
			WithAPIKey: true,
			Scopes:     scopes,
		})

		normalUser := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       models.RoleUser,
			WithAPIKey: true,
			Scopes:     scopes,
		})

		t.Run("valid_update", func(t *testing.T) {
			t.Parallel()

			updateAuthParams := models.UpdateUserAuthRequest{
				Role:   pointers.New(models.RoleManager),
				Scopes: nil,
			}
			ures, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.UserID.UUID, updateAuthParams, ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, ures.StatusCode(), string(ures.Body))

			res, err := srv.client.GetCurrentUserWithResponse(context.Background(), ReqWithAPIKey(normalUser.APIKey.APIKey))
			require.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))
			require.NoError(t, err)
			assert.EqualValues(t, *updateAuthParams.Role, res.JSON200.Role)
		})

		t.Run("insufficient_caller_scopes", func(t *testing.T) {
			t.Parallel()

			managerWithoutScopes := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleManager,
				WithAPIKey: true,
			})

			updateAuthParams := models.UpdateUserAuthRequest{
				Role:   pointers.New(models.RoleManager),
				Scopes: nil,
			}
			badres, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.UserID.UUID, updateAuthParams, ReqWithAPIKey(managerWithoutScopes.APIKey.APIKey))
			require.NoError(t, err)
			assert.Equal(t, http.StatusForbidden, badres.StatusCode())
		})

		t.Run("invalid_role_update", func(t *testing.T) {
			t.Parallel()
			updateAuthParams := models.UpdateUserAuthRequest{
				Role:   pointers.New(models.Role("bad")),
				Scopes: nil,
			}
			res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.UserID.UUID, updateAuthParams, ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode(), string(res.Body))
		})

		t.Run("invalid_scopes_update", func(t *testing.T) {
			t.Parallel()
			updateAuthParams := models.UpdateUserAuthRequest{
				Scopes: &[]models.Scope{models.Scope("bad")},
				Role:   nil,
			}
			res, err := srv.client.UpdateUserAuthorizationWithResponse(context.Background(), normalUser.UserID.UUID, updateAuthParams, ReqWithAPIKey(manager.APIKey.APIKey))

			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode(), string(res.Body))
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
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			normalUser := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
			})

			ures, err := srv.client.UpdateUserWithResponse(context.Background(), normalUser.UserID.UUID, tc.body, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			require.EqualValues(t, tc.status, ures.StatusCode(), string(ures.Body))

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(ures.Body), ve)
				}

				return
			}

			assert.EqualValues(t, normalUser.UserID, ures.JSON200.UserID)

			res, err := srv.client.GetCurrentUserWithResponse(context.Background(), ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.EqualValues(t, tc.body.FirstName, res.JSON200.FirstName)
			assert.EqualValues(t, tc.body.LastName, res.JSON200.LastName)
		})
	}
}
