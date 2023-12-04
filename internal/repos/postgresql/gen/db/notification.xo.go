package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Notification represents a row from 'public.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Notification struct {
	NotificationID   NotificationID   `json:"notificationID" db:"notification_id" required:"true" nullable:"false"`                                                 // notification_id
	ReceiverRank     *int             `json:"-" db:"receiver_rank"`                                                                                                 // receiver_rank
	Title            string           `json:"title" db:"title" required:"true" nullable:"false"`                                                                    // title
	Body             string           `json:"body" db:"body" required:"true" nullable:"false"`                                                                      // body
	Labels           []string         `json:"labels" db:"labels" required:"true" nullable:"false"`                                                                  // labels
	Link             *string          `json:"link" db:"link"`                                                                                                       // link
	CreatedAt        time.Time        `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                                                           // created_at
	Sender           UserID           `json:"sender" db:"sender" required:"true" nullable:"false"`                                                                  // sender
	Receiver         *UserID          `json:"receiver" db:"receiver"`                                                                                               // receiver
	NotificationType NotificationType `json:"notificationType" db:"notification_type" required:"true" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type

	ReceiverJoin                      *User               `json:"-" db:"user_receiver" openapi-go:"ignore"`      // O2O users (generated from M2O)
	SenderJoin                        *User               `json:"-" db:"user_sender" openapi-go:"ignore"`        // O2O users (generated from M2O)
	NotificationUserNotificationsJoin *[]UserNotification `json:"-" db:"user_notifications" openapi-go:"ignore"` // M2O notifications

}

// NotificationCreateParams represents insert params for 'public.notifications'.
type NotificationCreateParams struct {
	Body             string           `json:"body" required:"true" nullable:"false"`                                                         // body
	Labels           []string         `json:"labels" required:"true" nullable:"false"`                                                       // labels
	Link             *string          `json:"link"`                                                                                          // link
	NotificationType NotificationType `json:"notificationType" required:"true" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type
	Receiver         *UserID          `json:"receiver"`                                                                                      // receiver
	ReceiverRank     *int             `json:"-"`                                                                                             // receiver_rank
	Sender           UserID           `json:"sender" required:"true" nullable:"false"`                                                       // sender
	Title            string           `json:"title" required:"true" nullable:"false"`                                                        // title
}

type NotificationID int

// CreateNotification creates a new Notification in the database with the given params.
func CreateNotification(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	n := &Notification{
		Body:             params.Body,
		Labels:           params.Labels,
		Link:             params.Link,
		NotificationType: params.NotificationType,
		Receiver:         params.Receiver,
		ReceiverRank:     params.ReceiverRank,
		Sender:           params.Sender,
		Title:            params.Title,
	}

	return n.Insert(ctx, db)
}

// NotificationUpdateParams represents update params for 'public.notifications'.
type NotificationUpdateParams struct {
	Body             *string           `json:"body" nullable:"false"`                                                         // body
	Labels           *[]string         `json:"labels" nullable:"false"`                                                       // labels
	Link             **string          `json:"link"`                                                                          // link
	NotificationType *NotificationType `json:"notificationType" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type
	Receiver         **UserID          `json:"receiver"`                                                                      // receiver
	ReceiverRank     **int             `json:"-"`                                                                             // receiver_rank
	Sender           *UserID           `json:"sender" nullable:"false"`                                                       // sender
	Title            *string           `json:"title" nullable:"false"`                                                        // title
}

// SetUpdateParams updates public.notifications struct fields with the specified params.
func (n *Notification) SetUpdateParams(params *NotificationUpdateParams) {
	if params.Body != nil {
		n.Body = *params.Body
	}
	if params.Labels != nil {
		n.Labels = *params.Labels
	}
	if params.Link != nil {
		n.Link = *params.Link
	}
	if params.NotificationType != nil {
		n.NotificationType = *params.NotificationType
	}
	if params.Receiver != nil {
		n.Receiver = *params.Receiver
	}
	if params.ReceiverRank != nil {
		n.ReceiverRank = *params.ReceiverRank
	}
	if params.Sender != nil {
		n.Sender = *params.Sender
	}
	if params.Title != nil {
		n.Title = *params.Title
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

type NotificationOrderBy string

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
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
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

// WithNotificationFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
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

const notificationTableUserReceiverJoinSQL = `-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
`

const notificationTableUserReceiverSelectSQL = `(case when _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver`

const notificationTableUserReceiverGroupBySQL = `_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id`

const notificationTableUserSenderJoinSQL = `-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join users as _notifications_sender on _notifications_sender.user_id = notifications.sender
`

const notificationTableUserSenderSelectSQL = `(case when _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender`

const notificationTableUserSenderGroupBySQL = `_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id`

const notificationTableUserNotificationsJoinSQL = `-- M2O join generated from "user_notifications_notification_id_fkey"
left join (
  select
  notification_id as user_notifications_notification_id
    , array_agg(user_notifications.*) as user_notifications
  from
    user_notifications
  group by
        notification_id
) as joined_user_notifications on joined_user_notifications.user_notifications_notification_id = notifications.notification_id
`

const notificationTableUserNotificationsSelectSQL = `COALESCE(joined_user_notifications.user_notifications, '{}') as user_notifications`

const notificationTableUserNotificationsGroupBySQL = `joined_user_notifications.user_notifications, notifications.notification_id`

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) (*Notification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.notifications (
	body, labels, link, notification_type, receiver, receiver_rank, sender, title
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8
	) RETURNING * `
	// run
	logf(sqlstr, n.Body, n.Labels, n.Link, n.NotificationType, n.Receiver, n.ReceiverRank, n.Sender, n.Title)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Labels, n.Link, n.NotificationType, n.Receiver, n.ReceiverRank, n.Sender, n.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	*n = newn

	return n, nil
}

// Update updates a Notification in the database.
func (n *Notification) Update(ctx context.Context, db DB) (*Notification, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.notifications SET 
	body = $1, labels = $2, link = $3, notification_type = $4, receiver = $5, receiver_rank = $6, sender = $7, title = $8 
	WHERE notification_id = $9 
	RETURNING * `
	// run
	logf(sqlstr, n.Body, n.CreatedAt, n.Labels, n.Link, n.NotificationType, n.Receiver, n.ReceiverRank, n.Sender, n.Title, n.NotificationID)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Labels, n.Link, n.NotificationType, n.Receiver, n.ReceiverRank, n.Sender, n.Title, n.NotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}
	*n = newn

	return n, nil
}

// Upsert upserts a Notification in the database.
// Requires appropriate PK(s) to be set beforehand.
func (n *Notification) Upsert(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	var err error

	n.Body = params.Body
	n.Labels = params.Labels
	n.Link = params.Link
	n.NotificationType = params.NotificationType
	n.Receiver = params.Receiver
	n.ReceiverRank = params.ReceiverRank
	n.Sender = params.Sender
	n.Title = params.Title

	n, err = n.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Notification", Err: err})
			}
			n, err = n.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Notification", Err: err})
			}
		}
	}

	return n, err
}

// Delete deletes the Notification from the database.
func (n *Notification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.notifications 
	WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, n.NotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// NotificationPaginatedByNotificationID returns a cursor-paginated list of Notification.
func NotificationPaginatedByNotificationID(ctx context.Context, db DB, notificationID NotificationID, direction Direction, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, notificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, notificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserSenderGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, notificationTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserNotificationsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.created_at,
	notifications.labels,
	notifications.link,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.receiver_rank,
	notifications.sender,
	notifications.title %s 
	 FROM public.notifications %s 
	 WHERE notifications.notification_id %s $1
	 %s   %s 
  ORDER BY 
		notification_id %s `, selects, joins, operator, filters, groupbys, direction)
	sqlstr += c.limit
	sqlstr = "/* NotificationPaginatedByNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
	}
	return res, nil
}

// NotificationByNotificationID retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_pkey'.
func NotificationByNotificationID(ctx context.Context, db DB, notificationID NotificationID, opts ...NotificationSelectConfigOption) (*Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, notificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, notificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserSenderGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, notificationTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserNotificationsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.created_at,
	notifications.labels,
	notifications.link,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.receiver_rank,
	notifications.sender,
	notifications.title %s 
	 FROM public.notifications %s 
	 WHERE notifications.notification_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* NotificationByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	n, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	return &n, nil
}

// NotificationsByReceiverRankNotificationTypeCreatedAt retrieves a row from 'public.notifications' as a Notification.
//
// Generated from index 'notifications_receiver_rank_notification_type_created_at_idx'.
func NotificationsByReceiverRankNotificationTypeCreatedAt(ctx context.Context, db DB, receiverRank *int, notificationType NotificationType, createdAt time.Time, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, notificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, notificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserSenderGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, notificationTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, notificationTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, notificationTableUserNotificationsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.created_at,
	notifications.labels,
	notifications.link,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.receiver_rank,
	notifications.sender,
	notifications.title %s 
	 FROM public.notifications %s 
	 WHERE notifications.receiver_rank = $1 AND notifications.notification_type = $2 AND notifications.created_at = $3
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* NotificationsByReceiverRankNotificationTypeCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, receiverRank, notificationType, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{receiverRank, notificationType, createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsByReceiverRankNotificationTypeCreatedAt/Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsByReceiverRankNotificationTypeCreatedAt/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
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
