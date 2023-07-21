package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.cleanup(t)

	req, _ := http.NewRequest(http.MethodGet, resttestutil.MustConstructInternalPath("/ping"), nil)
	resp := httptest.NewRecorder()
	srv.server.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
