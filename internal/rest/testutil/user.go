package testutil

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/jackc/pgx/v4/pgxpool"
)

type createUserParams struct {
	usvc *services.User
	pool *pgxpool.Pool

	role          models.Role
	scopes        []models.Scope
	authenticated bool // if true, an access token is created and returned
}

func createUser(params createUserParams) {
	// TODO any value that has a  unique constraint in db must be generated
	// via randomXXX().
	// the only parameters createUser accepts are high level, at the `rest` layer only.
	// functions in this package only make use of SERVICES.
	// it also accepts
	// e.g. roles -> roles from deepmap gen server, not from services or anywhere else.
	// when everything is create we return the user as well as any external data associated to it
	// after creation

	//
}
