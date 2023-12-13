package rest_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestHandlers_CreateWorkItemTag(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(testPool, svc)

	requiredProject := models.ProjectDemo

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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			team, err := svc.Team.Create(context.Background(), testPool, postgresqltestutil.RandomTeamCreateParams(t, internal.ProjectIDByName[requiredProject]))
			require.NoError(t, err)
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})
			require.NoError(t, err)

			err = svc.User.AssignTeam(context.Background(), testPool, ufixture.User.UserID, team.TeamID)
			require.NoError(t, err)

			witCreateParams := postgresqltestutil.RandomWorkItemTagCreateParams(t, internal.ProjectIDByName[requiredProject])
			res, err := srv.client.CreateWorkItemTagWithResponse(context.Background(), requiredProject, rest.CreateWorkItemTagRequest{
				WorkItemTagCreateParams: *witCreateParams,
			}, ReqWithAPIKey(ufixture.APIKey.APIKey))
			fmt.Printf("ures.Body: %v\n", string(res.Body))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode())
		})
	}
}
