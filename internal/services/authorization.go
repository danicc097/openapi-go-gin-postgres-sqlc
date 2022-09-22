package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// Authorization represents a service for authorization.
type Authorization struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}

func NewAuthorization(logger *zap.Logger) *Authorization {
	return &Authorization{
		Logger: logger,
	}
}

// TODO RBAC: https://incidentio.notion.site/Proposal-Product-RBAC-265201563d884ec5aeecbb246c02ddc6
// but openapi friendly
// RolePermissions returns access levels per role.
func (a Authorization) RolePermissions() map[db.Role][]db.Role {
	return map[db.Role][]db.Role{
		db.RoleUser:    {db.RoleUser},
		db.RoleManager: {db.RoleUser, db.RoleManager},
		db.RoleAdmin:   {db.RoleUser, db.RoleManager, db.RoleAdmin},
	}
}

func (a Authorization) IsAuthorized(role, requiredRole db.Role) error {
	roles := a.RolePermissions()[role]

	if !slices.Contains(roles, requiredRole) {
		return internal.NewErrorf(internal.ErrorCodeUnauthorized, "access restricted")
	}

	return nil
}

/* TODO this is part of the authorization server.
For this app (resource server), all the auth server will do is ensure the user is registered.
Email, username, password changes, password reset requests, user verification is all done externally in
the auth server frontend.
Roles and other user data are up to us.
We should just have an inmemory mock of the auth server with predefined users and tokens
every time we start the app
and thats the end of it. we dont need /auth or /token routes in the resource server,
 just /user since we need
persistent storage for application specific data.


In the future, clone the project and implement the auth server openapi spec, etc.
for jwt, refresh token, redis...
https://developer.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
we would need a redis repo for auth
see https://github.com/joeferner/redis-commander for web interface
refresh_token brings increased security in case the auth server is better secured than
the resource server.
It also improves scalability and performance.
We can have access_token's stored in some fast, temporary storage (Redis).
and they will get deleted by Redis automatically based on expiration field (redis built-in functionality)
If refresh_token is valid, quickly grab the existing access_token or create a new one (scopes, etc. might change as well) and save it
*/
