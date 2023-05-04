package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestNotification_Create(t *testing.T) {
	t.Parallel()

	notificationRepo := postgresql.NewNotification()

	sender, _ := postgresqltestutil.NewRandomUser(t, testPool)

	t.Run("correct_personal_notification", func(t *testing.T) {
		t.Parallel()

		receiver, _ := postgresqltestutil.NewRandomUser(t, testPool)

		ncp := postgresqltestutil.RandomNotificationCreateParams(t, nil, sender.UserID, pointers.New(receiver.UserID), db.NotificationTypePersonal)

		ctx := context.Background()
		tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{}) // prevent fan out trigger from affecting other tests
		defer tx.Rollback(ctx)

		_, err := notificationRepo.Create(context.Background(), tx, ncp)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}

		params := db.GetUserNotificationsParams{UserID: receiver.UserID, NotificationType: db.NotificationTypePersonal}
		nn, err := notificationRepo.LatestUserNotifications(context.Background(), tx, &params)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}
		t.Logf("nn: %v", nn)

		assert.Equal(t, ncp.Body, nn[0].Body)
	})

	t.Run("correct_global_notification_by_rank", func(t *testing.T) {
		t.Parallel()

		receiverRank3, _ := postgresqltestutil.NewRandomUser(t, testPool)
		receiverRank3.RoleRank = 3
		receiverRank3.Update(context.Background(), testPool)
		receiverRank1, _ := postgresqltestutil.NewRandomUser(t, testPool)
		receiverRank1.RoleRank = 1
		receiverRank1.Update(context.Background(), testPool)

		receiverRank := pointers.New[int16](3)

		ncp := postgresqltestutil.RandomNotificationCreateParams(t, receiverRank, sender.UserID, nil, db.NotificationTypeGlobal)

		ctx := context.Background()
		tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{}) // prevent fan out trigger from affecting other tests
		defer tx.Rollback(ctx)

		_, err := notificationRepo.Create(context.Background(), tx, ncp)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}

		notificationCount := map[uuid.UUID]int{
			receiverRank1.UserID: 0,
			receiverRank3.UserID: 1,
		}

		for userID, count := range notificationCount {
			params := db.GetUserNotificationsParams{UserID: userID, NotificationType: db.NotificationTypeGlobal}
			nn, err := notificationRepo.LatestUserNotifications(context.Background(), tx, &params)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}

			assert.Equal(t, count, len(nn))
		}
	})

	t.Run("error_with_no_receiver_with_personal_notification", func(t *testing.T) {
		t.Parallel()

		ncp := postgresqltestutil.RandomNotificationCreateParams(t, nil, sender.UserID, nil, db.NotificationTypePersonal)

		errContains := errViolatesCheckConstraint

		_, err := notificationRepo.Create(context.Background(), testPool, ncp)
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})

	t.Run("error_with_no_rank_with_global_notification", func(t *testing.T) {
		t.Parallel()

		ncp := postgresqltestutil.RandomNotificationCreateParams(t, nil, sender.UserID, nil, db.NotificationTypeGlobal)

		errContains := errViolatesCheckConstraint

		_, err := notificationRepo.Create(context.Background(), testPool, ncp)
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})
}
