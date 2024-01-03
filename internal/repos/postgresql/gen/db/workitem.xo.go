package db

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

// WorkItem represents a row from 'public.work_items'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItem struct {
	WorkItemID     WorkItemID     `db:"work_item_id"      json:"workItemID"     nullable:"false" required:"true"` // work_item_id
	Title          string         `db:"title"             json:"title"          nullable:"false" required:"true"` // title
	Description    string         `db:"description"       json:"description"    nullable:"false" required:"true"` // description
	WorkItemTypeID WorkItemTypeID `db:"work_item_type_id" json:"workItemTypeID" nullable:"false" required:"true"` // work_item_type_id
	Metadata       map[string]any `db:"metadata"          json:"metadata"       nullable:"false" required:"true"` // metadata
	TeamID         TeamID         `db:"team_id"           json:"teamID"         nullable:"false" required:"true"` // team_id
	KanbanStepID   KanbanStepID   `db:"kanban_step_id"    json:"kanbanStepID"   nullable:"false" required:"true"` // kanban_step_id
	ClosedAt       *time.Time     `db:"closed_at"         json:"closedAt"`                                        // closed_at
	TargetDate     time.Time      `db:"target_date"       json:"targetDate"     nullable:"false" required:"true"` // target_date
	CreatedAt      time.Time      `db:"created_at"        json:"createdAt"      nullable:"false" required:"true"` // created_at
	UpdatedAt      time.Time      `db:"updated_at"        json:"updatedAt"      nullable:"false" required:"true"` // updated_at
	DeletedAt      *time.Time     `db:"deleted_at"        json:"deletedAt"`                                       // deleted_at

	DemoTwoWorkItemJoin          *DemoTwoWorkItem       `db:"demo_two_work_item_work_item_id"        json:"-" openapi-go:"ignore"` // O2O demo_two_work_items (inferred)
	DemoWorkItemJoin             *DemoWorkItem          `db:"demo_work_item_work_item_id"            json:"-" openapi-go:"ignore"` // O2O demo_work_items (inferred)
	WorkItemTimeEntriesJoin      *[]TimeEntry           `db:"time_entries"                           json:"-" openapi-go:"ignore"` // M2O work_items
	WorkItemAssignedUsersJoin    *[]User__WIAU_WorkItem `db:"work_item_assigned_user_assigned_users" json:"-" openapi-go:"ignore"` // M2M work_item_assigned_user
	WorkItemWorkItemCommentsJoin *[]WorkItemComment     `db:"work_item_comments"                     json:"-" openapi-go:"ignore"` // M2O work_items
	WorkItemWorkItemTagsJoin     *[]WorkItemTag         `db:"work_item_work_item_tag_work_item_tags" json:"-" openapi-go:"ignore"` // M2M work_item_work_item_tag
	KanbanStepJoin               *KanbanStep            `db:"kanban_step_kanban_step_id"             json:"-" openapi-go:"ignore"` // O2O kanban_steps (inferred)
	TeamJoin                     *Team                  `db:"team_team_id"                           json:"-" openapi-go:"ignore"` // O2O teams (inferred)
	WorkItemTypeJoin             *WorkItemType          `db:"work_item_type_work_item_type_id"       json:"-" openapi-go:"ignore"` // O2O work_item_types (inferred)
}

// WorkItemCreateParams represents insert params for 'public.work_items'.
type WorkItemCreateParams struct {
	ClosedAt       *time.Time     `json:"closedAt"`                                        // closed_at
	Description    string         `json:"description"    nullable:"false" required:"true"` // description
	KanbanStepID   KanbanStepID   `json:"kanbanStepID"   nullable:"false" required:"true"` // kanban_step_id
	Metadata       map[string]any `json:"metadata"       nullable:"false" required:"true"` // metadata
	TargetDate     time.Time      `json:"targetDate"     nullable:"false" required:"true"` // target_date
	TeamID         TeamID         `json:"teamID"         nullable:"false" required:"true"` // team_id
	Title          string         `json:"title"          nullable:"false" required:"true"` // title
	WorkItemTypeID WorkItemTypeID `json:"workItemTypeID" nullable:"false" required:"true"` // work_item_type_id
}

type WorkItemID int

// CreateWorkItem creates a new WorkItem in the database with the given params.
func CreateWorkItem(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	wi := &WorkItem{
		ClosedAt:       params.ClosedAt,
		Description:    params.Description,
		KanbanStepID:   params.KanbanStepID,
		Metadata:       params.Metadata,
		TargetDate:     params.TargetDate,
		TeamID:         params.TeamID,
		Title:          params.Title,
		WorkItemTypeID: params.WorkItemTypeID,
	}

	return wi.Insert(ctx, db)
}

// WorkItemUpdateParams represents update params for 'public.work_items'.
type WorkItemUpdateParams struct {
	ClosedAt       **time.Time     `json:"closedAt"`                        // closed_at
	Description    *string         `json:"description"    nullable:"false"` // description
	KanbanStepID   *KanbanStepID   `json:"kanbanStepID"   nullable:"false"` // kanban_step_id
	Metadata       *map[string]any `json:"metadata"       nullable:"false"` // metadata
	TargetDate     *time.Time      `json:"targetDate"     nullable:"false"` // target_date
	TeamID         *TeamID         `json:"teamID"         nullable:"false"` // team_id
	Title          *string         `json:"title"          nullable:"false"` // title
	WorkItemTypeID *WorkItemTypeID `json:"workItemTypeID" nullable:"false"` // work_item_type_id
}

// SetUpdateParams updates public.work_items struct fields with the specified params.
func (wi *WorkItem) SetUpdateParams(params *WorkItemUpdateParams) {
	if params.ClosedAt != nil {
		wi.ClosedAt = *params.ClosedAt
	}
	if params.Description != nil {
		wi.Description = *params.Description
	}
	if params.KanbanStepID != nil {
		wi.KanbanStepID = *params.KanbanStepID
	}
	if params.Metadata != nil {
		wi.Metadata = *params.Metadata
	}
	if params.TargetDate != nil {
		wi.TargetDate = *params.TargetDate
	}
	if params.TeamID != nil {
		wi.TeamID = *params.TeamID
	}
	if params.Title != nil {
		wi.Title = *params.Title
	}
	if params.WorkItemTypeID != nil {
		wi.WorkItemTypeID = *params.WorkItemTypeID
	}
}

type WorkItemSelectConfig struct {
	limit     string
	orderBy   string
	joins     WorkItemJoins
	filters   map[string][]any
	having    map[string][]any
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

type WorkItemOrderBy string

const (
	WorkItemClosedAtDescNullsFirst   WorkItemOrderBy = " closed_at DESC NULLS FIRST "
	WorkItemClosedAtDescNullsLast    WorkItemOrderBy = " closed_at DESC NULLS LAST "
	WorkItemClosedAtAscNullsFirst    WorkItemOrderBy = " closed_at ASC NULLS FIRST "
	WorkItemClosedAtAscNullsLast     WorkItemOrderBy = " closed_at ASC NULLS LAST "
	WorkItemCreatedAtDescNullsFirst  WorkItemOrderBy = " created_at DESC NULLS FIRST "
	WorkItemCreatedAtDescNullsLast   WorkItemOrderBy = " created_at DESC NULLS LAST "
	WorkItemCreatedAtAscNullsFirst   WorkItemOrderBy = " created_at ASC NULLS FIRST "
	WorkItemCreatedAtAscNullsLast    WorkItemOrderBy = " created_at ASC NULLS LAST "
	WorkItemDeletedAtDescNullsFirst  WorkItemOrderBy = " deleted_at DESC NULLS FIRST "
	WorkItemDeletedAtDescNullsLast   WorkItemOrderBy = " deleted_at DESC NULLS LAST "
	WorkItemDeletedAtAscNullsFirst   WorkItemOrderBy = " deleted_at ASC NULLS FIRST "
	WorkItemDeletedAtAscNullsLast    WorkItemOrderBy = " deleted_at ASC NULLS LAST "
	WorkItemTargetDateDescNullsFirst WorkItemOrderBy = " target_date DESC NULLS FIRST "
	WorkItemTargetDateDescNullsLast  WorkItemOrderBy = " target_date DESC NULLS LAST "
	WorkItemTargetDateAscNullsFirst  WorkItemOrderBy = " target_date ASC NULLS FIRST "
	WorkItemTargetDateAscNullsLast   WorkItemOrderBy = " target_date ASC NULLS LAST "
	WorkItemUpdatedAtDescNullsFirst  WorkItemOrderBy = " updated_at DESC NULLS FIRST "
	WorkItemUpdatedAtDescNullsLast   WorkItemOrderBy = " updated_at DESC NULLS LAST "
	WorkItemUpdatedAtAscNullsFirst   WorkItemOrderBy = " updated_at ASC NULLS FIRST "
	WorkItemUpdatedAtAscNullsLast    WorkItemOrderBy = " updated_at ASC NULLS LAST "
)

// WithWorkItemOrderBy orders results by the given columns.
func WithWorkItemOrderBy(rows ...WorkItemOrderBy) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
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

type WorkItemJoins struct {
	DemoTwoWorkItem  bool // O2O demo_two_work_items
	DemoWorkItem     bool // O2O demo_work_items
	TimeEntries      bool // M2O time_entries
	AssignedUsers    bool // M2M work_item_assigned_user
	WorkItemComments bool // M2O work_item_comments
	WorkItemTags     bool // M2M work_item_work_item_tag
	KanbanStep       bool // O2O kanban_steps
	Team             bool // O2O teams
	WorkItemType     bool // O2O work_item_types
}

// WithWorkItemJoin joins with the given tables.
func WithWorkItemJoin(joins WorkItemJoins) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.joins = WorkItemJoins{
			DemoTwoWorkItem:  s.joins.DemoTwoWorkItem || joins.DemoTwoWorkItem,
			DemoWorkItem:     s.joins.DemoWorkItem || joins.DemoWorkItem,
			TimeEntries:      s.joins.TimeEntries || joins.TimeEntries,
			AssignedUsers:    s.joins.AssignedUsers || joins.AssignedUsers,
			WorkItemComments: s.joins.WorkItemComments || joins.WorkItemComments,
			WorkItemTags:     s.joins.WorkItemTags || joins.WorkItemTags,
			KanbanStep:       s.joins.KanbanStep || joins.KanbanStep,
			Team:             s.joins.Team || joins.Team,
			WorkItemType:     s.joins.WorkItemType || joins.WorkItemType,
		}
	}
}

// User__WIAU_WorkItem represents a M2M join against "public.work_item_assigned_user".
type User__WIAU_WorkItem struct {
	User User                `db:"users" json:"user" required:"true"`
	Role models.WorkItemRole `db:"role"  json:"role" ref:"#/components/schemas/WorkItemRole" required:"true"`
}

// WithWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemFilters(filters map[string][]any) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.filters = filters
	}
}

func WithWorkItemHavingClause(clauses map[string][]any) WorkItemSelectConfigOption {
	return func(s *WorkItemSelectConfig) {
		s.having = clauses
	}
}

const workItemTableDemoTwoWorkItemJoinSQL = `-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join demo_two_work_items as _demo_two_work_items_work_item_id on _demo_two_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const workItemTableDemoTwoWorkItemSelectSQL = `(case when _demo_two_work_items_work_item_id.work_item_id is not null then row(_demo_two_work_items_work_item_id.*) end) as demo_two_work_item_work_item_id`

const workItemTableDemoTwoWorkItemGroupBySQL = `_demo_two_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const workItemTableDemoWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const workItemTableDemoWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id`

const workItemTableDemoWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const workItemTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id
) as joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
`

const workItemTableTimeEntriesSelectSQL = `COALESCE(joined_time_entries.time_entries, '{}') as time_entries`

const workItemTableTimeEntriesGroupBySQL = `joined_time_entries.time_entries, work_items.work_item_id`

const workItemTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
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
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_items.work_item_id
`

const workItemTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const workItemTableAssignedUsersGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

const workItemTableWorkItemCommentsJoinSQL = `-- M2O join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
  group by
        work_item_id
) as joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
`

const workItemTableWorkItemCommentsSelectSQL = `COALESCE(joined_work_item_comments.work_item_comments, '{}') as work_item_comments`

const workItemTableWorkItemCommentsGroupBySQL = `joined_work_item_comments.work_item_comments, work_items.work_item_id`

const workItemTableWorkItemTagsJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
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
) as joined_work_item_work_item_tag_work_item_tags on joined_work_item_work_item_tag_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id
`

const workItemTableWorkItemTagsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_item_tags.__work_item_tags
		)) filter (where joined_work_item_work_item_tag_work_item_tags.__work_item_tags_work_item_tag_id is not null), '{}') as work_item_work_item_tag_work_item_tags`

const workItemTableWorkItemTagsGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

const workItemTableKanbanStepJoinSQL = `-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join kanban_steps as _work_items_kanban_step_id on _work_items_kanban_step_id.kanban_step_id = work_items.kanban_step_id
`

const workItemTableKanbanStepSelectSQL = `(case when _work_items_kanban_step_id.kanban_step_id is not null then row(_work_items_kanban_step_id.*) end) as kanban_step_kanban_step_id`

const workItemTableKanbanStepGroupBySQL = `_work_items_kanban_step_id.kanban_step_id,
      _work_items_kanban_step_id.kanban_step_id,
	work_items.work_item_id`

const workItemTableTeamJoinSQL = `-- O2O join generated from "work_items_team_id_fkey (inferred)"
left join teams as _work_items_team_id on _work_items_team_id.team_id = work_items.team_id
`

const workItemTableTeamSelectSQL = `(case when _work_items_team_id.team_id is not null then row(_work_items_team_id.*) end) as team_team_id`

const workItemTableTeamGroupBySQL = `_work_items_team_id.team_id,
      _work_items_team_id.team_id,
	work_items.work_item_id`

const workItemTableWorkItemTypeJoinSQL = `-- O2O join generated from "work_items_work_item_type_id_fkey (inferred)"
left join work_item_types as _work_items_work_item_type_id on _work_items_work_item_type_id.work_item_type_id = work_items.work_item_type_id
`

const workItemTableWorkItemTypeSelectSQL = `(case when _work_items_work_item_type_id.work_item_type_id is not null then row(_work_items_work_item_type_id.*) end) as work_item_type_work_item_type_id`

const workItemTableWorkItemTypeGroupBySQL = `_work_items_work_item_type_id.work_item_type_id,
      _work_items_work_item_type_id.work_item_type_id,
	work_items.work_item_id`

// Insert inserts the WorkItem to the database.
func (wi *WorkItem) Insert(ctx context.Context, db DB) (*WorkItem, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_items (
	closed_at, deleted_at, description, kanban_step_id, metadata, target_date, team_id, title, work_item_type_id
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9
	) RETURNING * `
	// run
	logf(sqlstr, wi.ClosedAt, wi.DeletedAt, wi.Description, wi.KanbanStepID, wi.Metadata, wi.TargetDate, wi.TeamID, wi.Title, wi.WorkItemTypeID)

	rows, err := db.Query(ctx, sqlstr, wi.ClosedAt, wi.DeletedAt, wi.Description, wi.KanbanStepID, wi.Metadata, wi.TargetDate, wi.TeamID, wi.Title, wi.WorkItemTypeID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	*wi = newwi

	return wi, nil
}

// Update updates a WorkItem in the database.
func (wi *WorkItem) Update(ctx context.Context, db DB) (*WorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_items SET
	closed_at = $1, deleted_at = $2, description = $3, kanban_step_id = $4, metadata = $5, target_date = $6, team_id = $7, title = $8, work_item_type_id = $9
	WHERE work_item_id = $10
	RETURNING * `
	// run
	logf(sqlstr, wi.ClosedAt, wi.CreatedAt, wi.DeletedAt, wi.Description, wi.KanbanStepID, wi.Metadata, wi.TargetDate, wi.TeamID, wi.Title, wi.UpdatedAt, wi.WorkItemTypeID, wi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, wi.ClosedAt, wi.DeletedAt, wi.Description, wi.KanbanStepID, wi.Metadata, wi.TargetDate, wi.TeamID, wi.Title, wi.WorkItemTypeID, wi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	newwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}
	*wi = newwi

	return wi, nil
}

// Upsert upserts a WorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wi *WorkItem) Upsert(ctx context.Context, db DB, params *WorkItemCreateParams) (*WorkItem, error) {
	var err error

	wi.ClosedAt = params.ClosedAt
	wi.Description = params.Description
	wi.KanbanStepID = params.KanbanStepID
	wi.Metadata = params.Metadata
	wi.TargetDate = params.TargetDate
	wi.TeamID = params.TeamID
	wi.Title = params.Title
	wi.WorkItemTypeID = params.WorkItemTypeID

	wi, err = wi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item", Err: err})
			}
			wi, err = wi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item", Err: err})
			}
		}
	}

	return wi, err
}

// Delete deletes the WorkItem from the database.
func (wi *WorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_items
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the WorkItem from the database via 'deleted_at'.
func (wi *WorkItem) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.work_items
	SET deleted_at = NOW()
	WHERE work_item_id = $1 `
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
		return nil, logerror(fmt.Errorf("WorkItem/Restore/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return newwi, nil
}

// WorkItemPaginatedByWorkItemID returns a cursor-paginated list of WorkItem.
func WorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, direction models.Direction, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	havingClause := " HAVING true "
	if len(havingClauses) > 0 {
		havingClause = havingClause + " AND " + strings.Join(havingClauses, " AND ") + " "
	}

	fmt.Printf("havingClause: %v\n", havingClause)

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	FROM public.work_items %s
	WHERE work_items.work_item_id %s $1
	%s   AND work_items.deleted_at is %s
	%s
	%s
  ORDER BY
		work_item_id %s`, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* WorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItems retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index '[xo] base filter query'.
func WorkItems(ctx context.Context, db DB, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 0
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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	 FROM public.work_items %s
	 WHERE true
	 %s   AND work_items.deleted_at is %s  %s
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItems */\n" + sqlstr

	// run
	// logf(sqlstr, )
	rows, err := db.Query(ctx, sqlstr, append([]any{}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItemsByDeletedAt_WhereDeletedAtIsNotNull retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_deleted_at_idx'.
func WorkItemsByDeletedAt_WhereDeletedAtIsNotNull(ctx context.Context, db DB, deletedAt *time.Time, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " not null ", joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	 FROM public.work_items %s
	 WHERE work_items.deleted_at = $1 AND (deleted_at IS NOT NULL)
	 %s   AND work_items.deleted_at is %s  %s
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemsByDeletedAt_WhereDeletedAtIsNotNull */\n" + sqlstr

	// run
	// logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{deletedAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDeletedAt/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByDeletedAt/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItemByWorkItemID retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_pkey'.
func WorkItemByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...WorkItemSelectConfigOption) (*WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	 FROM public.work_items %s
	 WHERE work_items.work_item_id = $1
	 %s   AND work_items.deleted_at is %s  %s
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	wi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err}))
	}

	return &wi, nil
}

// WorkItemsByTeamID retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_team_id_idx'.
func WorkItemsByTeamID(ctx context.Context, db DB, teamID TeamID, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	 FROM public.work_items %s
	 WHERE work_items.team_id = $1
	 %s   AND work_items.deleted_at is %s  %s
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemsByTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{teamID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTeamID/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTeamID/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
}

// WorkItemsByTitle retrieves a row from 'public.work_items' as a WorkItem.
//
// Generated from index 'work_items_title_description_idx1'.
func WorkItemsByTitle(ctx context.Context, db DB, title string, opts ...WorkItemSelectConfigOption) ([]WorkItem, error) {
	c := &WorkItemSelectConfig{deletedAt: " null ", joins: WorkItemJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.DemoTwoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoTwoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoTwoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoTwoWorkItemGroupBySQL)
	}

	if c.joins.DemoWorkItem {
		selectClauses = append(selectClauses, workItemTableDemoWorkItemSelectSQL)
		joinClauses = append(joinClauses, workItemTableDemoWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableDemoWorkItemGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, workItemTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, workItemTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTimeEntriesGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, workItemTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, workItemTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableAssignedUsersGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, workItemTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemCommentsGroupBySQL)
	}

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTagsGroupBySQL)
	}

	if c.joins.KanbanStep {
		selectClauses = append(selectClauses, workItemTableKanbanStepSelectSQL)
		joinClauses = append(joinClauses, workItemTableKanbanStepJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableKanbanStepGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, workItemTableTeamSelectSQL)
		joinClauses = append(joinClauses, workItemTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableTeamGroupBySQL)
	}

	if c.joins.WorkItemType {
		selectClauses = append(selectClauses, workItemTableWorkItemTypeSelectSQL)
		joinClauses = append(joinClauses, workItemTableWorkItemTypeJoinSQL)
		groupByClauses = append(groupByClauses, workItemTableWorkItemTypeGroupBySQL)
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
	work_items.closed_at,
	work_items.created_at,
	work_items.deleted_at,
	work_items.description,
	work_items.kanban_step_id,
	work_items.metadata,
	work_items.target_date,
	work_items.team_id,
	work_items.title,
	work_items.updated_at,
	work_items.work_item_id,
	work_items.work_item_type_id %s
	 FROM public.work_items %s
	 WHERE work_items.title = $1
	 %s   AND work_items.deleted_at is %s  %s
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemsByTitle */\n" + sqlstr

	// run
	// logf(sqlstr, title)
	rows, err := db.Query(ctx, sqlstr, append([]any{title}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTitleDescription/Query: %w", &XoError{Entity: "Work item", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItem/WorkItemsByTitleDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err}))
	}
	return res, nil
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
