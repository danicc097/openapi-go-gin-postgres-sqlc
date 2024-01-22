package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkItemComment_Update(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)

	requiredProject := models.ProjectDemo

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	team, err := svc.Team.Create(context.Background(), testPool, postgresqlrandom.TeamCreateParams(internal.ProjectIDByName[requiredProject]))
	require.NoError(t, err)
	creator := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
	})

	demoWorkItemf := ff.CreateWorkItem(context.Background(), servicetestutil.CreateWorkItemParams{Project: models.ProjectDemo})
	require.NoError(t, err)

	creator.User, err = svc.User.AssignTeam(context.Background(), testPool, creator.User.UserID, demoWorkItemf.WorkItem.TeamID)
	require.NoError(t, err)

	workItemCommentCreateParams := postgresqlrandom.WorkItemCommentCreateParams(creator.User.UserID, demoWorkItemf.WorkItem.WorkItemID)
	workitemcomment, err := svc.WorkItemComment.Create(context.Background(), testPool, workItemCommentCreateParams)
	require.NoError(t, err)

	type args struct {
		params            *db.WorkItemCommentUpdateParams
		id                db.WorkItemCommentID
		withUserInProject bool
	}

	wantParams := postgresqlrandom.WorkItemCommentCreateParams(creator.User.UserID, demoWorkItemf.WorkItem.WorkItemID)

	tests := []struct {
		name          string
		args          args
		want          db.WorkItemCommentUpdateParams
		errorContains []string
	}{
		{
			name: "updated correctly",
			args: args{
				params: &db.WorkItemCommentUpdateParams{
					Message:    &wantParams.Message,
					UserID:     &wantParams.UserID,
					WorkItemID: &wantParams.WorkItemID,
				},
				withUserInProject: false, //
				id:                workitemcomment.WorkItemCommentID,
			},
			want: db.WorkItemCommentUpdateParams{
				// generating fields based on randomized createparams since it's a superset of updateparams.
				Message:    &wantParams.Message,
				UserID:     &wantParams.UserID,
				WorkItemID: &wantParams.WorkItemID,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos(t)
			repos.Notification = repostesting.NewFakeNotification() // unless we want to test notification integration

			ctx := context.Background()

			tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
			require.NoError(t, err)
			defer tx.Rollback(ctx) // rollback errors should be ignored

			user := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})

			if tc.args.withUserInProject {
				user.User, err = svc.User.AssignTeam(context.Background(), testPool, user.User.UserID, team.TeamID)
				require.NoError(t, err)
			}

			w := services.NewWorkItemComment(logger, repos)
			got, err := w.Update(ctx, tx, tc.args.id, tc.args.params)

			if (err != nil) && len(tc.errorContains) == 0 {
				t.Fatalf("unexpected error = %v", err)
			}

			if len(tc.errorContains) > 0 {
				for _, ve := range tc.errorContains {
					assert.ErrorContains(t, err, ve)
				}

				return
			}

			// loop all fields like in above
			// assert.Equal(t, *tc.want.<Field>, got.<Field>)
			assert.Equal(t, *tc.want.Message, got.Message)
			assert.Equal(t, *tc.want.UserID, got.UserID)
			assert.Equal(t, *tc.want.WorkItemID, got.WorkItemID)
		})
	}
}
