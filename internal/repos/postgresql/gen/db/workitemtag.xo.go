package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// WorkItemTag represents a row from 'public.work_item_tags'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type WorkItemTag struct {
	WorkItemTagID int    `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`           // project_id
	Name          string `json:"name" db:"name" required:"true"`                      // name
	Description   string `json:"description" db:"description" required:"true"`        // description
	Color         string `json:"color" db:"color" required:"true"`                    // color

	ProjectJoin              *Project    `json:"-" db:"project_project_id" openapi-go:"ignore"`                 // O2O (generated from M2O)
	WorkItemsJoinWorkItemTag *[]WorkItem `json:"-" db:"work_item_work_item_tag_work_items" openapi-go:"ignore"` // M2M

}

// WorkItemTagCreateParams represents insert params for 'public.work_item_tags'.
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
	Project              bool
	WorkItemsWorkItemTag bool
}

// WithWorkItemTagJoin joins with the given tables.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = WorkItemTagJoins{
			Project:              s.joins.Project || joins.Project,
			WorkItemsWorkItemTag: s.joins.WorkItemsWorkItemTag || joins.WorkItemsWorkItemTag,
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

// Upsert upserts a WorkItemTag in the database.
// Requires appropiate PK(s) to be set beforehand.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
	var err error

	wit.ProjectID = params.ProjectID
	wit.Name = params.Name
	wit.Description = params.Description
	wit.Color = params.Color

	wit, err = wit.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			wit, err = wit.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return wit, err
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

// WorkItemTagPaginatedByWorkItemTagID returns a cursor-paginated list of WorkItemTag.
func WorkItemTagPaginatedByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.work_item_tag_id > $3 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagPaginatedByProjectID returns a cursor-paginated list of WorkItemTag.
func WorkItemTagPaginatedByProjectID(ctx context.Context, db DB, projectID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.project_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color,
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.project_id > $3 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, projectID)
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
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.name = $3 AND work_item_tags.project_id = $4 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItemsWorkItemTag, name, projectID)
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
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.name = $3 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItemsWorkItemTag, name)
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
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.project_id = $3 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItemsWorkItemTag, projectID)
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
(case when $1::boolean = true and _projects_project_ids.project_id is not null then row(_projects_project_ids.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items is not null), '{}') end) as work_item_work_item_tag_work_items ` +
		`FROM public.work_item_tags ` +
		`-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _projects_project_ids on _projects_project_ids.project_id = work_item_tags.project_id
-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
			work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
			, row(work_items.*) as __work_items
		from
			work_item_work_item_tag
    join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
    group by
			work_item_work_item_tag_work_item_tag_id
			, work_items.work_item_id
  ) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
` +
		` WHERE work_item_tags.work_item_tag_id = $3 GROUP BY _projects_project_ids.project_id,
      _projects_project_ids.project_id,
	work_item_tags.work_item_tag_id, 
work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItemsWorkItemTag, workItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/pgx.CollectOneRow: %w", err))
	}

	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemTag's (ProjectID).
//
// Generated from foreign key 'work_item_tags_project_id_fkey'.
func (wit *WorkItemTag) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
