package services_test

import (
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAuthorization(t *testing.T) {
	t.Parallel()

	svc := services.NewAuthorization(zaptest.NewLogger(t))

	assert.ErrorContains(t, svc.IsAuthorized(db.RoleUser, db.RoleManager), "access restricted")
	assert.ErrorContains(t, svc.IsAuthorized(db.RoleUser, db.RoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.IsAuthorized(db.RoleManager, db.RoleAdmin), "access restricted")

	for _, r := range db.AllRoleValues() {
		assert.NoError(t, svc.IsAuthorized(r, r))
	}

	// there's no sane reason anyone would do this by accident
	// previousRolePermissions := svc.RolePermissions()[db.RoleUser]
	// svc.RolePermissions()[db.RoleUser] = []db.Role{}
	// assert.Equal(t, previousRolePermissions, svc.RolePermissions()[db.RoleUser])
}
