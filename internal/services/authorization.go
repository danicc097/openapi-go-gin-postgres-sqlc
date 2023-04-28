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

// NOTE: openapi Role enum is merely a string enum array for views. Ranks are for internal use.

// Role represents a predefined role that may be required
// for specific actions regardless of scopes assigned to a user.
// TODO It is also associated with a collection of scopes that get assigned/revoked upon role change.
type Role struct {
	Description string      `json:"description"`
	Rank        int16       `json:"rank"` // to avoid casting. postgres smallint with check > 0
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
	userScopes = []models.Scope{
		models.ScopeUsersRead,
		models.ScopeTestScope,
	}
	managerScopes = append(userScopes, models.ScopeWorkItemReview)
	adminScopes   = append(managerScopes, models.ScopeWorkItemReview)
)

// nolint:gochecknoglobals
var scopesByRole = map[models.Role][]models.Scope{
	models.RoleUser:    userScopes,
	models.RoleManager: managerScopes,
	models.RoleAdmin:   adminScopes,
}

// Authorization represents a service for authorization.
type Authorization struct {
	logger         *zap.Logger
	roles          roles
	scopes         scopes
	existingRoles  []models.Role
	existingScopes []models.Scope
}

// NewAuthorization returns a new Authorization service.
// Existing roles and scopes will be loaded from the given policy JSON file paths.
func NewAuthorization(logger *zap.Logger, scopePolicy string, rolePolicy string) (*Authorization, error) {
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
		return nil, fmt.Errorf("scope policy: %w", err)
	}
	if err := json.Unmarshal(roleBlob, &roles); err != nil {
		return nil, fmt.Errorf("role policy: %w", err)
	}

	return &Authorization{
		logger:         logger,
		roles:          roles,
		scopes:         scopes,
		existingRoles:  models.AllRoleValues(),
		existingScopes: models.AllScopeValues(),
	}, nil
}

func (a *Authorization) RoleByName(role string) (Role, error) {
	rl, ok := a.roles[models.Role(role)]
	if !ok {
		return Role{}, internal.NewErrorf(internal.ErrorCodeUnauthorized, "unknown role %s", role)
	}

	return rl, nil
}

func (a *Authorization) RoleByRank(rank int16) (Role, bool) {
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
		return Scope{}, internal.NewErrorf(internal.ErrorCodeInvalidScope, "unknown scope %s", scope)
	}

	return s, nil
}

func (a *Authorization) ValidateScopes(scopes ...string) error {
	for _, s := range scopes {
		_, ok := a.scopes[models.Scope(s)]
		if !ok {
			return internal.NewErrorf(internal.ErrorCodeInvalidScope, "unknown scope %s", s)
		}
	}

	return nil
}

func (a *Authorization) ValidateRole(role string) error {
	_, ok := a.roles[models.Role(role)]
	if !ok {
		return internal.NewErrorf(internal.ErrorCodeInvalidRole, "unknown role %s", role)
	}

	return nil
}

func (a *Authorization) HasRequiredRole(role Role, requiredRole models.Role) error {
	if err := a.ValidateRole(string(role.Name)); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnauthorized, "role is not valid")
	}
	rl, ok := a.roles[requiredRole]
	if !ok {
		return internal.NewErrorf(internal.ErrorCodeUnauthorized, "unknown role %s", requiredRole)
	}
	if role.Rank < rl.Rank {
		return internal.NewErrorf(internal.ErrorCodeUnauthorized, "access restricted: unauthorized role")
	}

	return nil
}

func (a *Authorization) HasRequiredScopes(scopes []string, requiredScopes []models.Scope) error {
	if err := a.ValidateScopes(scopes...); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnauthorized, "scopes are not valid")
	}
	for _, rs := range requiredScopes {
		if !slices.Contains(scopes, string(rs)) {
			return internal.NewErrorf(internal.ErrorCodeUnauthorized, fmt.Sprintf("access restricted: missing scope %s", rs))
		}
	}

	return nil
}

// DefaultScopes returns the default scopes for a role.
func (a *Authorization) DefaultScopes(role Role) (scopes []models.Scope) {
	if role.Rank > a.roles[models.RoleUser].Rank {
		scopes = append(scopes, models.ScopeUsersRead)
	}
	if role.Rank > a.roles[models.RoleAdvancedUser].Rank {
	}

	// TODO fill in

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
