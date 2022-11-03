package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// WorkItemField represents a row from 'public.work_item_fields'.
type WorkItemField struct {
	ProjectID int64  `json:"project_id"` // project_id
	Key       string `json:"key"`        // key
	// xo fields
	_exists, _deleted bool
}

type WorkItemFieldSelectConfig struct {
	limit    string
	orderBy  string
	joinWith []WorkItemFieldJoinBy
}

type WorkItemFieldSelectConfigOption func(*WorkItemFieldSelectConfig)

// WorkItemFieldWithLimit limits row selection.
func WorkItemFieldWithLimit(limit int) WorkItemFieldSelectConfigOption {
	return func(s *WorkItemFieldSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemFieldOrderBy = string

type WorkItemFieldJoinBy = string

// Exists returns true when the WorkItemField exists in the database.
func (wif *WorkItemField) Exists() bool {
	return wif._exists
}

// Deleted returns true when the WorkItemField has been marked for deletion from
// the database.
func (wif *WorkItemField) Deleted() bool {
	return wif._deleted
}

// Insert inserts the WorkItemField to the database.
func (wif *WorkItemField) Insert(ctx context.Context, db DB) error {
	switch {
	case wif._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wif._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_fields (` +
		`project_id, key` +
		`) VALUES (` +
		`$1, $2` +
		`) `
	// run
	logf(sqlstr, wif.ProjectID, wif.Key)
	if _, err := db.Exec(ctx, sqlstr, wif.ProjectID, wif.Key); err != nil {
		return logerror(err)
	}
	// set exists
	wif._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the WorkItemField from the database.
func (wif *WorkItemField) Delete(ctx context.Context, db DB) error {
	switch {
	case !wif._exists: // doesn't exist
		return nil
	case wif._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_fields ` +
		`WHERE project_id = $1 AND key = $2 `
	// run
	logf(sqlstr, wif.ProjectID, wif.Key)
	if _, err := db.Exec(ctx, sqlstr, wif.ProjectID, wif.Key); err != nil {
		return logerror(err)
	}
	// set deleted
	wif._deleted = true
	return nil
}

// WorkItemFieldByProjectIDKey retrieves a row from 'public.work_item_fields' as a WorkItemField.
//
// Generated from index 'work_item_fields_pkey'.
func WorkItemFieldByProjectIDKey(ctx context.Context, db DB, projectID int64, key string, opts ...WorkItemFieldSelectConfigOption) (*WorkItemField, error) {
	c := &WorkItemFieldSelectConfig{}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`project_id, key ` +
		`FROM public.work_item_fields ` +
		`WHERE project_id = $1 AND key = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, projectID, key)
	wif := WorkItemField{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, projectID, key).Scan(&wif.ProjectID, &wif.Key); err != nil {
		return nil, logerror(err)
	}
	return &wif, nil
}

// Project returns the Project associated with the WorkItemField's (ProjectID).
//
// Generated from foreign key 'work_item_fields_project_id_fkey'.
func (wif *WorkItemField) Project(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, int(wif.ProjectID))
}
