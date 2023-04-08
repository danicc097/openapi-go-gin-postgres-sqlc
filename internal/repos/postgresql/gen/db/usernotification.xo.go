package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UserNotification represents a row from 'public.user_notifications'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type UserNotification struct {
	UserNotificationID int64     `json:"userNotificationID" db:"user_notification_id"` // user_notification_id
	NotificationID     int       `json:"notificationID" db:"notification_id"`          // notification_id
	Read               bool      `json:"read" db:"read"`                               // read
	UserID             uuid.UUID `json:"userID" db:"user_id"`                          // user_id

	Notification *Notification `json:"notification" db:"notification"` // O2O
	// xo fields
	_exists, _deleted bool
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserNotificationOrderBy = string

const ()

type UserNotificationJoins struct {
	Notification bool
}

// WithUserNotificationJoin joins with the given tables.
func WithUserNotificationJoin(joins UserNotificationJoins) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the UserNotification exists in the database.
func (un *UserNotification) Exists() bool {
	return un._exists
}

// Deleted returns true when the UserNotification has been marked for deletion from
// the database.
func (un *UserNotification) Deleted() bool {
	return un._deleted
}

// Insert inserts the UserNotification to the database.

func (un *UserNotification) Insert(ctx context.Context, db DB) (*UserNotification, error) {
	switch {
	case un._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case un._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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
	newun._exists = true
	un = &newun

	return un, nil
}

// Update updates a UserNotification in the database.
func (un *UserNotification) Update(ctx context.Context, db DB) (*UserNotification, error) {
	switch {
	case !un._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case un._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
	newun._exists = true
	un = &newun

	return un, nil
}

// Save saves the UserNotification to the database.
func (un *UserNotification) Save(ctx context.Context, db DB) (*UserNotification, error) {
	if un.Exists() {
		return un.Update(ctx, db)
	}
	return un.Insert(ctx, db)
}

// Upsert performs an upsert for UserNotification.
func (un *UserNotification) Upsert(ctx context.Context, db DB) error {
	switch {
	case un._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.user_notifications (` +
		`user_notification_id, notification_id, read, user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (user_notification_id) DO ` +
		`UPDATE SET ` +
		`notification_id = EXCLUDED.notification_id, read = EXCLUDED.read, user_id = EXCLUDED.user_id  `
	// run
	logf(sqlstr, un.UserNotificationID, un.NotificationID, un.Read, un.UserID)
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID, un.NotificationID, un.Read, un.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	un._exists = true
	return nil
}

// Delete deletes the UserNotification from the database.
func (un *UserNotification) Delete(ctx context.Context, db DB) error {
	switch {
	case !un._exists: // doesn't exist
		return nil
	case un._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_notifications ` +
		`WHERE user_notification_id = $1 `
	// run
	logf(sqlstr, un.UserNotificationID)
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID); err != nil {
		return logerror(err)
	}
	// set deleted
	un._deleted = true
	return nil
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
(case when $1::boolean = true then row(notifications.*) end) as notification ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey"
left join notifications on notifications.notification_id = user_notifications.notification_id` +
		` WHERE user_notifications.notification_id = $2 AND user_notifications.user_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, notificationID, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, notificationID, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/pgx.CollectOneRow: %w", err))
	}
	un._exists = true
	return &un, nil
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
(case when $1::boolean = true then row(notifications.*) end) as notification ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey"
left join notifications on notifications.notification_id = user_notifications.notification_id` +
		` WHERE user_notifications.user_notification_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userNotificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, userNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/pgx.CollectOneRow: %w", err))
	}
	un._exists = true
	return &un, nil
}

// UserNotificationsByUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_user_id_idx'.
func UserNotificationsByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserNotificationSelectConfigOption) ([]*UserNotification, error) {
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
(case when $1::boolean = true then row(notifications.*) end) as notification ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey"
left join notifications on notifications.notification_id = user_notifications.notification_id` +
		` WHERE user_notifications.user_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[*UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
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
