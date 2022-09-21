package rest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(pool, []gin.HandlerFunc{func(ctx *gin.Context) {
		ctx.Set(string(userCtxKey), &db.Users{Role: db.RoleAdmin})
	}})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/admin/ping", nil)

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
