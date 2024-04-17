// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

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

// ExtraSchemaDummyJoin represents a row from 'extra_schema.dummy_join'.
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
type ExtraSchemaDummyJoin struct {
	DummyJoinID ExtraSchemaDummyJoinID `json:"dummyJoinID" db:"dummy_join_id" required:"true" nullable:"false"` // dummy_join_id
	Name        *string                `json:"name" db:"name"`                                                  // name

}

// ExtraSchemaDummyJoinCreateParams represents insert params for 'extra_schema.dummy_join'.
type ExtraSchemaDummyJoinCreateParams struct {
	Name *string `json:"name"` // name
}

// ExtraSchemaDummyJoinParams represents common params for both insert and update of 'extra_schema.dummy_join'.
type ExtraSchemaDummyJoinParams interface {
	GetName() *string
}

func (p ExtraSchemaDummyJoinCreateParams) GetName() *string {
	return p.Name
}
func (p ExtraSchemaDummyJoinUpdateParams) GetName() *string {
	if p.Name != nil {
		return *p.Name
	}
	return nil
}

type ExtraSchemaDummyJoinID int

// CreateExtraSchemaDummyJoin creates a new ExtraSchemaDummyJoin in the database with the given params.
func CreateExtraSchemaDummyJoin(ctx context.Context, db DB, params *ExtraSchemaDummyJoinCreateParams) (*ExtraSchemaDummyJoin, error) {
	esdj := &ExtraSchemaDummyJoin{
		Name: params.Name,
	}

	return esdj.Insert(ctx, db)
}

type ExtraSchemaDummyJoinSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   ExtraSchemaDummyJoinJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaDummyJoinSelectConfigOption func(*ExtraSchemaDummyJoinSelectConfig)

// WithExtraSchemaDummyJoinLimit limits row selection.
func WithExtraSchemaDummyJoinLimit(limit int) ExtraSchemaDummyJoinSelectConfigOption {
	return func(s *ExtraSchemaDummyJoinSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaDummyJoinOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaDummyJoinOrderBy(rows map[string]*models.Direction) ExtraSchemaDummyJoinSelectConfigOption {
	return func(s *ExtraSchemaDummyJoinSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaDummyJoin]
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

type ExtraSchemaDummyJoinJoins struct {
}

// WithExtraSchemaDummyJoinJoin joins with the given tables.
func WithExtraSchemaDummyJoinJoin(joins ExtraSchemaDummyJoinJoins) ExtraSchemaDummyJoinSelectConfigOption {
	return func(s *ExtraSchemaDummyJoinSelectConfig) {
		s.joins = ExtraSchemaDummyJoinJoins{}
	}
}

// WithExtraSchemaDummyJoinFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaDummyJoinFilters(filters map[string][]any) ExtraSchemaDummyJoinSelectConfigOption {
	return func(s *ExtraSchemaDummyJoinSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaDummyJoinHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaDummyJoinHavingClause(conditions map[string][]any) ExtraSchemaDummyJoinSelectConfigOption {
	return func(s *ExtraSchemaDummyJoinSelectConfig) {
		s.having = conditions
	}
}

// ExtraSchemaDummyJoinUpdateParams represents update params for 'extra_schema.dummy_join'.
type ExtraSchemaDummyJoinUpdateParams struct {
	Name **string `json:"name"` // name
}

// SetUpdateParams updates extra_schema.dummy_join struct fields with the specified params.
func (esdj *ExtraSchemaDummyJoin) SetUpdateParams(params *ExtraSchemaDummyJoinUpdateParams) {
	if params.Name != nil {
		esdj.Name = *params.Name
	}
}

// Insert inserts the ExtraSchemaDummyJoin to the database.
func (esdj *ExtraSchemaDummyJoin) Insert(ctx context.Context, db DB) (*ExtraSchemaDummyJoin, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.dummy_join (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, esdj.Name)

	rows, err := db.Query(ctx, sqlstr, esdj.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Insert/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newesdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	*esdj = newesdj

	return esdj, nil
}

// Update updates a ExtraSchemaDummyJoin in the database.
func (esdj *ExtraSchemaDummyJoin) Update(ctx context.Context, db DB) (*ExtraSchemaDummyJoin, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.dummy_join SET 
	name = $1 
	WHERE dummy_join_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, esdj.Name, esdj.DummyJoinID)

	rows, err := db.Query(ctx, sqlstr, esdj.Name, esdj.DummyJoinID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Update/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newesdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	*esdj = newesdj

	return esdj, nil
}

// Upsert upserts a ExtraSchemaDummyJoin in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esdj *ExtraSchemaDummyJoin) Upsert(ctx context.Context, db DB, params *ExtraSchemaDummyJoinCreateParams) (*ExtraSchemaDummyJoin, error) {
	var err error

	esdj.Name = params.Name

	esdj, err = esdj.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaDummyJoin/Insert: %w", &XoError{Entity: "Dummy join", Err: err})
			}
			esdj, err = esdj.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaDummyJoin/Update: %w", &XoError{Entity: "Dummy join", Err: err})
			}
		}
	}

	return esdj, err
}

// Delete deletes the ExtraSchemaDummyJoin from the database.
func (esdj *ExtraSchemaDummyJoin) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.dummy_join 
	WHERE dummy_join_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esdj.DummyJoinID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaDummyJoinPaginated returns a cursor-paginated list of ExtraSchemaDummyJoin.
// At least one cursor is required.
func ExtraSchemaDummyJoinPaginated(ctx context.Context, db DB, cursor models.PaginationCursor, opts ...ExtraSchemaDummyJoinSelectConfigOption) ([]ExtraSchemaDummyJoin, error) {
	c := &ExtraSchemaDummyJoinSelectConfig{joins: ExtraSchemaDummyJoinJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {

		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaDummyJoin][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Paginated/cursor: %w", &XoError{Entity: "Dummy join", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == models.DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("dummy_join.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Paginated/orderBy: %w", &XoError{Entity: "Dummy join", Err: fmt.Errorf("at least one sorted column is required")}))
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
	dummy_join.dummy_join_id,
	dummy_join.name %s 
	 FROM extra_schema.dummy_join %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaDummyJoinPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Paginated/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDummyJoin/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	return res, nil
}

// ExtraSchemaDummyJoinByDummyJoinID retrieves a row from 'extra_schema.dummy_join' as a ExtraSchemaDummyJoin.
//
// Generated from index 'dummy_join_pkey'.
func ExtraSchemaDummyJoinByDummyJoinID(ctx context.Context, db DB, dummyJoinID ExtraSchemaDummyJoinID, opts ...ExtraSchemaDummyJoinSelectConfigOption) (*ExtraSchemaDummyJoin, error) {
	c := &ExtraSchemaDummyJoinSelectConfig{joins: ExtraSchemaDummyJoinJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	dummy_join.dummy_join_id,
	dummy_join.name %s 
	 FROM extra_schema.dummy_join %s 
	 WHERE dummy_join.dummy_join_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaDummyJoinByDummyJoinID */\n" + sqlstr

	// run
	// logf(sqlstr, dummyJoinID)
	rows, err := db.Query(ctx, sqlstr, append([]any{dummyJoinID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	esdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	return &esdj, nil
}
