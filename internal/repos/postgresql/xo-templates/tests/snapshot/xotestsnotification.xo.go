package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsNotification represents a row from 'xo_tests.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsNotification struct {
	NotificationID XoTestsNotificationID `json:"notificationID" db:"notification_id" required:"true" nullable:"false"` // notification_id
	Body           string                `json:"-" db:"body" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                // body
	Sender         XoTestsUserID         `json:"sender" db:"sender" required:"true" nullable:"false"`                  // sender
	Receiver       *XoTestsUserID        `json:"receiver" db:"receiver"`                                               // receiver

	ReceiverJoin *XoTestsUser `json:"-" db:"user_receiver" openapi-go:"ignore"` // O2O users (generated from M2O)
	SenderJoin   *XoTestsUser `json:"-" db:"user_sender" openapi-go:"ignore"`   // O2O users (generated from M2O)
}

// XoTestsNotificationCreateParams represents insert params for 'xo_tests.notifications'.
type XoTestsNotificationCreateParams struct {
	Body     string         `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"` // body
	Receiver *XoTestsUserID `json:"receiver"`                                    // receiver
	Sender   XoTestsUserID  `json:"sender" required:"true" nullable:"false"`     // sender
}

type XoTestsNotificationID int

// CreateXoTestsNotification creates a new XoTestsNotification in the database with the given params.
func CreateXoTestsNotification(ctx context.Context, db DB, params *XoTestsNotificationCreateParams) (*XoTestsNotification, error) {
	xtn := &XoTestsNotification{
		Body:     params.Body,
		Receiver: params.Receiver,
		Sender:   params.Sender,
	}

	return xtn.Insert(ctx, db)
}

// XoTestsNotificationUpdateParams represents update params for 'xo_tests.notifications'.
type XoTestsNotificationUpdateParams struct {
	Body     *string         `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"` // body
	Receiver **XoTestsUserID `json:"receiver"`                                    // receiver
	Sender   *XoTestsUserID  `json:"sender" nullable:"false"`                     // sender
}

// SetUpdateParams updates xo_tests.notifications struct fields with the specified params.
func (xtn *XoTestsNotification) SetUpdateParams(params *XoTestsNotificationUpdateParams) {
	if params.Body != nil {
		xtn.Body = *params.Body
	}
	if params.Receiver != nil {
		xtn.Receiver = *params.Receiver
	}
	if params.Sender != nil {
		xtn.Sender = *params.Sender
	}
}

type XoTestsNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsNotificationJoins
	filters map[string][]any
}
type XoTestsNotificationSelectConfigOption func(*XoTestsNotificationSelectConfig)

// WithXoTestsNotificationLimit limits row selection.
func WithXoTestsNotificationLimit(limit int) XoTestsNotificationSelectConfigOption {
	return func(s *XoTestsNotificationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsNotificationOrderBy string

type XoTestsNotificationJoins struct {
	UserReceiver bool // O2O users
	UserSender   bool // O2O users
}

// WithXoTestsNotificationJoin joins with the given tables.
func WithXoTestsNotificationJoin(joins XoTestsNotificationJoins) XoTestsNotificationSelectConfigOption {
	return func(s *XoTestsNotificationSelectConfig) {
		s.joins = XoTestsNotificationJoins{
			UserReceiver: s.joins.UserReceiver || joins.UserReceiver,
			UserSender:   s.joins.UserSender || joins.UserSender,
		}
	}
}

// WithXoTestsNotificationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsNotificationFilters(filters map[string][]any) XoTestsNotificationSelectConfigOption {
	return func(s *XoTestsNotificationSelectConfig) {
		s.filters = filters
	}
}

const xoTestsNotificationTableUserReceiverJoinSQL = `-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join xo_tests.users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
`

const xoTestsNotificationTableUserReceiverSelectSQL = `(case when _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver`

const xoTestsNotificationTableUserReceiverGroupBySQL = `_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id`

const xoTestsNotificationTableUserSenderJoinSQL = `-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join xo_tests.users as _notifications_sender on _notifications_sender.user_id = notifications.sender
`

const xoTestsNotificationTableUserSenderSelectSQL = `(case when _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender`

const xoTestsNotificationTableUserSenderGroupBySQL = `_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id`

// Insert inserts the XoTestsNotification to the database.
func (xtn *XoTestsNotification) Insert(ctx context.Context, db DB) (*XoTestsNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.notifications (
	body, receiver, sender
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtn.Body, xtn.Receiver, xtn.Sender)

	rows, err := db.Query(ctx, sqlstr, xtn.Body, xtn.Receiver, xtn.Sender)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Insert/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newxtn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	*xtn = newxtn

	return xtn, nil
}

// Update updates a XoTestsNotification in the database.
func (xtn *XoTestsNotification) Update(ctx context.Context, db DB) (*XoTestsNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.notifications SET
	body = $1, receiver = $2, sender = $3
	WHERE notification_id = $4
	RETURNING * `
	// run
	logf(sqlstr, xtn.Body, xtn.Receiver, xtn.Sender, xtn.NotificationID)

	rows, err := db.Query(ctx, sqlstr, xtn.Body, xtn.Receiver, xtn.Sender, xtn.NotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Update/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newxtn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}
	*xtn = newxtn

	return xtn, nil
}

// Upsert upserts a XoTestsNotification in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtn *XoTestsNotification) Upsert(ctx context.Context, db DB, params *XoTestsNotificationCreateParams) (*XoTestsNotification, error) {
	var err error

	xtn.Body = params.Body
	xtn.Receiver = params.Receiver
	xtn.Sender = params.Sender

	xtn, err = xtn.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Notification", Err: err})
			}
			xtn, err = xtn.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Notification", Err: err})
			}
		}
	}

	return xtn, err
}

// Delete deletes the XoTestsNotification from the database.
func (xtn *XoTestsNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.notifications
	WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtn.NotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsNotificationPaginatedByNotificationID returns a cursor-paginated list of XoTestsNotification.
func XoTestsNotificationPaginatedByNotificationID(ctx context.Context, db DB, notificationID XoTestsNotificationID, direction models.Direction, opts ...XoTestsNotificationSelectConfigOption) ([]XoTestsNotification, error) {
	c := &XoTestsNotificationSelectConfig{joins: XoTestsNotificationJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, xoTestsNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, xoTestsNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserSenderGroupBySQL)
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
	if direction == models.DirectionAsc {
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
	sqlstr = "/* XoTestsNotificationPaginatedByNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Paginated/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
	}
	return res, nil
}

// XoTestsNotificationByNotificationID retrieves a row from 'xo_tests.notifications' as a XoTestsNotification.
//
// Generated from index 'notifications_pkey'.
func XoTestsNotificationByNotificationID(ctx context.Context, db DB, notificationID XoTestsNotificationID, opts ...XoTestsNotificationSelectConfigOption) (*XoTestsNotification, error) {
	c := &XoTestsNotificationSelectConfig{joins: XoTestsNotificationJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, xoTestsNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, xoTestsNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserSenderGroupBySQL)
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
	sqlstr = "/* XoTestsNotificationByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	xtn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	return &xtn, nil
}

// XoTestsNotificationsBySender retrieves a row from 'xo_tests.notifications' as a XoTestsNotification.
//
// Generated from index 'notifications_sender_idx'.
func XoTestsNotificationsBySender(ctx context.Context, db DB, sender XoTestsUserID, opts ...XoTestsNotificationSelectConfigOption) ([]XoTestsNotification, error) {
	c := &XoTestsNotificationSelectConfig{joins: XoTestsNotificationJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, xoTestsNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, xoTestsNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsNotificationTableUserSenderGroupBySQL)
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
	sqlstr = "/* XoTestsNotificationsBySender */\n" + sqlstr

	// run
	// logf(sqlstr, sender)
	rows, err := db.Query(ctx, sqlstr, append([]any{sender}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/NotificationsBySender/Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsNotification/NotificationsBySender/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
	}
	return res, nil
}

// FKUser_Receiver returns the User associated with the XoTestsNotification's (Receiver).
//
// Generated from foreign key 'notifications_receiver_fkey'.
func (xtn *XoTestsNotification) FKUser_Receiver(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, *xtn.Receiver)
}

// FKUser_Sender returns the User associated with the XoTestsNotification's (Sender).
//
// Generated from foreign key 'notifications_sender_fkey'.
func (xtn *XoTestsNotification) FKUser_Sender(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtn.Sender)
}
