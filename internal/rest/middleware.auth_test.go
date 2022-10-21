package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAuthorizationMiddleware(t *testing.T) {
	testCases := []struct {
		name         string
		role         db.Role
		requiredRole db.Role
		status       int
		body         string
	}{
		{
			name:         "unauthorized_user",
			role:         db.RoleUser,
			requiredRole: db.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "access restricted",
		},
		{
			name:         "unauthorized_manager",
			role:         db.RoleManager,
			requiredRole: db.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "access restricted",
		},
		{
			name:         "authorized",
			role:         db.RoleAdmin,
			requiredRole: db.RoleAdmin,
			status:       http.StatusForbidden,
			body:         "ok",
		},
	}

	for _, tc := range testCases {
		resp := httptest.NewRecorder()
		logger, _ := zap.NewDevelopment()
		_, engine := gin.CreateTestContext(resp)

		authMw := newAuthMiddleware(logger, pool)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		engine.Use(func(c *gin.Context) {
			ctxWithUser(c, &db.Users{Role: tc.role})
		})
		engine.Use(authMw.EnsureAuthorized(tc.requiredRole))
		engine.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		engine.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusForbidden, tc.status)
		assert.Contains(t, resp.Body.String(), tc.body)
	}
}
