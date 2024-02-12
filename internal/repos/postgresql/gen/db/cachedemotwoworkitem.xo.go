package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// CacheDemoTwoWorkItem represents a row from 'public.cache__demo_two_work_items'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type CacheDemoTwoWorkItem struct {
	CustomDateForProject2 *time.Time     `json:"customDateForProject2" db:"custom_date_for_project_2"`                   // custom_date_for_project_2
	WorkItemID            WorkItemID     `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`          // work_item_id
	Title                 string         `json:"title" db:"title" required:"true" nullable:"false"`                      // title
	Description           string         `json:"description" db:"description" required:"true" nullable:"false"`          // description
	WorkItemTypeID        WorkItemTypeID `json:"workItemTypeID" db:"work_item_type_id" required:"true" nullable:"false"` // work_item_type_id
	Metadata              map[string]any `json:"metadata" db:"metadata" required:"true" nullable:"false"`                // metadata
	TeamID                TeamID         `json:"teamID" db:"team_id" required:"true" nullable:"false"`                   // team_id
	KanbanStepID          KanbanStepID   `json:"kanbanStepID" db:"kanban_step_id" required:"true" nullable:"false"`      // kanban_step_id
	ClosedAt              *time.Time     `json:"closedAt" db:"closed_at"`                                                // closed_at
	TargetDate            time.Time      `json:"targetDate" db:"target_date" required:"true" nullable:"false"`           // target_date
	CreatedAt             time.Time      `json:"createdAt" db:"created_at" required:"true" nullable:"false"`             // created_at
	UpdatedAt             time.Time      `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`             // updated_at
	DeletedAt             *time.Time     `json:"deletedAt" db:"deleted_at"`                                              // deleted_at

	KanbanStepJoin               *KanbanStep                        `json:"-" db:"kanban_step_kanban_step_id" openapi-go:"ignore"`             // O2O kanban_steps (inferred)
	TeamJoin                     *Team                              `json:"-" db:"team_team_id" openapi-go:"ignore"`                           // O2O teams (inferred)
	WorkItemTypeJoin             *WorkItemType                      `json:"-" db:"work_item_type_work_item_type_id" openapi-go:"ignore"`       // O2O work_item_types (inferred)
	WorkItemTimeEntriesJoin      *[]TimeEntry                       `json:"-" db:"time_entries" openapi-go:"ignore"`                           // M2O cache__demo_two_work_items
	WorkItemAssignedUsersJoin    *[]User__WIAU_CacheDemoTwoWorkItem `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user
	WorkItemWorkItemCommentsJoin *[]WorkItemComment                 `json:"-" db:"work_item_comments" openapi-go:"ignore"`                     // M2O cache__demo_two_work_items
	WorkItemWorkItemTagsJoin     *[]WorkItemTag                     `json:"-" db:"work_item_work_item_tag_work_item_tags" openapi-go:"ignore"` // M2M work_item_work_item_tag

}

// CacheDemoTwoWorkItemCreateParams represents insert params for 'public.cache__demo_two_work_items'.
type CacheDemoTwoWorkItemCreateParams struct {
	ClosedAt              *time.Time     `json:"closedAt"`                                        // closed_at
	CustomDateForProject2 *time.Time     `json:"customDateForProject2"`                           // custom_date_for_project_2
	Description           string         `json:"description" required:"true" nullable:"false"`    // description
	KanbanStepID          KanbanStepID   `json:"kanbanStepID" required:"true" nullable:"false"`   // kanban_step_id
	Metadata              map[string]any `json:"metadata" required:"true" nullable:"false"`       // metadata
	TargetDate            time.Time      `json:"targetDate" required:"true" nullable:"false"`     // target_date
	TeamID                TeamID         `json:"teamID" required:"true" nullable:"false"`         // team_id
	Title                 string         `json:"title" required:"true" nullable:"false"`          // title
	WorkItemID            WorkItemID     `json:"-" required:"true" nullable:"false"`              // work_item_id
	WorkItemTypeID        WorkItemTypeID `json:"workItemTypeID" required:"true" nullable:"false"` // work_item_type_id
}

// CreateCacheDemoTwoWorkItem creates a new CacheDemoTwoWorkItem in the database with the given params.
func CreateCacheDemoTwoWorkItem(ctx context.Context, db DB, params *CacheDemoTwoWorkItemCreateParams) (*CacheDemoTwoWorkItem, error) {
	cdtwi := &CacheDemoTwoWorkItem{
		ClosedAt:              params.ClosedAt,
		CustomDateForProject2: params.CustomDateForProject2,
		Description:           params.Description,
		KanbanStepID:          params.KanbanStepID,
		Metadata:              params.Metadata,
		TargetDate:            params.TargetDate,
		TeamID:                params.TeamID,
		Title:                 params.Title,
		WorkItemID:            params.WorkItemID,
		WorkItemTypeID:        params.WorkItemTypeID,
	}

	return cdtwi.Insert(ctx, db)
}

type CacheDemoTwoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   CacheDemoTwoWorkItemJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
}
type CacheDemoTwoWorkItemSelectConfigOption func(*CacheDemoTwoWorkItemSelectConfig)

// WithCacheDemoTwoWorkItemLimit limits row selection.
func WithCacheDemoTwoWorkItemLimit(limit int) CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedCacheDemoTwoWorkItemOnly limits result to records marked as deleted.
func WithDeletedCacheDemoTwoWorkItemOnly() CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		s.deletedAt = " not null "
	}
}

type CacheDemoTwoWorkItemOrderBy string

const (
	CacheDemoTwoWorkItemClosedAtDescNullsFirst              CacheDemoTwoWorkItemOrderBy = " closed_at DESC NULLS FIRST "
	CacheDemoTwoWorkItemClosedAtDescNullsLast               CacheDemoTwoWorkItemOrderBy = " closed_at DESC NULLS LAST "
	CacheDemoTwoWorkItemClosedAtAscNullsFirst               CacheDemoTwoWorkItemOrderBy = " closed_at ASC NULLS FIRST "
	CacheDemoTwoWorkItemClosedAtAscNullsLast                CacheDemoTwoWorkItemOrderBy = " closed_at ASC NULLS LAST "
	CacheDemoTwoWorkItemCreatedAtDescNullsFirst             CacheDemoTwoWorkItemOrderBy = " created_at DESC NULLS FIRST "
	CacheDemoTwoWorkItemCreatedAtDescNullsLast              CacheDemoTwoWorkItemOrderBy = " created_at DESC NULLS LAST "
	CacheDemoTwoWorkItemCreatedAtAscNullsFirst              CacheDemoTwoWorkItemOrderBy = " created_at ASC NULLS FIRST "
	CacheDemoTwoWorkItemCreatedAtAscNullsLast               CacheDemoTwoWorkItemOrderBy = " created_at ASC NULLS LAST "
	CacheDemoTwoWorkItemCustomDateForProject2DescNullsFirst CacheDemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS FIRST "
	CacheDemoTwoWorkItemCustomDateForProject2DescNullsLast  CacheDemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS LAST "
	CacheDemoTwoWorkItemCustomDateForProject2AscNullsFirst  CacheDemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS FIRST "
	CacheDemoTwoWorkItemCustomDateForProject2AscNullsLast   CacheDemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS LAST "
	CacheDemoTwoWorkItemDeletedAtDescNullsFirst             CacheDemoTwoWorkItemOrderBy = " deleted_at DESC NULLS FIRST "
	CacheDemoTwoWorkItemDeletedAtDescNullsLast              CacheDemoTwoWorkItemOrderBy = " deleted_at DESC NULLS LAST "
	CacheDemoTwoWorkItemDeletedAtAscNullsFirst              CacheDemoTwoWorkItemOrderBy = " deleted_at ASC NULLS FIRST "
	CacheDemoTwoWorkItemDeletedAtAscNullsLast               CacheDemoTwoWorkItemOrderBy = " deleted_at ASC NULLS LAST "
	CacheDemoTwoWorkItemTargetDateDescNullsFirst            CacheDemoTwoWorkItemOrderBy = " target_date DESC NULLS FIRST "
	CacheDemoTwoWorkItemTargetDateDescNullsLast             CacheDemoTwoWorkItemOrderBy = " target_date DESC NULLS LAST "
	CacheDemoTwoWorkItemTargetDateAscNullsFirst             CacheDemoTwoWorkItemOrderBy = " target_date ASC NULLS FIRST "
	CacheDemoTwoWorkItemTargetDateAscNullsLast              CacheDemoTwoWorkItemOrderBy = " target_date ASC NULLS LAST "
	CacheDemoTwoWorkItemUpdatedAtDescNullsFirst             CacheDemoTwoWorkItemOrderBy = " updated_at DESC NULLS FIRST "
	CacheDemoTwoWorkItemUpdatedAtDescNullsLast              CacheDemoTwoWorkItemOrderBy = " updated_at DESC NULLS LAST "
	CacheDemoTwoWorkItemUpdatedAtAscNullsFirst              CacheDemoTwoWorkItemOrderBy = " updated_at ASC NULLS FIRST "
	CacheDemoTwoWorkItemUpdatedAtAscNullsLast               CacheDemoTwoWorkItemOrderBy = " updated_at ASC NULLS LAST "
)

// WithCacheDemoTwoWorkItemOrderBy orders results by the given columns.
func WithCacheDemoTwoWorkItemOrderBy(rows ...CacheDemoTwoWorkItemOrderBy) CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		if len(rows) > 0 {
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
		}
	}
}

type CacheDemoTwoWorkItemJoins struct {
	KanbanStep       bool // O2O kanban_steps
	Team             bool // O2O teams
	WorkItemType     bool // O2O work_item_types
	TimeEntries      bool // M2O time_entries
	AssignedUsers    bool // M2M work_item_assigned_user
	WorkItemComments bool // M2O work_item_comments
	WorkItemTags     bool // M2M work_item_work_item_tag
}

// WithCacheDemoTwoWorkItemJoin joins with the given tables.
func WithCacheDemoTwoWorkItemJoin(joins CacheDemoTwoWorkItemJoins) CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		s.joins = CacheDemoTwoWorkItemJoins{
			KanbanStep:       s.joins.KanbanStep || joins.KanbanStep,
			Team:             s.joins.Team || joins.Team,
			WorkItemType:     s.joins.WorkItemType || joins.WorkItemType,
			TimeEntries:      s.joins.TimeEntries || joins.TimeEntries,
			AssignedUsers:    s.joins.AssignedUsers || joins.AssignedUsers,
			WorkItemComments: s.joins.WorkItemComments || joins.WorkItemComments,
			WorkItemTags:     s.joins.WorkItemTags || joins.WorkItemTags,
		}
	}
}

// User__WIAU_CacheDemoTwoWorkItem represents a M2M join against "public.work_item_assigned_user"
type User__WIAU_CacheDemoTwoWorkItem struct {
	User User                `json:"user" db:"users" required:"true"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithCacheDemoTwoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithCacheDemoTwoWorkItemFilters(filters map[string][]any) CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithCacheDemoTwoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithCacheDemoTwoWorkItemHavingClause(conditions map[string][]any) CacheDemoTwoWorkItemSelectConfigOption {
	return func(s *CacheDemoTwoWorkItemSelectConfig) {
		s.having = conditions
	}
}

const cacheDemoTwoWorkItemTableKanbanStepJoinSQL = `-- O2O join generated from "cache__demo_two_work_items_kanban_step_id_fkey (inferred)"
left join kanban_steps as _cache__demo_two_work_items_kanban_step_id on _cache__demo_two_work_items_kanban_step_id.kanban_step_id = cache__demo_two_work_items.kanban_step_id
`

const cacheDemoTwoWorkItemTableKanbanStepSelectSQL = `(case when _cache__demo_two_work_items_kanban_step_id.kanban_step_id is not null then row(_cache__demo_two_work_items_kanban_step_id.*) end) as kanban_step_kanban_step_id`

const cacheDemoTwoWorkItemTableKanbanStepGroupBySQL = `_cache__demo_two_work_items_kanban_step_id.kanban_step_id,
      _cache__demo_two_work_items_kanban_step_id.kanban_step_id,
	cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableTeamJoinSQL = `-- O2O join generated from "cache__demo_two_work_items_team_id_fkey (inferred)"
left join teams as _cache__demo_two_work_items_team_id on _cache__demo_two_work_items_team_id.team_id = cache__demo_two_work_items.team_id
`

const cacheDemoTwoWorkItemTableTeamSelectSQL = `(case when _cache__demo_two_work_items_team_id.team_id is not null then row(_cache__demo_two_work_items_team_id.*) end) as team_team_id`

const cacheDemoTwoWorkItemTableTeamGroupBySQL = `_cache__demo_two_work_items_team_id.team_id,
      _cache__demo_two_work_items_team_id.team_id,
	cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableWorkItemTypeJoinSQL = `-- O2O join generated from "cache__demo_two_work_items_work_item_type_id_fkey (inferred)"
left join work_item_types as _cache__demo_two_work_items_work_item_type_id on _cache__demo_two_work_items_work_item_type_id.work_item_type_id = cache__demo_two_work_items.work_item_type_id
`

const cacheDemoTwoWorkItemTableWorkItemTypeSelectSQL = `(case when _cache__demo_two_work_items_work_item_type_id.work_item_type_id is not null then row(_cache__demo_two_work_items_work_item_type_id.*) end) as work_item_type_work_item_type_id`

const cacheDemoTwoWorkItemTableWorkItemTypeGroupBySQL = `_cache__demo_two_work_items_work_item_type_id.work_item_type_id,
      _cache__demo_two_work_items_work_item_type_id.work_item_type_id,
	cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_work_item_id_fkey-shared-ref-cache__demo_two_work_items"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id
) as joined_time_entries on joined_time_entries.time_entries_work_item_id = cache__demo_two_work_items.work_item_id
`

const cacheDemoTwoWorkItemTableTimeEntriesSelectSQL = `COALESCE(joined_time_entries.time_entries, '{}') as time_entries`

const cacheDemoTwoWorkItemTableTimeEntriesGroupBySQL = `joined_time_entries.time_entries, cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey-shared-ref-cache__demo_two_work_items"
left join (
	select
		work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
		, work_item_assigned_user.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		work_item_assigned_user
	join users on users.user_id = work_item_assigned_user.assigned_user
	group by
		work_item_assigned_user_work_item_id
		, users.user_id
		, role
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = cache__demo_two_work_items.work_item_id
`

const cacheDemoTwoWorkItemTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const cacheDemoTwoWorkItemTableAssignedUsersGroupBySQL = `cache__demo_two_work_items.work_item_id, cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableWorkItemCommentsJoinSQL = `-- M2O join generated from "work_item_comments_work_item_id_fkey-shared-ref-cache__demo_two_work_items"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
  group by
        work_item_id
) as joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = cache__demo_two_work_items.work_item_id
`

const cacheDemoTwoWorkItemTableWorkItemCommentsSelectSQL = `COALESCE(joined_work_item_comments.work_item_comments, '{}') as work_item_comments`

const cacheDemoTwoWorkItemTableWorkItemCommentsGroupBySQL = `joined_work_item_comments.work_item_comments, cache__demo_two_work_items.work_item_id`

const cacheDemoTwoWorkItemTableWorkItemTagsJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey-shared-ref-cache__demo_two_work_items"
left join (
	select
		work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
		, work_item_tags.work_item_tag_id as __work_item_tags_work_item_tag_id
		, row(work_item_tags.*) as __work_item_tags
	from
		work_item_work_item_tag
	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
	group by
		work_item_work_item_tag_work_item_id
		, work_item_tags.work_item_tag_id
) as joined_work_item_work_item_tag_work_item_tags on joined_work_item_work_item_tag_work_item_tags.work_item_work_item_tag_work_item_id = cache__demo_two_work_items.work_item_id
`

const cacheDemoTwoWorkItemTableWorkItemTagsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_item_tags.__work_item_tags
		)) filter (where joined_work_item_work_item_tag_work_item_tags.__work_item_tags_work_item_tag_id is not null), '{}') as work_item_work_item_tag_work_item_tags`

const cacheDemoTwoWorkItemTableWorkItemTagsGroupBySQL = `cache__demo_two_work_items.work_item_id, cache__demo_two_work_items.work_item_id`

// CacheDemoTwoWorkItemUpdateParams represents update params for 'public.cache__demo_two_work_items'.
type CacheDemoTwoWorkItemUpdateParams struct {
	ClosedAt              **time.Time     `json:"closedAt"`                        // closed_at
	CustomDateForProject2 **time.Time     `json:"customDateForProject2"`           // custom_date_for_project_2
	Description           *string         `json:"description" nullable:"false"`    // description
	KanbanStepID          *KanbanStepID   `json:"kanbanStepID" nullable:"false"`   // kanban_step_id
	Metadata              *map[string]any `json:"metadata" nullable:"false"`       // metadata
	TargetDate            *time.Time      `json:"targetDate" nullable:"false"`     // target_date
	TeamID                *TeamID         `json:"teamID" nullable:"false"`         // team_id
	Title                 *string         `json:"title" nullable:"false"`          // title
	WorkItemTypeID        *WorkItemTypeID `json:"workItemTypeID" nullable:"false"` // work_item_type_id
}

// SetUpdateParams updates public.cache__demo_two_work_items struct fields with the specified params.
func (cdtwi *CacheDemoTwoWorkItem) SetUpdateParams(params *CacheDemoTwoWorkItemUpdateParams) {
	if params.ClosedAt != nil {
		cdtwi.ClosedAt = *params.ClosedAt
	}
	if params.CustomDateForProject2 != nil {
		cdtwi.CustomDateForProject2 = *params.CustomDateForProject2
	}
	if params.Description != nil {
		cdtwi.Description = *params.Description
	}
	if params.KanbanStepID != nil {
		cdtwi.KanbanStepID = *params.KanbanStepID
	}
	if params.Metadata != nil {
		cdtwi.Metadata = *params.Metadata
	}
	if params.TargetDate != nil {
		cdtwi.TargetDate = *params.TargetDate
	}
	if params.TeamID != nil {
		cdtwi.TeamID = *params.TeamID
	}
	if params.Title != nil {
		cdtwi.Title = *params.Title
	}
	if params.WorkItemTypeID != nil {
		cdtwi.WorkItemTypeID = *params.WorkItemTypeID
	}
}

// Insert inserts the CacheDemoTwoWorkItem to the database.
func (cdtwi *CacheDemoTwoWorkItem) Insert(ctx context.Context, db DB) (*CacheDemoTwoWorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.cache__demo_two_work_items (
	closed_at, custom_date_for_project_2, deleted_at, description, kanban_step_id, metadata, target_date, team_id, title, work_item_id, work_item_type_id
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
	) RETURNING * `
	// run
	logf(sqlstr, cdtwi.ClosedAt, cdtwi.CustomDateForProject2, cdtwi.DeletedAt, cdtwi.Description, cdtwi.KanbanStepID, cdtwi.Metadata, cdtwi.TargetDate, cdtwi.TeamID, cdtwi.Title, cdtwi.WorkItemID, cdtwi.WorkItemTypeID)

	rows, err := db.Query(ctx, sqlstr, cdtwi.ClosedAt, cdtwi.CustomDateForProject2, cdtwi.DeletedAt, cdtwi.Description, cdtwi.KanbanStepID, cdtwi.Metadata, cdtwi.TargetDate, cdtwi.TeamID, cdtwi.Title, cdtwi.WorkItemID, cdtwi.WorkItemTypeID)
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Insert/db.Query: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	newcdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[CacheDemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}

	*cdtwi = newcdtwi

	return cdtwi, nil
}

// Update updates a CacheDemoTwoWorkItem in the database.
func (cdtwi *CacheDemoTwoWorkItem) Update(ctx context.Context, db DB) (*CacheDemoTwoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.cache__demo_two_work_items SET 
	closed_at = $1, custom_date_for_project_2 = $2, deleted_at = $3, description = $4, kanban_step_id = $5, metadata = $6, target_date = $7, team_id = $8, title = $9, work_item_type_id = $10 
	WHERE work_item_id = $11 
	RETURNING * `
	// run
	logf(sqlstr, cdtwi.ClosedAt, cdtwi.CreatedAt, cdtwi.CustomDateForProject2, cdtwi.DeletedAt, cdtwi.Description, cdtwi.KanbanStepID, cdtwi.Metadata, cdtwi.TargetDate, cdtwi.TeamID, cdtwi.Title, cdtwi.UpdatedAt, cdtwi.WorkItemTypeID, cdtwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, cdtwi.ClosedAt, cdtwi.CustomDateForProject2, cdtwi.DeletedAt, cdtwi.Description, cdtwi.KanbanStepID, cdtwi.Metadata, cdtwi.TargetDate, cdtwi.TeamID, cdtwi.Title, cdtwi.WorkItemTypeID, cdtwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Update/db.Query: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	newcdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[CacheDemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	*cdtwi = newcdtwi

	return cdtwi, nil
}

// Upsert upserts a CacheDemoTwoWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (cdtwi *CacheDemoTwoWorkItem) Upsert(ctx context.Context, db DB, params *CacheDemoTwoWorkItemCreateParams) (*CacheDemoTwoWorkItem, error) {
	var err error

	cdtwi.ClosedAt = params.ClosedAt
	cdtwi.CustomDateForProject2 = params.CustomDateForProject2
	cdtwi.Description = params.Description
	cdtwi.KanbanStepID = params.KanbanStepID
	cdtwi.Metadata = params.Metadata
	cdtwi.TargetDate = params.TargetDate
	cdtwi.TeamID = params.TeamID
	cdtwi.Title = params.Title
	cdtwi.WorkItemID = params.WorkItemID
	cdtwi.WorkItemTypeID = params.WorkItemTypeID

	cdtwi, err = cdtwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Cache  demo two work item", Err: err})
			}
			cdtwi, err = cdtwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Cache  demo two work item", Err: err})
			}
		}
	}

	return cdtwi, err
}

// Delete deletes the CacheDemoTwoWorkItem from the database.
func (cdtwi *CacheDemoTwoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.cache__demo_two_work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, cdtwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the CacheDemoTwoWorkItem from the database via 'deleted_at'.
func (cdtwi *CacheDemoTwoWorkItem) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.cache__demo_two_work_items 
	SET deleted_at = NOW() 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, cdtwi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	cdtwi.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted CacheDemoTwoWorkItem from the database.
func (cdtwi *CacheDemoTwoWorkItem) Restore(ctx context.Context, db DB) (*CacheDemoTwoWorkItem, error) {
	cdtwi.DeletedAt = nil
	newcdtwi, err := cdtwi.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Restore/pgx.CollectRows: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	return newcdtwi, nil
}

// CacheDemoTwoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of CacheDemoTwoWorkItem.
func CacheDemoTwoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID int, direction models.Direction, opts ...CacheDemoTwoWorkItemSelectConfigOption) ([]CacheDemoTwoWorkItem, error) {
	c := &CacheDemoTwoWorkItemSelectConfig{deletedAt: " null ", joins: CacheDemoTwoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemTypeGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemTagsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	cache__demo_two_work_items.closed_at,
	cache__demo_two_work_items.created_at,
	cache__demo_two_work_items.custom_date_for_project_2,
	cache__demo_two_work_items.deleted_at,
	cache__demo_two_work_items.description,
	cache__demo_two_work_items.kanban_step_id,
	cache__demo_two_work_items.metadata,
	cache__demo_two_work_items.target_date,
	cache__demo_two_work_items.team_id,
	cache__demo_two_work_items.title,
	cache__demo_two_work_items.updated_at,
	cache__demo_two_work_items.work_item_id,
	cache__demo_two_work_items.work_item_type_id %s 
	 FROM public.cache__demo_two_work_items %s 
	 WHERE cache__demo_two_work_items.work_item_id %s $1
	 %s   AND cache__demo_two_work_items.deleted_at is %s  %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* CacheDemoTwoWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[CacheDemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoTwoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	return res, nil
}

// CacheDemoTwoWorkItemByWorkItemID retrieves a row from 'public.cache__demo_two_work_items' as a CacheDemoTwoWorkItem.
//
// Generated from index 'cache__demo_two_work_items_pkey'.
func CacheDemoTwoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int, opts ...CacheDemoTwoWorkItemSelectConfigOption) (*CacheDemoTwoWorkItem, error) {
	c := &CacheDemoTwoWorkItemSelectConfig{deletedAt: " null ", joins: CacheDemoTwoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemTypeGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, cacheDemoTwoWorkItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, cacheDemoTwoWorkItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, cacheDemoTwoWorkItemTableWorkItemTagsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	cache__demo_two_work_items.closed_at,
	cache__demo_two_work_items.created_at,
	cache__demo_two_work_items.custom_date_for_project_2,
	cache__demo_two_work_items.deleted_at,
	cache__demo_two_work_items.description,
	cache__demo_two_work_items.kanban_step_id,
	cache__demo_two_work_items.metadata,
	cache__demo_two_work_items.target_date,
	cache__demo_two_work_items.team_id,
	cache__demo_two_work_items.title,
	cache__demo_two_work_items.updated_at,
	cache__demo_two_work_items.work_item_id,
	cache__demo_two_work_items.work_item_type_id %s 
	 FROM public.cache__demo_two_work_items %s 
	 WHERE cache__demo_two_work_items.work_item_id = $1
	 %s   AND cache__demo_two_work_items.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* CacheDemoTwoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("cache__demo_two_work_items/CacheDemoTwoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}
	cdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[CacheDemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("cache__demo_two_work_items/CacheDemoTwoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Cache  demo two work item", Err: err}))
	}

	return &cdtwi, nil
}

// FKKanbanStep_KanbanStepID returns the KanbanStep associated with the CacheDemoTwoWorkItem's (KanbanStepID).
//
// Generated from foreign key 'cache__demo_two_work_items_kanban_step_id_fkey'.
func (cdtwi *CacheDemoTwoWorkItem) FKKanbanStep_KanbanStepID(ctx context.Context, db DB) (*KanbanStep, error) {
	return KanbanStepByKanbanStepID(ctx, db, cdtwi.KanbanStepID)
}

// FKTeam_TeamID returns the Team associated with the CacheDemoTwoWorkItem's (TeamID).
//
// Generated from foreign key 'cache__demo_two_work_items_team_id_fkey'.
func (cdtwi *CacheDemoTwoWorkItem) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, cdtwi.TeamID)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the CacheDemoTwoWorkItem's (WorkItemID).
//
// Generated from foreign key 'cache__demo_two_work_items_work_item_id_fkey'.
func (cdtwi *CacheDemoTwoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, cdtwi.WorkItemID)
}

// FKWorkItemType_WorkItemTypeID returns the WorkItemType associated with the CacheDemoTwoWorkItem's (WorkItemTypeID).
//
// Generated from foreign key 'cache__demo_two_work_items_work_item_type_id_fkey'.
func (cdtwi *CacheDemoTwoWorkItem) FKWorkItemType_WorkItemTypeID(ctx context.Context, db DB) (*WorkItemType, error) {
	return WorkItemTypeByWorkItemTypeID(ctx, db, cdtwi.WorkItemTypeID)
}
