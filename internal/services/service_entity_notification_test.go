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
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestEntityNotification_Update(t *testing.T) {
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

	projectID := internal.ProjectIDByName[models.ProjectDemo]

	entityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t, projectID)
	entitynotification, err := svc.EntityNotification.Create(context.Background(), testPool, entityNotificationCreateParams)
	require.NoError(t, err)

	type args struct {
		params            *db.EntityNotificationUpdateParams
		id                db.EntityNotificationID
		withUserInProject bool
	}

	randomEntityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t, projectID)

	tests := []struct {
		name          string
		args          args
		want          db.EntityNotificationUpdateParams
		errorContains []string
	}{
		{
			name: "updated correctly",
			args: args{
				params: &db.EntityNotificationUpdateParams{
					ID:        &randomEntityNotificationCreateParams.ID,
					Message:   &randomEntityNotificationCreateParams.Message,
					ProjectID: &randomEntityNotificationCreateParams.ProjectID,
					Topic:     &randomEntityNotificationCreateParams.Topic,
				},
				withUserInProject: false, //
				id:                entitynotification.EntityNotificationID,
			},
			want: db.EntityNotificationUpdateParams{
				// generating fields based on randomized createparams since it's a superset of updateparams.
				ID:        &randomEntityNotificationCreateParams.ID,
				Message:   &randomEntityNotificationCreateParams.Message,
				ProjectID: &randomEntityNotificationCreateParams.ProjectID,
				Topic:     &randomEntityNotificationCreateParams.Topic,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repos := services.CreateTestRepos()
			repos.Notification = repostesting.NewFakeNotification() // unless we want to test notification integration

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

			w := services.NewEntityNotification(logger, repos)
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
			assert.Equal(t, *tc.want.ID, got.ID)
			assert.Equal(t, *tc.want.Message, got.Message)
			assert.Equal(t, *tc.want.ProjectID, got.ProjectID)
			assert.Equal(t, *tc.want.Topic, got.Topic)
		})
	}
}
