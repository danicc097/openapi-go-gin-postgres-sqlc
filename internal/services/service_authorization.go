package services

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
	"golang.org/x/exp/slices"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
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
//
// (1) append scopes:
//
//	update users
//	set    scopes = (select array_agg(distinct e) from unnest(scopes || '{"newscope-1","newscope-2"}') e)
//	where  not scopes @> '{"newscope-1","newscope-2"}' and role_rank >= @minimum_role_rank
//
// (2) add a new role:
//
//	update ... set rank = rank +1 where rank >= @new_role_rank
var (
	userScopes = models.Scopes{
		models.ScopeUsersRead,
	}
	managerScopes = append(
		userScopes,
		models.ScopeWorkItemReview,
		models.ScopeWorkItemTagCreate,
		models.ScopeWorkItemTagDelete,
		models.ScopeWorkItemTagEdit)
	adminScopes = append(managerScopes, models.ScopeUsersWrite)

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
func NewAuthorization(logger *zap.SugaredLogger) *Authorization {
	roles := make(roles)
	scopes := make(scopes)

	scopeBlob, err := os.ReadFile(internal.Config.ScopePolicyPath)
	if err != nil {
		panic(fmt.Sprintf("scope policy: %s", err))
	}
	roleBlob, err := os.ReadFile(internal.Config.RolePolicyPath)
	if err != nil {
		panic(fmt.Sprintf("role policy: %s", err))
	}
	if err := json.Unmarshal(scopeBlob, &scopes); err != nil {
		panic(fmt.Sprintf("scope policy loading: %s", err))
	}
	if err := json.Unmarshal(roleBlob, &roles); err != nil {
		panic(fmt.Sprintf("role policy loading: %s", err))
	}

	return &Authorization{
		logger:         logger,
		roles:          roles,
		scopes:         scopes,
		existingRoles:  models.AllRoleValues(),
		existingScopes: models.AllScopeValues(),
	}
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
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "%s", fmt.Sprintf("access restricted: missing scope %s", rs))
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
