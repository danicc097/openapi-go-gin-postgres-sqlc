package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// ProjectPublic represents fields that may be exposed from 'public.projects'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type ProjectPublic struct {
	ProjectID   int       `json:"projectID" required:"true"`   // project_id
	Name        string    `json:"name" required:"true"`        // name
	Description string    `json:"description" required:"true"` // description
	CreatedAt   time.Time `json:"createdAt" required:"true"`   // created_at
	UpdatedAt   time.Time `json:"updatedAt" required:"true"`   // updated_at
}

// Project represents a row from 'public.projects'.
type Project struct {
	ProjectID   int       `json:"project_id" db:"project_id"`   // project_id
	Name        string    `json:"name" db:"name"`               // name
	Description string    `json:"description" db:"description"` // description
	CreatedAt   time.Time `json:"created_at" db:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`   // updated_at

	Activities    *[]Activity     `json:"activities" db:"activities"`           // O2M
	KanbanSteps   *[]KanbanStep   `json:"kanban_steps" db:"kanban_steps"`       // O2M
	Teams         *[]Team         `json:"teams" db:"teams"`                     // O2M
	WorkItemTypes *[]WorkItemType `json:"work_item_types" db:"work_item_types"` // O2M
	// xo fields
	_exists, _deleted bool
}

func (x *Project) ToPublic() ProjectPublic {
	return ProjectPublic{
		ProjectID: x.ProjectID, Name: x.Name, Description: x.Description, CreatedAt: x.CreatedAt, UpdatedAt: x.UpdatedAt,
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
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
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type ProjectJoins struct {
	Activities    bool
	KanbanSteps   bool
	Teams         bool
	WorkItemTypes bool
}

// WithProjectJoin orders results by the given columns.
func WithProjectJoin(joins ProjectJoins) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Project exists in the database.
func (p *Project) Exists() bool {
	return p._exists
}

// Deleted returns true when the Project has been marked for deletion from
// the database.
func (p *Project) Deleted() bool {
	return p._deleted
}

// Insert inserts the Project to the database.
func (p *Project) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.projects (` +
		`name, description` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING project_id, created_at, updated_at `
	// run
	logf(sqlstr, p.Name, p.Description)
	if err := db.QueryRow(ctx, sqlstr, p.Name, p.Description).Scan(&p.ProjectID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a Project in the database.
func (p *Project) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.projects SET ` +
		`name = $1, description = $2 ` +
		`WHERE project_id = $3 ` +
		`RETURNING project_id, created_at, updated_at `
	// run
	logf(sqlstr, p.Name, p.Description, p.CreatedAt, p.UpdatedAt, p.ProjectID)
	if err := db.QueryRow(ctx, sqlstr, p.Name, p.Description, p.ProjectID).Scan(&p.ProjectID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Project to the database.
func (p *Project) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Project.
func (p *Project) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.projects (` +
		`project_id, name, description` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (project_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description  `
	// run
	logf(sqlstr, p.ProjectID, p.Name, p.Description)
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID, p.Name, p.Description); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Project from the database.
func (p *Project) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.projects ` +
		`WHERE project_id = $1 `
	// run
	logf(sqlstr, p.ProjectID)
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProjectByName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_name_key'.
func ProjectByName(ctx context.Context, db DB, name string, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then joined_activities.activities end)::jsonb as activities,
(case when $2::boolean = true then joined_kanban_steps.kanban_steps end)::jsonb as kanban_steps,
(case when $3::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $4::boolean = true then joined_work_item_types.work_item_types end)::jsonb as work_item_types ` +
		`FROM public.projects ` +
		`-- O2M join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , json_agg(activities.*) as activities
  from
    activities
   group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- O2M join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , json_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
   group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- O2M join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , json_agg(teams.*) as teams
  from
    teams
   group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- O2M join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , json_agg(work_item_types.*) as work_item_types
  from
    work_item_types
   group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.name = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name)
	p := Project{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTypes, name).Scan(&p.ProjectID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt, &p.Activities, &p.KanbanSteps, &p.Teams, &p.WorkItemTypes); err != nil {
		return nil, logerror(err)
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
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then joined_activities.activities end)::jsonb as activities,
(case when $2::boolean = true then joined_kanban_steps.kanban_steps end)::jsonb as kanban_steps,
(case when $3::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $4::boolean = true then joined_work_item_types.work_item_types end)::jsonb as work_item_types ` +
		`FROM public.projects ` +
		`-- O2M join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , json_agg(activities.*) as activities
  from
    activities
   group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- O2M join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , json_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
   group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- O2M join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , json_agg(teams.*) as teams
  from
    teams
   group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- O2M join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , json_agg(work_item_types.*) as work_item_types
  from
    work_item_types
   group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.project_id = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, projectID)
	p := Project{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTypes, projectID).Scan(&p.ProjectID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt, &p.Activities, &p.KanbanSteps, &p.Teams, &p.WorkItemTypes); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
