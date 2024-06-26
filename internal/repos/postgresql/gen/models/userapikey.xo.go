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

// UserAPIKey represents a row from 'public.user_api_keys'.
type UserAPIKey struct {
	UserAPIKeyID UserAPIKeyID `json:"-" db:"user_api_key_id" nullable:"false"`                    // user_api_key_id
	APIKey       string       `json:"apiKey" db:"api_key" required:"true" nullable:"false"`       // api_key
	ExpiresOn    time.Time    `json:"expiresOn" db:"expires_on" required:"true" nullable:"false"` // expires_on
	UserID       UserID       `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id

	UserJoin *User `json:"-" db:"user_user_id"` // O2O users (inferred)

}

// UserAPIKeyCreateParams represents insert params for 'public.user_api_keys'.
type UserAPIKeyCreateParams struct {
	APIKey    string    `json:"apiKey" required:"true" nullable:"false"`    // api_key
	ExpiresOn time.Time `json:"expiresOn" required:"true" nullable:"false"` // expires_on
	UserID    UserID    `json:"userID" required:"true" nullable:"false"`    // user_id
}

// UserAPIKeyParams represents common params for both insert and update of 'public.user_api_keys'.
type UserAPIKeyParams interface {
	GetAPIKey() *string
	GetExpiresOn() *time.Time
	GetUserID() *UserID
}

func (p UserAPIKeyCreateParams) GetAPIKey() *string {
	x := p.APIKey
	return &x
}
func (p UserAPIKeyUpdateParams) GetAPIKey() *string {
	return p.APIKey
}

func (p UserAPIKeyCreateParams) GetExpiresOn() *time.Time {
	x := p.ExpiresOn
	return &x
}
func (p UserAPIKeyUpdateParams) GetExpiresOn() *time.Time {
	return p.ExpiresOn
}

func (p UserAPIKeyCreateParams) GetUserID() *UserID {
	x := p.UserID
	return &x
}
func (p UserAPIKeyUpdateParams) GetUserID() *UserID {
	return p.UserID
}

type UserAPIKeyID int

// CreateUserAPIKey creates a new UserAPIKey in the database with the given params.
func CreateUserAPIKey(ctx context.Context, db DB, params *UserAPIKeyCreateParams) (*UserAPIKey, error) {
	uak := &UserAPIKey{
		APIKey:    params.APIKey,
		ExpiresOn: params.ExpiresOn,
		UserID:    params.UserID,
	}

	return uak.Insert(ctx, db)
}

type UserAPIKeySelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   UserAPIKeyJoins
	filters map[string][]any
	having  map[string][]any
}
type UserAPIKeySelectConfigOption func(*UserAPIKeySelectConfig)

// WithUserAPIKeyLimit limits row selection.
func WithUserAPIKeyLimit(limit int) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithUserAPIKeyOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithUserAPIKeyOrderBy(rows map[string]*Direction) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		te := EntityFields[TableEntityUserAPIKey]
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

type UserAPIKeyJoins struct {
	User bool `json:"user" required:"true" nullable:"false"` // O2O users
}

// WithUserAPIKeyJoin joins with the given tables.
func WithUserAPIKeyJoin(joins UserAPIKeyJoins) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.joins = UserAPIKeyJoins{
			User: s.joins.User || joins.User,
		}
	}
}

// WithUserAPIKeyFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserAPIKeyFilters(filters map[string][]any) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.filters = filters
	}
}

// WithUserAPIKeyHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithUserAPIKeyHavingClause(conditions map[string][]any) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.having = conditions
	}
}

const userAPIKeyTableUserJoinSQL = `-- O2O join generated from "user_api_keys_user_id_fkey (inferred)"
left join users as _user_api_keys_user_id on _user_api_keys_user_id.user_id = user_api_keys.user_id
`

const userAPIKeyTableUserSelectSQL = `(case when _user_api_keys_user_id.user_id is not null then row(_user_api_keys_user_id.*) end) as user_user_id`

const userAPIKeyTableUserGroupBySQL = `_user_api_keys_user_id.user_id,
      _user_api_keys_user_id.user_id,
	user_api_keys.user_api_key_id`

// UserAPIKeyUpdateParams represents update params for 'public.user_api_keys'.
type UserAPIKeyUpdateParams struct {
	APIKey    *string    `json:"apiKey" nullable:"false"`    // api_key
	ExpiresOn *time.Time `json:"expiresOn" nullable:"false"` // expires_on
	UserID    *UserID    `json:"userID" nullable:"false"`    // user_id
}

// SetUpdateParams updates public.user_api_keys struct fields with the specified params.
func (uak *UserAPIKey) SetUpdateParams(params *UserAPIKeyUpdateParams) {
	if params.APIKey != nil {
		uak.APIKey = *params.APIKey
	}
	if params.ExpiresOn != nil {
		uak.ExpiresOn = *params.ExpiresOn
	}
	if params.UserID != nil {
		uak.UserID = *params.UserID
	}
}

// Insert inserts the UserAPIKey to the database.
func (uak *UserAPIKey) Insert(ctx context.Context, db DB) (*UserAPIKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_api_keys (
	api_key, expires_on, user_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)

	rows, err := db.Query(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Insert/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	*uak = newuak

	return uak, nil
}

// Update updates a UserAPIKey in the database.
func (uak *UserAPIKey) Update(ctx context.Context, db DB) (*UserAPIKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.user_api_keys SET 
	api_key = $1, expires_on = $2, user_id = $3 
	WHERE user_api_key_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID)

	rows, err := db.Query(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Update/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}
	*uak = newuak

	return uak, nil
}

// Upsert upserts a UserAPIKey in the database.
// Requires appropriate PK(s) to be set beforehand.
func (uak *UserAPIKey) Upsert(ctx context.Context, db DB, params *UserAPIKeyCreateParams) (*UserAPIKey, error) {
	var err error

	uak.APIKey = params.APIKey
	uak.ExpiresOn = params.ExpiresOn
	uak.UserID = params.UserID

	uak, err = uak.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUserAPIKey/Insert: %w", &XoError{Entity: "User api key", Err: err})
			}
			uak, err = uak.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUserAPIKey/Update: %w", &XoError{Entity: "User api key", Err: err})
			}
		}
	}

	return uak, err
}

// Delete deletes the UserAPIKey from the database.
func (uak *UserAPIKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_api_keys 
	WHERE user_api_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, uak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserAPIKeyPaginated returns a cursor-paginated list of UserAPIKey.
// At least one cursor is required.
func UserAPIKeyPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...UserAPIKeySelectConfigOption) ([]UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{},
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
	field, ok := EntityFields[TableEntityUserAPIKey][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/cursor: %w", &XoError{Entity: "User api key", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("user_api_keys.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/orderBy: %w", &XoError{Entity: "User api key", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.User {
		selectClauses = append(selectClauses, userAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, userAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM public.user_api_keys %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* UserAPIKeyPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User api key", Err: err}))
	}
	return res, nil
}

// UserAPIKeyByAPIKey retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_api_key_key'.
func UserAPIKeyByAPIKey(ctx context.Context, db DB, apiKey string, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, userAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, userAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM public.user_api_keys %s 
	 WHERE user_api_keys.api_key = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* UserAPIKeyByAPIKey */\n" + sqlstr

	// run
	// logf(sqlstr, apiKey)
	rows, err := db.Query(ctx, sqlstr, append([]any{apiKey}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &uak, nil
}

// UserAPIKeyByUserAPIKeyID retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_pkey'.
func UserAPIKeyByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID UserAPIKeyID, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, userAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, userAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM public.user_api_keys %s 
	 WHERE user_api_keys.user_api_key_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* UserAPIKeyByUserAPIKeyID */\n" + sqlstr

	// run
	// logf(sqlstr, userAPIKeyID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userAPIKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &uak, nil
}

// UserAPIKeyByUserID retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_user_id_key'.
func UserAPIKeyByUserID(ctx context.Context, db DB, userID UserID, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, userAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, userAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, userAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM public.user_api_keys %s 
	 WHERE user_api_keys.user_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* UserAPIKeyByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &uak, nil
}

// FKUser_UserID returns the User associated with the UserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (uak *UserAPIKey) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, uak.UserID)
}
