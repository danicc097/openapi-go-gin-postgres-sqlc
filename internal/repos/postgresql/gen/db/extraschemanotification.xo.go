

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

// ExtraSchemaNotification represents a row from 'extra_schema.notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaNotification struct {
	NotificationID   ExtraSchemaNotificationID `json:"notificationID" db:"notification_id" required:"true" nullable:"false"`     // notification_id
	Body             string                    `json:"-" db:"body" nullable:"false" pattern:"^[A-Za-z0-9]*$"`                    // body
	Sender           UserID                    `json:"sender" db:"sender" required:"true" nullable:"false"`                      // sender
	Receiver         *UserID                   `json:"receiver" db:"receiver"`                                                   // receiver
	NotificationType NotificationType          `json:"notificationType" db:"notification_type" required:"true" nullable:"false"` // notification_type

	ReceiverJoin *User `json:"-" db:"user_receiver" openapi-go:"ignore"` // O2O users (generated from M2O)
	SenderJoin   *User `json:"-" db:"user_sender" openapi-go:"ignore"`   // O2O users (generated from M2O)

}

// ExtraSchemaNotificationCreateParams represents insert params for 'extra_schema.notifications'.
type ExtraSchemaNotificationCreateParams struct {
	Body             string           `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"`       // body
	NotificationType NotificationType `json:"notificationType" required:"true" nullable:"false"` // notification_type
	Receiver         *UserID          `json:"receiver"`                                          // receiver
	Sender           UserID           `json:"sender" required:"true" nullable:"false"`           // sender
}

type ExtraSchemaNotificationID int

// CreateExtraSchemaNotification creates a new ExtraSchemaNotification in the database with the given params.
func CreateExtraSchemaNotification(ctx context.Context, db DB, params *ExtraSchemaNotificationCreateParams) (*ExtraSchemaNotification, error) {
	esn := &ExtraSchemaNotification{
		Body:             params.Body,
		NotificationType: params.NotificationType,
		Receiver:         params.Receiver,
		Sender:           params.Sender,
	}

	return esn.Insert(ctx, db)
}

// ExtraSchemaNotificationUpdateParams represents update params for 'extra_schema.notifications'.
type ExtraSchemaNotificationUpdateParams struct {
	Body             *string           `json:"-" nullable:"false" pattern:"^[A-Za-z0-9]*$"` // body
	NotificationType *NotificationType `json:"notificationType" nullable:"false"`           // notification_type
	Receiver         **UserID          `json:"receiver"`                                    // receiver
	Sender           *UserID           `json:"sender" nullable:"false"`                     // sender
}

// SetUpdateParams updates extra_schema.notifications struct fields with the specified params.
func (esn *ExtraSchemaNotification) SetUpdateParams(params *ExtraSchemaNotificationUpdateParams) {
	if params.Body != nil {
		esn.Body = *params.Body
	}
	if params.NotificationType != nil {
		esn.NotificationType = *params.NotificationType
	}
	if params.Receiver != nil {
		esn.Receiver = *params.Receiver
	}
	if params.Sender != nil {
		esn.Sender = *params.Sender
	}
}

type ExtraSchemaNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaNotificationJoins
	filters map[string][]any
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

type ExtraSchemaNotificationOrderBy string

const ()

type ExtraSchemaNotificationJoins struct {
	UserReceiver bool // O2O users
	UserSender   bool // O2O users
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

// WithExtraSchemaNotificationFilters adds the given filters, which can be dynamically parameterized
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

// Insert inserts the ExtraSchemaNotification to the database.
func (esn *ExtraSchemaNotification) Insert(ctx context.Context, db DB) (*ExtraSchemaNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.notifications (
	body, notification_type, receiver, sender
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, esn.Body, esn.NotificationType, esn.Receiver, esn.Sender)

	rows, err := db.Query(ctx, sqlstr, esn.Body, esn.NotificationType, esn.Receiver, esn.Sender)
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
	logf(sqlstr, esn.Body, esn.NotificationType, esn.Receiver, esn.Sender, esn.NotificationID)

	rows, err := db.Query(ctx, sqlstr, esn.Body, esn.NotificationType, esn.Receiver, esn.Sender, esn.NotificationID)
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
	esn.NotificationType = params.NotificationType
	esn.Receiver = params.Receiver
	esn.Sender = params.Sender

	esn, err = esn.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Notification", Err: err})
			}
			esn, err = esn.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Notification", Err: err})
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

// ExtraSchemaNotificationPaginatedByNotificationID returns a cursor-paginated list of ExtraSchemaNotification.
func ExtraSchemaNotificationPaginatedByNotificationID(ctx context.Context, db DB, notificationID ExtraSchemaNotificationID, direction models.Direction, opts ...ExtraSchemaNotificationSelectConfigOption) ([]ExtraSchemaNotification, error) {
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{}, filters: make(map[string][]any)}

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
	notifications.notification_type,
	notifications.receiver,
	notifications.sender %s 
	 FROM extra_schema.notifications %s 
	 WHERE notifications.notification_id %s $1
	 %s   %s 
  ORDER BY 
		notification_id %s `, selects, joins, operator, filters, groupbys, direction)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationPaginatedByNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
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
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{}, filters: make(map[string][]any)}

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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, filterParams...)...)
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
func ExtraSchemaNotificationsBySender(ctx context.Context, db DB, sender UserID, opts ...ExtraSchemaNotificationSelectConfigOption) ([]ExtraSchemaNotification, error) {
	c := &ExtraSchemaNotificationSelectConfig{joins: ExtraSchemaNotificationJoins{}, filters: make(map[string][]any)}

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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaNotificationsBySender */\n" + sqlstr

	// run
	// logf(sqlstr, sender)
	rows, err := db.Query(ctx, sqlstr, append([]any{sender}, filterParams...)...)
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
func (esn *ExtraSchemaNotification) FKUser_Receiver(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, *esn.Receiver)
}

// FKUser_Sender returns the User associated with the ExtraSchemaNotification's (Sender).
//
// Generated from foreign key 'notifications_sender_fkey'.
func (esn *ExtraSchemaNotification) FKUser_Sender(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, esn.Sender)
}
