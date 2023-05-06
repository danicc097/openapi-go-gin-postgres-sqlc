package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// WorkItemTag represents a row from 'public.work_item_tags'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type WorkItemTag struct {
	WorkItemTagID int    `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`           // project_id
	Name          string `json:"name" db:"name" required:"true"`                      // name
	Description   string `json:"description" db:"description" required:"true"`        // description
	Color         string `json:"color" db:"color" required:"true"`                    // color

	ProjectJoin      *Project       `json:"-" db:"project" openapi-go:"ignore"`        // O2O (generated from M2O)
	WorkItemsJoin    *[]WorkItem    `json:"-" db:"work_items" openapi-go:"ignore"`     // M2M
	WorkItemTagJoin  *WorkItemTag   `json:"-" db:"work_item_tag" openapi-go:"ignore"`  // O2O (generated from M2O)
	WorkItemTagsJoin *[]WorkItemTag `json:"-" db:"work_item_tags" openapi-go:"ignore"` // M2O

}

// WorkItemTagCreateParams represents insert params for 'public.work_item_tags'
type WorkItemTagCreateParams struct {
	ProjectID   int    `json:"projectID" required:"true"`   // project_id
	Name        string `json:"name" required:"true"`        // name
	Description string `json:"description" required:"true"` // description
	Color       string `json:"color" required:"true"`       // color
}

// CreateWorkItemTag creates a new WorkItemTag in the database with the given params.
func CreateWorkItemTag(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
	wit := &WorkItemTag{
		ProjectID:   params.ProjectID,
		Name:        params.Name,
		Description: params.Description,
		Color:       params.Color,
	}

	return wit.Insert(ctx, db)
}

// WorkItemTagUpdateParams represents update params for 'public.work_item_tags'
type WorkItemTagUpdateParams struct {
	ProjectID   *int    `json:"projectID" required:"true"`   // project_id
	Name        *string `json:"name" required:"true"`        // name
	Description *string `json:"description" required:"true"` // description
	Color       *string `json:"color" required:"true"`       // color
}

// SetUpdateParams updates public.work_item_tags struct fields with the specified params.
func (wit *WorkItemTag) SetUpdateParams(params *WorkItemTagUpdateParams) {
	if params.ProjectID != nil {
		wit.ProjectID = *params.ProjectID
	}
	if params.Name != nil {
		wit.Name = *params.Name
	}
	if params.Description != nil {
		wit.Description = *params.Description
	}
	if params.Color != nil {
		wit.Color = *params.Color
	}
}

type WorkItemTagSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTagJoins
}
type WorkItemTagSelectConfigOption func(*WorkItemTagSelectConfig)

// WithWorkItemTagLimit limits row selection.
func WithWorkItemTagLimit(limit int) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemTagOrderBy = string

const ()

type WorkItemTagJoins struct {
	Project      bool
	WorkItems    bool
	WorkItemTag  bool
	WorkItemTags bool
}

// WithWorkItemTagJoin joins with the given tables.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = WorkItemTagJoins{

			Project:      s.joins.Project || joins.Project,
			WorkItems:    s.joins.WorkItems || joins.WorkItems,
			WorkItemTag:  s.joins.WorkItemTag || joins.WorkItemTag,
			WorkItemTags: s.joins.WorkItemTags || joins.WorkItemTags,
		}
	}
}

// Insert inserts the WorkItemTag to the database.
func (wit *WorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemTag, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/pgx.CollectOneRow: %w", err))
	}

	*wit = newwit

	return wit, nil
}

// Update updates a WorkItemTag in the database.
func (wit *WorkItemTag) Update(ctx context.Context, db DB) (*WorkItemTag, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_tags SET ` +
		`project_id = $1, name = $2, description = $3, color = $4 ` +
		`WHERE work_item_tag_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/pgx.CollectOneRow: %w", err))
	}
	*wit = newwit

	return wit, nil
}

// Upsert performs an upsert for WorkItemTag.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`work_item_tag_id, project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_tag_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color ` +
		` RETURNING * `
	// run
	logf(sqlstr, wit.WorkItemTagID, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID, wit.ProjectID, wit.Name, wit.Description, wit.Color); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the WorkItemTag from the database.
func (wit *WorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_tags ` +
		`WHERE work_item_tag_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	return nil
}

// PaginatedWorkItemTagByWorkItemTagID returns a cursor-paginated list of WorkItemTag.
func (wit *WorkItemTag) PaginatedWorkItemTagByWorkItemTagID(ctx context.Context, db DB) ([]WorkItemTag, error) {
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.work_item_tag_id > $5 `
	// run

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// PaginatedWorkItemTagByProjectID returns a cursor-paginated list of WorkItemTag.
func (wit *WorkItemTag) PaginatedWorkItemTagByProjectID(ctx context.Context, db DB) ([]WorkItemTag, error) {
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.project_id > $5 `
	// run

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// PaginatedWorkItemTagByProjectID returns a cursor-paginated list of WorkItemTag.
func (wit *WorkItemTag) PaginatedWorkItemTagByProjectID(ctx context.Context, db DB) ([]WorkItemTag, error) {
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.project_id > $5 `
	// run

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagByNameProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.name = $5 AND work_item_tags.project_id = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItems, c.joins.WorkItemTag, c.joins.WorkItemTags, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/pgx.CollectOneRow: %w", err))
	}

	return &wit, nil
}

// WorkItemTagsByName retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByName(ctx context.Context, db DB, name string, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.name = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItems, c.joins.WorkItemTag, c.joins.WorkItemTags, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagsByProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByProjectID(ctx context.Context, db DB, projectID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.project_id = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItems, c.joins.WorkItemTag, c.joins.WorkItemTags, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagByWorkItemTagID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_pkey'.
func WorkItemTagByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $3::boolean = true and work_item_tags.name is not null then row(work_item_tags.*) end) as work_item_tag,
(case when $4::boolean = true then COALESCE(joined_work_item_tags.work_item_tags, '{}') end) as work_item_tags ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_work_item_tag
    	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by work_item_work_item_tag_work_item_tag_id
  ) as joined_work_items on joined_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id

-- O2O join generated from "work_item_tags_name_project_id_key (Generated from M2O)"
left join work_item_tags on work_item_tags.name = work_item_tags.project_id
-- M2O join generated from "work_item_tags_name_project_id_key"
left join (
  select
  name as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
  group by
        name) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = work_item_tags.project_id` +
		` WHERE work_item_tags.work_item_tag_id = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItems, c.joins.WorkItemTag, c.joins.WorkItemTags, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/pgx.CollectOneRow: %w", err))
	}

	return &wit, nil
}
