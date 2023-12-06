package rest

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
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

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   "valid tag creation",
			status: http.StatusCreated,
			role:   models.RoleAdmin,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ufixture, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})
			require.NoError(t, err, "ff.CreateUser: %s")

			project := models.ProjectDemo
			witCreateParams := postgresqltestutil.RandomWorkItemTagCreateParams(t, internal.ProjectIDByName[project])
			res, err := srv.client.CreateWorkItemTagWithResponse(context.Background(), project, models.WorkItemTagCreateRequest{
				Color:       witCreateParams.Color,
				Description: witCreateParams.Description,
				Name:        witCreateParams.Name,
			}, resttestutil.ReqWithAPIKey(ufixture.APIKey.APIKey))
			fmt.Printf("ures.Body: %v\n", string(res.Body))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode())
		})
	}
}
