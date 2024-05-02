package rest_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaginatedNotificationsRoute(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("user notifications", func(t *testing.T) {
		t.Parallel()

		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			WithAPIKey: true,
		})
		require.NoError(t, err)

		notification := ff.CreatePersonalNotification(context.Background(), servicetestutil.CreateNotificationParams{Receiver: &ufixture.UserID})

		p := &models.GetPaginatedNotificationsParams{Limit: 5, Direction: models.DirectionAsc, Cursor: pointers.New("0")}
		nres, err := srv.client.GetPaginatedNotificationsWithResponse(context.Background(), p, ReqWithAPIKey(ufixture.APIKey.APIKey))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, nres.StatusCode())

		// we are already validating structure via response validator. now we should just focus on
		// testing elements intrinsic to rest layer in handlers, such as status codes, pagination next cursor returned...
		body := nres.JSON200
		assert.EqualValues(t, fmt.Sprint(notification.UserNotificationID), *body.Page.NextCursor)
		// this would actually be a duplicated test
		assert.Len(t, *body.Items, 1)
		assert.True(t, (*body.Items)[0].UserID == ufixture.UserID)
	})
}
