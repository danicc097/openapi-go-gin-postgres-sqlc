package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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

		usvc := services.NewUser(postgresql.NewUser(), logger)
		authzsvc, err := services.NewAuthorization(logger, "testdata/scopes.json", "testdata/roles.json")
		if err != nil {
			t.Fatalf("services.NewAuthorization: %v", err)
		}
		authnsvc := services.NewAuthentication(logger, usvc)

		authMw := newAuthMiddleware(logger, pool, authnsvc, authzsvc, usvc)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		engine.Use(func(c *gin.Context) {
			r, err := authzsvc.RoleByName(string(tc.role))
			if err != nil {
				t.Fatalf("authzsvc.RoleByName: %v", err)
			}
			ctxWithUser(c, &db.User{RoleRank: r.Rank})
		})

		engine.Use(authMw.EnsureAuthorized(AuthRestriction{MinimumRole: tc.requiredRole}))
		engine.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		engine.ServeHTTP(resp, req)

		assert.Equal(t, tc.status, resp.Code)
		assert.Contains(t, resp.Body.String(), tc.body)
	}
}
