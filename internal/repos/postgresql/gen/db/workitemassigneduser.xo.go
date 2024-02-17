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

// WorkItemAssignedUser represents a row from 'public.work_item_assigned_user'.
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
type WorkItemAssignedUser struct {
	WorkItemID   WorkItemID          `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                           // work_item_id
	AssignedUser UserID              `json:"assignedUser" db:"assigned_user" required:"true" nullable:"false"`                        // assigned_user
	Role         models.WorkItemRole `json:"role" db:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role

	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_WorkItemAssignedUser `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"`     // M2M work_item_assigned_user
	WorkItemAssignedUsersJoin *[]User__WIAU_WorkItemAssignedUser     `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user

}

// WorkItemAssignedUserCreateParams represents insert params for 'public.work_item_assigned_user'.
type WorkItemAssignedUserCreateParams struct {
	AssignedUser UserID              `json:"assignedUser" required:"true" nullable:"false"`                                 // assigned_user
	Role         models.WorkItemRole `json:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID   WorkItemID          `json:"workItemID" required:"true" nullable:"false"`                                   // work_item_id
}

// CreateWorkItemAssignedUser creates a new WorkItemAssignedUser in the database with the given params.
func CreateWorkItemAssignedUser(ctx context.Context, db DB, params *WorkItemAssignedUserCreateParams) (*WorkItemAssignedUser, error) {
	wiau := &WorkItemAssignedUser{
		AssignedUser: params.AssignedUser,
		Role:         params.Role,
		WorkItemID:   params.WorkItemID,
	}

	return wiau.Insert(ctx, db)
}

type WorkItemAssignedUserSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemAssignedUserJoins
	filters map[string][]any
	having  map[string][]any
}
type WorkItemAssignedUserSelectConfigOption func(*WorkItemAssignedUserSelectConfig)

// WithWorkItemAssignedUserLimit limits row selection.
func WithWorkItemAssignedUserLimit(limit int) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemAssignedUserOrderBy string

const ()

type WorkItemAssignedUserJoins struct {
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
	AssignedUsers         bool // M2M work_item_assigned_user
}

// WithWorkItemAssignedUserJoin joins with the given tables.
func WithWorkItemAssignedUserJoin(joins WorkItemAssignedUserJoins) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		s.joins = WorkItemAssignedUserJoins{
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
			AssignedUsers:         s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// WorkItem__WIAU_WorkItemAssignedUser represents a M2M join against "public.work_item_assigned_user"
type WorkItem__WIAU_WorkItemAssignedUser struct {
	WorkItem WorkItem            `json:"workItem" db:"work_items" required:"true"`
	Role     models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// User__WIAU_WorkItemAssignedUser represents a M2M join against "public.work_item_assigned_user"
type User__WIAU_WorkItemAssignedUser struct {
	User User                `json:"user" db:"users" required:"true"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithWorkItemAssignedUserFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemAssignedUserFilters(filters map[string][]any) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		s.filters = filters
	}
}

// WithWorkItemAssignedUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithWorkItemAssignedUserHavingClause(conditions map[string][]any) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		s.having = conditions
	}
}

const workItemAssignedUserTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
		work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
		, work_item_assigned_user.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		work_item_assigned_user
	join work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
	group by
		work_item_assigned_user_assigned_user
		, work_items.work_item_id
		, role
) as xo_join_work_item_assigned_user_work_items on xo_join_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user
`

const workItemAssignedUserTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assigned_user_work_items.__work_items
		, xo_join_work_item_assigned_user_work_items.role
		)) filter (where xo_join_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const workItemAssignedUserTableWorkItemsAssignedUserGroupBySQL = `work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

const workItemAssignedUserTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
		work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
		, work_item_assigned_user.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		work_item_assigned_user
	join users on users.user_id = work_item_assigned_user.assigned_user
	group by
		work_item_assigned_user_work_item_id
		, users.user_id
		, role
) as xo_join_work_item_assigned_user_assigned_users on xo_join_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
`

const workItemAssignedUserTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assigned_user_assigned_users.__users
		, xo_join_work_item_assigned_user_assigned_users.role
		)) filter (where xo_join_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const workItemAssignedUserTableAssignedUsersGroupBySQL = `work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

// WorkItemAssignedUserUpdateParams represents update params for 'public.work_item_assigned_user'.
type WorkItemAssignedUserUpdateParams struct {
	AssignedUser *UserID              `json:"assignedUser" nullable:"false"`                                 // assigned_user
	Role         *models.WorkItemRole `json:"role" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID   *WorkItemID          `json:"workItemID" nullable:"false"`                                   // work_item_id
}

// SetUpdateParams updates public.work_item_assigned_user struct fields with the specified params.
func (wiau *WorkItemAssignedUser) SetUpdateParams(params *WorkItemAssignedUserUpdateParams) {
	if params.AssignedUser != nil {
		wiau.AssignedUser = *params.AssignedUser
	}
	if params.Role != nil {
		wiau.Role = *params.Role
	}
	if params.WorkItemID != nil {
		wiau.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the WorkItemAssignedUser to the database.
func (wiau *WorkItemAssignedUser) Insert(ctx context.Context, db DB) (*WorkItemAssignedUser, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_assigned_user (
	assigned_user, role, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, wiau.AssignedUser, wiau.Role, wiau.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, wiau.AssignedUser, wiau.Role, wiau.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Insert/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	newwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*wiau = newwiau

	return wiau, nil
}

// Update updates a WorkItemAssignedUser in the database.
func (wiau *WorkItemAssignedUser) Update(ctx context.Context, db DB) (*WorkItemAssignedUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_assigned_user SET 
	role = $1 
	WHERE work_item_id = $2  AND assigned_user = $3 
	RETURNING * `
	// run
	logf(sqlstr, wiau.Role, wiau.WorkItemID, wiau.AssignedUser)

	rows, err := db.Query(ctx, sqlstr, wiau.Role, wiau.WorkItemID, wiau.AssignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Update/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	newwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*wiau = newwiau

	return wiau, nil
}

// Upsert upserts a WorkItemAssignedUser in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wiau *WorkItemAssignedUser) Upsert(ctx context.Context, db DB, params *WorkItemAssignedUserCreateParams) (*WorkItemAssignedUser, error) {
	var err error

	wiau.AssignedUser = params.AssignedUser
	wiau.Role = params.Role
	wiau.WorkItemID = params.WorkItemID

	wiau, err = wiau.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
			wiau, err = wiau.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
		}
	}

	return wiau, err
}

// Delete deletes the WorkItemAssignedUser from the database.
func (wiau *WorkItemAssignedUser) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_assigned_user 
	WHERE work_item_id = $1 AND assigned_user = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wiau.WorkItemID, wiau.AssignedUser); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemAssignedUsersByAssignedUserWorkItemID retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_assigned_user_work_item_id_idx'.
func WorkItemAssignedUsersByAssignedUserWorkItemID(ctx context.Context, db DB, assignedUser UserID, workItemID WorkItemID, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableAssignedUsersGroupBySQL)
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
	 FROM public.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1 AND work_item_assigned_user.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssignedUsersByAssignedUserWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// WorkItemAssignedUserByWorkItemIDAssignedUser retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUserByWorkItemIDAssignedUser(ctx context.Context, db DB, workItemID WorkItemID, assignedUser UserID, opts ...WorkItemAssignedUserSelectConfigOption) (*WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableAssignedUsersGroupBySQL)
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
	 FROM public.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1 AND work_item_assigned_user.assigned_user = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssignedUserByWorkItemIDAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	wiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}

	return &wiau, nil
}

// WorkItemAssignedUsersByWorkItemID retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUsersByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableAssignedUsersGroupBySQL)
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
	 FROM public.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssignedUsersByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// WorkItemAssignedUsersByAssignedUser retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUsersByAssignedUser(ctx context.Context, db DB, assignedUser UserID, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssignedUserTableAssignedUsersGroupBySQL)
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
	 FROM public.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssignedUsersByAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// FKUser_AssignedUser returns the User associated with the WorkItemAssignedUser's (AssignedUser).
//
// Generated from foreign key 'work_item_assigned_user_assigned_user_fkey'.
func (wiau *WorkItemAssignedUser) FKUser_AssignedUser(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wiau.AssignedUser)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemAssignedUser's (WorkItemID).
//
// Generated from foreign key 'work_item_assigned_user_work_item_id_fkey'.
func (wiau *WorkItemAssignedUser) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wiau.WorkItemID)
}
