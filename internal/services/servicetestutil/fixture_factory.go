package servicetestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// FixtureFactory provides fixtures to create randomized elements
// in the data store.
type FixtureFactory struct {
	t   *testing.T
	d   models.DBTX
	svc *services.Services
}

// NewFixtureFactory returns a new FixtureFactory.
func NewFixtureFactory(
	t *testing.T,
	d models.DBTX,
	svc *services.Services,
) *FixtureFactory {
	return &FixtureFactory{
		t:   t,
		d:   d,
		svc: svc,
	}
}
