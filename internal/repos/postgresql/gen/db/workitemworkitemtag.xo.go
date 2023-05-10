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
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type WorkItemWorkItemTag struct {
	WorkItemTagID int   `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID" db:"work_item_id" required:"true"`        // work_item_id

	WorkItemTagsJoin *[]WorkItemTag `json:"-" db:"work_item_tags" openapi-go:"ignore"` // M2M
	WorkItemsJoin    *[]WorkItem    `json:"-" db:"work_items" openapi-go:"ignore"`     // M2M

}

// WorkItemWorkItemTagCreateParams represents insert params for 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagCreateParams struct {
	WorkItemTagID int   `json:"workItemTagID" required:"true"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID" required:"true"`    // work_item_id
}

// CreateWorkItemWorkItemTag creates a new WorkItemWorkItemTag in the database with the given params.
func CreateWorkItemWorkItemTag(ctx context.Context, db DB, params *WorkItemWorkItemTagCreateParams) (*WorkItemWorkItemTag, error) {
	wiwit := &WorkItemWorkItemTag{
		WorkItemTagID: params.WorkItemTagID,
		WorkItemID:    params.WorkItemID,
	}

	return wiwit.Insert(ctx, db)
}

// WorkItemWorkItemTagUpdateParams represents update params for 'public.work_item_work_item_tag'
type WorkItemWorkItemTagUpdateParams struct {
	WorkItemTagID *int   `json:"workItemTagID" required:"true"` // work_item_tag_id
	WorkItemID    *int64 `json:"workItemID" required:"true"`    // work_item_id
}

// SetUpdateParams updates public.work_item_work_item_tag struct fields with the specified params.
func (wiwit *WorkItemWorkItemTag) SetUpdateParams(params *WorkItemWorkItemTagUpdateParams) {
	if params.WorkItemTagID != nil {
		wiwit.WorkItemTagID = *params.WorkItemTagID
	}
	if params.WorkItemID != nil {
		wiwit.WorkItemID = *params.WorkItemID
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemWorkItemTagOrderBy = string

const ()

type WorkItemWorkItemTagJoins struct {
	WorkItemTags bool
	WorkItems    bool
}

// WithWorkItemWorkItemTagJoin joins with the given tables.
func WithWorkItemWorkItemTagJoin(joins WorkItemWorkItemTagJoins) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.joins = WorkItemWorkItemTagJoins{
			WorkItemTags: s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItems:    s.joins.WorkItems || joins.WorkItems,
		}
	}
}

// Insert inserts the WorkItemWorkItemTag to the database.
func (wiwit *WorkItemWorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemWorkItemTag, error) {
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
	*wiwit = newwiwit

	return wiwit, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the WorkItemWorkItemTag from the database.
func (wiwit *WorkItemWorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_work_item_tag ` +
		`WHERE work_item_tag_id = $1 AND work_item_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemID returns a cursor-paginated list of WorkItemWorkItemTag.
func WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id,
(case when $1::boolean = true then ARRAY_AGG((
		joined_work_item_tags.__work_item_tags
		)) end) as work_item_tags,
(case when $2::boolean = true then ARRAY_AGG((
		joined_work_items.__work_items
		)) end) as work_items ` +
		`FROM public.work_item_work_item_tag ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, row(work_item_tags.*) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
			work_item_work_item_tag_work_item_id
			, work_item_tags.work_item_tag_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id

-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
` +
		` WHERE work_item_work_item_tag.work_item_tag_id > $3 AND work_item_work_item_tag.work_item_id > $4 `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemTagID, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
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
work_item_work_item_tag.work_item_id,
(case when $1::boolean = true then ARRAY_AGG((
		joined_work_item_tags.__work_item_tags
		)) end) as work_item_tags,
(case when $2::boolean = true then ARRAY_AGG((
		joined_work_items.__work_items
		)) end) as work_items ` +
		`FROM public.work_item_work_item_tag ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, row(work_item_tags.*) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
			work_item_work_item_tag_work_item_id
			, work_item_tags.work_item_tag_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id

-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
` +
		` WHERE work_item_work_item_tag.work_item_id = $3 AND work_item_work_item_tag.work_item_tag_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemTags, c.joins.WorkItems, workItemID, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/db.Query: %w", err))
	}
	wiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectOneRow: %w", err))
	}

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
work_item_work_item_tag.work_item_id,
(case when $1::boolean = true then ARRAY_AGG((
		joined_work_item_tags.__work_item_tags
		)) end) as work_item_tags,
(case when $2::boolean = true then ARRAY_AGG((
		joined_work_items.__work_items
		)) end) as work_items ` +
		`FROM public.work_item_work_item_tag ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, row(work_item_tags.*) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
			work_item_work_item_tag_work_item_id
			, work_item_tags.work_item_tag_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id

-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
` +
		` WHERE work_item_work_item_tag.work_item_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemTags, c.joins.WorkItems, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", err))
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
work_item_work_item_tag.work_item_id,
(case when $1::boolean = true then ARRAY_AGG((
		joined_work_item_tags.__work_item_tags
		)) end) as work_item_tags,
(case when $2::boolean = true then ARRAY_AGG((
		joined_work_items.__work_items
		)) end) as work_items ` +
		`FROM public.work_item_work_item_tag ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, row(work_item_tags.*) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
			work_item_work_item_tag_work_item_id
			, work_item_tags.work_item_tag_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id

-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
` +
		` WHERE work_item_work_item_tag.work_item_tag_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemTags, c.joins.WorkItems, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", err))
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
work_item_work_item_tag.work_item_id,
(case when $1::boolean = true then ARRAY_AGG((
		joined_work_item_tags.__work_item_tags
		)) end) as work_item_tags,
(case when $2::boolean = true then ARRAY_AGG((
		joined_work_items.__work_items
		)) end) as work_items ` +
		`FROM public.work_item_work_item_tag ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
			, row(work_item_tags.*) as __work_item_tags
		from work_item_work_item_tag
    	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
			work_item_work_item_tag_work_item_id
			, work_item_tags.work_item_tag_id
  ) as joined_work_item_tags on joined_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id

-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
` +
		` WHERE work_item_work_item_tag.work_item_tag_id = $3 AND work_item_work_item_tag.work_item_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemTags, c.joins.WorkItems, workItemTagID, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/pgx.CollectRows: %w", err))
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
