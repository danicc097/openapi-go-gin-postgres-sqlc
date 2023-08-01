package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// Role represents a predefined role that may be required
// for specific actions regardless of scopes assigned to a user.
type Role struct {
	Description string      `json:"description"`
	Rank        int         `json:"rank"`
	Name        models.Role `json:"name"`
}

type Scope struct {
	Description string `json:"description"`
}

type (
	roles  = map[models.Role]Role
	scopes = map[models.Scope]Scope
)

// nolint:gochecknoglobals
// NOTE: ensure any changes are followed by an appropriate migration.
var (
	userScopes = models.Scopes{
		models.ScopeUsersRead,
	}
	managerScopes = append(userScopes, models.ScopeWorkItemReview)
	adminScopes   = append(managerScopes, models.ScopeUsersWrite)

	// scopesByRole represents user scopes by role.
	scopesByRole = map[models.Role]models.Scopes{
		models.RoleGuest:        {},
		models.RoleUser:         userScopes,
		models.RoleAdvancedUser: userScopes,
		models.RoleManager:      managerScopes,
		models.RoleAdmin:        adminScopes,
		models.RoleSuperAdmin:   adminScopes,
	}
)

// Authorization represents a service for authorization.
type Authorization struct {
	logger         *zap.SugaredLogger
	roles          roles
	scopes         scopes
	existingRoles  []models.Role
	existingScopes models.Scopes
}

// NewAuthorization returns a new Authorization service.
// Existing roles and scopes will be loaded from the given policy JSON file paths.
func NewAuthorization(logger *zap.SugaredLogger, scopePolicy string, rolePolicy string) (*Authorization, error) {
	roles := make(roles)
	scopes := make(scopes)

	scopeBlob, err := os.ReadFile(scopePolicy)
	if err != nil {
		return nil, fmt.Errorf("scope policy: %w", err)
	}
	roleBlob, err := os.ReadFile(rolePolicy)
	if err != nil {
		return nil, fmt.Errorf("role policy: %w", err)
	}
	if err := json.Unmarshal(scopeBlob, &scopes); err != nil {
		return nil, fmt.Errorf("scope policy loading: %w", err)
	}
	if err := json.Unmarshal(roleBlob, &roles); err != nil {
		return nil, fmt.Errorf("role policy loading: %w", err)
	}

	return &Authorization{
		logger:         logger,
		roles:          roles,
		scopes:         scopes,
		existingRoles:  models.AllRoleValues(),
		existingScopes: models.AllScopeValues(),
	}, nil
}

func (a *Authorization) RoleByName(role models.Role) Role {
	return a.roles[role]
}

func (a *Authorization) RoleByRank(rank int) (Role, bool) {
	for _, r := range a.roles {
		if r.Rank == rank {
			return r, true
		}
	}

	return Role{}, false
}

func (a *Authorization) ScopeByName(scope string) (Scope, error) {
	s, ok := a.scopes[models.Scope(scope)]
	if !ok {
		return Scope{}, internal.NewErrorf(models.ErrorCodeInvalidScope, "unknown scope %s", scope)
	}

	return s, nil
}

func (a *Authorization) HasRequiredRole(role Role, requiredRole models.Role) error {
	if role.Rank < a.roles[requiredRole].Rank {
		return internal.NewErrorf(models.ErrorCodeUnauthorized, "access restricted: unauthorized role")
	}

	return nil
}

func (a *Authorization) HasRequiredScopes(scopes models.Scopes, requiredScopes models.Scopes) error {
	for _, rs := range requiredScopes {
		if !slices.Contains(scopes, rs) {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, fmt.Sprintf("access restricted: missing scope %s", rs))
		}
	}

	return nil
}

// DefaultScopes returns the default scopes for a role.
// Scopes are assigned/revoked upon role change (reset completely).
func (a *Authorization) DefaultScopes(role models.Role) models.Scopes {
	scopes := models.Scopes{}

	if defaultScopes, ok := scopesByRole[role]; ok {
		scopes = defaultScopes
	}

	return scopes
}

/*
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
