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

	// xo fields
	_exists, _deleted bool
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

// KanbanStepUpdateParams represents update params for 'public.kanban_steps'
type KanbanStepUpdateParams struct {
	ProjectID     *int    `json:"projectID"`     // project_id
	StepOrder     **int   `json:"stepOrder"`     // step_order
	Name          *string `json:"name"`          // name
	Description   *string `json:"description"`   // description
	Color         *string `json:"color"`         // color
	TimeTrackable *bool   `json:"timeTrackable"` // time_trackable
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
}

// WithKanbanStepJoin joins with the given tables.
func WithKanbanStepJoin(joins KanbanStepJoins) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the KanbanStep to the database.
func (ks *KanbanStep) Insert(ctx context.Context, db DB) (*KanbanStep, error) {
	switch {
	case ks._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ks._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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

	newks._exists = true
	*ks = newks

	return ks, nil
}

// Update updates a KanbanStep in the database.
func (ks *KanbanStep) Update(ctx context.Context, db DB) (*KanbanStep, error) {
	switch {
	case !ks._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ks._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
	newks._exists = true
	*ks = newks

	return ks, nil
}

// Save saves the KanbanStep to the database.
func (ks *KanbanStep) Save(ctx context.Context, db DB) (*KanbanStep, error) {
	if ks._exists {
		return ks.Update(ctx, db)
	}
	return ks.Insert(ctx, db)
}

// Upsert performs an upsert for KanbanStep.
func (ks *KanbanStep) Upsert(ctx context.Context, db DB) error {
	switch {
	case ks._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
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
	ks._exists = true
	return nil
}

// Delete deletes the KanbanStep from the database.
func (ks *KanbanStep) Delete(ctx context.Context, db DB) error {
	switch {
	case !ks._exists: // doesn't exist
		return nil
	case ks._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.kanban_steps ` +
		`WHERE kanban_step_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID); err != nil {
		return logerror(err)
	}
	// set deleted
	ks._deleted = true
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.kanban_step_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, kanbanStepID)
	rows, err := db.Query(ctx, sqlstr, kanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/pgx.CollectOneRow: %w", err))
	}
	ks._exists = true

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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, projectID)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 AND kanban_steps.name = $3 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, name)
	rows, err := db.Query(ctx, sqlstr, projectID, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDName/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDName/pgx.CollectOneRow: %w", err))
	}
	ks._exists = true

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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.name = $2 AND (step_order IS NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, name)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 AND kanban_steps.name = $3 AND kanban_steps.step_order = $4 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, name, stepOrder)
	rows, err := db.Query(ctx, sqlstr, projectID, name, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/pgx.CollectOneRow: %w", err))
	}
	ks._exists = true

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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, projectID)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.name = $2 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, name)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.step_order = $2 AND (step_order IS NOT NULL) `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, stepOrder)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 AND kanban_steps.step_order = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID, stepOrder)
	rows, err := db.Query(ctx, sqlstr, projectID, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/db.Query: %w", err))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/pgx.CollectOneRow: %w", err))
	}
	ks._exists = true

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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.project_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, projectID)
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project ` +
		`FROM public.kanban_steps ` +
		`-- automatic join generated from foreign key on "project_id"
left join projects on projects.project_id = kanban_steps.project_id` +
		` WHERE kanban_steps.step_order = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, stepOrder)
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

// FKProject_ProjectID returns the Project associated with the KanbanStep's (ProjectID).
//
// Generated from foreign key 'kanban_steps_project_id_fkey'.
func (ks *KanbanStep) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, ks.ProjectID)
}
