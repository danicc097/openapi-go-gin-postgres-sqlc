package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// PagElement represents a row from 'xo_tests.pag_element'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type PagElement struct {
	PaginatedElementID uuid.UUID `json:"paginatedElementID" db:"paginated_element_id" required:"true" nullable:"false"` // paginated_element_id
	Name               string    `json:"name" db:"name" required:"true" nullable:"false"`                               // name
	CreatedAt          time.Time `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                    // created_at
	Dummy              *int      `json:"dummy" db:"dummy"`                                                              // dummy

	DummyJoin *DummyJoin `json:"-" db:"dummy_join_dummy" openapi-go:"ignore"` // O2O dummy_join (inferred)
}

// PagElementCreateParams represents insert params for 'xo_tests.pag_element'.
type PagElementCreateParams struct {
	Dummy *int   `json:"dummy"`                                 // dummy
	Name  string `json:"name" required:"true" nullable:"false"` // name
}

// CreatePagElement creates a new PagElement in the database with the given params.
func CreatePagElement(ctx context.Context, db DB, params *PagElementCreateParams) (*PagElement, error) {
	pe := &PagElement{
		Dummy: params.Dummy,
		Name:  params.Name,
	}

	return pe.Insert(ctx, db)
}

// PagElementUpdateParams represents update params for 'xo_tests.pag_element'.
type PagElementUpdateParams struct {
	Dummy **int   `json:"dummy"`                 // dummy
	Name  *string `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.pag_element struct fields with the specified params.
func (pe *PagElement) SetUpdateParams(params *PagElementUpdateParams) {
	if params.Dummy != nil {
		pe.Dummy = *params.Dummy
	}
	if params.Name != nil {
		pe.Name = *params.Name
	}
}

type PagElementSelectConfig struct {
	limit   string
	orderBy string
	joins   PagElementJoins
	filters map[string][]any
}
type PagElementSelectConfigOption func(*PagElementSelectConfig)

// WithPagElementLimit limits row selection.
func WithPagElementLimit(limit int) PagElementSelectConfigOption {
	return func(s *PagElementSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type PagElementOrderBy string

const (
	PagElementCreatedAtDescNullsFirst PagElementOrderBy = " created_at DESC NULLS FIRST "
	PagElementCreatedAtDescNullsLast  PagElementOrderBy = " created_at DESC NULLS LAST "
	PagElementCreatedAtAscNullsFirst  PagElementOrderBy = " created_at ASC NULLS FIRST "
	PagElementCreatedAtAscNullsLast   PagElementOrderBy = " created_at ASC NULLS LAST "
)

// WithPagElementOrderBy orders results by the given columns.
func WithPagElementOrderBy(rows ...PagElementOrderBy) PagElementSelectConfigOption {
	return func(s *PagElementSelectConfig) {
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

type PagElementJoins struct {
	DummyJoin bool // O2O dummy_join
}

// WithPagElementJoin joins with the given tables.
func WithPagElementJoin(joins PagElementJoins) PagElementSelectConfigOption {
	return func(s *PagElementSelectConfig) {
		s.joins = PagElementJoins{
			DummyJoin: s.joins.DummyJoin || joins.DummyJoin,
		}
	}
}

// WithPagElementFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithPagElementFilters(filters map[string][]any) PagElementSelectConfigOption {
	return func(s *PagElementSelectConfig) {
		s.filters = filters
	}
}

const pagElementTableDummyJoinJoinSQL = `-- O2O join generated from "pag_element_dummy_fkey (inferred)"
left join xo_tests.dummy_join as _pag_element_dummy on _pag_element_dummy.dummy_join_id = pag_element.dummy
`

const pagElementTableDummyJoinSelectSQL = `(case when _pag_element_dummy.dummy_join_id is not null then row(_pag_element_dummy.*) end) as dummy_join_dummy`

const pagElementTableDummyJoinGroupBySQL = `_pag_element_dummy.dummy_join_id,
      _pag_element_dummy.dummy_join_id,
	pag_element.paginated_element_id`

// Insert inserts the PagElement to the database.
func (pe *PagElement) Insert(ctx context.Context, db DB) (*PagElement, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.pag_element (
	dummy, name
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, pe.Dummy, pe.Name)

	rows, err := db.Query(ctx, sqlstr, pe.Dummy, pe.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Insert/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	*pe = newpe

	return pe, nil
}

// Update updates a PagElement in the database.
func (pe *PagElement) Update(ctx context.Context, db DB) (*PagElement, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.pag_element SET 
	dummy = $1, name = $2 
	WHERE paginated_element_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, pe.CreatedAt, pe.Dummy, pe.Name, pe.PaginatedElementID)

	rows, err := db.Query(ctx, sqlstr, pe.Dummy, pe.Name, pe.PaginatedElementID)
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Update/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	*pe = newpe

	return pe, nil
}

// Upsert upserts a PagElement in the database.
// Requires appropriate PK(s) to be set beforehand.
func (pe *PagElement) Upsert(ctx context.Context, db DB, params *PagElementCreateParams) (*PagElement, error) {
	var err error

	pe.Dummy = params.Dummy
	pe.Name = params.Name

	pe, err = pe.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Pag element", Err: err})
			}
			pe, err = pe.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Pag element", Err: err})
			}
		}
	}

	return pe, err
}

// Delete deletes the PagElement from the database.
func (pe *PagElement) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.pag_element 
	WHERE paginated_element_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, pe.PaginatedElementID); err != nil {
		return logerror(err)
	}
	return nil
}

// PagElementPaginatedByCreatedAtAsc returns a cursor-paginated list of PagElement in Asc order.
func PagElementPaginatedByCreatedAtAsc(ctx context.Context, db DB, createdAt time.Time, opts ...PagElementSelectConfigOption) ([]PagElement, error) {
	c := &PagElementSelectConfig{joins: PagElementJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, pagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, pagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, pagElementTableDummyJoinGroupBySQL)
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
	pag_element.created_at,
	pag_element.dummy,
	pag_element.name,
	pag_element.paginated_element_id %s 
	 FROM xo_tests.pag_element %s 
	 WHERE pag_element.created_at > $1
	 %s   %s 
  ORDER BY 
		created_at Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* PagElementPaginatedByCreatedAtAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Paginated/Asc/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	return res, nil
}

// PagElementPaginatedByCreatedAtDesc returns a cursor-paginated list of PagElement in Desc order.
func PagElementPaginatedByCreatedAtDesc(ctx context.Context, db DB, createdAt time.Time, opts ...PagElementSelectConfigOption) ([]PagElement, error) {
	c := &PagElementSelectConfig{joins: PagElementJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, pagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, pagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, pagElementTableDummyJoinGroupBySQL)
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
	pag_element.created_at,
	pag_element.dummy,
	pag_element.name,
	pag_element.paginated_element_id %s 
	 FROM xo_tests.pag_element %s 
	 WHERE pag_element.created_at < $1
	 %s   %s 
  ORDER BY 
		created_at Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* PagElementPaginatedByCreatedAtDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Paginated/Desc/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("PagElement/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	return res, nil
}

// PagElementByCreatedAt retrieves a row from 'xo_tests.pag_element' as a PagElement.
//
// Generated from index 'pag_element_created_at_key'.
func PagElementByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...PagElementSelectConfigOption) (*PagElement, error) {
	c := &PagElementSelectConfig{joins: PagElementJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, pagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, pagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, pagElementTableDummyJoinGroupBySQL)
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
	pag_element.created_at,
	pag_element.dummy,
	pag_element.name,
	pag_element.paginated_element_id %s 
	 FROM xo_tests.pag_element %s 
	 WHERE pag_element.created_at = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* PagElementByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	pe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &pe, nil
}

// PagElementByPaginatedElementID retrieves a row from 'xo_tests.pag_element' as a PagElement.
//
// Generated from index 'pag_element_pkey'.
func PagElementByPaginatedElementID(ctx context.Context, db DB, paginatedElementID uuid.UUID, opts ...PagElementSelectConfigOption) (*PagElement, error) {
	c := &PagElementSelectConfig{joins: PagElementJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, pagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, pagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, pagElementTableDummyJoinGroupBySQL)
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
	pag_element.created_at,
	pag_element.dummy,
	pag_element.name,
	pag_element.paginated_element_id %s 
	 FROM xo_tests.pag_element %s 
	 WHERE pag_element.paginated_element_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* PagElementByPaginatedElementID */\n" + sqlstr

	// run
	// logf(sqlstr, paginatedElementID)
	rows, err := db.Query(ctx, sqlstr, append([]any{paginatedElementID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	pe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[PagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &pe, nil
}

// FKDummyJoin_Dummy returns the DummyJoin associated with the PagElement's (Dummy).
//
// Generated from foreign key 'pag_element_dummy_fkey'.
func (pe *PagElement) FKDummyJoin_Dummy(ctx context.Context, db DB) (*DummyJoin, error) {
	return DummyJoinByDummyJoinID(ctx, db, *pe.Dummy)
}
