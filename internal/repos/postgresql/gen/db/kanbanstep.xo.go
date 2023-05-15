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

// KanbanStep represents a row from 'public.kanban_steps'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type KanbanStep struct {
	KanbanStepID  int    `json:"kanbanStepID" db:"kanban_step_id" required:"true"`  // kanban_step_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`         // project_id
	StepOrder     int    `json:"stepOrder" db:"step_order" required:"true"`         // step_order
	Name          string `json:"name" db:"name" required:"true"`                    // name
	Description   string `json:"description" db:"description" required:"true"`      // description
	Color         string `json:"color" db:"color" required:"true"`                  // color
	TimeTrackable bool   `json:"timeTrackable" db:"time_trackable" required:"true"` // time_trackable

	ProjectJoin  *Project  `json:"-" db:"project_project_id" openapi-go:"ignore"`       // O2O projects (generated from M2O)
	WorkItemJoin *WorkItem `json:"-" db:"work_item_kanban_step_id" openapi-go:"ignore"` // O2O work_items (inferred)

}

// KanbanStepCreateParams represents insert params for 'public.kanban_steps'.
type KanbanStepCreateParams struct {
	ProjectID     int    `json:"projectID" required:"true"`     // project_id
	StepOrder     int    `json:"stepOrder" required:"true"`     // step_order
	Name          string `json:"name" required:"true"`          // name
	Description   string `json:"description" required:"true"`   // description
	Color         string `json:"color" required:"true"`         // color
	TimeTrackable bool   `json:"timeTrackable" required:"true"` // time_trackable
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
	ProjectID     *int    `json:"projectID" required:"true"`     // project_id
	StepOrder     *int    `json:"stepOrder" required:"true"`     // step_order
	Name          *string `json:"name" required:"true"`          // name
	Description   *string `json:"description" required:"true"`   // description
	Color         *string `json:"color" required:"true"`         // color
	TimeTrackable *bool   `json:"timeTrackable" required:"true"` // time_trackable
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
	filters map[string][]any
}
type KanbanStepSelectConfigOption func(*KanbanStepSelectConfig)

// WithKanbanStepLimit limits row selection.
func WithKanbanStepLimit(limit int) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type KanbanStepOrderBy = string

const ()

type KanbanStepJoins struct {
	Project  bool // O2O projects
	WorkItem bool // O2O work_items
}

// WithKanbanStepJoin joins with the given tables.
func WithKanbanStepJoin(joins KanbanStepJoins) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.joins = KanbanStepJoins{
			Project:  s.joins.Project || joins.Project,
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithKanbanStepFilters adds the given filters, which may be parameterized.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`col.created_at > $i AND
//		col.created_at < $i`: {time.Now().Add(-24 * time.Hour), time.Now().Add(24 * time.Hour)},
//	}
func WithKanbanStepFilters(filters map[string][]any) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.filters = filters
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

// Upsert upserts a KanbanStep in the database.
// Requires appropiate PK(s) to be set beforehand.
func (ks *KanbanStep) Upsert(ctx context.Context, db DB, params *KanbanStepCreateParams) (*KanbanStep, error) {
	var err error

	ks.ProjectID = params.ProjectID
	ks.StepOrder = params.StepOrder
	ks.Name = params.Name
	ks.Description = params.Description
	ks.Color = params.Color
	ks.TimeTrackable = params.TimeTrackable

	ks, err = ks.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			ks, err = ks.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return ks, err
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

// KanbanStepPaginatedByKanbanStepIDAsc returns a cursor-paginated list of KanbanStep in Asc order.
func KanbanStepPaginatedByKanbanStepIDAsc(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.kanban_step_id > $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		kanban_step_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, kanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepPaginatedByProjectIDAsc returns a cursor-paginated list of KanbanStep in Asc order.
func KanbanStepPaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.project_id > $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		project_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepPaginatedByStepOrderAsc returns a cursor-paginated list of KanbanStep in Asc order.
func KanbanStepPaginatedByStepOrderAsc(ctx context.Context, db DB, stepOrder int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.step_order > $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		step_order Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepPaginatedByKanbanStepIDDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByKanbanStepIDDesc(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.kanban_step_id < $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		kanban_step_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, kanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepPaginatedByProjectIDDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.project_id < $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		project_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepPaginatedByStepOrderDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByStepOrderDesc(ctx context.Context, db DB, stepOrder int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.step_order < $3`+
		` %s  GROUP BY kanban_steps.kanban_step_id, 
kanban_steps.project_id, 
kanban_steps.step_order, 
kanban_steps.name, 
kanban_steps.description, 
kanban_steps.color, 
kanban_steps.time_trackable, 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id ORDER BY 
		step_order Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepByKanbanStepID retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_pkey'.
func KanbanStepByKanbanStepID(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.kanban_step_id = $3`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
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

// KanbanStepByProjectIDNameStepOrder retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepByProjectIDNameStepOrder(ctx context.Context, db DB, projectID int, name string, stepOrder int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.project_id = $3 AND kanban_steps.name = $4 AND kanban_steps.step_order = $5`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
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

// KanbanStepsByProjectID retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByProjectID(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.project_id = $3`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepsByName retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByName(ctx context.Context, db DB, name string, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.name = $3`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepsByStepOrder retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_name_step_order_idx'.
func KanbanStepsByStepOrder(ctx context.Context, db DB, stepOrder int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.step_order = $3`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.WorkItem, stepOrder)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// KanbanStepByProjectIDStepOrder retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_project_id_step_order_key'.
func KanbanStepByProjectIDStepOrder(ctx context.Context, db DB, projectID int, stepOrder int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`kanban_steps.kanban_step_id,
kanban_steps.project_id,
kanban_steps.step_order,
kanban_steps.name,
kanban_steps.description,
kanban_steps.color,
kanban_steps.time_trackable,
(case when $1::boolean = true and _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id,
(case when $2::boolean = true and _kanban_steps_kanban_step_id.kanban_step_id is not null then row(_kanban_steps_kanban_step_id.*) end) as work_item_kanban_step_id `+
		`FROM public.kanban_steps `+
		`-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join work_items as _kanban_steps_kanban_step_id on _kanban_steps_kanban_step_id.kanban_step_id = kanban_steps.kanban_step_id`+
		` WHERE kanban_steps.project_id = $3 AND kanban_steps.step_order = $4`+
		` %s  GROUP BY 
_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id, 
_kanban_steps_kanban_step_id.kanban_step_id,
      _kanban_steps_kanban_step_id.work_item_id,
	kanban_steps.kanban_step_id `, filters)
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

// FKProject_ProjectID returns the Project associated with the KanbanStep's (ProjectID).
//
// Generated from foreign key 'kanban_steps_project_id_fkey'.
func (ks *KanbanStep) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, ks.ProjectID)
}
