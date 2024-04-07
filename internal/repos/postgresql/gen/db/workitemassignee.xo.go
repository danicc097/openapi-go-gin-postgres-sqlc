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

// WorkItemAssignee represents a row from 'public.work_item_assignee'.
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
type WorkItemAssignee struct {
	WorkItemID WorkItemID          `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                           // work_item_id
	Assignee   UserID              `json:"assignee" db:"assignee" required:"true" nullable:"false"`                                 // assignee
	Role       models.WorkItemRole `json:"role" db:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role

	WorkItemsJoin *[]WorkItemAssigneeM2MWorkItemWIA `json:"-" db:"work_item_assignee_work_items" openapi-go:"ignore"` // M2M work_item_assignee
	AssigneesJoin *[]WorkItemAssigneeM2MAssigneeWIA `json:"-" db:"work_item_assignee_assignees" openapi-go:"ignore"`  // M2M work_item_assignee

}

// WorkItemAssigneeCreateParams represents insert params for 'public.work_item_assignee'.
type WorkItemAssigneeCreateParams struct {
	Assignee   UserID              `json:"assignee" required:"true" nullable:"false"`                                     // assignee
	Role       models.WorkItemRole `json:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID WorkItemID          `json:"workItemID" required:"true" nullable:"false"`                                   // work_item_id
}

// WorkItemAssigneeParams represents common params for both insert and update of 'public.work_item_assignee'.
type WorkItemAssigneeParams interface {
	GetAssignee() *UserID
	GetRole() *models.WorkItemRole
	GetWorkItemID() *WorkItemID
}

func (p WorkItemAssigneeCreateParams) GetAssignee() *UserID {
	x := p.Assignee
	return &x
}
func (p WorkItemAssigneeUpdateParams) GetAssignee() *UserID {
	return p.Assignee
}

func (p WorkItemAssigneeCreateParams) GetRole() *models.WorkItemRole {
	x := p.Role
	return &x
}
func (p WorkItemAssigneeUpdateParams) GetRole() *models.WorkItemRole {
	return p.Role
}

func (p WorkItemAssigneeCreateParams) GetWorkItemID() *WorkItemID {
	x := p.WorkItemID
	return &x
}
func (p WorkItemAssigneeUpdateParams) GetWorkItemID() *WorkItemID {
	return p.WorkItemID
}

// CreateWorkItemAssignee creates a new WorkItemAssignee in the database with the given params.
func CreateWorkItemAssignee(ctx context.Context, db DB, params *WorkItemAssigneeCreateParams) (*WorkItemAssignee, error) {
	wia := &WorkItemAssignee{
		Assignee:   params.Assignee,
		Role:       params.Role,
		WorkItemID: params.WorkItemID,
	}

	return wia.Insert(ctx, db)
}

type WorkItemAssigneeSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   WorkItemAssigneeJoins
	filters map[string][]any
	having  map[string][]any
}
type WorkItemAssigneeSelectConfigOption func(*WorkItemAssigneeSelectConfig)

// WithWorkItemAssigneeLimit limits row selection.
func WithWorkItemAssigneeLimit(limit int) WorkItemAssigneeSelectConfigOption {
	return func(s *WorkItemAssigneeSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithWorkItemAssigneeOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithWorkItemAssigneeOrderBy(rows map[string]*models.Direction) WorkItemAssigneeSelectConfigOption {
	return func(s *WorkItemAssigneeSelectConfig) {
		te := EntityFields[TableEntityWorkItemAssignee]
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

type WorkItemAssigneeJoins struct {
	WorkItems bool `json:"workItems" required:"true" nullable:"false"` // M2M work_item_assignee
	Assignees bool `json:"assignees" required:"true" nullable:"false"` // M2M work_item_assignee
}

// WithWorkItemAssigneeJoin joins with the given tables.
func WithWorkItemAssigneeJoin(joins WorkItemAssigneeJoins) WorkItemAssigneeSelectConfigOption {
	return func(s *WorkItemAssigneeSelectConfig) {
		s.joins = WorkItemAssigneeJoins{
			WorkItems: s.joins.WorkItems || joins.WorkItems,
			Assignees: s.joins.Assignees || joins.Assignees,
		}
	}
}

// WorkItemAssigneeM2MWorkItemWIA represents a M2M join against "public.work_item_assignee"
type WorkItemAssigneeM2MWorkItemWIA struct {
	WorkItem WorkItem            `json:"workItem" db:"work_items" required:"true"`
	Role     models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WorkItemAssigneeM2MAssigneeWIA represents a M2M join against "public.work_item_assignee"
type WorkItemAssigneeM2MAssigneeWIA struct {
	User User                `json:"user" db:"users" required:"true"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithWorkItemAssigneeFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemAssigneeFilters(filters map[string][]any) WorkItemAssigneeSelectConfigOption {
	return func(s *WorkItemAssigneeSelectConfig) {
		s.filters = filters
	}
}

// WithWorkItemAssigneeHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithWorkItemAssigneeHavingClause(conditions map[string][]any) WorkItemAssigneeSelectConfigOption {
	return func(s *WorkItemAssigneeSelectConfig) {
		s.having = conditions
	}
}

const workItemAssigneeTableWorkItemsJoinSQL = `-- M2M join generated from "work_item_assignee_work_item_id_fkey"
left join (
	select
		work_item_assignee.assignee as work_item_assignee_assignee
		, work_item_assignee.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		work_item_assignee
	join work_items on work_items.work_item_id = work_item_assignee.work_item_id
	group by
		work_item_assignee_assignee
		, work_items.work_item_id
		, role
) as xo_join_work_item_assignee_work_items on xo_join_work_item_assignee_work_items.work_item_assignee_assignee = work_item_assignee.assignee
`

const workItemAssigneeTableWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_work_items.__work_items
		, xo_join_work_item_assignee_work_items.role
		)) filter (where xo_join_work_item_assignee_work_items.__work_items_work_item_id is not null), '{}') as work_item_assignee_work_items`

const workItemAssigneeTableWorkItemsGroupBySQL = `work_item_assignee.assignee, work_item_assignee.work_item_id, work_item_assignee.assignee`

const workItemAssigneeTableAssigneesJoinSQL = `-- M2M join generated from "work_item_assignee_assignee_fkey"
left join (
	select
		work_item_assignee.work_item_id as work_item_assignee_work_item_id
		, work_item_assignee.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		work_item_assignee
	join users on users.user_id = work_item_assignee.assignee
	group by
		work_item_assignee_work_item_id
		, users.user_id
		, role
) as xo_join_work_item_assignee_assignees on xo_join_work_item_assignee_assignees.work_item_assignee_work_item_id = work_item_assignee.work_item_id
`

const workItemAssigneeTableAssigneesSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_assignees.__users
		, xo_join_work_item_assignee_assignees.role
		)) filter (where xo_join_work_item_assignee_assignees.__users_user_id is not null), '{}') as work_item_assignee_assignees`

const workItemAssigneeTableAssigneesGroupBySQL = `work_item_assignee.work_item_id, work_item_assignee.work_item_id, work_item_assignee.assignee`

// WorkItemAssigneeUpdateParams represents update params for 'public.work_item_assignee'.
type WorkItemAssigneeUpdateParams struct {
	Assignee   *UserID              `json:"assignee" nullable:"false"`                                     // assignee
	Role       *models.WorkItemRole `json:"role" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID *WorkItemID          `json:"workItemID" nullable:"false"`                                   // work_item_id
}

// SetUpdateParams updates public.work_item_assignee struct fields with the specified params.
func (wia *WorkItemAssignee) SetUpdateParams(params *WorkItemAssigneeUpdateParams) {
	if params.Assignee != nil {
		wia.Assignee = *params.Assignee
	}
	if params.Role != nil {
		wia.Role = *params.Role
	}
	if params.WorkItemID != nil {
		wia.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the WorkItemAssignee to the database.
func (wia *WorkItemAssignee) Insert(ctx context.Context, db DB) (*WorkItemAssignee, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_assignee (
	assignee, role, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, wia.Assignee, wia.Role, wia.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, wia.Assignee, wia.Role, wia.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Insert/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	newwia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	*wia = newwia

	return wia, nil
}

// Update updates a WorkItemAssignee in the database.
func (wia *WorkItemAssignee) Update(ctx context.Context, db DB) (*WorkItemAssignee, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_assignee SET 
	role = $1 
	WHERE work_item_id = $2  AND assignee = $3 
	RETURNING * `
	// run
	logf(sqlstr, wia.Role, wia.WorkItemID, wia.Assignee)

	rows, err := db.Query(ctx, sqlstr, wia.Role, wia.WorkItemID, wia.Assignee)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Update/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	newwia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	*wia = newwia

	return wia, nil
}

// Upsert upserts a WorkItemAssignee in the database.
// Requires appropriate PK(s) to be set beforehand.
func (wia *WorkItemAssignee) Upsert(ctx context.Context, db DB, params *WorkItemAssigneeCreateParams) (*WorkItemAssignee, error) {
	var err error

	wia.Assignee = params.Assignee
	wia.Role = params.Role
	wia.WorkItemID = params.WorkItemID

	wia, err = wia.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertWorkItemAssignee/Insert: %w", &XoError{Entity: "Work item assignee", Err: err})
			}
			wia, err = wia.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertWorkItemAssignee/Update: %w", &XoError{Entity: "Work item assignee", Err: err})
			}
		}
	}

	return wia, err
}

// Delete deletes the WorkItemAssignee from the database.
func (wia *WorkItemAssignee) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_assignee 
	WHERE work_item_id = $1 AND assignee = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wia.WorkItemID, wia.Assignee); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemAssigneePaginated returns a cursor-paginated list of WorkItemAssignee.
// At least one cursor is required.
func WorkItemAssigneePaginated(ctx context.Context, db DB, cursors []Cursor, opts ...WorkItemAssigneeSelectConfigOption) ([]WorkItemAssignee, error) {
	c := &WorkItemAssigneeSelectConfig{joins: WorkItemAssigneeJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		field, ok := EntityFields[TableEntityWorkItemAssignee][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("WorkItemAssignee/Paginated/cursor: %w", &XoError{Entity: "Work item assignee", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("work_item_assignee.%s %s $i", field.Db, op)] = []any{cursor.Value}
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
		filters += " where "
	}
	if len(filterClauses) > 0 {
		filters += strings.Join(filterClauses, " AND ") + " "
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
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Paginated/orderBy: %w", &XoError{Entity: "Work item assignee", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, workItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableAssigneesGroupBySQL)
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
	work_item_assignee.assignee,
	work_item_assignee.role,
	work_item_assignee.work_item_id %s 
	 FROM public.work_item_assignee %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssigneePaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Paginated/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// WorkItemAssigneesByAssigneeWorkItemID retrieves a row from 'public.work_item_assignee' as a WorkItemAssignee.
//
// Generated from index 'work_item_assignee_assignee_work_item_id_idx'.
func WorkItemAssigneesByAssigneeWorkItemID(ctx context.Context, db DB, assignee UserID, workItemID WorkItemID, opts ...WorkItemAssigneeSelectConfigOption) ([]WorkItemAssignee, error) {
	c := &WorkItemAssigneeSelectConfig{joins: WorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, workItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableAssigneesGroupBySQL)
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
	work_item_assignee.assignee,
	work_item_assignee.role,
	work_item_assignee.work_item_id %s 
	 FROM public.work_item_assignee %s 
	 WHERE work_item_assignee.assignee = $1 AND work_item_assignee.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssigneesByAssigneeWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, assignee, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignee, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByAssigneeWorkItemID/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByAssigneeWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// WorkItemAssigneeByWorkItemIDAssignee retrieves a row from 'public.work_item_assignee' as a WorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func WorkItemAssigneeByWorkItemIDAssignee(ctx context.Context, db DB, workItemID WorkItemID, assignee UserID, opts ...WorkItemAssigneeSelectConfigOption) (*WorkItemAssignee, error) {
	c := &WorkItemAssigneeSelectConfig{joins: WorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, workItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableAssigneesGroupBySQL)
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
	work_item_assignee.assignee,
	work_item_assignee.role,
	work_item_assignee.work_item_id %s 
	 FROM public.work_item_assignee %s 
	 WHERE work_item_assignee.work_item_id = $1 AND work_item_assignee.assignee = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssigneeByWorkItemIDAssignee */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, assignee)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, assignee}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assignee/WorkItemAssigneeByWorkItemIDAssignee/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	wia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}

	return &wia, nil
}

// WorkItemAssigneesByWorkItemID retrieves a row from 'public.work_item_assignee' as a WorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func WorkItemAssigneesByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...WorkItemAssigneeSelectConfigOption) ([]WorkItemAssignee, error) {
	c := &WorkItemAssigneeSelectConfig{joins: WorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, workItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableAssigneesGroupBySQL)
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
	work_item_assignee.assignee,
	work_item_assignee.role,
	work_item_assignee.work_item_id %s 
	 FROM public.work_item_assignee %s 
	 WHERE work_item_assignee.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssigneesByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// WorkItemAssigneesByAssignee retrieves a row from 'public.work_item_assignee' as a WorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func WorkItemAssigneesByAssignee(ctx context.Context, db DB, assignee UserID, opts ...WorkItemAssigneeSelectConfigOption) ([]WorkItemAssignee, error) {
	c := &WorkItemAssigneeSelectConfig{joins: WorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, workItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, workItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, workItemAssigneeTableAssigneesGroupBySQL)
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
	work_item_assignee.assignee,
	work_item_assignee.role,
	work_item_assignee.work_item_id %s 
	 FROM public.work_item_assignee %s 
	 WHERE work_item_assignee.assignee = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemAssigneesByAssignee */\n" + sqlstr

	// run
	// logf(sqlstr, assignee)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignee}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// FKUser_Assignee returns the User associated with the WorkItemAssignee's (Assignee).
//
// Generated from foreign key 'work_item_assignee_assignee_fkey'.
func (wia *WorkItemAssignee) FKUser_Assignee(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wia.Assignee)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemAssignee's (WorkItemID).
//
// Generated from foreign key 'work_item_assignee_work_item_id_fkey'.
func (wia *WorkItemAssignee) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wia.WorkItemID)
}
