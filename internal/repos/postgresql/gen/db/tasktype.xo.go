package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// TaskType represents a row from 'public.task_types'.
type TaskType struct {
	TaskTypeID  int    `json:"task_type_id" db:"task_type_id"` // task_type_id
	TeamID      int64  `json:"team_id" db:"team_id"`           // team_id
	Name        string `json:"name" db:"name"`                 // name
	Description string `json:"description" db:"description"`   // description
	Color       string `json:"color" db:"color"`               // color

	// xo fields
	_exists, _deleted bool
}

type TaskTypeSelectConfig struct {
	limit   string
	orderBy string
	joins   TaskTypeJoins
}

type TaskTypeSelectConfigOption func(*TaskTypeSelectConfig)

// TaskTypeWithLimit limits row selection.
func TaskTypeWithLimit(limit int) TaskTypeSelectConfigOption {
	return func(s *TaskTypeSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type TaskTypeOrderBy = string

type TaskTypeJoins struct{}

// TaskTypeWithJoin orders results by the given columns.
func TaskTypeWithJoin(joins TaskTypeJoins) TaskTypeSelectConfigOption {
	return func(s *TaskTypeSelectConfig) {
		s.joins = joins
	}
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
	sqlstr := `INSERT INTO public.task_types (` +
		`team_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING task_type_id `
	// run
	logf(sqlstr, tt.TeamID, tt.Name, tt.Description, tt.Color)
	if err := db.QueryRow(ctx, sqlstr, tt.TeamID, tt.Name, tt.Description, tt.Color).Scan(&tt.TaskTypeID); err != nil {
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
	sqlstr := `UPDATE public.task_types SET ` +
		`team_id = $1, name = $2, description = $3, color = $4 ` +
		`WHERE task_type_id = $5 `
	// run
	logf(sqlstr, tt.TeamID, tt.Name, tt.Description, tt.Color, tt.TaskTypeID)
	if _, err := db.Exec(ctx, sqlstr, tt.TeamID, tt.Name, tt.Description, tt.Color, tt.TaskTypeID); err != nil {
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
	sqlstr := `INSERT INTO public.task_types (` +
		`task_type_id, team_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (task_type_id) DO ` +
		`UPDATE SET ` +
		`team_id = EXCLUDED.team_id, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color  `
	// run
	logf(sqlstr, tt.TaskTypeID, tt.TeamID, tt.Name, tt.Description, tt.Color)
	if _, err := db.Exec(ctx, sqlstr, tt.TaskTypeID, tt.TeamID, tt.Name, tt.Description, tt.Color); err != nil {
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
	sqlstr := `DELETE FROM public.task_types ` +
		`WHERE task_type_id = $1 `
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
func TaskTypeByTaskTypeID(ctx context.Context, db DB, taskTypeID int, opts ...TaskTypeSelectConfigOption) (*TaskType, error) {
	c := &TaskTypeSelectConfig{
		joins: TaskTypeJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`task_types.task_type_id,
task_types.team_id,
task_types.name,
task_types.description,
task_types.color ` +
		`FROM public.task_types ` +
		`` +
		` WHERE task_type_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, taskTypeID)
	tt := TaskType{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, taskTypeID).Scan(&tt.TaskTypeID, &tt.TeamID, &tt.Name, &tt.Description, &tt.Color); err != nil {
		return nil, logerror(err)
	}
	return &tt, nil
}

// TaskTypeByTeamIDName retrieves a row from 'public.task_types' as a TaskType.
//
// Generated from index 'task_types_team_id_name_key'.
func TaskTypeByTeamIDName(ctx context.Context, db DB, teamID int64, name string, opts ...TaskTypeSelectConfigOption) (*TaskType, error) {
	c := &TaskTypeSelectConfig{
		joins: TaskTypeJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`task_types.task_type_id,
task_types.team_id,
task_types.name,
task_types.description,
task_types.color ` +
		`FROM public.task_types ` +
		`` +
		` WHERE team_id = $1 AND name = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, teamID, name)
	tt := TaskType{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, teamID, name).Scan(&tt.TaskTypeID, &tt.TeamID, &tt.Name, &tt.Description, &tt.Color); err != nil {
		return nil, logerror(err)
	}
	return &tt, nil
}

// Team returns the Team associated with the TaskType's (TeamID).
//
// Generated from foreign key 'task_types_team_id_fkey'.
func (tt *TaskType) Team(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, int(tt.TeamID))
}
