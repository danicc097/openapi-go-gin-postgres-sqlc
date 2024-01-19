package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotification_Create(t *testing.T) {
	t.Parallel()

	notificationRepo := postgresql.NewNotification()

	sender := newRandomUser(t, testPool)

	t.Run("correct_personal_notification", func(t *testing.T) {
		t.Parallel()

		receiver := newRandomUser(t, testPool)

		ncp := postgresqlrandom.NotificationCreateParams(nil, sender.UserID, pointers.New(receiver.UserID), db.NotificationTypePersonal)

		ctx := context.Background()
		// prevent fan out trigger from affecting other tests
		tx, err := testPool.BeginTx(ctx, pgx.TxOptions{})
		require.NoError(t, err)
		defer tx.Rollback(ctx) // rollback errors should be ignored

		_, err = notificationRepo.Create(context.Background(), tx, ncp)
		require.NoError(t, err)

		params := db.GetUserNotificationsParams{UserID: receiver.UserID.UUID, NotificationType: db.NotificationTypePersonal}
		nn, err := notificationRepo.LatestNotifications(context.Background(), tx, &params)
		require.NoError(t, err)

		assert.Equal(t, ncp.Body, nn[0].Body)
	})

	t.Run("correct_global_notification_by_rank", func(t *testing.T) {
		t.Parallel()

		var err error

		receiverRank3 := newRandomUser(t, testPool)
		require.NoError(t, err)
		receiverRank3.RoleRank = 3
		_, err = receiverRank3.Update(context.Background(), testPool)
		require.NoError(t, err)

		receiverRank1 := newRandomUser(t, testPool)
		require.NoError(t, err)
		receiverRank1.RoleRank = 1
		_, err = receiverRank1.Update(context.Background(), testPool)
		require.NoError(t, err)

		receiverRank := pointers.New(3)

		ncp := postgresqlrandom.NotificationCreateParams(receiverRank, sender.UserID, nil, db.NotificationTypeGlobal)

		ctx := context.Background()
		tx, err := testPool.BeginTx(ctx, pgx.TxOptions{}) // prevent fan out trigger from affecting other tests
		require.NoError(t, err)
		defer tx.Rollback(ctx) // rollback errors should be ignored

		_, err = notificationRepo.Create(context.Background(), tx, ncp)
		require.NoError(t, err)

		notificationCount := map[db.UserID]int{
			receiverRank1.UserID: 0,
			receiverRank3.UserID: 1,
		}

		for userID, count := range notificationCount {
			params := db.GetUserNotificationsParams{UserID: userID.UUID, NotificationType: db.NotificationTypeGlobal}
			nn, err := notificationRepo.LatestNotifications(context.Background(), tx, &params)
			require.NoError(t, err)

			assert.Equal(t, count, len(nn))
		}
	})

	t.Run("error_with_no_receiver_with_personal_notification", func(t *testing.T) {
		t.Parallel()

		ncp := postgresqlrandom.NotificationCreateParams(nil, sender.UserID, nil, db.NotificationTypePersonal)

		_, err := notificationRepo.Create(context.Background(), testPool, ncp)
		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})

	t.Run("error_with_no_rank_with_global_notification", func(t *testing.T) {
		t.Parallel()

		ncp := postgresqlrandom.NotificationCreateParams(nil, sender.UserID, nil, db.NotificationTypeGlobal)

		_, err := notificationRepo.Create(context.Background(), testPool, ncp)
		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
