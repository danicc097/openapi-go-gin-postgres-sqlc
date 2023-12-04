package got

// Code generated by xo. DO NOT EDIT.

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

// DummyJoin represents a row from 'xo_tests.dummy_join'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type DummyJoin struct {
	DummyJoinID DummyJoinID `json:"dummyJoinID" db:"dummy_join_id" required:"true" nullable:"false"` // dummy_join_id
	Name        *string     `json:"name" db:"name"`                                                  // name
}

// DummyJoinCreateParams represents insert params for 'xo_tests.dummy_join'.
type DummyJoinCreateParams struct {
	Name *string `json:"name"` // name
}

type DummyJoinID int

// CreateDummyJoin creates a new DummyJoin in the database with the given params.
func CreateDummyJoin(ctx context.Context, db DB, params *DummyJoinCreateParams) (*DummyJoin, error) {
	dj := &DummyJoin{
		Name: params.Name,
	}

	return dj.Insert(ctx, db)
}

// DummyJoinUpdateParams represents update params for 'xo_tests.dummy_join'.
type DummyJoinUpdateParams struct {
	Name **string `json:"name"` // name
}

// SetUpdateParams updates xo_tests.dummy_join struct fields with the specified params.
func (dj *DummyJoin) SetUpdateParams(params *DummyJoinUpdateParams) {
	if params.Name != nil {
		dj.Name = *params.Name
	}
}

type DummyJoinSelectConfig struct {
	limit   string
	orderBy string
	joins   DummyJoinJoins
	filters map[string][]any
}
type DummyJoinSelectConfigOption func(*DummyJoinSelectConfig)

// WithDummyJoinLimit limits row selection.
func WithDummyJoinLimit(limit int) DummyJoinSelectConfigOption {
	return func(s *DummyJoinSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type DummyJoinOrderBy string

type DummyJoinJoins struct{}

// WithDummyJoinJoin joins with the given tables.
func WithDummyJoinJoin(joins DummyJoinJoins) DummyJoinSelectConfigOption {
	return func(s *DummyJoinSelectConfig) {
		s.joins = DummyJoinJoins{}
	}
}

// WithDummyJoinFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithDummyJoinFilters(filters map[string][]any) DummyJoinSelectConfigOption {
	return func(s *DummyJoinSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the DummyJoin to the database.
func (dj *DummyJoin) Insert(ctx context.Context, db DB) (*DummyJoin, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.dummy_join (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, dj.Name)

	rows, err := db.Query(ctx, sqlstr, dj.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Insert/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	*dj = newdj

	return dj, nil
}

// Update updates a DummyJoin in the database.
func (dj *DummyJoin) Update(ctx context.Context, db DB) (*DummyJoin, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.dummy_join SET 
	name = $1 
	WHERE dummy_join_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, dj.Name, dj.DummyJoinID)

	rows, err := db.Query(ctx, sqlstr, dj.Name, dj.DummyJoinID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Update/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	newdj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	*dj = newdj

	return dj, nil
}

// Upsert upserts a DummyJoin in the database.
// Requires appropriate PK(s) to be set beforehand.
func (dj *DummyJoin) Upsert(ctx context.Context, db DB, params *DummyJoinCreateParams) (*DummyJoin, error) {
	var err error

	dj.Name = params.Name

	dj, err = dj.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Dummy join", Err: err})
			}
			dj, err = dj.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Dummy join", Err: err})
			}
		}
	}

	return dj, err
}

// Delete deletes the DummyJoin from the database.
func (dj *DummyJoin) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.dummy_join 
	WHERE dummy_join_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dj.DummyJoinID); err != nil {
		return logerror(err)
	}
	return nil
}

// DummyJoinPaginatedByDummyJoinID returns a cursor-paginated list of DummyJoin.
func DummyJoinPaginatedByDummyJoinID(ctx context.Context, db DB, dummyJoinID DummyJoinID, direction Direction, opts ...DummyJoinSelectConfigOption) ([]DummyJoin, error) {
	c := &DummyJoinSelectConfig{joins: DummyJoinJoins{}, filters: make(map[string][]any)}

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
	if direction == DirectionAsc {
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
	sqlstr = "/* DummyJoinPaginatedByDummyJoinID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{dummyJoinID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Paginated/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("DummyJoin/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	return res, nil
}

// DummyJoinByDummyJoinID retrieves a row from 'xo_tests.dummy_join' as a DummyJoin.
//
// Generated from index 'dummy_join_pkey'.
func DummyJoinByDummyJoinID(ctx context.Context, db DB, dummyJoinID DummyJoinID, opts ...DummyJoinSelectConfigOption) (*DummyJoin, error) {
	c := &DummyJoinSelectConfig{joins: DummyJoinJoins{}, filters: make(map[string][]any)}

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
	sqlstr = "/* DummyJoinByDummyJoinID */\n" + sqlstr

	// run
	// logf(sqlstr, dummyJoinID)
	rows, err := db.Query(ctx, sqlstr, append([]any{dummyJoinID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/db.Query: %w", &XoError{Entity: "Dummy join", Err: err}))
	}
	dj, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DummyJoin])
	if err != nil {
		return nil, logerror(fmt.Errorf("dummy_join/DummyJoinByDummyJoinID/pgx.CollectOneRow: %w", &XoError{Entity: "Dummy join", Err: err}))
	}

	return &dj, nil
}
