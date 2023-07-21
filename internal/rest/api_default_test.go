package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()

	srv, _, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	require.NoError(t, err, "Couldn't run test server: %s\n")
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, resttestutil.MustConstructInternalPath("/ping"), nil)
	resp := httptest.NewRecorder()
	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
