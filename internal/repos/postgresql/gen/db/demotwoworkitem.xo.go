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
	orderBy string
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

type DemoTwoWorkItemOrderBy string

const (
	DemoTwoWorkItemCustomDateForProject2DescNullsFirst DemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS FIRST "
	DemoTwoWorkItemCustomDateForProject2DescNullsLast  DemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS LAST "
	DemoTwoWorkItemCustomDateForProject2AscNullsFirst  DemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS FIRST "
	DemoTwoWorkItemCustomDateForProject2AscNullsLast   DemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS LAST "
)

// WithDemoTwoWorkItemOrderBy orders results by the given columns.
func WithDemoTwoWorkItemOrderBy(rows ...DemoTwoWorkItemOrderBy) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
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

type DemoTwoWorkItemJoins struct {
	WorkItem bool // O2O work_items
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
//	// See joins db tag to use the appropriate aliases.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Demo two work item", Err: err})
			}
			dtwi, err = dtwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Demo two work item", Err: err})
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

// DemoTwoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of DemoTwoWorkItem.
func DemoTwoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, direction models.Direction, opts ...DemoTwoWorkItemSelectConfigOption) ([]DemoTwoWorkItem, error) {
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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	demo_two_work_items.custom_date_for_project_2,
	demo_two_work_items.work_item_id %s 
	 FROM public.demo_two_work_items %s 
	 WHERE demo_two_work_items.work_item_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* DemoTwoWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
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
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	demo_two_work_items.custom_date_for_project_2,
	demo_two_work_items.work_item_id %s 
	 FROM public.demo_two_work_items %s 
	 WHERE demo_two_work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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
