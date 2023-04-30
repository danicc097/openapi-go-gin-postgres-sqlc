package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAuthorizationMiddleware_Roles(t *testing.T) {
	t.Parallel()

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
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resp := httptest.NewRecorder()
			logger, _ := zap.NewDevelopment()
			_, engine := gin.CreateTestContext(resp)

			authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
			if err != nil {
				t.Fatalf("services.NewAuthorization: %v", err)
			}
			usvc := services.NewUser(logger, postgresql.NewUser(), postgresql.NewNotification(), authzsvc)
			authnsvc := services.NewAuthentication(logger, usvc, testPool)

			authMw := newAuthMiddleware(logger, testPool, authnsvc, authzsvc, usvc)

			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			ff := servicetestutil.NewFixtureFactory(usvc, testPool, authnsvc, authzsvc)
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
			})
			if err != nil {
				t.Fatalf("ff.CreateUser: %s", err)
			}

			engine.Use(func(c *gin.Context) {
				ctxWithUser(c, ufixture.User)
			})

			engine.Use(authMw.EnsureAuthorized(AuthRestriction{MinimumRole: tc.requiredRole}))
			engine.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})
			engine.ServeHTTP(resp, req)

			assert.Equal(t, tc.status, resp.Code)
			assert.Contains(t, resp.Body.String(), tc.body)
		})
	}
}
