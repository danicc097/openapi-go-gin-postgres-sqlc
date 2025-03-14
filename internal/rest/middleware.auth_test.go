package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorizationMiddleware(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		role           models.Role
		requiredRole   models.Role
		scopes         models.Scopes
		requiredScopes models.Scopes
		status         int
		body           string
		withoutUser    bool
	}{
		{
			name:         "unauthorized_user",
			role:         models.RoleUser,
			requiredRole: models.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "Unauthorized",
		},
		{
			name:         "unauthorized_manager",
			role:         models.RoleManager,
			requiredRole: models.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "Unauthorized",
		},
		{
			name:         "authorized",
			role:         models.RoleAdmin,
			requiredRole: models.RoleAdmin,
			status:       http.StatusOK,
			body:         "ok",
		},
		{
			name:           "authorized_with_missing_scopes_but_valid_rank",
			role:           models.RoleAdmin,
			requiredScopes: []models.Scope{models.ScopeUsersWrite},
			requiredRole:   models.RoleAdmin,
			status:         http.StatusOK,
			body:           "ok",
		},
		{
			name:           "authorized_with_missing_role_but_valid_scopes",
			role:           models.RoleUser,
			scopes:         []models.Scope{models.ScopeUsersWrite},
			requiredScopes: []models.Scope{models.ScopeUsersWrite},
			requiredRole:   models.RoleAdmin,
			status:         http.StatusOK,
			body:           "ok",
		},
		{
			name:           "unauthorized_with_missing_scope",
			scopes:         []models.Scope{models.ScopeUsersWrite},
			requiredScopes: []models.Scope{models.ScopeUsersWrite, models.ScopeScopesWrite},
			status:         http.StatusForbidden,
			body:           "unauthorized: scope(s) users:write and scopes:write is required",
		},
		{
			name:           "unauthorized_with_missing_scopes_and_role",
			scopes:         []models.Scope{models.ScopeUsersWrite},
			requiredScopes: []models.Scope{models.ScopeUsersWrite, models.ScopeScopesWrite},
			role:           models.RoleUser,
			requiredRole:   models.RoleAdmin,
			status:         http.StatusForbidden,
			body:           "unauthorized: either role admin or scope(s) users:write and scopes:write are required",
		},
		{
			name:        "unauthorized_if_no_user",
			withoutUser: true,
			status:      http.StatusInternalServerError,
		},
	}

	logger := testutil.NewLogger(t)

	svcs := services.New(logger, services.CreateTestRepos(t), testPool)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := httptest.NewRecorder()
			_, engine := gin.CreateTestContext(res)

			authMw := rest.NewAuthMiddleware(logger, testPool, svcs)

			ff := servicetestutil.NewFixtureFactory(t, testPool, svcs)
			ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				Scopes:     tc.scopes,
				WithAPIKey: true,
			})

			if !tc.withoutUser {
				engine.Use(func(c *gin.Context) {
					rest.CtxWithUserCaller(c, ufixture.User)
				})
			}

			engine.Use(authMw.EnsureAuthorized(rest.AuthRestriction{
				MinimumRole:    tc.requiredRole,
				RequiredScopes: tc.requiredScopes,
			}))
			engine.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})

			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			engine.ServeHTTP(res, req)

			require.Equal(t, tc.status, res.Code)
			if tc.body != "" {
				assert.Contains(t, res.Body.String(), tc.body)
			}
		})
	}
}
