package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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
			body:           "Unauthorized",
		},
		{
			name:        "unauthorized_if_no_user",
			withoutUser: true,
			status:      http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resp := httptest.NewRecorder()
			logger, _ := zap.NewDevelopment()
			_, engine := gin.CreateTestContext(resp)

			authzsvc, err := services.NewAuthorization(logger.Sugar())
			require.NoError(t, err, "newTestAuthService")

			repos := services.CreateTestRepos()
			usvc := services.NewUser(logger.Sugar(), repos)
			authnsvc := services.NewAuthentication(logger.Sugar(), repos, testPool)

			authMw := newAuthMiddleware(logger.Sugar(), testPool, authnsvc, authzsvc, usvc)

			ff := servicetestutil.NewFixtureFactory(usvc, testPool, authnsvc, authzsvc)
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				Scopes:     tc.scopes,
				WithAPIKey: true,
			})
			if err != nil {
				t.Fatalf("ff.CreateUser: %s", err)
			}

			if !tc.withoutUser {
				engine.Use(func(c *gin.Context) {
					ctxWithUser(c, ufixture.User)
				})
			}

			engine.Use(authMw.EnsureAuthorized(AuthRestriction{
				MinimumRole:    tc.requiredRole,
				RequiredScopes: tc.requiredScopes,
			}))
			engine.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})

			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			engine.ServeHTTP(resp, req)

			assert.Equal(t, tc.status, resp.Code)
			if tc.body != "" {
				assert.Contains(t, resp.Body.String(), tc.body)
			}
		})
	}
}
