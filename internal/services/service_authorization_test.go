package services_test

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestAuthorization_Roles(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	authzsvc := services.NewAuthorization(logger)

	userRole := authzsvc.RoleByName(models.RoleUser)
	managerRole := authzsvc.RoleByName(models.RoleManager)

	require.ErrorContains(t, authzsvc.HasRequiredRole(userRole, models.RoleManager), "access restricted")
	require.ErrorContains(t, authzsvc.HasRequiredRole(userRole, models.RoleAdmin), "access restricted")
	require.ErrorContains(t, authzsvc.HasRequiredRole(managerRole, models.RoleAdmin), "access restricted")
	require.ErrorContains(t, authzsvc.HasRequiredRole(services.Role{}, models.RoleAdmin), "access restricted")

	require.NoError(t, authzsvc.HasRequiredRole(services.Role{Rank: managerRole.Rank, Name: models.RoleManager}, models.RoleManager))
}

func TestAuthorization_Scopes(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	authzsvc := services.NewAuthorization(logger)

	req := models.Scopes{models.ScopeTeamSettingsWrite}
	require.ErrorContains(t, authzsvc.HasRequiredScopes(models.Scopes{}, req), "access restricted")
	require.ErrorContains(t, authzsvc.HasRequiredScopes(models.Scopes{models.ScopeUsersRead}, req), "access restricted")
	require.NoError(t, authzsvc.HasRequiredScopes(models.Scopes{models.ScopeTeamSettingsWrite}, req))

	req = models.Scopes{models.ScopeTeamSettingsWrite, models.ScopeUsersRead}
	require.ErrorContains(t, authzsvc.HasRequiredScopes(models.Scopes{models.ScopeUsersRead}, req), "access restricted")
}
