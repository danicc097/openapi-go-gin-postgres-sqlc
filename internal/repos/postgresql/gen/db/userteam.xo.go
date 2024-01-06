package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// UserTeam represents a row from 'public.user_team'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type UserTeam struct {
	TeamID TeamID `json:"teamID" db:"team_id" required:"true" nullable:"false"` // team_id
	Member UserID `json:"member" db:"member" required:"true" nullable:"false"`  // member

	MemberTeamsJoin *[]Team `json:"-" db:"user_team_teams" openapi-go:"ignore"`   // M2M user_team
	TeamMembersJoin *[]User `json:"-" db:"user_team_members" openapi-go:"ignore"` // M2M user_team

}

// UserTeamCreateParams represents insert params for 'public.user_team'.
type UserTeamCreateParams struct {
	Member UserID `json:"member" required:"true" nullable:"false"` // member
	TeamID TeamID `json:"teamID" required:"true" nullable:"false"` // team_id
}

// CreateUserTeam creates a new UserTeam in the database with the given params.
func CreateUserTeam(ctx context.Context, db DB, params *UserTeamCreateParams) (*UserTeam, error) {
	ut := &UserTeam{
		Member: params.Member,
		TeamID: params.TeamID,
	}

	return ut.Insert(ctx, db)
}

// UserTeamUpdateParams represents update params for 'public.user_team'.
type UserTeamUpdateParams struct {
	Member *UserID `json:"member" nullable:"false"` // member
	TeamID *TeamID `json:"teamID" nullable:"false"` // team_id
}

// SetUpdateParams updates public.user_team struct fields with the specified params.
func (ut *UserTeam) SetUpdateParams(params *UserTeamUpdateParams) {
	if params.Member != nil {
		ut.Member = *params.Member
	}
	if params.TeamID != nil {
		ut.TeamID = *params.TeamID
	}
}

type UserTeamSelectConfig struct {
	limit   string
	orderBy string
	joins   UserTeamJoins
	filters map[string][]any
	having  map[string][]any
}
type UserTeamSelectConfigOption func(*UserTeamSelectConfig)

// WithUserTeamLimit limits row selection.
func WithUserTeamLimit(limit int) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserTeamOrderBy string

const ()

type UserTeamJoins struct {
	TeamsMember bool // M2M user_team
	MembersTeam bool // M2M user_team
}

// WithUserTeamJoin joins with the given tables.
func WithUserTeamJoin(joins UserTeamJoins) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.joins = UserTeamJoins{
			TeamsMember: s.joins.TeamsMember || joins.TeamsMember,
			MembersTeam: s.joins.MembersTeam || joins.MembersTeam,
		}
	}
}

// WithUserTeamFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserTeamFilters(filters map[string][]any) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.filters = filters
	}
}

// WithUserTeamHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithUserTeamHavingClause(conditions map[string][]any) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.having = conditions
	}
}

const userTeamTableTeamsMemberJoinSQL = `-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.member as user_team_member
		, teams.team_id as __teams_team_id
		, row(teams.*) as __teams
	from
		user_team
	join teams on teams.team_id = user_team.team_id
	group by
		user_team_member
		, teams.team_id
) as joined_user_team_teams on joined_user_team_teams.user_team_member = user_team.team_id
`

const userTeamTableTeamsMemberSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_user_team_teams.__teams
		)) filter (where joined_user_team_teams.__teams_team_id is not null), '{}') as user_team_teams`

const userTeamTableTeamsMemberGroupBySQL = `user_team.team_id, user_team.team_id, user_team.member`

const userTeamTableMembersTeamJoinSQL = `-- M2M join generated from "user_team_member_fkey"
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
) as joined_user_team_members on joined_user_team_members.user_team_team_id = user_team.member
`

const userTeamTableMembersTeamSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_user_team_members.__users
		)) filter (where joined_user_team_members.__users_user_id is not null), '{}') as user_team_members`

const userTeamTableMembersTeamGroupBySQL = `user_team.member, user_team.team_id, user_team.member`

// Insert inserts the UserTeam to the database.
func (ut *UserTeam) Insert(ctx context.Context, db DB) (*UserTeam, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.user_team (
	member, team_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, ut.Member, ut.TeamID)
	rows, err := db.Query(ctx, sqlstr, ut.Member, ut.TeamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/Insert/db.Query: %w", &XoError{Entity: "User team", Err: err}))
	}
	newut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User team", Err: err}))
	}
	*ut = newut

	return ut, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the UserTeam from the database.
func (ut *UserTeam) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.user_team 
	WHERE team_id = $1 AND member = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ut.TeamID, ut.Member); err != nil {
		return logerror(err)
	}
	return nil
}

// UserTeamsByMember retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_member_idx'.
func UserTeamsByMember(ctx context.Context, db DB, member UserID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.TeamsMember {
		selectClauses = append(selectClauses, userTeamTableTeamsMemberSelectSQL)
		joinClauses = append(joinClauses, userTeamTableTeamsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableTeamsMemberGroupBySQL)
	}

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, userTeamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, userTeamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableMembersTeamGroupBySQL)
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
	user_team.member,
	user_team.team_id %s 
	 FROM public.user_team %s 
	 WHERE user_team.member = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserTeamsByMember */\n" + sqlstr

	// run
	// logf(sqlstr, member)
	rows, err := db.Query(ctx, sqlstr, append([]any{member}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMember/Query: %w", &XoError{Entity: "User team", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMember/pgx.CollectRows: %w", &XoError{Entity: "User team", Err: err}))
	}
	return res, nil
}

// UserTeamByMemberTeamID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_pkey'.
func UserTeamByMemberTeamID(ctx context.Context, db DB, member UserID, teamID TeamID, opts ...UserTeamSelectConfigOption) (*UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.TeamsMember {
		selectClauses = append(selectClauses, userTeamTableTeamsMemberSelectSQL)
		joinClauses = append(joinClauses, userTeamTableTeamsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableTeamsMemberGroupBySQL)
	}

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, userTeamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, userTeamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableMembersTeamGroupBySQL)
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
	user_team.member,
	user_team.team_id %s 
	 FROM public.user_team %s 
	 WHERE user_team.member = $1 AND user_team.team_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserTeamByMemberTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, member, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{member, teamID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByMemberTeamID/db.Query: %w", &XoError{Entity: "User team", Err: err}))
	}
	ut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByMemberTeamID/pgx.CollectOneRow: %w", &XoError{Entity: "User team", Err: err}))
	}

	return &ut, nil
}

// UserTeamsByTeamID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_pkey'.
func UserTeamsByTeamID(ctx context.Context, db DB, teamID TeamID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.TeamsMember {
		selectClauses = append(selectClauses, userTeamTableTeamsMemberSelectSQL)
		joinClauses = append(joinClauses, userTeamTableTeamsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableTeamsMemberGroupBySQL)
	}

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, userTeamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, userTeamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableMembersTeamGroupBySQL)
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
	user_team.member,
	user_team.team_id %s 
	 FROM public.user_team %s 
	 WHERE user_team.team_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserTeamsByTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{teamID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMemberTeamID/Query: %w", &XoError{Entity: "User team", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMemberTeamID/pgx.CollectRows: %w", &XoError{Entity: "User team", Err: err}))
	}
	return res, nil
}

// UserTeamsByTeamIDMember retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_team_id_member_idx'.
func UserTeamsByTeamIDMember(ctx context.Context, db DB, teamID TeamID, member UserID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.TeamsMember {
		selectClauses = append(selectClauses, userTeamTableTeamsMemberSelectSQL)
		joinClauses = append(joinClauses, userTeamTableTeamsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableTeamsMemberGroupBySQL)
	}

	if c.joins.MembersTeam {
		selectClauses = append(selectClauses, userTeamTableMembersTeamSelectSQL)
		joinClauses = append(joinClauses, userTeamTableMembersTeamJoinSQL)
		groupByClauses = append(groupByClauses, userTeamTableMembersTeamGroupBySQL)
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
	user_team.member,
	user_team.team_id %s 
	 FROM public.user_team %s 
	 WHERE user_team.team_id = $1 AND user_team.member = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserTeamsByTeamIDMember */\n" + sqlstr

	// run
	// logf(sqlstr, teamID, member)
	rows, err := db.Query(ctx, sqlstr, append([]any{teamID, member}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDMember/Query: %w", &XoError{Entity: "User team", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDMember/pgx.CollectRows: %w", &XoError{Entity: "User team", Err: err}))
	}
	return res, nil
}

// FKUser_Member returns the User associated with the UserTeam's (Member).
//
// Generated from foreign key 'user_team_member_fkey'.
func (ut *UserTeam) FKUser_Member(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ut.Member)
}

// FKTeam_TeamID returns the Team associated with the UserTeam's (TeamID).
//
// Generated from foreign key 'user_team_team_id_fkey'.
func (ut *UserTeam) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, ut.TeamID)
}
