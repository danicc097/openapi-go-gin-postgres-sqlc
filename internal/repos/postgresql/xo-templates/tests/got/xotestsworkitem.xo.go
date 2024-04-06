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

// XoTestsWorkItem represents a row from 'xo_tests.work_items'.
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
type XoTestsWorkItem struct {
	WorkItemID  XoTestsWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Title       *string           `json:"title" db:"title"`                                              // title
	Description *string           `json:"description" db:"description"`                                  // description
	TeamID      XoTestsTeamID     `json:"teamID" db:"team_id" required:"true" nullable:"false"`          // team_id

	DemoWorkItemJoin     *XoTestsDemoWorkItem             `json:"-" db:"demo_work_item_work_item_id" openapi-go:"ignore"`  // O2O demo_work_items (inferred)
	TimeEntriesJoin      *[]XoTestsTimeEntry              `json:"-" db:"time_entries" openapi-go:"ignore"`                 // M2O work_items
	AssigneesJoin        *[]XoTestsWorkItemM2MAssigneeWIA `json:"-" db:"work_item_assignee_assignees" openapi-go:"ignore"` // M2M work_item_assignee
	WorkItemCommentsJoin *[]XoTestsWorkItemComment        `json:"-" db:"work_item_comments" openapi-go:"ignore"`           // M2O work_items
	TeamJoin             *XoTestsTeam                     `json:"-" db:"team_team_id" openapi-go:"ignore"`                 // O2O teams (inferred)
}

// XoTestsWorkItemCreateParams represents insert params for 'xo_tests.work_items'.
type XoTestsWorkItemCreateParams struct {
	Description *string       `json:"description"`                             // description
	TeamID      XoTestsTeamID `json:"teamID" required:"true" nullable:"false"` // team_id
	Title       *string       `json:"title"`                                   // title
}

// XoTestsWorkItemParams represents common params for both insert and update of 'xo_tests.work_items'.
type XoTestsWorkItemParams interface {
	GetDescription() *string
	GetTeamID() *XoTestsTeamID
	GetTitle() *string
}

func (p XoTestsWorkItemCreateParams) GetDescription() *string {
	return p.Description
}

func (p XoTestsWorkItemUpdateParams) GetDescription() *string {
	if p.Description != nil {
		return *p.Description
	}
	return nil
}

func (p XoTestsWorkItemCreateParams) GetTeamID() *XoTestsTeamID {
	x := p.TeamID
	return &x
}

func (p XoTestsWorkItemUpdateParams) GetTeamID() *XoTestsTeamID {
	return p.TeamID
}

func (p XoTestsWorkItemCreateParams) GetTitle() *string {
	return p.Title
}

func (p XoTestsWorkItemUpdateParams) GetTitle() *string {
	if p.Title != nil {
		return *p.Title
	}
	return nil
}

type XoTestsWorkItemID int

// CreateXoTestsWorkItem creates a new XoTestsWorkItem in the database with the given params.
func CreateXoTestsWorkItem(ctx context.Context, db DB, params *XoTestsWorkItemCreateParams) (*XoTestsWorkItem, error) {
	xtwi := &XoTestsWorkItem{
		Description: params.Description,
		TeamID:      params.TeamID,
		Title:       params.Title,
	}

	return xtwi.Insert(ctx, db)
}

type XoTestsWorkItemSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   XoTestsWorkItemJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsWorkItemSelectConfigOption func(*XoTestsWorkItemSelectConfig)

// WithXoTestsWorkItemLimit limits row selection.
func WithXoTestsWorkItemLimit(limit int) XoTestsWorkItemSelectConfigOption {
	return func(s *XoTestsWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsWorkItemOrderBy string

type XoTestsWorkItemJoins struct {
	DemoWorkItem     bool `json:"demoWorkItem" required:"true" nullable:"false"`     // O2O demo_work_items
	TimeEntries      bool `json:"timeEntries" required:"true" nullable:"false"`      // M2O time_entries
	Assignees        bool `json:"assignees" required:"true" nullable:"false"`        // M2M work_item_assignee
	WorkItemComments bool `json:"workItemComments" required:"true" nullable:"false"` // M2O work_item_comments
	Team             bool `json:"team" required:"true" nullable:"false"`             // O2O teams
}

// WithXoTestsWorkItemJoin joins with the given tables.
func WithXoTestsWorkItemJoin(joins XoTestsWorkItemJoins) XoTestsWorkItemSelectConfigOption {
	return func(s *XoTestsWorkItemSelectConfig) {
		s.joins = XoTestsWorkItemJoins{
			DemoWorkItem:     s.joins.DemoWorkItem || joins.DemoWorkItem,
			TimeEntries:      s.joins.TimeEntries || joins.TimeEntries,
			Assignees:        s.joins.Assignees || joins.Assignees,
			WorkItemComments: s.joins.WorkItemComments || joins.WorkItemComments,
			Team:             s.joins.Team || joins.Team,
		}
	}
}

// XoTestsWorkItemM2MAssigneeWIA represents a M2M join against "xo_tests.work_item_assignee"
type XoTestsWorkItemM2MAssigneeWIA struct {
	User XoTestsUser          `json:"user" db:"users" required:"true"`
	Role *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithXoTestsWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsWorkItemFilters(filters map[string][]any) XoTestsWorkItemSelectConfigOption {
	return func(s *XoTestsWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsWorkItemHavingClause(conditions map[string][]any) XoTestsWorkItemSelectConfigOption {
	return func(s *XoTestsWorkItemSelectConfig) {
		s.having = conditions
	}
}

const xoTestsWorkItemTableDemoWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const xoTestsWorkItemTableDemoWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id`

const xoTestsWorkItemTableDemoWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const xoTestsWorkItemTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , row(time_entries.*) as __time_entries
  from
    xo_tests.time_entries
  group by
	  time_entries_work_item_id, xo_tests.time_entries.time_entry_id
) as xo_join_time_entries on xo_join_time_entries.time_entries_work_item_id = work_items.work_item_id
`

const xoTestsWorkItemTableTimeEntriesSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_time_entries.__time_entries)) filter (where xo_join_time_entries.time_entries_work_item_id is not null), '{}') as time_entries`

const xoTestsWorkItemTableTimeEntriesGroupBySQL = `work_items.work_item_id`

const xoTestsWorkItemTableAssigneesJoinSQL = `-- M2M join generated from "work_item_assignee_assignee_fkey"
left join (
	select
		work_item_assignee.work_item_id as work_item_assignee_work_item_id
		, work_item_assignee.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.work_item_assignee
	join xo_tests.users on users.user_id = work_item_assignee.assignee
	group by
		work_item_assignee_work_item_id
		, users.user_id
		, role
) as xo_join_work_item_assignee_assignees on xo_join_work_item_assignee_assignees.work_item_assignee_work_item_id = work_items.work_item_id
`

const xoTestsWorkItemTableAssigneesSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_assignees.__users
		, xo_join_work_item_assignee_assignees.role
		)) filter (where xo_join_work_item_assignee_assignees.__users_user_id is not null), '{}') as work_item_assignee_assignees`

const xoTestsWorkItemTableAssigneesGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

const xoTestsWorkItemTableWorkItemCommentsJoinSQL = `-- M2O join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , row(work_item_comments.*) as __work_item_comments
  from
    xo_tests.work_item_comments
  group by
	  work_item_comments_work_item_id, xo_tests.work_item_comments.work_item_comment_id
) as xo_join_work_item_comments on xo_join_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
`

const xoTestsWorkItemTableWorkItemCommentsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_work_item_comments.__work_item_comments)) filter (where xo_join_work_item_comments.work_item_comments_work_item_id is not null), '{}') as work_item_comments`

const xoTestsWorkItemTableWorkItemCommentsGroupBySQL = `work_items.work_item_id`

const xoTestsWorkItemTableTeamJoinSQL = `-- O2O join generated from "work_items_team_id_fkey (inferred)"
left join xo_tests.teams as _work_items_team_id on _work_items_team_id.team_id = work_items.team_id
`

const xoTestsWorkItemTableTeamSelectSQL = `(case when _work_items_team_id.team_id is not null then row(_work_items_team_id.*) end) as team_team_id`

const xoTestsWorkItemTableTeamGroupBySQL = `_work_items_team_id.team_id,
      _work_items_team_id.team_id,
	work_items.work_item_id`

// XoTestsWorkItemUpdateParams represents update params for 'xo_tests.work_items'.
type XoTestsWorkItemUpdateParams struct {
	Description **string       `json:"description"`             // description
	TeamID      *XoTestsTeamID `json:"teamID" nullable:"false"` // team_id
	Title       **string       `json:"title"`                   // title
}

// SetUpdateParams updates xo_tests.work_items struct fields with the specified params.
func (xtwi *XoTestsWorkItem) SetUpdateParams(params *XoTestsWorkItemUpdateParams) {
	if params.Description != nil {
		xtwi.Description = *params.Description
	}
	if params.TeamID != nil {
		xtwi.TeamID = *params.TeamID
	}
	if params.Title != nil {
		xtwi.Title = *params.Title
	}
}

// Insert inserts the XoTestsWorkItem to the database.
func (xtwi *XoTestsWorkItem) Insert(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.work_items (
	description, team_id, title
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtwi.Description, xtwi.TeamID, xtwi.Title)

	rows, err := db.Query(ctx, sqlstr, xtwi.Description, xtwi.TeamID, xtwi.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Insert/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newxtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	*xtwi = newxtwi

	return xtwi, nil
}

// Update updates a XoTestsWorkItem in the database.
func (xtwi *XoTestsWorkItem) Update(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.work_items SET 
	description = $1, team_id = $2, title = $3 
	WHERE work_item_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, xtwi.Description, xtwi.TeamID, xtwi.Title, xtwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, xtwi.Description, xtwi.TeamID, xtwi.Title, xtwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Update/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newxtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}
	*xtwi = newxtwi

	return xtwi, nil
}

// Upsert upserts a XoTestsWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtwi *XoTestsWorkItem) Upsert(ctx context.Context, db DB, params *XoTestsWorkItemCreateParams) (*XoTestsWorkItem, error) {
	var err error

	xtwi.Description = params.Description
	xtwi.TeamID = params.TeamID
	xtwi.Title = params.Title

	xtwi, err = xtwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item", Err: err})
			}
			xtwi, err = xtwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item", Err: err})
			}
		}
	}

	return xtwi, err
}

// Delete deletes the XoTestsWorkItem from the database.
func (xtwi *XoTestsWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsWorkItemPaginatedByWorkItemID returns a cursor-paginated list of XoTestsWorkItem.
func XoTestsWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, direction models.Direction, opts ...XoTestsWorkItemSelectConfigOption) ([]XoTestsWorkItem, error) {
	c := &XoTestsWorkItemSelectConfig{
		joins:   XoTestsWorkItemJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

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
		selectClauses = append(selectClauses, xoTestsWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, xoTestsWorkItemTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableAssigneesGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTeamGroupBySQL)
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
	work_items.team_id,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.work_item_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// XoTestsWorkItems retrieves a row from 'xo_tests.work_items' as a XoTestsWorkItem.
//
// Generated from index '[xo] base filter query'.
func XoTestsWorkItems(ctx context.Context, db DB, opts ...XoTestsWorkItemSelectConfigOption) ([]XoTestsWorkItem, error) {
	c := &XoTestsWorkItemSelectConfig{joins: XoTestsWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, xoTestsWorkItemTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableAssigneesGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTeamGroupBySQL)
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
	work_items.team_id,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE true
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItems */\n" + sqlstr

	// run
	// logf(sqlstr, )
	rows, err := db.Query(ctx, sqlstr, append([]any{}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/WorkItemsByDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/WorkItemsByDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// XoTestsWorkItemByWorkItemID retrieves a row from 'xo_tests.work_items' as a XoTestsWorkItem.
//
// Generated from index 'work_items_pkey'.
func XoTestsWorkItemByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, opts ...XoTestsWorkItemSelectConfigOption) (*XoTestsWorkItem, error) {
	c := &XoTestsWorkItemSelectConfig{joins: XoTestsWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, xoTestsWorkItemTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableAssigneesGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTeamGroupBySQL)
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
	work_items.team_id,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	xtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	return &xtwi, nil
}

// XoTestsWorkItemsByTitle retrieves a row from 'xo_tests.work_items' as a XoTestsWorkItem.
//
// Generated from index 'work_items_title_description_idx1'.
func XoTestsWorkItemsByTitle(ctx context.Context, db DB, title *string, opts ...XoTestsWorkItemSelectConfigOption) ([]XoTestsWorkItem, error) {
	c := &XoTestsWorkItemSelectConfig{joins: XoTestsWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, xoTestsWorkItemTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableAssigneesGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, xoTestsWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, xoTestsWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemTableTeamGroupBySQL)
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
	work_items.team_id,
	work_items.title,
	work_items.work_item_id %s 
	 FROM xo_tests.work_items %s 
	 WHERE work_items.title = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemsByTitle */\n" + sqlstr

	// run
	// logf(sqlstr, title)
	rows, err := db.Query(ctx, sqlstr, append([]any{title}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/WorkItemsByTitleDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItem/WorkItemsByTitleDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// FKTeam_TeamID returns the Team associated with the XoTestsWorkItem's (TeamID).
//
// Generated from foreign key 'work_items_team_id_fkey'.
func (xtwi *XoTestsWorkItem) FKTeam_TeamID(ctx context.Context, db DB) (*XoTestsTeam, error) {
	return XoTestsTeamByTeamID(ctx, db, xtwi.TeamID)
}
