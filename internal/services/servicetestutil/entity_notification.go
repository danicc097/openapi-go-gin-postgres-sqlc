package servicetestutil

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
)

type CreateEntityNotificationParams struct {
	ProjectID  db.ProjectID
	// DeletedAt allows returning a soft deleted entity notification when a deleted_at column exists.
	// Note that the service Delete call should make use of the SoftDelete method.
	DeletedAt  *time.Time
}

type CreateEntityNotificationFixture struct {
	EntityNotification   *db.EntityNotification
}

// CreateEntityNotification creates a new random entity notification with the given configuration.
func (ff *FixtureFactory) CreateEntityNotification(ctx context.Context, params CreateEntityNotificationParams) (*CreateEntityNotificationFixture, error) {
	randomRepoCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(ff.t) // , params.ProjectID
	// don't use repos for tests
	entityNotification, err := ff.svc.EntityNotification.Create(ctx, ff.db, randomRepoCreateParams)
	if err != nil {
		return nil, fmt.Errorf("svc.EntityNotification.Create: %w", err)
	}

	if params.DeletedAt != nil {
		entityNotification, err = ff.svc.EntityNotification.Delete(ctx, ff.db, entityNotification.EntityNotificationID)
		if err != nil {
			return nil, fmt.Errorf("svc.EntityNotification.Delete: %w", err)
		}
	}

	return &CreateEntityNotificationFixture{
		EntityNotification:   entityNotification,
	}, nil
}


