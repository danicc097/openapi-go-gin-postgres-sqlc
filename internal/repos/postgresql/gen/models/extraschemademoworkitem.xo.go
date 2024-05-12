// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

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

// ExtraSchemaDemoWorkItem represents a row from 'extra_schema.demo_work_items'.
type ExtraSchemaDemoWorkItem struct {
	WorkItemID ExtraSchemaWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Checked    bool                  `json:"checked" db:"checked" required:"true" nullable:"false"`         // checked

	WorkItemJoin *ExtraSchemaWorkItem `json:"-" db:"work_item_work_item_id"` // O2O work_items (inferred)

}

// ExtraSchemaDemoWorkItemCreateParams represents insert params for 'extra_schema.demo_work_items'.
type ExtraSchemaDemoWorkItemCreateParams struct {
	Checked    bool                  `json:"checked" required:"true" nullable:"false"` // checked
	WorkItemID ExtraSchemaWorkItemID `json:"-" required:"true" nullable:"false"`       // work_item_id
}

// ExtraSchemaDemoWorkItemParams represents common params for both insert and update of 'extra_schema.demo_work_items'.
type ExtraSchemaDemoWorkItemParams interface {
	GetChecked() *bool
}

func (p ExtraSchemaDemoWorkItemCreateParams) GetChecked() *bool {
	x := p.Checked
	return &x
}
func (p ExtraSchemaDemoWorkItemUpdateParams) GetChecked() *bool {
	return p.Checked
}

// CreateExtraSchemaDemoWorkItem creates a new ExtraSchemaDemoWorkItem in the database with the given params.
func CreateExtraSchemaDemoWorkItem(ctx context.Context, db DB, params *ExtraSchemaDemoWorkItemCreateParams) (*ExtraSchemaDemoWorkItem, error) {
	esdwi := &ExtraSchemaDemoWorkItem{
		Checked:    params.Checked,
		WorkItemID: params.WorkItemID,
	}

	return esdwi.Insert(ctx, db)
}

type ExtraSchemaDemoWorkItemSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   ExtraSchemaDemoWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaDemoWorkItemSelectConfigOption func(*ExtraSchemaDemoWorkItemSelectConfig)

// WithExtraSchemaDemoWorkItemLimit limits row selection.
func WithExtraSchemaDemoWorkItemLimit(limit int) ExtraSchemaDemoWorkItemSelectConfigOption {
	return func(s *ExtraSchemaDemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaDemoWorkItemOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaDemoWorkItemOrderBy(rows map[string]*Direction) ExtraSchemaDemoWorkItemSelectConfigOption {
	return func(s *ExtraSchemaDemoWorkItemSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaDemoWorkItem]
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

type ExtraSchemaDemoWorkItemJoins struct {
	WorkItem bool `json:"workItem" required:"true" nullable:"false"` // O2O work_items
}

// WithExtraSchemaDemoWorkItemJoin joins with the given tables.
func WithExtraSchemaDemoWorkItemJoin(joins ExtraSchemaDemoWorkItemJoins) ExtraSchemaDemoWorkItemSelectConfigOption {
	return func(s *ExtraSchemaDemoWorkItemSelectConfig) {
		s.joins = ExtraSchemaDemoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithExtraSchemaDemoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaDemoWorkItemFilters(filters map[string][]any) ExtraSchemaDemoWorkItemSelectConfigOption {
	return func(s *ExtraSchemaDemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaDemoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaDemoWorkItemHavingClause(conditions map[string][]any) ExtraSchemaDemoWorkItemSelectConfigOption {
	return func(s *ExtraSchemaDemoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaDemoWorkItemTableWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey (inferred)"
left join extra_schema.work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = demo_work_items.work_item_id
`

const extraSchemaDemoWorkItemTableWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as work_item_work_item_id`

const extraSchemaDemoWorkItemTableWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
      _demo_work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id`

// ExtraSchemaDemoWorkItemUpdateParams represents update params for 'extra_schema.demo_work_items'.
type ExtraSchemaDemoWorkItemUpdateParams struct {
	Checked *bool `json:"checked" nullable:"false"` // checked
}

// SetUpdateParams updates extra_schema.demo_work_items struct fields with the specified params.
func (esdwi *ExtraSchemaDemoWorkItem) SetUpdateParams(params *ExtraSchemaDemoWorkItemUpdateParams) {
	if params.Checked != nil {
		esdwi.Checked = *params.Checked
	}
}

// Insert inserts the ExtraSchemaDemoWorkItem to the database.
func (esdwi *ExtraSchemaDemoWorkItem) Insert(ctx context.Context, db DB) (*ExtraSchemaDemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.demo_work_items (
	checked, work_item_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, esdwi.Checked, esdwi.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, esdwi.Checked, esdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newesdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*esdwi = newesdwi

	return esdwi, nil
}

// Update updates a ExtraSchemaDemoWorkItem in the database.
func (esdwi *ExtraSchemaDemoWorkItem) Update(ctx context.Context, db DB) (*ExtraSchemaDemoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.demo_work_items SET 
	checked = $1 
	WHERE work_item_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, esdwi.Checked, esdwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, esdwi.Checked, esdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Update/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newesdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*esdwi = newesdwi

	return esdwi, nil
}

// Upsert upserts a ExtraSchemaDemoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esdwi *ExtraSchemaDemoWorkItem) Upsert(ctx context.Context, db DB, params *ExtraSchemaDemoWorkItemCreateParams) (*ExtraSchemaDemoWorkItem, error) {
	var err error

	esdwi.Checked = params.Checked
	esdwi.WorkItemID = params.WorkItemID

	esdwi, err = esdwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaDemoWorkItem/Insert: %w", &XoError{Entity: "Demo work item", Err: err})
			}
			esdwi, err = esdwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaDemoWorkItem/Update: %w", &XoError{Entity: "Demo work item", Err: err})
			}
		}
	}

	return esdwi, err
}

// Delete deletes the ExtraSchemaDemoWorkItem from the database.
func (esdwi *ExtraSchemaDemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.demo_work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esdwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaDemoWorkItemPaginated returns a cursor-paginated list of ExtraSchemaDemoWorkItem.
// At least one cursor is required.
func ExtraSchemaDemoWorkItemPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaDemoWorkItemSelectConfigOption) ([]ExtraSchemaDemoWorkItem, error) {
	c := &ExtraSchemaDemoWorkItemSelectConfig{joins: ExtraSchemaDemoWorkItemJoins{},
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
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaDemoWorkItem][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Paginated/cursor: %w", &XoError{Entity: "Demo work item", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("demo_work_items.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Paginated/orderBy: %w", &XoError{Entity: "Demo work item", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, extraSchemaDemoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaDemoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaDemoWorkItemTableWorkItemGroupBySQL)
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
	demo_work_items.checked,
	demo_work_items.work_item_id %s 
	 FROM extra_schema.demo_work_items %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaDemoWorkItemPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaDemoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	return res, nil
}

// ExtraSchemaDemoWorkItemByWorkItemID retrieves a row from 'extra_schema.demo_work_items' as a ExtraSchemaDemoWorkItem.
//
// Generated from index 'demo_work_items_pkey'.
func ExtraSchemaDemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaDemoWorkItemSelectConfigOption) (*ExtraSchemaDemoWorkItem, error) {
	c := &ExtraSchemaDemoWorkItemSelectConfig{joins: ExtraSchemaDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, extraSchemaDemoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaDemoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaDemoWorkItemTableWorkItemGroupBySQL)
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
	demo_work_items.checked,
	demo_work_items.work_item_id %s 
	 FROM extra_schema.demo_work_items %s 
	 WHERE demo_work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaDemoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	esdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}

	return &esdwi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the ExtraSchemaDemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_work_items_work_item_id_fkey'.
func (esdwi *ExtraSchemaDemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	return ExtraSchemaWorkItemByWorkItemID(ctx, db, esdwi.WorkItemID)
}
