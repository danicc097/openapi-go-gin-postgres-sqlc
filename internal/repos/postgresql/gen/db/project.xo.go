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
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Project struct {
	ProjectID          int                  `json:"projectID" db:"project_id" required:"true"`                                              // project_id
	Name               models.Project       `json:"name" db:"name" required:"true" ref:"#/components/schemas/Project"`                      // name
	Description        string               `json:"description" db:"description" required:"true"`                                           // description
	WorkItemsTableName string               `json:"-" db:"work_items_table_name"`                                                           // work_items_table_name
	BoardConfig        models.ProjectConfig `json:"boardConfig" db:"board_config" required:"true" ref:"#/components/schemas/ProjectConfig"` // board_config
	CreatedAt          time.Time            `json:"createdAt" db:"created_at" required:"true"`                                              // created_at
	UpdatedAt          time.Time            `json:"updatedAt" db:"updated_at" required:"true"`                                              // updated_at

	ProjectActivitiesJoin    *[]Activity     `json:"-" db:"activities" openapi-go:"ignore"`      // M2O projects
	ProjectKanbanStepsJoin   *[]KanbanStep   `json:"-" db:"kanban_steps" openapi-go:"ignore"`    // M2O projects
	ProjectTeamsJoin         *[]Team         `json:"-" db:"teams" openapi-go:"ignore"`           // M2O projects
	ProjectWorkItemTagsJoin  *[]WorkItemTag  `json:"-" db:"work_item_tags" openapi-go:"ignore"`  // M2O projects
	ProjectWorkItemTypesJoin *[]WorkItemType `json:"-" db:"work_item_types" openapi-go:"ignore"` // M2O projects

}

// ProjectCreateParams represents insert params for 'public.projects'.
type ProjectCreateParams struct {
	Name               models.Project       `json:"name" required:"true" ref:"#/components/schemas/Project"`              // name
	Description        string               `json:"description" required:"true"`                                          // description
	WorkItemsTableName string               `json:"-"`                                                                    // work_items_table_name
	BoardConfig        models.ProjectConfig `json:"boardConfig" required:"true" ref:"#/components/schemas/ProjectConfig"` // board_config
}

// CreateProject creates a new Project in the database with the given params.
func CreateProject(ctx context.Context, db DB, params *ProjectCreateParams) (*Project, error) {
	p := &Project{
		Name:               params.Name,
		Description:        params.Description,
		WorkItemsTableName: params.WorkItemsTableName,
		BoardConfig:        params.BoardConfig,
	}

	return p.Insert(ctx, db)
}

// ProjectUpdateParams represents update params for 'public.projects'.
type ProjectUpdateParams struct {
	Name               *models.Project       `json:"name" required:"true" ref:"#/components/schemas/Project"`              // name
	Description        *string               `json:"description" required:"true"`                                          // description
	WorkItemsTableName *string               `json:"-"`                                                                    // work_items_table_name
	BoardConfig        *models.ProjectConfig `json:"boardConfig" required:"true" ref:"#/components/schemas/ProjectConfig"` // board_config
}

// SetUpdateParams updates public.projects struct fields with the specified params.
func (p *Project) SetUpdateParams(params *ProjectUpdateParams) {
	if params.Name != nil {
		p.Name = *params.Name
	}
	if params.Description != nil {
		p.Description = *params.Description
	}
	if params.WorkItemsTableName != nil {
		p.WorkItemsTableName = *params.WorkItemsTableName
	}
	if params.BoardConfig != nil {
		p.BoardConfig = *params.BoardConfig
	}
}

type ProjectSelectConfig struct {
	limit   string
	orderBy string
	joins   ProjectJoins
	filters map[string][]any
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
	Activities    bool // M2O activities
	KanbanSteps   bool // M2O kanban_steps
	Teams         bool // M2O teams
	WorkItemTags  bool // M2O work_item_tags
	WorkItemTypes bool // M2O work_item_types
}

// WithProjectJoin joins with the given tables.
func WithProjectJoin(joins ProjectJoins) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.joins = ProjectJoins{
			Activities:    s.joins.Activities || joins.Activities,
			KanbanSteps:   s.joins.KanbanSteps || joins.KanbanSteps,
			Teams:         s.joins.Teams || joins.Teams,
			WorkItemTags:  s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItemTypes: s.joins.WorkItemTypes || joins.WorkItemTypes,
		}
	}
}

// WithProjectFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
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

// Insert inserts the Project to the database.
func (p *Project) Insert(ctx context.Context, db DB) (*Project, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.projects (
	name, description, work_items_table_name, board_config
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig)
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
	name = $1, description = $2, work_items_table_name = $3, board_config = $4 
	WHERE project_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig, p.CreatedAt, p.UpdatedAt, p.ProjectID)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig, p.ProjectID)
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
// Requires appropiate PK(s) to be set beforehand.
func (p *Project) Upsert(ctx context.Context, db DB, params *ProjectCreateParams) (*Project, error) {
	var err error

	p.Name = params.Name
	p.Description = params.Description
	p.WorkItemsTableName = params.WorkItemsTableName
	p.BoardConfig = params.BoardConfig

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

// ProjectPaginatedByProjectIDAsc returns a cursor-paginated list of Project in Asc order.
func ProjectPaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) ([]Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any)}

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
	projects.project_id,
	projects.name,
	projects.description,
	projects.work_items_table_name,
	projects.board_config,
	projects.created_at,
	projects.updated_at %s 
	 FROM public.projects %s 
	 WHERE projects.project_id > $1
	 %s   %s 
  ORDER BY 
		project_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* ProjectPaginatedByProjectIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/Asc/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Project", Err: err}))
	}
	return res, nil
}

// ProjectPaginatedByProjectIDDesc returns a cursor-paginated list of Project in Desc order.
func ProjectPaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) ([]Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any)}

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
	projects.project_id,
	projects.name,
	projects.description,
	projects.work_items_table_name,
	projects.board_config,
	projects.created_at,
	projects.updated_at %s 
	 FROM public.projects %s 
	 WHERE projects.project_id < $1
	 %s   %s 
  ORDER BY 
		project_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* ProjectPaginatedByProjectIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/Desc/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Project", Err: err}))
	}
	return res, nil
}

// ProjectByName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_name_key'.
func ProjectByName(ctx context.Context, db DB, name models.Project, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any)}

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
	projects.project_id,
	projects.name,
	projects.description,
	projects.work_items_table_name,
	projects.board_config,
	projects.created_at,
	projects.updated_at %s 
	 FROM public.projects %s 
	 WHERE projects.name = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
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
func ProjectByProjectID(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any)}

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
	projects.project_id,
	projects.name,
	projects.description,
	projects.work_items_table_name,
	projects.board_config,
	projects.created_at,
	projects.updated_at %s 
	 FROM public.projects %s 
	 WHERE projects.project_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
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
	c := &ProjectSelectConfig{joins: ProjectJoins{}, filters: make(map[string][]any)}

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
	projects.project_id,
	projects.name,
	projects.description,
	projects.work_items_table_name,
	projects.board_config,
	projects.created_at,
	projects.updated_at %s 
	 FROM public.projects %s 
	 WHERE projects.work_items_table_name = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ProjectByWorkItemsTableName */\n" + sqlstr

	// run
	// logf(sqlstr, workItemsTableName)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemsTableName}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/db.Query: %w", &XoError{Entity: "Project", Err: err}))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/pgx.CollectOneRow: %w", &XoError{Entity: "Project", Err: err}))
	}

	return &p, nil
}
