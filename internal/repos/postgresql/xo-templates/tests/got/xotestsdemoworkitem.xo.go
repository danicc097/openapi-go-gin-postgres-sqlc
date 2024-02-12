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

// XoTestsDemoWorkItem represents a row from 'xo_tests.demo_work_items'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsDemoWorkItem struct {
	WorkItemID XoTestsWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Checked    bool              `json:"checked" db:"checked" required:"true" nullable:"false"`         // checked

	WorkItemJoin              *XoTestsWorkItem                  `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"`                 // O2O work_items (inferred)
	WorkItemJoinWII           *XoTestsWorkItem                  `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"`                 // O2O work_items (inferred)
	WorkItemAssignedUsersJoin *[]User__WIAU_XoTestsDemoWorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// XoTestsDemoWorkItemCreateParams represents insert params for 'xo_tests.demo_work_items'.
type XoTestsDemoWorkItemCreateParams struct {
	Checked    bool              `json:"checked" required:"true" nullable:"false"` // checked
	WorkItemID XoTestsWorkItemID `json:"-" required:"true" nullable:"false"`       // work_item_id
}

// CreateXoTestsDemoWorkItem creates a new XoTestsDemoWorkItem in the database with the given params.
func CreateXoTestsDemoWorkItem(ctx context.Context, db DB, params *XoTestsDemoWorkItemCreateParams) (*XoTestsDemoWorkItem, error) {
	xtdwi := &XoTestsDemoWorkItem{
		Checked:    params.Checked,
		WorkItemID: params.WorkItemID,
	}

	return xtdwi.Insert(ctx, db)
}

type XoTestsDemoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsDemoWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsDemoWorkItemSelectConfigOption func(*XoTestsDemoWorkItemSelectConfig)

// WithXoTestsDemoWorkItemLimit limits row selection.
func WithXoTestsDemoWorkItemLimit(limit int) XoTestsDemoWorkItemSelectConfigOption {
	return func(s *XoTestsDemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsDemoWorkItemOrderBy string

type XoTestsDemoWorkItemJoins struct {
	WorkItem          bool // O2O work_items
	WorkItemWorkItems bool // O2O work_items
	AssignedUsers     bool // M2M work_item_assigned_user
}

// WithXoTestsDemoWorkItemJoin joins with the given tables.
func WithXoTestsDemoWorkItemJoin(joins XoTestsDemoWorkItemJoins) XoTestsDemoWorkItemSelectConfigOption {
	return func(s *XoTestsDemoWorkItemSelectConfig) {
		s.joins = XoTestsDemoWorkItemJoins{
			WorkItem:          s.joins.WorkItem || joins.WorkItem,
			WorkItemWorkItems: s.joins.WorkItemWorkItems || joins.WorkItemWorkItems,
			AssignedUsers:     s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// User__WIAU_XoTestsDemoWorkItem represents a M2M join against "xo_tests.work_item_assigned_user"
type User__WIAU_XoTestsDemoWorkItem struct {
	User XoTestsUser          `json:"user" db:"users" required:"true"`
	Role *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithXoTestsDemoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsDemoWorkItemFilters(filters map[string][]any) XoTestsDemoWorkItemSelectConfigOption {
	return func(s *XoTestsDemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsDemoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsDemoWorkItemHavingClause(conditions map[string][]any) XoTestsDemoWorkItemSelectConfigOption {
	return func(s *XoTestsDemoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const xoTestsDemoWorkItemTableWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey (inferred)"
left join xo_tests.work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = demo_work_items.work_item_id
`

const xoTestsDemoWorkItemTableWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as work_item_work_item_id`

const xoTestsDemoWorkItemTableWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
      _demo_work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id`

const xoTestsDemoWorkItemTableWorkItemWorkItemsJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey (inferred)-shared-ref-demo_work_items"
left join xo_tests.work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = demo_work_items.work_item_id
`

const xoTestsDemoWorkItemTableWorkItemWorkItemsSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as work_item_work_item_id`

const xoTestsDemoWorkItemTableWorkItemWorkItemsGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
      _demo_work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id`

const xoTestsDemoWorkItemTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey-shared-ref-demo_work_items"
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
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = demo_work_items.work_item_id
`

const xoTestsDemoWorkItemTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const xoTestsDemoWorkItemTableAssignedUsersGroupBySQL = `demo_work_items.work_item_id, demo_work_items.work_item_id`

// XoTestsDemoWorkItemUpdateParams represents update params for 'xo_tests.demo_work_items'.
type XoTestsDemoWorkItemUpdateParams struct {
	Checked *bool `json:"checked" nullable:"false"` // checked
}

// SetUpdateParams updates xo_tests.demo_work_items struct fields with the specified params.
func (xtdwi *XoTestsDemoWorkItem) SetUpdateParams(params *XoTestsDemoWorkItemUpdateParams) {
	if params.Checked != nil {
		xtdwi.Checked = *params.Checked
	}
}

// Insert inserts the XoTestsDemoWorkItem to the database.
func (xtdwi *XoTestsDemoWorkItem) Insert(ctx context.Context, db DB) (*XoTestsDemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.demo_work_items (
	checked, work_item_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtdwi.Checked, xtdwi.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, xtdwi.Checked, xtdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newxtdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*xtdwi = newxtdwi

	return xtdwi, nil
}

// Update updates a XoTestsDemoWorkItem in the database.
func (xtdwi *XoTestsDemoWorkItem) Update(ctx context.Context, db DB) (*XoTestsDemoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.demo_work_items SET 
	checked = $1 
	WHERE work_item_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, xtdwi.Checked, xtdwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, xtdwi.Checked, xtdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Update/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	newxtdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	*xtdwi = newxtdwi

	return xtdwi, nil
}

// Upsert upserts a XoTestsDemoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtdwi *XoTestsDemoWorkItem) Upsert(ctx context.Context, db DB, params *XoTestsDemoWorkItemCreateParams) (*XoTestsDemoWorkItem, error) {
	var err error

	xtdwi.Checked = params.Checked
	xtdwi.WorkItemID = params.WorkItemID

	xtdwi, err = xtdwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Demo work item", Err: err})
			}
			xtdwi, err = xtdwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Demo work item", Err: err})
			}
		}
	}

	return xtdwi, err
}

// Delete deletes the XoTestsDemoWorkItem from the database.
func (xtdwi *XoTestsDemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.demo_work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtdwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsDemoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of XoTestsDemoWorkItem.
func XoTestsDemoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, direction models.Direction, opts ...XoTestsDemoWorkItemSelectConfigOption) ([]XoTestsDemoWorkItem, error) {
	c := &XoTestsDemoWorkItemSelectConfig{joins: XoTestsDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableWorkItemGroupBySQL)
	}

	if c.joins.WorkItemWorkItems {
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableAssignedUsersGroupBySQL)
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
	demo_work_items.checked,
	demo_work_items.work_item_id %s 
	 FROM xo_tests.demo_work_items %s 
	 WHERE demo_work_items.work_item_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsDemoWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsDemoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	return res, nil
}

// XoTestsDemoWorkItemByWorkItemID retrieves a row from 'xo_tests.demo_work_items' as a XoTestsDemoWorkItem.
//
// Generated from index 'demo_work_items_pkey'.
func XoTestsDemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, opts ...XoTestsDemoWorkItemSelectConfigOption) (*XoTestsDemoWorkItem, error) {
	c := &XoTestsDemoWorkItemSelectConfig{joins: XoTestsDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableWorkItemGroupBySQL)
	}

	if c.joins.WorkItemWorkItems {
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableWorkItemWorkItemsGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsDemoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsDemoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsDemoWorkItemTableAssignedUsersGroupBySQL)
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
	demo_work_items.checked,
	demo_work_items.work_item_id %s 
	 FROM xo_tests.demo_work_items %s 
	 WHERE demo_work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsDemoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	xtdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}

	return &xtdwi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the XoTestsDemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_work_items_work_item_id_fkey'.
func (xtdwi *XoTestsDemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	return XoTestsWorkItemByWorkItemID(ctx, db, xtdwi.WorkItemID)
}
