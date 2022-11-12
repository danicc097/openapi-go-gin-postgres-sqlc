package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgtype"
	"gopkg.in/guregu/null.v4"
)

// WorkItem represents a row from 'public.work_items'.
type WorkItem struct {
	WorkItemID   int64        `json:"work_item_id" db:"work_item_id"`     // work_item_id
	Title        string       `json:"title" db:"title"`                   // title
	Metadata     pgtype.JSONB `json:"metadata" db:"metadata"`             // metadata
	TeamID       int          `json:"team_id" db:"team_id"`               // team_id
	KanbanStepID int          `json:"kanban_step_id" db:"kanban_step_id"` // kanban_step_id
	Closed       bool         `json:"closed" db:"closed"`                 // closed
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`         // created_at
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`         // updated_at
	DeletedAt    null.Time    `json:"deleted_at" db:"deleted_at"`         // deleted_at

	WorkItemComments *[]WorkItemComment `json:"work_item_comments"` // O2M
	Users            *[]User            `json:"users"`              // M2M
	// xo fields
	_exists, _deleted bool
}

type WorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemJoins
}

type WorkItemSelectConfigOption func(*WorkItemSelectConfig)

// WorkItemWithLimit limits row selection.
func WorkItemWithLimit(limit int) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemOrderBy = string

const (
	WorkItemCreatedAtDescNullsFirst WorkItemOrderBy = "created_at DESC NULLS FIRST"
	WorkItemCreatedAtDescNullsLast  WorkItemOrderBy = "created_at DESC NULLS LAST"
	WorkItemCreatedAtAscNullsFirst  WorkItemOrderBy = "created_at ASC NULLS FIRST"
	WorkItemCreatedAtAscNullsLast   WorkItemOrderBy = "created_at ASC NULLS LAST"
	WorkItemUpdatedAtDescNullsFirst WorkItemOrderBy = "updated_at DESC NULLS FIRST"
	WorkItemUpdatedAtDescNullsLast  WorkItemOrderBy = "updated_at DESC NULLS LAST"
	WorkItemUpdatedAtAscNullsFirst  WorkItemOrderBy = "updated_at ASC NULLS FIRST"
	WorkItemUpdatedAtAscNullsLast   WorkItemOrderBy = "updated_at ASC NULLS LAST"
	WorkItemDeletedAtDescNullsFirst WorkItemOrderBy = "deleted_at DESC NULLS FIRST"
	WorkItemDeletedAtDescNullsLast  WorkItemOrderBy = "deleted_at DESC NULLS LAST"
	WorkItemDeletedAtAscNullsFirst  WorkItemOrderBy = "deleted_at ASC NULLS FIRST"
	WorkItemDeletedAtAscNullsLast   WorkItemOrderBy = "deleted_at ASC NULLS LAST"
)

// WorkItemWithOrderBy orders results by the given columns.
func WorkItemWithOrderBy(rows ...WorkItemOrderBy) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type WorkItemJoins struct {
	WorkItemComments bool
	Users            bool
}

// WorkItemWithJoin orders results by the given columns.
func WorkItemWithJoin(joins WorkItemJoins) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItem exists in the database.
func (wi *WorkItem) Exists() bool {
	return wi._exists
}

// Deleted returns true when the WorkItem has been marked for deletion from
// the database.
func (wi *WorkItem) Deleted() bool {
	return wi._deleted
}

// Insert inserts the WorkItem to the database.
func (wi *WorkItem) Insert(ctx context.Context, db DB) error {
	switch {
	case wi._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wi._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_items (` +
		`title, metadata, team_id, kanban_step_id, closed, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING work_item_id `
	// run
	logf(sqlstr, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt)
	if err := db.QueryRow(ctx, sqlstr, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt).Scan(&wi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set exists
	wi._exists = true
	return nil
}

// Update updates a WorkItem in the database.
func (wi *WorkItem) Update(ctx context.Context, db DB) error {
	switch {
	case !wi._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wi._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_items SET ` +
		`title = $1, metadata = $2, team_id = $3, kanban_step_id = $4, closed = $5, deleted_at = $6 ` +
		`WHERE work_item_id = $7 `
	// run
	logf(sqlstr, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.CreatedAt, wi.UpdatedAt, wi.DeletedAt, wi.WorkItemID)
	if _, err := db.Exec(ctx, sqlstr, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.CreatedAt, wi.UpdatedAt, wi.DeletedAt, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the WorkItem to the database.
func (wi *WorkItem) Save(ctx context.Context, db DB) error {
	if wi.Exists() {
		return wi.Update(ctx, db)
	}
	return wi.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItem.
func (wi *WorkItem) Upsert(ctx context.Context, db DB) error {
	switch {
	case wi._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_items (` +
		`work_item_id, title, metadata, team_id, kanban_step_id, closed, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, metadata = EXCLUDED.metadata, team_id = EXCLUDED.team_id, kanban_step_id = EXCLUDED.kanban_step_id, closed = EXCLUDED.closed, deleted_at = EXCLUDED.deleted_at  `
	// run
	logf(sqlstr, wi.WorkItemID, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID, wi.Title, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt); err != nil {
		return logerror(err)
	}
	// set exists
	wi._exists = true
	return nil
}

// Delete deletes the WorkItem from the database.
func (wi *WorkItem) Delete(ctx context.Context, db DB) error {
	switch {
	case !wi._exists: // doesn't exist
		return nil
	case wi._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_items ` +
		`WHERE work_item_id = $1 `
	// run
	logf(sqlstr, wi.WorkItemID)
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	wi._deleted = true
	return nil
}

// WorkItemByWorkItemID retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_pkey'.
func WorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemSelectConfigOption) (*WorkItem, error) {
	c := &WorkItemSelectConfig{
		joins: WorkItemJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_items.work_item_id,
work_items.title,
work_items.metadata,
work_items.team_id,
work_items.kanban_step_id,
work_items.closed,
work_items.created_at,
work_items.updated_at,
work_items.deleted_at,
(case when $1::boolean = true then joined_work_item_comments.work_item_comments end)::jsonb as work_item_comments,
(case when $2::boolean = true then joined_users.users end)::jsonb as users ` +
		`FROM public.work_items ` +
		`-- O2M join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , json_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
   group by
        work_item_id) joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
		work_item_id as users_work_item_id
		, json_agg(users.*) as users
	from
		work_item_member
		join users using (user_id)
	where
		work_item_id in (
			select
				work_item_id
			from
				work_item_member
			where
				user_id = any (
					select
						user_id
					from
						users))
			group by
				work_item_id) joined_users on joined_users.users_work_item_id = work_items.work_item_id` +
		` WHERE work_item_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID)
	wi := WorkItem{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, c.joins.WorkItemComments, c.joins.Users, workItemID).Scan(&wi.WorkItemID, &wi.Title, &wi.Metadata, &wi.TeamID, &wi.KanbanStepID, &wi.Closed, &wi.CreatedAt, &wi.UpdatedAt, &wi.DeletedAt); err != nil {
		return nil, logerror(err)
	}
	return &wi, nil
}

// KanbanStep returns the KanbanStep associated with the WorkItem's (KanbanStepID).
//
// Generated from foreign key 'work_items_kanban_step_id_fkey'.
func (wi *WorkItem) KanbanStep(ctx context.Context, db DB) (*KanbanStep, error) {
	return KanbanStepByKanbanStepID(ctx, db, wi.KanbanStepID)
}

// Team returns the Team associated with the WorkItem's (TeamID).
//
// Generated from foreign key 'work_items_team_id_fkey'.
func (wi *WorkItem) Team(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, wi.TeamID)
}
