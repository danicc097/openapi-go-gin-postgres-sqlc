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
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItem struct {
	WorkItemID int64   `json:"workItemID" db:"work_item_id" required:"true"` // work_item_id
	Title      *string `json:"title" db:"title" required:"true"`             // title

	DemoWorkItemJoin          *DemoWorkItem          `json:"-" db:"demo_work_item_work_item_id" openapi-go:"ignore"`            // O2O demo_work_items (inferred)
	WorkItemAssignedUsersJoin *[]User__WIAU_WorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// WorkItemCreateParams represents insert params for 'xo_tests.work_items'.
type WorkItemCreateParams struct {
	Title *string `json:"title" required:"true"` // title
}

// CreateWorkItem creates a new WorkItem in the database with the given params.
func CreateWorkItem(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	wi := &WorkItem{
		Title: params.Title,
	}

	return wi.Insert(ctx, db)
}

// WorkItemUpdateParams represents update params for 'xo_tests.work_items'.
type WorkItemUpdateParams struct {
	Title **string `json:"title" required:"true"` // title
}

// SetUpdateParams updates xo_tests.work_items struct fields with the specified params.
func (wi *WorkItem) SetUpdateParams(params *WorkItemUpdateParams) {
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

// WithWorkItemFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
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

// Insert inserts the WorkItem to the database.
func (wi *WorkItem) Insert(ctx context.Context, db DB) (*WorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.work_items (` +
		`title` +
		`) VALUES (` +
		`$1` +
		`) RETURNING * `
	// run
	logf(sqlstr, wi.Title)

	rows, err := db.Query(ctx, sqlstr, wi.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/db.Query: %w", err))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/pgx.CollectOneRow: %w", err))
	}

	*wi = newwi

	return wi, nil
}

// Update updates a WorkItem in the database.
func (wi *WorkItem) Update(ctx context.Context, db DB) (*WorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.work_items SET ` +
		`title = $1 ` +
		`WHERE work_item_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wi.Title, wi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, wi.Title, wi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/db.Query: %w", err))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/pgx.CollectOneRow: %w", err))
	}
	*wi = newwi

	return wi, nil
}

// Upsert upserts a WorkItem in the database.
// Requires appropiate PK(s) to be set beforehand.
func (wi *WorkItem) Upsert(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	var err error

	wi.Title = params.Title

	wi, err = wi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			wi, err = wi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return wi, err
}

// Delete deletes the WorkItem from the database.
func (wi *WorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemPaginatedByWorkItemIDAsc returns a cursor-paginated list of WorkItem in Asc order.
func WorkItemPaginatedByWorkItemIDAsc(ctx context.Context, db DB, workItemID int64, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_items.work_item_id,
work_items.title,
(case when $1::boolean = true and _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') end) as work_item_assigned_user_assigned_users `+
		`FROM xo_tests.work_items `+
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
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
`+
		` WHERE work_items.work_item_id > $3`+
		` %s  GROUP BY work_items.work_item_id,
work_items.title,
_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id,
work_items.work_item_id, work_items.work_item_id ORDER BY
		work_item_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.DemoWorkItem, c.joins.AssignedUsers, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemPaginatedByWorkItemIDDesc returns a cursor-paginated list of WorkItem in Desc order.
func WorkItemPaginatedByWorkItemIDDesc(ctx context.Context, db DB, workItemID int64, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_items.work_item_id,
work_items.title,
(case when $1::boolean = true and _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') end) as work_item_assigned_user_assigned_users `+
		`FROM xo_tests.work_items `+
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
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
`+
		` WHERE work_items.work_item_id < $3`+
		` %s  GROUP BY work_items.work_item_id,
work_items.title,
_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id,
work_items.work_item_id, work_items.work_item_id ORDER BY
		work_item_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.DemoWorkItem, c.joins.AssignedUsers, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemByWorkItemID retrieves a row from 'xo_tests.work_items' as a WorkItem.
//
// Generated from index 'work_items_pkey'.
func WorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemSelectConfigOption) (*WorkItem, error) {
	c := &WorkItemSelectConfig{joins: WorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_items.work_item_id,
work_items.title,
(case when $1::boolean = true and _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') end) as work_item_assigned_user_assigned_users `+
		`FROM xo_tests.work_items `+
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
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
`+
		` WHERE work_items.work_item_id = $3`+
		` %s  GROUP BY
_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id,
work_items.work_item_id, work_items.work_item_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.DemoWorkItem, c.joins.AssignedUsers, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", err))
	}
	wi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", err))
	}

	return &wi, nil
}
