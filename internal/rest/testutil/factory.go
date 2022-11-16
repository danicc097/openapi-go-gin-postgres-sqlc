package testutil

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/jackc/pgx/v4/pgxpool"
)

// FixtureFactory provides fixtures to create randomized elements
// in the data store.
type FixtureFactory struct {
	usvc     *services.User
	pool     *pgxpool.Pool
	authnsvc *services.Authentication
	authzsvc *services.Authorization
}

// NewFixtureFactory returns a new FixtureFactory.
func NewFixtureFactory(
	usvc *services.User,
	pool *pgxpool.Pool,
	authnsvc *services.Authentication,
	authzsvc *services.Authorization,
) *FixtureFactory {
	return &FixtureFactory{
		usvc:     usvc,
		pool:     pool,
		authnsvc: authnsvc,
		authzsvc: authzsvc,
	}
}
