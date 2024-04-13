package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, context.Background(), testPool)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.setupCleanup(t)

	res, err := srv.client.PingWithResponse(context.Background())
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))
	assert.Equal(t, "pong", string(res.Body))
}
