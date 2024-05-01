package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkItemTag_Update(t *testing.T) {
	t.Parallel()

	var err error

	logger := testutil.NewLogger(t)

	requiredProject := models.ProjectNameDemo

	svc := services.New(logger, services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	teamf := ff.CreateTeam(context.Background(), servicetestutil.CreateTeamParams{Project: requiredProject})
	tagCreator := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
	})

	tagCreator.User, err = svc.User.AssignTeam(context.Background(), testPool, tagCreator.UserID, teamf.TeamID)
	require.NoError(t, err)

	witCreateParams := postgresqlrandom.WorkItemTagCreateParams(internal.ProjectIDByName[requiredProject])
	wit, err := svc.WorkItemTag.Create(context.Background(), testPool, *services.NewCtxUser(tagCreator.User), witCreateParams)
	require.NoError(t, err)

	type args struct {
		params            *models.WorkItemTagUpdateParams
		id                models.WorkItemTagID
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
				params: &models.WorkItemTagUpdateParams{
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
				params:            &models.WorkItemTagUpdateParams{},
				withUserInProject: false,
				id:                wit.WorkItemTagID,
			},
			errorContains: "user is not a member of project",
		},
		{
			name: "tag not found",
			args: args{
				params:            &models.WorkItemTagUpdateParams{},
				withUserInProject: true,
				id:                models.WorkItemTagID(-1),
			},
			errorContains: "Work item tag not found",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos(t)
			repos.Notification = repostesting.NewFakeNotification()

			ctx := context.Background()
			tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
			require.NoError(t, err)
			defer tx.Rollback(ctx) // rollback errors should be ignored

			user := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
				WithAPIKey: true,
			})

			if tc.args.withUserInProject {
				user.User, err = svc.User.AssignTeam(context.Background(), testPool, user.UserID, teamf.TeamID)
				require.NoError(t, err)
			}

			w := services.NewWorkItemTag(logger, repos)
			got, err := w.Update(ctx, tx, *services.NewCtxUser(user.User), tc.args.id, tc.args.params)
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
