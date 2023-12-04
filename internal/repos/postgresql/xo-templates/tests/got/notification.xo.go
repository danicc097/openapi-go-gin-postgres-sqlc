package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Notification represents a row from 'xo_tests.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Notification struct {
	NotificationID NotificationID `json:"notificationID" db:"notification_id" required:"true" nullable:"false"` // notification_id
	Body           string         `json:"-" db:"body" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                // body
	Sender         UserID         `json:"sender" db:"sender" required:"true" nullable:"false"`                  // sender
	Receiver       *UserID        `json:"receiver" db:"receiver"`                                               // receiver

	ReceiverJoin *User `json:"-" db:"user_receiver" openapi-go:"ignore"` // O2O users (generated from M2O)
	SenderJoin   *User `json:"-" db:"user_sender" openapi-go:"ignore"`   // O2O users (generated from M2O)
}

// NotificationCreateParams represents insert params for 'xo_tests.notifications'.
type NotificationCreateParams struct {
	Body     string  `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"` // body
	Receiver *UserID `json:"receiver"`                                    // receiver
	Sender   UserID  `json:"sender" required:"true" nullable:"false"`     // sender
}

type NotificationID int

// CreateNotification creates a new Notification in the database with the given params.
func CreateNotification(ctx context.Context, db DB, params *NotificationCreateParams) (*Notification, error) {
	n := &Notification{
		Body:     params.Body,
		Receiver: params.Receiver,
		Sender:   params.Sender,
	}

	return n.Insert(ctx, db)
}

// NotificationUpdateParams represents update params for 'xo_tests.notifications'.
type NotificationUpdateParams struct {
	Body     *string  `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"` // body
	Receiver **UserID `json:"receiver"`                                    // receiver
	Sender   *UserID  `json:"sender" nullable:"false"`                     // sender
}

// SetUpdateParams updates xo_tests.notifications struct fields with the specified params.
func (n *Notification) SetUpdateParams(params *NotificationUpdateParams) {
	if params.Body != nil {
		n.Body = *params.Body
	}
	if params.Receiver != nil {
		n.Receiver = *params.Receiver
	}
	if params.Sender != nil {
		n.Sender = *params.Sender
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
	body, receiver, sender
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, n.Body, n.Receiver, n.Sender)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Receiver, n.Sender)
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
	sqlstr := `UPDATE xo_tests.notifications SET 
	body = $1, receiver = $2, sender = $3 
	WHERE notification_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, n.Body, n.Receiver, n.Sender, n.NotificationID)

	rows, err := db.Query(ctx, sqlstr, n.Body, n.Receiver, n.Sender, n.NotificationID)
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
	n.Receiver = params.Receiver
	n.Sender = params.Sender

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
	sqlstr := `DELETE FROM xo_tests.notifications 
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
	notifications.notification_id,
	notifications.receiver,
	notifications.sender %s 
	 FROM xo_tests.notifications %s 
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

// NotificationByNotificationID retrieves a row from 'xo_tests.notifications' as a Notification.
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
	notifications.notification_id,
	notifications.receiver,
	notifications.sender %s 
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
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	n, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	return &n, nil
}

// NotificationsBySender retrieves a row from 'xo_tests.notifications' as a Notification.
//
// Generated from index 'notifications_sender_idx'.
func NotificationsBySender(ctx context.Context, db DB, sender UserID, opts ...NotificationSelectConfigOption) ([]Notification, error) {
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
	notifications.body,
	notifications.notification_id,
	notifications.receiver,
	notifications.sender %s 
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
		return nil, logerror(fmt.Errorf("Notification/NotificationsBySender/Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Notification])
	if err != nil {
		return nil, logerror(fmt.Errorf("Notification/NotificationsBySender/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
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
