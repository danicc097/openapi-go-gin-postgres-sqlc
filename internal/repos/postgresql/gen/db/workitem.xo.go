package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgtype"
)

// WorkItem represents a row from 'public.work_items'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type WorkItem struct {
	WorkItemID     int64        `json:"workItemID" db:"work_item_id"`          // work_item_id
	Title          string       `json:"title" db:"title"`                      // title
	Description    string       `json:"description" db:"description"`          // description
	WorkItemTypeID int          `json:"workItemTypeID" db:"work_item_type_id"` // work_item_type_id
	Metadata       pgtype.JSONB `json:"metadata" db:"metadata"`                // metadata
	TeamID         int          `json:"teamID" db:"team_id"`                   // team_id
	KanbanStepID   int          `json:"kanbanStepID" db:"kanban_step_id"`      // kanban_step_id
	Closed         *time.Time   `json:"closed" db:"closed"`                    // closed
	TargetDate     time.Time    `json:"targetDate" db:"target_date"`           // target_date
	CreatedAt      time.Time    `json:"createdAt" db:"created_at"`             // created_at
	UpdatedAt      time.Time    `json:"updatedAt" db:"updated_at"`             // updated_at
	DeletedAt      *time.Time   `json:"deletedAt" db:"deleted_at"`             // deleted_at

	DemoProjectWorkItem *DemoProjectWorkItem `json:"demoProjectWorkItem" db:"demo_project_work_item"` // O2O
	Project2WorkItem    *Project2WorkItem    `json:"project2workItem" db:"project_2_work_item"`       // O2O
	TimeEntries         *[]TimeEntry         `json:"timeEntries" db:"time_entries"`                   // O2M
	WorkItemComments    *[]WorkItemComment   `json:"workItemComments" db:"work_item_comments"`        // O2M
	Members             *[]User              `json:"members" db:"members"`                            // M2M
	WorkItemTags        *[]WorkItemTag       `json:"workItemTags" db:"work_item_tags"`                // M2M
	WorkItemType        *WorkItemType        `json:"workItemType" db:"work_item_type"`                // O2O
	// xo fields
	_exists, _deleted bool
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
	WorkItemClosedDescNullsFirst     WorkItemOrderBy = " closed DESC NULLS FIRST "
	WorkItemClosedDescNullsLast      WorkItemOrderBy = " closed DESC NULLS LAST "
	WorkItemClosedAscNullsFirst      WorkItemOrderBy = " closed ASC NULLS FIRST "
	WorkItemClosedAscNullsLast       WorkItemOrderBy = " closed ASC NULLS LAST "
	WorkItemTargetDateDescNullsFirst WorkItemOrderBy = " target_date DESC NULLS FIRST "
	WorkItemTargetDateDescNullsLast  WorkItemOrderBy = " target_date DESC NULLS LAST "
	WorkItemTargetDateAscNullsFirst  WorkItemOrderBy = " target_date ASC NULLS FIRST "
	WorkItemTargetDateAscNullsLast   WorkItemOrderBy = " target_date ASC NULLS LAST "
	WorkItemCreatedAtDescNullsFirst  WorkItemOrderBy = " created_at DESC NULLS FIRST "
	WorkItemCreatedAtDescNullsLast   WorkItemOrderBy = " created_at DESC NULLS LAST "
	WorkItemCreatedAtAscNullsFirst   WorkItemOrderBy = " created_at ASC NULLS FIRST "
	WorkItemCreatedAtAscNullsLast    WorkItemOrderBy = " created_at ASC NULLS LAST "
	WorkItemUpdatedAtDescNullsFirst  WorkItemOrderBy = " updated_at DESC NULLS FIRST "
	WorkItemUpdatedAtDescNullsLast   WorkItemOrderBy = " updated_at DESC NULLS LAST "
	WorkItemUpdatedAtAscNullsFirst   WorkItemOrderBy = " updated_at ASC NULLS FIRST "
	WorkItemUpdatedAtAscNullsLast    WorkItemOrderBy = " updated_at ASC NULLS LAST "
	WorkItemDeletedAtDescNullsFirst  WorkItemOrderBy = " deleted_at DESC NULLS FIRST "
	WorkItemDeletedAtDescNullsLast   WorkItemOrderBy = " deleted_at DESC NULLS LAST "
	WorkItemDeletedAtAscNullsFirst   WorkItemOrderBy = " deleted_at ASC NULLS FIRST "
	WorkItemDeletedAtAscNullsLast    WorkItemOrderBy = " deleted_at ASC NULLS LAST "
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
	DemoProjectWorkItem bool
	Project2WorkItem    bool
	TimeEntries         bool
	WorkItemComments    bool
	Members             bool
	WorkItemTags        bool
	WorkItemType        bool
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
		`title, description, work_item_type_id, metadata, team_id, kanban_step_id, closed, target_date, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING work_item_id, created_at, updated_at `
	// run
	logf(sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt)
	if err := db.QueryRow(ctx, sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt).Scan(&wi.WorkItemID, &wi.CreatedAt, &wi.UpdatedAt); err != nil {
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
		`title = $1, description = $2, work_item_type_id = $3, metadata = $4, team_id = $5, kanban_step_id = $6, closed = $7, target_date = $8, deleted_at = $9 ` +
		`WHERE work_item_id = $10 ` +
		`RETURNING work_item_id, created_at, updated_at `
	// run
	logf(sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.CreatedAt, wi.UpdatedAt, wi.DeletedAt, wi.WorkItemID)
	if err := db.QueryRow(ctx, sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt, wi.WorkItemID).Scan(&wi.WorkItemID, &wi.CreatedAt, &wi.UpdatedAt); err != nil {
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
		`work_item_id, title, description, work_item_type_id, metadata, team_id, kanban_step_id, closed, target_date, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, description = EXCLUDED.description, work_item_type_id = EXCLUDED.work_item_type_id, metadata = EXCLUDED.metadata, team_id = EXCLUDED.team_id, kanban_step_id = EXCLUDED.kanban_step_id, closed = EXCLUDED.closed, target_date = EXCLUDED.target_date, deleted_at = EXCLUDED.deleted_at  `
	// run
	logf(sqlstr, wi.WorkItemID, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt); err != nil {
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

// WorkItemsByDeletedAt_work_items_deleted_at_idx retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_deleted_at_idx'.
func WorkItemsByDeletedAt_work_items_deleted_at_idx(ctx context.Context, db DB, deletedAt *time.Time, opts ...WorkItemSelectConfigOption) ([]*WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`work_items.work_item_id,
work_items.title,
work_items.description,
work_items.work_item_type_id,
work_items.metadata,
work_items.team_id,
work_items.kanban_step_id,
work_items.closed,
work_items.target_date,
work_items.created_at,
work_items.updated_at,
work_items.deleted_at,
(case when $1::boolean = true then row(demo_project_work_items.*) end)::jsonb as demo_project_work_item,
(case when $2::boolean = true then row(project_2_work_items.*) end)::jsonb as project_2_work_item,
(case when $3::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $4::boolean = true then joined_work_item_comments.work_item_comments end)::jsonb as work_item_comments,
(case when $5::boolean = true then joined_users.users end)::jsonb as users,
(case when $6::boolean = true then joined_work_item_tags.work_item_tags end)::jsonb as work_item_tags,
(case when $7::boolean = true then row(work_item_types.*) end)::jsonb as work_item_type `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_project_work_items_work_item_id_fkey"
left join demo_project_work_items on demo_project_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "project_2_work_items_work_item_id_fkey"
left join project_2_work_items on project_2_work_items.work_item_id = work_items.work_item_id
-- O2M join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- O2M join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
   group by
        work_item_id) joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
		work_item_id as users_work_item_id
		, array_agg(users.*) as users
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
				work_item_id) joined_users on joined_users.users_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
		work_item_id as work_item_tags_work_item_id
		, array_agg(work_item_tags.*) as work_item_tags
	from
		work_item_work_item_tag
		join work_item_tags using (work_item_tag_id)
	where
		work_item_id in (
			select
				work_item_id
			from
				work_item_work_item_tag
			where
				work_item_tag_id = any (
					select
						work_item_tag_id
					from
						work_item_tags))
			group by
				work_item_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_item_types on work_item_types.work_item_type_id = work_items.work_item_type_id`+
		` WHERE work_items.deleted_at = $8 AND (deleted_at IS NOT NULL)  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.DemoProjectWorkItem, c.joins.Project2WorkItem, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Members, c.joins.WorkItemTags, c.joins.WorkItemType, deletedAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*WorkItem
	for rows.Next() {
		wi := WorkItem{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&wi.WorkItemID, &wi.Title, &wi.Description, &wi.WorkItemTypeID, &wi.Metadata, &wi.TeamID, &wi.KanbanStepID, &wi.Closed, &wi.TargetDate, &wi.CreatedAt, &wi.UpdatedAt, &wi.DeletedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &wi)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
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
work_items.description,
work_items.work_item_type_id,
work_items.metadata,
work_items.team_id,
work_items.kanban_step_id,
work_items.closed,
work_items.target_date,
work_items.created_at,
work_items.updated_at,
work_items.deleted_at,
(case when $1::boolean = true then row(demo_project_work_items.*) end)::jsonb as demo_project_work_item,
(case when $2::boolean = true then row(project_2_work_items.*) end)::jsonb as project_2_work_item,
(case when $3::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $4::boolean = true then joined_work_item_comments.work_item_comments end)::jsonb as work_item_comments,
(case when $5::boolean = true then joined_users.users end)::jsonb as users,
(case when $6::boolean = true then joined_work_item_tags.work_item_tags end)::jsonb as work_item_tags,
(case when $7::boolean = true then row(work_item_types.*) end)::jsonb as work_item_type `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_project_work_items_work_item_id_fkey"
left join demo_project_work_items on demo_project_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "project_2_work_items_work_item_id_fkey"
left join project_2_work_items on project_2_work_items.work_item_id = work_items.work_item_id
-- O2M join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- O2M join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
   group by
        work_item_id) joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
		work_item_id as users_work_item_id
		, array_agg(users.*) as users
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
				work_item_id) joined_users on joined_users.users_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
		work_item_id as work_item_tags_work_item_id
		, array_agg(work_item_tags.*) as work_item_tags
	from
		work_item_work_item_tag
		join work_item_tags using (work_item_tag_id)
	where
		work_item_id in (
			select
				work_item_id
			from
				work_item_work_item_tag
			where
				work_item_tag_id = any (
					select
						work_item_tag_id
					from
						work_item_tags))
			group by
				work_item_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_item_types on work_item_types.work_item_type_id = work_items.work_item_type_id`+
		` WHERE work_items.work_item_id = $8  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID)
	wi := WorkItem{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.DemoProjectWorkItem, c.joins.Project2WorkItem, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Members, c.joins.WorkItemTags, c.joins.WorkItemType, workItemID).Scan(&wi.WorkItemID, &wi.Title, &wi.Description, &wi.WorkItemTypeID, &wi.Metadata, &wi.TeamID, &wi.KanbanStepID, &wi.Closed, &wi.TargetDate, &wi.CreatedAt, &wi.UpdatedAt, &wi.DeletedAt, &wi.DemoProjectWorkItem, &wi.Project2WorkItem, &wi.TimeEntries, &wi.WorkItemComments, &wi.Members, &wi.WorkItemTags, &wi.WorkItemType); err != nil {
		return nil, logerror(err)
	}
	return &wi, nil
}

// FKKanbanStep_KanbanStepID returns the KanbanStep associated with the WorkItem's (KanbanStepID).
//
// Generated from foreign key 'work_items_kanban_step_id_fkey'.
func (wi *WorkItem) FKKanbanStep_KanbanStepID(ctx context.Context, db DB) (*KanbanStep, error) {
	return KanbanStepByKanbanStepID(ctx, db, wi.KanbanStepID)
}

// FKTeam_TeamID returns the Team associated with the WorkItem's (TeamID).
//
// Generated from foreign key 'work_items_team_id_fkey'.
func (wi *WorkItem) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, wi.TeamID)
}

// FKWorkItemType_WorkItemTypeID returns the WorkItemType associated with the WorkItem's (WorkItemTypeID).
//
// Generated from foreign key 'work_items_work_item_type_id_fkey'.
func (wi *WorkItem) FKWorkItemType_WorkItemTypeID(ctx context.Context, db DB) (*WorkItemType, error) {
	return WorkItemTypeByWorkItemTypeID(ctx, db, wi.WorkItemTypeID)
}
