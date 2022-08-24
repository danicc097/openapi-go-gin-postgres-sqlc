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

	assert.Equal(t, false, svc.IsAuthorized(db.RoleUser, db.RoleManager))
	assert.Equal(t, false, svc.IsAuthorized(db.RoleUser, db.RoleAdmin))
	assert.Equal(t, false, svc.IsAuthorized(db.RoleManager, db.RoleAdmin))

	for _, r := range db.AllRoleValues() {
		assert.Equal(t, true, svc.IsAuthorized(r, r))
	}

	previousRolePermissions := svc.RolePermissions()[db.RoleUser]
	svc.RolePermissions()[db.RoleUser] = []db.Role{}
	assert.Equal(t, previousRolePermissions, svc.RolePermissions()[db.RoleUser])
}
