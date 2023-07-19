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

// WorkItemType represents a row from 'public.work_item_types'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItemType struct {
	WorkItemTypeID int    `json:"workItemTypeID" db:"work_item_type_id" required:"true"`                         // work_item_type_id
	ProjectID      int    `json:"-" db:"project_id"`                                                             // project_id
	Name           string `json:"name" db:"name" required:"true"`                                                // name
	Description    string `json:"description" db:"description" required:"true"`                                  // description
	Color          string `json:"color" db:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color

	ProjectJoin *Project `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)

}

// WorkItemTypeCreateParams represents insert params for 'public.work_item_types'.
type WorkItemTypeCreateParams struct {
	ProjectID   int    `json:"-"`                                                                  // project_id
	Name        string `json:"name" required:"true"`                                               // name
	Description string `json:"description" required:"true"`                                        // description
	Color       string `json:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
}

// CreateWorkItemType creates a new WorkItemType in the database with the given params.
func CreateWorkItemType(ctx context.Context, db DB, params *WorkItemTypeCreateParams) (*WorkItemType, error) {
	wit := &WorkItemType{
		ProjectID:   params.ProjectID,
		Name:        params.Name,
		Description: params.Description,
		Color:       params.Color,
	}

	return wit.Insert(ctx, db)
}

// WorkItemTypeUpdateParams represents update params for 'public.work_item_types'.
type WorkItemTypeUpdateParams struct {
	ProjectID   *int    `json:"-"`                                                                  // project_id
	Name        *string `json:"name" required:"true"`                                               // name
	Description *string `json:"description" required:"true"`                                        // description
	Color       *string `json:"color" required:"true" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
}

// SetUpdateParams updates public.work_item_types struct fields with the specified params.
func (wit *WorkItemType) SetUpdateParams(params *WorkItemTypeUpdateParams) {
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

type WorkItemTypeSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTypeJoins
	filters map[string][]any
}
type WorkItemTypeSelectConfigOption func(*WorkItemTypeSelectConfig)

// WithWorkItemTypeLimit limits row selection.
func WithWorkItemTypeLimit(limit int) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemTypeOrderBy string

const ()

type WorkItemTypeJoins struct {
	Project bool // O2O projects
}

// WithWorkItemTypeJoin joins with the given tables.
func WithWorkItemTypeJoin(joins WorkItemTypeJoins) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.joins = WorkItemTypeJoins{
			Project: s.joins.Project || joins.Project,
		}
	}
}

// WithWorkItemTypeFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemTypeFilters(filters map[string][]any) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.filters = filters
	}
}

const workItemTypeTableProjectJoinSQL = `-- O2O join generated from "work_item_types_project_id_fkey (Generated from M2O)"
left join projects as _work_item_types_project_id on _work_item_types_project_id.project_id = work_item_types.project_id
`

const workItemTypeTableProjectSelectSQL = `(case when _work_item_types_project_id.project_id is not null then row(_work_item_types_project_id.*) end) as project_project_id`

const workItemTypeTableProjectGroupBySQL = `_work_item_types_project_id.project_id,
      _work_item_types_project_id.project_id,
	work_item_types.work_item_type_id`

// Insert inserts the WorkItemType to the database.
func (wit *WorkItemType) Insert(ctx context.Context, db DB) (*WorkItemType, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_types (
	project_id, name, description, color
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Insert/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item type", Err: err}))
	}

	*wit = newwit

	return wit, nil
}

// Update updates a WorkItemType in the database.
func (wit *WorkItemType) Update(ctx context.Context, db DB) (*WorkItemType, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_types SET 
	project_id = $1, name = $2, description = $3, color = $4 
	WHERE work_item_type_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTypeID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Update/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	*wit = newwit

	return wit, nil
}

// Upsert upserts a WorkItemType in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wit *WorkItemType) Upsert(ctx context.Context, db DB, params *WorkItemTypeCreateParams) (*WorkItemType, error) {
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
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item type", Err: err})
			}
			wit, err = wit.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item type", Err: err})
			}
		}
	}

	return wit, err
}

// Delete deletes the WorkItemType from the database.
func (wit *WorkItemType) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_types 
	WHERE work_item_type_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTypeID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemTypePaginatedByWorkItemTypeIDAsc returns a cursor-paginated list of WorkItemType in Asc order.
func WorkItemTypePaginatedByWorkItemTypeIDAsc(ctx context.Context, db DB, workItemTypeID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.work_item_type_id > $1
	 %s   %s 
  ORDER BY 
		work_item_type_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypePaginatedByWorkItemTypeIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTypeID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Asc/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypePaginatedByProjectIDAsc returns a cursor-paginated list of WorkItemType in Asc order.
func WorkItemTypePaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.project_id > $1
	 %s   %s 
  ORDER BY 
		project_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypePaginatedByProjectIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Asc/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypePaginatedByWorkItemTypeIDDesc returns a cursor-paginated list of WorkItemType in Desc order.
func WorkItemTypePaginatedByWorkItemTypeIDDesc(ctx context.Context, db DB, workItemTypeID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.work_item_type_id < $1
	 %s   %s 
  ORDER BY 
		work_item_type_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypePaginatedByWorkItemTypeIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTypeID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Desc/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypePaginatedByProjectIDDesc returns a cursor-paginated list of WorkItemType in Desc order.
func WorkItemTypePaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.project_id < $1
	 %s   %s 
  ORDER BY 
		project_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypePaginatedByProjectIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Desc/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypeByNameProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypeByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.name = $1 AND work_item_types.project_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypeByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByNameProjectID/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByNameProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item type", Err: err}))
	}

	return &wit, nil
}

// WorkItemTypesByName retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypesByName(ctx context.Context, db DB, name string, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.name = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypesByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/WorkItemTypeByNameProjectID/Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/WorkItemTypeByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypesByProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypesByProjectID(ctx context.Context, db DB, projectID int, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.project_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypesByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/WorkItemTypeByNameProjectID/Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/WorkItemTypeByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypeByWorkItemTypeID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_pkey'.
func WorkItemTypeByWorkItemTypeID(ctx context.Context, db DB, workItemTypeID int, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTypeTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTypeTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTypeTableProjectGroupBySQL)
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
	work_item_types.work_item_type_id,
	work_item_types.project_id,
	work_item_types.name,
	work_item_types.description,
	work_item_types.color %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.work_item_type_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypeByWorkItemTypeID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTypeID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTypeID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByWorkItemTypeID/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_types/WorkItemTypeByWorkItemTypeID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item type", Err: err}))
	}

	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemType's (ProjectID).
//
// Generated from foreign key 'work_item_types_project_id_fkey'.
func (wit *WorkItemType) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
