package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
)

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
	IsAuthorized(role, requiredRole db.Role) error
}

type AuthenticationService interface {
	// TODO Authentication delegated to auth server.
	// will use inmemory tokens for predefined users for simplicity.
}
