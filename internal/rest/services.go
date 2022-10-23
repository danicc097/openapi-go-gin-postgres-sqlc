// grouped for generation caching purposes
package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type UserService interface {
	Upsert(ctx context.Context, user *db.User) error
	UserByEmail(ctx context.Context, email string) (*db.User, error)
	Register(ctx context.Context, user *db.User) error
	// +anything related to users
}

// TODO custom RBAC.
// casbin: strange api
type AuthorizationService interface {
	IsAuthorized(role, requiredRole db.Role) error
}

// TODO oidc server and client from zitadel
type AuthenticationService interface{}
