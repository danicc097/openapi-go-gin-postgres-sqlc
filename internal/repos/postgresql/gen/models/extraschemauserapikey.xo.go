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

// ExtraSchemaUserAPIKey represents a row from 'extra_schema.user_api_keys'.
type ExtraSchemaUserAPIKey struct {
	UserAPIKeyID ExtraSchemaUserAPIKeyID `json:"-" db:"user_api_key_id" nullable:"false"`                    // user_api_key_id
	APIKey       string                  `json:"apiKey" db:"api_key" required:"true" nullable:"false"`       // api_key
	ExpiresOn    time.Time               `json:"expiresOn" db:"expires_on" required:"true" nullable:"false"` // expires_on
	UserID       ExtraSchemaUserID       `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id

	UserJoin *ExtraSchemaUser `json:"-" db:"user_user_id"` // O2O users (inferred)

}

// ExtraSchemaUserAPIKeyCreateParams represents insert params for 'extra_schema.user_api_keys'.
type ExtraSchemaUserAPIKeyCreateParams struct {
	APIKey    string            `json:"apiKey" required:"true" nullable:"false"`    // api_key
	ExpiresOn time.Time         `json:"expiresOn" required:"true" nullable:"false"` // expires_on
	UserID    ExtraSchemaUserID `json:"userID" required:"true" nullable:"false"`    // user_id
}

// ExtraSchemaUserAPIKeyParams represents common params for both insert and update of 'extra_schema.user_api_keys'.
type ExtraSchemaUserAPIKeyParams interface {
	GetAPIKey() *string
	GetExpiresOn() *time.Time
	GetUserID() *ExtraSchemaUserID
}

func (p ExtraSchemaUserAPIKeyCreateParams) GetAPIKey() *string {
	x := p.APIKey
	return &x
}
func (p ExtraSchemaUserAPIKeyUpdateParams) GetAPIKey() *string {
	return p.APIKey
}

func (p ExtraSchemaUserAPIKeyCreateParams) GetExpiresOn() *time.Time {
	x := p.ExpiresOn
	return &x
}
func (p ExtraSchemaUserAPIKeyUpdateParams) GetExpiresOn() *time.Time {
	return p.ExpiresOn
}

func (p ExtraSchemaUserAPIKeyCreateParams) GetUserID() *ExtraSchemaUserID {
	x := p.UserID
	return &x
}
func (p ExtraSchemaUserAPIKeyUpdateParams) GetUserID() *ExtraSchemaUserID {
	return p.UserID
}

type ExtraSchemaUserAPIKeyID int

// CreateExtraSchemaUserAPIKey creates a new ExtraSchemaUserAPIKey in the database with the given params.
func CreateExtraSchemaUserAPIKey(ctx context.Context, db DB, params *ExtraSchemaUserAPIKeyCreateParams) (*ExtraSchemaUserAPIKey, error) {
	esuak := &ExtraSchemaUserAPIKey{
		APIKey:    params.APIKey,
		ExpiresOn: params.ExpiresOn,
		UserID:    params.UserID,
	}

	return esuak.Insert(ctx, db)
}

type ExtraSchemaUserAPIKeySelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   ExtraSchemaUserAPIKeyJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaUserAPIKeySelectConfigOption func(*ExtraSchemaUserAPIKeySelectConfig)

// WithExtraSchemaUserAPIKeyLimit limits row selection.
func WithExtraSchemaUserAPIKeyLimit(limit int) ExtraSchemaUserAPIKeySelectConfigOption {
	return func(s *ExtraSchemaUserAPIKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaUserAPIKeyOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaUserAPIKeyOrderBy(rows map[string]*Direction) ExtraSchemaUserAPIKeySelectConfigOption {
	return func(s *ExtraSchemaUserAPIKeySelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaUserAPIKey]
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

type ExtraSchemaUserAPIKeyJoins struct {
	User bool `json:"user" required:"true" nullable:"false"` // O2O users
}

// WithExtraSchemaUserAPIKeyJoin joins with the given tables.
func WithExtraSchemaUserAPIKeyJoin(joins ExtraSchemaUserAPIKeyJoins) ExtraSchemaUserAPIKeySelectConfigOption {
	return func(s *ExtraSchemaUserAPIKeySelectConfig) {
		s.joins = ExtraSchemaUserAPIKeyJoins{
			User: s.joins.User || joins.User,
		}
	}
}

// WithExtraSchemaUserAPIKeyFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaUserAPIKeyFilters(filters map[string][]any) ExtraSchemaUserAPIKeySelectConfigOption {
	return func(s *ExtraSchemaUserAPIKeySelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaUserAPIKeyHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaUserAPIKeyHavingClause(conditions map[string][]any) ExtraSchemaUserAPIKeySelectConfigOption {
	return func(s *ExtraSchemaUserAPIKeySelectConfig) {
		s.having = conditions
	}
}

const extraSchemaUserAPIKeyTableUserJoinSQL = `-- O2O join generated from "user_api_keys_user_id_fkey (inferred)"
left join extra_schema.users as _user_api_keys_user_id on _user_api_keys_user_id.user_id = user_api_keys.user_id
`

const extraSchemaUserAPIKeyTableUserSelectSQL = `(case when _user_api_keys_user_id.user_id is not null then row(_user_api_keys_user_id.*) end) as user_user_id`

const extraSchemaUserAPIKeyTableUserGroupBySQL = `_user_api_keys_user_id.user_id,
      _user_api_keys_user_id.user_id,
	user_api_keys.user_api_key_id`

// ExtraSchemaUserAPIKeyUpdateParams represents update params for 'extra_schema.user_api_keys'.
type ExtraSchemaUserAPIKeyUpdateParams struct {
	APIKey    *string            `json:"apiKey" nullable:"false"`    // api_key
	ExpiresOn *time.Time         `json:"expiresOn" nullable:"false"` // expires_on
	UserID    *ExtraSchemaUserID `json:"userID" nullable:"false"`    // user_id
}

// SetUpdateParams updates extra_schema.user_api_keys struct fields with the specified params.
func (esuak *ExtraSchemaUserAPIKey) SetUpdateParams(params *ExtraSchemaUserAPIKeyUpdateParams) {
	if params.APIKey != nil {
		esuak.APIKey = *params.APIKey
	}
	if params.ExpiresOn != nil {
		esuak.ExpiresOn = *params.ExpiresOn
	}
	if params.UserID != nil {
		esuak.UserID = *params.UserID
	}
}

// Insert inserts the ExtraSchemaUserAPIKey to the database.
func (esuak *ExtraSchemaUserAPIKey) Insert(ctx context.Context, db DB) (*ExtraSchemaUserAPIKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.user_api_keys (
	api_key, expires_on, user_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, esuak.APIKey, esuak.ExpiresOn, esuak.UserID)

	rows, err := db.Query(ctx, sqlstr, esuak.APIKey, esuak.ExpiresOn, esuak.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Insert/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newesuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	*esuak = newesuak

	return esuak, nil
}

// Update updates a ExtraSchemaUserAPIKey in the database.
func (esuak *ExtraSchemaUserAPIKey) Update(ctx context.Context, db DB) (*ExtraSchemaUserAPIKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.user_api_keys SET 
	api_key = $1, expires_on = $2, user_id = $3 
	WHERE user_api_key_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, esuak.APIKey, esuak.ExpiresOn, esuak.UserID, esuak.UserAPIKeyID)

	rows, err := db.Query(ctx, sqlstr, esuak.APIKey, esuak.ExpiresOn, esuak.UserID, esuak.UserAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Update/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newesuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}
	*esuak = newesuak

	return esuak, nil
}

// Upsert upserts a ExtraSchemaUserAPIKey in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esuak *ExtraSchemaUserAPIKey) Upsert(ctx context.Context, db DB, params *ExtraSchemaUserAPIKeyCreateParams) (*ExtraSchemaUserAPIKey, error) {
	var err error

	esuak.APIKey = params.APIKey
	esuak.ExpiresOn = params.ExpiresOn
	esuak.UserID = params.UserID

	esuak, err = esuak.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaUserAPIKey/Insert: %w", &XoError{Entity: "User api key", Err: err})
			}
			esuak, err = esuak.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaUserAPIKey/Update: %w", &XoError{Entity: "User api key", Err: err})
			}
		}
	}

	return esuak, err
}

// Delete deletes the ExtraSchemaUserAPIKey from the database.
func (esuak *ExtraSchemaUserAPIKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.user_api_keys 
	WHERE user_api_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esuak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaUserAPIKeyPaginated returns a cursor-paginated list of ExtraSchemaUserAPIKey.
// At least one cursor is required.
func ExtraSchemaUserAPIKeyPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaUserAPIKeySelectConfigOption) ([]ExtraSchemaUserAPIKey, error) {
	c := &ExtraSchemaUserAPIKeySelectConfig{joins: ExtraSchemaUserAPIKeyJoins{},
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
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaUserAPIKey][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Paginated/cursor: %w", &XoError{Entity: "User api key", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
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
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Paginated/orderBy: %w", &XoError{Entity: "User api key", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, extraSchemaUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserAPIKeyTableUserGroupBySQL)
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
	 FROM extra_schema.user_api_keys %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserAPIKeyPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Paginated/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUserAPIKey/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User api key", Err: err}))
	}
	return res, nil
}

// ExtraSchemaUserAPIKeyByAPIKey retrieves a row from 'extra_schema.user_api_keys' as a ExtraSchemaUserAPIKey.
//
// Generated from index 'user_api_keys_api_key_key'.
func ExtraSchemaUserAPIKeyByAPIKey(ctx context.Context, db DB, apiKey string, opts ...ExtraSchemaUserAPIKeySelectConfigOption) (*ExtraSchemaUserAPIKey, error) {
	c := &ExtraSchemaUserAPIKeySelectConfig{joins: ExtraSchemaUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserAPIKeyTableUserGroupBySQL)
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
	 FROM extra_schema.user_api_keys %s 
	 WHERE user_api_keys.api_key = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserAPIKeyByAPIKey */\n" + sqlstr

	// run
	// logf(sqlstr, apiKey)
	rows, err := db.Query(ctx, sqlstr, append([]any{apiKey}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	esuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &esuak, nil
}

// ExtraSchemaUserAPIKeyByUserAPIKeyID retrieves a row from 'extra_schema.user_api_keys' as a ExtraSchemaUserAPIKey.
//
// Generated from index 'user_api_keys_pkey'.
func ExtraSchemaUserAPIKeyByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID ExtraSchemaUserAPIKeyID, opts ...ExtraSchemaUserAPIKeySelectConfigOption) (*ExtraSchemaUserAPIKey, error) {
	c := &ExtraSchemaUserAPIKeySelectConfig{joins: ExtraSchemaUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserAPIKeyTableUserGroupBySQL)
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
	 FROM extra_schema.user_api_keys %s 
	 WHERE user_api_keys.user_api_key_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserAPIKeyByUserAPIKeyID */\n" + sqlstr

	// run
	// logf(sqlstr, userAPIKeyID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userAPIKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	esuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &esuak, nil
}

// ExtraSchemaUserAPIKeyByUserID retrieves a row from 'extra_schema.user_api_keys' as a ExtraSchemaUserAPIKey.
//
// Generated from index 'user_api_keys_user_id_key'.
func ExtraSchemaUserAPIKeyByUserID(ctx context.Context, db DB, userID ExtraSchemaUserID, opts ...ExtraSchemaUserAPIKeySelectConfigOption) (*ExtraSchemaUserAPIKey, error) {
	c := &ExtraSchemaUserAPIKeySelectConfig{joins: ExtraSchemaUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserAPIKeyTableUserGroupBySQL)
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
	 FROM extra_schema.user_api_keys %s 
	 WHERE user_api_keys.user_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserAPIKeyByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	esuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &esuak, nil
}

// FKUser_UserID returns the User associated with the ExtraSchemaUserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (esuak *ExtraSchemaUserAPIKey) FKUser_UserID(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, esuak.UserID)
}
