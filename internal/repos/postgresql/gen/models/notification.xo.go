// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

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

	UserReceiverJoin      *User               `json:"-" db:"user_receiver"`      // O2O users (generated from M2O)
	UserSenderJoin        *User               `json:"-" db:"user_sender"`        // O2O users (generated from M2O)
	UserNotificationsJoin *[]UserNotification `json:"-" db:"user_notifications"` // M2O notifications

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

// NotificationParams represents common params for both insert and update of 'public.notifications'.
type NotificationParams interface {
	GetBody() *string
	GetLabels() *[]string
	GetLink() *string
	GetNotificationType() *NotificationType
	GetReceiver() *UserID
	GetReceiverRank() *int
	GetSender() *UserID
	GetTitle() *string
}

func (p NotificationCreateParams) GetBody() *string {
	x := p.Body
	return &x
}
func (p NotificationUpdateParams) GetBody() *string {
	return p.Body
}

func (p NotificationCreateParams) GetLabels() *[]string {
	x := p.Labels
	return &x
}
func (p NotificationUpdateParams) GetLabels() *[]string {
	return p.Labels
}

func (p NotificationCreateParams) GetLink() *string {
	return p.Link
}
func (p NotificationUpdateParams) GetLink() *string {
	if p.Link != nil {
		return *p.Link
	}
	return nil
}

func (p NotificationCreateParams) GetNotificationType() *NotificationType {
	x := p.NotificationType
	return &x
}
func (p NotificationUpdateParams) GetNotificationType() *NotificationType {
	return p.NotificationType
}

func (p NotificationCreateParams) GetReceiver() *UserID {
	return p.Receiver
}
func (p NotificationUpdateParams) GetReceiver() *UserID {
	if p.Receiver != nil {
		return *p.Receiver
	}
	return nil
}

func (p NotificationCreateParams) GetReceiverRank() *int {
	return p.ReceiverRank
}
func (p NotificationUpdateParams) GetReceiverRank() *int {
	if p.ReceiverRank != nil {
		return *p.ReceiverRank
	}
	return nil
}

func (p NotificationCreateParams) GetSender() *UserID {
	x := p.Sender
	return &x
}
func (p NotificationUpdateParams) GetSender() *UserID {
	return p.Sender
}

func (p NotificationCreateParams) GetTitle() *string {
	x := p.Title
	return &x
}
func (p NotificationUpdateParams) GetTitle() *string {
	return p.Title
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

type NotificationSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   NotificationJoins
	filters map[string][]any
	having  map[string][]any
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

// WithNotificationOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithNotificationOrderBy(rows map[string]*Direction) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		te := EntityFields[TableEntityNotification]
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

type NotificationJoins struct {
	UserReceiver      bool `json:"userReceiver" required:"true" nullable:"false"`      // O2O users
	UserSender        bool `json:"userSender" required:"true" nullable:"false"`        // O2O users
	UserNotifications bool `json:"userNotifications" required:"true" nullable:"false"` // M2O user_notifications
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

// WithNotificationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
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

// WithNotificationHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithNotificationHavingClause(conditions map[string][]any) NotificationSelectConfigOption {
	return func(s *NotificationSelectConfig) {
		s.having = conditions
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
    , row(user_notifications.*) as __user_notifications
  from
    user_notifications
  group by
	  user_notifications_notification_id, user_notifications.user_notification_id
) as xo_join_user_notifications on xo_join_user_notifications.user_notifications_notification_id = notifications.notification_id
`

const notificationTableUserNotificationsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_user_notifications.__user_notifications)) filter (where xo_join_user_notifications.user_notifications_notification_id is not null), '{}') as user_notifications`

const notificationTableUserNotificationsGroupBySQL = `notifications.notification_id`

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
	logf(sqlstr, n.Body, n.Labels, n.Link, n.NotificationType, n.Receiver, n.ReceiverRank, n.Sender, n.Title, n.NotificationID)

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
				return nil, fmt.Errorf("UpsertNotification/Insert: %w", &XoError{Entity: "Notification", Err: err})
			}
			n, err = n.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertNotification/Update: %w", &XoError{Entity: "Notification", Err: err})
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

// NotificationPaginated returns a cursor-paginated list of Notification.
// At least one cursor is required.
func NotificationPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...NotificationSelectConfigOption) ([]Notification, error) {
	c := &NotificationSelectConfig{joins: NotificationJoins{},
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
	field, ok := EntityFields[TableEntityNotification][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("Notification/Paginated/cursor: %w", &XoError{Entity: "Notification", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
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
		return nil, logerror(fmt.Errorf("Notification/Paginated/orderBy: %w", &XoError{Entity: "Notification", Err: fmt.Errorf("at least one sorted column is required")}))
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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* NotificationPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
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
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* NotificationByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, append(filterParams, havingParams...)...)...)
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
	c := &NotificationSelectConfig{joins: NotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* NotificationsByReceiverRankNotificationTypeCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, receiverRank, notificationType, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{receiverRank, notificationType, createdAt}, append(filterParams, havingParams...)...)...)
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
