package db

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

// ExtraSchemaWorkItem represents a row from 'extra_schema.work_items'.
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
type ExtraSchemaWorkItem struct {
	WorkItemID  ExtraSchemaWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Title       *string               `json:"title" db:"title"`                                              // title
	Description *string               `json:"description" db:"description"`                                  // description

	DemoWorkItemJoin          *ExtraSchemaDemoWorkItem          `json:"-" db:"demo_work_item_work_item_id" openapi-go:"ignore"`            // O2O demo_work_items (inferred)
	WorkItemAdminsJoin        *[]ExtraSchemaUser                `json:"-" db:"work_item_admin_admins" openapi-go:"ignore"`                 // M2M work_item_admin
	WorkItemAssignedUsersJoin *[]User__WIAU_ExtraSchemaWorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user

}

// ExtraSchemaWorkItemCreateParams represents insert params for 'extra_schema.work_items'.
type ExtraSchemaWorkItemCreateParams struct {
	Description *string `json:"description"` // description
	Title       *string `json:"title"`       // title
}

type ExtraSchemaWorkItemID int

// CreateExtraSchemaWorkItem creates a new ExtraSchemaWorkItem in the database with the given params.
func CreateExtraSchemaWorkItem(ctx context.Context, db DB, params *ExtraSchemaWorkItemCreateParams) (*ExtraSchemaWorkItem, error) {
	eswi := &ExtraSchemaWorkItem{
		Description: params.Description,
		Title:       params.Title,
	}

	return eswi.Insert(ctx, db)
}

type ExtraSchemaWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaWorkItemSelectConfigOption func(*ExtraSchemaWorkItemSelectConfig)

// WithExtraSchemaWorkItemLimit limits row selection.
func WithExtraSchemaWorkItemLimit(limit int) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaWorkItemOrderBy string

const ()

type ExtraSchemaWorkItemJoins struct {
	DemoWorkItem          bool // O2O demo_work_items
	WorkItemAdmins        bool // M2M work_item_admin
	WorkItemAssignedUsers bool // M2M work_item_assigned_user
}

// WithExtraSchemaWorkItemJoin joins with the given tables.
func WithExtraSchemaWorkItemJoin(joins ExtraSchemaWorkItemJoins) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.joins = ExtraSchemaWorkItemJoins{
			DemoWorkItem:          s.joins.DemoWorkItem || joins.DemoWorkItem,
			WorkItemAdmins:        s.joins.WorkItemAdmins || joins.WorkItemAdmins,
			WorkItemAssignedUsers: s.joins.WorkItemAssignedUsers || joins.WorkItemAssignedUsers,
		}
	}
}

// User__WIAU_ExtraSchemaWorkItem represents a M2M join against "extra_schema.work_item_assigned_user"
type User__WIAU_ExtraSchemaWorkItem struct {
	User ExtraSchemaUser          `json:"user" db:"users" required:"true"`
	Role *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithExtraSchemaWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaWorkItemFilters(filters map[string][]any) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaWorkItemHavingClause(conditions map[string][]any) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaWorkItemTableDemoWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join extra_schema.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableDemoWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id`

const extraSchemaWorkItemTableDemoWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const extraSchemaWorkItemTableWorkItemAdminsJoinSQL = `-- M2M join generated from "work_item_admin_admin_fkey"
left join (
	select
		work_item_admin.work_item_id as work_item_admin_work_item_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.work_item_admin
	join extra_schema.users on users.user_id = work_item_admin.admin
	group by
		work_item_admin_work_item_id
		, users.user_id
) as xo_join_work_item_admin_admins on xo_join_work_item_admin_admins.work_item_admin_work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableWorkItemAdminsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_admin_admins.__users
		)) filter (where xo_join_work_item_admin_admins.__users_user_id is not null), '{}') as work_item_admin_admins`

const extraSchemaWorkItemTableWorkItemAdminsGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

const extraSchemaWorkItemTableWorkItemAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
		work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
		, work_item_assigned_user.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.work_item_assigned_user
	join extra_schema.users on users.user_id = work_item_assigned_user.assigned_user
	group by
		work_item_assigned_user_work_item_id
		, users.user_id
		, role
) as xo_join_work_item_assigned_user_assigned_users on xo_join_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableWorkItemAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assigned_user_assigned_users.__users
		, xo_join_work_item_assigned_user_assigned_users.role
		)) filter (where xo_join_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const extraSchemaWorkItemTableWorkItemAssignedUsersGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

// ExtraSchemaWorkItemUpdateParams represents update params for 'extra_schema.work_items'.
type ExtraSchemaWorkItemUpdateParams struct {
	Description **string `json:"description"` // description
	Title       **string `json:"title"`       // title
}

// SetUpdateParams updates extra_schema.work_items struct fields with the specified params.
func (eswi *ExtraSchemaWorkItem) SetUpdateParams(params *ExtraSchemaWorkItemUpdateParams) {
	if params.Description != nil {
		eswi.Description = *params.Description
	}
	if params.Title != nil {
		eswi.Title = *params.Title
	}
}

// Insert inserts the ExtraSchemaWorkItem to the database.
func (eswi *ExtraSchemaWorkItem) Insert(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.work_items (
	description, title
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, eswi.Description, eswi.Title)

	rows, err := db.Query(ctx, sqlstr, eswi.Description, eswi.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Insert/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	neweswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	*eswi = neweswi

	return eswi, nil
}

// Update updates a ExtraSchemaWorkItem in the database.
func (eswi *ExtraSchemaWorkItem) Update(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.work_items SET 
	description = $1, title = $2 
	WHERE work_item_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, eswi.Description, eswi.Title, eswi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, eswi.Description, eswi.Title, eswi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Update/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	neweswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}
	*eswi = neweswi

	return eswi, nil
}

// Upsert upserts a ExtraSchemaWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (eswi *ExtraSchemaWorkItem) Upsert(ctx context.Context, db DB, params *ExtraSchemaWorkItemCreateParams) (*ExtraSchemaWorkItem, error) {
	var err error

	eswi.Description = params.Description
	eswi.Title = params.Title

	eswi, err = eswi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item", Err: err})
			}
			eswi, err = eswi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item", Err: err})
			}
		}
	}

	return eswi, err
}

// Delete deletes the ExtraSchemaWorkItem from the database.
func (eswi *ExtraSchemaWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, eswi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaWorkItemPaginatedByWorkItemID returns a cursor-paginated list of ExtraSchemaWorkItem.
func ExtraSchemaWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, direction models.Direction, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAdminsGroupBySQL)
	}

	if c.joins.WorkItemAssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAssignedUsersGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM extra_schema.work_items %s 
	 WHERE work_items.work_item_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItems retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index '[xo] base filter query'.
func ExtraSchemaWorkItems(ctx context.Context, db DB, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAdminsGroupBySQL)
	}

	if c.joins.WorkItemAssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAssignedUsersGroupBySQL)
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
	 FROM extra_schema.work_items %s 
	 WHERE true
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItems */\n" + sqlstr

	// run
	// logf(sqlstr, )
	rows, err := db.Query(ctx, sqlstr, append([]any{}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemByWorkItemID retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index 'work_items_pkey'.
func ExtraSchemaWorkItemByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemSelectConfigOption) (*ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAdminsGroupBySQL)
	}

	if c.joins.WorkItemAssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAssignedUsersGroupBySQL)
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
	 FROM extra_schema.work_items %s 
	 WHERE work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	eswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	return &eswi, nil
}

// ExtraSchemaWorkItemsByTitle retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index 'work_items_title_description_idx1'.
func ExtraSchemaWorkItemsByTitle(ctx context.Context, db DB, title *string, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAdminsGroupBySQL)
	}

	if c.joins.WorkItemAssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemTableWorkItemAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemTableWorkItemAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemTableWorkItemAssignedUsersGroupBySQL)
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
	 FROM extra_schema.work_items %s 
	 WHERE work_items.title = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemsByTitle */\n" + sqlstr

	// run
	// logf(sqlstr, title)
	rows, err := db.Query(ctx, sqlstr, append([]any{title}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByTitleDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByTitleDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}
