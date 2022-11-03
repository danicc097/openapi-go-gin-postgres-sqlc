package services_test

import (
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAuthorization(t *testing.T) {
	t.Parallel()

	svc := services.NewAuthorization(zaptest.NewLogger(t))

	assert.ErrorContains(t, svc.IsAuthorized(db.UserRoleUser, db.UserRoleManager), "access restricted")
	assert.ErrorContains(t, svc.IsAuthorized(db.UserRoleUser, db.UserRoleAdmin), "access restricted")
	assert.ErrorContains(t, svc.IsAuthorized(db.UserRoleManager, db.UserRoleAdmin), "access restricted")

	for _, r := range db.AllUserRoleValues() {
		assert.NoError(t, svc.IsAuthorized(r, r))
	}

	// there's no sane reason anyone would do this by accident
	// previousRolePermissions := svc.RolePermissions()[db.UserRoleUser]
	// svc.RolePermissions()[db.UserRoleUser] = []db.UserRole{}
	// assert.Equal(t, previousRolePermissions, svc.RolePermissions()[db.UserRoleUser])
}
