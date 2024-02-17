package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Team represents a row from 'public.teams'.
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
type Team struct {
	TeamID      TeamID    `json:"teamID" db:"team_id" required:"true" nullable:"false"`          // team_id
	ProjectID   ProjectID `json:"projectID" db:"project_id" required:"true" nullable:"false"`    // project_id
	Name        string    `json:"name" db:"name" required:"true" nullable:"false"`               // name
	Description string    `json:"description" db:"description" required:"true" nullable:"false"` // description
	CreatedAt   time.Time `json:"createdAt" db:"created_at" required:"true" nullable:"false"`    // created_at
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`    // updated_at

	ProjectJoin         *Project     `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)
	TeamTimeEntriesJoin *[]TimeEntry `json:"-" db:"time_entries" openapi-go:"ignore"`       // M2O teams
	TeamMembersJoin     *[]User      `json:"-" db:"user_team_members" openapi-go:"ignore"`  // M2M user_team

}

// TeamCreateParams represents insert params for 'public.teams'.
type TeamCreateParams struct {
	Description string    `json:"description" required:"true" nullable:"false"` // description
	Name        string    `json:"name" required:"true" nullable:"false"`        // name
	ProjectID   ProjectID `json:"-" openapi-go:"ignore"`                        // project_id
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
	orderBy string
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

type TeamOrderBy string

const (
	TeamCreatedAtDescNullsFirst TeamOrderBy = " created_at DESC NULLS FIRST "
	TeamCreatedAtDescNullsLast  TeamOrderBy = " created_at DESC NULLS LAST "
	TeamCreatedAtAscNullsFirst  TeamOrderBy = " created_at ASC NULLS FIRST "
	TeamCreatedAtAscNullsLast   TeamOrderBy = " created_at ASC NULLS LAST "
	TeamUpdatedAtDescNullsFirst TeamOrderBy = " updated_at DESC NULLS FIRST "
	TeamUpdatedAtDescNullsLast  TeamOrderBy = " updated_at DESC NULLS LAST "
	TeamUpdatedAtAscNullsFirst  TeamOrderBy = " updated_at ASC NULLS FIRST "
	TeamUpdatedAtAscNullsLast   TeamOrderBy = " updated_at ASC NULLS LAST "
)

// WithTeamOrderBy orders results by the given columns.
func WithTeamOrderBy(rows ...TeamOrderBy) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		if len(rows) > 0 {
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
		}
	}
}

type TeamJoins struct {
	Project     bool // O2O projects
	TimeEntries bool // M2O time_entries
	MembersTeam bool // M2M user_team
}

// WithTeamJoin joins with the given tables.
func WithTeamJoin(joins TeamJoins) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.joins = TeamJoins{
			Project:     s.joins.Project || joins.Project,
			TimeEntries: s.joins.TimeEntries || joins.TimeEntries,
			MembersTeam: s.joins.MembersTeam || joins.MembersTeam,
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
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        team_id
) as xo_join_time_entries on xo_join_time_entries.time_entries_team_id = teams.team_id
`

const teamTableTimeEntriesSelectSQL = `COALESCE(xo_join_time_entries.time_entries, '{}') as time_entries`

const teamTableTimeEntriesGroupBySQL = `xo_join_time_entries.time_entries, teams.team_id`

const teamTableMembersTeamJoinSQL = `-- M2M join generated from "user_team_member_fkey"
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

const teamTableMembersTeamSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_team_members.__users
		)) filter (where xo_join_user_team_members.__users_user_id is not null), '{}') as user_team_members`

const teamTableMembersTeamGroupBySQL = `teams.team_id, teams.team_id`

// TeamUpdateParams represents update params for 'public.teams'.
type TeamUpdateParams struct {
	Description *string    `json:"description" nullable:"false"` // description
	Name        *string    `json:"name" nullable:"false"`        // name
	ProjectID   *ProjectID `json:"-" openapi-go:"ignore"`        // project_id
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Team", Err: err})
			}
			t, err = t.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Team", Err: err})
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

// TeamPaginatedByTeamID returns a cursor-paginated list of Team.
func TeamPaginatedByTeamID(ctx context.Context, db DB, teamID TeamID, direction models.Direction, opts ...TeamSelectConfigOption) ([]Team, error) {
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.team_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		team_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* TeamPaginatedByTeamID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{teamID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Team", Err: err}))
	}
	return res, nil
}

// TeamPaginatedByProjectID returns a cursor-paginated list of Team.
func TeamPaginatedByProjectID(ctx context.Context, db DB, projectID ProjectID, direction models.Direction, opts ...TeamSelectConfigOption) ([]Team, error) {
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
	teams.created_at,
	teams.description,
	teams.name,
	teams.project_id,
	teams.team_id,
	teams.updated_at %s 
	 FROM public.teams %s 
	 WHERE teams.project_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		project_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* TeamPaginatedByProjectID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, teamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, teamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, teamTableMembersTeamGroupBySQL)
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
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
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
