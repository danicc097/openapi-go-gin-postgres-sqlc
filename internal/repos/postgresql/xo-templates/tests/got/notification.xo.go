package got

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

// Notification represents a row from 'xo_tests.notifications'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type Notification struct {
	NotificationID int        `json:"notificationID" db:"notification_id" required:"true"` // notification_id
	Body           string     `json:"body" db:"body" required:"true"`                      // body
	Sender         uuid.UUID  `json:"sender" db:"sender" required:"true"`                  // sender
	Receiver       *uuid.UUID `json:"receiver" db:"receiver" required:"true"`              // receiver

	UserReceiverJoin *User `json:"-" db:"user_receiver" openapi-go:"ignore"` // O2O users (generated from M2O)
	UserSenderJoin   *User `json:"-" db:"user_sender" openapi-go:"ignore"`   // O2O users (generated from M2O)
}

// NotificationCreateParams represents insert params for 'xo_tests.notifications'.
type NotificationCreateParams struct {
	Body     string     `json:"body" required:"true"`     // body
	Sender   uuid.UUID  `json:"sender" required:"true"`   // sender
	Receiver *uuid.UUID `json:"receiver" required:"true"` // receiver
}

// CreateNotification creates a new Notification in the database with the given params.
func CreateNotification(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	n := &Notification{
		Body:     params.Body,
		Sender:   params.Sender,
		Receiver: params.Receiver,
	}

	return n.Insert(ctx, db)
}

// NotificationUpdateParams represents update params for 'xo_tests.notifications'
type NotificationUpdateParams struct {
	Body     *string     `json:"body" required:"true"`     // body
	Sender   *uuid.UUID  `json:"sender" required:"true"`   // sender
	Receiver **uuid.UUID `json:"receiver" required:"true"` // receiver
}

// SetUpdateParams updates xo_tests.notifications struct fields with the specified params.
func (n *Notification) SetUpdateParams(params *NotificationUpdateParams) {
	if params.Body != nil {
		n.Body = *params.Body
	}
	if params.Sender != nil {
		n.Sender = *params.Sender
	}
	if params.Receiver != nil {
		n.Receiver = *params.Receiver
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type NotificationOrderBy = string

type NotificationJoins struct {
	UserReceiver bool // O2O users
	UserSender   bool // O2O users
}

// WithNotificationJoin joins with the given tables.
func WithNotificationJoin(joins NotificationJoins) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.joins = NotificationJoins{
			UserReceiver: s.joins.UserReceiver || joins.UserReceiver,
			UserSender:   s.joins.UserSender || joins.UserSender,
		}
	}
}

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) (*Notification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.notifications (` +
		`body, sender, receiver` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, n.Body, n.Sender, n.Receiver)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Sender, n.Receiver)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/db.Query: %w", err))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/pgx.CollectOneRow: %w", err))
	}

	*n = newn

	return n, nil
}

// Update updates a Notification in the database.
func (n *Notification) Update(ctx context.Context, db DB) (*Notification, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.notifications SET ` +
		`body = $1, sender = $2, receiver = $3 ` +
		`WHERE notification_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, n.Body, n.Sender, n.Receiver, n.NotificationID)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Sender, n.Receiver, n.NotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/db.Query: %w", err))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/pgx.CollectOneRow: %w", err))
	}
	*n = newn

	return n, nil
}

// Upsert upserts a Notification in the database.
// Requires appropiate PK(s) to be set beforehand.
func (n *Notification) Upsert(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	var err error

	n.Body = params.Body
	n.Sender = params.Sender
	n.Receiver = params.Receiver

	n, err = n.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			n, err = n.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return n, err
}

// Delete deletes the Notification from the database.
func (n *Notification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.notifications ` +
		`WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, n.NotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// NotificationPaginatedByNotificationID returns a cursor-paginated list of Notification.
func NotificationPaginatedByNotificationID(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`notifications.notification_id,
notifications.body,
notifications.sender,
notifications.receiver,
(case when $1::boolean = true and _users_receivers.user_id is not null then row(_users_receivers.*) end) as user_receiver,
(case when $2::boolean = true and _users_senders.user_id is not null then row(_users_senders.*) end) as user_sender ` +
		`FROM xo_tests.notifications ` +
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join xo_tests.users as _users_receivers on _users_receivers.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join xo_tests.users as _users_senders on _users_senders.user_id = notifications.sender` +
		` WHERE notifications.notification_id > $3 GROUP BY _users_receivers.user_id,
      _users_receivers.user_id,
	notifications.notification_id, 
_users_senders.user_id,
      _users_senders.user_id,
	notifications.notification_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// NotificationByNotificationID retrieves a row from 'xo_tests.notifications' as a Notification.
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
notifications.body,
notifications.sender,
notifications.receiver,
(case when $1::boolean = true and _users_receivers.user_id is not null then row(_users_receivers.*) end) as user_receiver,
(case when $2::boolean = true and _users_senders.user_id is not null then row(_users_senders.*) end) as user_sender ` +
		`FROM xo_tests.notifications ` +
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join xo_tests.users as _users_receivers on _users_receivers.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join xo_tests.users as _users_senders on _users_senders.user_id = notifications.sender` +
		` WHERE notifications.notification_id = $3 GROUP BY _users_receivers.user_id,
      _users_receivers.user_id,
	notifications.notification_id, 
_users_senders.user_id,
      _users_senders.user_id,
	notifications.notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", err))
	}
	n, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", err))
	}

	return &n, nil
}

// NotificationsBySender retrieves a row from 'xo_tests.notifications' as a Notification.
//
// Generated from index 'notifications_sender_idx'.
func NotificationsBySender(ctx context.Context, db DB, sender uuid.UUID, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`notifications.notification_id,
notifications.body,
notifications.sender,
notifications.receiver,
(case when $1::boolean = true and _users_receivers.user_id is not null then row(_users_receivers.*) end) as user_receiver,
(case when $2::boolean = true and _users_senders.user_id is not null then row(_users_senders.*) end) as user_sender ` +
		`FROM xo_tests.notifications ` +
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join xo_tests.users as _users_receivers on _users_receivers.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join xo_tests.users as _users_senders on _users_senders.user_id = notifications.sender` +
		` WHERE notifications.sender = $3 GROUP BY _users_receivers.user_id,
      _users_receivers.user_id,
	notifications.notification_id, 
_users_senders.user_id,
      _users_senders.user_id,
	notifications.notification_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, sender)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, sender)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsBySender/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsBySender/pgx.CollectRows: %w", err))
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
