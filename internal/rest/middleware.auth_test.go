package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAuthorizationMiddleware(t *testing.T) {
	testCases := []struct {
		name         string
		role         models.Role
		requiredRole models.Role
		status       int
		body         string
	}{
		{
			name:         "unauthorized_user",
			role:         models.RoleUser,
			requiredRole: models.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "Unauthorized.",
		},
		{
			name:         "unauthorized_manager",
			role:         models.RoleManager,
			requiredRole: models.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "Unauthorized.",
		},
		{
			name:         "authorized",
			role:         models.RoleAdmin,
			requiredRole: models.RoleAdmin,
			status:       http.StatusOK,
			body:         "ok",
		},
	}

	for _, tc := range testCases {
		resp := httptest.NewRecorder()
		logger, _ := zap.NewDevelopment()
		_, engine := gin.CreateTestContext(resp)

		usvc := services.NewUser(repos.NewUserWrapped(postgresql.NewUser(), otelName, repos.UserWrappedConfig{}, nil), logger)
		authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
		if err != nil {
			t.Fatalf("services.NewAuthorization: %v", err)
		}
		authnsvc := services.NewAuthentication(logger, usvc, pool)

		authMw := newAuthMiddleware(logger, pool, authnsvc, authzsvc, usvc)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		ff := testutil.NewFixtureFactory(usvc, pool, authnsvc, authzsvc)
		user, err := ff.CreateUser(context.Background(), testutil.CreateUserParams{
			Role:       tc.role,
			WithAPIKey: true,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}

		engine.Use(func(c *gin.Context) {
			ctxWithUser(c, user)
		})

		engine.Use(authMw.EnsureAuthorized(AuthRestriction{MinimumRole: tc.requiredRole}))
		engine.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
		engine.ServeHTTP(resp, req)

		assert.Equal(t, tc.status, resp.Code)
		assert.Contains(t, resp.Body.String(), tc.body)
	}
}
