package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Notification represents a row from 'xo_tests.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Notification struct {
	NotificationID int        `json:"notificationID" db:"notification_id" required:"true"` // notification_id
	Body           string     `json:"-" db:"body" pattern:"^[A-Za-z0-9]*$"`                // body
	Sender         uuid.UUID  `json:"sender" db:"sender" required:"true"`                  // sender
	Receiver       *uuid.UUID `json:"receiver" db:"receiver" required:"true"`              // receiver

	ReceiverJoin *User `json:"-" db:"user_receiver" openapi-go:"ignore"` // O2O users (generated from M2O)
	SenderJoin   *User `json:"-" db:"user_sender" openapi-go:"ignore"`   // O2O users (generated from M2O)
}

// NotificationCreateParams represents insert params for 'xo_tests.notifications'.
type NotificationCreateParams struct {
	Body     string     `json:"-" pattern:"^[A-Za-z0-9]*$"` // body
	Sender   uuid.UUID  `json:"sender" required:"true"`     // sender
	Receiver *uuid.UUID `json:"receiver" required:"true"`   // receiver
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

// NotificationUpdateParams represents update params for 'xo_tests.notifications'.
type NotificationUpdateParams struct {
	Body     *string     `json:"-" pattern:"^[A-Za-z0-9]*$"` // body
	Sender   *uuid.UUID  `json:"sender" required:"true"`     // sender
	Receiver **uuid.UUID `json:"receiver" required:"true"`   // receiver
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

const notificationTableUserReceiverJoinSQL = `-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join xo_tests.users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
`

const notificationTableUserReceiverSelectSQL = `(case when _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver`

const notificationTableUserReceiverGroupBySQL = `_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id`

const notificationTableUserSenderJoinSQL = `-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join xo_tests.users as _notifications_sender on _notifications_sender.user_id = notifications.sender
`

const notificationTableUserSenderSelectSQL = `(case when _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender`

const notificationTableUserSenderGroupBySQL = `_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id`

// Insert inserts the Notification to the database.
func (n *Notification) Insert(ctx context.Context, db DB) (*Notification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.notifications (
	body, sender, receiver
	) VALUES (
	$1, $2, $3
	) RETURNING * `
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
	sqlstr := `UPDATE xo_tests.notifications SET 
	body = $1, sender = $2, receiver = $3 
	WHERE notification_id = $4 
	RETURNING * `
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
	sqlstr := `DELETE FROM xo_tests.notifications 
	WHERE notification_id = $1 `
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
	notifications.notification_id,
	notifications.body,
	notifications.sender,
	notifications.receiver %s 
	 FROM xo_tests.notifications %s 
	 WHERE notifications.notification_id > $1
	 %s   %s 
  ORDER BY 
		notification_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* NotificationPaginatedByNotificationIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
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
	notifications.notification_id,
	notifications.body,
	notifications.sender,
	notifications.receiver %s 
	 FROM xo_tests.notifications %s 
	 WHERE notifications.notification_id < $1
	 %s   %s 
  ORDER BY 
		notification_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* NotificationPaginatedByNotificationIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// NotificationByNotificationID retrieves a row from 'xo_tests.notifications' as a Notification.
//
// Generated from index 'notifications_pkey'.
func NotificationByNotificationID(ctx context.Context, db DB, notificationID int, opts ...NotificationSelectConfigOption) (*Notification, error) {
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
	notifications.notification_id,
	notifications.body,
	notifications.sender,
	notifications.receiver %s 
	 FROM xo_tests.notifications %s 
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
	notifications.notification_id,
	notifications.body,
	notifications.sender,
	notifications.receiver %s 
	 FROM xo_tests.notifications %s 
	 WHERE notifications.sender = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* NotificationsBySender */\n" + sqlstr

	// run
	// logf(sqlstr, sender)
	rows, err := db.Query(ctx, sqlstr, append([]any{sender}, filterParams...)...)
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
