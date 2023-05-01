package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// WorkItemWorkItemTag represents a row from 'public.work_item_work_item_tag'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type WorkItemWorkItemTag struct {
	WorkItemTagID int   `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID" db:"work_item_id" required:"true"`        // work_item_id

	// xo fields
	_exists, _deleted bool
}

// WorkItemWorkItemTagCreateParams represents insert params for 'public.work_item_work_item_tag'
type WorkItemWorkItemTagCreateParams struct {
	WorkItemTagID int   `json:"workItemTagID"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID"`    // work_item_id
}

// WorkItemWorkItemTagUpdateParams represents update params for 'public.work_item_work_item_tag'
type WorkItemWorkItemTagUpdateParams struct {
	WorkItemTagID *int   `json:"workItemTagID"` // work_item_tag_id
	WorkItemID    *int64 `json:"workItemID"`    // work_item_id
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

const ()

type WorkItemWorkItemTagJoins struct {
}

// WithWorkItemWorkItemTagJoin joins with the given tables.
func WithWorkItemWorkItemTagJoin(joins WorkItemWorkItemTagJoins) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the WorkItemWorkItemTag to the database.
func (wiwit *WorkItemWorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemWorkItemTag, error) {
	switch {
	case wiwit._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wiwit._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_work_item_tag (` +
		`work_item_tag_id, work_item_id` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/db.Query: %w", err))
	}
	newwiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/pgx.CollectOneRow: %w", err))
	}
	newwiwit._exists = true
	*wiwit = newwiwit

	return wiwit, nil
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
	// logf(sqlstr, workItemID, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, workItemID, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/db.Query: %w", err))
	}
	wiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectOneRow: %w", err))
	}
	wiwit._exists = true

	return &wiwit, nil
}

// WorkItemWorkItemTagsByWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
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
		` WHERE work_item_work_item_tag.work_item_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, workItemID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
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
		` WHERE work_item_work_item_tag.work_item_tag_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, workItemTagID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagIDWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_work_item_tag_id_work_item_id_idx'.
func WorkItemWorkItemTagsByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
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
	// logf(sqlstr, workItemTagID, workItemID)
	rows, err := db.Query(ctx, sqlstr, workItemTagID, workItemID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemWorkItemTag's (WorkItemID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wiwit.WorkItemID)
}

// FKWorkItemTag_WorkItemTagID returns the WorkItemTag associated with the WorkItemWorkItemTag's (WorkItemTagID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_tag_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItemTag_WorkItemTagID(ctx context.Context, db DB) (*WorkItemTag, error) {
	return WorkItemTagByWorkItemTagID(ctx, db, wiwit.WorkItemTagID)
}
