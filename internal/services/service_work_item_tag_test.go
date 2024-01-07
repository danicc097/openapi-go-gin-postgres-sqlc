package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestWorkItemTag_Update(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	requiredProject := models.ProjectDemo

	svc := services.New(logger, services.CreateTestRepos(), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	team, err := svc.Team.Create(context.Background(), testPool, postgresqltestutil.RandomTeamCreateParams(t, internal.ProjectIDByName[requiredProject]))
	require.NoError(t, err)
	tagCreator, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
	})
	require.NoError(t, err)

	err = svc.User.AssignTeam(context.Background(), testPool, tagCreator.User.UserID, team.TeamID)
	require.NoError(t, err)

	witCreateParams := postgresqltestutil.RandomWorkItemTagCreateParams(t, internal.ProjectIDByName[requiredProject])
	wit, err := svc.WorkItemTag.Create(context.Background(), testPool, services.CtxUser{User: *tagCreator.User}, witCreateParams)
	require.NoError(t, err)

	type args struct {
		params            *db.WorkItemTagUpdateParams
		id                db.WorkItemTagID
		withUserInProject bool
	}
	type want struct {
		Name *string
	}

	tests := []struct {
		name          string
		args          args
		want          want
		errorContains string
	}{
		{
			name: "updated correctly",
			args: args{
				params: &db.WorkItemTagUpdateParams{
					Name: pointers.New("changed"),
				},
				withUserInProject: true,
				id:                wit.WorkItemTagID,
			},
			want: want{
				Name: pointers.New("changed"),
			},
		},
		{
			name: "user not in project",
			args: args{
				params:            &db.WorkItemTagUpdateParams{},
				withUserInProject: false,
				id:                wit.WorkItemTagID,
			},
			errorContains: "user is not a member of project",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos()
			repos.Notification = repostesting.NewFakeNotification()

			ctx := context.Background()
			tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
			defer tx.Rollback(ctx)

			user, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})
			require.NoError(t, err)

			if tc.args.withUserInProject {
				err = svc.User.AssignTeam(context.Background(), testPool, user.User.UserID, team.TeamID)
				require.NoError(t, err)
			}

			w := services.NewWorkItemTag(logger, repos)
			got, err := w.Update(ctx, tx, services.CtxUser{User: *user.User}, tc.args.id, tc.args.params)
			if (err != nil) && tc.errorContains == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.errorContains != "" {
				require.Error(t, err)

				assert.Contains(t, err.Error(), tc.errorContains)

				return
			}

			assert.Equal(t, *tc.want.Name, got.Name)
		})
	}
}
