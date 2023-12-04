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

// WorkItem represents a row from 'xo_tests.work_items'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItem struct {
	WorkItemID  WorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Title       *string    `json:"title" db:"title"`                                              // title
	Description *string    `json:"description" db:"description"`                                  // description

	DemoWorkItemJoin          *DemoWorkItem          `json:"-" db:"demo_work_item_work_item_id" openapi-go:"ignore"`            // O2O demo_work_items (inferred)
	WorkItemAssignedUsersJoin *[]User__WIAU_WorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// WorkItemCreateParams represents insert params for 'xo_tests.work_items'.
type WorkItemCreateParams struct {
	Description *string `json:"description"` // description
	Title       *string `json:"title"`       // title
}

type WorkItemID int

// CreateWorkItem creates a new WorkItem in the database with the given params.
func CreateWorkItem(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	wi := &WorkItem{
		Description: params.Description,
		Title:       params.Title,
	}

	return wi.Insert(ctx, db)
}

// WorkItemUpdateParams represents update params for 'xo_tests.work_items'.
type WorkItemUpdateParams struct {
	Description **string `json:"description"` // description
	Title       **string `json:"title"`       // title
}

// SetUpdateParams updates xo_tests.work_items struct fields with the specified params.
func (wi *WorkItem) SetUpdateParams(params *WorkItemUpdateParams) {
	if params.Description != nil {
		wi.Description = *params.Description
	}
	if params.Title != nil {
		wi.Title = *params.Title
	}
}

type WorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemJoins
	filters map[string][]any
}
type WorkItemSelectConfigOption func(*WorkItemSelectConfig)

// WithWorkItemLimit limits row selection.
func WithWorkItemLimit(limit int) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemOrderBy string

type WorkItemJoins struct {
	DemoWorkItem  bool // O2O demo_work_items
	AssignedUsers bool // M2M work_item_assigned_user
}

// WithWorkItemJoin joins with the given tables.
func WithWorkItemJoin(joins WorkItemJoins) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.joins = WorkItemJoins{
			DemoWorkItem:  s.joins.DemoWorkItem || joins.DemoWorkItem,
			AssignedUsers: s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// User__WIAU_WorkItem represents a M2M join against "xo_tests.work_item_assigned_user"
type User__WIAU_WorkItem struct {
	User User             `json:"user" db:"users" required:"true"`
	Role NullWorkItemRole `json:"role" db:"role" required:"true" `
}

// WithWorkItemFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemFilters(filters map[string][]any) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.filters = filters
	}
}

const workItemTableDemoWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const workItemTableDemoWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id`

const workItemTableDemoWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const workItemTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
		work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
		, work_item_assigned_user.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.work_item_assigned_user
	join xo_tests.users on users.user_id = work_item_assigned_user.assigned_user
	group by
		work_item_assigned_user_work_item_id
		, users.user_id
		, role
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_items.work_item_id
`

const workItemTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const workItemTableAssignedUsersGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

// Insert inserts the WorkItem to the database.
func (wi *WorkItem) Insert(ctx context.Context, db DB) (*WorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.work_items (
	description, title
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, wi.Description, wi.Title)

	rows, err := db.Query(ctx, sqlstr, wi.Description, wi.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	*wi = newwi

	return wi, nil
}

// Update updates a WorkItem in the database.
func (wi *WorkItem) Update(ctx context.Context, db DB) (*WorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.work_items SET 
	description = $1, title = $2 
	WHERE work_item_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, wi.Description, wi.Title, wi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, wi.Description, wi.Title, wi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}
	*wi = newwi

	return wi, nil
}

// Upsert upserts a WorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wi *WorkItem) Upsert(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	var err error

	wi.Description = params.Description
	wi.Title = params.Title

	wi, err = wi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item", Err: err})
			}
			wi, err = wi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item", Err: err})
			}
		}
	}

	return wi, err
}

// Delete deletes the WorkItem from the database.
func (wi *WorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemPaginatedByWorkItemID returns a cursor-paginated list of WorkItem.
func WorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, direction Direction, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
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
	if direction == DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.work_item_id %s $1
	 %s   %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, groupbys, direction)
	sqlstr += c.limit
	sqlstr = "/* WorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItems retrieves a row from 'xo_tests.work_items' as a WorkItem.
//
// Generated from index '[xo] base filter query'.
func WorkItems(ctx context.Context, db DB, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 0
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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE true
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItems */\n" + sqlstr

	// run
	// logf(sqlstr, )
	rows, err := db.Query(ctx, sqlstr, append([]any{}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItemByWorkItemID retrieves a row from 'xo_tests.work_items' as a WorkItem.
//
// Generated from index 'work_items_pkey'.
func WorkItemByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...WorkItemSelectConfigOption) (*WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.work_item_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	wi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	return &wi, nil
}

// WorkItemsByTitle retrieves a row from 'xo_tests.work_items' as a WorkItem.
//
// Generated from index 'work_items_title_description_idx1'.
func WorkItemsByTitle(ctx context.Context, db DB, title *string, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.title = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemsByTitle */\n" + sqlstr

	// run
	// logf(sqlstr, title)
	rows, err := db.Query(ctx, sqlstr, append([]any{title}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTitleDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTitleDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}
