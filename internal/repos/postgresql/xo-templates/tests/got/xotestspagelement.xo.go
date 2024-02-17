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

	"github.com/google/uuid"
)

// XoTestsPagElement represents a row from 'xo_tests.pag_element'.
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
type XoTestsPagElement struct {
	PaginatedElementID XoTestsPagElementID `json:"paginatedElementID" db:"paginated_element_id" required:"true" nullable:"false"` // paginated_element_id
	Name               string              `json:"name" db:"name" required:"true" nullable:"false"`                               // name
	CreatedAt          time.Time           `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                    // created_at
	Dummy              *XoTestsDummyJoinID `json:"dummy" db:"dummy"`                                                              // dummy

	DummyJoin *XoTestsDummyJoin `json:"-" db:"dummy_join_dummy" openapi-go:"ignore"` // O2O dummy_join (inferred)
}

// XoTestsPagElementCreateParams represents insert params for 'xo_tests.pag_element'.
type XoTestsPagElementCreateParams struct {
	Dummy *XoTestsDummyJoinID `json:"dummy"`                                 // dummy
	Name  string              `json:"name" required:"true" nullable:"false"` // name
}

type XoTestsPagElementID struct {
	uuid.UUID
}

func NewXoTestsPagElementID(id uuid.UUID) XoTestsPagElementID {
	return XoTestsPagElementID{
		UUID: id,
	}
}

// CreateXoTestsPagElement creates a new XoTestsPagElement in the database with the given params.
func CreateXoTestsPagElement(ctx context.Context, db DB, params *XoTestsPagElementCreateParams) (*XoTestsPagElement, error) {
	xtpe := &XoTestsPagElement{
		Dummy: params.Dummy,
		Name:  params.Name,
	}

	return xtpe.Insert(ctx, db)
}

type XoTestsPagElementSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsPagElementJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsPagElementSelectConfigOption func(*XoTestsPagElementSelectConfig)

// WithXoTestsPagElementLimit limits row selection.
func WithXoTestsPagElementLimit(limit int) XoTestsPagElementSelectConfigOption {
	return func(s *XoTestsPagElementSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsPagElementOrderBy string

const (
	XoTestsPagElementCreatedAtDescNullsFirst XoTestsPagElementOrderBy = " created_at DESC NULLS FIRST "
	XoTestsPagElementCreatedAtDescNullsLast  XoTestsPagElementOrderBy = " created_at DESC NULLS LAST "
	XoTestsPagElementCreatedAtAscNullsFirst  XoTestsPagElementOrderBy = " created_at ASC NULLS FIRST "
	XoTestsPagElementCreatedAtAscNullsLast   XoTestsPagElementOrderBy = " created_at ASC NULLS LAST "
)

// WithXoTestsPagElementOrderBy orders results by the given columns.
func WithXoTestsPagElementOrderBy(rows ...XoTestsPagElementOrderBy) XoTestsPagElementSelectConfigOption {
	return func(s *XoTestsPagElementSelectConfig) {
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

type XoTestsPagElementJoins struct {
	DummyJoin bool // O2O dummy_join
}

// WithXoTestsPagElementJoin joins with the given tables.
func WithXoTestsPagElementJoin(joins XoTestsPagElementJoins) XoTestsPagElementSelectConfigOption {
	return func(s *XoTestsPagElementSelectConfig) {
		s.joins = XoTestsPagElementJoins{
			DummyJoin: s.joins.DummyJoin || joins.DummyJoin,
		}
	}
}

// WithXoTestsPagElementFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsPagElementFilters(filters map[string][]any) XoTestsPagElementSelectConfigOption {
	return func(s *XoTestsPagElementSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsPagElementHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsPagElementHavingClause(conditions map[string][]any) XoTestsPagElementSelectConfigOption {
	return func(s *XoTestsPagElementSelectConfig) {
		s.having = conditions
	}
}

const xoTestsPagElementTableDummyJoinJoinSQL = `-- O2O join generated from "pag_element_dummy_fkey (inferred)"
left join xo_tests.dummy_join as _pag_element_dummy on _pag_element_dummy.dummy_join_id = pag_element.dummy
`

const xoTestsPagElementTableDummyJoinSelectSQL = `(case when _pag_element_dummy.dummy_join_id is not null then row(_pag_element_dummy.*) end) as dummy_join_dummy`

const xoTestsPagElementTableDummyJoinGroupBySQL = `_pag_element_dummy.dummy_join_id,
      _pag_element_dummy.dummy_join_id,
	pag_element.paginated_element_id`

// XoTestsPagElementUpdateParams represents update params for 'xo_tests.pag_element'.
type XoTestsPagElementUpdateParams struct {
	Dummy **XoTestsDummyJoinID `json:"dummy"`                 // dummy
	Name  *string              `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.pag_element struct fields with the specified params.
func (xtpe *XoTestsPagElement) SetUpdateParams(params *XoTestsPagElementUpdateParams) {
	if params.Dummy != nil {
		xtpe.Dummy = *params.Dummy
	}
	if params.Name != nil {
		xtpe.Name = *params.Name
	}
}

// Insert inserts the XoTestsPagElement to the database.
func (xtpe *XoTestsPagElement) Insert(ctx context.Context, db DB) (*XoTestsPagElement, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.pag_element (
	dummy, name
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, xtpe.Dummy, xtpe.Name)

	rows, err := db.Query(ctx, sqlstr, xtpe.Dummy, xtpe.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Insert/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newxtpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	*xtpe = newxtpe

	return xtpe, nil
}

// Update updates a XoTestsPagElement in the database.
func (xtpe *XoTestsPagElement) Update(ctx context.Context, db DB) (*XoTestsPagElement, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.pag_element SET
	dummy = $1, name = $2
	WHERE paginated_element_id = $3
	RETURNING * `
	// run
	logf(sqlstr, xtpe.CreatedAt, xtpe.Dummy, xtpe.Name, xtpe.PaginatedElementID)

	rows, err := db.Query(ctx, sqlstr, xtpe.Dummy, xtpe.Name, xtpe.PaginatedElementID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Update/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newxtpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	*xtpe = newxtpe

	return xtpe, nil
}

// Upsert upserts a XoTestsPagElement in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtpe *XoTestsPagElement) Upsert(ctx context.Context, db DB, params *XoTestsPagElementCreateParams) (*XoTestsPagElement, error) {
	var err error

	xtpe.Dummy = params.Dummy
	xtpe.Name = params.Name

	xtpe, err = xtpe.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Pag element", Err: err})
			}
			xtpe, err = xtpe.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Pag element", Err: err})
			}
		}
	}

	return xtpe, err
}

// Delete deletes the XoTestsPagElement from the database.
func (xtpe *XoTestsPagElement) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.pag_element
	WHERE paginated_element_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtpe.PaginatedElementID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsPagElementPaginatedByCreatedAt returns a cursor-paginated list of XoTestsPagElement.
func XoTestsPagElementPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, direction models.Direction, opts ...XoTestsPagElementSelectConfigOption) ([]XoTestsPagElement, error) {
	c := &XoTestsPagElementSelectConfig{joins: XoTestsPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, xoTestsPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, xoTestsPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsPagElementTableDummyJoinGroupBySQL)
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
	pag_element.created_at,
	pag_element.dummy,
	pag_element.name,
	pag_element.paginated_element_id %s
	 FROM xo_tests.pag_element %s
	 WHERE pag_element.created_at %s $1
	 %s   %s
  %s
  ORDER BY
		created_at %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsPagElementPaginatedByCreatedAt */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Paginated/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsPagElement/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	return res, nil
}

// XoTestsPagElementByCreatedAt retrieves a row from 'xo_tests.pag_element' as a XoTestsPagElement.
//
// Generated from index 'pag_element_created_at_key'.
func XoTestsPagElementByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...XoTestsPagElementSelectConfigOption) (*XoTestsPagElement, error) {
	c := &XoTestsPagElementSelectConfig{joins: XoTestsPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, xoTestsPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, xoTestsPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsPagElementTableDummyJoinGroupBySQL)
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
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsPagElementByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	xtpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &xtpe, nil
}

// XoTestsPagElementByPaginatedElementID retrieves a row from 'xo_tests.pag_element' as a XoTestsPagElement.
//
// Generated from index 'pag_element_pkey'.
func XoTestsPagElementByPaginatedElementID(ctx context.Context, db DB, paginatedElementID XoTestsPagElementID, opts ...XoTestsPagElementSelectConfigOption) (*XoTestsPagElement, error) {
	c := &XoTestsPagElementSelectConfig{joins: XoTestsPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DummyJoin {
		selectClauses = append(selectClauses, xoTestsPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, xoTestsPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsPagElementTableDummyJoinGroupBySQL)
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
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsPagElementByPaginatedElementID */\n" + sqlstr

	// run
	// logf(sqlstr, paginatedElementID)
	rows, err := db.Query(ctx, sqlstr, append([]any{paginatedElementID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	xtpe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &xtpe, nil
}

// FKDummyJoin_Dummy returns the DummyJoin associated with the XoTestsPagElement's (Dummy).
//
// Generated from foreign key 'pag_element_dummy_fkey'.
func (xtpe *XoTestsPagElement) FKDummyJoin_Dummy(ctx context.Context, db DB) (*XoTestsDummyJoin, error) {
	return XoTestsDummyJoinByDummyJoinID(ctx, db, *xtpe.Dummy)
}
