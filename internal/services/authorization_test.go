package services_test

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAuthorization_Roles(t *testing.T) {
	t.Parallel()

	svc, err := services.NewAuthorization(zaptest.NewLogger(t).Sugar(), "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	userRole := svc.RoleByName(models.RoleUser)
	managerRole := svc.RoleByName(models.RoleManager)

	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleManager), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(managerRole, models.RoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(services.Role{}, models.RoleAdmin), "access restricted")

	assert.NoError(t, svc.HasRequiredRole(services.Role{Rank: managerRole.Rank, Name: models.RoleManager}, models.RoleManager))
}

func TestAuthorization_Scopes(t *testing.T) {
	t.Parallel()

	svc, err := services.NewAuthorization(zaptest.NewLogger(t).Sugar(), "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	req := models.Scopes{models.ScopeTeamSettingsWrite}
	assert.ErrorContains(t, svc.HasRequiredScopes(models.Scopes{}, req), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredScopes(models.Scopes{models.ScopeUsersRead}, req), "access restricted")
	assert.NoError(t, svc.HasRequiredScopes(models.Scopes{models.ScopeTeamSettingsWrite}, req))

	req = models.Scopes{models.ScopeTeamSettingsWrite, models.ScopeUsersRead}
	assert.ErrorContains(t, svc.HasRequiredScopes(models.Scopes{models.ScopeUsersRead}, req), "access restricted")
}
