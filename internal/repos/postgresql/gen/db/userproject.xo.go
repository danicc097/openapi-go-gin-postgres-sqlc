package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// UserProject represents a row from 'public.user_project'.
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
type UserProject struct {
	ProjectID ProjectID `json:"projectID" db:"project_id" required:"true" nullable:"false"` // project_id
	Member    UserID    `json:"member" db:"member" required:"true" nullable:"false"`        // member

	MemberProjectsJoin *[]Project `json:"-" db:"user_project_projects" openapi-go:"ignore"` // M2M user_project
	ProjectMembersJoin *[]User    `json:"-" db:"user_project_members" openapi-go:"ignore"`  // M2M user_project

}

// UserProjectCreateParams represents insert params for 'public.user_project'.
type UserProjectCreateParams struct {
	Member    UserID    `json:"member" required:"true" nullable:"false"`    // member
	ProjectID ProjectID `json:"projectID" required:"true" nullable:"false"` // project_id
}

// CreateUserProject creates a new UserProject in the database with the given params.
func CreateUserProject(ctx context.Context, db DB, params *UserProjectCreateParams) (*UserProject, error) {
	up := &UserProject{
		Member:    params.Member,
		ProjectID: params.ProjectID,
	}

	return up.Insert(ctx, db)
}

type UserProjectSelectConfig struct {
	limit   string
	orderBy string
	joins   UserProjectJoins
	filters map[string][]any
	having  map[string][]any
}
type UserProjectSelectConfigOption func(*UserProjectSelectConfig)

// WithUserProjectLimit limits row selection.
func WithUserProjectLimit(limit int) UserProjectSelectConfigOption {
	return func(s *UserProjectSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserProjectOrderBy string

const ()

type UserProjectJoins struct {
	ProjectsMember bool // M2M user_project
	MembersProject bool // M2M user_project
}

// WithUserProjectJoin joins with the given tables.
func WithUserProjectJoin(joins UserProjectJoins) UserProjectSelectConfigOption {
	return func(s *UserProjectSelectConfig) {
		s.joins = UserProjectJoins{
			ProjectsMember: s.joins.ProjectsMember || joins.ProjectsMember,
			MembersProject: s.joins.MembersProject || joins.MembersProject,
		}
	}
}

// WithUserProjectFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserProjectFilters(filters map[string][]any) UserProjectSelectConfigOption {
	return func(s *UserProjectSelectConfig) {
		s.filters = filters
	}
}

// WithUserProjectHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithUserProjectHavingClause(conditions map[string][]any) UserProjectSelectConfigOption {
	return func(s *UserProjectSelectConfig) {
		s.having = conditions
	}
}

const userProjectTableProjectsMemberJoinSQL = `-- M2M join generated from "user_project_project_id_fkey"
left join (
	select
		user_project.member as user_project_member
		, projects.project_id as __projects_project_id
		, row(projects.*) as __projects
	from
		user_project
	join projects on projects.project_id = user_project.project_id
	group by
		user_project_member
		, projects.project_id
) as xo_join_user_project_projects on xo_join_user_project_projects.user_project_member = user_project.project_id
`

const userProjectTableProjectsMemberSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_project_projects.__projects
		)) filter (where xo_join_user_project_projects.__projects_project_id is not null), '{}') as user_project_projects`

const userProjectTableProjectsMemberGroupBySQL = `user_project.project_id, user_project.project_id, user_project.member`

const userProjectTableMembersProjectJoinSQL = `-- M2M join generated from "user_project_member_fkey"
left join (
	select
		user_project.project_id as user_project_project_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		user_project
	join users on users.user_id = user_project.member
	group by
		user_project_project_id
		, users.user_id
) as xo_join_user_project_members on xo_join_user_project_members.user_project_project_id = user_project.member
`

const userProjectTableMembersProjectSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_project_members.__users
		)) filter (where xo_join_user_project_members.__users_user_id is not null), '{}') as user_project_members`

const userProjectTableMembersProjectGroupBySQL = `user_project.member, user_project.project_id, user_project.member`

// UserProjectUpdateParams represents update params for 'public.user_project'.
type UserProjectUpdateParams struct {
	Member    *UserID    `json:"member" nullable:"false"`    // member
	ProjectID *ProjectID `json:"projectID" nullable:"false"` // project_id
}

// SetUpdateParams updates public.user_project struct fields with the specified params.
func (up *UserProject) SetUpdateParams(params *UserProjectUpdateParams) {
	if params.Member != nil {
		up.Member = *params.Member
	}
	if params.ProjectID != nil {
		up.ProjectID = *params.ProjectID
	}
}

// Insert inserts the UserProject to the database.
func (up *UserProject) Insert(ctx context.Context, db DB) (*UserProject, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.user_project (
	member, project_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, up.Member, up.ProjectID)
	rows, err := db.Query(ctx, sqlstr, up.Member, up.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/Insert/db.Query: %w", &XoError{Entity: "User project", Err: err}))
	}
	newup, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserProject])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User project", Err: err}))
	}
	*up = newup

	return up, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key or generated fields

// Delete deletes the UserProject from the database.
func (up *UserProject) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.user_project
	WHERE project_id = $1 AND member = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, up.ProjectID, up.Member); err != nil {
		return logerror(err)
	}
	return nil
}

// UserProjectsByMember retrieves a row from 'public.user_project' as a UserProject.
//
// Generated from index 'user_project_member_idx'.
func UserProjectsByMember(ctx context.Context, db DB, member UserID, opts ...UserProjectSelectConfigOption) ([]UserProject, error) {
	c := &UserProjectSelectConfig{joins: UserProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.ProjectsMember {
		selectClauses = append(selectClauses, userProjectTableProjectsMemberSelectSQL)
		joinClauses = append(joinClauses, userProjectTableProjectsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableProjectsMemberGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, userProjectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, userProjectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableMembersProjectGroupBySQL)
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
	user_project.member,
	user_project.project_id %s
	 FROM public.user_project %s
	 WHERE user_project.member = $1
	 %s   %s
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserProjectsByMember */\n" + sqlstr

	// run
	// logf(sqlstr, member)
	rows, err := db.Query(ctx, sqlstr, append([]any{member}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByMember/Query: %w", &XoError{Entity: "User project", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserProject])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByMember/pgx.CollectRows: %w", &XoError{Entity: "User project", Err: err}))
	}
	return res, nil
}

// UserProjectByMemberProjectID retrieves a row from 'public.user_project' as a UserProject.
//
// Generated from index 'user_project_pkey'.
func UserProjectByMemberProjectID(ctx context.Context, db DB, member UserID, projectID ProjectID, opts ...UserProjectSelectConfigOption) (*UserProject, error) {
	c := &UserProjectSelectConfig{joins: UserProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.ProjectsMember {
		selectClauses = append(selectClauses, userProjectTableProjectsMemberSelectSQL)
		joinClauses = append(joinClauses, userProjectTableProjectsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableProjectsMemberGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, userProjectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, userProjectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableMembersProjectGroupBySQL)
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
	user_project.member,
	user_project.project_id %s
	 FROM public.user_project %s
	 WHERE user_project.member = $1 AND user_project.project_id = $2
	 %s   %s
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserProjectByMemberProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, member, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{member, projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_project/UserProjectByMemberProjectID/db.Query: %w", &XoError{Entity: "User project", Err: err}))
	}
	up, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserProject])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_project/UserProjectByMemberProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "User project", Err: err}))
	}

	return &up, nil
}

// UserProjectsByProjectID retrieves a row from 'public.user_project' as a UserProject.
//
// Generated from index 'user_project_pkey'.
func UserProjectsByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...UserProjectSelectConfigOption) ([]UserProject, error) {
	c := &UserProjectSelectConfig{joins: UserProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.ProjectsMember {
		selectClauses = append(selectClauses, userProjectTableProjectsMemberSelectSQL)
		joinClauses = append(joinClauses, userProjectTableProjectsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableProjectsMemberGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, userProjectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, userProjectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableMembersProjectGroupBySQL)
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
	user_project.member,
	user_project.project_id %s
	 FROM public.user_project %s
	 WHERE user_project.project_id = $1
	 %s   %s
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserProjectsByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByMemberProjectID/Query: %w", &XoError{Entity: "User project", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserProject])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByMemberProjectID/pgx.CollectRows: %w", &XoError{Entity: "User project", Err: err}))
	}
	return res, nil
}

// UserProjectsByProjectIDMember retrieves a row from 'public.user_project' as a UserProject.
//
// Generated from index 'user_project_project_id_member_idx'.
func UserProjectsByProjectIDMember(ctx context.Context, db DB, projectID ProjectID, member UserID, opts ...UserProjectSelectConfigOption) ([]UserProject, error) {
	c := &UserProjectSelectConfig{joins: UserProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.ProjectsMember {
		selectClauses = append(selectClauses, userProjectTableProjectsMemberSelectSQL)
		joinClauses = append(joinClauses, userProjectTableProjectsMemberJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableProjectsMemberGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, userProjectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, userProjectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, userProjectTableMembersProjectGroupBySQL)
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
	user_project.member,
	user_project.project_id %s
	 FROM public.user_project %s
	 WHERE user_project.project_id = $1 AND user_project.member = $2
	 %s   %s
  %s
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserProjectsByProjectIDMember */\n" + sqlstr

	// run
	// logf(sqlstr, projectID, member)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID, member}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByProjectIDMember/Query: %w", &XoError{Entity: "User project", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserProject])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserProject/UserProjectByProjectIDMember/pgx.CollectRows: %w", &XoError{Entity: "User project", Err: err}))
	}
	return res, nil
}

// FKUser_Member returns the User associated with the UserProject's (Member).
//
// Generated from foreign key 'user_project_member_fkey'.
func (up *UserProject) FKUser_Member(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, up.Member)
}

// FKProject_ProjectID returns the Project associated with the UserProject's (ProjectID).
//
// Generated from foreign key 'user_project_project_id_fkey'.
func (up *UserProject) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, up.ProjectID)
}
