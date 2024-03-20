package db

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

// UserNotification represents a row from 'public.user_notifications'.
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
type UserNotification struct {
	UserNotificationID UserNotificationID `json:"userNotificationID" db:"user_notification_id" required:"true" nullable:"false"` // user_notification_id
	NotificationID     NotificationID     `json:"notificationID" db:"notification_id" required:"true" nullable:"false"`          // notification_id
	Read               bool               `json:"read" db:"read" required:"true" nullable:"false"`                               // read
	UserID             UserID             `json:"userID" db:"user_id" required:"true" nullable:"false"`                          // user_id

	NotificationJoin *Notification `json:"-" db:"notification_notification_id" openapi-go:"ignore"` // O2O notifications (generated from M2O)
	UserJoin         *User         `json:"-" db:"user_user_id" openapi-go:"ignore"`                 // O2O users (generated from M2O)

}

// UserNotificationCreateParams represents insert params for 'public.user_notifications'.
type UserNotificationCreateParams struct {
	NotificationID NotificationID `json:"notificationID" required:"true" nullable:"false"` // notification_id
	Read           bool           `json:"read" required:"true" nullable:"false"`           // read
	UserID         UserID         `json:"userID" required:"true" nullable:"false"`         // user_id
}

// UserNotificationParams represents common params for both insert and update of 'public.user_notifications'.
type UserNotificationParams interface {
	GetNotificationID() *NotificationID
	GetRead() *bool
	GetUserID() *UserID
}

func (p UserNotificationCreateParams) GetNotificationID() *NotificationID {
	x := p.NotificationID
	return &x
}
func (p UserNotificationUpdateParams) GetNotificationID() *NotificationID {
	return p.NotificationID
}

func (p UserNotificationCreateParams) GetRead() *bool {
	x := p.Read
	return &x
}
func (p UserNotificationUpdateParams) GetRead() *bool {
	return p.Read
}

func (p UserNotificationCreateParams) GetUserID() *UserID {
	x := p.UserID
	return &x
}
func (p UserNotificationUpdateParams) GetUserID() *UserID {
	return p.UserID
}

type UserNotificationID int

// CreateUserNotification creates a new UserNotification in the database with the given params.
func CreateUserNotification(ctx context.Context, db DB, params *UserNotificationCreateParams) (*UserNotification, error) {
	un := &UserNotification{
		NotificationID: params.NotificationID,
		Read:           params.Read,
		UserID:         params.UserID,
	}

	return un.Insert(ctx, db)
}

type UserNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   UserNotificationJoins
	filters map[string][]any
	having  map[string][]any
}
type UserNotificationSelectConfigOption func(*UserNotificationSelectConfig)

// WithUserNotificationLimit limits row selection.
func WithUserNotificationLimit(limit int) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserNotificationOrderBy string

const ()

type UserNotificationJoins struct {
	Notification bool `json:"notification" required:"true" nullable:"false"` // O2O notifications
	User         bool `json:"user" required:"true" nullable:"false"`         // O2O users
}

// WithUserNotificationJoin joins with the given tables.
func WithUserNotificationJoin(joins UserNotificationJoins) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.joins = UserNotificationJoins{
			Notification: s.joins.Notification || joins.Notification,
			User:         s.joins.User || joins.User,
		}
	}
}

// WithUserNotificationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserNotificationFilters(filters map[string][]any) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.filters = filters
	}
}

// WithUserNotificationHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithUserNotificationHavingClause(conditions map[string][]any) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.having = conditions
	}
}

const userNotificationTableNotificationJoinSQL = `-- O2O join generated from "user_notifications_notification_id_fkey (Generated from M2O)"
left join notifications as _user_notifications_notification_id on _user_notifications_notification_id.notification_id = user_notifications.notification_id
`

const userNotificationTableNotificationSelectSQL = `(case when _user_notifications_notification_id.notification_id is not null then row(_user_notifications_notification_id.*) end) as notification_notification_id`

const userNotificationTableNotificationGroupBySQL = `_user_notifications_notification_id.notification_id,
      _user_notifications_notification_id.notification_id,
	user_notifications.user_notification_id`

const userNotificationTableUserJoinSQL = `-- O2O join generated from "user_notifications_user_id_fkey (Generated from M2O)"
left join users as _user_notifications_user_id on _user_notifications_user_id.user_id = user_notifications.user_id
`

const userNotificationTableUserSelectSQL = `(case when _user_notifications_user_id.user_id is not null then row(_user_notifications_user_id.*) end) as user_user_id`

const userNotificationTableUserGroupBySQL = `_user_notifications_user_id.user_id,
      _user_notifications_user_id.user_id,
	user_notifications.user_notification_id`

// UserNotificationUpdateParams represents update params for 'public.user_notifications'.
type UserNotificationUpdateParams struct {
	NotificationID *NotificationID `json:"notificationID" nullable:"false"` // notification_id
	Read           *bool           `json:"read" nullable:"false"`           // read
	UserID         *UserID         `json:"userID" nullable:"false"`         // user_id
}

// SetUpdateParams updates public.user_notifications struct fields with the specified params.
func (un *UserNotification) SetUpdateParams(params *UserNotificationUpdateParams) {
	if params.NotificationID != nil {
		un.NotificationID = *params.NotificationID
	}
	if params.Read != nil {
		un.Read = *params.Read
	}
	if params.UserID != nil {
		un.UserID = *params.UserID
	}
}

// Insert inserts the UserNotification to the database.
func (un *UserNotification) Insert(ctx context.Context, db DB) (*UserNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_notifications (
	notification_id, read, user_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User notification", Err: err}))
	}

	*un = newun

	return un, nil
}

// Update updates a UserNotification in the database.
func (un *UserNotification) Update(ctx context.Context, db DB) (*UserNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.user_notifications SET 
	notification_id = $1, read = $2, user_id = $3 
	WHERE user_notification_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User notification", Err: err}))
	}
	*un = newun

	return un, nil
}

// Upsert upserts a UserNotification in the database.
// Requires appropriate PK(s) to be set beforehand.
func (un *UserNotification) Upsert(ctx context.Context, db DB, params *UserNotificationCreateParams) (*UserNotification, error) {
	var err error

	un.NotificationID = params.NotificationID
	un.Read = params.Read
	un.UserID = params.UserID

	un, err = un.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User notification", Err: err})
			}
			un, err = un.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User notification", Err: err})
			}
		}
	}

	return un, err
}

// Delete deletes the UserNotification from the database.
func (un *UserNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_notifications 
	WHERE user_notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserNotificationPaginatedByUserNotificationID returns a cursor-paginated list of UserNotification.
func UserNotificationPaginatedByUserNotificationID(ctx context.Context, db DB, userNotificationID UserNotificationID, direction models.Direction, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.user_notification_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		user_notification_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* UserNotificationPaginatedByUserNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{userNotificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User notification", Err: err}))
	}
	return res, nil
}

// UserNotificationPaginatedByNotificationID returns a cursor-paginated list of UserNotification.
func UserNotificationPaginatedByNotificationID(ctx context.Context, db DB, notificationID NotificationID, direction models.Direction, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.notification_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		notification_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* UserNotificationPaginatedByNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User notification", Err: err}))
	}
	return res, nil
}

// UserNotificationByNotificationIDUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationByNotificationIDUserID(ctx context.Context, db DB, notificationID NotificationID, userID UserID, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.notification_id = $1 AND user_notifications.user_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserNotificationByNotificationIDUserID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID, userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User notification", Err: err}))
	}

	return &un, nil
}

// UserNotificationsByNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationsByNotificationID(ctx context.Context, db DB, notificationID NotificationID, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.notification_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserNotificationsByNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{notificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationByNotificationIDUserID/Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationByNotificationIDUserID/pgx.CollectRows: %w", &XoError{Entity: "User notification", Err: err}))
	}
	return res, nil
}

// UserNotificationByUserNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_pkey'.
func UserNotificationByUserNotificationID(ctx context.Context, db DB, userNotificationID UserNotificationID, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.user_notification_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserNotificationByUserNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, userNotificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userNotificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/db.Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "User notification", Err: err}))
	}

	return &un, nil
}

// UserNotificationsByUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_user_id_idx'.
func UserNotificationsByUserID(ctx context.Context, db DB, userID UserID, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Notification {
		selectClauses = append(selectClauses, userNotificationTableNotificationSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableNotificationJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableNotificationGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, userNotificationTableUserSelectSQL)
		joinClauses = append(joinClauses, userNotificationTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userNotificationTableUserGroupBySQL)
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
	user_notifications.notification_id,
	user_notifications.read,
	user_notifications.user_id,
	user_notifications.user_notification_id %s 
	 FROM public.user_notifications %s 
	 WHERE user_notifications.user_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserNotificationsByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationsByUserID/Query: %w", &XoError{Entity: "User notification", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/UserNotificationsByUserID/pgx.CollectRows: %w", &XoError{Entity: "User notification", Err: err}))
	}
	return res, nil
}

// FKNotification_NotificationID returns the Notification associated with the UserNotification's (NotificationID).
//
// Generated from foreign key 'user_notifications_notification_id_fkey'.
func (un *UserNotification) FKNotification_NotificationID(ctx context.Context, db DB) (*Notification, error) {
	return NotificationByNotificationID(ctx, db, un.NotificationID)
}

// FKUser_UserID returns the User associated with the UserNotification's (UserID).
//
// Generated from foreign key 'user_notifications_user_id_fkey'.
func (un *UserNotification) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, un.UserID)
}
