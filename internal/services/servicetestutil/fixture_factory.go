package servicetestutil

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// FixtureFactory provides fixtures to create randomized elements
// in the data store.
type FixtureFactory struct {
	usvc     *services.User
	d        db.DBTX
	authnsvc *services.Authentication
	authzsvc *services.Authorization
}

// NewFixtureFactory returns a new FixtureFactory.
func NewFixtureFactory(
	usvc *services.User,
	d db.DBTX,
	authnsvc *services.Authentication,
	authzsvc *services.Authorization,
) *FixtureFactory {
	return &FixtureFactory{
		usvc:     usvc,
		d:        d,
		authnsvc: authnsvc,
		authzsvc: authzsvc,
	}
}
