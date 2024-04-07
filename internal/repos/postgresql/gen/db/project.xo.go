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

	ActivitiesJoin    *[]Activity     `json:"-" db:"activities" openapi-go:"ignore"`           // M2O projects
	KanbanStepsJoin   *[]KanbanStep   `json:"-" db:"kanban_steps" openapi-go:"ignore"`         // M2O projects
	TeamsJoin         *[]Team         `json:"-" db:"teams" openapi-go:"ignore"`                // M2O projects
	MembersJoin       *[]User         `json:"-" db:"user_project_members" openapi-go:"ignore"` // M2M user_project
	WorkItemTagsJoin  *[]WorkItemTag  `json:"-" db:"work_item_tags" openapi-go:"ignore"`       // M2O projects
	WorkItemTypesJoin *[]WorkItemType `json:"-" db:"work_item_types" openapi-go:"ignore"`      // M2O projects

}

// ProjectCreateParams represents insert params for 'public.projects'.
type ProjectCreateParams struct {
	BoardConfig        models.ProjectConfig `json:"boardConfig" required:"true" nullable:"false" ref:"#/components/schemas/ProjectConfig"` // board_config
	Description        string               `json:"description" required:"true" nullable:"false"`                                          // description
	Name               models.Project       `json:"name" required:"true" nullable:"false" ref:"#/components/schemas/Project"`              // name
	WorkItemsTableName string               `json:"-" nullable:"false"`                                                                    // work_items_table_name
}

// ProjectParams represents common params for both insert and update of 'public.projects'.
type ProjectParams interface {
	GetBoardConfig() *models.ProjectConfig
	GetDescription() *string
	GetName() *models.Project
	GetWorkItemsTableName() *string
}

func (p ProjectCreateParams) GetBoardConfig() *models.ProjectConfig {
	x := p.BoardConfig
	return &x
}
func (p ProjectUpdateParams) GetBoardConfig() *models.ProjectConfig {
	return p.BoardConfig
}

func (p ProjectCreateParams) GetDescription() *string {
	x := p.Description
	return &x
}
func (p ProjectUpdateParams) GetDescription() *string {
	return p.Description
}

func (p ProjectCreateParams) GetName() *models.Project {
	x := p.Name
	return &x
}
func (p ProjectUpdateParams) GetName() *models.Project {
	return p.Name
}

func (p ProjectCreateParams) GetWorkItemsTableName() *string {
	x := p.WorkItemsTableName
	return &x
}
func (p ProjectUpdateParams) GetWorkItemsTableName() *string {
	return p.WorkItemsTableName
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
	orderBy map[string]models.Direction
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

// WithProjectOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithProjectOrderBy(rows map[string]*models.Direction) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		te := EntityFields[TableEntityProject]
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

type ProjectJoins struct {
	Activities    bool `json:"activities" required:"true" nullable:"false"`    // M2O activities
	KanbanSteps   bool `json:"kanbanSteps" required:"true" nullable:"false"`   // M2O kanban_steps
	Teams         bool `json:"teams" required:"true" nullable:"false"`         // M2O teams
	Members       bool `json:"members" required:"true" nullable:"false"`       // M2M user_project
	WorkItemTags  bool `json:"workItemTags" required:"true" nullable:"false"`  // M2O work_item_tags
	WorkItemTypes bool `json:"workItemTypes" required:"true" nullable:"false"` // M2O work_item_types
}

// WithProjectJoin joins with the given tables.
func WithProjectJoin(joins ProjectJoins) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.joins = ProjectJoins{
			Activities:    s.joins.Activities || joins.Activities,
			KanbanSteps:   s.joins.KanbanSteps || joins.KanbanSteps,
			Teams:         s.joins.Teams || joins.Teams,
			Members:       s.joins.Members || joins.Members,
			WorkItemTags:  s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItemTypes: s.joins.WorkItemTypes || joins.WorkItemTypes,
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
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
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
    , row(activities.*) as __activities
  from
    activities
  group by
	  activities_project_id, activities.activity_id
) as xo_join_activities on xo_join_activities.activities_project_id = projects.project_id
`

const projectTableActivitiesSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_activities.__activities)) filter (where xo_join_activities.activities_project_id is not null), '{}') as activities`

const projectTableActivitiesGroupBySQL = `projects.project_id`

const projectTableKanbanStepsJoinSQL = `-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , row(kanban_steps.*) as __kanban_steps
  from
    kanban_steps
  group by
	  kanban_steps_project_id, kanban_steps.kanban_step_id
) as xo_join_kanban_steps on xo_join_kanban_steps.kanban_steps_project_id = projects.project_id
`

const projectTableKanbanStepsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_kanban_steps.__kanban_steps)) filter (where xo_join_kanban_steps.kanban_steps_project_id is not null), '{}') as kanban_steps`

const projectTableKanbanStepsGroupBySQL = `projects.project_id`

const projectTableTeamsJoinSQL = `-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , row(teams.*) as __teams
  from
    teams
  group by
	  teams_project_id, teams.team_id
) as xo_join_teams on xo_join_teams.teams_project_id = projects.project_id
`

const projectTableTeamsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_teams.__teams)) filter (where xo_join_teams.teams_project_id is not null), '{}') as teams`

const projectTableTeamsGroupBySQL = `projects.project_id`

const projectTableMembersJoinSQL = `-- M2M join generated from "user_project_member_fkey"
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
) as xo_join_user_project_members on xo_join_user_project_members.user_project_project_id = projects.project_id
`

const projectTableMembersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_project_members.__users
		)) filter (where xo_join_user_project_members.__users_user_id is not null), '{}') as user_project_members`

const projectTableMembersGroupBySQL = `projects.project_id, projects.project_id`

const projectTableWorkItemTagsJoinSQL = `-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , row(work_item_tags.*) as __work_item_tags
  from
    work_item_tags
  group by
	  work_item_tags_project_id, work_item_tags.work_item_tag_id
) as xo_join_work_item_tags on xo_join_work_item_tags.work_item_tags_project_id = projects.project_id
`

const projectTableWorkItemTagsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_work_item_tags.__work_item_tags)) filter (where xo_join_work_item_tags.work_item_tags_project_id is not null), '{}') as work_item_tags`

const projectTableWorkItemTagsGroupBySQL = `projects.project_id`

const projectTableWorkItemTypesJoinSQL = `-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , row(work_item_types.*) as __work_item_types
  from
    work_item_types
  group by
	  work_item_types_project_id, work_item_types.work_item_type_id
) as xo_join_work_item_types on xo_join_work_item_types.work_item_types_project_id = projects.project_id
`

const projectTableWorkItemTypesSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_work_item_types.__work_item_types)) filter (where xo_join_work_item_types.work_item_types_project_id is not null), '{}') as work_item_types`

const projectTableWorkItemTypesGroupBySQL = `projects.project_id`

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
				return nil, fmt.Errorf("UpsertProject/Insert: %w", &XoError{Entity: "Project", Err: err})
			}
			p, err = p.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertProject/Update: %w", &XoError{Entity: "Project", Err: err})
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

// ProjectPaginated returns a cursor-paginated list of Project.
// At least one cursor is required.
func ProjectPaginated(ctx context.Context, db DB, cursors []Cursor, opts ...ProjectSelectConfigOption) ([]Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		field, ok := EntityFields[TableEntityProject][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("Project/Paginated/cursor: %w", &XoError{Entity: "Project", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("projects.%s %s $i", field.Db, op)] = []any{cursor.Value}
		c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts
	}

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
		return nil, logerror(fmt.Errorf("Project/Paginated/orderBy: %w", &XoError{Entity: "Project", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.Members {
		selectClauses = append(selectClauses, projectTableMembersSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersGroupBySQL)
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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ProjectPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
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

	if c.joins.Members {
		selectClauses = append(selectClauses, projectTableMembersSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersGroupBySQL)
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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
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

	if c.joins.Members {
		selectClauses = append(selectClauses, projectTableMembersSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersGroupBySQL)
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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
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

	if c.joins.Members {
		selectClauses = append(selectClauses, projectTableMembersSelectSQL)
		joinClauses = append(joinClauses, projectTableMembersJoinSQL)
		groupByClauses = append(groupByClauses, projectTableMembersGroupBySQL)
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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
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
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
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
