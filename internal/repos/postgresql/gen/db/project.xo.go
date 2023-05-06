package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Project represents a row from 'public.projects'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type Project struct {
	ProjectID          int                  `json:"projectID" db:"project_id" required:"true"`                                              // project_id
	Name               models.Project       `json:"name" db:"name" required:"true" ref:"#/components/schemas/Project"`                      // name
	Description        string               `json:"description" db:"description" required:"true"`                                           // description
	WorkItemsTableName string               `json:"-" db:"work_items_table_name"`                                                           // work_items_table_name
	BoardConfig        models.ProjectConfig `json:"boardConfig" db:"board_config" required:"true" ref:"#/components/schemas/ProjectConfig"` // board_config
	CreatedAt          time.Time            `json:"createdAt" db:"created_at" required:"true"`                                              // created_at
	UpdatedAt          time.Time            `json:"updatedAt" db:"updated_at" required:"true"`                                              // updated_at

	ActivitiesJoin    *[]Activity     `json:"-" db:"activities" openapi-go:"ignore"`      // M2O
	KanbanStepsJoin   *[]KanbanStep   `json:"-" db:"kanban_steps" openapi-go:"ignore"`    // M2O
	TeamsJoin         *[]Team         `json:"-" db:"teams" openapi-go:"ignore"`           // M2O
	WorkItemTagsJoin  *[]WorkItemTag  `json:"-" db:"work_item_tags" openapi-go:"ignore"`  // M2O
	WorkItemTypesJoin *[]WorkItemType `json:"-" db:"work_item_types" openapi-go:"ignore"` // M2O

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

// ProjectUpdateParams represents update params for 'public.projects'
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

type ProjectOrderBy = string

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
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type ProjectJoins struct {
	Activities    bool
	KanbanSteps   bool
	Teams         bool
	WorkItemTags  bool
	WorkItemTypes bool
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

// Insert inserts the Project to the database.
func (p *Project) Insert(ctx context.Context, db DB) (*Project, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.projects (` +
		`name, description, work_items_table_name, board_config` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/db.Query: %w", err))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/pgx.CollectOneRow: %w", err))
	}

	*p = newp

	return p, nil
}

// Update updates a Project in the database.
func (p *Project) Update(ctx context.Context, db DB) (*Project, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.projects SET ` +
		`name = $1, description = $2, work_items_table_name = $3, board_config = $4 ` +
		`WHERE project_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig, p.CreatedAt, p.UpdatedAt, p.ProjectID)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.BoardConfig, p.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/db.Query: %w", err))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/pgx.CollectOneRow: %w", err))
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			p, err = p.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return p, err
}

// Delete deletes the Project from the database.
func (p *Project) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.projects ` +
		`WHERE project_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID); err != nil {
		return logerror(err)
	}
	return nil
}

// ProjectPaginatedByProjectID returns a cursor-paginated list of Project.
func ProjectPaginatedByProjectID(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) ([]Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then COALESCE(joined_activities.activities, '{}') end) as activities,
(case when $2::boolean = true then COALESCE(joined_kanban_steps.kanban_steps, '{}') end) as kanban_steps,
(case when $3::boolean = true then COALESCE(joined_teams.teams, '{}') end) as teams,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags,
(case when $5::boolean = true then COALESCE(joined_work_item_types.work_item_types, '{}') end) as work_item_types ` +
		`FROM public.projects ` +
		`-- M2O join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
  group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
  group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
  group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
  group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.project_id > $6` +
		` ORDER BY 
		project_id DESC `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ProjectByName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_name_key'.
func ProjectByName(ctx context.Context, db DB, name models.Project, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then COALESCE(joined_activities.activities, '{}') end) as activities,
(case when $2::boolean = true then COALESCE(joined_kanban_steps.kanban_steps, '{}') end) as kanban_steps,
(case when $3::boolean = true then COALESCE(joined_teams.teams, '{}') end) as teams,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags,
(case when $5::boolean = true then COALESCE(joined_work_item_types.work_item_types, '{}') end) as work_item_types ` +
		`FROM public.projects ` +
		`-- M2O join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
  group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
  group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
  group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
  group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.name = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTags, c.joins.WorkItemTypes, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/db.Query: %w", err))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/pgx.CollectOneRow: %w", err))
	}

	return &p, nil
}

// ProjectByProjectID retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_pkey'.
func ProjectByProjectID(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then COALESCE(joined_activities.activities, '{}') end) as activities,
(case when $2::boolean = true then COALESCE(joined_kanban_steps.kanban_steps, '{}') end) as kanban_steps,
(case when $3::boolean = true then COALESCE(joined_teams.teams, '{}') end) as teams,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags,
(case when $5::boolean = true then COALESCE(joined_work_item_types.work_item_types, '{}') end) as work_item_types ` +
		`FROM public.projects ` +
		`-- M2O join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
  group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
  group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
  group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
  group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.project_id = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTags, c.joins.WorkItemTypes, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/db.Query: %w", err))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/pgx.CollectOneRow: %w", err))
	}

	return &p, nil
}

// ProjectByWorkItemsTableName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_work_items_table_name_key'.
func ProjectByWorkItemsTableName(ctx context.Context, db DB, workItemsTableName string, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then COALESCE(joined_activities.activities, '{}') end) as activities,
(case when $2::boolean = true then COALESCE(joined_kanban_steps.kanban_steps, '{}') end) as kanban_steps,
(case when $3::boolean = true then COALESCE(joined_teams.teams, '{}') end) as teams,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags,
(case when $5::boolean = true then COALESCE(joined_work_item_types.work_item_types, '{}') end) as work_item_types ` +
		`FROM public.projects ` +
		`-- M2O join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
  group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- M2O join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
  group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- M2O join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
  group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- M2O join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- M2O join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
  group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.work_items_table_name = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemsTableName)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTags, c.joins.WorkItemTypes, workItemsTableName)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/db.Query: %w", err))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByWorkItemsTableName/pgx.CollectOneRow: %w", err))
	}

	return &p, nil
}
