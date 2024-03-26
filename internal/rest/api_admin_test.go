package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleAdmin,
		WithAPIKey: true,
	})

	srv, err := runTestServer(t, testPool)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.setupCleanup(t)

	t.Run("authorized", func(t *testing.T) {
		t.Parallel()

		res, err := srv.client.AdminPingWithResponse(context.Background(), ReqWithAPIKey(ufixture.APIKey.APIKey))
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))
		assert.Equal(t, "pong", string(res.Body))
	})
	t.Run("missing_auth_header", func(t *testing.T) {
		t.Parallel()

		res, err := srv.client.AdminPingWithResponse(context.Background())
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode(), string(res.Body))
		assert.EqualValues(t, models.ErrorCodeUnauthenticated, res.JSON4XX.Type)
	})
}
