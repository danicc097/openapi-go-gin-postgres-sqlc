package rest

import (
	"fmt"
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

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Set(string(userCtxKey), &db.Users{Role: db.RoleUser})
	req, _ := http.NewRequestWithContext(c, http.MethodGet, os.Getenv("API_VERSION")+"/admin/ping", nil)

	user := GetUser(c)
	fmt.Printf("userrrr: %#v\n", user)
	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
