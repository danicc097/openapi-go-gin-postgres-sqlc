package services

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAuthorization(t *testing.T) {
	t.Parallel()

	svc, err := NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}
	userRole, ok := svc.roles[models.RoleUser]
	if !ok {
		t.Fatalf("role does not exist: %v", err)
	}
	managerRole, ok := svc.roles[models.RoleUser]
	if !ok {
		t.Fatalf("role does not exist: %v", err)
	}
	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleManager), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(userRole, models.RoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.HasRequiredRole(managerRole, models.RoleAdmin), "access restricted")

	roles := make([]models.Role, 0, len(svc.roles))
	for r := range svc.roles {
		roles = append(roles, r)
	}

	for _, r := range roles {
		assert.NoError(t, svc.HasRequiredRole(svc.roles[r], r))
	}

	// there's no sane reason anyone would do this by accident
	// previousRolePermissions := svc.RolePermissions()[db.UserRoleUser]
	// svc.RolePermissions()[db.UserRoleUser] = []db.UserRole{}
	// assert.Equal(t, previousRolePermissions, svc.RolePermissions()[db.UserRoleUser])
}
