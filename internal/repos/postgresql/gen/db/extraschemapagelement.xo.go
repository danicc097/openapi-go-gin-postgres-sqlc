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

	"github.com/google/uuid"
)

// ExtraSchemaPagElement represents a row from 'extra_schema.pag_element'.
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
type ExtraSchemaPagElement struct {
	PaginatedElementID ExtraSchemaPagElementID `json:"paginatedElementID" db:"paginated_element_id" required:"true" nullable:"false"` // paginated_element_id
	Name               string                  `json:"name" db:"name" required:"true" nullable:"false"`                               // name
	CreatedAt          time.Time               `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                    // created_at
	Dummy              *ExtraSchemaDummyJoinID `json:"dummy" db:"dummy"`                                                              // dummy

	DummyJoinJoin *ExtraSchemaDummyJoin `json:"-" db:"dummy_join_dummy" openapi-go:"ignore"` // O2O dummy_join (inferred)

}

// ExtraSchemaPagElementCreateParams represents insert params for 'extra_schema.pag_element'.
type ExtraSchemaPagElementCreateParams struct {
	Dummy *ExtraSchemaDummyJoinID `json:"dummy"`                                 // dummy
	Name  string                  `json:"name" required:"true" nullable:"false"` // name
}

// ExtraSchemaPagElementParams represents common params for both insert and update of 'extra_schema.pag_element'.
type ExtraSchemaPagElementParams interface {
	GetDummy() *ExtraSchemaDummyJoinID
	GetName() *string
}

func (p ExtraSchemaPagElementCreateParams) GetDummy() *ExtraSchemaDummyJoinID {
	return p.Dummy
}
func (p ExtraSchemaPagElementUpdateParams) GetDummy() *ExtraSchemaDummyJoinID {
	if p.Dummy != nil {
		return *p.Dummy
	}
	return nil
}

func (p ExtraSchemaPagElementCreateParams) GetName() *string {
	x := p.Name
	return &x
}
func (p ExtraSchemaPagElementUpdateParams) GetName() *string {
	return p.Name
}

type ExtraSchemaPagElementID struct {
	uuid.UUID
}

func NewExtraSchemaPagElementID(id uuid.UUID) ExtraSchemaPagElementID {
	return ExtraSchemaPagElementID{
		UUID: id,
	}
}

// CreateExtraSchemaPagElement creates a new ExtraSchemaPagElement in the database with the given params.
func CreateExtraSchemaPagElement(ctx context.Context, db DB, params *ExtraSchemaPagElementCreateParams) (*ExtraSchemaPagElement, error) {
	espe := &ExtraSchemaPagElement{
		Dummy: params.Dummy,
		Name:  params.Name,
	}

	return espe.Insert(ctx, db)
}

type ExtraSchemaPagElementSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaPagElementJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaPagElementSelectConfigOption func(*ExtraSchemaPagElementSelectConfig)

// WithExtraSchemaPagElementLimit limits row selection.
func WithExtraSchemaPagElementLimit(limit int) ExtraSchemaPagElementSelectConfigOption {
	return func(s *ExtraSchemaPagElementSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaPagElementOrderBy string

const (
	ExtraSchemaPagElementCreatedAtDescNullsFirst ExtraSchemaPagElementOrderBy = " created_at DESC NULLS FIRST "
	ExtraSchemaPagElementCreatedAtDescNullsLast  ExtraSchemaPagElementOrderBy = " created_at DESC NULLS LAST "
	ExtraSchemaPagElementCreatedAtAscNullsFirst  ExtraSchemaPagElementOrderBy = " created_at ASC NULLS FIRST "
	ExtraSchemaPagElementCreatedAtAscNullsLast   ExtraSchemaPagElementOrderBy = " created_at ASC NULLS LAST "
)

// WithExtraSchemaPagElementOrderBy orders results by the given columns.
func WithExtraSchemaPagElementOrderBy(rows ...ExtraSchemaPagElementOrderBy) ExtraSchemaPagElementSelectConfigOption {
	return func(s *ExtraSchemaPagElementSelectConfig) {
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

type ExtraSchemaPagElementJoins struct {
	DummyJoin bool `json:"dummyJoin" required:"true" nullable:"false"` // O2O dummy_join
}

// WithExtraSchemaPagElementJoin joins with the given tables.
func WithExtraSchemaPagElementJoin(joins ExtraSchemaPagElementJoins) ExtraSchemaPagElementSelectConfigOption {
	return func(s *ExtraSchemaPagElementSelectConfig) {
		s.joins = ExtraSchemaPagElementJoins{
			DummyJoin: s.joins.DummyJoin || joins.DummyJoin,
		}
	}
}

// WithExtraSchemaPagElementFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaPagElementFilters(filters map[string][]any) ExtraSchemaPagElementSelectConfigOption {
	return func(s *ExtraSchemaPagElementSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaPagElementHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaPagElementHavingClause(conditions map[string][]any) ExtraSchemaPagElementSelectConfigOption {
	return func(s *ExtraSchemaPagElementSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaPagElementTableDummyJoinJoinSQL = `-- O2O join generated from "pag_element_dummy_fkey (inferred)"
left join extra_schema.dummy_join as _pag_element_dummy on _pag_element_dummy.dummy_join_id = pag_element.dummy
`

const extraSchemaPagElementTableDummyJoinSelectSQL = `(case when _pag_element_dummy.dummy_join_id is not null then row(_pag_element_dummy.*) end) as dummy_join_dummy`

const extraSchemaPagElementTableDummyJoinGroupBySQL = `_pag_element_dummy.dummy_join_id,
      _pag_element_dummy.dummy_join_id,
	pag_element.paginated_element_id`

// ExtraSchemaPagElementUpdateParams represents update params for 'extra_schema.pag_element'.
type ExtraSchemaPagElementUpdateParams struct {
	Dummy **ExtraSchemaDummyJoinID `json:"dummy"`                 // dummy
	Name  *string                  `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates extra_schema.pag_element struct fields with the specified params.
func (espe *ExtraSchemaPagElement) SetUpdateParams(params *ExtraSchemaPagElementUpdateParams) {
	if params.Dummy != nil {
		espe.Dummy = *params.Dummy
	}
	if params.Name != nil {
		espe.Name = *params.Name
	}
}

// Insert inserts the ExtraSchemaPagElement to the database.
func (espe *ExtraSchemaPagElement) Insert(ctx context.Context, db DB) (*ExtraSchemaPagElement, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.pag_element (
	dummy, name
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, espe.Dummy, espe.Name)

	rows, err := db.Query(ctx, sqlstr, espe.Dummy, espe.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Insert/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newespe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	*espe = newespe

	return espe, nil
}

// Update updates a ExtraSchemaPagElement in the database.
func (espe *ExtraSchemaPagElement) Update(ctx context.Context, db DB) (*ExtraSchemaPagElement, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.pag_element SET 
	dummy = $1, name = $2 
	WHERE paginated_element_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, espe.CreatedAt, espe.Dummy, espe.Name, espe.PaginatedElementID)

	rows, err := db.Query(ctx, sqlstr, espe.Dummy, espe.Name, espe.PaginatedElementID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Update/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	newespe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	*espe = newespe

	return espe, nil
}

// Upsert upserts a ExtraSchemaPagElement in the database.
// Requires appropriate PK(s) to be set beforehand.
func (espe *ExtraSchemaPagElement) Upsert(ctx context.Context, db DB, params *ExtraSchemaPagElementCreateParams) (*ExtraSchemaPagElement, error) {
	var err error

	espe.Dummy = params.Dummy
	espe.Name = params.Name

	espe, err = espe.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Pag element", Err: err})
			}
			espe, err = espe.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Pag element", Err: err})
			}
		}
	}

	return espe, err
}

// Delete deletes the ExtraSchemaPagElement from the database.
func (espe *ExtraSchemaPagElement) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.pag_element 
	WHERE paginated_element_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, espe.PaginatedElementID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaPagElementPaginatedByCreatedAt returns a cursor-paginated list of ExtraSchemaPagElement.
func ExtraSchemaPagElementPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, direction models.Direction, opts ...ExtraSchemaPagElementSelectConfigOption) ([]ExtraSchemaPagElement, error) {
	c := &ExtraSchemaPagElementSelectConfig{joins: ExtraSchemaPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, extraSchemaPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaPagElementTableDummyJoinGroupBySQL)
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
	 FROM extra_schema.pag_element %s 
	 WHERE pag_element.created_at %s $1
	 %s   %s 
  %s 
  ORDER BY 
		created_at %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaPagElementPaginatedByCreatedAt */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Paginated/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaPagElement/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	return res, nil
}

// ExtraSchemaPagElementByCreatedAt retrieves a row from 'extra_schema.pag_element' as a ExtraSchemaPagElement.
//
// Generated from index 'pag_element_created_at_key'.
func ExtraSchemaPagElementByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...ExtraSchemaPagElementSelectConfigOption) (*ExtraSchemaPagElement, error) {
	c := &ExtraSchemaPagElementSelectConfig{joins: ExtraSchemaPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, extraSchemaPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaPagElementTableDummyJoinGroupBySQL)
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
	 FROM extra_schema.pag_element %s 
	 WHERE pag_element.created_at = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaPagElementByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	espe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &espe, nil
}

// ExtraSchemaPagElementByPaginatedElementID retrieves a row from 'extra_schema.pag_element' as a ExtraSchemaPagElement.
//
// Generated from index 'pag_element_pkey'.
func ExtraSchemaPagElementByPaginatedElementID(ctx context.Context, db DB, paginatedElementID ExtraSchemaPagElementID, opts ...ExtraSchemaPagElementSelectConfigOption) (*ExtraSchemaPagElement, error) {
	c := &ExtraSchemaPagElementSelectConfig{joins: ExtraSchemaPagElementJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaPagElementTableDummyJoinSelectSQL)
		joinClauses = append(joinClauses, extraSchemaPagElementTableDummyJoinJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaPagElementTableDummyJoinGroupBySQL)
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
	 FROM extra_schema.pag_element %s 
	 WHERE pag_element.paginated_element_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaPagElementByPaginatedElementID */\n" + sqlstr

	// run
	// logf(sqlstr, paginatedElementID)
	rows, err := db.Query(ctx, sqlstr, append([]any{paginatedElementID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/db.Query: %w", &XoError{Entity: "Pag element", Err: err}))
	}
	espe, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaPagElement])
	if err != nil {
		return nil, logerror(fmt.Errorf("pag_element/PagElementByPaginatedElementID/pgx.CollectOneRow: %w", &XoError{Entity: "Pag element", Err: err}))
	}

	return &espe, nil
}

// FKDummyJoin_Dummy returns the DummyJoin associated with the ExtraSchemaPagElement's (Dummy).
//
// Generated from foreign key 'pag_element_dummy_fkey'.
func (espe *ExtraSchemaPagElement) FKDummyJoin_Dummy(ctx context.Context, db DB) (*ExtraSchemaDummyJoin, error) {
	return ExtraSchemaDummyJoinByDummyJoinID(ctx, db, *espe.Dummy)
}
