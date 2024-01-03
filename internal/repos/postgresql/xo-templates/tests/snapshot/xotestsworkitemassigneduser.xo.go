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

// XoTestsWorkItemAssignedUser represents a row from 'xo_tests.work_item_assigned_user'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsWorkItemAssignedUser struct {
	WorkItemID   XoTestsWorkItemID    `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                           // work_item_id
	AssignedUser XoTestsUserID        `json:"assignedUser" db:"assigned_user" required:"true" nullable:"false"`                        // assigned_user
	XoTestsRole  *XoTestsWorkItemRole `json:"role" db:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role

	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_XoTestsWorkItemAssignedUser `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"`     // M2M work_item_assigned_user
	WorkItemAssignedUsersJoin *[]User__WIAU_XoTestsWorkItemAssignedUser     `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// XoTestsWorkItemAssignedUserCreateParams represents insert params for 'xo_tests.work_item_assigned_user'.
type XoTestsWorkItemAssignedUserCreateParams struct {
	AssignedUser XoTestsUserID        `json:"assignedUser" required:"true" nullable:"false"`                                 // assigned_user
	XoTestsRole  *XoTestsWorkItemRole `json:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID   XoTestsWorkItemID    `json:"workItemID" required:"true" nullable:"false"`                                   // work_item_id
}

// CreateXoTestsWorkItemAssignedUser creates a new XoTestsWorkItemAssignedUser in the database with the given params.
func CreateXoTestsWorkItemAssignedUser(ctx context.Context, db DB, params *XoTestsWorkItemAssignedUserCreateParams) (*XoTestsWorkItemAssignedUser, error) {
	xtwiau := &XoTestsWorkItemAssignedUser{
		AssignedUser: params.AssignedUser,
		XoTestsRole:  params.XoTestsRole,
		WorkItemID:   params.WorkItemID,
	}

	return xtwiau.Insert(ctx, db)
}

// XoTestsWorkItemAssignedUserUpdateParams represents update params for 'xo_tests.work_item_assigned_user'.
type XoTestsWorkItemAssignedUserUpdateParams struct {
	AssignedUser *XoTestsUserID        `json:"assignedUser" nullable:"false"`                                 // assigned_user
	XoTestsRole  **XoTestsWorkItemRole `json:"role" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID   *XoTestsWorkItemID    `json:"workItemID" nullable:"false"`                                   // work_item_id
}

// SetUpdateParams updates xo_tests.work_item_assigned_user struct fields with the specified params.
func (xtwiau *XoTestsWorkItemAssignedUser) SetUpdateParams(params *XoTestsWorkItemAssignedUserUpdateParams) {
	if params.AssignedUser != nil {
		xtwiau.AssignedUser = *params.AssignedUser
	}
	if params.XoTestsRole != nil {
		xtwiau.XoTestsRole = *params.XoTestsRole
	}
	if params.WorkItemID != nil {
		xtwiau.WorkItemID = *params.WorkItemID
	}
}

type XoTestsWorkItemAssignedUserSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsWorkItemAssignedUserJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsWorkItemAssignedUserSelectConfigOption func(*XoTestsWorkItemAssignedUserSelectConfig)

// WithXoTestsWorkItemAssignedUserLimit limits row selection.
func WithXoTestsWorkItemAssignedUserLimit(limit int) XoTestsWorkItemAssignedUserSelectConfigOption {
	return func(s *XoTestsWorkItemAssignedUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsWorkItemAssignedUserOrderBy string

type XoTestsWorkItemAssignedUserJoins struct {
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
	AssignedUsers         bool // M2M work_item_assigned_user
}

// WithXoTestsWorkItemAssignedUserJoin joins with the given tables.
func WithXoTestsWorkItemAssignedUserJoin(joins XoTestsWorkItemAssignedUserJoins) XoTestsWorkItemAssignedUserSelectConfigOption {
	return func(s *XoTestsWorkItemAssignedUserSelectConfig) {
		s.joins = XoTestsWorkItemAssignedUserJoins{
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
			AssignedUsers:         s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// WorkItem__WIAU_XoTestsWorkItemAssignedUser represents a M2M join against "xo_tests.work_item_assigned_user"
type WorkItem__WIAU_XoTestsWorkItemAssignedUser struct {
	WorkItem XoTestsWorkItem      `json:"workItem" db:"work_items" required:"true"`
	Role     *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// User__WIAU_XoTestsWorkItemAssignedUser represents a M2M join against "xo_tests.work_item_assigned_user"
type User__WIAU_XoTestsWorkItemAssignedUser struct {
	User XoTestsUser          `json:"user" db:"users" required:"true"`
	Role *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithXoTestsWorkItemAssignedUserFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsWorkItemAssignedUserFilters(filters map[string][]any) XoTestsWorkItemAssignedUserSelectConfigOption {
	return func(s *XoTestsWorkItemAssignedUserSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsWorkItemAssignedUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// // filter a given aggregate of assigned users to return results where at least one of them has id of userId
//
//	filters := map[string][]any{
//		"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsWorkItemAssignedUserHavingClause(conditions map[string][]any) XoTestsWorkItemAssignedUserSelectConfigOption {
	return func(s *XoTestsWorkItemAssignedUserSelectConfig) {
		s.having = conditions
	}
}

const xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
		work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
		, work_item_assigned_user.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		xo_tests.work_item_assigned_user
	join xo_tests.work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
	group by
		work_item_assigned_user_assigned_user
		, work_items.work_item_id
		, role
) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user
`

const xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL = `work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

const xoTestsWorkItemAssignedUserTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
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
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
`

const xoTestsWorkItemAssignedUserTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const xoTestsWorkItemAssignedUserTableAssignedUsersGroupBySQL = `work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

// Insert inserts the XoTestsWorkItemAssignedUser to the database.
func (xtwiau *XoTestsWorkItemAssignedUser) Insert(ctx context.Context, db DB) (*XoTestsWorkItemAssignedUser, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.work_item_assigned_user (
	assigned_user, role, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtwiau.AssignedUser, xtwiau.XoTestsRole, xtwiau.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, xtwiau.AssignedUser, xtwiau.XoTestsRole, xtwiau.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/Insert/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	newxtwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*xtwiau = newxtwiau

	return xtwiau, nil
}

// Update updates a XoTestsWorkItemAssignedUser in the database.
func (xtwiau *XoTestsWorkItemAssignedUser) Update(ctx context.Context, db DB) (*XoTestsWorkItemAssignedUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.work_item_assigned_user SET 
	role = $1 
	WHERE work_item_id = $2  AND assigned_user = $3 
	RETURNING * `
	// run
	logf(sqlstr, xtwiau.XoTestsRole, xtwiau.WorkItemID, xtwiau.AssignedUser)

	rows, err := db.Query(ctx, sqlstr, xtwiau.XoTestsRole, xtwiau.WorkItemID, xtwiau.AssignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/Update/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	newxtwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*xtwiau = newxtwiau

	return xtwiau, nil
}

// Upsert upserts a XoTestsWorkItemAssignedUser in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtwiau *XoTestsWorkItemAssignedUser) Upsert(ctx context.Context, db DB, params *XoTestsWorkItemAssignedUserCreateParams) (*XoTestsWorkItemAssignedUser, error) {
	var err error

	xtwiau.AssignedUser = params.AssignedUser
	xtwiau.XoTestsRole = params.XoTestsRole
	xtwiau.WorkItemID = params.WorkItemID

	xtwiau, err = xtwiau.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
			xtwiau, err = xtwiau.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
		}
	}

	return xtwiau, err
}

// Delete deletes the XoTestsWorkItemAssignedUser from the database.
func (xtwiau *XoTestsWorkItemAssignedUser) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.work_item_assigned_user 
	WHERE work_item_id = $1 AND assigned_user = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtwiau.WorkItemID, xtwiau.AssignedUser); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsWorkItemAssignedUsersByAssignedUserWorkItemID retrieves a row from 'xo_tests.work_item_assigned_user' as a XoTestsWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_assigned_user_work_item_id_idx'.
func XoTestsWorkItemAssignedUsersByAssignedUserWorkItemID(ctx context.Context, db DB, assignedUser XoTestsUserID, workItemID XoTestsWorkItemID, opts ...XoTestsWorkItemAssignedUserSelectConfigOption) ([]XoTestsWorkItemAssignedUser, error) {
	c := &XoTestsWorkItemAssignedUserSelectConfig{joins: XoTestsWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM xo_tests.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1 AND work_item_assigned_user.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemAssignedUsersByAssignedUserWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// XoTestsWorkItemAssignedUserByWorkItemIDAssignedUser retrieves a row from 'xo_tests.work_item_assigned_user' as a XoTestsWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func XoTestsWorkItemAssignedUserByWorkItemIDAssignedUser(ctx context.Context, db DB, workItemID XoTestsWorkItemID, assignedUser XoTestsUserID, opts ...XoTestsWorkItemAssignedUserSelectConfigOption) (*XoTestsWorkItemAssignedUser, error) {
	c := &XoTestsWorkItemAssignedUserSelectConfig{joins: XoTestsWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM xo_tests.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1 AND work_item_assigned_user.assigned_user = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemAssignedUserByWorkItemIDAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	xtwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}

	return &xtwiau, nil
}

// XoTestsWorkItemAssignedUsersByWorkItemID retrieves a row from 'xo_tests.work_item_assigned_user' as a XoTestsWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func XoTestsWorkItemAssignedUsersByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, opts ...XoTestsWorkItemAssignedUserSelectConfigOption) ([]XoTestsWorkItemAssignedUser, error) {
	c := &XoTestsWorkItemAssignedUserSelectConfig{joins: XoTestsWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM xo_tests.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemAssignedUsersByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// XoTestsWorkItemAssignedUsersByAssignedUser retrieves a row from 'xo_tests.work_item_assigned_user' as a XoTestsWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func XoTestsWorkItemAssignedUsersByAssignedUser(ctx context.Context, db DB, assignedUser XoTestsUserID, opts ...XoTestsWorkItemAssignedUserSelectConfigOption) ([]XoTestsWorkItemAssignedUser, error) {
	c := &XoTestsWorkItemAssignedUserSelectConfig{joins: XoTestsWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM xo_tests.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemAssignedUsersByAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// FKUser_AssignedUser returns the User associated with the XoTestsWorkItemAssignedUser's (AssignedUser).
//
// Generated from foreign key 'work_item_assigned_user_assigned_user_fkey'.
func (xtwiau *XoTestsWorkItemAssignedUser) FKUser_AssignedUser(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtwiau.AssignedUser)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the XoTestsWorkItemAssignedUser's (WorkItemID).
//
// Generated from foreign key 'work_item_assigned_user_work_item_id_fkey'.
func (xtwiau *XoTestsWorkItemAssignedUser) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	return XoTestsWorkItemByWorkItemID(ctx, db, xtwiau.WorkItemID)
}
