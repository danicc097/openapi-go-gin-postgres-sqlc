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

// DemoTwoWorkItem represents a row from 'public.demo_two_work_items'.
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
type DemoTwoWorkItem struct {
	WorkItemID            WorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	CustomDateForProject2 *time.Time `json:"customDateForProject2" db:"custom_date_for_project_2"`          // custom_date_for_project_2

	WorkItemJoin *WorkItem `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"` // O2O work_items (inferred)

}

// DemoTwoWorkItemCreateParams represents insert params for 'public.demo_two_work_items'.
type DemoTwoWorkItemCreateParams struct {
	CustomDateForProject2 *time.Time `json:"customDateForProject2"`              // custom_date_for_project_2
	WorkItemID            WorkItemID `json:"-" required:"true" nullable:"false"` // work_item_id
}

// DemoTwoWorkItemParams represents common params for both insert and update of 'public.demo_two_work_items'.
type DemoTwoWorkItemParams interface {
	GetCustomDateForProject2() *time.Time
}

func (p DemoTwoWorkItemCreateParams) GetCustomDateForProject2() *time.Time {
	return p.CustomDateForProject2
}
func (p DemoTwoWorkItemUpdateParams) GetCustomDateForProject2() *time.Time {
	if p.CustomDateForProject2 != nil {
		return *p.CustomDateForProject2
	}
	return nil
}

// CreateDemoTwoWorkItem creates a new DemoTwoWorkItem in the database with the given params.
func CreateDemoTwoWorkItem(ctx context.Context, db DB, params *DemoTwoWorkItemCreateParams) (*DemoTwoWorkItem, error) {
	dtwi := &DemoTwoWorkItem{
		CustomDateForProject2: params.CustomDateForProject2,
		WorkItemID:            params.WorkItemID,
	}

	return dtwi.Insert(ctx, db)
}

type DemoTwoWorkItemSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   DemoTwoWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type DemoTwoWorkItemSelectConfigOption func(*DemoTwoWorkItemSelectConfig)

// WithDemoTwoWorkItemLimit limits row selection.
func WithDemoTwoWorkItemLimit(limit int) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDemoTwoWorkItemOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithDemoTwoWorkItemOrderBy(rows map[string]*models.Direction) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		te := EntityFields[TableEntityDemoTwoWorkItem]
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

type DemoTwoWorkItemJoins struct {
	WorkItem bool `json:"workItem" required:"true" nullable:"false"` // O2O work_items
}

// WithDemoTwoWorkItemJoin joins with the given tables.
func WithDemoTwoWorkItemJoin(joins DemoTwoWorkItemJoins) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		s.joins = DemoTwoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithDemoTwoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithDemoTwoWorkItemFilters(filters map[string][]any) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithDemoTwoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithDemoTwoWorkItemHavingClause(conditions map[string][]any) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const demoTwoWorkItemTableWorkItemJoinSQL = `-- O2O join generated from "demo_two_work_items_work_item_id_fkey (inferred)"
left join work_items as _demo_two_work_items_work_item_id on _demo_two_work_items_work_item_id.work_item_id = demo_two_work_items.work_item_id
`

const demoTwoWorkItemTableWorkItemSelectSQL = `(case when _demo_two_work_items_work_item_id.work_item_id is not null then row(_demo_two_work_items_work_item_id.*) end) as work_item_work_item_id`

const demoTwoWorkItemTableWorkItemGroupBySQL = `_demo_two_work_items_work_item_id.work_item_id,
      _demo_two_work_items_work_item_id.work_item_id,
	demo_two_work_items.work_item_id`

// DemoTwoWorkItemUpdateParams represents update params for 'public.demo_two_work_items'.
type DemoTwoWorkItemUpdateParams struct {
	CustomDateForProject2 **time.Time `json:"customDateForProject2"` // custom_date_for_project_2
}

// SetUpdateParams updates public.demo_two_work_items struct fields with the specified params.
func (dtwi *DemoTwoWorkItem) SetUpdateParams(params *DemoTwoWorkItemUpdateParams) {
	if params.CustomDateForProject2 != nil {
		dtwi.CustomDateForProject2 = *params.CustomDateForProject2
	}
}

// Insert inserts the DemoTwoWorkItem to the database.
func (dtwi *DemoTwoWorkItem) Insert(ctx context.Context, db DB) (*DemoTwoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.demo_two_work_items (
	custom_date_for_project_2, work_item_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	newdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	*dtwi = newdtwi

	return dtwi, nil
}

// Update updates a DemoTwoWorkItem in the database.
func (dtwi *DemoTwoWorkItem) Update(ctx context.Context, db DB) (*DemoTwoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.demo_two_work_items SET 
	custom_date_for_project_2 = $1 
	WHERE work_item_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Update/db.Query: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	newdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	*dtwi = newdtwi

	return dtwi, nil
}

// Upsert upserts a DemoTwoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (dtwi *DemoTwoWorkItem) Upsert(ctx context.Context, db DB, params *DemoTwoWorkItemCreateParams) (*DemoTwoWorkItem, error) {
	var err error

	dtwi.CustomDateForProject2 = params.CustomDateForProject2
	dtwi.WorkItemID = params.WorkItemID

	dtwi, err = dtwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertDemoTwoWorkItem/Insert: %w", &XoError{Entity: "Demo two work item", Err: err})
			}
			dtwi, err = dtwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertDemoTwoWorkItem/Update: %w", &XoError{Entity: "Demo two work item", Err: err})
			}
		}
	}

	return dtwi, err
}

// Delete deletes the DemoTwoWorkItem from the database.
func (dtwi *DemoTwoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_two_work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dtwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// DemoTwoWorkItemPaginated returns a cursor-paginated list of DemoTwoWorkItem.
// At least one cursor is required.
func DemoTwoWorkItemPaginated(ctx context.Context, db DB, cursors models.PaginationCursors, opts ...DemoTwoWorkItemSelectConfigOption) ([]DemoTwoWorkItem, error) {
	c := &DemoTwoWorkItemSelectConfig{joins: DemoTwoWorkItemJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		if cursor.Value == nil {

			return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
		}
		field, ok := EntityFields[TableEntityDemoTwoWorkItem][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/cursor: %w", &XoError{Entity: "Demo two work item", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("demo_two_work_items.%s %s $i", field.Db, op)] = []any{*cursor.Value}
		c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts
	}

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
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/orderBy: %w", &XoError{Entity: "Demo two work item", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, demoTwoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, demoTwoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, demoTwoWorkItemTableWorkItemGroupBySQL)
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
	demo_two_work_items.custom_date_for_project_2,
	demo_two_work_items.work_item_id %s 
	 FROM public.demo_two_work_items %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* DemoTwoWorkItemPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	return res, nil
}

// DemoTwoWorkItemByWorkItemID retrieves a row from 'public.demo_two_work_items' as a DemoTwoWorkItem.
//
// Generated from index 'demo_two_work_items_pkey'.
func DemoTwoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...DemoTwoWorkItemSelectConfigOption) (*DemoTwoWorkItem, error) {
	c := &DemoTwoWorkItemSelectConfig{joins: DemoTwoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, demoTwoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, demoTwoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, demoTwoWorkItemTableWorkItemGroupBySQL)
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
	demo_two_work_items.custom_date_for_project_2,
	demo_two_work_items.work_item_id %s 
	 FROM public.demo_two_work_items %s 
	 WHERE demo_two_work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* DemoTwoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_two_work_items/DemoTwoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}
	dtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_two_work_items/DemoTwoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Demo two work item", Err: err}))
	}

	return &dtwi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoTwoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_two_work_items_work_item_id_fkey'.
func (dtwi *DemoTwoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dtwi.WorkItemID)
}
