package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestHandlers_DeleteEntityNotification(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   "valid entity notification deletion",
			status: http.StatusNoContent,
			scopes: []models.Scope{models.ScopeEntityNotificationDelete},
		},
		{
			name:   "unauthorized entity notification call",
			status: http.StatusForbidden,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})
			require.NoError(t, err, "ff.CreateUser: %s")

			entityNotification, err := ff.CreateEntityNotification(context.Background(), servicetestutil.CreateEntityNotificationParams{})
			require.NoError(t, err, "ff.CreateEntityNotification: %s")

			id := entityNotification.EntityNotification.EntityNotificationID
			res, err := srv.client.DeleteEntityNotificationWithResponse(context.Background(), int(id), ReqWithAPIKey(ufixture.APIKey.APIKey))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode(), string(res.Body))
		})
	}
}

func TestHandlers_CreateEntityNotification(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{models.ScopeEntityNotificationCreate}

		ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		require.NoError(t, err, "ff.CreateUser: %s")

		randomEntityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t)
		body := rest.CreateEntityNotificationRequest{
			EntityNotificationCreateParams: *randomEntityNotificationCreateParams,
		}

		res, err := srv.client.CreateEntityNotificationWithResponse(context.Background(), body, ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode(), string(res.Body))
		assert.EqualValues(t, randomEntityNotificationCreateParams.ID, res.JSON201.ID)
		assert.EqualValues(t, randomEntityNotificationCreateParams.Message, res.JSON201.Message)
		assert.EqualValues(t, randomEntityNotificationCreateParams.Topic, res.JSON201.Topic)
	})
}

func TestHandlers_GetEntityNotification(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{}

		ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		require.NoError(t, err, "ff.CreateUser: %s")

		entityNotification, err := ff.CreateEntityNotification(context.Background(), servicetestutil.CreateEntityNotificationParams{})
		require.NoError(t, err, "ff.CreateEntityNotification: %s")

		id := entityNotification.EntityNotification.EntityNotificationID
		res, err := srv.client.GetEntityNotificationWithResponse(context.Background(), int(id), ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))

		got, err := json.Marshal(res.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&rest.EntityNotification{EntityNotification: *entityNotification.EntityNotification})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "") // ignore private JSON fields
	})
}

func TestHandlers_UpdateEntityNotification(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name                    string
		status                  int
		body                    rest.UpdateEntityNotificationRequest
		validationErrorContains []string
	}{
		{
			name:   "valid entity notification update",
			status: http.StatusOK,
			body: func() rest.UpdateEntityNotificationRequest {
				randomEntityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t)

				return rest.UpdateEntityNotificationRequest{
					EntityNotificationUpdateParams: db.EntityNotificationUpdateParams{
						ID:      pointers.New(randomEntityNotificationCreateParams.ID),
						Message: pointers.New(randomEntityNotificationCreateParams.Message),
						Topic:   pointers.New(randomEntityNotificationCreateParams.Topic),
					},
				}
			}(),
		},
		// NOTE: we do need to test spec validation
		// {
		// 	name:                    "invalid entity notification update param",
		// 	status:                  http.StatusBadRequest,
		// 	body:                    rest.UpdateEntityNotificationRequest{},
		// 	validationErrorContains: []string{"[\" field <JSON >\"]", "<error>"},
		// },
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			normalUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
				Scopes:     []models.Scope{models.ScopeEntityNotificationEdit},
			})
			require.NoError(t, err, "ff.CreateUser: %s")

			entityNotification, err := ff.CreateEntityNotification(context.Background(), servicetestutil.CreateEntityNotificationParams{})
			require.NoError(t, err, "ff.CreateEntityNotification: %s")

			id := entityNotification.EntityNotification.EntityNotificationID
			updateRes, err := srv.client.UpdateEntityNotificationWithResponse(context.Background(), int(id), tc.body, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			require.EqualValues(t, tc.status, updateRes.StatusCode(), string(updateRes.Body))

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(updateRes.Body), ve)
				}

				return
			}

			assert.EqualValues(t, id, updateRes.JSON200.EntityNotificationID)

			res, err := srv.client.GetEntityNotificationWithResponse(context.Background(), int(id), ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.EqualValues(t, *tc.body.ID, res.JSON200.ID)
			assert.EqualValues(t, *tc.body.Message, res.JSON200.Message)
			assert.EqualValues(t, *tc.body.Topic, res.JSON200.Topic)
		})
	}
}
