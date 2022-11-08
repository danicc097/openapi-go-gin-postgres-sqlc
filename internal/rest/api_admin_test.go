package rest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	authzsvc, err := services.NewAuthorization(zaptest.NewLogger(t), "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("services.NewAuthorization: %v", err)
	}

	srv, err := runTestServer(t, pool, []gin.HandlerFunc{func(c *gin.Context) {
		r, err := authzsvc.RoleByName(string(models.RoleAdmin))
		if err != nil {
			t.Fatalf("authzsvc.RoleByName: %v", err)
		}
		ctxWithUser(c, &db.User{RoleRank: r.Rank})
	}})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/admin/ping", nil)
	req.Header.Add("x-api-key", "dummy-key")

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
