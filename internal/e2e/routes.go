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
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}
