package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgtype"
)

// Project represents a row from 'public.projects'.
type Project struct {
	ProjectID   int          `json:"project_id" db:"project_id"`   // project_id
	Name        string       `json:"name" db:"name"`               // name
	Description string       `json:"description" db:"description"` // description
	Metadata    pgtype.JSONB `json:"metadata" db:"metadata"`       // metadata
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`   // created_at
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
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

type ProjectJoins struct{}

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
		`name, description, metadata` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING project_id, created_at, updated_at `
	// run
	logf(sqlstr, p.Name, p.Description, p.Metadata)
	if err := db.QueryRow(ctx, sqlstr, p.Name, p.Description, p.Metadata).Scan(&p.ProjectID, &p.CreatedAt, &p.UpdatedAt); err != nil {
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
		`name = $1, description = $2, metadata = $3 ` +
		`WHERE project_id = $4 `
	// run
	logf(sqlstr, p.Name, p.Description, p.Metadata, p.CreatedAt, p.UpdatedAt, p.ProjectID)
	if _, err := db.Exec(ctx, sqlstr, p.Name, p.Description, p.Metadata, p.ProjectID); err != nil {
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
		`project_id, name, description, metadata` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (project_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description, metadata = EXCLUDED.metadata  `
	// run
	logf(sqlstr, p.ProjectID, p.Name, p.Description, p.Metadata)
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID, p.Name, p.Description, p.Metadata); err != nil {
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
projects.metadata,
projects.created_at,
projects.updated_at ` +
		`FROM public.projects ` +
		`` +
		` WHERE projects.name = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name)
	p := Project{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, name).Scan(&p.ProjectID, &p.Name, &p.Description, &p.Metadata, &p.CreatedAt, &p.UpdatedAt); err != nil {
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
projects.metadata,
projects.created_at,
projects.updated_at ` +
		`FROM public.projects ` +
		`` +
		` WHERE projects.project_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, projectID)
	p := Project{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, projectID).Scan(&p.ProjectID, &p.Name, &p.Description, &p.Metadata, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
