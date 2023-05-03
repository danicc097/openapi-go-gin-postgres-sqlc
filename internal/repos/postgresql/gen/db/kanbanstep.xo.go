package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// KanbanStep represents a row from 'public.kanban_steps'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type KanbanStep struct {
	KanbanStepID  int    `json:"kanbanStepID" db:"kanban_step_id" required:"true"`  // kanban_step_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`         // project_id
	StepOrder     *int   `json:"stepOrder" db:"step_order" required:"true"`         // step_order
	Name          string `json:"name" db:"name" required:"true"`                    // name
	Description   string `json:"description" db:"description" required:"true"`      // description
	Color         string `json:"color" db:"color" required:"true"`                  // color
	TimeTrackable bool   `json:"timeTrackable" db:"time_trackable" required:"true"` // time_trackable

	ProjectJoin  *Project  `json:"-" db:"project" openapi-go:"ignore"`   // O2O (generated from M2O)
	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O (inferred)

}

// KanbanStepCreateParams represents insert params for 'public.kanban_steps'
type KanbanStepCreateParams struct {
	ProjectID     int    `json:"projectID"`     // project_id
	StepOrder     *int   `json:"stepOrder"`     // step_order
	Name          string `json:"name"`          // name
	Description   string `json:"description"`   // description
	Color         string `json:"color"`         // color
	TimeTrackable bool   `json:"timeTrackable"` // time_trackable
}

// CreateKanbanStep creates a new KanbanStep in the database with the given params.
func CreateKanbanStep(ctx context.Context, db DB, params *KanbanStepCreateParams) (*KanbanStep, error) {
	ks := &KanbanStep{
		ProjectID:     params.ProjectID,
		StepOrder:     params.StepOrder,
		Name:          params.Name,
		Description:   params.Description,
		Color:         params.Color,
		TimeTrackable: params.TimeTrackable,
	}

	return ks.Insert(ctx, db)
}

// KanbanStepUpdateParams represents update params for 'public.kanban_steps'
type KanbanStepUpdateParams struct {
	ProjectID     *int    `json:"projectID"`     // project_id
	StepOrder     **int   `json:"stepOrder"`     // step_order
	Name          *string `json:"name"`          // name
	Description   *string `json:"description"`   // description
	Color         *string `json:"color"`         // color
	TimeTrackable *bool   `json:"timeTrackable"` // time_trackable
}

// SetUpdateParams updates public.kanban_steps struct fields with the specified params.
func (ks *KanbanStep) SetUpdateParams(params *KanbanStepUpdateParams) {
	if params.ProjectID != nil {
		ks.ProjectID = *params.ProjectID
	}
	if params.StepOrder != nil {
		ks.StepOrder = *params.StepOrder
	}
	if params.Name != nil {
		ks.Name = *params.Name
	}
	if params.Description != nil {
		ks.Description = *params.Description
	}
	if params.Color != nil {
		ks.Color = *params.Color
	}
	if params.TimeTrackable != nil {
		ks.TimeTrackable = *params.TimeTrackable
	}
}

type KanbanStepSelectConfig struct {
	limit   string
	orderBy string
	joins   KanbanStepJoins
}
type KanbanStepSelectConfigOption func(*KanbanStepSelectConfig)

// WithKanbanStepLimit limits row selection.
func WithKanbanStepLimit(limit int) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type KanbanStepOrderBy = string

const ()

type KanbanStepJoins struct {
	Project  bool
	WorkItem bool
}

// WithKanbanStepJoin joins with the given tables.
func WithKanbanStepJoin(joins KanbanStepJoins) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the KanbanStep to the database.
func (ks *KanbanStep) Insert(ctx context.Context, db DB) (*KanbanStep, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.kanban_steps (` +
		`project_id, step_order, name, description, color, time_trackable` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING * `
	// run
	logf(sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable)

	rows, err := db.Query(ctx, sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Insert/db.Query: %w", err))
	}
	newks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Insert/pgx.CollectOneRow: %w", err))
	}

	*ks = newks

	return ks, nil
}

// Update updates a KanbanStep in the database.
func (ks *KanbanStep) Update(ctx context.Context, db DB) (*KanbanStep, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.kanban_steps SET ` +
		`project_id = $1, step_order = $2, name = $3, description = $4, color = $5, time_trackable = $6 ` +
		`WHERE kanban_step_id = $7 ` +
		`RETURNING * `
	// run
	logf(sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable, ks.KanbanStepID)

	rows, err := db.Query(ctx, sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable, ks.KanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Update/db.Query: %w", err))
	}
	newks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Update/pgx.CollectOneRow: %w", err))
	}
	*ks = newks

	return ks, nil
}

// Upsert performs an upsert for KanbanStep.
func (ks *KanbanStep) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.kanban_steps (` +
		`kanban_step_id, project_id, step_order, name, description, color, time_trackable` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (kanban_step_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, step_order = EXCLUDED.step_order, name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color, time_trackable = EXCLUDED.time_trackable ` +
		` RETURNING * `
	// run
	logf(sqlstr, ks.KanbanStepID, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable)
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the KanbanStep from the database.
func (ks *KanbanStep) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.kanban_steps ` +
		`WHERE kanban_step_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID); err != nil {
		return logerror(err)
	}
	return nil
}

// KanbanStepByKanbanStepID retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_pkey'.
func KanbanStepByKanbanStepID(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.kanban_step_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, kanbanStepID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, kanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/pgx.CollectOneRow: %w", err))
	}

	return &ks, nil
}

// KanbanStepsByProjectID_WhereStepOrderIsNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_idx'.
func KanbanStepsByProjectID_WhereStepOrderIsNull(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepByProjectIDName_WhereStepOrderIsNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_idx'.
func KanbanStepByProjectIDName_WhereStepOrderIsNull(ctx context.Context, db DB, projectID int, name string, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 AND kanban_steps.name = $4 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDName/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDName/pgx.CollectOneRow: %w", err))
	}

	return &ks, nil
}

// KanbanStepsByName_WhereStepOrderIsNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_idx'.
func KanbanStepsByName_WhereStepOrderIsNull(ctx context.Context, db DB, name string, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.name = $3 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepByProjectIDNameStepOrder_WhereStepOrderIsNotNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepByProjectIDNameStepOrder_WhereStepOrderIsNotNull(ctx context.Context, db DB, projectID int, name string, stepOrder *int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 AND kanban_steps.name = $4 AND kanban_steps.step_order = $5 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, name, stepOrder)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID, name, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/pgx.CollectOneRow: %w", err))
	}

	return &ks, nil
}

// KanbanStepsByProjectID_WhereStepOrderIsNotNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByProjectID_WhereStepOrderIsNotNull(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepsByName_WhereStepOrderIsNotNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByName_WhereStepOrderIsNotNull(ctx context.Context, db DB, name string, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.name = $3 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepsByStepOrder_WhereStepOrderIsNotNull retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByStepOrder_WhereStepOrderIsNotNull(ctx context.Context, db DB, stepOrder *int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.step_order = $3 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, stepOrder)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepByProjectIDStepOrder retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_step_order_key'.
func KanbanStepByProjectIDStepOrder(ctx context.Context, db DB, projectID int, stepOrder *int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 AND kanban_steps.step_order = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, stepOrder)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/pgx.CollectOneRow: %w", err))
	}

	return &ks, nil
}

// KanbanStepsByProjectID retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_step_order_key'.
func KanbanStepsByProjectID(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepsByStepOrder retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_step_order_key'.
func KanbanStepsByStepOrder(ctx context.Context, db DB, stepOrder *int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true and work_items.kanban_step_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.kanban_steps ` +
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey(O2O inferred)"
left join work_items on work_items.kanban_step_id = kanban_steps.kanban_step_id` +
		` WHERE kanban_steps.step_order = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, stepOrder)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}
