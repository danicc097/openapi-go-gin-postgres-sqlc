package servicetestutil

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// FixtureFactory provides fixtures to create randomized elements
// in the data store.
type FixtureFactory struct {
	d   db.DBTX
	svc *services.Services
}

// NewFixtureFactory returns a new FixtureFactory.
func NewFixtureFactory(
	d db.DBTX,
	svc *services.Services,
) *FixtureFactory {
	return &FixtureFactory{
		d:   d,
		svc: svc,
	}
}
