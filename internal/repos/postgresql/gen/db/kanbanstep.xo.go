package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// KanbanStep represents a row from 'public.kanban_steps'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type KanbanStep struct {
	KanbanStepID  int    `json:"kanbanStepID" db:"kanban_step_id" required:"true"`                              // kanban_step_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`                                     // project_id
	StepOrder     int    `json:"stepOrder" db:"step_order" required:"true"`                                     // step_order
	Name          string `json:"name" db:"name" required:"true"`                                                // name
	Description   string `json:"description" db:"description" required:"true"`                                  // description
	Color         string `json:"color" db:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	TimeTrackable bool   `json:"timeTrackable" db:"time_trackable" required:"true"`                             // time_trackable

	ProjectJoin *Project `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)

}

// KanbanStepCreateParams represents insert params for 'public.kanban_steps'.
type KanbanStepCreateParams struct {
	ProjectID     int    `json:"projectID"`                                                          // project_id
	StepOrder     int    `json:"stepOrder" required:"true"`                                          // step_order
	Name          string `json:"name" required:"true"`                                               // name
	Description   string `json:"description" required:"true"`                                        // description
	Color         string `json:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	TimeTrackable bool   `json:"timeTrackable" required:"true"`                                      // time_trackable
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

// KanbanStepUpdateParams represents update params for 'public.kanban_steps'.
type KanbanStepUpdateParams struct {
	ProjectID     *int    `json:"projectID"`                                                          // project_id
	StepOrder     *int    `json:"stepOrder" required:"true"`                                          // step_order
	Name          *string `json:"name" required:"true"`                                               // name
	Description   *string `json:"description" required:"true"`                                        // description
	Color         *string `json:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	TimeTrackable *bool   `json:"timeTrackable" required:"true"`                                      // time_trackable
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

type KanbanStepOrderBy string

const ()

type KanbanStepJoins struct {
	Project bool // O2O projects
}

// WithKanbanStepJoin joins with the given tables.
func WithKanbanStepJoin(joins KanbanStepJoins) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.joins = KanbanStepJoins{
			Project: s.joins.Project || joins.Project,
		}
	}
}

// WithKanbanStepFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithKanbanStepFilters(filters map[string][]any) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.filters = filters
	}
}

const kanbanStepTableProjectJoinSQL = `-- O2O join generated from "kanban_steps_project_id_fkey (Generated from M2O)"
left join projects as _kanban_steps_project_id on _kanban_steps_project_id.project_id = kanban_steps.project_id
`

const kanbanStepTableProjectSelectSQL = `(case when _kanban_steps_project_id.project_id is not null then row(_kanban_steps_project_id.*) end) as project_project_id`

const kanbanStepTableProjectGroupBySQL = `_kanban_steps_project_id.project_id,
      _kanban_steps_project_id.project_id,
	kanban_steps.kanban_step_id`

// Insert inserts the KanbanStep to the database.
func (ks *KanbanStep) Insert(ctx context.Context, db DB) (*KanbanStep, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.kanban_steps (
	project_id, step_order, name, description, color, time_trackable
	) VALUES (
	$1, $2, $3, $4, $5, $6
	) RETURNING * `
	// run
	logf(sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable)

	rows, err := db.Query(ctx, sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Insert/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	newks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Kanban step", Err: err}))
	}

	*ks = newks

	return ks, nil
}

// Update updates a KanbanStep in the database.
func (ks *KanbanStep) Update(ctx context.Context, db DB) (*KanbanStep, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.kanban_steps SET 
	project_id = $1, step_order = $2, name = $3, description = $4, color = $5, time_trackable = $6 
	WHERE kanban_step_id = $7 
	RETURNING * `
	// run
	logf(sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable, ks.KanbanStepID)

	rows, err := db.Query(ctx, sqlstr, ks.ProjectID, ks.StepOrder, ks.Name, ks.Description, ks.Color, ks.TimeTrackable, ks.KanbanStepID)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Update/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	newks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	*ks = newks

	return ks, nil
}

// Upsert upserts a KanbanStep in the database.
// Requires appropriate PK(s) to be set beforehand.
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Kanban step", Err: err})
			}
			ks, err = ks.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Kanban step", Err: err})
			}
		}
	}

	return ks, err
}

// Delete deletes the KanbanStep from the database.
func (ks *KanbanStep) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.kanban_steps 
	WHERE kanban_step_id = $1 `
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.kanban_step_id > $1
	 %s   %s 
  ORDER BY 
		kanban_step_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByKanbanStepIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{kanbanStepID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	return res, nil
}

// KanbanStepPaginatedByProjectIDAsc returns a cursor-paginated list of KanbanStep in Asc order.
func KanbanStepPaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.project_id > $1
	 %s   %s 
  ORDER BY 
		project_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByProjectIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	return res, nil
}

// KanbanStepPaginatedByStepOrderAsc returns a cursor-paginated list of KanbanStep in Asc order.
func KanbanStepPaginatedByStepOrderAsc(ctx context.Context, db DB, stepOrder int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.step_order > $1
	 %s   %s 
  ORDER BY 
		step_order Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByStepOrderAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{stepOrder}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	return res, nil
}

// KanbanStepPaginatedByKanbanStepIDDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByKanbanStepIDDesc(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.kanban_step_id < $1
	 %s   %s 
  ORDER BY 
		kanban_step_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByKanbanStepIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{kanbanStepID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	return res, nil
}

// KanbanStepPaginatedByProjectIDDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.project_id < $1
	 %s   %s 
  ORDER BY 
		project_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByProjectIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	return res, nil
}

// KanbanStepPaginatedByStepOrderDesc returns a cursor-paginated list of KanbanStep in Desc order.
func KanbanStepPaginatedByStepOrderDesc(ctx context.Context, db DB, stepOrder int, opts ...KanbanStepSelectConfigOption) ([]KanbanStep, error) {
	c := &KanbanStepSelectConfig{joins: KanbanStepJoins{}, filters: make(map[string][]any)}

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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.step_order < $1
	 %s   %s 
  ORDER BY 
		step_order Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* KanbanStepPaginatedByStepOrderDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{stepOrder}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.kanban_step_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepByKanbanStepID */\n" + sqlstr

	// run
	// logf(sqlstr, kanbanStepID)
	rows, err := db.Query(ctx, sqlstr, append([]any{kanbanStepID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByKanbanStepID/pgx.CollectOneRow: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	paramStart := 3
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.project_id = $1 AND kanban_steps.name = $2 AND kanban_steps.step_order = $3
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepByProjectIDNameStepOrder */\n" + sqlstr

	// run
	// logf(sqlstr, projectID, name, stepOrder)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID, name, stepOrder}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDNameStepOrder/pgx.CollectOneRow: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.project_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepsByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.name = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepsByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.step_order = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepsByStepOrder */\n" + sqlstr

	// run
	// logf(sqlstr, stepOrder)
	rows, err := db.Query(ctx, sqlstr, append([]any{stepOrder}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("KanbanStep/KanbanStepByProjectIDNameStepOrder/pgx.CollectRows: %w", &XoError{Entity: "Kanban step", Err: err}))
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

	paramStart := 2
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

	if c.joins.Project {
		selectClauses = append(selectClauses, kanbanStepTableProjectSelectSQL)
		joinClauses = append(joinClauses, kanbanStepTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, kanbanStepTableProjectGroupBySQL)
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
	kanban_steps.kanban_step_id,
	kanban_steps.project_id,
	kanban_steps.step_order,
	kanban_steps.name,
	kanban_steps.description,
	kanban_steps.color,
	kanban_steps.time_trackable %s 
	 FROM public.kanban_steps %s 
	 WHERE kanban_steps.project_id = $1 AND kanban_steps.step_order = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* KanbanStepByProjectIDStepOrder */\n" + sqlstr

	// run
	// logf(sqlstr, projectID, stepOrder)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID, stepOrder}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/db.Query: %w", &XoError{Entity: "Kanban step", Err: err}))
	}
	ks, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[KanbanStep])
	if err != nil {
		return nil, logerror(fmt.Errorf("kanban_steps/KanbanStepByProjectIDStepOrder/pgx.CollectOneRow: %w", &XoError{Entity: "Kanban step", Err: err}))
	}

	return &ks, nil
}

// FKProject_ProjectID returns the Project associated with the KanbanStep's (ProjectID).
//
// Generated from foreign key 'kanban_steps_project_id_fkey'.
func (ks *KanbanStep) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, ks.ProjectID)
}
