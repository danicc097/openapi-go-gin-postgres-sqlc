package e2e_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := gen.NewRouter()

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
