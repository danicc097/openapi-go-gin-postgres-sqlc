package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// WorkItemType represents a row from 'public.work_item_types'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItemType struct {
	WorkItemTypeID WorkItemTypeID `json:"workItemTypeID" db:"work_item_type_id" required:"true" nullable:"false"`                         // work_item_type_id
	ProjectID      ProjectID      `json:"projectID" db:"project_id" required:"true" nullable:"false"`                                     // project_id
	Name           string         `json:"name" db:"name" required:"true" nullable:"false"`                                                // name
	Description    string         `json:"description" db:"description" required:"true" nullable:"false"`                                  // description
	Color          string         `json:"color" db:"color" required:"true" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color

	ProjectJoin *Project `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)

}

// WorkItemTypeCreateParams represents insert params for 'public.work_item_types'.
type WorkItemTypeCreateParams struct {
	Color       string    `json:"color" required:"true" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	Description string    `json:"description" required:"true" nullable:"false"`                                        // description
	Name        string    `json:"name" required:"true" nullable:"false"`                                               // name
	ProjectID   ProjectID `json:"-" openapi-go:"ignore"`                                                               // project_id
}

// WorkItemTypeParams represents common params for both insert and update of 'public.work_item_types'.
type WorkItemTypeParams interface {
	GetColor() *string
	GetDescription() *string
	GetName() *string
	GetProjectID() *ProjectID
}

func (p WorkItemTypeCreateParams) GetColor() *string {
	x := p.Color
	return &x
}
func (p WorkItemTypeUpdateParams) GetColor() *string {
	return p.Color
}

func (p WorkItemTypeCreateParams) GetDescription() *string {
	x := p.Description
	return &x
}
func (p WorkItemTypeUpdateParams) GetDescription() *string {
	return p.Description
}

func (p WorkItemTypeCreateParams) GetName() *string {
	x := p.Name
	return &x
}
func (p WorkItemTypeUpdateParams) GetName() *string {
	return p.Name
}

func (p WorkItemTypeCreateParams) GetProjectID() *ProjectID {
	x := p.ProjectID
	return &x
}
func (p WorkItemTypeUpdateParams) GetProjectID() *ProjectID {
	return p.ProjectID
}

type WorkItemTypeID int

// CreateWorkItemType creates a new WorkItemType in the database with the given params.
func CreateWorkItemType(ctx context.Context, db DB, params *WorkItemTypeCreateParams) (*WorkItemType, error) {
	wit := &WorkItemType{
		Color:       params.Color,
		Description: params.Description,
		Name:        params.Name,
		ProjectID:   params.ProjectID,
	}

	return wit.Insert(ctx, db)
}

type WorkItemTypeSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   WorkItemTypeJoins
	filters map[string][]any
	having  map[string][]any
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

// WithWorkItemTypeOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithWorkItemTypeOrderBy(rows map[string]*models.Direction) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		te := EntityFields[TableEntityWorkItemType]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type WorkItemTypeJoins struct {
	Project bool `json:"project" required:"true" nullable:"false"` // O2O projects
}

// WithWorkItemTypeJoin joins with the given tables.
func WithWorkItemTypeJoin(joins WorkItemTypeJoins) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.joins = WorkItemTypeJoins{
			Project: s.joins.Project || joins.Project,
		}
	}
}

// WithWorkItemTypeFilters adds the given WHERE clause conditions, which can be dynamically parameterized
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

// WithWorkItemTypeHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithWorkItemTypeHavingClause(conditions map[string][]any) WorkItemTypeSelectConfigOption {
	return func(s *WorkItemTypeSelectConfig) {
		s.having = conditions
	}
}

const workItemTypeTableProjectJoinSQL = `-- O2O join generated from "work_item_types_project_id_fkey (Generated from M2O)"
left join projects as _work_item_types_project_id on _work_item_types_project_id.project_id = work_item_types.project_id
`

const workItemTypeTableProjectSelectSQL = `(case when _work_item_types_project_id.project_id is not null then row(_work_item_types_project_id.*) end) as project_project_id`

const workItemTypeTableProjectGroupBySQL = `_work_item_types_project_id.project_id,
      _work_item_types_project_id.project_id,
	work_item_types.work_item_type_id`

// WorkItemTypeUpdateParams represents update params for 'public.work_item_types'.
type WorkItemTypeUpdateParams struct {
	Color       *string    `json:"color" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	Description *string    `json:"description" nullable:"false"`                                        // description
	Name        *string    `json:"name" nullable:"false"`                                               // name
	ProjectID   *ProjectID `json:"-" openapi-go:"ignore"`                                               // project_id
}

// SetUpdateParams updates public.work_item_types struct fields with the specified params.
func (wit *WorkItemType) SetUpdateParams(params *WorkItemTypeUpdateParams) {
	if params.Color != nil {
		wit.Color = *params.Color
	}
	if params.Description != nil {
		wit.Description = *params.Description
	}
	if params.Name != nil {
		wit.Name = *params.Name
	}
	if params.ProjectID != nil {
		wit.ProjectID = *params.ProjectID
	}
}

// Insert inserts the WorkItemType to the database.
func (wit *WorkItemType) Insert(ctx context.Context, db DB) (*WorkItemType, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_types (
	color, description, name, project_id
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, wit.Color, wit.Description, wit.Name, wit.ProjectID)

	rows, err := db.Query(ctx, sqlstr, wit.Color, wit.Description, wit.Name, wit.ProjectID)
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
	color = $1, description = $2, name = $3, project_id = $4 
	WHERE work_item_type_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, wit.Color, wit.Description, wit.Name, wit.ProjectID, wit.WorkItemTypeID)

	rows, err := db.Query(ctx, sqlstr, wit.Color, wit.Description, wit.Name, wit.ProjectID, wit.WorkItemTypeID)
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

	wit.Color = params.Color
	wit.Description = params.Description
	wit.Name = params.Name
	wit.ProjectID = params.ProjectID

	wit, err = wit.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertWorkItemType/Insert: %w", &XoError{Entity: "Work item type", Err: err})
			}
			wit, err = wit.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertWorkItemType/Update: %w", &XoError{Entity: "Work item type", Err: err})
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

// WorkItemTypePaginated returns a cursor-paginated list of WorkItemType.
// At least one cursor is required.
func WorkItemTypePaginated(ctx context.Context, db DB, cursors models.PaginationCursors, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		if cursor.Value == nil {

			return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
		}
		field, ok := EntityFields[TableEntityWorkItemType][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("WorkItemType/Paginated/cursor: %w", &XoError{Entity: "Work item type", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("work_item_types.%s %s $i", field.Db, op)] = []any{*cursor.Value}
		c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts
	}

	paramStart := 0 // all filters will come from the user
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
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/orderBy: %w", &XoError{Entity: "Work item type", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_types.color,
	work_item_types.description,
	work_item_types.name,
	work_item_types.project_id,
	work_item_types.work_item_type_id %s 
	 FROM public.work_item_types %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypePaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/db.Query: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemType])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemType/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item type", Err: err}))
	}
	return res, nil
}

// WorkItemTypeByNameProjectID retrieves a row from 'public.work_item_types' as a WorkItemType.
//
// Generated from index 'work_item_types_name_project_id_key'.
func WorkItemTypeByNameProjectID(ctx context.Context, db DB, name string, projectID ProjectID, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_types.color,
	work_item_types.description,
	work_item_types.name,
	work_item_types.project_id,
	work_item_types.work_item_type_id %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.name = $1 AND work_item_types.project_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypeByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, append(filterParams, havingParams...)...)...)
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
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_types.color,
	work_item_types.description,
	work_item_types.name,
	work_item_types.project_id,
	work_item_types.work_item_type_id %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.name = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypesByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
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
func WorkItemTypesByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...WorkItemTypeSelectConfigOption) ([]WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_types.color,
	work_item_types.description,
	work_item_types.name,
	work_item_types.project_id,
	work_item_types.work_item_type_id %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.project_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypesByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
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
func WorkItemTypeByWorkItemTypeID(ctx context.Context, db DB, workItemTypeID WorkItemTypeID, opts ...WorkItemTypeSelectConfigOption) (*WorkItemType, error) {
	c := &WorkItemTypeSelectConfig{joins: WorkItemTypeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

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
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_types.color,
	work_item_types.description,
	work_item_types.name,
	work_item_types.project_id,
	work_item_types.work_item_type_id %s 
	 FROM public.work_item_types %s 
	 WHERE work_item_types.work_item_type_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTypeByWorkItemTypeID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTypeID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTypeID}, append(filterParams, havingParams...)...)...)
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
