package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// NotificationPublic represents fields that may be exposed from 'public.notifications'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type NotificationPublic struct {
	NotificationID   int              `json:"notificationID" required:"true"`   // notification_id
	ReceiverRank     *int16           `json:"receiverRank" required:"true"`     // receiver_rank
	Title            string           `json:"title" required:"true"`            // title
	Body             string           `json:"body" required:"true"`             // body
	Label            string           `json:"label" required:"true"`            // label
	Link             *string          `json:"link" required:"true"`             // link
	CreatedAt        time.Time        `json:"createdAt" required:"true"`        // created_at
	Sender           uuid.UUID        `json:"sender" required:"true"`           // sender
	Receiver         *uuid.UUID       `json:"receiver" required:"true"`         // receiver
	NotificationType NotificationType `json:"notificationType" required:"true"` // notification_type
}

// Notification represents a row from 'public.notifications'.
type Notification struct {
	NotificationID   int              `json:"notification_id" db:"notification_id"`     // notification_id
	ReceiverRank     *int16           `json:"receiver_rank" db:"receiver_rank"`         // receiver_rank
	Title            string           `json:"title" db:"title"`                         // title
	Body             string           `json:"body" db:"body"`                           // body
	Label            string           `json:"label" db:"label"`                         // label
	Link             *string          `json:"link" db:"link"`                           // link
	CreatedAt        time.Time        `json:"created_at" db:"created_at"`               // created_at
	Sender           uuid.UUID        `json:"sender" db:"sender"`                       // sender
	Receiver         *uuid.UUID       `json:"receiver" db:"receiver"`                   // receiver
	NotificationType NotificationType `json:"notification_type" db:"notification_type"` // notification_type

	// xo fields
	_exists, _deleted bool
}

func (x *Notification) ToPublic() NotificationPublic {
	return NotificationPublic{
		NotificationID: x.NotificationID, ReceiverRank: x.ReceiverRank, Title: x.Title, Body: x.Body, Label: x.Label, Link: x.Link, CreatedAt: x.CreatedAt, Sender: x.Sender, Receiver: x.Receiver, NotificationType: x.NotificationType,
	}
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

type NotificationJoins struct{}

// WithNotificationJoin orders results by the given columns.
func WithNotificationJoin(joins NotificationJoins) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Notification exists in the database.
func (n *Notification) Exists() bool {
	return n._exists
}

// Deleted returns true when the Notification has been marked for deletion from
// the database.
func (n *Notification) Deleted() bool {
	return n._deleted
}

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) error {
	switch {
	case n._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case n._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.notifications (` +
		`receiver_rank, title, body, label, link, sender, receiver, notification_type` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING notification_id, created_at `
	// run
	logf(sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType)
	if err := db.QueryRow(ctx, sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType).Scan(&n.NotificationID, &n.CreatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	n._exists = true
	return nil
}

// Update updates a Notification in the database.
func (n *Notification) Update(ctx context.Context, db DB) error {
	switch {
	case !n._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case n._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.notifications SET ` +
		`receiver_rank = $1, title = $2, body = $3, label = $4, link = $5, sender = $6, receiver = $7, notification_type = $8 ` +
		`WHERE notification_id = $9 ` +
		`RETURNING notification_id, created_at `
	// run
	logf(sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.CreatedAt, n.Sender, n.Receiver, n.NotificationType, n.NotificationID)
	if err := db.QueryRow(ctx, sqlstr, n.ReceiverRank, n.Title, n.Body, n.Label, n.Link, n.Sender, n.Receiver, n.NotificationType, n.NotificationID).Scan(&n.NotificationID, &n.CreatedAt); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Notification to the database.
func (n *Notification) Save(ctx context.Context, db DB) error {
	if n.Exists() {
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
		`receiver_rank = EXCLUDED.receiver_rank, title = EXCLUDED.title, body = EXCLUDED.body, label = EXCLUDED.label, link = EXCLUDED.link, sender = EXCLUDED.sender, receiver = EXCLUDED.receiver, notification_type = EXCLUDED.notification_type  `
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
	logf(sqlstr, n.NotificationID)
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
notifications.notification_type ` +
		`FROM public.notifications ` +
		`` +
		` WHERE notifications.notification_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, notificationID)
	n := Notification{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, notificationID).Scan(&n.NotificationID, &n.ReceiverRank, &n.Title, &n.Body, &n.Label, &n.Link, &n.CreatedAt, &n.Sender, &n.Receiver, &n.NotificationType); err != nil {
		return nil, logerror(err)
	}
	return &n, nil
}

// NotificationsByReceiverRankNotificationTypeCreatedAt retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_receiver_rank_notification_type_created_at_idx'.
func NotificationsByReceiverRankNotificationTypeCreatedAt(ctx context.Context, db DB, receiverRank *int16, notificationType NotificationType, createdAt time.Time, opts ...NotificationSelectConfigOption) ([]*Notification, error) {
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
notifications.notification_type ` +
		`FROM public.notifications ` +
		`` +
		` WHERE notifications.receiver_rank = $1 AND notifications.notification_type = $2 AND notifications.created_at = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, receiverRank, notificationType, createdAt)
	rows, err := db.Query(ctx, sqlstr, receiverRank, notificationType, createdAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Notification
	for rows.Next() {
		n := Notification{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&n.NotificationID, &n.ReceiverRank, &n.Title, &n.Body, &n.Label, &n.Link, &n.CreatedAt, &n.Sender, &n.Receiver, &n.NotificationType); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &n)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
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
