package db

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

// UserAPIKey represents a row from 'public.user_api_keys'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type UserAPIKey struct {
	UserAPIKeyID UserAPIKeyID `json:"-" db:"user_api_key_id" nullable:"false"`                    // user_api_key_id
	APIKey       string       `json:"apiKey" db:"api_key" required:"true" nullable:"false"`       // api_key
	ExpiresOn    time.Time    `json:"expiresOn" db:"expires_on" required:"true" nullable:"false"` // expires_on
	UserID       UserID       `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id

	UserJoin *User `json:"-" db:"user_user_id" openapi-go:"ignore"` // O2O users (inferred)

}

// UserAPIKeyCreateParams represents insert params for 'public.user_api_keys'.
type UserAPIKeyCreateParams struct {
	APIKey    string    `json:"apiKey" required:"true" nullable:"false"`    // api_key
	ExpiresOn time.Time `json:"expiresOn" required:"true" nullable:"false"` // expires_on
	UserID    UserID    `json:"userID" required:"true" nullable:"false"`    // user_id
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
	orderBy string
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

type UserAPIKeyOrderBy string

const (
	UserAPIKeyExpiresOnDescNullsFirst UserAPIKeyOrderBy = " expires_on DESC NULLS FIRST "
	UserAPIKeyExpiresOnDescNullsLast  UserAPIKeyOrderBy = " expires_on DESC NULLS LAST "
	UserAPIKeyExpiresOnAscNullsFirst  UserAPIKeyOrderBy = " expires_on ASC NULLS FIRST "
	UserAPIKeyExpiresOnAscNullsLast   UserAPIKeyOrderBy = " expires_on ASC NULLS LAST "
)

// WithUserAPIKeyOrderBy orders results by the given columns.
func WithUserAPIKeyOrderBy(rows ...UserAPIKeyOrderBy) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
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

type UserAPIKeyJoins struct {
	User bool // O2O users
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
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User api key", Err: err})
			}
			uak, err = uak.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User api key", Err: err})
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

// UserAPIKeyPaginatedByUserAPIKeyID returns a cursor-paginated list of UserAPIKey.
func UserAPIKeyPaginatedByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID UserAPIKeyID, direction models.Direction, opts ...UserAPIKeySelectConfigOption) ([]UserAPIKey, error) {
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
	 FROM public.user_api_keys %s 
	 WHERE user_api_keys.user_api_key_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		user_api_key_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* UserAPIKeyPaginatedByUserAPIKeyID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{userAPIKeyID}, append(filterParams, havingParams...)...)...)
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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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
