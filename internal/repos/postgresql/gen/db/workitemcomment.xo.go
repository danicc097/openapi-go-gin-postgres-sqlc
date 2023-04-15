package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// WorkItemComment represents a row from 'public.work_item_comments'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type WorkItemComment struct {
	WorkItemCommentID int64     `json:"workItemCommentID" db:"work_item_comment_id" required:"true"` // work_item_comment_id
	WorkItemID        int64     `json:"workItemID" db:"work_item_id" required:"true"`                // work_item_id
	UserID            uuid.UUID `json:"userID" db:"user_id" required:"true"`                         // user_id
	Message           string    `json:"message" db:"message" required:"true"`                        // message
	CreatedAt         time.Time `json:"createdAt" db:"created_at" required:"true"`                   // created_at
	UpdatedAt         time.Time `json:"updatedAt" db:"updated_at" required:"true"`                   // updated_at

	// xo fields
	_exists, _deleted bool
}

// WorkItemCommentCreateParams represents insert params for 'public.work_item_comments'
type WorkItemCommentCreateParams struct {
	WorkItemID int64     `json:"workItemID"` // work_item_id
	UserID     uuid.UUID `json:"userID"`     // user_id
	Message    string    `json:"message"`    // message
}

// WorkItemCommentUpdateParams represents update params for 'public.work_item_comments'
type WorkItemCommentUpdateParams struct {
	WorkItemID *int64     `json:"workItemID"` // work_item_id
	UserID     *uuid.UUID `json:"userID"`     // user_id
	Message    *string    `json:"message"`    // message
}

type WorkItemCommentSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemCommentJoins
}
type WorkItemCommentSelectConfigOption func(*WorkItemCommentSelectConfig)

// WithWorkItemCommentLimit limits row selection.
func WithWorkItemCommentLimit(limit int) WorkItemCommentSelectConfigOption {
	return func(s *WorkItemCommentSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemCommentOrderBy = string

const (
	WorkItemCommentCreatedAtDescNullsFirst WorkItemCommentOrderBy = " created_at DESC NULLS FIRST "
	WorkItemCommentCreatedAtDescNullsLast  WorkItemCommentOrderBy = " created_at DESC NULLS LAST "
	WorkItemCommentCreatedAtAscNullsFirst  WorkItemCommentOrderBy = " created_at ASC NULLS FIRST "
	WorkItemCommentCreatedAtAscNullsLast   WorkItemCommentOrderBy = " created_at ASC NULLS LAST "
	WorkItemCommentUpdatedAtDescNullsFirst WorkItemCommentOrderBy = " updated_at DESC NULLS FIRST "
	WorkItemCommentUpdatedAtDescNullsLast  WorkItemCommentOrderBy = " updated_at DESC NULLS LAST "
	WorkItemCommentUpdatedAtAscNullsFirst  WorkItemCommentOrderBy = " updated_at ASC NULLS FIRST "
	WorkItemCommentUpdatedAtAscNullsLast   WorkItemCommentOrderBy = " updated_at ASC NULLS LAST "
)

// WithWorkItemCommentOrderBy orders results by the given columns.
func WithWorkItemCommentOrderBy(rows ...WorkItemCommentOrderBy) WorkItemCommentSelectConfigOption {
	return func(s *WorkItemCommentSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type WorkItemCommentJoins struct {
}

// WithWorkItemCommentJoin joins with the given tables.
func WithWorkItemCommentJoin(joins WorkItemCommentJoins) WorkItemCommentSelectConfigOption {
	return func(s *WorkItemCommentSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemComment exists in the database.
func (wic *WorkItemComment) Exists() bool {
	return wic._exists
}

// Deleted returns true when the WorkItemComment has been marked for deletion from
// the database.
func (wic *WorkItemComment) Deleted() bool {
	return wic._deleted
}

// Insert inserts the WorkItemComment to the database.
func (wic *WorkItemComment) Insert(ctx context.Context, db DB) (*WorkItemComment, error) {
	switch {
	case wic._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wic._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_comments (` +
		`work_item_id, user_id, message` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, wic.WorkItemID, wic.UserID, wic.Message)

	rows, err := db.Query(ctx, sqlstr, wic.WorkItemID, wic.UserID, wic.Message)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Insert/db.Query: %w", err))
	}
	newwic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Insert/pgx.CollectOneRow: %w", err))
	}
	newwic._exists = true
	*wic = newwic

	return wic, nil
}

// Update updates a WorkItemComment in the database.
func (wic *WorkItemComment) Update(ctx context.Context, db DB) (*WorkItemComment, error) {
	switch {
	case !wic._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wic._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_comments SET ` +
		`work_item_id = $1, user_id = $2, message = $3 ` +
		`WHERE work_item_comment_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wic.WorkItemID, wic.UserID, wic.Message, wic.CreatedAt, wic.UpdatedAt, wic.WorkItemCommentID)

	rows, err := db.Query(ctx, sqlstr, wic.WorkItemID, wic.UserID, wic.Message, wic.WorkItemCommentID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Update/db.Query: %w", err))
	}
	newwic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Update/pgx.CollectOneRow: %w", err))
	}
	newwic._exists = true
	*wic = newwic

	return wic, nil
}

// Save saves the WorkItemComment to the database.
func (wic *WorkItemComment) Save(ctx context.Context, db DB) (*WorkItemComment, error) {
	if wic.Exists() {
		return wic.Update(ctx, db)
	}
	return wic.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemComment.
func (wic *WorkItemComment) Upsert(ctx context.Context, db DB) error {
	switch {
	case wic._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_comments (` +
		`work_item_comment_id, work_item_id, user_id, message` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (work_item_comment_id) DO ` +
		`UPDATE SET ` +
		`work_item_id = EXCLUDED.work_item_id, user_id = EXCLUDED.user_id, message = EXCLUDED.message ` +
		` RETURNING * `
	// run
	logf(sqlstr, wic.WorkItemCommentID, wic.WorkItemID, wic.UserID, wic.Message)
	if _, err := db.Exec(ctx, sqlstr, wic.WorkItemCommentID, wic.WorkItemID, wic.UserID, wic.Message); err != nil {
		return logerror(err)
	}
	// set exists
	wic._exists = true
	return nil
}

// Delete deletes the WorkItemComment from the database.
func (wic *WorkItemComment) Delete(ctx context.Context, db DB) error {
	switch {
	case !wic._exists: // doesn't exist
		return nil
	case wic._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_comments ` +
		`WHERE work_item_comment_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wic.WorkItemCommentID); err != nil {
		return logerror(err)
	}
	// set deleted
	wic._deleted = true
	return nil
}

// WorkItemCommentByWorkItemCommentID retrieves a row from 'public.work_item_comments' as a WorkItemComment.
//
// Generated from index 'work_item_comments_pkey'.
func WorkItemCommentByWorkItemCommentID(ctx context.Context, db DB, workItemCommentID int64, opts ...WorkItemCommentSelectConfigOption) (*WorkItemComment, error) {
	c := &WorkItemCommentSelectConfig{joins: WorkItemCommentJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_comments.work_item_comment_id,
work_item_comments.work_item_id,
work_item_comments.user_id,
work_item_comments.message,
work_item_comments.created_at,
work_item_comments.updated_at ` +
		`FROM public.work_item_comments ` +
		`` +
		` WHERE work_item_comments.work_item_comment_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemCommentID)
	rows, err := db.Query(ctx, sqlstr, workItemCommentID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/db.Query: %w", err))
	}
	wic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/pgx.CollectOneRow: %w", err))
	}
	wic._exists = true
	return &wic, nil
}

// WorkItemCommentsByWorkItemID retrieves a row from 'public.work_item_comments' as a WorkItemComment.
//
// Generated from index 'work_item_comments_work_item_id_idx'.
func WorkItemCommentsByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemCommentSelectConfigOption) ([]WorkItemComment, error) {
	c := &WorkItemCommentSelectConfig{joins: WorkItemCommentJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_comments.work_item_comment_id,
work_item_comments.work_item_id,
work_item_comments.user_id,
work_item_comments.message,
work_item_comments.created_at,
work_item_comments.updated_at ` +
		`FROM public.work_item_comments ` +
		`` +
		` WHERE work_item_comments.work_item_id = $1 `
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

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKUser_UserID returns the User associated with the WorkItemComment's (UserID).
//
// Generated from foreign key 'work_item_comments_user_id_fkey'.
func (wic *WorkItemComment) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wic.UserID)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemComment's (WorkItemID).
//
// Generated from foreign key 'work_item_comments_work_item_id_fkey'.
func (wic *WorkItemComment) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wic.WorkItemID)
}
