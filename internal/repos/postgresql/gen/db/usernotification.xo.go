package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// UserNotification represents a row from 'public.user_notifications'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type UserNotification struct {
	UserNotificationID int64     `json:"userNotificationID" db:"user_notification_id" required:"true"` // user_notification_id
	NotificationID     int       `json:"notificationID" db:"notification_id" required:"true"`          // notification_id
	Read               bool      `json:"read" db:"read" required:"true"`                               // read
	UserID             uuid.UUID `json:"userID" db:"user_id" required:"true"`                          // user_id

	NotificationJoin *Notification `json:"-" db:"notification_notification_id" openapi-go:"ignore"` // O2O notifications (generated from M2O)
	UserJoin         *User         `json:"-" db:"user_user_id" openapi-go:"ignore"`                 // O2O users (generated from M2O)

}

// UserNotificationCreateParams represents insert params for 'public.user_notifications'.
type UserNotificationCreateParams struct {
	NotificationID int       `json:"notificationID" required:"true"` // notification_id
	Read           bool      `json:"read" required:"true"`           // read
	UserID         uuid.UUID `json:"userID" required:"true"`         // user_id
}

// CreateUserNotification creates a new UserNotification in the database with the given params.
func CreateUserNotification(ctx context.Context, db DB, params *UserNotificationCreateParams) (*UserNotification, error) {
	un := &UserNotification{
		NotificationID: params.NotificationID,
		Read:           params.Read,
		UserID:         params.UserID,
	}

	return un.Insert(ctx, db)
}

// UserNotificationUpdateParams represents update params for 'public.user_notifications'
type UserNotificationUpdateParams struct {
	NotificationID *int       `json:"notificationID" required:"true"` // notification_id
	Read           *bool      `json:"read" required:"true"`           // read
	UserID         *uuid.UUID `json:"userID" required:"true"`         // user_id
}

// SetUpdateParams updates public.user_notifications struct fields with the specified params.
func (un *UserNotification) SetUpdateParams(params *UserNotificationUpdateParams) {
	if params.NotificationID != nil {
		un.NotificationID = *params.NotificationID
	}
	if params.Read != nil {
		un.Read = *params.Read
	}
	if params.UserID != nil {
		un.UserID = *params.UserID
	}
}

type UserNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   UserNotificationJoins
}
type UserNotificationSelectConfigOption func(*UserNotificationSelectConfig)

// WithUserNotificationLimit limits row selection.
func WithUserNotificationLimit(limit int) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserNotificationOrderBy = string

const ()

type UserNotificationJoins struct {
	Notification bool // O2O notifications
	User         bool // O2O users
}

// WithUserNotificationJoin joins with the given tables.
func WithUserNotificationJoin(joins UserNotificationJoins) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.joins = UserNotificationJoins{
			Notification: s.joins.Notification || joins.Notification,
			User:         s.joins.User || joins.User,
		}
	}
}

// Insert inserts the UserNotification to the database.
func (un *UserNotification) Insert(ctx context.Context, db DB) (*UserNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_notifications (` +
		`notification_id, read, user_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/db.Query: %w", err))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/pgx.CollectOneRow: %w", err))
	}

	*un = newun

	return un, nil
}

// Update updates a UserNotification in the database.
func (un *UserNotification) Update(ctx context.Context, db DB) (*UserNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.user_notifications SET ` +
		`notification_id = $1, read = $2, user_id = $3 ` +
		`WHERE user_notification_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/db.Query: %w", err))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/pgx.CollectOneRow: %w", err))
	}
	*un = newun

	return un, nil
}

// Upsert upserts a UserNotification in the database.
// Requires appropiate PK(s) to be set beforehand.
func (un *UserNotification) Upsert(ctx context.Context, db DB, params *UserNotificationCreateParams) (*UserNotification, error) {
	var err error

	un.NotificationID = params.NotificationID
	un.Read = params.Read
	un.UserID = params.UserID

	un, err = un.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			un, err = un.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return un, err
}

// Delete deletes the UserNotification from the database.
func (un *UserNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_notifications ` +
		`WHERE user_notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserNotificationPaginatedByUserNotificationIDAsc returns a cursor-paginated list of UserNotification in Asc order.
func UserNotificationPaginatedByUserNotificationIDAsc(ctx context.Context, db DB, userNotificationID int64, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_notification_id > $3 GROUP BY user_notifications.user_notification_id, 
user_notifications.notification_id, 
user_notifications.read, 
user_notifications.user_id, 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id ORDER BY 
		user_notification_id Asc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationPaginatedByNotificationIDAsc returns a cursor-paginated list of UserNotification in Asc order.
func UserNotificationPaginatedByNotificationIDAsc(ctx context.Context, db DB, notificationID int, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id > $3 GROUP BY user_notifications.user_notification_id, 
user_notifications.notification_id, 
user_notifications.read, 
user_notifications.user_id, 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id ORDER BY 
		notification_id Asc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationPaginatedByUserNotificationIDDesc returns a cursor-paginated list of UserNotification in Desc order.
func UserNotificationPaginatedByUserNotificationIDDesc(ctx context.Context, db DB, userNotificationID int64, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_notification_id < $3 GROUP BY user_notifications.user_notification_id, 
user_notifications.notification_id, 
user_notifications.read, 
user_notifications.user_id, 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id ORDER BY 
		user_notification_id Desc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationPaginatedByNotificationIDDesc returns a cursor-paginated list of UserNotification in Desc order.
func UserNotificationPaginatedByNotificationIDDesc(ctx context.Context, db DB, notificationID int, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id < $3 GROUP BY user_notifications.user_notification_id, 
user_notifications.notification_id, 
user_notifications.read, 
user_notifications.user_id, 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id ORDER BY 
		notification_id Desc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationByNotificationIDUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationByNotificationIDUserID(ctx context.Context, db DB, notificationID int, userID uuid.UUID, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id = $3 AND user_notifications.user_id = $4 GROUP BY 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/pgx.CollectOneRow: %w", err))
	}

	return &un, nil
}

// UserNotificationsByNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationsByNotificationID(ctx context.Context, db DB, notificationID int, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id = $3 GROUP BY 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationByNotificationIDUserID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationByNotificationIDUserID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationByUserNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_pkey'.
func UserNotificationByUserNotificationID(ctx context.Context, db DB, userNotificationID int64, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_notification_id = $3 GROUP BY 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userNotificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/pgx.CollectOneRow: %w", err))
	}

	return &un, nil
}

// UserNotificationsByUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_user_id_idx'.
func UserNotificationsByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id,
(case when $2::boolean = true and _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_id = $3 GROUP BY 
_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id, 
_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationsByUserID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationsByUserID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKNotification_NotificationID returns the Notification associated with the UserNotification's (NotificationID).
//
// Generated from foreign key 'user_notifications_notification_id_fkey'.
func (un *UserNotification) FKNotification_NotificationID(ctx context.Context, db DB) (*Notification, error) {
	return NotificationByNotificationID(ctx, db, un.NotificationID)
}

// FKUser_UserID returns the User associated with the UserNotification's (UserID).
//
// Generated from foreign key 'user_notifications_user_id_fkey'.
func (un *UserNotification) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, un.UserID)
}
