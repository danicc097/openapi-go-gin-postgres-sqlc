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

// Project represents a row from 'public.projects'.
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
type Project struct {
	ProjectID          ProjectID            `json:"projectID" db:"project_id" required:"true" nullable:"false"`                                              // project_id
	Name               models.Project       `json:"name" db:"name" required:"true" nullable:"false" ref:"#/components/schemas/Project"`                      // name
	Description        string               `json:"description" db:"description" required:"true" nullable:"false"`                                           // description
	WorkItemsTableName string               `json:"-" db:"work_items_table_name" nullable:"false"`                                                           // work_items_table_name
	BoardConfig        models.ProjectConfig `json:"boardConfig" db:"board_config" required:"true" nullable:"false" ref:"#/components/schemas/ProjectConfig"` // board_config
	CreatedAt          time.Time            `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                                              // created_at
	UpdatedAt          time.Time            `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`                                              // updated_at

	ProjectActivitiesJoin    *[]Activity     `json:"-" db:"activities" openapi-go:"ignore"`           // M2O projects
	ProjectKanbanStepsJoin   *[]KanbanStep   `json:"-" db:"kanban_steps" openapi-go:"ignore"`         // M2O projects
	ProjectTeamsJoin         *[]Team         `json:"-" db:"teams" openapi-go:"ignore"`                // M2O projects
	ProjectMembersJoin       *[]User         `json:"-" db:"user_project_members" openapi-go:"ignore"` // M2M user_project
	ProjectWorkItemTagsJoin  *[]WorkItemTag  `json:"-" db:"work_item_tags" openapi-go:"ignore"`       // M2O projects
	ProjectWorkItemTypesJoin *[]WorkItemType `json:"-" db:"work_item_types" openapi-go:"ignore"`      // M2O projects

}

// ProjectCreateParams represents insert params for 'public.projects'.
type ProjectCreateParams struct {
	BoardConfig        models.ProjectConfig `json:"boardConfig" required:"true" nullable:"false" ref:"#/components/schemas/ProjectConfig"` // board_config
	Description        string               `json:"description" required:"true" nullable:"false"`                                          // description
	Name               models.Project       `json:"name" required:"true" nullable:"false" ref:"#/components/schemas/Project"`              // name
	WorkItemsTableName string               `json:"-" nullable:"false"`                                                                    // work_items_table_name
}

type ProjectID int

// CreateProject creates a new Project in the database with the given params.
func CreateProject(ctx context.Context, db DB, params *ProjectCreateParams) (*Project, error) {
	p := &Project{
		BoardConfig:        params.BoardConfig,
		Description:        params.Description,
		Name:               params.Name,
		WorkItemsTableName: params.WorkItemsTableName,
	}

	return p.Insert(ctx, db)
}

type ProjectSelectConfig struct {
	limit   string
	orderBy string
	joins   ProjectJoins
	filters map[string][]any
	having  map[string][]any
}
type ProjectSelectConfigOption func(*ProjectSelectConfig)

// WithProjectLimit limits row selection.
func WithProjectLimit(limit int) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ProjectOrderBy string

const (
	ProjectCreatedAtDescNullsFirst ProjectOrderBy = " created_at DESC NULLS FIRST "
	ProjectCreatedAtDescNullsLast  ProjectOrderBy = " created_at DESC NULLS LAST "
	ProjectCreatedAtAscNullsFirst  ProjectOrderBy = " created_at ASC NULLS FIRST "
	ProjectCreatedAtAscNullsLast   ProjectOrderBy = " created_at ASC NULLS LAST "
	ProjectUpdatedAtDescNullsFirst ProjectOrderBy = " updated_at DESC NULLS FIRST "
	ProjectUpdatedAtDescNullsLast  ProjectOrderBy = " updated_at DESC NULLS LAST "
	ProjectUpdatedAtAscNullsFirst  ProjectOrderBy = " updated_at ASC NULLS FIRST "
	ProjectUpdatedAtAscNullsLast   ProjectOrderBy = " updated_at ASC NULLS LAST "
)

// WithProjectOrderBy orders results by the given columns.
func WithProjectOrderBy(rows ...ProjectOrderBy) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
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

type ProjectJoins struct {
	Activities     bool // M2O activities
	KanbanSteps    bool // M2O kanban_steps
	Teams          bool // M2O teams
	MembersProject bool // M2M user_project
	WorkItemTags   bool // M2O work_item_tags
	WorkItemTypes  bool // M2O work_item_types
}

// WithProjectJoin joins with the given tables.
func WithProjectJoin(joins ProjectJoins) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.joins = ProjectJoins{
			Activities:     s.joins.Activities || joins.Activities,
			KanbanSteps:    s.joins.KanbanSteps || joins.KanbanSteps,
			Teams:          s.joins.Teams || joins.Teams,
			MembersProject: s.joins.MembersProject || joins.MembersProject,
			WorkItemTags:   s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItemTypes:  s.joins.WorkItemTypes || joins.WorkItemTypes,
		}
	}
}

// WithProjectFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithProjectFilters(filters map[string][]any) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.filters = filters
	}
}

// WithProjectHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithProjectHavingClause(conditions map[string][]any) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.having = conditions
	}
}

const projectTableActivitiesJoinSQL = `-- M2O join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
  group by
        project_id
) as joined_activities on joined_activities.activities_project_id = projects.project_id
`

const projectTableActivitiesSelectSQL = `COALESCE(joined_activities.activities, '{}') as activities`

const projectTableActivitiesGroupBySQL = `joined_activities.activities, projects.project_id`

const projectTableKanbanStepsJoinSQL = `-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
  group by
        project_id
) as joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
`

const projectTableKanbanStepsSelectSQL = `COALESCE(joined_kanban_steps.kanban_steps, '{}') as kanban_steps`

const projectTableKanbanStepsGroupBySQL = `joined_kanban_steps.kanban_steps, projects.project_id`

const projectTableTeamsJoinSQL = `-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
  group by
        project_id
) as joined_teams on joined_teams.teams_project_id = projects.project_id
`

const projectTableTeamsSelectSQL = `COALESCE(joined_teams.teams, '{}') as teams`

const projectTableTeamsGroupBySQL = `joined_teams.teams, projects.project_id`

const projectTableMembersProjectJoinSQL = `-- M2M join generated from "user_project_member_fkey"
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
) as joined_user_project_members on joined_user_project_members.user_project_project_id = projects.project_id
`

const projectTableMembersProjectSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_user_project_members.__users
		)) filter (where joined_user_project_members.__users_user_id is not null), '{}') as user_project_members`

const projectTableMembersProjectGroupBySQL = `projects.project_id, projects.project_id`

const projectTableWorkItemTagsJoinSQL = `-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        project_id
) as joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
`

const projectTableWorkItemTagsSelectSQL = `COALESCE(joined_work_item_tags.work_item_tags, '{}') as work_item_tags`

const projectTableWorkItemTagsGroupBySQL = `joined_work_item_tags.work_item_tags, projects.project_id`

const projectTableWorkItemTypesJoinSQL = `-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
  group by
        project_id
) as joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id
`

const projectTableWorkItemTypesSelectSQL = `COALESCE(joined_work_item_types.work_item_types, '{}') as work_item_types`

const projectTableWorkItemTypesGroupBySQL = `joined_work_item_types.work_item_types, projects.project_id`

// ProjectUpdateParams represents update params for 'public.projects'.
type ProjectUpdateParams struct {
	BoardConfig        *models.ProjectConfig `json:"boardConfig" nullable:"false" ref:"#/components/schemas/ProjectConfig"` // board_config
	Description        *string               `json:"description" nullable:"false"`                                          // description
	Name               *models.Project       `json:"name" nullable:"false" ref:"#/components/schemas/Project"`              // name
	WorkItemsTableName *string               `json:"-" nullable:"false"`                                                    // work_items_table_name
}

// SetUpdateParams updates public.projects struct fields with the specified params.
func (p *Project) SetUpdateParams(params *ProjectUpdateParams) {
	if params.BoardConfig != nil {
		p.BoardConfig = *params.BoardConfig
	}
	if params.Description != nil {
		p.Description = *params.Description
	}
	if params.Name != nil {
		p.Name = *params.Name
	}
	if params.WorkItemsTableName != nil {
		p.WorkItemsTableName = *params.WorkItemsTableName
	}
}

// Insert inserts the Project to the database.
func (p *Project) Insert(ctx context.Context, db DB) (*Project, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.projects (
	board_config, description, name, work_items_table_name
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, p.BoardConfig, p.Description, p.Name, p.WorkItemsTableName)

	rows, err := db.Query(ctx, sqlstr, p.BoardConfig, p.Description, p.Name, p.WorkItemsTableName)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}

	*p = newp

	return p, nil
}

// Update updates a Project in the database.
func (p *Project) Update(ctx context.Context, db DB) (*Project, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.projects SET 
	board_config = $1, description = $2, name = $3, work_items_table_name = $4 
	WHERE project_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, p.BoardConfig, p.Description, p.Name, p.WorkItemsTableName, p.ProjectID)

	rows, err := db.Query(ctx, sqlstr, p.BoardConfig, p.Description, p.Name, p.WorkItemsTableName, p.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}
	*p = newp

	return p, nil
}

// Upsert upserts a Project in the database.
// Requires appropriate PK(s) to be set beforehand.
func (p *Project) Upsert(ctx context.Context, db DB, params *ProjectCreateParams) (*Project, error) {
	var err error

	p.BoardConfig = params.BoardConfig
	p.Description = params.Description
	p.Name = params.Name
	p.WorkItemsTableName = params.WorkItemsTableName

	p, err = p.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Project", Err: err})
			}
			p, err = p.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Project", Err: err})
			}
		}
	}

	return p, err
}

// Delete deletes the Project from the database.
func (p *Project) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.projects 
	WHERE project_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID); err != nil {
		return logerror(err)
	}
	return nil
}

// ProjectPaginatedByProjectID returns a cursor-paginated list of Project.
func ProjectPaginatedByProjectID(ctx context.Context, db DB, projectID ProjectID, direction models.Direction, opts ...ProjectSelectConfigOption) ([]Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Activities {
		selectClauses = append(selectClauses, projectTableActivitiesSelectSQL)
		joinClauses = append(joinClauses, projectTableActivitiesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableActivitiesGroupBySQL)
	}

	if c.joins.KanbanSteps {
		selectClauses = append(selectClauses, projectTableKanbanStepsSelectSQL)
		joinClauses = append(joinClauses, projectTableKanbanStepsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableKanbanStepsGroupBySQL)
	}

	if c.joins.Teams {
		selectClauses = append(selectClauses, projectTableTeamsSelectSQL)
		joinClauses = append(joinClauses, projectTableTeamsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableTeamsGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, projectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersProjectGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, projectTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemTypes {
		selectClauses = append(selectClauses, projectTableWorkItemTypesSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTypesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTypesGroupBySQL)
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
	projects.board_config,
	projects.created_at,
	projects.description,
	projects.name,
	projects.project_id,
	projects.updated_at,
	projects.work_items_table_name %s 
	 FROM public.projects %s 
	 WHERE projects.project_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		project_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ProjectPaginatedByProjectID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Project", Err: err}))
	}
	return res, nil
}

// ProjectByName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_name_key'.
func ProjectByName(ctx context.Context, db DB, name models.Project, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Activities {
		selectClauses = append(selectClauses, projectTableActivitiesSelectSQL)
		joinClauses = append(joinClauses, projectTableActivitiesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableActivitiesGroupBySQL)
	}

	if c.joins.KanbanSteps {
		selectClauses = append(selectClauses, projectTableKanbanStepsSelectSQL)
		joinClauses = append(joinClauses, projectTableKanbanStepsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableKanbanStepsGroupBySQL)
	}

	if c.joins.Teams {
		selectClauses = append(selectClauses, projectTableTeamsSelectSQL)
		joinClauses = append(joinClauses, projectTableTeamsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableTeamsGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, projectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersProjectGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, projectTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemTypes {
		selectClauses = append(selectClauses, projectTableWorkItemTypesSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTypesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTypesGroupBySQL)
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
	projects.board_config,
	projects.created_at,
	projects.description,
	projects.name,
	projects.project_id,
	projects.updated_at,
	projects.work_items_table_name %s 
	 FROM public.projects %s 
	 WHERE projects.name = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}

	return &p, nil
}

// ProjectByProjectID retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_pkey'.
func ProjectByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Activities {
		selectClauses = append(selectClauses, projectTableActivitiesSelectSQL)
		joinClauses = append(joinClauses, projectTableActivitiesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableActivitiesGroupBySQL)
	}

	if c.joins.KanbanSteps {
		selectClauses = append(selectClauses, projectTableKanbanStepsSelectSQL)
		joinClauses = append(joinClauses, projectTableKanbanStepsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableKanbanStepsGroupBySQL)
	}

	if c.joins.Teams {
		selectClauses = append(selectClauses, projectTableTeamsSelectSQL)
		joinClauses = append(joinClauses, projectTableTeamsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableTeamsGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, projectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersProjectGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, projectTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemTypes {
		selectClauses = append(selectClauses, projectTableWorkItemTypesSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTypesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTypesGroupBySQL)
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
	projects.board_config,
	projects.created_at,
	projects.description,
	projects.name,
	projects.project_id,
	projects.updated_at,
	projects.work_items_table_name %s 
	 FROM public.projects %s 
	 WHERE projects.project_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}

	return &p, nil
}

// ProjectByWorkItemsTableName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_work_items_table_name_key'.
func ProjectByWorkItemsTableName(ctx context.Context, db DB, workItemsTableName string, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Activities {
		selectClauses = append(selectClauses, projectTableActivitiesSelectSQL)
		joinClauses = append(joinClauses, projectTableActivitiesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableActivitiesGroupBySQL)
	}

	if c.joins.KanbanSteps {
		selectClauses = append(selectClauses, projectTableKanbanStepsSelectSQL)
		joinClauses = append(joinClauses, projectTableKanbanStepsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableKanbanStepsGroupBySQL)
	}

	if c.joins.Teams {
		selectClauses = append(selectClauses, projectTableTeamsSelectSQL)
		joinClauses = append(joinClauses, projectTableTeamsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableTeamsGroupBySQL)
	}

	if c.joins.MembersProject {
		selectClauses = append(selectClauses, projectTableMembersProjectSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersProjectJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersProjectGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, projectTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemTypes {
		selectClauses = append(selectClauses, projectTableWorkItemTypesSelectSQL)
		joinClauses = append(joinClauses, projectTableWorkItemTypesJoinSQL)
		groupByClauses = append(groupByClauses, projectTableWorkItemTypesGroupBySQL)
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
	projects.board_config,
	projects.created_at,
	projects.description,
	projects.name,
	projects.project_id,
	projects.updated_at,
	projects.work_items_table_name %s 
	 FROM public.projects %s 
	 WHERE projects.work_items_table_name = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByWorkItemsTableName */\n" + sqlstr

	// run
	// logf(sqlstr, workItemsTableName)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemsTableName}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}

	return &p, nil
}
