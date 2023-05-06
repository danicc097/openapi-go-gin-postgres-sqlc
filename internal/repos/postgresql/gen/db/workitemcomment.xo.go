package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// WorkItemComment represents a row from 'public.work_item_comments'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type WorkItemComment struct {
	WorkItemCommentID int64     `json:"workItemCommentID" db:"work_item_comment_id" required:"true"` // work_item_comment_id
	WorkItemID        int64     `json:"workItemID" db:"work_item_id" required:"true"`                // work_item_id
	UserID            uuid.UUID `json:"userID" db:"user_id" required:"true"`                         // user_id
	Message           string    `json:"message" db:"message" required:"true"`                        // message
	CreatedAt         time.Time `json:"createdAt" db:"created_at" required:"true"`                   // created_at
	UpdatedAt         time.Time `json:"updatedAt" db:"updated_at" required:"true"`                   // updated_at

	UserJoin     *User     `json:"-" db:"user" openapi-go:"ignore"`      // O2O (generated from M2O)
	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O (generated from M2O)

}

// WorkItemCommentCreateParams represents insert params for 'public.work_item_comments'.
type WorkItemCommentCreateParams struct {
	WorkItemID int64     `json:"workItemID" required:"true"` // work_item_id
	UserID     uuid.UUID `json:"userID" required:"true"`     // user_id
	Message    string    `json:"message" required:"true"`    // message
}

// CreateWorkItemComment creates a new WorkItemComment in the database with the given params.
func CreateWorkItemComment(ctx context.Context, db DB, params *WorkItemCommentCreateParams) (*WorkItemComment, error) {
	wic := &WorkItemComment{
		WorkItemID: params.WorkItemID,
		UserID:     params.UserID,
		Message:    params.Message,
	}

	return wic.Insert(ctx, db)
}

// WorkItemCommentUpdateParams represents update params for 'public.work_item_comments'
type WorkItemCommentUpdateParams struct {
	WorkItemID *int64     `json:"workItemID" required:"true"` // work_item_id
	UserID     *uuid.UUID `json:"userID" required:"true"`     // user_id
	Message    *string    `json:"message" required:"true"`    // message
}

// SetUpdateParams updates public.work_item_comments struct fields with the specified params.
func (wic *WorkItemComment) SetUpdateParams(params *WorkItemCommentUpdateParams) {
	if params.WorkItemID != nil {
		wic.WorkItemID = *params.WorkItemID
	}
	if params.UserID != nil {
		wic.UserID = *params.UserID
	}
	if params.Message != nil {
		wic.Message = *params.Message
	}
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
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
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type WorkItemCommentJoins struct {
	User     bool
	WorkItem bool
}

// WithWorkItemCommentJoin joins with the given tables.
func WithWorkItemCommentJoin(joins WorkItemCommentJoins) WorkItemCommentSelectConfigOption {
	return func(s *WorkItemCommentSelectConfig) {
		s.joins = WorkItemCommentJoins{
			User:     s.joins.User || joins.User,
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// Insert inserts the WorkItemComment to the database.
func (wic *WorkItemComment) Insert(ctx context.Context, db DB) (*WorkItemComment, error) {
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

	*wic = newwic

	return wic, nil
}

// Update updates a WorkItemComment in the database.
func (wic *WorkItemComment) Update(ctx context.Context, db DB) (*WorkItemComment, error) {
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
	*wic = newwic

	return wic, nil
}

// Upsert upserts a WorkItemComment in the database.
// Requires appropiate PK(s) to be set beforehand.
func (wic *WorkItemComment) Upsert(ctx context.Context, db DB, params *WorkItemCommentCreateParams) (*WorkItemComment, error) {
	var err error

	wic.WorkItemID = params.WorkItemID
	wic.UserID = params.UserID
	wic.Message = params.Message

	wic, err = wic.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			wic, err = wic.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return wic, err
}

// Delete deletes the WorkItemComment from the database.
func (wic *WorkItemComment) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_comments ` +
		`WHERE work_item_comment_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wic.WorkItemCommentID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemCommentPaginatedByWorkItemCommentID returns a cursor-paginated list of WorkItemComment.
func WorkItemCommentPaginatedByWorkItemCommentID(ctx context.Context, db DB, workItemCommentID int64, opts ...WorkItemCommentSelectConfigOption) ([]WorkItemComment, error) {
	c := &WorkItemCommentSelectConfig{joins: WorkItemCommentJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`work_item_comments.work_item_comment_id,
work_item_comments.work_item_id,
work_item_comments.user_id,
work_item_comments.message,
work_item_comments.created_at,
work_item_comments.updated_at,
(case when $1::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $2::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_comments ` +
		`-- O2O join generated from "work_item_comments_user_id_fkey (Generated from M2O)"
left join users on users.user_id = work_item_comments.user_id
-- O2O join generated from "work_item_comments_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = work_item_comments.work_item_id` +
		` WHERE work_item_comments.work_item_comment_id > $3 `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemCommentID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
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
work_item_comments.updated_at,
(case when $1::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $2::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_comments ` +
		`-- O2O join generated from "work_item_comments_user_id_fkey (Generated from M2O)"
left join users on users.user_id = work_item_comments.user_id
-- O2O join generated from "work_item_comments_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = work_item_comments.work_item_id` +
		` WHERE work_item_comments.work_item_comment_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemCommentID)
	rows, err := db.Query(ctx, sqlstr, c.joins.User, c.joins.WorkItem, workItemCommentID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/db.Query: %w", err))
	}
	wic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/pgx.CollectOneRow: %w", err))
	}

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
work_item_comments.updated_at,
(case when $1::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $2::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_comments ` +
		`-- O2O join generated from "work_item_comments_user_id_fkey (Generated from M2O)"
left join users on users.user_id = work_item_comments.user_id
-- O2O join generated from "work_item_comments_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = work_item_comments.work_item_id` +
		` WHERE work_item_comments.work_item_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.User, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/WorkItemCommentsByWorkItemID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemComment/WorkItemCommentsByWorkItemID/pgx.CollectRows: %w", err))
	}
	return res, nil
}
