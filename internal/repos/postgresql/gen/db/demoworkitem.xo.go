// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

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

// DemoWorkItem represents a row from 'public.demo_work_items'.
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
type DemoWorkItem struct {
	WorkItemID    WorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`       // work_item_id
	Ref           string     `json:"ref" db:"ref" required:"true" nullable:"false" pattern:"^[0-9]{8}$"`  // ref
	Line          string     `json:"line" db:"line" required:"true" nullable:"false"`                     // line
	LastMessageAt time.Time  `json:"lastMessageAt" db:"last_message_at" required:"true" nullable:"false"` // last_message_at
	Reopened      bool       `json:"reopened" db:"reopened" required:"true" nullable:"false"`             // reopened

	WorkItemJoin *WorkItem `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"` // O2O work_items (inferred)

}

// DemoWorkItemCreateParams represents insert params for 'public.demo_work_items'.
type DemoWorkItemCreateParams struct {
	LastMessageAt time.Time  `json:"lastMessageAt" required:"true" nullable:"false"`            // last_message_at
	Line          string     `json:"line" required:"true" nullable:"false"`                     // line
	Ref           string     `json:"ref" required:"true" nullable:"false" pattern:"^[0-9]{8}$"` // ref
	Reopened      bool       `json:"reopened" required:"true" nullable:"false"`                 // reopened
	WorkItemID    WorkItemID `json:"-" required:"true" nullable:"false"`                        // work_item_id
}

// DemoWorkItemParams represents common params for both insert and update of 'public.demo_work_items'.
type DemoWorkItemParams interface {
	GetLastMessageAt() *time.Time
	GetLine() *string
	GetRef() *string
	GetReopened() *bool
}

func (p DemoWorkItemCreateParams) GetLastMessageAt() *time.Time {
	x := p.LastMessageAt
	return &x
}
func (p DemoWorkItemUpdateParams) GetLastMessageAt() *time.Time {
	return p.LastMessageAt
}

func (p DemoWorkItemCreateParams) GetLine() *string {
	x := p.Line
	return &x
}
func (p DemoWorkItemUpdateParams) GetLine() *string {
	return p.Line
}

func (p DemoWorkItemCreateParams) GetRef() *string {
	x := p.Ref
	return &x
}
func (p DemoWorkItemUpdateParams) GetRef() *string {
	return p.Ref
}

func (p DemoWorkItemCreateParams) GetReopened() *bool {
	x := p.Reopened
	return &x
}
func (p DemoWorkItemUpdateParams) GetReopened() *bool {
	return p.Reopened
}

// CreateDemoWorkItem creates a new DemoWorkItem in the database with the given params.
func CreateDemoWorkItem(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	dwi := &DemoWorkItem{
		LastMessageAt: params.LastMessageAt,
		Line:          params.Line,
		Ref:           params.Ref,
		Reopened:      params.Reopened,
		WorkItemID:    params.WorkItemID,
	}

	return dwi.Insert(ctx, db)
}

type DemoWorkItemSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   DemoWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type DemoWorkItemSelectConfigOption func(*DemoWorkItemSelectConfig)

// WithDemoWorkItemLimit limits row selection.
func WithDemoWorkItemLimit(limit int) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDemoWorkItemOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithDemoWorkItemOrderBy(rows map[string]*models.Direction) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		te := EntityFields[TableEntityDemoWorkItem]
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

type DemoWorkItemJoins struct {
	WorkItem bool `json:"workItem" required:"true" nullable:"false"` // O2O work_items
}

// WithDemoWorkItemJoin joins with the given tables.
func WithDemoWorkItemJoin(joins DemoWorkItemJoins) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.joins = DemoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithDemoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithDemoWorkItemFilters(filters map[string][]any) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithDemoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithDemoWorkItemHavingClause(conditions map[string][]any) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const demoWorkItemTableWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey (inferred)"
left join work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = demo_work_items.work_item_id
`

const demoWorkItemTableWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as work_item_work_item_id`

const demoWorkItemTableWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
      _demo_work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id`

// DemoWorkItemUpdateParams represents update params for 'public.demo_work_items'.
type DemoWorkItemUpdateParams struct {
	LastMessageAt *time.Time `json:"lastMessageAt" nullable:"false"`            // last_message_at
	Line          *string    `json:"line" nullable:"false"`                     // line
	Ref           *string    `json:"ref" nullable:"false" pattern:"^[0-9]{8}$"` // ref
	Reopened      *bool      `json:"reopened" nullable:"false"`                 // reopened
}

// SetUpdateParams updates public.demo_work_items struct fields with the specified params.
func (dwi *DemoWorkItem) SetUpdateParams(params *DemoWorkItemUpdateParams) {
	if params.LastMessageAt != nil {
		dwi.LastMessageAt = *params.LastMessageAt
	}
	if params.Line != nil {
		dwi.Line = *params.Line
	}
	if params.Ref != nil {
		dwi.Ref = *params.Ref
	}
	if params.Reopened != nil {
		dwi.Reopened = *params.Reopened
	}
}

// Insert inserts the DemoWorkItem to the database.
func (dwi *DemoWorkItem) Insert(ctx context.Context, db DB) (*DemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.demo_work_items (
	last_message_at, line, ref, reopened, work_item_id
	) VALUES (
	$1, $2, $3, $4, $5
	)
	 RETURNING * `
	// run
	logf(sqlstr, dwi.LastMessageAt, dwi.Line, dwi.Ref, dwi.Reopened, dwi.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, dwi.LastMessageAt, dwi.Line, dwi.Ref, dwi.Reopened, dwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*dwi = newdwi

	return dwi, nil
}

// Update updates a DemoWorkItem in the database.
func (dwi *DemoWorkItem) Update(ctx context.Context, db DB) (*DemoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.demo_work_items SET 
	last_message_at = $1, line = $2, ref = $3, reopened = $4 
	WHERE work_item_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, dwi.LastMessageAt, dwi.Line, dwi.Ref, dwi.Reopened, dwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dwi.LastMessageAt, dwi.Line, dwi.Ref, dwi.Reopened, dwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Update/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*dwi = newdwi

	return dwi, nil
}

// Upsert upserts a DemoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (dwi *DemoWorkItem) Upsert(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	var err error

	dwi.LastMessageAt = params.LastMessageAt
	dwi.Line = params.Line
	dwi.Ref = params.Ref
	dwi.Reopened = params.Reopened
	dwi.WorkItemID = params.WorkItemID

	dwi, err = dwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertDemoWorkItem/Insert: %w", &XoError{Entity: "Demo work item", Err: err})
			}
			dwi, err = dwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertDemoWorkItem/Update: %w", &XoError{Entity: "Demo work item", Err: err})
			}
		}
	}

	return dwi, err
}

// Delete deletes the DemoWorkItem from the database.
func (dwi *DemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// DemoWorkItemPaginated returns a cursor-paginated list of DemoWorkItem.
// At least one cursor is required.
func DemoWorkItemPaginated(ctx context.Context, db DB, cursor models.PaginationCursor, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{},
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
	field, ok := EntityFields[TableEntityDemoWorkItem][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/cursor: %w", &XoError{Entity: "Demo work item", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == models.DirectionAsc {
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
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/orderBy: %w", &XoError{Entity: "Demo work item", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, demoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, demoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, demoWorkItemTableWorkItemGroupBySQL)
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
	demo_work_items.last_message_at,
	demo_work_items.line,
	demo_work_items.ref,
	demo_work_items.reopened,
	demo_work_items.work_item_id %s 
	 FROM public.demo_work_items %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* DemoWorkItemPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	return res, nil
}

// DemoWorkItemByWorkItemID retrieves a row from 'public.demo_work_items' as a DemoWorkItem.
//
// Generated from index 'demo_work_items_pkey'.
func DemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...DemoWorkItemSelectConfigOption) (*DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, demoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, demoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, demoWorkItemTableWorkItemGroupBySQL)
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
	demo_work_items.last_message_at,
	demo_work_items.line,
	demo_work_items.ref,
	demo_work_items.reopened,
	demo_work_items.work_item_id %s 
	 FROM public.demo_work_items %s 
	 WHERE demo_work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* DemoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	dwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}

	return &dwi, nil
}

// DemoWorkItemsByRefLine retrieves a row from 'public.demo_work_items' as a DemoWorkItem.
//
// Generated from index 'demo_work_items_ref_line_idx'.
func DemoWorkItemsByRefLine(ctx context.Context, db DB, ref string, line string, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
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
		selectClauses = append(selectClauses, demoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, demoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, demoWorkItemTableWorkItemGroupBySQL)
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
	demo_work_items.last_message_at,
	demo_work_items.line,
	demo_work_items.ref,
	demo_work_items.reopened,
	demo_work_items.work_item_id %s 
	 FROM public.demo_work_items %s 
	 WHERE demo_work_items.ref = $1 AND demo_work_items.line = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* DemoWorkItemsByRefLine */\n" + sqlstr

	// run
	// logf(sqlstr, ref, line)
	rows, err := db.Query(ctx, sqlstr, append([]any{ref, line}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/DemoWorkItemsByRefLine/Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/DemoWorkItemsByRefLine/pgx.CollectRows: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	return res, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_work_items_work_item_id_fkey'.
func (dwi *DemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dwi.WorkItemID)
}
