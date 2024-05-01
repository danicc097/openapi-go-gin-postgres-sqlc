// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

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

// ExtraSchemaNotification represents a row from 'extra_schema.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaNotification struct {
	NotificationID              ExtraSchemaNotificationID   `json:"notificationID" db:"notification_id" required:"true" nullable:"false"`                                                 // notification_id
	Body                        string                      `json:"-" db:"body" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                                                                // body
	Sender                      ExtraSchemaUserID           `json:"sender" db:"sender" required:"true" nullable:"false"`                                                                  // sender
	Receiver                    *ExtraSchemaUserID          `json:"receiver" db:"receiver"`                                                                                               // receiver
	ExtraSchemaNotificationType ExtraSchemaNotificationType `json:"notificationType" db:"notification_type" required:"true" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type

	UserReceiverJoin *ExtraSchemaUser `json:"-" db:"user_receiver"` // O2O users (generated from M2O)
	UserSenderJoin   *ExtraSchemaUser `json:"-" db:"user_sender"`   // O2O users (generated from M2O)

}

// ExtraSchemaNotificationCreateParams represents insert params for 'extra_schema.notifications'.
type ExtraSchemaNotificationCreateParams struct {
	Body                        string                      `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                                                   // body
	ExtraSchemaNotificationType ExtraSchemaNotificationType `json:"notificationType" required:"true" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type
	Receiver                    *ExtraSchemaUserID          `json:"receiver"`                                                                                      // receiver
	Sender                      ExtraSchemaUserID           `json:"sender" required:"true" nullable:"false"`                                                       // sender
}

// ExtraSchemaNotificationParams represents common params for both insert and update of 'extra_schema.notifications'.
type ExtraSchemaNotificationParams interface {
	GetBody() *string
	GetExtraSchemaNotificationType() *ExtraSchemaNotificationType
	GetReceiver() *ExtraSchemaUserID
	GetSender() *ExtraSchemaUserID
}

func (p ExtraSchemaNotificationCreateParams) GetBody() *string {
	x := p.Body
	return &x
}
func (p ExtraSchemaNotificationUpdateParams) GetBody() *string {
	return p.Body
}

func (p ExtraSchemaNotificationCreateParams) GetExtraSchemaNotificationType() *ExtraSchemaNotificationType {
	x := p.ExtraSchemaNotificationType
	return &x
}
func (p ExtraSchemaNotificationUpdateParams) GetExtraSchemaNotificationType() *ExtraSchemaNotificationType {
	return p.ExtraSchemaNotificationType
}

func (p ExtraSchemaNotificationCreateParams) GetReceiver() *ExtraSchemaUserID {
	return p.Receiver
}
func (p ExtraSchemaNotificationUpdateParams) GetReceiver() *ExtraSchemaUserID {
	if p.Receiver != nil {
		return *p.Receiver
	}
	return nil
}

func (p ExtraSchemaNotificationCreateParams) GetSender() *ExtraSchemaUserID {
	x := p.Sender
	return &x
}
func (p ExtraSchemaNotificationUpdateParams) GetSender() *ExtraSchemaUserID {
	return p.Sender
}

type ExtraSchemaNotificationID int

// CreateExtraSchemaNotification creates a new ExtraSchemaNotification in the database with the given params.
func CreateExtraSchemaNotification(ctx context.Context, db DB, params *ExtraSchemaNotificationCreateParams) (*ExtraSchemaNotification, error) {
	esn := &ExtraSchemaNotification{
		Body:                        params.Body,
		ExtraSchemaNotificationType: params.ExtraSchemaNotificationType,
		Receiver:                    params.Receiver,
		Sender:                      params.Sender,
	}

	return esn.Insert(ctx, db)
}

type ExtraSchemaNotificationSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   ExtraSchemaNotificationJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaNotificationSelectConfigOption func(*ExtraSchemaNotificationSelectConfig)

// WithExtraSchemaNotificationLimit limits row selection.
func WithExtraSchemaNotificationLimit(limit int) ExtraSchemaNotificationSelectConfigOption {
	return func(s *ExtraSchemaNotificationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaNotificationOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaNotificationOrderBy(rows map[string]*Direction) ExtraSchemaNotificationSelectConfigOption {
	return func(s *ExtraSchemaNotificationSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaNotification]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type ExtraSchemaNotificationJoins struct {
	UserReceiver bool `json:"userReceiver" required:"true" nullable:"false"` // O2O users
	UserSender   bool `json:"userSender" required:"true" nullable:"false"`   // O2O users
}

// WithExtraSchemaNotificationJoin joins with the given tables.
func WithExtraSchemaNotificationJoin(joins ExtraSchemaNotificationJoins) ExtraSchemaNotificationSelectConfigOption {
	return func(s *ExtraSchemaNotificationSelectConfig) {
		s.joins = ExtraSchemaNotificationJoins{
			UserReceiver: s.joins.UserReceiver || joins.UserReceiver,
			UserSender:   s.joins.UserSender || joins.UserSender,
		}
	}
}

// WithExtraSchemaNotificationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaNotificationFilters(filters map[string][]any) ExtraSchemaNotificationSelectConfigOption {
	return func(s *ExtraSchemaNotificationSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaNotificationHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithExtraSchemaNotificationHavingClause(conditions map[string][]any) ExtraSchemaNotificationSelectConfigOption {
	return func(s *ExtraSchemaNotificationSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaNotificationTableUserReceiverJoinSQL = `-- O2O join generated from "notifications_receiver_fkey (Generated from M2O)"
left join extra_schema.users as _notifications_receiver on _notifications_receiver.user_id = notifications.receiver
`

const extraSchemaNotificationTableUserReceiverSelectSQL = `(case when _notifications_receiver.user_id is not null then row(_notifications_receiver.*) end) as user_receiver`

const extraSchemaNotificationTableUserReceiverGroupBySQL = `_notifications_receiver.user_id,
      _notifications_receiver.user_id,
	notifications.notification_id`

const extraSchemaNotificationTableUserSenderJoinSQL = `-- O2O join generated from "notifications_sender_fkey (Generated from M2O)"
left join extra_schema.users as _notifications_sender on _notifications_sender.user_id = notifications.sender
`

const extraSchemaNotificationTableUserSenderSelectSQL = `(case when _notifications_sender.user_id is not null then row(_notifications_sender.*) end) as user_sender`

const extraSchemaNotificationTableUserSenderGroupBySQL = `_notifications_sender.user_id,
      _notifications_sender.user_id,
	notifications.notification_id`

// ExtraSchemaNotificationUpdateParams represents update params for 'extra_schema.notifications'.
type ExtraSchemaNotificationUpdateParams struct {
	Body                        *string                      `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                                   // body
	ExtraSchemaNotificationType *ExtraSchemaNotificationType `json:"notificationType" nullable:"false" ref:"#/components/schemas/NotificationType"` // notification_type
	Receiver                    **ExtraSchemaUserID          `json:"receiver"`                                                                      // receiver
	Sender                      *ExtraSchemaUserID           `json:"sender" nullable:"false"`                                                       // sender
}

// SetUpdateParams updates extra_schema.notifications struct fields with the specified params.
func (esn *ExtraSchemaNotification) SetUpdateParams(params *ExtraSchemaNotificationUpdateParams) {
	if params.Body != nil {
		esn.Body = *params.Body
	}
	if params.ExtraSchemaNotificationType != nil {
		esn.ExtraSchemaNotificationType = *params.ExtraSchemaNotificationType
	}
	if params.Receiver != nil {
		esn.Receiver = *params.Receiver
	}
	if params.Sender != nil {
		esn.Sender = *params.Sender
	}
}

// Insert inserts the ExtraSchemaNotification to the database.
func (esn *ExtraSchemaNotification) Insert(ctx context.Context, db DB) (*ExtraSchemaNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.notifications (
	body, notification_type, receiver, sender
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, esn.Body, esn.ExtraSchemaNotificationType, esn.Receiver, esn.Sender)

	rows, err := db.Query(ctx, sqlstr, esn.Body, esn.ExtraSchemaNotificationType, esn.Receiver, esn.Sender)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Insert/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newesn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	*esn = newesn

	return esn, nil
}

// Update updates a ExtraSchemaNotification in the database.
func (esn *ExtraSchemaNotification) Update(ctx context.Context, db DB) (*ExtraSchemaNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.notifications SET 
	body = $1, notification_type = $2, receiver = $3, sender = $4 
	WHERE notification_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, esn.Body, esn.ExtraSchemaNotificationType, esn.Receiver, esn.Sender, esn.NotificationID)

	rows, err := db.Query(ctx, sqlstr, esn.Body, esn.ExtraSchemaNotificationType, esn.Receiver, esn.Sender, esn.NotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Update/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	newesn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}
	*esn = newesn

	return esn, nil
}

// Upsert upserts a ExtraSchemaNotification in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esn *ExtraSchemaNotification) Upsert(ctx context.Context, db DB, params *ExtraSchemaNotificationCreateParams) (*ExtraSchemaNotification, error) {
	var err error

	esn.Body = params.Body
	esn.ExtraSchemaNotificationType = params.ExtraSchemaNotificationType
	esn.Receiver = params.Receiver
	esn.Sender = params.Sender

	esn, err = esn.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaNotification/Insert: %w", &XoError{Entity: "Notification", Err: err})
			}
			esn, err = esn.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaNotification/Update: %w", &XoError{Entity: "Notification", Err: err})
			}
		}
	}

	return esn, err
}

// Delete deletes the ExtraSchemaNotification from the database.
func (esn *ExtraSchemaNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.notifications 
	WHERE notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esn.NotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaNotificationPaginated returns a cursor-paginated list of ExtraSchemaNotification.
// At least one cursor is required.
func ExtraSchemaNotificationPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaNotificationSelectConfigOption) ([]ExtraSchemaNotification, error) {
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {

		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaNotification][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Paginated/cursor: %w", &XoError{Entity: "Notification", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("notifications.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

	paramStart := 0 // all filters will come from the user
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
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Paginated/orderBy: %w", &XoError{Entity: "Notification", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserSenderGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.sender %s 
	 FROM extra_schema.notifications %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Paginated/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
	}
	return res, nil
}

// ExtraSchemaNotificationByNotificationID retrieves a row from 'extra_schema.notifications' as a ExtraSchemaNotification.
//
// Generated from index 'notifications_pkey'.
func ExtraSchemaNotificationByNotificationID(ctx context.Context, db DB, notificationID ExtraSchemaNotificationID, opts ...ExtraSchemaNotificationSelectConfigOption) (*ExtraSchemaNotification, error) {
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserSenderGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.sender %s 
	 FROM extra_schema.notifications %s 
	 WHERE notifications.notification_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/db.Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	esn, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("notifications/NotificationByNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "Notification", Err: err}))
	}

	return &esn, nil
}

// ExtraSchemaNotificationsBySender retrieves a row from 'extra_schema.notifications' as a ExtraSchemaNotification.
//
// Generated from index 'notifications_sender_idx'.
func ExtraSchemaNotificationsBySender(ctx context.Context, db DB, sender ExtraSchemaUserID, opts ...ExtraSchemaNotificationSelectConfigOption) ([]ExtraSchemaNotification, error) {
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.UserReceiver {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserReceiverGroupBySQL)
	}

	if c.joins.UserSender {
		selectClauses = append(selectClauses, extraSchemaNotificationTableUserSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaNotificationTableUserSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaNotificationTableUserSenderGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	notifications.body,
	notifications.notification_id,
	notifications.notification_type,
	notifications.receiver,
	notifications.sender %s 
	 FROM extra_schema.notifications %s 
	 WHERE notifications.sender = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationsBySender */\n" + sqlstr

	// run
	// logf(sqlstr, sender)
	rows, err := db.Query(ctx, sqlstr, append([]any{sender}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/NotificationsBySender/Query: %w", &XoError{Entity: "Notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaNotification/NotificationsBySender/pgx.CollectRows: %w", &XoError{Entity: "Notification", Err: err}))
	}
	return res, nil
}

// FKUser_Receiver returns the User associated with the ExtraSchemaNotification's (Receiver).
//
// Generated from foreign key 'notifications_receiver_fkey'.
func (esn *ExtraSchemaNotification) FKUser_Receiver(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, *esn.Receiver)
}

// FKUser_Sender returns the User associated with the ExtraSchemaNotification's (Sender).
//
// Generated from foreign key 'notifications_sender_fkey'.
func (esn *ExtraSchemaNotification) FKUser_Sender(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, esn.Sender)
}
