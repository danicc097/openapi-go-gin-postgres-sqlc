package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgtype"
)

// WorkItemPublic represents fields that may be exposed from 'public.work_items'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type WorkItemPublic struct {
	WorkItemID     int64        `json:"workItemID" required:"true"`     // work_item_id
	Title          string       `json:"title" required:"true"`          // title
	WorkItemTypeID int          `json:"workItemTypeID" required:"true"` // work_item_type_id
	Metadata       pgtype.JSONB `json:"metadata" required:"true"`       // metadata
	TeamID         int          `json:"teamID" required:"true"`         // team_id
	KanbanStepID   int          `json:"kanbanStepID" required:"true"`   // kanban_step_id
	Closed         bool         `json:"closed" required:"true"`         // closed
	CreatedAt      time.Time    `json:"createdAt" required:"true"`      // created_at
	UpdatedAt      time.Time    `json:"updatedAt" required:"true"`      // updated_at
	DeletedAt      *time.Time   `json:"deletedAt" required:"true"`      // deleted_at
}

// WorkItem represents a row from 'public.work_items'.
type WorkItem struct {
	WorkItemID     int64        `json:"work_item_id" db:"work_item_id"`           // work_item_id
	Title          string       `json:"title" db:"title"`                         // title
	WorkItemTypeID int          `json:"work_item_type_id" db:"work_item_type_id"` // work_item_type_id
	Metadata       pgtype.JSONB `json:"metadata" db:"metadata"`                   // metadata
	TeamID         int          `json:"team_id" db:"team_id"`                     // team_id
	KanbanStepID   int          `json:"kanban_step_id" db:"kanban_step_id"`       // kanban_step_id
	Closed         bool         `json:"closed" db:"closed"`                       // closed
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`               // created_at
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`               // updated_at
	DeletedAt      *time.Time   `json:"deleted_at" db:"deleted_at"`               // deleted_at

	TimeEntries      *[]TimeEntry       `json:"time_entries" db:"time_entries"`             // O2M
	WorkItemComments *[]WorkItemComment `json:"work_item_comments" db:"work_item_comments"` // O2M
	Users            *[]User            `json:"users" db:"users"`                           // M2M
	// xo fields
	_exists, _deleted bool
}

func (x *WorkItem) ToPublic() WorkItemPublic {
	return WorkItemPublic{
		WorkItemID: x.WorkItemID, Title: x.Title, WorkItemTypeID: x.WorkItemTypeID, Metadata: x.Metadata, TeamID: x.TeamID, KanbanStepID: x.KanbanStepID, Closed: x.Closed, CreatedAt: x.CreatedAt, UpdatedAt: x.UpdatedAt, DeletedAt: x.DeletedAt,
	}
}

type WorkItemSelectConfig struct {
	limit     string
	orderBy   string
	joins     WorkItemJoins
	deletedAt string
}
type WorkItemSelectConfigOption func(*WorkItemSelectConfig)

// WithWorkItemLimit limits row selection.
func WithWorkItemLimit(limit int) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

// WithDeletedWorkItemOnly limits result to records marked as deleted.
func WithDeletedWorkItemOnly() WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.deletedAt = " not null "
	}
}

type WorkItemOrderBy = string

const (
	WorkItemCreatedAtDescNullsFirst WorkItemOrderBy = " created_at DESC NULLS FIRST "
	WorkItemCreatedAtDescNullsLast  WorkItemOrderBy = " created_at DESC NULLS LAST "
	WorkItemCreatedAtAscNullsFirst  WorkItemOrderBy = " created_at ASC NULLS FIRST "
	WorkItemCreatedAtAscNullsLast   WorkItemOrderBy = " created_at ASC NULLS LAST "
	WorkItemUpdatedAtDescNullsFirst WorkItemOrderBy = " updated_at DESC NULLS FIRST "
	WorkItemUpdatedAtDescNullsLast  WorkItemOrderBy = " updated_at DESC NULLS LAST "
	WorkItemUpdatedAtAscNullsFirst  WorkItemOrderBy = " updated_at ASC NULLS FIRST "
	WorkItemUpdatedAtAscNullsLast   WorkItemOrderBy = " updated_at ASC NULLS LAST "
	WorkItemDeletedAtDescNullsFirst WorkItemOrderBy = " deleted_at DESC NULLS FIRST "
	WorkItemDeletedAtDescNullsLast  WorkItemOrderBy = " deleted_at DESC NULLS LAST "
	WorkItemDeletedAtAscNullsFirst  WorkItemOrderBy = " deleted_at ASC NULLS FIRST "
	WorkItemDeletedAtAscNullsLast   WorkItemOrderBy = " deleted_at ASC NULLS LAST "
)

// WithWorkItemOrderBy orders results by the given columns.
func WithWorkItemOrderBy(rows ...WorkItemOrderBy) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type WorkItemJoins struct {
	TimeEntries      bool
	WorkItemComments bool
	Users            bool
}

// WithWorkItemJoin orders results by the given columns.
func WithWorkItemJoin(joins WorkItemJoins) WorkItemSelectConfigOption {
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
		`title, work_item_type_id, metadata, team_id, kanban_step_id, closed, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING work_item_id, created_at, updated_at `
	// run
	logf(sqlstr, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt)
	if err := db.QueryRow(ctx, sqlstr, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt).Scan(&wi.WorkItemID, &wi.CreatedAt, &wi.UpdatedAt); err != nil {
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
		`title = $1, work_item_type_id = $2, metadata = $3, team_id = $4, kanban_step_id = $5, closed = $6, deleted_at = $7 ` +
		`WHERE work_item_id = $8 ` +
		`RETURNING work_item_id, created_at, updated_at `
	// run
	logf(sqlstr, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.CreatedAt, wi.UpdatedAt, wi.DeletedAt, wi.WorkItemID)
	if err := db.QueryRow(ctx, sqlstr, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt, wi.WorkItemID).Scan(&wi.WorkItemID, &wi.CreatedAt, &wi.UpdatedAt); err != nil {
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
		`work_item_id, title, work_item_type_id, metadata, team_id, kanban_step_id, closed, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, work_item_type_id = EXCLUDED.work_item_type_id, metadata = EXCLUDED.metadata, team_id = EXCLUDED.team_id, kanban_step_id = EXCLUDED.kanban_step_id, closed = EXCLUDED.closed, deleted_at = EXCLUDED.deleted_at  `
	// run
	logf(sqlstr, wi.WorkItemID, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID, wi.Title, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.DeletedAt); err != nil {
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
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`work_items.work_item_id,
work_items.title,
work_items.work_item_type_id,
work_items.metadata,
work_items.team_id,
work_items.kanban_step_id,
work_items.closed,
work_items.created_at,
work_items.updated_at,
work_items.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_work_item_comments.work_item_comments end)::jsonb as work_item_comments,
(case when $3::boolean = true then joined_users.users end)::jsonb as users `+
		`FROM public.work_items `+
		`-- O2M join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- O2M join generated from "work_item_comments_work_item_id_fkey"
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
				work_item_id) joined_users on joined_users.users_work_item_id = work_items.work_item_id`+
		` WHERE work_items.work_item_id = $4  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID)
	wi := WorkItem{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Users, workItemID).Scan(&wi.WorkItemID, &wi.Title, &wi.WorkItemTypeID, &wi.Metadata, &wi.TeamID, &wi.KanbanStepID, &wi.Closed, &wi.CreatedAt, &wi.UpdatedAt, &wi.DeletedAt, &wi.TimeEntries, &wi.WorkItemComments, &wi.Users); err != nil {
		return nil, logerror(err)
	}
	return &wi, nil
}

// FKKanbanStep returns the KanbanStep associated with the WorkItem's (KanbanStepID).
//
// Generated from foreign key 'work_items_kanban_step_id_fkey'.
func (wi *WorkItem) FKKanbanStep(ctx context.Context, db DB) (*KanbanStep, error) {
	return KanbanStepByKanbanStepID(ctx, db, wi.KanbanStepID)
}

// FKTeam returns the Team associated with the WorkItem's (TeamID).
//
// Generated from foreign key 'work_items_team_id_fkey'.
func (wi *WorkItem) FKTeam(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, wi.TeamID)
}

// FKWorkItemType returns the WorkItemType associated with the WorkItem's (WorkItemTypeID).
//
// Generated from foreign key 'work_items_work_item_type_id_fkey'.
func (wi *WorkItem) FKWorkItemType(ctx context.Context, db DB) (*WorkItemType, error) {
	return WorkItemTypeByWorkItemTypeID(ctx, db, wi.WorkItemTypeID)
}
