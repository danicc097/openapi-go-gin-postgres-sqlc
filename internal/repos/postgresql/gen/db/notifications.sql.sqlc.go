// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: notifications.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const GetUserNotifications = `-- name: GetUserNotifications :many
select
  user_notifications.user_notification_id, user_notifications.notification_id, user_notifications.read, user_notifications.user_id
  , notifications.notification_type
  , notifications.sender
  , notifications.title
  , notifications.body
  , notifications.labels
  , notifications.link
from
  user_notifications
  inner join notifications using (notification_id)
where
  user_notifications.user_id = $1
  and notifications.notification_type = $2::notification_type
  and (created_at > $3
    or $3 is null) -- first search null, infinite query using last created_at
order by
  created_at desc
limit $4
`

type GetUserNotificationsParams struct {
	UserID           uuid.UUID        `db:"user_id" json:"user_id"`
	NotificationType NotificationType `db:"notification_type" json:"notification_type"`
	MinCreatedAt     *time.Time       `db:"min_created_at" json:"min_created_at"`
	Lim              *int32           `db:"lim" json:"lim"`
}

type GetUserNotificationsRow struct {
	UserNotificationID int64                `db:"user_notification_id" json:"user_notification_id"`
	NotificationID     int32                `db:"notification_id" json:"notification_id"`
	Read               bool                 `db:"read" json:"read"`
	UserID             uuid.UUID            `db:"user_id" json:"user_id"`
	NotificationType   NotificationType     `db:"notification_type" json:"notification_type"`
	Sender             uuid.UUID            `db:"sender" json:"sender"`
	Title              string               `db:"title" json:"title"`
	Body               string               `db:"body" json:"body"`
	Labels             pgtype.Array[string] `db:"labels" json:"labels"`
	Link               *string              `db:"link" json:"link"`
}

func (q *Queries) GetUserNotifications(ctx context.Context, db DBTX, arg GetUserNotificationsParams) ([]GetUserNotificationsRow, error) {
	rows, err := db.Query(ctx, GetUserNotifications,
		arg.UserID,
		arg.NotificationType,
		arg.MinCreatedAt,
		arg.Lim,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserNotificationsRow{}
	for rows.Next() {
		var i GetUserNotificationsRow
		if err := rows.Scan(
			&i.UserNotificationID,
			&i.NotificationID,
			&i.Read,
			&i.UserID,
			&i.NotificationType,
			&i.Sender,
			&i.Title,
			&i.Body,
			&i.Labels,
			&i.Link,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
