package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func TestWorkItemComment_Update(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	requiredProject := models.ProjectNameDemo

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	type args struct {
		params            *models.WorkItemCommentUpdateParams
		withUserInProject bool
	}

	wantParams := postgresqlrandom.WorkItemCommentCreateParams(models.UserID{UUID: uuid.UUID{}}, models.WorkItemID(-1))

	tests := []struct {
		name          string
		args          args
		want          models.WorkItemCommentUpdateParams
		errorContains []string
	}{
		{
			name: "updated correctly",
			args: args{
				params: &models.WorkItemCommentUpdateParams{
					Message: &wantParams.Message,
				},
				withUserInProject: false, //
			},
			want: models.WorkItemCommentUpdateParams{
				// generating fields based on randomized createparams since it's a superset of updateparams.
				Message: &wantParams.Message,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos(t)
			repos.Notification = repostesting.NewFakeNotification() // unless we want to test notification integration

			ctx := t.Context()

			tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
			require.NoError(t, err)
			defer tx.Rollback(ctx) // rollback errors should be ignored

			teamf := ff.CreateTeam(t.Context(), servicetestutil.CreateTeamParams{Project: requiredProject})
			user := ff.CreateUser(t.Context(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})

			if tc.args.withUserInProject {
				user.User, err = svc.User.AssignTeam(t.Context(), testPool, user.UserID, teamf.TeamID)
				require.NoError(t, err)
			}

			creator := ff.CreateUser(t.Context(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
				TeamIDs:    []models.TeamID{teamf.TeamID},
			})

			demoWorkItemf := ff.CreateWorkItem(t.Context(), requiredProject, *services.NewCtxUser(creator.User), teamf.TeamID)

			workItemCommentCreateParams := postgresqlrandom.WorkItemCommentCreateParams(creator.UserID, demoWorkItemf.WorkItemID)
			workitemcomment, err := svc.WorkItemComment.Create(t.Context(), testPool, workItemCommentCreateParams)
			require.NoError(t, err)

			w := services.NewWorkItemComment(logger, repos)
			got, err := w.Update(ctx, tx, *services.NewCtxUser(user.User), workitemcomment.WorkItemCommentID, tc.args.params)

			if (err != nil) && len(tc.errorContains) == 0 {
				t.Fatalf("unexpected error = %v", err)
			}

			if len(tc.errorContains) > 0 {
				for _, ve := range tc.errorContains {
					require.ErrorContains(t, err, ve)
				}

				return
			}

			assert.Equal(t, creator.UserID, got.UserID)
			assert.Equal(t, demoWorkItemf.WorkItemID, got.WorkItemID)

			// assert.Equal(t, *tc.want.<Field>, got.<Field>)
			assert.Equal(t, *tc.want.Message, got.Message)
		})
	}
}
