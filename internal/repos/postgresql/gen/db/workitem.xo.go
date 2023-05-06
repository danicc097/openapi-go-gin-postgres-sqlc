package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgx/v5"
)

// WorkItem represents a row from 'public.work_items'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type WorkItem struct {
	WorkItemID     int64      `json:"workItemID" db:"work_item_id" required:"true"`          // work_item_id
	Title          string     `json:"title" db:"title" required:"true"`                      // title
	Description    string     `json:"description" db:"description" required:"true"`          // description
	WorkItemTypeID int        `json:"workItemTypeID" db:"work_item_type_id" required:"true"` // work_item_type_id
	Metadata       []byte     `json:"metadata" db:"metadata" required:"true"`                // metadata
	TeamID         int        `json:"teamID" db:"team_id" required:"true"`                   // team_id
	KanbanStepID   int        `json:"kanbanStepID" db:"kanban_step_id" required:"true"`      // kanban_step_id
	Closed         *time.Time `json:"closed" db:"closed" required:"true"`                    // closed
	TargetDate     time.Time  `json:"targetDate" db:"target_date" required:"true"`           // target_date
	CreatedAt      time.Time  `json:"createdAt" db:"created_at" required:"true"`             // created_at
	UpdatedAt      time.Time  `json:"updatedAt" db:"updated_at" required:"true"`             // updated_at
	DeletedAt      *time.Time `json:"deletedAt" db:"deleted_at" required:"true"`             // deleted_at

	DemoTwoWorkItemJoin  *DemoTwoWorkItem   `json:"-" db:"demo_two_work_item" openapi-go:"ignore"` // O2O
	DemoWorkItemJoin     *DemoWorkItem      `json:"-" db:"demo_work_item" openapi-go:"ignore"`     // O2O
	TimeEntriesJoin      *[]TimeEntry       `json:"-" db:"time_entries" openapi-go:"ignore"`       // M2O
	WorkItemCommentsJoin *[]WorkItemComment `json:"-" db:"work_item_comments" openapi-go:"ignore"` // M2O
	MembersJoin          *[]WorkItem_Member `json:"-" db:"members" openapi-go:"ignore"`            // M2M
	WorkItemTagsJoin     *[]WorkItemTag     `json:"-" db:"work_item_tags" openapi-go:"ignore"`     // M2M
	WorkItemJoin         *WorkItem          `json:"-" db:"work_item" openapi-go:"ignore"`          // O2O
	WorkItemJoin         *WorkItem          `json:"-" db:"work_item" openapi-go:"ignore"`          // O2O

}

// WorkItemCreateParams represents insert params for 'public.work_items'
type WorkItemCreateParams struct {
	Title          string     `json:"title" required:"true"`          // title
	Description    string     `json:"description" required:"true"`    // description
	WorkItemTypeID int        `json:"workItemTypeID" required:"true"` // work_item_type_id
	Metadata       []byte     `json:"metadata" required:"true"`       // metadata
	TeamID         int        `json:"teamID" required:"true"`         // team_id
	KanbanStepID   int        `json:"kanbanStepID" required:"true"`   // kanban_step_id
	Closed         *time.Time `json:"closed" required:"true"`         // closed
	TargetDate     time.Time  `json:"targetDate" required:"true"`     // target_date
}

// CreateWorkItem creates a new WorkItem in the database with the given params.
func CreateWorkItem(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	wi := &WorkItem{
		Title:          params.Title,
		Description:    params.Description,
		WorkItemTypeID: params.WorkItemTypeID,
		Metadata:       params.Metadata,
		TeamID:         params.TeamID,
		KanbanStepID:   params.KanbanStepID,
		Closed:         params.Closed,
		TargetDate:     params.TargetDate,
	}

	return wi.Insert(ctx, db)
}

// WorkItemUpdateParams represents update params for 'public.work_items'
type WorkItemUpdateParams struct {
	Title          *string     `json:"title" required:"true"`          // title
	Description    *string     `json:"description" required:"true"`    // description
	WorkItemTypeID *int        `json:"workItemTypeID" required:"true"` // work_item_type_id
	Metadata       *[]byte     `json:"metadata" required:"true"`       // metadata
	TeamID         *int        `json:"teamID" required:"true"`         // team_id
	KanbanStepID   *int        `json:"kanbanStepID" required:"true"`   // kanban_step_id
	Closed         **time.Time `json:"closed" required:"true"`         // closed
	TargetDate     *time.Time  `json:"targetDate" required:"true"`     // target_date
}

// SetUpdateParams updates public.work_items struct fields with the specified params.
func (wi *WorkItem) SetUpdateParams(params *WorkItemUpdateParams) {
	if params.Title != nil {
		wi.Title = *params.Title
	}
	if params.Description != nil {
		wi.Description = *params.Description
	}
	if params.WorkItemTypeID != nil {
		wi.WorkItemTypeID = *params.WorkItemTypeID
	}
	if params.Metadata != nil {
		wi.Metadata = *params.Metadata
	}
	if params.TeamID != nil {
		wi.TeamID = *params.TeamID
	}
	if params.KanbanStepID != nil {
		wi.KanbanStepID = *params.KanbanStepID
	}
	if params.Closed != nil {
		wi.Closed = *params.Closed
	}
	if params.TargetDate != nil {
		wi.TargetDate = *params.TargetDate
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
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
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type WorkItemJoins struct {
	DemoTwoWorkItem  bool
	DemoWorkItem     bool
	TimeEntries      bool
	WorkItemComments bool
	Members          bool
	WorkItemTags     bool
	WorkItem         bool
	WorkItem         bool
}

// WithWorkItemJoin joins with the given tables.
func WithWorkItemJoin(joins WorkItemJoins) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.joins = WorkItemJoins{

			DemoTwoWorkItem:  s.joins.DemoTwoWorkItem || joins.DemoTwoWorkItem,
			DemoWorkItem:     s.joins.DemoWorkItem || joins.DemoWorkItem,
			TimeEntries:      s.joins.TimeEntries || joins.TimeEntries,
			WorkItemComments: s.joins.WorkItemComments || joins.WorkItemComments,
			Members:          s.joins.Members || joins.Members,
			WorkItemTags:     s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItem:         s.joins.WorkItem || joins.WorkItem,
			WorkItem:         s.joins.WorkItem || joins.WorkItem,
		}
	}
}

type WorkItem_Member struct {
	User User                `json:"user" db:"users"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"`
}

// Insert inserts the WorkItem to the database.
func (wi *WorkItem) Insert(ctx context.Context, db DB) (*WorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_items (` +
		`title, description, work_item_type_id, metadata, team_id, kanban_step_id, closed, target_date, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING * `
	// run
	logf(sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt)

	rows, err := db.Query(ctx, sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/db.Query: %w", err))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/pgx.CollectOneRow: %w", err))
	}

	*wi = newwi

	return wi, nil
}

// Update updates a WorkItem in the database.
func (wi *WorkItem) Update(ctx context.Context, db DB) (*WorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_items SET ` +
		`title = $1, description = $2, work_item_type_id = $3, metadata = $4, team_id = $5, kanban_step_id = $6, closed = $7, target_date = $8, deleted_at = $9 ` +
		`WHERE work_item_id = $10 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.CreatedAt, wi.UpdatedAt, wi.DeletedAt, wi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt, wi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/db.Query: %w", err))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/pgx.CollectOneRow: %w", err))
	}
	*wi = newwi

	return wi, nil
}

// Upsert performs an upsert for WorkItem.
func (wi *WorkItem) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.work_items (` +
		`work_item_id, title, description, work_item_type_id, metadata, team_id, kanban_step_id, closed, target_date, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, description = EXCLUDED.description, work_item_type_id = EXCLUDED.work_item_type_id, metadata = EXCLUDED.metadata, team_id = EXCLUDED.team_id, kanban_step_id = EXCLUDED.kanban_step_id, closed = EXCLUDED.closed, target_date = EXCLUDED.target_date, deleted_at = EXCLUDED.deleted_at ` +
		` RETURNING * `
	// run
	logf(sqlstr, wi.WorkItemID, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID, wi.Title, wi.Description, wi.WorkItemTypeID, wi.Metadata, wi.TeamID, wi.KanbanStepID, wi.Closed, wi.TargetDate, wi.DeletedAt); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the WorkItem from the database.
func (wi *WorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the WorkItem from the database via 'deleted_at'.
func (wi *WorkItem) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.work_items ` +
		`SET deleted_at = NOW() ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	wi.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted WorkItem from the database.
func (wi *WorkItem) Restore(ctx context.Context, db DB) (*WorkItem, error) {
	wi.DeletedAt = nil
	newwi, err := wi.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Restore/pgx.CollectRows: %w", err))
	}
	return newwi, nil
}

// WorkItemPaginatedByWorkItemID returns a cursor-paginated list of WorkItem.
func WorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

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
(case when $1::boolean = true and demo_two_work_items.work_item_id is not null then row(demo_two_work_items.*) end) as demo_two_work_item,
(case when $2::boolean = true and demo_work_items.work_item_id is not null then row(demo_work_items.*) end) as demo_work_item,
(case when $3::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $4::boolean = true then COALESCE(joined_work_item_comments.work_item_comments, '{}') end) as work_item_comments,
(case when $5::boolean = true then COALESCE(joined_members.__users, '{}') end) as members,
(case when $6::boolean = true then COALESCE(joined_work_item_tags.__work_item_tags, '{}') end) as work_item_tags,
(case when $7::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item,
(case when $8::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O reference)"
left join demo_two_work_items on demo_two_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O reference)"
left join demo_work_items on demo_work_items.work_item_id = work_items.work_item_id
-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- M2O join generated from "work_item_comments_work_item_id_fkey"
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
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_items.work_item_id

-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, array_agg(work_item_tags.*) filter (where work_item_tags.* is not null) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by work_item_work_item_tag_work_item_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id

-- O2O join generated from "work_items_pkey(O2O reference)"
left join work_items on work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_pkey"
left join work_items on work_items.work_item_id = work_items.work_item_id`+
		` WHERE work_items.work_item_id > $9  AND work_items.deleted_at is %s `, c.deletedAt)
	// TODO order by hardcoded default desc, if specific index  found generate reversed where ... < $i order by ... asc
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemsByDeletedAt_WhereDeletedAtIsNotNull retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_deleted_at_idx'.
func WorkItemsByDeletedAt_WhereDeletedAtIsNotNull(ctx context.Context, db DB, deletedAt *time.Time, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " not null ", joins: WorkItemJoins{}}

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
(case when $1::boolean = true and demo_two_work_items.work_item_id is not null then row(demo_two_work_items.*) end) as demo_two_work_item,
(case when $2::boolean = true and demo_work_items.work_item_id is not null then row(demo_work_items.*) end) as demo_work_item,
(case when $3::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $4::boolean = true then COALESCE(joined_work_item_comments.work_item_comments, '{}') end) as work_item_comments,
(case when $5::boolean = true then COALESCE(joined_members.__users, '{}') end) as members,
(case when $6::boolean = true then COALESCE(joined_work_item_tags.__work_item_tags, '{}') end) as work_item_tags,
(case when $7::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item,
(case when $8::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O reference)"
left join demo_two_work_items on demo_two_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O reference)"
left join demo_work_items on demo_work_items.work_item_id = work_items.work_item_id
-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- M2O join generated from "work_item_comments_work_item_id_fkey"
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
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_items.work_item_id

-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, array_agg(work_item_tags.*) filter (where work_item_tags.* is not null) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by work_item_work_item_tag_work_item_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id

-- O2O join generated from "work_items_pkey(O2O reference)"
left join work_items on work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_pkey"
left join work_items on work_items.work_item_id = work_items.work_item_id`+
		` WHERE work_items.deleted_at = $9 AND (deleted_at IS NOT NULL)  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.DemoTwoWorkItem, c.joins.DemoWorkItem, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Members, c.joins.WorkItemTags, c.joins.WorkItem, c.joins.WorkItem, deletedAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDeletedAt/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDeletedAt/pgx.CollectRows: %w", err))
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
(case when $1::boolean = true and demo_two_work_items.work_item_id is not null then row(demo_two_work_items.*) end) as demo_two_work_item,
(case when $2::boolean = true and demo_work_items.work_item_id is not null then row(demo_work_items.*) end) as demo_work_item,
(case when $3::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $4::boolean = true then COALESCE(joined_work_item_comments.work_item_comments, '{}') end) as work_item_comments,
(case when $5::boolean = true then COALESCE(joined_members.__users, '{}') end) as members,
(case when $6::boolean = true then COALESCE(joined_work_item_tags.__work_item_tags, '{}') end) as work_item_tags,
(case when $7::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item,
(case when $8::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O reference)"
left join demo_two_work_items on demo_two_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O reference)"
left join demo_work_items on demo_work_items.work_item_id = work_items.work_item_id
-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- M2O join generated from "work_item_comments_work_item_id_fkey"
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
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_items.work_item_id

-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, array_agg(work_item_tags.*) filter (where work_item_tags.* is not null) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by work_item_work_item_tag_work_item_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id

-- O2O join generated from "work_items_pkey(O2O reference)"
left join work_items on work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_pkey"
left join work_items on work_items.work_item_id = work_items.work_item_id`+
		` WHERE work_items.work_item_id = $9  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.DemoTwoWorkItem, c.joins.DemoWorkItem, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Members, c.joins.WorkItemTags, c.joins.WorkItem, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", err))
	}
	wi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", err))
	}

	return &wi, nil
}

// WorkItemsByTeamID retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_team_id_idx'.
func WorkItemsByTeamID(ctx context.Context, db DB, teamID int, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
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
(case when $1::boolean = true and demo_two_work_items.work_item_id is not null then row(demo_two_work_items.*) end) as demo_two_work_item,
(case when $2::boolean = true and demo_work_items.work_item_id is not null then row(demo_work_items.*) end) as demo_work_item,
(case when $3::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $4::boolean = true then COALESCE(joined_work_item_comments.work_item_comments, '{}') end) as work_item_comments,
(case when $5::boolean = true then COALESCE(joined_members.__users, '{}') end) as members,
(case when $6::boolean = true then COALESCE(joined_work_item_tags.__work_item_tags, '{}') end) as work_item_tags,
(case when $7::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item,
(case when $8::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item `+
		`FROM public.work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O reference)"
left join demo_two_work_items on demo_two_work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O reference)"
left join demo_work_items on demo_work_items.work_item_id = work_items.work_item_id
-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- M2O join generated from "work_item_comments_work_item_id_fkey"
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
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_items.work_item_id

-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, array_agg(work_item_tags.*) filter (where work_item_tags.* is not null) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by work_item_work_item_tag_work_item_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id

-- O2O join generated from "work_items_pkey(O2O reference)"
left join work_items on work_items.work_item_id = work_items.work_item_id
-- O2O join generated from "work_items_pkey"
left join work_items on work_items.work_item_id = work_items.work_item_id`+
		` WHERE work_items.team_id = $9  AND work_items.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.DemoTwoWorkItem, c.joins.DemoWorkItem, c.joins.TimeEntries, c.joins.WorkItemComments, c.joins.Members, c.joins.WorkItemTags, c.joins.WorkItem, c.joins.WorkItem, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTeamID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTeamID/pgx.CollectRows: %w", err))
	}
	return res, nil
}
