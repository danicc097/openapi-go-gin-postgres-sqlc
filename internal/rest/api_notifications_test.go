package rest

import (
	"context"
	"fmt"
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

		notification, err := ff.CreatePersonalNotification(context.Background(), servicetestutil.CreateNotificationParams{Receiver: &ufixture.User.UserID})
		require.NoError(t, err)

		p := &models.GetPaginatedNotificationsParams{Limit: 5, Direction: models.DirectionAsc, Cursor: "0"}
		nres, err := srv.client.GetPaginatedNotificationsWithResponse(context.Background(), p, resttestutil.ReqWithAPIKey(ufixture.APIKey.APIKey))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, nres.StatusCode())

		got := nres.JSON200
		assert.Equal(t, fmt.Sprint(notification.UserNotificationID), *got.Page.NextCursor)
		assert.Len(t, *got.Items, 1)
		assert.True(t, (*got.Items)[0].UserID == ufixture.User.UserID.UUID)
		assert.True(t, (*got.Items)[0].Read == false)
	})
}
