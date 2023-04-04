package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// WorkItemType represents a row from 'public.work_item_types'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type WorkItemType struct {
	WorkItemTypeID int    `json:"workItemTypeID" db:"work_item_type_id"` // work_item_type_id
	ProjectID      int    `json:"projectID" db:"project_id"`             // project_id
	Name           string `json:"name" db:"name"`                        // name
	Description    string `json:"description" db:"description"`          // description
	Color          string `json:"color" db:"color"`                      // color

	WorkItem *WorkItem `json:"work_item" db:"work_item"` // O2O
	// xo fields
	_exists, _deleted bool
}

type WorkItemTypeSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTypeJoins
}
type WorkItemTypeSelectConfigOption func(*WorkItemTypeSelectConfig)

// WithWorkItemTypeLimit limits row selection.
func WithWorkItemTypeLimit(limit int) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemTypeOrderBy = string

type WorkItemTypeJoins struct {
	WorkItem bool
}

// WithWorkItemTypeJoin orders results by the given columns.
func WithWorkItemTypeJoin(joins WorkItemTypeJoins) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemType exists in the database.
func (wit *WorkItemType) Exists() bool {
	return wit._exists
}

// Deleted returns true when the WorkItemType has been marked for deletion from
// the database.
func (wit *WorkItemType) Deleted() bool {
	return wit._deleted
}

// Insert inserts the WorkItemType to the database.
func (wit *WorkItemType) Insert(ctx context.Context, db DB) error {
	switch {
	case wit._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wit._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_types (` +
		`project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING work_item_type_id `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err := db.QueryRow(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color).Scan(&wit.WorkItemTypeID); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Update updates a WorkItemType in the database.
func (wit *WorkItemType) Update(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wit._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_types SET ` +
		`project_id = $1, name = $2, description = $3, color = $4 ` +
		`WHERE work_item_type_id = $5 ` +
		`RETURNING work_item_type_id `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID)
	if err := db.QueryRow(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID).Scan(&wit.WorkItemTypeID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the WorkItemType to the database.
func (wit *WorkItemType) Save(ctx context.Context, db DB) error {
	if wit.Exists() {
		return wit.Update(ctx, db)
	}
	return wit.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemType.
func (wit *WorkItemType) Upsert(ctx context.Context, db DB) error {
	switch {
	case wit._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_types (` +
		`work_item_type_id, project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_type_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color  `
	// run
	logf(sqlstr, wit.WorkItemTypeID, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTypeID, wit.ProjectID, wit.Name, wit.Description, wit.Color); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Delete deletes the WorkItemType from the database.
func (wit *WorkItemType) Delete(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return nil
	case wit._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_types ` +
		`WHERE work_item_type_id = $1 `
	// run
	logf(sqlstr, wit.WorkItemTypeID)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTypeID); err != nil {
		return logerror(err)
	}
	// set deleted
	wit._deleted = true
	return nil
}

// WorkItemTypeByNameProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypeByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true then row(work_items.*) end)::jsonb as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.name = $2 AND work_item_types.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name, projectID)
	wit := WorkItemType{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.WorkItem, name, projectID).Scan(&wit.WorkItemTypeID, &wit.ProjectID, &wit.Name, &wit.Description, &wit.Color, &wit.WorkItem); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
}

// WorkItemTypeByWorkItemTypeID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_pkey'.
func WorkItemTypeByWorkItemTypeID(ctx context.Context, db DB, workItemTypeID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true then row(work_items.*) end)::jsonb as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.work_item_type_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemTypeID)
	wit := WorkItemType{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.WorkItem, workItemTypeID).Scan(&wit.WorkItemTypeID, &wit.ProjectID, &wit.Name, &wit.Description, &wit.Color, &wit.WorkItem); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemType's (ProjectID).
//
// Generated from foreign key 'work_item_types_project_id_fkey'.
func (wit *WorkItemType) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
