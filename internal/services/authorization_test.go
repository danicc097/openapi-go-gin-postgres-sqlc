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

	svc, err := services.NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}
	userRole, err := svc.RoleByName(string(models.RoleUser))
	if err != nil {
		t.Fatalf("role does not exist: %v", err)
	}
	managerRole, err := svc.RoleByName(string(models.RoleManager))
	if err != nil {
		t.Fatalf("role does not exist: %v", err)
	}
	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleManager), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(managerRole, models.RoleAdmin), "access restricted")

	assert.NoError(t, svc.HasRequiredRole(services.Role{Rank: managerRole.Rank}, models.RoleManager))
}

func TestAuthorization_Scopes(t *testing.T) {
	t.Parallel()

	svc, err := services.NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	req := []models.Scope{models.ScopeTeamSettingsWrite}
	assert.ErrorContains(t, svc.HasRequiredScopes([]string{}, req), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredScopes([]string{string("")}, req), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredScopes([]string{string(models.ScopeUsersRead)}, req), "access restricted")
	assert.NoError(t, svc.HasRequiredScopes([]string{string(models.ScopeTeamSettingsWrite)}, req))

	req = []models.Scope{models.ScopeTeamSettingsWrite, models.ScopeUsersRead}
	assert.ErrorContains(t, svc.HasRequiredScopes([]string{string(models.ScopeUsersRead)}, req), "access restricted")
}
