package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsUserAPIKey represents a row from 'xo_tests.user_api_keys'.
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
type XoTestsUserAPIKey struct {
	UserAPIKeyID XoTestsUserAPIKeyID `json:"-" db:"user_api_key_id" nullable:"false"`                    // user_api_key_id
	APIKey       string              `json:"apiKey" db:"api_key" required:"true" nullable:"false"`       // api_key
	ExpiresOn    time.Time           `json:"expiresOn" db:"expires_on" required:"true" nullable:"false"` // expires_on
	UserID       XoTestsUserID       `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id

	UserJoin *XoTestsUser `json:"-" db:"user_user_id" openapi-go:"ignore"` // O2O users (inferred)
}

// XoTestsUserAPIKeyCreateParams represents insert params for 'xo_tests.user_api_keys'.
type XoTestsUserAPIKeyCreateParams struct {
	APIKey    string        `json:"apiKey" required:"true" nullable:"false"`    // api_key
	ExpiresOn time.Time     `json:"expiresOn" required:"true" nullable:"false"` // expires_on
	UserID    XoTestsUserID `json:"userID" required:"true" nullable:"false"`    // user_id
}

// XoTestsUserAPIKeyParams represents common params for both insert and update of 'xo_tests.user_api_keys'.
type XoTestsUserAPIKeyParams interface {
	GetAPIKey() *string
	GetExpiresOn() *time.Time
	GetUserID() *XoTestsUserID
}

func (p XoTestsUserAPIKeyCreateParams) GetAPIKey() *string {
	x := p.APIKey
	return &x
}

func (p XoTestsUserAPIKeyUpdateParams) GetAPIKey() *string {
	return p.APIKey
}

func (p XoTestsUserAPIKeyCreateParams) GetExpiresOn() *time.Time {
	x := p.ExpiresOn
	return &x
}

func (p XoTestsUserAPIKeyUpdateParams) GetExpiresOn() *time.Time {
	return p.ExpiresOn
}

func (p XoTestsUserAPIKeyCreateParams) GetUserID() *XoTestsUserID {
	x := p.UserID
	return &x
}

func (p XoTestsUserAPIKeyUpdateParams) GetUserID() *XoTestsUserID {
	return p.UserID
}

type XoTestsUserAPIKeyID int

// CreateXoTestsUserAPIKey creates a new XoTestsUserAPIKey in the database with the given params.
func CreateXoTestsUserAPIKey(ctx context.Context, db DB, params *XoTestsUserAPIKeyCreateParams) (*XoTestsUserAPIKey, error) {
	xtuak := &XoTestsUserAPIKey{
		APIKey:    params.APIKey,
		ExpiresOn: params.ExpiresOn,
		UserID:    params.UserID,
	}

	return xtuak.Insert(ctx, db)
}

type XoTestsUserAPIKeySelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsUserAPIKeyJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsUserAPIKeySelectConfigOption func(*XoTestsUserAPIKeySelectConfig)

// WithXoTestsUserAPIKeyLimit limits row selection.
func WithXoTestsUserAPIKeyLimit(limit int) XoTestsUserAPIKeySelectConfigOption {
	return func(s *XoTestsUserAPIKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsUserAPIKeyOrderBy string

const (
	XoTestsUserAPIKeyExpiresOnDescNullsFirst XoTestsUserAPIKeyOrderBy = " expires_on DESC NULLS FIRST "
	XoTestsUserAPIKeyExpiresOnDescNullsLast  XoTestsUserAPIKeyOrderBy = " expires_on DESC NULLS LAST "
	XoTestsUserAPIKeyExpiresOnAscNullsFirst  XoTestsUserAPIKeyOrderBy = " expires_on ASC NULLS FIRST "
	XoTestsUserAPIKeyExpiresOnAscNullsLast   XoTestsUserAPIKeyOrderBy = " expires_on ASC NULLS LAST "
)

// WithXoTestsUserAPIKeyOrderBy orders results by the given columns.
func WithXoTestsUserAPIKeyOrderBy(rows ...XoTestsUserAPIKeyOrderBy) XoTestsUserAPIKeySelectConfigOption {
	return func(s *XoTestsUserAPIKeySelectConfig) {
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

type XoTestsUserAPIKeyJoins struct {
	User bool // O2O users
}

// WithXoTestsUserAPIKeyJoin joins with the given tables.
func WithXoTestsUserAPIKeyJoin(joins XoTestsUserAPIKeyJoins) XoTestsUserAPIKeySelectConfigOption {
	return func(s *XoTestsUserAPIKeySelectConfig) {
		s.joins = XoTestsUserAPIKeyJoins{
			User: s.joins.User || joins.User,
		}
	}
}

// WithXoTestsUserAPIKeyFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsUserAPIKeyFilters(filters map[string][]any) XoTestsUserAPIKeySelectConfigOption {
	return func(s *XoTestsUserAPIKeySelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsUserAPIKeyHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsUserAPIKeyHavingClause(conditions map[string][]any) XoTestsUserAPIKeySelectConfigOption {
	return func(s *XoTestsUserAPIKeySelectConfig) {
		s.having = conditions
	}
}

const xoTestsUserAPIKeyTableUserJoinSQL = `-- O2O join generated from "user_api_keys_user_id_fkey (inferred)"
left join xo_tests.users as _user_api_keys_user_id on _user_api_keys_user_id.user_id = user_api_keys.user_id
`

const xoTestsUserAPIKeyTableUserSelectSQL = `(case when _user_api_keys_user_id.user_id is not null then row(_user_api_keys_user_id.*) end) as user_user_id`

const xoTestsUserAPIKeyTableUserGroupBySQL = `_user_api_keys_user_id.user_id,
      _user_api_keys_user_id.user_id,
	user_api_keys.user_api_key_id`

// XoTestsUserAPIKeyUpdateParams represents update params for 'xo_tests.user_api_keys'.
type XoTestsUserAPIKeyUpdateParams struct {
	APIKey    *string        `json:"apiKey" nullable:"false"`    // api_key
	ExpiresOn *time.Time     `json:"expiresOn" nullable:"false"` // expires_on
	UserID    *XoTestsUserID `json:"userID" nullable:"false"`    // user_id
}

// SetUpdateParams updates xo_tests.user_api_keys struct fields with the specified params.
func (xtuak *XoTestsUserAPIKey) SetUpdateParams(params *XoTestsUserAPIKeyUpdateParams) {
	if params.APIKey != nil {
		xtuak.APIKey = *params.APIKey
	}
	if params.ExpiresOn != nil {
		xtuak.ExpiresOn = *params.ExpiresOn
	}
	if params.UserID != nil {
		xtuak.UserID = *params.UserID
	}
}

// Insert inserts the XoTestsUserAPIKey to the database.
func (xtuak *XoTestsUserAPIKey) Insert(ctx context.Context, db DB) (*XoTestsUserAPIKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.user_api_keys (
	api_key, expires_on, user_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtuak.APIKey, xtuak.ExpiresOn, xtuak.UserID)

	rows, err := db.Query(ctx, sqlstr, xtuak.APIKey, xtuak.ExpiresOn, xtuak.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Insert/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newxtuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	*xtuak = newxtuak

	return xtuak, nil
}

// Update updates a XoTestsUserAPIKey in the database.
func (xtuak *XoTestsUserAPIKey) Update(ctx context.Context, db DB) (*XoTestsUserAPIKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.user_api_keys SET 
	api_key = $1, expires_on = $2, user_id = $3 
	WHERE user_api_key_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, xtuak.APIKey, xtuak.ExpiresOn, xtuak.UserID, xtuak.UserAPIKeyID)

	rows, err := db.Query(ctx, sqlstr, xtuak.APIKey, xtuak.ExpiresOn, xtuak.UserID, xtuak.UserAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Update/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	newxtuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}
	*xtuak = newxtuak

	return xtuak, nil
}

// Upsert upserts a XoTestsUserAPIKey in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtuak *XoTestsUserAPIKey) Upsert(ctx context.Context, db DB, params *XoTestsUserAPIKeyCreateParams) (*XoTestsUserAPIKey, error) {
	var err error

	xtuak.APIKey = params.APIKey
	xtuak.ExpiresOn = params.ExpiresOn
	xtuak.UserID = params.UserID

	xtuak, err = xtuak.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User api key", Err: err})
			}
			xtuak, err = xtuak.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User api key", Err: err})
			}
		}
	}

	return xtuak, err
}

// Delete deletes the XoTestsUserAPIKey from the database.
func (xtuak *XoTestsUserAPIKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.user_api_keys 
	WHERE user_api_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtuak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsUserAPIKeyPaginatedByUserAPIKeyID returns a cursor-paginated list of XoTestsUserAPIKey.
func XoTestsUserAPIKeyPaginatedByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID XoTestsUserAPIKeyID, direction models.Direction, opts ...XoTestsUserAPIKeySelectConfigOption) ([]XoTestsUserAPIKey, error) {
	c := &XoTestsUserAPIKeySelectConfig{joins: XoTestsUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM xo_tests.user_api_keys %s 
	 WHERE user_api_keys.user_api_key_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		user_api_key_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserAPIKeyPaginatedByUserAPIKeyID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{userAPIKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Paginated/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUserAPIKey/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User api key", Err: err}))
	}
	return res, nil
}

// XoTestsUserAPIKeyByAPIKey retrieves a row from 'xo_tests.user_api_keys' as a XoTestsUserAPIKey.
//
// Generated from index 'user_api_keys_api_key_key'.
func XoTestsUserAPIKeyByAPIKey(ctx context.Context, db DB, apiKey string, opts ...XoTestsUserAPIKeySelectConfigOption) (*XoTestsUserAPIKey, error) {
	c := &XoTestsUserAPIKeySelectConfig{joins: XoTestsUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM xo_tests.user_api_keys %s 
	 WHERE user_api_keys.api_key = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserAPIKeyByAPIKey */\n" + sqlstr

	// run
	// logf(sqlstr, apiKey)
	rows, err := db.Query(ctx, sqlstr, append([]any{apiKey}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	xtuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &xtuak, nil
}

// XoTestsUserAPIKeyByUserAPIKeyID retrieves a row from 'xo_tests.user_api_keys' as a XoTestsUserAPIKey.
//
// Generated from index 'user_api_keys_pkey'.
func XoTestsUserAPIKeyByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID XoTestsUserAPIKeyID, opts ...XoTestsUserAPIKeySelectConfigOption) (*XoTestsUserAPIKey, error) {
	c := &XoTestsUserAPIKeySelectConfig{joins: XoTestsUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM xo_tests.user_api_keys %s 
	 WHERE user_api_keys.user_api_key_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserAPIKeyByUserAPIKeyID */\n" + sqlstr

	// run
	// logf(sqlstr, userAPIKeyID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userAPIKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	xtuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &xtuak, nil
}

// XoTestsUserAPIKeyByUserID retrieves a row from 'xo_tests.user_api_keys' as a XoTestsUserAPIKey.
//
// Generated from index 'user_api_keys_user_id_key'.
func XoTestsUserAPIKeyByUserID(ctx context.Context, db DB, userID XoTestsUserID, opts ...XoTestsUserAPIKeySelectConfigOption) (*XoTestsUserAPIKey, error) {
	c := &XoTestsUserAPIKeySelectConfig{joins: XoTestsUserAPIKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsUserAPIKeyTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserAPIKeyTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserAPIKeyTableUserGroupBySQL)
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
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_api_key_id,
	user_api_keys.user_id %s 
	 FROM xo_tests.user_api_keys %s 
	 WHERE user_api_keys.user_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserAPIKeyByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/db.Query: %w", &XoError{Entity: "User api key", Err: err}))
	}
	xtuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User api key", Err: err}))
	}

	return &xtuak, nil
}

// FKUser_UserID returns the User associated with the XoTestsUserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (xtuak *XoTestsUserAPIKey) FKUser_UserID(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtuak.UserID)
}
