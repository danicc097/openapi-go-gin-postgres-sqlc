package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Notification represents a row from 'public.notifications'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
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

	ReceiverJoin                      *User               `json:"-" db:"user_receiver" openapi-go:"ignore"`      // O2O users (generated from M2O)
	SenderJoin                        *User               `json:"-" db:"user_sender" openapi-go:"ignore"`        // O2O users (generated from M2O)
	NotificationUserNotificationsJoin *[]UserNotification `json:"-" db:"user_notifications" openapi-go:"ignore"` // M2O notifications

}

// NotificationCreateParams represents insert params for 'public.notifications'.
type NotificationCreateParams struct {
	ReceiverRank     *int16           `json:"receiverRank" required:"true"`                                                 // receiver_rank
	Title            string           `json:"title" required:"true"`                                                        // title
	Body             string           `json:"body" required:"true"`                                                         // body
	Label            string           `json:"label" required:"true"`                                                        // label
	Link             *string          `json:"link" required:"true"`                                                         // link
	Sender           uuid.UUID        `json:"sender" required:"true"`                                                       // sender
	Receiver         *uuid.UUID       `json:"receiver" required:"true"`                                                     // receiver
	NotificationType NotificationType `json:"notificationType" required:"true" ref:"#/components/schemas/NotificationType"` // notification_type
}

// CreateNotification creates a new Notification in the database with the given params.
func CreateNotification(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	n := &Notification{
		ReceiverRank:     params.ReceiverRank,
		Title:            params.Title,
		Body:             params.Body,
		Label:            params.Label,
		Link:             params.Link,
		Sender:           params.Sender,
		Receiver:         params.Receiver,
		NotificationType: params.NotificationType,
	}

	return n.Insert(ctx, db)
}

// NotificationUpdateParams represents update params for 'public.notifications'
type NotificationUpdateParams struct {
	ReceiverRank     **int16           `json:"receiverRank" required:"true"`                                                 // receiver_rank
	Title            *string           `json:"title" required:"true"`                                                        // title
	Body             *string           `json:"body" required:"true"`                                                         // body
	Label            *string           `json:"label" required:"true"`                                                        // label
	Link             **string          `json:"link" required:"true"`                                                         // link
	Sender           *uuid.UUID        `json:"sender" required:"true"`                                                       // sender
	Receiver         **uuid.UUID       `json:"receiver" required:"true"`                                                     // receiver
	NotificationType *NotificationType `json:"notificationType" required:"true" ref:"#/components/schemas/NotificationType"` // notification_type
}

// SetUpdateParams updates public.notifications struct fields with the specified params.
func (n *Notification) SetUpdateParams(params *NotificationUpdateParams) {
	if params.ReceiverRank != nil {
		n.ReceiverRank = *params.ReceiverRank
	}
	if params.Title != nil {
		n.Title = *params.Title
	}
	if params.Body != nil {
		n.Body = *params.Body
	}
	if params.Label != nil {
		n.Label = *params.Label
	}
	if params.Link != nil {
		n.Link = *params.Link
	}
	if params.Sender != nil {
		n.Sender = *params.Sender
	}
	if params.Receiver != nil {
		n.Receiver = *params.Receiver
	}
	if params.NotificationType != nil {
		n.NotificationType = *params.NotificationType
	}
}

type NotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   NotificationJoins
	filters map[string][]any
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

const (
	NotificationCreatedAtDescNullsFirst NotificationOrderBy = " created_at DESC NULLS FIRST "
	NotificationCreatedAtDescNullsLast  NotificationOrderBy = " created_at DESC NULLS LAST "
	NotificationCreatedAtAscNullsFirst  NotificationOrderBy = " created_at ASC NULLS FIRST "
	NotificationCreatedAtAscNullsLast   NotificationOrderBy = " created_at ASC NULLS LAST "
)

// WithNotificationOrderBy orders results by the given columns.
func WithNotificationOrderBy(rows ...NotificationOrderBy) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type NotificationJoins struct {
	UserReceiver      bool // O2O users
	UserSender        bool // O2O users
	UserNotifications bool // M2O user_notifications
}

// WithNotificationJoin joins with the given tables.
func WithNotificationJoin(joins NotificationJoins) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.joins = NotificationJoins{
			UserReceiver:      s.joins.UserReceiver || joins.UserReceiver,
			UserSender:        s.joins.UserSender || joins.UserSender,
			UserNotifications: s.joins.UserNotifications || joins.UserNotifications,
		}
	}
}

// WithNotificationFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithNotificationFilters(filters map[string][]any) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) (*Notification, error) {
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

	*n = newn

	return n, nil
}

// Update updates a Notification in the database.
func (n *Notification) Update(ctx context.Context, db DB) (*Notification, error) {
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
	*n = newn

	return n, nil
}

// Upsert upserts a Notification in the database.
// Requires appropiate PK(s) to be set beforehand.
func (n *Notification) Upsert(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	var err error

	n.ReceiverRank = params.ReceiverRank
	n.Title = params.Title
	n.Body = params.Body
	n.Label = params.Label
	n.Link = params.Link
	n.Sender = params.Sender
	n.Receiver = params.Receiver
	n.NotificationType = params.NotificationType

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
	sqlstr := `DELETE FROM public.notifications ` +
		`WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, n.NotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// NotificationPaginatedByNotificationIDAsc returns a cursor-paginated list of Notification in Asc order.
func NotificationPaginatedByNotificationIDAsc(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
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
(case when $1::boolean = true and _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver,
(case when $2::boolean = true and _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender,
(case when $3::boolean = true then COALESCE(joined_user_notifications.user_notifications, '{}') end) as user_notifications `+
		`FROM public.notifications `+
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join users as _notifications_sender on _notifications_sender.user_id = notifications.sender
-- M2O join generated from "user_notifications_notification_id_fkey"
left join (
  select
  notification_id as user_notifications_notification_id
    , array_agg(user_notifications.*) as user_notifications
  from
    user_notifications
  group by
        notification_id) joined_user_notifications on joined_user_notifications.user_notifications_notification_id = notifications.notification_id`+
		` WHERE notifications.notification_id > $4`+
		` %s  GROUP BY notifications.notification_id, 
notifications.receiver_rank, 
notifications.title, 
notifications.body, 
notifications.label, 
notifications.link, 
notifications.created_at, 
notifications.sender, 
notifications.receiver, 
notifications.notification_type, 
_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id, 
_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id, 
joined_user_notifications.user_notifications, notifications.notification_id ORDER BY 
		notification_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, c.joins.UserNotifications, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// NotificationPaginatedByNotificationIDDesc returns a cursor-paginated list of Notification in Desc order.
func NotificationPaginatedByNotificationIDDesc(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
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
(case when $1::boolean = true and _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver,
(case when $2::boolean = true and _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender,
(case when $3::boolean = true then COALESCE(joined_user_notifications.user_notifications, '{}') end) as user_notifications `+
		`FROM public.notifications `+
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join users as _notifications_sender on _notifications_sender.user_id = notifications.sender
-- M2O join generated from "user_notifications_notification_id_fkey"
left join (
  select
  notification_id as user_notifications_notification_id
    , array_agg(user_notifications.*) as user_notifications
  from
    user_notifications
  group by
        notification_id) joined_user_notifications on joined_user_notifications.user_notifications_notification_id = notifications.notification_id`+
		` WHERE notifications.notification_id < $4`+
		` %s  GROUP BY notifications.notification_id, 
notifications.receiver_rank, 
notifications.title, 
notifications.body, 
notifications.label, 
notifications.link, 
notifications.created_at, 
notifications.sender, 
notifications.receiver, 
notifications.notification_type, 
_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id, 
_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id, 
joined_user_notifications.user_notifications, notifications.notification_id ORDER BY 
		notification_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, c.joins.UserNotifications, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// NotificationByNotificationID retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_pkey'.
func NotificationByNotificationID(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) (*Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
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
(case when $1::boolean = true and _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver,
(case when $2::boolean = true and _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender,
(case when $3::boolean = true then COALESCE(joined_user_notifications.user_notifications, '{}') end) as user_notifications `+
		`FROM public.notifications `+
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join users as _notifications_sender on _notifications_sender.user_id = notifications.sender
-- M2O join generated from "user_notifications_notification_id_fkey"
left join (
  select
  notification_id as user_notifications_notification_id
    , array_agg(user_notifications.*) as user_notifications
  from
    user_notifications
  group by
        notification_id) joined_user_notifications on joined_user_notifications.user_notifications_notification_id = notifications.notification_id`+
		` WHERE notifications.notification_id = $4`+
		` %s  GROUP BY 
_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id, 
_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id, 
joined_user_notifications.user_notifications, notifications.notification_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, c.joins.UserNotifications, notificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", err))
	}
	n, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", err))
	}

	return &n, nil
}

// NotificationsByReceiverRankNotificationTypeCreatedAt retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_receiver_rank_notification_type_created_at_idx'.
func NotificationsByReceiverRankNotificationTypeCreatedAt(ctx context.Context, db DB, receiverRank *int16, notificationType NotificationType, createdAt time.Time, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
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
(case when $1::boolean = true and _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver,
(case when $2::boolean = true and _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender,
(case when $3::boolean = true then COALESCE(joined_user_notifications.user_notifications, '{}') end) as user_notifications `+
		`FROM public.notifications `+
		`-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join users as _notifications_sender on _notifications_sender.user_id = notifications.sender
-- M2O join generated from "user_notifications_notification_id_fkey"
left join (
  select
  notification_id as user_notifications_notification_id
    , array_agg(user_notifications.*) as user_notifications
  from
    user_notifications
  group by
        notification_id) joined_user_notifications on joined_user_notifications.user_notifications_notification_id = notifications.notification_id`+
		` WHERE notifications.receiver_rank = $4 AND notifications.notification_type = $5 AND notifications.created_at = $6`+
		` %s  GROUP BY 
_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id, 
_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id, 
joined_user_notifications.user_notifications, notifications.notification_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, receiverRank, notificationType, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.UserReceiver, c.joins.UserSender, c.joins.UserNotifications, receiverRank, notificationType, createdAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsByReceiverRankNotificationTypeCreatedAt/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsByReceiverRankNotificationTypeCreatedAt/pgx.CollectRows: %w", err))
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
