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

// XoTestsCacheDemoWorkItem represents a row from 'xo_tests.cache__demo_work_items'.
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
type XoTestsCacheDemoWorkItem struct {
	WorkItemID XoTestsWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Title      *string           `json:"title" db:"title"`                                              // title
	TeamID     XoTestsTeamID     `json:"teamID" db:"team_id" required:"true" nullable:"false"`          // team_id

	TeamJoin                     *XoTestsTeam                           `json:"-" db:"team_team_id" openapi-go:"ignore"`                           // O2O teams (inferred)
	WorkItemAssignedUsersJoin    *[]User__WIAU_XoTestsCacheDemoWorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
	WorkItemWorkItemCommentsJoin *[]XoTestsWorkItemComment              `json:"-" db:"work_item_comments" openapi-go:"ignore"`                     // M2O cache__demo_work_items
}

// XoTestsCacheDemoWorkItemCreateParams represents insert params for 'xo_tests.cache__demo_work_items'.
type XoTestsCacheDemoWorkItemCreateParams struct {
	TeamID     XoTestsTeamID     `json:"teamID" required:"true" nullable:"false"` // team_id
	Title      *string           `json:"title"`                                   // title
	WorkItemID XoTestsWorkItemID `json:"-" required:"true" nullable:"false"`      // work_item_id
}

// CreateXoTestsCacheDemoWorkItem creates a new XoTestsCacheDemoWorkItem in the database with the given params.
func CreateXoTestsCacheDemoWorkItem(ctx context.Context, db DB, params *XoTestsCacheDemoWorkItemCreateParams) (*XoTestsCacheDemoWorkItem, error) {
	xtcdwi := &XoTestsCacheDemoWorkItem{
		TeamID:     params.TeamID,
		Title:      params.Title,
		WorkItemID: params.WorkItemID,
	}

	return xtcdwi.Insert(ctx, db)
}

type XoTestsCacheDemoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsCacheDemoWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsCacheDemoWorkItemSelectConfigOption func(*XoTestsCacheDemoWorkItemSelectConfig)

// WithXoTestsCacheDemoWorkItemLimit limits row selection.
func WithXoTestsCacheDemoWorkItemLimit(limit int) XoTestsCacheDemoWorkItemSelectConfigOption {
	return func(s *XoTestsCacheDemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsCacheDemoWorkItemOrderBy string

type XoTestsCacheDemoWorkItemJoins struct {
	Team             bool // O2O teams
	AssignedUsers    bool // M2M work_item_assigned_user
	WorkItemComments bool // M2O work_item_comments
}

// WithXoTestsCacheDemoWorkItemJoin joins with the given tables.
func WithXoTestsCacheDemoWorkItemJoin(joins XoTestsCacheDemoWorkItemJoins) XoTestsCacheDemoWorkItemSelectConfigOption {
	return func(s *XoTestsCacheDemoWorkItemSelectConfig) {
		s.joins = XoTestsCacheDemoWorkItemJoins{
			Team:             s.joins.Team || joins.Team,
			AssignedUsers:    s.joins.AssignedUsers || joins.AssignedUsers,
			WorkItemComments: s.joins.WorkItemComments || joins.WorkItemComments,
		}
	}
}

// User__WIAU_XoTestsCacheDemoWorkItem represents a M2M join against "xo_tests.work_item_assigned_user"
type User__WIAU_XoTestsCacheDemoWorkItem struct {
	User XoTestsUser          `json:"user" db:"users" required:"true"`
	Role *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithXoTestsCacheDemoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsCacheDemoWorkItemFilters(filters map[string][]any) XoTestsCacheDemoWorkItemSelectConfigOption {
	return func(s *XoTestsCacheDemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsCacheDemoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See joins db tag to use the appropriate aliases.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsCacheDemoWorkItemHavingClause(conditions map[string][]any) XoTestsCacheDemoWorkItemSelectConfigOption {
	return func(s *XoTestsCacheDemoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const xoTestsCacheDemoWorkItemTableTeamJoinSQL = `-- O2O join generated from "cache__demo_work_items_team_id_fkey (inferred)"
left join xo_tests.teams as _cache__demo_work_items_team_id on _cache__demo_work_items_team_id.team_id = cache__demo_work_items.team_id
`

const xoTestsCacheDemoWorkItemTableTeamSelectSQL = `(case when _cache__demo_work_items_team_id.team_id is not null then row(_cache__demo_work_items_team_id.*) end) as team_team_id`

const xoTestsCacheDemoWorkItemTableTeamGroupBySQL = `_cache__demo_work_items_team_id.team_id,
      _cache__demo_work_items_team_id.team_id,
	cache__demo_work_items.work_item_id`

const xoTestsCacheDemoWorkItemTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey-shared-ref-cache__demo_work_items"
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
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = cache__demo_work_items.work_item_id
`

const xoTestsCacheDemoWorkItemTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const xoTestsCacheDemoWorkItemTableAssignedUsersGroupBySQL = `cache__demo_work_items.work_item_id, cache__demo_work_items.work_item_id`

const xoTestsCacheDemoWorkItemTableWorkItemCommentsJoinSQL = `-- M2O join generated from "work_item_comments_work_item_id_fkey-shared-ref-cache__demo_work_items"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    xo_tests.work_item_comments
  group by
        work_item_id
) as joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = cache__demo_work_items.work_item_id
`

const xoTestsCacheDemoWorkItemTableWorkItemCommentsSelectSQL = `COALESCE(joined_work_item_comments.work_item_comments, '{}') as work_item_comments`

const xoTestsCacheDemoWorkItemTableWorkItemCommentsGroupBySQL = `joined_work_item_comments.work_item_comments, cache__demo_work_items.work_item_id`

// XoTestsCacheDemoWorkItemUpdateParams represents update params for 'xo_tests.cache__demo_work_items'.
type XoTestsCacheDemoWorkItemUpdateParams struct {
	TeamID *XoTestsTeamID `json:"teamID" nullable:"false"` // team_id
	Title  **string       `json:"title"`                   // title
}

// SetUpdateParams updates xo_tests.cache__demo_work_items struct fields with the specified params.
func (xtcdwi *XoTestsCacheDemoWorkItem) SetUpdateParams(params *XoTestsCacheDemoWorkItemUpdateParams) {
	if params.TeamID != nil {
		xtcdwi.TeamID = *params.TeamID
	}
	if params.Title != nil {
		xtcdwi.Title = *params.Title
	}
}

// Insert inserts the XoTestsCacheDemoWorkItem to the database.
func (xtcdwi *XoTestsCacheDemoWorkItem) Insert(ctx context.Context, db DB) (*XoTestsCacheDemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.cache__demo_work_items (
	team_id, title, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtcdwi.TeamID, xtcdwi.Title, xtcdwi.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, xtcdwi.TeamID, xtcdwi.Title, xtcdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	newxtcdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsCacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	*xtcdwi = newxtcdwi

	return xtcdwi, nil
}

// Update updates a XoTestsCacheDemoWorkItem in the database.
func (xtcdwi *XoTestsCacheDemoWorkItem) Update(ctx context.Context, db DB) (*XoTestsCacheDemoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.cache__demo_work_items SET
	team_id = $1, title = $2
	WHERE work_item_id = $3
	RETURNING * `
	// run
	logf(sqlstr, xtcdwi.TeamID, xtcdwi.Title, xtcdwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, xtcdwi.TeamID, xtcdwi.Title, xtcdwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Update/db.Query: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	newxtcdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsCacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	*xtcdwi = newxtcdwi

	return xtcdwi, nil
}

// Upsert upserts a XoTestsCacheDemoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtcdwi *XoTestsCacheDemoWorkItem) Upsert(ctx context.Context, db DB, params *XoTestsCacheDemoWorkItemCreateParams) (*XoTestsCacheDemoWorkItem, error) {
	var err error

	xtcdwi.TeamID = params.TeamID
	xtcdwi.Title = params.Title
	xtcdwi.WorkItemID = params.WorkItemID

	xtcdwi, err = xtcdwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Cache  demo work item", Err: err})
			}
			xtcdwi, err = xtcdwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Cache  demo work item", Err: err})
			}
		}
	}

	return xtcdwi, err
}

// Delete deletes the XoTestsCacheDemoWorkItem from the database.
func (xtcdwi *XoTestsCacheDemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.cache__demo_work_items
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtcdwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsCacheDemoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of XoTestsCacheDemoWorkItem.
func XoTestsCacheDemoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID int, direction models.Direction, opts ...XoTestsCacheDemoWorkItemSelectConfigOption) ([]XoTestsCacheDemoWorkItem, error) {
	c := &XoTestsCacheDemoWorkItemSelectConfig{joins: XoTestsCacheDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableTeamGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsGroupBySQL)
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
	cache__demo_work_items.team_id,
	cache__demo_work_items.title,
	cache__demo_work_items.work_item_id %s
	 FROM xo_tests.cache__demo_work_items %s
	 WHERE cache__demo_work_items.work_item_id %s $1
	 %s   %s
  %s
  ORDER BY
		work_item_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsCacheDemoWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsCacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsCacheDemoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	return res, nil
}

// XoTestsCacheDemoWorkItemByWorkItemID retrieves a row from 'xo_tests.cache__demo_work_items' as a XoTestsCacheDemoWorkItem.
//
// Generated from index 'cache__demo_work_items_pkey'.
func XoTestsCacheDemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int, opts ...XoTestsCacheDemoWorkItemSelectConfigOption) (*XoTestsCacheDemoWorkItem, error) {
	c := &XoTestsCacheDemoWorkItemSelectConfig{joins: XoTestsCacheDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableTeamGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsCacheDemoWorkItemTableWorkItemCommentsGroupBySQL)
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
	cache__demo_work_items.team_id,
	cache__demo_work_items.title,
	cache__demo_work_items.work_item_id %s
	 FROM xo_tests.cache__demo_work_items %s
	 WHERE cache__demo_work_items.work_item_id = $1
	 %s   %s
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsCacheDemoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("cache__demo_work_items/CacheDemoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}
	xtcdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsCacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("cache__demo_work_items/CacheDemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo work item", Err: err}))
	}

	return &xtcdwi, nil
}

// FKTeam_TeamID returns the Team associated with the XoTestsCacheDemoWorkItem's (TeamID).
//
// Generated from foreign key 'cache__demo_work_items_team_id_fkey'.
func (xtcdwi *XoTestsCacheDemoWorkItem) FKTeam_TeamID(ctx context.Context, db DB) (*XoTestsTeam, error) {
	return XoTestsTeamByTeamID(ctx, db, xtcdwi.TeamID)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the XoTestsCacheDemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'cache__demo_work_items_work_item_id_fkey'.
func (xtcdwi *XoTestsCacheDemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	return XoTestsWorkItemByWorkItemID(ctx, db, xtcdwi.WorkItemID)
}
