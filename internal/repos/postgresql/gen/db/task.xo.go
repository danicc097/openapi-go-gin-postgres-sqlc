package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgtype"
)

// Task represents a row from 'public.tasks'.
type Task struct {
	TaskID     int64        `json:"task_id" db:"task_id"`           // task_id
	TaskTypeID int          `json:"task_type_id" db:"task_type_id"` // task_type_id
	WorkItemID int64        `json:"work_item_id" db:"work_item_id"` // work_item_id
	Title      string       `json:"title" db:"title"`               // title
	Metadata   pgtype.JSONB `json:"metadata" db:"metadata"`         // metadata
	Finished   *bool        `json:"finished" db:"finished"`         // finished
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`     // created_at
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`     // updated_at
	DeletedAt  *time.Time   `json:"deleted_at" db:"deleted_at"`     // deleted_at

	TaskType    *TaskType    `json:"task_type"`    // O2O
	TimeEntries *[]TimeEntry `json:"time_entries"` // O2M
	// xo fields
	_exists, _deleted bool
}

type TaskSelectConfig struct {
	limit   string
	orderBy string
	joins   TaskJoins
}

type TaskSelectConfigOption func(*TaskSelectConfig)

// TaskWithLimit limits row selection.
func TaskWithLimit(limit int) TaskSelectConfigOption {
	return func(s *TaskSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type TaskOrderBy = string

const (
	TaskCreatedAtDescNullsFirst TaskOrderBy = " created_at DESC NULLS FIRST "
	TaskCreatedAtDescNullsLast  TaskOrderBy = " created_at DESC NULLS LAST "
	TaskCreatedAtAscNullsFirst  TaskOrderBy = " created_at ASC NULLS FIRST "
	TaskCreatedAtAscNullsLast   TaskOrderBy = " created_at ASC NULLS LAST "
	TaskUpdatedAtDescNullsFirst TaskOrderBy = " updated_at DESC NULLS FIRST "
	TaskUpdatedAtDescNullsLast  TaskOrderBy = " updated_at DESC NULLS LAST "
	TaskUpdatedAtAscNullsFirst  TaskOrderBy = " updated_at ASC NULLS FIRST "
	TaskUpdatedAtAscNullsLast   TaskOrderBy = " updated_at ASC NULLS LAST "
	TaskDeletedAtDescNullsFirst TaskOrderBy = " deleted_at DESC NULLS FIRST "
	TaskDeletedAtDescNullsLast  TaskOrderBy = " deleted_at DESC NULLS LAST "
	TaskDeletedAtAscNullsFirst  TaskOrderBy = " deleted_at ASC NULLS FIRST "
	TaskDeletedAtAscNullsLast   TaskOrderBy = " deleted_at ASC NULLS LAST "
)

// TaskWithOrderBy orders results by the given columns.
func TaskWithOrderBy(rows ...TaskOrderBy) TaskSelectConfigOption {
	return func(s *TaskSelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type TaskJoins struct {
	TaskType    bool
	TimeEntries bool
}

// TaskWithJoin orders results by the given columns.
func TaskWithJoin(joins TaskJoins) TaskSelectConfigOption {
	return func(s *TaskSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Task exists in the database.
func (t *Task) Exists() bool {
	return t._exists
}

// Deleted returns true when the Task has been marked for deletion from
// the database.
func (t *Task) Deleted() bool {
	return t._deleted
}

// Insert inserts the Task to the database.
func (t *Task) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.tasks (` +
		`task_type_id, work_item_id, title, metadata, finished, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING task_id `
	// run
	logf(sqlstr, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.DeletedAt)
	if err := db.QueryRow(ctx, sqlstr, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.DeletedAt).Scan(&t.TaskID); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Update updates a Task in the database.
func (t *Task) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.tasks SET ` +
		`task_type_id = $1, work_item_id = $2, title = $3, metadata = $4, finished = $5, deleted_at = $6 ` +
		`WHERE task_id = $7 `
	// run
	logf(sqlstr, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.CreatedAt, t.UpdatedAt, t.DeletedAt, t.TaskID)
	if _, err := db.Exec(ctx, sqlstr, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.CreatedAt, t.UpdatedAt, t.DeletedAt, t.TaskID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Task to the database.
func (t *Task) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Task.
func (t *Task) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.tasks (` +
		`task_id, task_type_id, work_item_id, title, metadata, finished, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (task_id) DO ` +
		`UPDATE SET ` +
		`task_type_id = EXCLUDED.task_type_id, work_item_id = EXCLUDED.work_item_id, title = EXCLUDED.title, metadata = EXCLUDED.metadata, finished = EXCLUDED.finished, deleted_at = EXCLUDED.deleted_at  `
	// run
	logf(sqlstr, t.TaskID, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, t.TaskID, t.TaskTypeID, t.WorkItemID, t.Title, t.Metadata, t.Finished, t.DeletedAt); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Task from the database.
func (t *Task) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.tasks ` +
		`WHERE task_id = $1 `
	// run
	logf(sqlstr, t.TaskID)
	if _, err := db.Exec(ctx, sqlstr, t.TaskID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TaskByTaskID retrieves a row from 'public.tasks' as a Task.
//
// Generated from index 'tasks_pkey'.
func TaskByTaskID(ctx context.Context, db DB, taskID int64, opts ...TaskSelectConfigOption) (*Task, error) {
	c := &TaskSelectConfig{
		joins: TaskJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`tasks.task_id,
tasks.task_type_id,
tasks.work_item_id,
tasks.title,
tasks.metadata,
tasks.finished,
tasks.created_at,
tasks.updated_at,
tasks.deleted_at,
(case when $1::boolean = true then row_to_json(task_types.*) end)::jsonb as task_type,
(case when $2::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries ` +
		`FROM public.tasks ` +
		`-- O2O join generated from "tasks_task_type_id_fkey"
left join task_types on task_types.task_type_id = tasks.task_type_id
-- O2M join generated from "time_entries_task_id_fkey"
left join (
  select
  task_id as time_entries_task_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        task_id) joined_time_entries on joined_time_entries.time_entries_task_id = tasks.task_id` +
		` WHERE tasks.task_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, taskID)
	t := Task{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TaskType, c.joins.TimeEntries, taskID).Scan(&t.TaskID, &t.TaskTypeID, &t.WorkItemID, &t.Title, &t.Metadata, &t.Finished, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt, &t.TaskType, &t.TimeEntries); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// FKTaskType returns the TaskType associated with the Task's (TaskTypeID).
//
// Generated from foreign key 'tasks_task_type_id_fkey'.
func (t *Task) FKTaskType(ctx context.Context, db DB) (*TaskType, error) {
	return TaskTypeByTaskTypeID(ctx, db, t.TaskTypeID)
}

// FKWorkItem returns the WorkItem associated with the Task's (WorkItemID).
//
// Generated from foreign key 'tasks_work_item_id_fkey'.
func (t *Task) FKWorkItem(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, t.WorkItemID)
}
