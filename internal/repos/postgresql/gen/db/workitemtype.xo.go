package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// WorkItemType represents a row from 'public.work_item_types'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type WorkItemType struct {
	WorkItemTypeID int    `json:"workItemTypeID" db:"work_item_type_id" required:"true"` // work_item_type_id
	ProjectID      int    `json:"projectID" db:"project_id" required:"true"`             // project_id
	Name           string `json:"name" db:"name" required:"true"`                        // name
	Description    string `json:"description" db:"description" required:"true"`          // description
	Color          string `json:"color" db:"color" required:"true"`                      // color

	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O (inferred O2O - modify via `cardinality:` column comment)
	// xo fields
	_exists, _deleted bool
}

// WorkItemTypeCreateParams represents insert params for 'public.work_item_types'
type WorkItemTypeCreateParams struct {
	ProjectID   int    `json:"projectID"`   // project_id
	Name        string `json:"name"`        // name
	Description string `json:"description"` // description
	Color       string `json:"color"`       // color
}

// WorkItemTypeUpdateParams represents update params for 'public.work_item_types'
type WorkItemTypeUpdateParams struct {
	ProjectID   *int    `json:"projectID"`   // project_id
	Name        *string `json:"name"`        // name
	Description *string `json:"description"` // description
	Color       *string `json:"color"`       // color
}

type WorkItemTypeSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTypeJoins
}
type WorkItemTypeSelectConfigOption func(*WorkItemTypeSelectConfig)

// WithWorkItemTypeLimit limits row selection.
func WithWorkItemTypeLimit(limit int) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemTypeOrderBy = string

const ()

type WorkItemTypeJoins struct {
	WorkItem bool
}

// WithWorkItemTypeJoin joins with the given tables.
func WithWorkItemTypeJoin(joins WorkItemTypeJoins) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the WorkItemType to the database.
func (wit *WorkItemType) Insert(ctx context.Context, db DB) (*WorkItemType, error) {
	switch {
	case wit._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wit._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_types (` +
		`project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Insert/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Insert/pgx.CollectOneRow: %w", err))
	}

	newwit._exists = true
	*wit = newwit

	return wit, nil
}

// Update updates a WorkItemType in the database.
func (wit *WorkItemType) Update(ctx context.Context, db DB) (*WorkItemType, error) {
	switch {
	case !wit._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wit._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_types SET ` +
		`project_id = $1, name = $2, description = $3, color = $4 ` +
		`WHERE work_item_type_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Update/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Update/pgx.CollectOneRow: %w", err))
	}
	newwit._exists = true
	*wit = newwit

	return wit, nil
}

// Save saves the WorkItemType to the database.
func (wit *WorkItemType) Save(ctx context.Context, db DB) (*WorkItemType, error) {
	if wit._exists {
		return wit.Update(ctx, db)
	}
	return wit.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemType.
func (wit *WorkItemType) Upsert(ctx context.Context, db DB) error {
	switch {
	case wit._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_types (` +
		`work_item_type_id, project_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_type_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color ` +
		` RETURNING * `
	// run
	logf(sqlstr, wit.WorkItemTypeID, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTypeID, wit.ProjectID, wit.Name, wit.Description, wit.Color); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Delete deletes the WorkItemType from the database.
func (wit *WorkItemType) Delete(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return nil
	case wit._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_types ` +
		`WHERE work_item_type_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTypeID); err != nil {
		return logerror(err)
	}
	// set deleted
	wit._deleted = true
	return nil
}

// WorkItemTypeByNameProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypeByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true and work_items.work_item_type_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.name = $2 AND work_item_types.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByNameProjectID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByNameProjectID/pgx.CollectOneRow: %w", err))
	}
	wit._exists = true

	return &wit, nil
}

// WorkItemTypesByName retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypesByName(ctx context.Context, db DB, name string, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true and work_items.work_item_type_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.name = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTypesByProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypesByProjectID(ctx context.Context, db DB, projectID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true and work_items.work_item_type_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.project_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTypeByWorkItemTypeID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_pkey'.
func WorkItemTypeByWorkItemTypeID(ctx context.Context, db DB, workItemTypeID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_types.work_item_type_id,
work_item_types.project_id,
work_item_types.name,
work_item_types.description,
work_item_types.color,
(case when $1::boolean = true and work_items.work_item_type_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.work_item_types ` +
		`-- O2O join generated from "work_items_work_item_type_id_fkey"
left join work_items on work_items.work_item_type_id = work_item_types.work_item_type_id` +
		` WHERE work_item_types.work_item_type_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTypeID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemTypeID)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByWorkItemTypeID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByWorkItemTypeID/pgx.CollectOneRow: %w", err))
	}
	wit._exists = true

	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemType's (ProjectID).
//
// Generated from foreign key 'work_item_types_project_id_fkey'.
func (wit *WorkItemType) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
