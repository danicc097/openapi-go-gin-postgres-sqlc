package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func NewRandomEntityNotification(t *testing.T, d db.DBTX, projectID db.ProjectID) (*db.EntityNotification, error) {
	t.Helper()

	entityNotificationRepo := postgresql.NewEntityNotification()

	ucp := RandomEntityNotificationCreateParams(t, projectID)

	entityNotification, err := entityNotificationRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return entityNotification, nil
}

func RandomEntityNotificationCreateParams(t *testing.T, projectID db.ProjectID) *db.EntityNotificationCreateParams {
	t.Helper()

	return &db.EntityNotificationCreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
		ProjectID: projectID,
		ID:        testutil.RandomString(3432),
		Message:   testutil.RandomString(3432),
		Topic:     models.TopicsGlobalAlerts,
	}
}
