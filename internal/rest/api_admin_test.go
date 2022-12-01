package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	ff := newTestFixtureFactory(t)

	ufixture, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
		Role:       models.RoleAdmin,
		Scopes:     []models.Scope{models.ScopeTestScope},
		WithAPIKey: true,
	})
	if err != nil {
		t.Fatalf("ff.CreateUser: %s", err)
	}

	srv, err := runTestServer(t, testpool, []gin.HandlerFunc{func(c *gin.Context) {
		c.Next()
	}})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/admin/ping", nil)
	req.Header.Add("x-api-key", ufixture.APIKey.APIKey)

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
