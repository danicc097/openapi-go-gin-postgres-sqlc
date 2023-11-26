package rest

import (
	"context"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	svc := services.New(zap.S(), services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(testPool, svc)

	ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleAdmin,
		WithAPIKey: true,
	})
	if err != nil {
		t.Fatalf("ff.CreateUser: %s", err)
	}

	srv, err := runTestServer(t, testPool)
	require.NoError(t, err, "Couldn't run test server: %s\n")
	srv.cleanup(t)

	t.Run("authorized", func(t *testing.T) {
		t.Parallel()

		res, err := srv.client.AdminPingWithResponse(context.Background(), resttestutil.ReqWithAPIKey(ufixture.APIKey.APIKey))
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, "pong", string(res.Body))
	})
	t.Run("missing_auth_header", func(t *testing.T) {
		t.Parallel()

		res, err := srv.client.AdminPingWithResponse(context.Background())
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
		assert.Equal(t, models.ErrorCodeRequestValidation, res.JSON4XX.Type)
	})
}
