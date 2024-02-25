package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// ExtraSchemaWorkItemAdmin represents a row from 'extra_schema.work_item_admin'.
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
type ExtraSchemaWorkItemAdmin struct {
	WorkItemID ExtraSchemaWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Admin      ExtraSchemaUserID     `json:"admin" db:"admin" required:"true" nullable:"false"`             // admin

	AdminWorkItemsJoin *[]ExtraSchemaWorkItem `json:"-" db:"work_item_admin_work_items" openapi-go:"ignore"` // M2M work_item_admin
	WorkItemAdminsJoin *[]ExtraSchemaUser     `json:"-" db:"work_item_admin_admins" openapi-go:"ignore"`     // M2M work_item_admin

}

// ExtraSchemaWorkItemAdminCreateParams represents insert params for 'extra_schema.work_item_admin'.
type ExtraSchemaWorkItemAdminCreateParams struct {
	Admin      ExtraSchemaUserID     `json:"admin" required:"true" nullable:"false"`      // admin
	WorkItemID ExtraSchemaWorkItemID `json:"workItemID" required:"true" nullable:"false"` // work_item_id
}

// CreateExtraSchemaWorkItemAdmin creates a new ExtraSchemaWorkItemAdmin in the database with the given params.
func CreateExtraSchemaWorkItemAdmin(ctx context.Context, db DB, params *ExtraSchemaWorkItemAdminCreateParams) (*ExtraSchemaWorkItemAdmin, error) {
	eswia := &ExtraSchemaWorkItemAdmin{
		Admin:      params.Admin,
		WorkItemID: params.WorkItemID,
	}

	return eswia.Insert(ctx, db)
}

type ExtraSchemaWorkItemAdminSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaWorkItemAdminJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaWorkItemAdminSelectConfigOption func(*ExtraSchemaWorkItemAdminSelectConfig)

// WithExtraSchemaWorkItemAdminLimit limits row selection.
func WithExtraSchemaWorkItemAdminLimit(limit int) ExtraSchemaWorkItemAdminSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAdminSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaWorkItemAdminOrderBy string

const ()

type ExtraSchemaWorkItemAdminJoins struct {
	AdminWorkItems bool // M2M work_item_admin
	WorkItemAdmins bool // M2M work_item_admin
}

// WithExtraSchemaWorkItemAdminJoin joins with the given tables.
func WithExtraSchemaWorkItemAdminJoin(joins ExtraSchemaWorkItemAdminJoins) ExtraSchemaWorkItemAdminSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAdminSelectConfig) {
		s.joins = ExtraSchemaWorkItemAdminJoins{
			AdminWorkItems: s.joins.AdminWorkItems || joins.AdminWorkItems,
			WorkItemAdmins: s.joins.WorkItemAdmins || joins.WorkItemAdmins,
		}
	}
}

// WithExtraSchemaWorkItemAdminFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaWorkItemAdminFilters(filters map[string][]any) ExtraSchemaWorkItemAdminSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAdminSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaWorkItemAdminHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaWorkItemAdminHavingClause(conditions map[string][]any) ExtraSchemaWorkItemAdminSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAdminSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaWorkItemAdminTableAdminWorkItemsJoinSQL = `-- M2M join generated from "work_item_admin_work_item_id_fkey"
left join (
	select
		work_item_admin.admin as work_item_admin_admin
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		extra_schema.work_item_admin
	join extra_schema.work_items on work_items.work_item_id = work_item_admin.work_item_id
	group by
		work_item_admin_admin
		, work_items.work_item_id
) as xo_join_work_item_admin_work_items on xo_join_work_item_admin_work_items.work_item_admin_admin = work_item_admin.admin
`

const extraSchemaWorkItemAdminTableAdminWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_admin_work_items.__work_items
		)) filter (where xo_join_work_item_admin_work_items.__work_items_work_item_id is not null), '{}') as work_item_admin_work_items`

const extraSchemaWorkItemAdminTableAdminWorkItemsGroupBySQL = `work_item_admin.admin, work_item_admin.work_item_id, work_item_admin.admin`

const extraSchemaWorkItemAdminTableWorkItemAdminsJoinSQL = `-- M2M join generated from "work_item_admin_admin_fkey"
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
) as xo_join_work_item_admin_admins on xo_join_work_item_admin_admins.work_item_admin_work_item_id = work_item_admin.work_item_id
`

const extraSchemaWorkItemAdminTableWorkItemAdminsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_admin_admins.__users
		)) filter (where xo_join_work_item_admin_admins.__users_user_id is not null), '{}') as work_item_admin_admins`

const extraSchemaWorkItemAdminTableWorkItemAdminsGroupBySQL = `work_item_admin.work_item_id, work_item_admin.work_item_id, work_item_admin.admin`

// ExtraSchemaWorkItemAdminUpdateParams represents update params for 'extra_schema.work_item_admin'.
type ExtraSchemaWorkItemAdminUpdateParams struct {
	Admin      *ExtraSchemaUserID     `json:"admin" nullable:"false"`      // admin
	WorkItemID *ExtraSchemaWorkItemID `json:"workItemID" nullable:"false"` // work_item_id
}

// SetUpdateParams updates extra_schema.work_item_admin struct fields with the specified params.
func (eswia *ExtraSchemaWorkItemAdmin) SetUpdateParams(params *ExtraSchemaWorkItemAdminUpdateParams) {
	if params.Admin != nil {
		eswia.Admin = *params.Admin
	}
	if params.WorkItemID != nil {
		eswia.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the ExtraSchemaWorkItemAdmin to the database.
func (eswia *ExtraSchemaWorkItemAdmin) Insert(ctx context.Context, db DB) (*ExtraSchemaWorkItemAdmin, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.work_item_admin (
	admin, work_item_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, eswia.Admin, eswia.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, eswia.Admin, eswia.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/Insert/db.Query: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	neweswia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAdmin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	*eswia = neweswia

	return eswia, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key or generated fields

// Delete deletes the ExtraSchemaWorkItemAdmin from the database.
func (eswia *ExtraSchemaWorkItemAdmin) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM extra_schema.work_item_admin 
	WHERE work_item_id = $1 AND admin = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, eswia.WorkItemID, eswia.Admin); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaWorkItemAdminsByAdminWorkItemID retrieves a row from 'extra_schema.work_item_admin' as a ExtraSchemaWorkItemAdmin.
//
// Generated from index 'work_item_admin_admin_work_item_id_idx'.
func ExtraSchemaWorkItemAdminsByAdminWorkItemID(ctx context.Context, db DB, admin ExtraSchemaUserID, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAdminSelectConfigOption) ([]ExtraSchemaWorkItemAdmin, error) {
	c := &ExtraSchemaWorkItemAdminSelectConfig{joins: ExtraSchemaWorkItemAdminJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.AdminWorkItems {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableAdminWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableAdminWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableAdminWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableWorkItemAdminsGroupBySQL)
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
	work_item_admin.admin,
	work_item_admin.work_item_id %s 
	 FROM extra_schema.work_item_admin %s 
	 WHERE work_item_admin.admin = $1 AND work_item_admin.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAdminsByAdminWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, admin, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{admin, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByAdminWorkItemID/Query: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAdmin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByAdminWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAdminByWorkItemIDAdmin retrieves a row from 'extra_schema.work_item_admin' as a ExtraSchemaWorkItemAdmin.
//
// Generated from index 'work_item_admin_pkey'.
func ExtraSchemaWorkItemAdminByWorkItemIDAdmin(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, admin ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAdminSelectConfigOption) (*ExtraSchemaWorkItemAdmin, error) {
	c := &ExtraSchemaWorkItemAdminSelectConfig{joins: ExtraSchemaWorkItemAdminJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.AdminWorkItems {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableAdminWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableAdminWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableAdminWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableWorkItemAdminsGroupBySQL)
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
	work_item_admin.admin,
	work_item_admin.work_item_id %s 
	 FROM extra_schema.work_item_admin %s 
	 WHERE work_item_admin.work_item_id = $1 AND work_item_admin.admin = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAdminByWorkItemIDAdmin */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, admin)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, admin}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_admin/WorkItemAdminByWorkItemIDAdmin/db.Query: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	eswia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAdmin])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_admin/WorkItemAdminByWorkItemIDAdmin/pgx.CollectOneRow: %w", &XoError{Entity: "Work item admin", Err: err}))
	}

	return &eswia, nil
}

// ExtraSchemaWorkItemAdminsByWorkItemID retrieves a row from 'extra_schema.work_item_admin' as a ExtraSchemaWorkItemAdmin.
//
// Generated from index 'work_item_admin_pkey'.
func ExtraSchemaWorkItemAdminsByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAdminSelectConfigOption) ([]ExtraSchemaWorkItemAdmin, error) {
	c := &ExtraSchemaWorkItemAdminSelectConfig{joins: ExtraSchemaWorkItemAdminJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.AdminWorkItems {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableAdminWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableAdminWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableAdminWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableWorkItemAdminsGroupBySQL)
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
	work_item_admin.admin,
	work_item_admin.work_item_id %s 
	 FROM extra_schema.work_item_admin %s 
	 WHERE work_item_admin.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAdminsByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByWorkItemIDAdmin/Query: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAdmin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByWorkItemIDAdmin/pgx.CollectRows: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAdminsByAdmin retrieves a row from 'extra_schema.work_item_admin' as a ExtraSchemaWorkItemAdmin.
//
// Generated from index 'work_item_admin_pkey'.
func ExtraSchemaWorkItemAdminsByAdmin(ctx context.Context, db DB, admin ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAdminSelectConfigOption) ([]ExtraSchemaWorkItemAdmin, error) {
	c := &ExtraSchemaWorkItemAdminSelectConfig{joins: ExtraSchemaWorkItemAdminJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.AdminWorkItems {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableAdminWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableAdminWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableAdminWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemAdmins {
		selectClauses = append(selectClauses, extraSchemaWorkItemAdminTableWorkItemAdminsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAdminTableWorkItemAdminsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAdminTableWorkItemAdminsGroupBySQL)
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
	work_item_admin.admin,
	work_item_admin.work_item_id %s 
	 FROM extra_schema.work_item_admin %s 
	 WHERE work_item_admin.admin = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAdminsByAdmin */\n" + sqlstr

	// run
	// logf(sqlstr, admin)
	rows, err := db.Query(ctx, sqlstr, append([]any{admin}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByWorkItemIDAdmin/Query: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAdmin])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAdmin/WorkItemAdminByWorkItemIDAdmin/pgx.CollectRows: %w", &XoError{Entity: "Work item admin", Err: err}))
	}
	return res, nil
}

// FKUser_Admin returns the User associated with the ExtraSchemaWorkItemAdmin's (Admin).
//
// Generated from foreign key 'work_item_admin_admin_fkey'.
func (eswia *ExtraSchemaWorkItemAdmin) FKUser_Admin(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, eswia.Admin)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the ExtraSchemaWorkItemAdmin's (WorkItemID).
//
// Generated from foreign key 'work_item_admin_work_item_id_fkey'.
func (eswia *ExtraSchemaWorkItemAdmin) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	return ExtraSchemaWorkItemByWorkItemID(ctx, db, eswia.WorkItemID)
}
