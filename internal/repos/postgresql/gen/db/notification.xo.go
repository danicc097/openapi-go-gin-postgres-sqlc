package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Notification represents a row from 'public.notifications'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Notification struct {
	NotificationID   int              `json:"notificationID" db:"notification_id" required:"true"`                                                 // notification_id
	ReceiverRank     *int16           `json:"receiverRank" db:"receiver_rank" required:"true"`                                                     // receiver_rank
	Title            string           `json:"title" db:"title" required:"true"`                                                                    // title
	Body             string           `json:"body" db:"body" required:"true"`                                                                      // body
	Label            string           `json:"label" db:"label" required:"true"`                                                                    // label
	Link             *string          `json:"link" db:"link" required:"true"`                                                                      // link
	CreatedAt        time.Time        `json:"createdAt" db:"created_at" required:"true"`                                                           // created_at
	Sender           uuid.UUID        `json:"sender" db:"sender" required:"true"`                                                                  // sender
	Receiver         *uuid.UUID       `json:"receiver" db:"receiver" required:"true"`                                                              // receiver
	NotificationType NotificationType `json:"notificationType" db:"notification_type" required:"true" ref:"#/components/schemas/NotificationType"` // notification_type

	UserNotification *UserNotification `json:"userNotification" db:"user_notification"` // O2O
	// xo fields
	_exists, _deleted bool
}

// NotificationCreateParams represents insert params for 'public.notifications'
type NotificationCreateParams struct {
	ReceiverRank     *int16           `json:"receiverRank"`     // receiver_rank
	Title            string           `json:"title"`            // title
	Body             string           `json:"body"`             // body
	Label            string           `json:"label"`            // label
	Link             *string          `json:"link"`             // link
	Sender           uuid.UUID        `json:"sender"`           // sender
	Receiver         *uuid.UUID       `json:"receiver"`         // receiver
	NotificationType NotificationType `json:"notificationType"` // notification_type
}

// NotificationUpdateParams represents update params for 'public.notifications'
type NotificationUpdateParams struct {
	ReceiverRank     **int16           `json:"receiverRank"`     // receiver_rank
	Title            *string           `json:"title"`            // title
	Body             *string           `json:"body"`             // body
	Label            *string           `json:"label"`            // label
	Link             **string          `json:"link"`             // link
	Sender           *uuid.UUID        `json:"sender"`           // sender
	Receiver         **uuid.UUID       `json:"receiver"`         // receiver
	NotificationType *NotificationType `json:"notificationType"` // notification_type
}

type NotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   NotificationJoins
}
type NotificationSelectConfigOption func(*NotificationSelectConfig)

// WithNotificationLimit limits row selection.
func WithNotificationLimit(limit int) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type NotificationOrderBy = string

const (
	NotificationCreatedAtDescNullsFirst NotificationOrderBy = " created_at DESC NULLS FIRST "
	NotificationCreatedAtDescNullsLast  NotificationOrderBy = " created_at DESC NULLS LAST "
	NotificationCreatedAtAscNullsFirst  NotificationOrderBy = " created_at ASC NULLS FIRST "
	NotificationCreatedAtAscNullsLast   NotificationOrderBy = " created_at ASC NULLS LAST "
)

// WithNotificationOrderBy orders results by the given columns.
func WithNotificationOrderBy(rows ...NotificationOrderBy) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type NotificationJoins struct {
	UserNotification bool
}

// WithNotificationJoin joins with the given tables.
func WithNotificationJoin(joins NotificationJoins) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) (*Notification, error) {
	switch {
	case n._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case n._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.notifications (` +
		`receiver_rank, title, body, label, link, sender, receiver, notification_type` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING * `
	// run
	logf(sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType)

	rows, err := db.Query(ctx, sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/db.Query: %w", err))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/pgx.CollectOneRow: %w", err))
	}
	newn._exists = true
	*n = newn

	return n, nil
}

// Update updates a Notification in the database.
func (n *Notification) Update(ctx context.Context, db DB) (*Notification, error) {
	switch {
	case !n._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case n._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.notifications SET ` +
		`receiver_rank = $1, title = $2, body = $3, label = $4, link = $5, sender = $6, receiver = $7, notification_type = $8 ` +
		`WHERE notification_id = $9 ` +
		`RETURNING * `
	// run
	logf(sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.CreatedAt, n.Sender, n.Receiver, n.NotificationType, n.NotificationID)

	rows, err := db.Query(ctx, sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType, n.NotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/db.Query: %w", err))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/pgx.CollectOneRow: %w", err))
	}
	newn._exists = true
	*n = newn

	return n, nil
}

// Save saves the Notification to the database.
func (n *Notification) Save(ctx context.Context, db DB) (*Notification, error) {
	if n._exists {
		return n.Update(ctx, db)
	}
	return n.Insert(ctx, db)
}

// Upsert performs an upsert for Notification.
func (n *Notification) Upsert(ctx context.Context, db DB) error {
	switch {
	case n._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.notifications (` +
		`notification_id, receiver_rank, title, body, label, link, sender, receiver, notification_type` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)` +
		` ON CONFLICT (notification_id) DO ` +
		`UPDATE SET ` +
		`receiver_rank = EXCLUDED.receiver_rank, title = EXCLUDED.title, body = EXCLUDED.body, label = EXCLUDED.label, link = EXCLUDED.link, sender = EXCLUDED.sender, receiver = EXCLUDED.receiver, notification_type = EXCLUDED.notification_type ` +
		` RETURNING * `
	// run
	logf(sqlstr, n.NotificationID, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType)
	if _, err := db.Exec(ctx, sqlstr, n.NotificationID, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType); err != nil {
		return logerror(err)
	}
	// set exists
	n._exists = true
	return nil
}

// Delete deletes the Notification from the database.
func (n *Notification) Delete(ctx context.Context, db DB) error {
	switch {
	case !n._exists: // doesn't exist
		return nil
	case n._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.notifications ` +
		`WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, n.NotificationID); err != nil {
		return logerror(err)
	}
	// set deleted
	n._deleted = true
	return nil
}

// NotificationByNotificationID retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_pkey'.
func NotificationByNotificationID(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) (*Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`notifications.notification_id,
notifications.receiver_rank,
notifications.title,
notifications.body,
notifications.label,
notifications.link,
notifications.created_at,
notifications.sender,
notifications.receiver,
notifications.notification_type,
(case when $1::boolean = true then row(user_notifications.*) end) as user_notification ` +
		`FROM public.notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey"
left join user_notifications on user_notifications.notification_id = notifications.notification_id` +
		` WHERE notifications.notification_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserNotification, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", err))
	}
	n, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", err))
	}
	n._exists = true
	return &n, nil
}

// NotificationsByReceiverRankNotificationTypeCreatedAt retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_receiver_rank_notification_type_created_at_idx'.
func NotificationsByReceiverRankNotificationTypeCreatedAt(ctx context.Context, db DB, receiverRank *int16, notificationType NotificationType, createdAt time.Time, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`notifications.notification_id,
notifications.receiver_rank,
notifications.title,
notifications.body,
notifications.label,
notifications.link,
notifications.created_at,
notifications.sender,
notifications.receiver,
notifications.notification_type,
(case when $1::boolean = true then row(user_notifications.*) end) as user_notification ` +
		`FROM public.notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey"
left join user_notifications on user_notifications.notification_id = notifications.notification_id` +
		` WHERE notifications.receiver_rank = $2 AND notifications.notification_type = $3 AND notifications.created_at = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, receiverRank, notificationType, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserNotification, receiverRank, notificationType, createdAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKUser_Receiver returns the User associated with the Notification's (Receiver).
//
// Generated from foreign key 'notifications_receiver_fkey'.
func (n *Notification) FKUser_Receiver(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, *n.Receiver)
}

// FKUser_Sender returns the User associated with the Notification's (Sender).
//
// Generated from foreign key 'notifications_sender_fkey'.
func (n *Notification) FKUser_Sender(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, n.Sender)
}
