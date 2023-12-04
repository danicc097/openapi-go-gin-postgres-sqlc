package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestGetPaginatedNotificationsRoute(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.cleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(testPool, svc)

	t.Run("all notifications", func(t *testing.T) {
		t.Parallel()

		ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			WithAPIKey: true,
		})
		require.NoError(t, err, "ff.CreateUser: %s")

		_, err = ff.CreatePersonalNotification(context.Background(), servicetestutil.CreateNotificationParams{Receiver: &ufixture.User.UserID})
		require.NoError(t, err)

		p := &models.GetPaginatedNotificationsParams{Limit: 5, Direction: models.GetPaginatedNotificationsParamsDirectionAsc, Cursor: "0"}
		nres, err := srv.client.GetPaginatedNotificationsWithResponse(context.Background(), p, resttestutil.ReqWithAPIKey(ufixture.APIKey.APIKey))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, nres.StatusCode())
		t.Log(string(nres.Body))
		got, err := json.Marshal(nres.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&PaginatedNotificationsResponse{Page: PaginationPage{NextCursor: }})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "")
	})
}
