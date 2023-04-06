package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// WorkItemTag represents a row from 'public.work_item_tags'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type WorkItemTag struct {
	WorkItemTagID int    `json:"workItemTagID" db:"work_item_tag_id"` // work_item_tag_id
	ProjectID     int    `json:"projectID" db:"project_id"`           // project_id
	Name          string `json:"name" db:"name"`                      // name
	Description   string `json:"description" db:"description"`        // description
	Color         string `json:"color" db:"color"`                    // color

	WorkItems *[]WorkItem `json:"workItems" db:"work_items"` // M2M
	// xo fields
	_exists, _deleted bool
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemTagOrderBy = string

type WorkItemTagJoins struct {
	WorkItems bool
}

// WithWorkItemTagJoin orders results by the given columns.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemTag exists in the database.
func (wit *WorkItemTag) Exists() bool {
	return wit._exists
}

// Deleted returns true when the WorkItemTag has been marked for deletion from
// the database.
func (wit *WorkItemTag) Deleted() bool {
	return wit._deleted
}

// Insert inserts the WorkItemTag to the database.
func (wit *WorkItemTag) Insert(ctx context.Context, db DB) error {
	switch {
	case wit._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wit._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING work_item_tag_id `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err := db.QueryRow(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color).Scan(&wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Update updates a WorkItemTag in the database.
func (wit *WorkItemTag) Update(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wit._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_tags SET ` +
		`project_id = $1, name = $2, description = $3, color = $4 ` +
		`WHERE work_item_tag_id = $5 ` +
		`RETURNING work_item_tag_id `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err := db.QueryRow(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID).Scan(&wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the WorkItemTag to the database.
func (wit *WorkItemTag) Save(ctx context.Context, db DB) error {
	if wit.Exists() {
		return wit.Update(ctx, db)
	}
	return wit.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemTag.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB) error {
	switch {
	case wit._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`work_item_tag_id, project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_tag_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color  `
	// run
	logf(sqlstr, wit.WorkItemTagID, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID, wit.ProjectID, wit.Name, wit.Description, wit.Color); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Delete deletes the WorkItemTag from the database.
func (wit *WorkItemTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return nil
	case wit._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_tags ` +
		`WHERE work_item_tag_id = $1 `
	// run
	logf(sqlstr, wit.WorkItemTagID)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	// set deleted
	wit._deleted = true
	return nil
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
(case when $1::boolean = true then joined_work_items.work_items end)::jsonb as work_items ` +
		`FROM public.work_item_tags ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
		work_item_tag_id as work_items_work_item_tag_id
		, array_agg(work_items.*) as work_items
	from
		work_item_work_item_tag
		join work_items using (work_item_id)
	where
		work_item_tag_id in (
			select
				work_item_tag_id
			from
				work_item_work_item_tag
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				work_item_tag_id) joined_work_items on joined_work_items.work_items_work_item_tag_id = work_item_tags.work_item_tag_id` +
		` WHERE work_item_tags.name = $2 AND work_item_tags.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name, projectID)
	wit := WorkItemTag{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.WorkItems, name, projectID).Scan(&wit.WorkItemTagID, &wit.ProjectID, &wit.Name, &wit.Description, &wit.Color, &wit.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
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
(case when $1::boolean = true then joined_work_items.work_items end)::jsonb as work_items ` +
		`FROM public.work_item_tags ` +
		`-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
		work_item_tag_id as work_items_work_item_tag_id
		, array_agg(work_items.*) as work_items
	from
		work_item_work_item_tag
		join work_items using (work_item_id)
	where
		work_item_tag_id in (
			select
				work_item_tag_id
			from
				work_item_work_item_tag
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				work_item_tag_id) joined_work_items on joined_work_items.work_items_work_item_tag_id = work_item_tags.work_item_tag_id` +
		` WHERE work_item_tags.work_item_tag_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemTagID)
	wit := WorkItemTag{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.WorkItems, workItemTagID).Scan(&wit.WorkItemTagID, &wit.ProjectID, &wit.Name, &wit.Description, &wit.Color, &wit.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemTag's (ProjectID).
//
// Generated from foreign key 'work_item_tags_project_id_fkey'.
func (wit *WorkItemTag) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
