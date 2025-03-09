package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestHandlers_CreateWorkItemTag(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	requiredProject := models.ProjectNameDemo

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   "valid tag creation",
			status: http.StatusCreated,
			scopes: []models.Scope{models.ScopeWorkItemTagCreate},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			teamf := ff.CreateTeam(context.Background(), servicetestutil.CreateTeamParams{Project: requiredProject})
			userf := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
				TeamIDs:    []models.TeamID{teamf.TeamID},
			})
			require.NoError(t, err)

			witCreateParams := postgresqlrandom.WorkItemTagCreateParams(internal.ProjectIDByName[requiredProject])
			res, err := srv.client.CreateWorkItemTagWithResponse(context.Background(), requiredProject, rest.CreateWorkItemTagRequest{
				WorkItemTagCreateParams: *witCreateParams,
			}, ReqWithAPIKey(userf.APIKey.APIKey))

			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode(), string(res.Body))
		})
	}
}
