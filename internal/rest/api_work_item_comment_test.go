package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_DeleteWorkItemComment(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	tests := []struct {
		name   string
		status int
		role   models.Role
		scopes models.Scopes
	}{
		{
			name:   "valid work item comment deletion",
			status: http.StatusNoContent,
			scopes: []models.Scope{models.ScopeWorkItemCommentDelete},
		},
		{
			name:   "unauthorized work item comment call",
			status: http.StatusForbidden,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       tc.role,
				WithAPIKey: true,
				Scopes:     tc.scopes,
			})

			workItemCommentf := ff.CreateWorkItemComment(context.Background(), servicetestutil.CreateWorkItemCommentParams{Project: models.ProjectDemo, UserID: ufixture.User.UserID})

			id := workItemCommentf.WorkItemComment.WorkItemCommentID
			res, err := srv.client.DeleteWorkItemCommentWithResponse(context.Background(), workItemCommentf.WorkItem.WorkItemID, id, ReqWithAPIKey(ufixture.APIKey.APIKey))
			require.NoError(t, err)
			require.Equal(t, tc.status, res.StatusCode(), string(res.Body))
		})
	}
}

func TestHandlers_CreateWorkItemComment(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{models.ScopeWorkItemCommentCreate}

		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		demoWorkItemf := ff.CreateWorkItem(context.Background(), servicetestutil.CreateWorkItemParams{Project: models.ProjectDemo})
		require.NoError(t, err)

		ufixture.User, err = svc.User.AssignTeam(context.Background(), testPool, ufixture.User.UserID, demoWorkItemf.WorkItem.TeamID)
		require.NoError(t, err)

		randomWorkItemCommentCreateParams := postgresqltestutil.RandomWorkItemCommentCreateParams(t, ufixture.User.UserID, demoWorkItemf.WorkItem.WorkItemID)
		body := rest.CreateWorkItemCommentRequest{
			WorkItemCommentCreateParams: *randomWorkItemCommentCreateParams,
		}

		res, err := srv.client.CreateWorkItemCommentWithResponse(context.Background(), int(demoWorkItemf.WorkItem.WorkItemID), body, ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode(), string(res.Body))
		assert.EqualValues(t, randomWorkItemCommentCreateParams.Message, res.JSON201.Message)
		assert.EqualValues(t, randomWorkItemCommentCreateParams.UserID, res.JSON201.UserID)
		assert.EqualValues(t, randomWorkItemCommentCreateParams.WorkItemID, res.JSON201.WorkItemID)
	})
}

func TestHandlers_GetWorkItemComment(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	t.Run("authenticated_user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleUser
		scopes := models.Scopes{} // no scope needed to read

		ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})

		workItemCommentf := ff.CreateWorkItemComment(context.Background(), servicetestutil.CreateWorkItemCommentParams{Project: models.ProjectDemo, UserID: ufixture.User.UserID})

		id := workItemCommentf.WorkItemComment.WorkItemCommentID
		res, err := srv.client.GetWorkItemCommentWithResponse(context.Background(), workItemCommentf.WorkItem.WorkItemID, id, ReqWithAPIKey(ufixture.APIKey.APIKey))

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode(), string(res.Body))

		got, err := json.Marshal(res.JSON200)
		require.NoError(t, err)
		want, err := json.Marshal(&rest.WorkItemComment{WorkItemComment: *workItemCommentf.WorkItemComment})
		require.NoError(t, err)

		assert.JSONEqf(t, string(want), string(got), "") // ignore private JSON fields
	})
}

func TestHandlers_UpdateWorkItemComment(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	demoWorkItemf := ff.CreateWorkItem(context.Background(), servicetestutil.CreateWorkItemParams{Project: models.ProjectDemo})
	require.NoError(t, err)

	ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
		Scopes:     []models.Scope{models.ScopeWorkItemCommentEdit}, // TODO: most crud should be via roles, else cumbersome testing
	})

	ufixture.User, err = svc.User.AssignTeam(context.Background(), testPool, ufixture.User.UserID, demoWorkItemf.WorkItem.TeamID)
	require.NoError(t, err)

	tests := []struct {
		name                    string
		status                  int
		body                    rest.UpdateWorkItemCommentRequest
		validationErrorContains []string
	}{
		{
			name:   "valid work item comment update",
			status: http.StatusOK,
			body: func() rest.UpdateWorkItemCommentRequest {
				randomWorkItemCommentCreateParams := postgresqltestutil.RandomWorkItemCommentCreateParams(t, ufixture.User.UserID, demoWorkItemf.WorkItem.WorkItemID)

				return rest.UpdateWorkItemCommentRequest{
					WorkItemCommentUpdateParams: db.WorkItemCommentUpdateParams{
						Message:    pointers.New(randomWorkItemCommentCreateParams.Message),
						UserID:     pointers.New(randomWorkItemCommentCreateParams.UserID),
						WorkItemID: pointers.New(randomWorkItemCommentCreateParams.WorkItemID),
					},
				}
			}(),
		},
		// NOTE: we do need to test spec validation
		// {
		// 	name:                    "invalid work item comment update param",
		// 	status:                  http.StatusBadRequest,
		// 	body:                    rest.UpdateWorkItemCommentRequest{},
		// 	validationErrorContains: []string{"[\" field <JSON >\"]", "<error>"},
		// },
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			normalUser := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				Role:       models.RoleUser,
				WithAPIKey: true,
				Scopes:     []models.Scope{models.ScopeWorkItemCommentEdit},
			})

			workItemCommentf := ff.CreateWorkItemComment(context.Background(), servicetestutil.CreateWorkItemCommentParams{Project: models.ProjectDemo, UserID: *tc.body.UserID})

			id := workItemCommentf.WorkItemComment.WorkItemCommentID
			updateRes, err := srv.client.UpdateWorkItemCommentWithResponse(context.Background(), workItemCommentf.WorkItem.WorkItemID, id, tc.body, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			require.EqualValues(t, tc.status, updateRes.StatusCode(), string(updateRes.Body))

			if len(tc.validationErrorContains) > 0 {
				for _, ve := range tc.validationErrorContains {
					assert.Contains(t, string(updateRes.Body), ve)
				}

				return
			}

			assert.EqualValues(t, id, updateRes.JSON200.WorkItemCommentID)

			res, err := srv.client.GetWorkItemCommentWithResponse(context.Background(), workItemCommentf.WorkItem.WorkItemID, id, ReqWithAPIKey(normalUser.APIKey.APIKey))

			require.NoError(t, err)
			assert.EqualValues(t, *tc.body.Message, res.JSON200.Message)
			assert.EqualValues(t, *tc.body.UserID, res.JSON200.UserID)
			assert.EqualValues(t, *tc.body.WorkItemID, res.JSON200.WorkItemID)
		})
	}
}
