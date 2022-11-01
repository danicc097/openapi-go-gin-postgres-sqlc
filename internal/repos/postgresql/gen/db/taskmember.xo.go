package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/google/uuid"
)

// TaskMember represents a row from 'public.task_member'.
type TaskMember struct {
	TaskID int64     `json:"task_id"` // task_id
	Member uuid.UUID `json:"member"`  // member
	// xo fields
	_exists, _deleted bool
}

// TODO only create if exists
// GetMostRecentTaskMember returns n most recent rows from 'task_member',
// ordered by "created_at" in descending order.
func GetMostRecentTaskMember(ctx context.Context, db DB, n int) ([]*TaskMember, error) {
	// list
	const sqlstr = `SELECT ` +
		`task_id, member ` +
		`FROM public.task_member ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*TaskMember
	for rows.Next() {
		tm := TaskMember{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&tm.TaskID, &tm.Member); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &tm)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Exists returns true when the TaskMember exists in the database.
func (tm *TaskMember) Exists() bool {
	return tm._exists
}

// Deleted returns true when the TaskMember has been marked for deletion from
// the database.
func (tm *TaskMember) Deleted() bool {
	return tm._deleted
}

// Insert inserts the TaskMember to the database.
func (tm *TaskMember) Insert(ctx context.Context, db DB) error {
	switch {
	case tm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case tm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.task_member (` +
		`task_id, member` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, tm.TaskID, tm.Member)
	if _, err := db.Exec(ctx, sqlstr, tm.TaskID, tm.Member); err != nil {
		return logerror(err)
	}
	// set exists
	tm._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the TaskMember from the database.
func (tm *TaskMember) Delete(ctx context.Context, db DB) error {
	switch {
	case !tm._exists: // doesn't exist
		return nil
	case tm._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.task_member ` +
		`WHERE task_id = $1 AND member = $2`
	// run
	logf(sqlstr, tm.TaskID, tm.Member)
	if _, err := db.Exec(ctx, sqlstr, tm.TaskID, tm.Member); err != nil {
		return logerror(err)
	}
	// set deleted
	tm._deleted = true
	return nil
}

// TaskMemberByMemberTaskID retrieves a row from 'public.task_member' as a TaskMember.
//
// Generated from index 'task_member_member_task_id_idx'.
func TaskMemberByMemberTaskID(ctx context.Context, db DB, member uuid.UUID, taskID int64) ([]*TaskMember, error) {
	// query
	const sqlstr = `SELECT ` +
		`task_id, member ` +
		`FROM public.task_member ` +
		`WHERE member = $1 AND task_id = $2`
	// run
	logf(sqlstr, member, taskID)
	rows, err := db.Query(ctx, sqlstr, member, taskID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*TaskMember
	for rows.Next() {
		tm := TaskMember{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&tm.TaskID, &tm.Member); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &tm)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// TaskMemberByTaskIDMember retrieves a row from 'public.task_member' as a TaskMember.
//
// Generated from index 'task_member_pkey'.
func TaskMemberByTaskIDMember(ctx context.Context, db DB, taskID int64, member uuid.UUID) (*TaskMember, error) {
	// query
	const sqlstr = `SELECT ` +
		`task_id, member ` +
		`FROM public.task_member ` +
		`WHERE task_id = $1 AND member = $2`
	// run
	logf(sqlstr, taskID, member)
	tm := TaskMember{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, taskID, member).Scan(&tm.TaskID, &tm.Member); err != nil {
		return nil, logerror(err)
	}
	return &tm, nil
}

// User returns the User associated with the TaskMember's (Member).
//
// Generated from foreign key 'task_member_member_fkey'.
func (tm *TaskMember) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, tm.Member)
}

// Task returns the Task associated with the TaskMember's (TaskID).
//
// Generated from foreign key 'task_member_task_id_fkey'.
func (tm *TaskMember) Task(ctx context.Context, db DB) (*Task, error) {
	return TaskByTaskID(ctx, db, tm.TaskID)
}
