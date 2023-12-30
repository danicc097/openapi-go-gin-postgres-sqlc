package servicetestutil

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
)

type CreateEntityNotificationParams struct {
	ProjectID db.ProjectID
}

type CreateEntityNotificationFixture struct {
	EntityNotification *db.EntityNotification
}

// CreateEntityNotification creates a new random entity notification with the given configuration.
func (ff *FixtureFactory) CreateEntityNotification(ctx context.Context, params CreateEntityNotificationParams) (*CreateEntityNotificationFixture, error) {
	randomRepoCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(ff.t, params.ProjectID)
	// don't use repos for tests
	entityNotification, err := ff.svc.EntityNotification.Create(ctx, ff.db, randomRepoCreateParams)
	if err != nil {
		return nil, fmt.Errorf("svc.EntityNotification.Create: %w", err)
	}

	return &CreateEntityNotificationFixture{
		EntityNotification: entityNotification,
	}, nil
}
