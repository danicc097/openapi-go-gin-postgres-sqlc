// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Team represents a row from 'public.teams'.
type Team struct {
	TeamID      TeamID    `json:"teamID" db:"team_id" required:"true" nullable:"false"`          // team_id
	ProjectID   ProjectID `json:"projectID" db:"project_id" required:"true" nullable:"false"`    // project_id
	Name        string    `json:"name" db:"name" required:"true" nullable:"false"`               // name
	Description string    `json:"description" db:"description" required:"true" nullable:"false"` // description
	CreatedAt   time.Time `json:"createdAt" db:"created_at" required:"true" nullable:"false"`    // created_at
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`    // updated_at

	ProjectJoin     *Project     `json:"-" db:"project_project_id"` // O2O projects (generated from M2O)
	TimeEntriesJoin *[]TimeEntry `json:"-" db:"time_entries"`       // M2O teams
	MembersJoin     *[]User      `json:"-" db:"user_team_members"`  // M2M user_team

}

// TeamCreateParams represents insert params for 'public.teams'.
type TeamCreateParams struct {
	Description string    `json:"description" required:"true" nullable:"false"` // description
	Name        string    `json:"name" required:"true" nullable:"false"`        // name
	ProjectID   ProjectID `json:"-"`                                            // project_id
}

// TeamParams represents common params for both insert and update of 'public.teams'.
type TeamParams interface {
	GetDescription() *string
	GetName() *string
	GetProjectID() *ProjectID
}

func (p TeamCreateParams) GetDescription() *string {
	x := p.Description
	return &x
}
func (p TeamUpdateParams) GetDescription() *string {
	return p.Description
}

func (p TeamCreateParams) GetName() *string {
	x := p.Name
	return &x
}
func (p TeamUpdateParams) GetName() *string {
	return p.Name
}

func (p TeamCreateParams) GetProjectID() *ProjectID {
	x := p.ProjectID
	return &x
}
func (p TeamUpdateParams) GetProjectID() *ProjectID {
	return p.ProjectID
}

type TeamID int

// CreateTeam creates a new Team in the database with the given params.
func CreateTeam(ctx context.Context, db DB, params *TeamCreateParams) (*Team, error) {
	t := &Team{
		Description: params.Description,
		Name:        params.Name,
		ProjectID:   params.ProjectID,
	}

	return t.Insert(ctx, db)
}

type TeamSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   TeamJoins
	filters map[string][]any
	having  map[string][]any
}
type TeamSelectConfigOption func(*TeamSelectConfig)

// WithTeamLimit limits row selection.
func WithTeamLimit(limit int) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithTeamOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithTeamOrderBy(rows map[string]*Direction) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		te := EntityFields[TableEntityTeam]
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

type TeamJoins struct {
	Project     bool `json:"project" required:"true" nullable:"false"`     // O2O projects
	TimeEntries bool `json:"timeEntries" required:"true" nullable:"false"` // M2O time_entries
	Members     bool `json:"members" required:"true" nullable:"false"`     // M2M user_team
}

// WithTeamJoin joins with the given tables.
func WithTeamJoin(joins TeamJoins) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.joins = TeamJoins{
			Project:     s.joins.Project || joins.Project,
			TimeEntries: s.joins.TimeEntries || joins.TimeEntries,
			Members:     s.joins.Members || joins.Members,
		}
	}
}

// WithTeamFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithTeamFilters(filters map[string][]any) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.filters = filters
	}
}

// WithTeamHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithTeamHavingClause(conditions map[string][]any) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.having = conditions
	}
}

const teamTableProjectJoinSQL = `-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects as _teams_project_id on _teams_project_id.project_id = teams.project_id
`

const teamTableProjectSelectSQL = `(case when _teams_project_id.project_id is not null then row(_teams_project_id.*) end) as project_project_id`

const teamTableProjectGroupBySQL = `_teams_project_id.project_id,
      _teams_project_id.project_id,
	teams.team_id`

const teamTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_team_id_fkey"
left join (
  select
  team_id as time_entries_team_id
    , row(time_entries.*) as __time_entries
  from
    time_entries
  group by
	  time_entries_team_id, time_entries.time_entry_id
) as xo_join_time_entries on xo_join_time_entries.time_entries_team_id = teams.team_id
`

const teamTableTimeEntriesSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_time_entries.__time_entries)) filter (where xo_join_time_entries.time_entries_team_id is not null), '{}') as time_entries`

const teamTableTimeEntriesGroupBySQL = `teams.team_id`

const teamTableMembersJoinSQL = `-- M2M join generated from "user_team_member_fkey"
left join (
	select
		user_team.team_id as user_team_team_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		user_team
	join users on users.user_id = user_team.member
	group by
		user_team_team_id
		, users.user_id
) as xo_join_user_team_members on xo_join_user_team_members.user_team_team_id = teams.team_id
`

const teamTableMembersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_team_members.__users
		)) filter (where xo_join_user_team_members.__users_user_id is not null), '{}') as user_team_members`

const teamTableMembersGroupBySQL = `teams.team_id, teams.team_id`

// TeamUpdateParams represents update params for 'public.teams'.
type TeamUpdateParams struct {
	Description *string    `json:"description" nullable:"false"` // description
	Name        *string    `json:"name" nullable:"false"`        // name
	ProjectID   *ProjectID `json:"-"`                            // project_id
}

// SetUpdateParams updates public.teams struct fields with the specified params.
func (t *Team) SetUpdateParams(params *TeamUpdateParams) {
	if params.Description != nil {
		t.Description = *params.Description
	}
	if params.Name != nil {
		t.Name = *params.Name
	}
	if params.ProjectID != nil {
		t.ProjectID = *params.ProjectID
	}
}

// Insert inserts the Team to the database.
func (t *Team) Insert(ctx context.Context, db DB) (*Team, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.teams (
	description, name, project_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, t.Description, t.Name, t.ProjectID)

	rows, err := db.Query(ctx, sqlstr, t.Description, t.Name, t.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Insert/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	newt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}

	*t = newt

	return t, nil
}

// Update updates a Team in the database.
func (t *Team) Update(ctx context.Context, db DB) (*Team, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.teams SET 
	description = $1, name = $2, project_id = $3 
	WHERE team_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, t.Description, t.Name, t.ProjectID, t.TeamID)

	rows, err := db.Query(ctx, sqlstr, t.Description, t.Name, t.ProjectID, t.TeamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Update/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	newt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}
	*t = newt

	return t, nil
}

// Upsert upserts a Team in the database.
// Requires appropriate PK(s) to be set beforehand.
func (t *Team) Upsert(ctx context.Context, db DB, params *TeamCreateParams) (*Team, error) {
	var err error

	t.Description = params.Description
	t.Name = params.Name
	t.ProjectID = params.ProjectID

	t, err = t.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertTeam/Insert: %w", &XoError{Entity: "Team", Err: err})
			}
			t, err = t.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertTeam/Update: %w", &XoError{Entity: "Team", Err: err})
			}
		}
	}

	return t, err
}

// Delete deletes the Team from the database.
func (t *Team) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.teams 
	WHERE team_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, t.TeamID); err != nil {
		return logerror(err)
	}
	return nil
}

// TeamPaginated returns a cursor-paginated list of Team.
// At least one cursor is required.
func TeamPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...TeamSelectConfigOption) ([]Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {

		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := EntityFields[TableEntityTeam][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("Team/Paginated/cursor: %w", &XoError{Entity: "Team", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("teams.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("Team/Paginated/orderBy: %w", &XoError{Entity: "Team", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.Project {
		selectClauses = append(selectClauses, teamTableProjectSelectSQL)
		joinClauses = append(joinClauses, teamTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, teamTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, teamTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, teamTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, teamTableTimeEntriesGroupBySQL)
	}

	if c.joins.Members {
		selectClauses = append(selectClauses, teamTableMembersSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* TeamPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Team", Err: err}))
	}
	return res, nil
}

// TeamByNameProjectID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamByNameProjectID(ctx context.Context, db DB, name string, projectID ProjectID, opts ...TeamSelectConfigOption) (*Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, teamTableProjectSelectSQL)
		joinClauses = append(joinClauses, teamTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, teamTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, teamTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, teamTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, teamTableTimeEntriesGroupBySQL)
	}

	if c.joins.Members {
		selectClauses = append(selectClauses, teamTableMembersSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.name = $1 AND teams.project_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* TeamByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}

	return &t, nil
}

// TeamsByName retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamsByName(ctx context.Context, db DB, name string, opts ...TeamSelectConfigOption) ([]Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, teamTableProjectSelectSQL)
		joinClauses = append(joinClauses, teamTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, teamTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, teamTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, teamTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, teamTableTimeEntriesGroupBySQL)
	}

	if c.joins.Members {
		selectClauses = append(selectClauses, teamTableMembersSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.name = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* TeamsByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Team", Err: err}))
	}
	return res, nil
}

// TeamsByProjectID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamsByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...TeamSelectConfigOption) ([]Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, teamTableProjectSelectSQL)
		joinClauses = append(joinClauses, teamTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, teamTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, teamTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, teamTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, teamTableTimeEntriesGroupBySQL)
	}

	if c.joins.Members {
		selectClauses = append(selectClauses, teamTableMembersSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.project_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* TeamsByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Team", Err: err}))
	}
	return res, nil
}

// TeamByTeamID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_pkey'.
func TeamByTeamID(ctx context.Context, db DB, teamID TeamID, opts ...TeamSelectConfigOption) (*Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, teamTableProjectSelectSQL)
		joinClauses = append(joinClauses, teamTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, teamTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, teamTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, teamTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, teamTableTimeEntriesGroupBySQL)
	}

	if c.joins.Members {
		selectClauses = append(selectClauses, teamTableMembersSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.team_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* TeamByTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{teamID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}

	return &t, nil
}

// FKProject_ProjectID returns the Project associated with the Team's (ProjectID).
//
// Generated from foreign key 'teams_project_id_fkey'.
func (t *Team) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, t.ProjectID)
}
