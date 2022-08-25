package handlers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
)

// one per package
//go:generate counterfeiter -generate

//counterfeiter:generate -o handlerstesting/authorization.gen.go . AuthorizationService
//counterfeiter:generate -o handlerstesting/authentication.gen.go . AuthenticationService

type UserService interface {
	Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error)
}

/* authentication/authorization based on specific requirements
and out of scope of this app.
 -- delegated to auth server
 -- oauth2 where resource and auth servers are the same
 -- sessions and cookies
*/

type AuthorizationService interface {
	IsAuthorized(role, requiredRole db.Role) bool
}

type AuthenticationService interface {
	// TODO Authentication delegated to auth server.
	// will use inmemory tokens for predefined users for simplicity.
}
