package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// TaskType represents a row from 'public.task_types'.
type TaskType struct {
	TaskTypeID int    `json:"task_type_id"` // task_type_id
	ProjectID  int64  `json:"project_id"`   // project_id
	Name       string `json:"name"`         // name
	// xo fields
	_exists, _deleted bool
}

// TODO only create if exists
// GetMostRecentTaskType returns n most recent rows from 'task_types',
// ordered by "created_at" in descending order.
func GetMostRecentTaskType(ctx context.Context, db DB, n int) ([]*TaskType, error) {
	// list
	const sqlstr = `SELECT ` +
		`task_type_id, project_id, name ` +
		`FROM public.task_types ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*TaskType
	for rows.Next() {
		tt := TaskType{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&tt.TaskTypeID, &tt.ProjectID, &tt.Name); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &tt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Exists returns true when the TaskType exists in the database.
func (tt *TaskType) Exists() bool {
	return tt._exists
}

// Deleted returns true when the TaskType has been marked for deletion from
// the database.
func (tt *TaskType) Deleted() bool {
	return tt._deleted
}

// Insert inserts the TaskType to the database.
func (tt *TaskType) Insert(ctx context.Context, db DB) error {
	switch {
	case tt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case tt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.task_types (` +
		`project_id, name` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING task_type_id`
	// run
	logf(sqlstr, tt.ProjectID, tt.Name)
	if err := db.QueryRow(ctx, sqlstr, tt.ProjectID, tt.Name).Scan(&tt.TaskTypeID); err != nil {
		return logerror(err)
	}
	// set exists
	tt._exists = true
	return nil
}

// Update updates a TaskType in the database.
func (tt *TaskType) Update(ctx context.Context, db DB) error {
	switch {
	case !tt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case tt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.task_types SET ` +
		`project_id = $1, name = $2 ` +
		`WHERE task_type_id = $3`
	// run
	logf(sqlstr, tt.ProjectID, tt.Name, tt.TaskTypeID)
	if _, err := db.Exec(ctx, sqlstr, tt.ProjectID, tt.Name, tt.TaskTypeID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the TaskType to the database.
func (tt *TaskType) Save(ctx context.Context, db DB) error {
	if tt.Exists() {
		return tt.Update(ctx, db)
	}
	return tt.Insert(ctx, db)
}

// Upsert performs an upsert for TaskType.
func (tt *TaskType) Upsert(ctx context.Context, db DB) error {
	switch {
	case tt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.task_types (` +
		`task_type_id, project_id, name` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (task_type_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name `
	// run
	logf(sqlstr, tt.TaskTypeID, tt.ProjectID, tt.Name)
	if _, err := db.Exec(ctx, sqlstr, tt.TaskTypeID, tt.ProjectID, tt.Name); err != nil {
		return logerror(err)
	}
	// set exists
	tt._exists = true
	return nil
}

// Delete deletes the TaskType from the database.
func (tt *TaskType) Delete(ctx context.Context, db DB) error {
	switch {
	case !tt._exists: // doesn't exist
		return nil
	case tt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.task_types ` +
		`WHERE task_type_id = $1`
	// run
	logf(sqlstr, tt.TaskTypeID)
	if _, err := db.Exec(ctx, sqlstr, tt.TaskTypeID); err != nil {
		return logerror(err)
	}
	// set deleted
	tt._deleted = true
	return nil
}

// TaskTypeByTaskTypeID retrieves a row from 'public.task_types' as a TaskType.
//
// Generated from index 'task_types_pkey'.
func TaskTypeByTaskTypeID(ctx context.Context, db DB, taskTypeID int) (*TaskType, error) {
	// query
	const sqlstr = `SELECT ` +
		`task_type_id, project_id, name ` +
		`FROM public.task_types ` +
		`WHERE task_type_id = $1`
	// run
	logf(sqlstr, taskTypeID)
	tt := TaskType{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, taskTypeID).Scan(&tt.TaskTypeID, &tt.ProjectID, &tt.Name); err != nil {
		return nil, logerror(err)
	}
	return &tt, nil
}

// TaskTypeByProjectIDName retrieves a row from 'public.task_types' as a TaskType.
//
// Generated from index 'task_types_project_id_name_key'.
func TaskTypeByProjectIDName(ctx context.Context, db DB, projectID int64, name string) (*TaskType, error) {
	// query
	const sqlstr = `SELECT ` +
		`task_type_id, project_id, name ` +
		`FROM public.task_types ` +
		`WHERE project_id = $1 AND name = $2`
	// run
	logf(sqlstr, projectID, name)
	tt := TaskType{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, projectID, name).Scan(&tt.TaskTypeID, &tt.ProjectID, &tt.Name); err != nil {
		return nil, logerror(err)
	}
	return &tt, nil
}

// Project returns the Project associated with the TaskType's (ProjectID).
//
// Generated from foreign key 'task_types_project_id_fkey'.
func (tt *TaskType) Project(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, int(tt.ProjectID))
}
