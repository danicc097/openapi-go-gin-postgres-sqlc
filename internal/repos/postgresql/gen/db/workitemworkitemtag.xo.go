package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// WorkItemWorkItemTagPublic represents fields that may be exposed from 'public.work_item_work_item_tag'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type WorkItemWorkItemTagPublic struct {
	WorkItemTagID int   `json:"workItemTagID"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID"`    // work_item_id
}

// WorkItemWorkItemTag represents a row from 'public.work_item_work_item_tag'.
type WorkItemWorkItemTag struct {
	WorkItemTagID int   `json:"work_item_tag_id" db:"work_item_tag_id"` // work_item_tag_id
	WorkItemID    int64 `json:"work_item_id" db:"work_item_id"`         // work_item_id

	// xo fields
	_exists, _deleted bool
}

func (x *WorkItemWorkItemTag) ToPublic() WorkItemWorkItemTagPublic {
	return WorkItemWorkItemTagPublic{
		WorkItemTagID: x.WorkItemTagID, WorkItemID: x.WorkItemID,
	}
}

type WorkItemWorkItemTagSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemWorkItemTagJoins
}
type WorkItemWorkItemTagSelectConfigOption func(*WorkItemWorkItemTagSelectConfig)

// WithWorkItemWorkItemTagLimit limits row selection.
func WithWorkItemWorkItemTagLimit(limit int) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemWorkItemTagOrderBy = string

type WorkItemWorkItemTagJoins struct{}

// WithWorkItemWorkItemTagJoin orders results by the given columns.
func WithWorkItemWorkItemTagJoin(joins WorkItemWorkItemTagJoins) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemWorkItemTag exists in the database.
func (wiwit *WorkItemWorkItemTag) Exists() bool {
	return wiwit._exists
}

// Deleted returns true when the WorkItemWorkItemTag has been marked for deletion from
// the database.
func (wiwit *WorkItemWorkItemTag) Deleted() bool {
	return wiwit._deleted
}

// Insert inserts the WorkItemWorkItemTag to the database.
func (wiwit *WorkItemWorkItemTag) Insert(ctx context.Context, db DB) error {
	switch {
	case wiwit._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wiwit._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_work_item_tag (` +
		`work_item_tag_id, work_item_id` +
		`) VALUES (` +
		`$1, $2` +
		`) `
	// run
	logf(sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	if _, err := db.Exec(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID); err != nil {
		return logerror(err)
	}
	// set exists
	wiwit._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the WorkItemWorkItemTag from the database.
func (wiwit *WorkItemWorkItemTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !wiwit._exists: // doesn't exist
		return nil
	case wiwit._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_work_item_tag ` +
		`WHERE work_item_tag_id = $1 AND work_item_id = $2 `
	// run
	logf(sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	if _, err := db.Exec(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	wiwit._deleted = true
	return nil
}

// WorkItemWorkItemTagByWorkItemIDWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagByWorkItemIDWorkItemTagID(ctx context.Context, db DB, workItemID int64, workItemTagID int, opts ...WorkItemWorkItemTagSelectConfigOption) (*WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id ` +
		`FROM public.work_item_work_item_tag ` +
		`` +
		` WHERE work_item_work_item_tag.work_item_id = $1 AND work_item_work_item_tag.work_item_tag_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID, workItemTagID)
	wiwit := WorkItemWorkItemTag{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, workItemID, workItemTagID).Scan(&wiwit.WorkItemTagID, &wiwit.WorkItemID); err != nil {
		return nil, logerror(err)
	}
	return &wiwit, nil
}

// WorkItemWorkItemTagByWorkItemTagIDWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_work_item_tag_id_work_item_id_idx'.
func WorkItemWorkItemTagByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]*WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id ` +
		`FROM public.work_item_work_item_tag ` +
		`` +
		` WHERE work_item_work_item_tag.work_item_tag_id = $1 AND work_item_work_item_tag.work_item_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemTagID, workItemID)
	rows, err := db.Query(ctx, sqlstr, workItemTagID, workItemID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*WorkItemWorkItemTag
	for rows.Next() {
		wiwit := WorkItemWorkItemTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&wiwit.WorkItemTagID, &wiwit.WorkItemID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &wiwit)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// FKWorkItem returns the WorkItem associated with the WorkItemWorkItemTag's (WorkItemID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItem(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wiwit.WorkItemID)
}

// FKWorkItemTag returns the WorkItemTag associated with the WorkItemWorkItemTag's (WorkItemTagID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_tag_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItemTag(ctx context.Context, db DB) (*WorkItemTag, error) {
	return WorkItemTagByWorkItemTagID(ctx, db, wiwit.WorkItemTagID)
}
