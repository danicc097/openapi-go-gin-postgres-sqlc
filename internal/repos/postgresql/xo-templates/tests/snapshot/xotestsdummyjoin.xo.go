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

// XoTestsDummyJoin represents a row from 'xo_tests.dummy_join'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsDummyJoin struct {
	DummyJoinID XoTestsDummyJoinID `json:"dummyJoinID" db:"dummy_join_id" required:"true" nullable:"false"` // dummy_join_id
	Name        *string            `json:"name" db:"name"`                                                  // name
}

// XoTestsDummyJoinCreateParams represents insert params for 'xo_tests.dummy_join'.
type XoTestsDummyJoinCreateParams struct {
	Name *string `json:"name"` // name
}

type XoTestsDummyJoinID int

// CreateXoTestsDummyJoin creates a new XoTestsDummyJoin in the database with the given params.
func CreateXoTestsDummyJoin(ctx context.Context, db DB, params *XoTestsDummyJoinCreateParams) (*XoTestsDummyJoin, error) {
	xtdj := &XoTestsDummyJoin{
		Name: params.Name,
	}

	return xtdj.Insert(ctx, db)
}

// XoTestsDummyJoinUpdateParams represents update params for 'xo_tests.dummy_join'.
type XoTestsDummyJoinUpdateParams struct {
	Name **string `json:"name"` // name
}

// SetUpdateParams updates xo_tests.dummy_join struct fields with the specified params.
func (xtdj *XoTestsDummyJoin) SetUpdateParams(params *XoTestsDummyJoinUpdateParams) {
	if params.Name != nil {
		xtdj.Name = *params.Name
	}
}

type XoTestsDummyJoinSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsDummyJoinJoins
	filters map[string][]any
}
type XoTestsDummyJoinSelectConfigOption func(*XoTestsDummyJoinSelectConfig)

// WithXoTestsDummyJoinLimit limits row selection.
func WithXoTestsDummyJoinLimit(limit int) XoTestsDummyJoinSelectConfigOption {
	return func(s *XoTestsDummyJoinSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsDummyJoinOrderBy string

type XoTestsDummyJoinJoins struct{}

// WithXoTestsDummyJoinJoin joins with the given tables.
func WithXoTestsDummyJoinJoin(joins XoTestsDummyJoinJoins) XoTestsDummyJoinSelectConfigOption {
	return func(s *XoTestsDummyJoinSelectConfig) {
		s.joins = XoTestsDummyJoinJoins{}
	}
}

// WithXoTestsDummyJoinFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsDummyJoinFilters(filters map[string][]any) XoTestsDummyJoinSelectConfigOption {
	return func(s *XoTestsDummyJoinSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the XoTestsDummyJoin to the database.
func (xtdj *XoTestsDummyJoin) Insert(ctx context.Context, db DB) (*XoTestsDummyJoin, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.dummy_join (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, xtdj.Name)

	rows, err := db.Query(ctx, sqlstr, xtdj.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Insert/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newxtdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	*xtdj = newxtdj

	return xtdj, nil
}

// Update updates a XoTestsDummyJoin in the database.
func (xtdj *XoTestsDummyJoin) Update(ctx context.Context, db DB) (*XoTestsDummyJoin, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.dummy_join SET
	name = $1
	WHERE dummy_join_id = $2
	RETURNING * `
	// run
	logf(sqlstr, xtdj.Name, xtdj.DummyJoinID)

	rows, err := db.Query(ctx, sqlstr, xtdj.Name, xtdj.DummyJoinID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Update/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newxtdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	*xtdj = newxtdj

	return xtdj, nil
}

// Upsert upserts a XoTestsDummyJoin in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtdj *XoTestsDummyJoin) Upsert(ctx context.Context, db DB, params *XoTestsDummyJoinCreateParams) (*XoTestsDummyJoin, error) {
	var err error

	xtdj.Name = params.Name

	xtdj, err = xtdj.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Dummy join", Err: err})
			}
			xtdj, err = xtdj.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Dummy join", Err: err})
			}
		}
	}

	return xtdj, err
}

// Delete deletes the XoTestsDummyJoin from the database.
func (xtdj *XoTestsDummyJoin) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.dummy_join
	WHERE dummy_join_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtdj.DummyJoinID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsDummyJoinPaginatedByDummyJoinID returns a cursor-paginated list of XoTestsDummyJoin.
func XoTestsDummyJoinPaginatedByDummyJoinID(ctx context.Context, db DB, dummyJoinID XoTestsDummyJoinID, direction models.Direction, opts ...XoTestsDummyJoinSelectConfigOption) ([]XoTestsDummyJoin, error) {
	c := &XoTestsDummyJoinSelectConfig{joins: XoTestsDummyJoinJoins{}, filters: make(map[string][]any)}

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
	dummy_join.dummy_join_id,
	dummy_join.name %s
	 FROM xo_tests.dummy_join %s
	 WHERE dummy_join.dummy_join_id %s $1
	 %s   %s
  ORDER BY
		dummy_join_id %s `, selects, joins, operator, filters, groupbys, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsDummyJoinPaginatedByDummyJoinID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{dummyJoinID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Paginated/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDummyJoin/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	return res, nil
}

// XoTestsDummyJoinByDummyJoinID retrieves a row from 'xo_tests.dummy_join' as a XoTestsDummyJoin.
//
// Generated from index 'dummy_join_pkey'.
func XoTestsDummyJoinByDummyJoinID(ctx context.Context, db DB, dummyJoinID XoTestsDummyJoinID, opts ...XoTestsDummyJoinSelectConfigOption) (*XoTestsDummyJoin, error) {
	c := &XoTestsDummyJoinSelectConfig{joins: XoTestsDummyJoinJoins{}, filters: make(map[string][]any)}

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
	dummy_join.dummy_join_id,
	dummy_join.name %s
	 FROM xo_tests.dummy_join %s
	 WHERE dummy_join.dummy_join_id = $1
	 %s   %s
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsDummyJoinByDummyJoinID */\n" + sqlstr

	// run
	// logf(sqlstr, dummyJoinID)
	rows, err := db.Query(ctx, sqlstr, append([]any{dummyJoinID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	xtdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	return &xtdj, nil
}
