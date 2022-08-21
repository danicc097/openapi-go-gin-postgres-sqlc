package services_test

import (
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAuthorization(t *testing.T) {
	t.Parallel()

	authService := services.NewAuthorization(zaptest.NewLogger(t))

	assert.Equal(t, false, authService.IsAuthorized(db.RoleUser, db.RoleManager))
	assert.Equal(t, false, authService.IsAuthorized(db.RoleUser, db.RoleAdmin))
	assert.Equal(t, false, authService.IsAuthorized(db.RoleManager, db.RoleAdmin))

	for _, r := range db.AllRoleValues() {
		assert.Equal(t, true, authService.IsAuthorized(r, r))
	}

	previousRolePermissions := authService.RolePermissions()[db.RoleUser]
	authService.RolePermissions()[db.RoleUser] = []db.Role{}
	assert.Equal(t, previousRolePermissions, authService.RolePermissions()[db.RoleUser])
}
