// grouped for generation caching purposes
package rest

import (
	"context"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
)

type UserService interface {
	Upsert(ctx context.Context, user *crud.User) error
	UserByEmail(ctx context.Context, email string) (*crud.User, error)
	Create(ctx context.Context, user *crud.User) error
	// +anything related to users
}

// TODO custom RBAC.
// casbin: strange api
type AuthorizationService interface {
	IsAuthorized(role, requiredRole db.Role) error
}

// TODO oidc server and client from zitadel
type AuthenticationService interface{}
