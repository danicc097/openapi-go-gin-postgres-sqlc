package queries

import (
	"github.com/google/uuid"

	. "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

func GetUserNotificationsByUserID(userID uuid.UUID) SelectStatement {
	return SELECT(
		UserNotifications.AllColumns,
		Notifications.AllColumns,
	).FROM(
		UserNotifications.
			INNER_JOIN(Notifications, Notifications.NotificationID.EQ(UserNotifications.NotificationID)),
	).WHERE(
		UserNotifications.UserID.EQ(UUID(userID)),
	).ORDER_BY(
		Notifications.CreatedAt.DESC(),
	)
}
