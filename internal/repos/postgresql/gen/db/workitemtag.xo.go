package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// WorkItemTag represents a row from 'public.work_item_tags'.
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
type WorkItemTag struct {
	WorkItemTagID WorkItemTagID `json:"workItemTagID" db:"work_item_tag_id" required:"true" nullable:"false"`                           // work_item_tag_id
	ProjectID     ProjectID     `json:"projectID" db:"project_id" required:"true" nullable:"false"`                                     // project_id
	Name          string        `json:"name" db:"name" required:"true" nullable:"false"`                                                // name
	Description   string        `json:"description" db:"description" required:"true" nullable:"false"`                                  // description
	Color         string        `json:"color" db:"color" required:"true" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	DeletedAt     *time.Time    `json:"deletedAt" db:"deleted_at"`                                                                      // deleted_at

	ProjectJoin   *Project    `json:"-" db:"project_project_id" openapi-go:"ignore"`                 // O2O projects (generated from M2O)
	WorkItemsJoin *[]WorkItem `json:"-" db:"work_item_work_item_tag_work_items" openapi-go:"ignore"` // M2M work_item_work_item_tag

}

// WorkItemTagCreateParams represents insert params for 'public.work_item_tags'.
type WorkItemTagCreateParams struct {
	Color       string    `json:"color" required:"true" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	Description string    `json:"description" required:"true" nullable:"false"`                                        // description
	Name        string    `json:"name" required:"true" nullable:"false"`                                               // name
	ProjectID   ProjectID `json:"-" openapi-go:"ignore"`                                                               // project_id
}

// WorkItemTagParams represents common params for both insert and update of 'public.work_item_tags'.
type WorkItemTagParams interface {
	GetColor() *string
	GetDescription() *string
	GetName() *string
	GetProjectID() *ProjectID
}

func (p WorkItemTagCreateParams) GetColor() *string {
	x := p.Color
	return &x
}
func (p WorkItemTagUpdateParams) GetColor() *string {
	return p.Color
}

func (p WorkItemTagCreateParams) GetDescription() *string {
	x := p.Description
	return &x
}
func (p WorkItemTagUpdateParams) GetDescription() *string {
	return p.Description
}

func (p WorkItemTagCreateParams) GetName() *string {
	x := p.Name
	return &x
}
func (p WorkItemTagUpdateParams) GetName() *string {
	return p.Name
}

func (p WorkItemTagCreateParams) GetProjectID() *ProjectID {
	x := p.ProjectID
	return &x
}
func (p WorkItemTagUpdateParams) GetProjectID() *ProjectID {
	return p.ProjectID
}

type WorkItemTagID int

// CreateWorkItemTag creates a new WorkItemTag in the database with the given params.
func CreateWorkItemTag(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
	wit := &WorkItemTag{
		Color:       params.Color,
		Description: params.Description,
		Name:        params.Name,
		ProjectID:   params.ProjectID,
	}

	return wit.Insert(ctx, db)
}

type WorkItemTagSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   WorkItemTagJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
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

// WithDeletedWorkItemTagOnly limits result to records marked as deleted.
func WithDeletedWorkItemTagOnly() WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.deletedAt = " not null "
	}
}

// WithWorkItemTagOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithWorkItemTagOrderBy(rows map[string]*models.Direction) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		te := EntityFields[TableEntityWorkItemTag]
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

type WorkItemTagJoins struct {
	Project   bool `json:"project" required:"true" nullable:"false"`   // O2O projects
	WorkItems bool `json:"workItems" required:"true" nullable:"false"` // M2M work_item_work_item_tag
}

// WithWorkItemTagJoin joins with the given tables.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = WorkItemTagJoins{
			Project:   s.joins.Project || joins.Project,
			WorkItems: s.joins.WorkItems || joins.WorkItems,
		}
	}
}

// WithWorkItemTagFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemTagFilters(filters map[string][]any) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.filters = filters
	}
}

// WithWorkItemTagHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithWorkItemTagHavingClause(conditions map[string][]any) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.having = conditions
	}
}

const workItemTagTableProjectJoinSQL = `-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _work_item_tags_project_id on _work_item_tags_project_id.project_id = work_item_tags.project_id
`

const workItemTagTableProjectSelectSQL = `(case when _work_item_tags_project_id.project_id is not null then row(_work_item_tags_project_id.*) end) as project_project_id`

const workItemTagTableProjectGroupBySQL = `_work_item_tags_project_id.project_id,
      _work_item_tags_project_id.project_id,
	work_item_tags.work_item_tag_id`

const workItemTagTableWorkItemsJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
		work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		work_item_work_item_tag
	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
	group by
		work_item_work_item_tag_work_item_tag_id
		, work_items.work_item_id
) as xo_join_work_item_work_item_tag_work_items on xo_join_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
`

const workItemTagTableWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_work_item_tag_work_items.__work_items
		)) filter (where xo_join_work_item_work_item_tag_work_items.__work_items_work_item_id is not null), '{}') as work_item_work_item_tag_work_items`

const workItemTagTableWorkItemsGroupBySQL = `work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id`

// WorkItemTagUpdateParams represents update params for 'public.work_item_tags'.
type WorkItemTagUpdateParams struct {
	Color       *string    `json:"color" nullable:"false" pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"` // color
	Description *string    `json:"description" nullable:"false"`                                        // description
	Name        *string    `json:"name" nullable:"false"`                                               // name
	ProjectID   *ProjectID `json:"-" openapi-go:"ignore"`                                               // project_id
}

// SetUpdateParams updates public.work_item_tags struct fields with the specified params.
func (wit *WorkItemTag) SetUpdateParams(params *WorkItemTagUpdateParams) {
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

// Insert inserts the WorkItemTag to the database.
func (wit *WorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemTag, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_tags (
	color, deleted_at, description, name, project_id
	) VALUES (
	$1, $2, $3, $4, $5
	) RETURNING * `
	// run
	logf(sqlstr, wit.Color, wit.DeletedAt, wit.Description, wit.Name, wit.ProjectID)

	rows, err := db.Query(ctx, sqlstr, wit.Color, wit.DeletedAt, wit.Description, wit.Name, wit.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/db.Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item tag", Err: err}))
	}

	*wit = newwit

	return wit, nil
}

// Update updates a WorkItemTag in the database.
func (wit *WorkItemTag) Update(ctx context.Context, db DB) (*WorkItemTag, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_tags SET 
	color = $1, deleted_at = $2, description = $3, name = $4, project_id = $5 
	WHERE work_item_tag_id = $6 
	RETURNING * `
	// run
	logf(sqlstr, wit.Color, wit.DeletedAt, wit.Description, wit.Name, wit.ProjectID, wit.WorkItemTagID)

	rows, err := db.Query(ctx, sqlstr, wit.Color, wit.DeletedAt, wit.Description, wit.Name, wit.ProjectID, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/db.Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	*wit = newwit

	return wit, nil
}

// Upsert upserts a WorkItemTag in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
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
				return nil, fmt.Errorf("UpsertWorkItemTag/Insert: %w", &XoError{Entity: "Work item tag", Err: err})
			}
			wit, err = wit.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertWorkItemTag/Update: %w", &XoError{Entity: "Work item tag", Err: err})
			}
		}
	}

	return wit, err
}

// Delete deletes the WorkItemTag from the database.
func (wit *WorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_tags 
	WHERE work_item_tag_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the WorkItemTag from the database via 'deleted_at'.
func (wit *WorkItemTag) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.work_item_tags 
	SET deleted_at = NOW() 
	WHERE work_item_tag_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	// set deleted
	wit.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted WorkItemTag from the database.
func (wit *WorkItemTag) Restore(ctx context.Context, db DB) (*WorkItemTag, error) {
	wit.DeletedAt = nil
	newwit, err := wit.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Restore/pgx.CollectRows: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	return newwit, nil
}

// WorkItemTagPaginated returns a cursor-paginated list of WorkItemTag.
// At least one cursor is required.
func WorkItemTagPaginated(ctx context.Context, db DB, cursors models.PaginationCursors, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{deletedAt: " null ", joins: WorkItemTagJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		field, ok := EntityFields[TableEntityWorkItemTag][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/cursor: %w", &XoError{Entity: "Work item tag", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("work_item_tags.%s %s $i", field.Db, op)] = []any{cursor.Value}
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
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/orderBy: %w", &XoError{Entity: "Work item tag", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsGroupBySQL)
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
	work_item_tags.color,
	work_item_tags.deleted_at,
	work_item_tags.description,
	work_item_tags.name,
	work_item_tags.project_id,
	work_item_tags.work_item_tag_id %s 
	 FROM public.work_item_tags %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/db.Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemTagByNameProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagByNameProjectID(ctx context.Context, db DB, name string, projectID ProjectID, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{deletedAt: " null ", joins: WorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsGroupBySQL)
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
	work_item_tags.color,
	work_item_tags.deleted_at,
	work_item_tags.description,
	work_item_tags.name,
	work_item_tags.project_id,
	work_item_tags.work_item_tag_id %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.name = $1 AND work_item_tags.project_id = $2
	 %s   AND work_item_tags.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/db.Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item tag", Err: err}))
	}

	return &wit, nil
}

// WorkItemTagsByName retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByName(ctx context.Context, db DB, name string, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{deletedAt: " null ", joins: WorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsGroupBySQL)
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
	work_item_tags.color,
	work_item_tags.deleted_at,
	work_item_tags.description,
	work_item_tags.name,
	work_item_tags.project_id,
	work_item_tags.work_item_tag_id %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.name = $1
	 %s   AND work_item_tags.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagsByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemTagsByProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{deletedAt: " null ", joins: WorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsGroupBySQL)
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
	work_item_tags.color,
	work_item_tags.deleted_at,
	work_item_tags.description,
	work_item_tags.name,
	work_item_tags.project_id,
	work_item_tags.work_item_tag_id %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.project_id = $1
	 %s   AND work_item_tags.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagsByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemTagByWorkItemTagID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_pkey'.
func WorkItemTagByWorkItemTagID(ctx context.Context, db DB, workItemTagID WorkItemTagID, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{deletedAt: " null ", joins: WorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsGroupBySQL)
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
	work_item_tags.color,
	work_item_tags.deleted_at,
	work_item_tags.description,
	work_item_tags.name,
	work_item_tags.project_id,
	work_item_tags.work_item_tag_id %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.work_item_tag_id = $1
	 %s   AND work_item_tags.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagByWorkItemTagID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/db.Query: %w", &XoError{Entity: "Work item tag", Err: err}))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item tag", Err: err}))
	}

	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemTag's (ProjectID).
//
// Generated from foreign key 'work_item_tags_project_id_fkey'.
func (wit *WorkItemTag) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
