// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const GetUser = `-- name: GetUser :one
select
  username
  , email
  , role_rank
  , created_at
  , updated_at
  , user_id
  -- case when @get_db_data::boolean then
  --   (user_id)
  -- end as user_id, -- TODO sqlc.yaml overrides sql.NullInt64
from
  users
where (email = LOWER($1)::text
  or $1::text is null)
and (username = $2::text
  or $2::text is null)
and (user_id = $3::uuid
  or $3::uuid is null)
limit 1
`

type GetUserParams struct {
	Email    *string     `db:"email" json:"email"`
	Username *string     `db:"username" json:"username"`
	UserID   pgtype.UUID `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID          `db:"user_id" json:"user_id"`
}

func (q *Queries) GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error) {
	row := db.QueryRow(ctx, GetUser, arg.Email, arg.Username, arg.UserID)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const GetUserPersonalNotificationsByUserID = `-- name: GetUserPersonalNotificationsByUserID :many
select
  user_notifications.user_notification_id, user_notifications.notification_id, user_notifications.read, user_notifications.created_at, user_notifications.user_id
  , notifications.notification_type
  , notifications.sender
  , notifications.title
  , notifications.body
  , notifications.label
  , notifications.link
from
  user_notifications
  inner join notifications using (notification_id)
where
  user_notifications.user_id = $1
  and notifications.notification_type = 'personal'
order by
  user_notifications.created_at desc
limit $2
`

type GetUserPersonalNotificationsByUserIDParams struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	Lim    int32     `db:"lim" json:"lim"`
}

type GetUserPersonalNotificationsByUserIDRow struct {
	UserNotificationID int64              `db:"user_notification_id" json:"user_notification_id"`
	NotificationID     int32              `db:"notification_id" json:"notification_id"`
	Read               bool               `db:"read" json:"read"`
	CreatedAt          pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UserID             uuid.UUID          `db:"user_id" json:"user_id"`
	NotificationType   NotificationType   `db:"notification_type" json:"notification_type"`
	Sender             uuid.UUID          `db:"sender" json:"sender"`
	Title              string             `db:"title" json:"title"`
	Body               string             `db:"body" json:"body"`
	Label              string             `db:"label" json:"label"`
	Link               *string            `db:"link" json:"link"`
}

func (q *Queries) GetUserPersonalNotificationsByUserID(ctx context.Context, db DBTX, arg GetUserPersonalNotificationsByUserIDParams) ([]GetUserPersonalNotificationsByUserIDRow, error) {
	rows, err := db.Query(ctx, GetUserPersonalNotificationsByUserID, arg.UserID, arg.Lim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPersonalNotificationsByUserIDRow{}
	for rows.Next() {
		var i GetUserPersonalNotificationsByUserIDRow
		if err := rows.Scan(
			&i.UserNotificationID,
			&i.NotificationID,
			&i.Read,
			&i.CreatedAt,
			&i.UserID,
			&i.NotificationType,
			&i.Sender,
			&i.Title,
			&i.Body,
			&i.Label,
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

const RegisterNewUser = `-- name: RegisterNewUser :one
insert into users (
  username
  , email
  , role_rank)
values (
  $1
  , $2
  , $3)
returning
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
`

type RegisterNewUserParams struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	RoleRank int16  `db:"role_rank" json:"role_rank"`
}

type RegisterNewUserRow struct {
	UserID    uuid.UUID          `db:"user_id" json:"user_id"`
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

// plpgsql-language-server:disable
func (q *Queries) RegisterNewUser(ctx context.Context, db DBTX, arg RegisterNewUserParams) (RegisterNewUserRow, error) {
	row := db.QueryRow(ctx, RegisterNewUser, arg.Username, arg.Email, arg.RoleRank)
	var i RegisterNewUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const Test = `-- name: Test :exec
select
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
from
  users
`

type TestRow struct {
	UserID    uuid.UUID          `db:"user_id" json:"user_id"`
	Username  string             `db:"username" json:"username"`
	Email     string             `db:"email" json:"email"`
	RoleRank  int16              `db:"role_rank" json:"role_rank"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

// update
//
//	users
//
// set
//
//	username = null
//	, email = COALESCE(LOWER(sqlc.narg('email')) , email)
//
// where
//
//	user_id = @user_id;
func (q *Queries) Test(ctx context.Context, db DBTX) error {
	_, err := db.Exec(ctx, Test)
	return err
}
