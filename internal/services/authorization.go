package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// Role represents a predefined role that may be required
// for specific actions regardless of scopes assigned to a user.
// It is also associated with a collection of scopes that get assigned/revoked upon role change.
type Role struct {
	Name        db.UserRole
	Description string
	Rank        uint
}

var (
	RoleGuest = Role{
		Name:        db.UserRoleGuest,
		Description: "Users with limited read-only permissions.",
		Rank:        1,
	}
	RoleUser = Role{
		Name:        db.UserRoleUser,
		Description: "Regular users, with no special permissions.",
		Rank:        2,
	}
	RoleAdvancedUser = Role{
		Name:        db.UserRoleAdvanceduser,
		Description: "Users with additional permissions.",
		Rank:        3,
	}
	RoleManager = Role{
		Name:        db.UserRoleManager,
		Description: "Managers have privileged access for team management.",
		Rank:        4,
	}
	RoleAdmin = Role{
		Name:        db.UserRoleAdmin,
		Description: "Admins can manage all settings on a per-project basis.",
		Rank:        5,
	}
	RoleOwner = Role{
		Name:        db.UserRoleSuperadmin,
		Description: "Superadmins have unrestricted access to any project.",
		Rank:        6,
	}
)

// Authorization represents a service for authorization.
type Authorization struct {
	logger *zap.Logger
}

func NewAuthorization(logger *zap.Logger) *Authorization {
	return &Authorization{
		logger: logger,
	}
}

// TODO ABAC:
// for scope structure references (not roles logic obv) see:
// https://incidentio.notion.site/Proposal-Product-RBAC-265201563d884ec5aeecbb246c02ddc6
// last resort: casbin. too much scope, poor docs, maintenance
// for frontend https://casbin.org/docs/en/frontend
// load policy from db: https://github.com/casbin/casbin-pg-adapter

// RolePermissions returns access levels per role.
func (a Authorization) RolePermissions() map[db.UserRole][]db.UserRole {
	return roles
}

func (a Authorization) IsAuthorized(role, requiredRole db.UserRole) error {
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
