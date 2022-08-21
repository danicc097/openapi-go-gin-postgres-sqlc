package services

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql/gen"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// Authorization represents a service for authorization.
type Authorization struct {
	Logger *zap.Logger
}

func NewAuthorization(logger *zap.Logger) *Authorization {
	return &Authorization{
		Logger: logger,
	}
}

// AuthorizationService
// authentication/authorization based on specific requirements
// and out of scope of this app.
//  -- delegated to auth server
//  -- oauth2 where resource and auth servers are the same
//  -- sessions and cookies
type AuthorizationService interface {
	IsAuthorized(role, requiredRole db.Role) bool
	RolePermissions() map[db.Role][]db.Role
}

// we dont want accidental edits to permissions
// TODO memoize
func (a *Authorization) RolePermissions() map[db.Role][]db.Role {
	rolePermissions := map[db.Role][]db.Role{
		db.RoleUser:    {db.RoleUser},
		db.RoleManager: {db.RoleUser, db.RoleManager},
		db.RoleAdmin:   {db.RoleUser, db.RoleManager, db.RoleAdmin},
	}

	return rolePermissions
}

func (a *Authorization) IsAuthorized(role, requiredRole db.Role) bool {
	roles := a.RolePermissions()[role]

	return slices.Contains(roles, requiredRole)
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
