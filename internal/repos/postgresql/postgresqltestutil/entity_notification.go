package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/stretchr/testify/require"
)

func NewRandomEntityNotification(t *testing.T, d db.DBTX) (*db.EntityNotification, error) {
	t.Helper()

	entityNotificationRepo := postgresql.NewEntityNotification()

	ucp := RandomEntityNotificationCreateParams(t)

	entityNotification, err := entityNotificationRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return entityNotification, nil
}

func RandomEntityNotificationCreateParams(t *testing.T) *db.EntityNotificationCreateParams {
	t.Helper()

	return &db.EntityNotificationCreateParams{
		// TODO: fill in with testutil randomizer helpers
	}
}

