// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

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

// ExtraSchemaWorkItemAssignee represents a row from 'extra_schema.work_item_assignee'.
type ExtraSchemaWorkItemAssignee struct {
	WorkItemID      ExtraSchemaWorkItemID    `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                           // work_item_id
	Assignee        ExtraSchemaUserID        `json:"assignee" db:"assignee" required:"true" nullable:"false"`                                 // assignee
	ExtraSchemaRole *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role

	WorkItemsJoin *[]ExtraSchemaWorkItemAssigneeM2MWorkItemWIA `json:"-" db:"work_item_assignee_work_items"` // M2M work_item_assignee
	AssigneesJoin *[]ExtraSchemaWorkItemAssigneeM2MAssigneeWIA `json:"-" db:"work_item_assignee_assignees"`  // M2M work_item_assignee

}

// ExtraSchemaWorkItemAssigneeCreateParams represents insert params for 'extra_schema.work_item_assignee'.
type ExtraSchemaWorkItemAssigneeCreateParams struct {
	Assignee        ExtraSchemaUserID        `json:"assignee" required:"true" nullable:"false"`                                     // assignee
	ExtraSchemaRole *ExtraSchemaWorkItemRole `json:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID      ExtraSchemaWorkItemID    `json:"workItemID" required:"true" nullable:"false"`                                   // work_item_id
}

// ExtraSchemaWorkItemAssigneeParams represents common params for both insert and update of 'extra_schema.work_item_assignee'.
type ExtraSchemaWorkItemAssigneeParams interface {
	GetAssignee() *ExtraSchemaUserID
	GetExtraSchemaRole() *ExtraSchemaWorkItemRole
	GetWorkItemID() *ExtraSchemaWorkItemID
}

func (p ExtraSchemaWorkItemAssigneeCreateParams) GetAssignee() *ExtraSchemaUserID {
	x := p.Assignee
	return &x
}
func (p ExtraSchemaWorkItemAssigneeUpdateParams) GetAssignee() *ExtraSchemaUserID {
	return p.Assignee
}

func (p ExtraSchemaWorkItemAssigneeCreateParams) GetExtraSchemaRole() *ExtraSchemaWorkItemRole {
	return p.ExtraSchemaRole
}
func (p ExtraSchemaWorkItemAssigneeUpdateParams) GetExtraSchemaRole() *ExtraSchemaWorkItemRole {
	if p.ExtraSchemaRole != nil {
		return *p.ExtraSchemaRole
	}
	return nil
}

func (p ExtraSchemaWorkItemAssigneeCreateParams) GetWorkItemID() *ExtraSchemaWorkItemID {
	x := p.WorkItemID
	return &x
}
func (p ExtraSchemaWorkItemAssigneeUpdateParams) GetWorkItemID() *ExtraSchemaWorkItemID {
	return p.WorkItemID
}

// CreateExtraSchemaWorkItemAssignee creates a new ExtraSchemaWorkItemAssignee in the database with the given params.
func CreateExtraSchemaWorkItemAssignee(ctx context.Context, db DB, params *ExtraSchemaWorkItemAssigneeCreateParams) (*ExtraSchemaWorkItemAssignee, error) {
	eswia := &ExtraSchemaWorkItemAssignee{
		Assignee:        params.Assignee,
		ExtraSchemaRole: params.ExtraSchemaRole,
		WorkItemID:      params.WorkItemID,
	}

	return eswia.Insert(ctx, db)
}

type ExtraSchemaWorkItemAssigneeSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   ExtraSchemaWorkItemAssigneeJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaWorkItemAssigneeSelectConfigOption func(*ExtraSchemaWorkItemAssigneeSelectConfig)

// WithExtraSchemaWorkItemAssigneeLimit limits row selection.
func WithExtraSchemaWorkItemAssigneeLimit(limit int) ExtraSchemaWorkItemAssigneeSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssigneeSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaWorkItemAssigneeOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaWorkItemAssigneeOrderBy(rows map[string]*Direction) ExtraSchemaWorkItemAssigneeSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssigneeSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaWorkItemAssignee]
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

type ExtraSchemaWorkItemAssigneeJoins struct {
	WorkItems bool `json:"workItems" required:"true" nullable:"false"` // M2M work_item_assignee
	Assignees bool `json:"assignees" required:"true" nullable:"false"` // M2M work_item_assignee
}

// WithExtraSchemaWorkItemAssigneeJoin joins with the given tables.
func WithExtraSchemaWorkItemAssigneeJoin(joins ExtraSchemaWorkItemAssigneeJoins) ExtraSchemaWorkItemAssigneeSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssigneeSelectConfig) {
		s.joins = ExtraSchemaWorkItemAssigneeJoins{
			WorkItems: s.joins.WorkItems || joins.WorkItems,
			Assignees: s.joins.Assignees || joins.Assignees,
		}
	}
}

// ExtraSchemaWorkItemAssigneeM2MWorkItemWIA represents a M2M join against "extra_schema.work_item_assignee"
type ExtraSchemaWorkItemAssigneeM2MWorkItemWIA struct {
	WorkItem ExtraSchemaWorkItem      `json:"workItem" db:"work_items" required:"true"`
	Role     *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// ExtraSchemaWorkItemAssigneeM2MAssigneeWIA represents a M2M join against "extra_schema.work_item_assignee"
type ExtraSchemaWorkItemAssigneeM2MAssigneeWIA struct {
	User ExtraSchemaUser          `json:"user" db:"users" required:"true"`
	Role *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithExtraSchemaWorkItemAssigneeFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaWorkItemAssigneeFilters(filters map[string][]any) ExtraSchemaWorkItemAssigneeSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssigneeSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaWorkItemAssigneeHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaWorkItemAssigneeHavingClause(conditions map[string][]any) ExtraSchemaWorkItemAssigneeSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssigneeSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL = `-- M2M join generated from "work_item_assignee_work_item_id_fkey"
left join (
	select
		work_item_assignee.assignee as work_item_assignee_assignee
		, work_item_assignee.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		extra_schema.work_item_assignee
	join extra_schema.work_items on work_items.work_item_id = work_item_assignee.work_item_id
	group by
		work_item_assignee_assignee
		, work_items.work_item_id
		, role
) as xo_join_work_item_assignee_work_items on xo_join_work_item_assignee_work_items.work_item_assignee_assignee = work_item_assignee.assignee
`

const extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_work_items.__work_items
		, xo_join_work_item_assignee_work_items.role
		)) filter (where xo_join_work_item_assignee_work_items.__work_items_work_item_id is not null), '{}') as work_item_assignee_work_items`

const extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL = `work_item_assignee.assignee, work_item_assignee.work_item_id, work_item_assignee.assignee`

const extraSchemaWorkItemAssigneeTableAssigneesJoinSQL = `-- M2M join generated from "work_item_assignee_assignee_fkey"
left join (
	select
		work_item_assignee.work_item_id as work_item_assignee_work_item_id
		, work_item_assignee.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.work_item_assignee
	join extra_schema.users on users.user_id = work_item_assignee.assignee
	group by
		work_item_assignee_work_item_id
		, users.user_id
		, role
) as xo_join_work_item_assignee_assignees on xo_join_work_item_assignee_assignees.work_item_assignee_work_item_id = work_item_assignee.work_item_id
`

const extraSchemaWorkItemAssigneeTableAssigneesSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_assignees.__users
		, xo_join_work_item_assignee_assignees.role
		)) filter (where xo_join_work_item_assignee_assignees.__users_user_id is not null), '{}') as work_item_assignee_assignees`

const extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL = `work_item_assignee.work_item_id, work_item_assignee.work_item_id, work_item_assignee.assignee`

// ExtraSchemaWorkItemAssigneeUpdateParams represents update params for 'extra_schema.work_item_assignee'.
type ExtraSchemaWorkItemAssigneeUpdateParams struct {
	Assignee        *ExtraSchemaUserID        `json:"assignee" nullable:"false"`                                     // assignee
	ExtraSchemaRole **ExtraSchemaWorkItemRole `json:"role" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID      *ExtraSchemaWorkItemID    `json:"workItemID" nullable:"false"`                                   // work_item_id
}

// SetUpdateParams updates extra_schema.work_item_assignee struct fields with the specified params.
func (eswia *ExtraSchemaWorkItemAssignee) SetUpdateParams(params *ExtraSchemaWorkItemAssigneeUpdateParams) {
	if params.Assignee != nil {
		eswia.Assignee = *params.Assignee
	}
	if params.ExtraSchemaRole != nil {
		eswia.ExtraSchemaRole = *params.ExtraSchemaRole
	}
	if params.WorkItemID != nil {
		eswia.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the ExtraSchemaWorkItemAssignee to the database.
func (eswia *ExtraSchemaWorkItemAssignee) Insert(ctx context.Context, db DB) (*ExtraSchemaWorkItemAssignee, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.work_item_assignee (
	assignee, role, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, eswia.Assignee, eswia.ExtraSchemaRole, eswia.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, eswia.Assignee, eswia.ExtraSchemaRole, eswia.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Insert/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	neweswia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	*eswia = neweswia

	return eswia, nil
}

// Update updates a ExtraSchemaWorkItemAssignee in the database.
func (eswia *ExtraSchemaWorkItemAssignee) Update(ctx context.Context, db DB) (*ExtraSchemaWorkItemAssignee, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.work_item_assignee SET 
	role = $1 
	WHERE work_item_id = $2  AND assignee = $3 
	RETURNING * `
	// run
	logf(sqlstr, eswia.ExtraSchemaRole, eswia.WorkItemID, eswia.Assignee)

	rows, err := db.Query(ctx, sqlstr, eswia.ExtraSchemaRole, eswia.WorkItemID, eswia.Assignee)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Update/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	neweswia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	*eswia = neweswia

	return eswia, nil
}

// Upsert upserts a ExtraSchemaWorkItemAssignee in the database.
// Requires appropriate PK(s) to be set beforehand.
func (eswia *ExtraSchemaWorkItemAssignee) Upsert(ctx context.Context, db DB, params *ExtraSchemaWorkItemAssigneeCreateParams) (*ExtraSchemaWorkItemAssignee, error) {
	var err error

	eswia.Assignee = params.Assignee
	eswia.ExtraSchemaRole = params.ExtraSchemaRole
	eswia.WorkItemID = params.WorkItemID

	eswia, err = eswia.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaWorkItemAssignee/Insert: %w", &XoError{Entity: "Work item assignee", Err: err})
			}
			eswia, err = eswia.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaWorkItemAssignee/Update: %w", &XoError{Entity: "Work item assignee", Err: err})
			}
		}
	}

	return eswia, err
}

// Delete deletes the ExtraSchemaWorkItemAssignee from the database.
func (eswia *ExtraSchemaWorkItemAssignee) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM extra_schema.work_item_assignee 
	WHERE work_item_id = $1 AND assignee = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, eswia.WorkItemID, eswia.Assignee); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaWorkItemAssigneePaginated returns a cursor-paginated list of ExtraSchemaWorkItemAssignee.
// At least one cursor is required.
func ExtraSchemaWorkItemAssigneePaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaWorkItemAssigneeSelectConfigOption) ([]ExtraSchemaWorkItemAssignee, error) {
	c := &ExtraSchemaWorkItemAssigneeSelectConfig{joins: ExtraSchemaWorkItemAssigneeJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {

		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaWorkItemAssignee][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Paginated/cursor: %w", &XoError{Entity: "Work item assignee", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("work_item_assignee.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

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
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Paginated/orderBy: %w", &XoError{Entity: "Work item assignee", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL)
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
	 FROM extra_schema.work_item_assignee %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssigneePaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Paginated/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAssigneesByAssigneeWorkItemID retrieves a row from 'extra_schema.work_item_assignee' as a ExtraSchemaWorkItemAssignee.
//
// Generated from index 'work_item_assignee_assignee_work_item_id_idx'.
func ExtraSchemaWorkItemAssigneesByAssigneeWorkItemID(ctx context.Context, db DB, assignee ExtraSchemaUserID, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAssigneeSelectConfigOption) ([]ExtraSchemaWorkItemAssignee, error) {
	c := &ExtraSchemaWorkItemAssigneeSelectConfig{joins: ExtraSchemaWorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL)
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
	 FROM extra_schema.work_item_assignee %s 
	 WHERE work_item_assignee.assignee = $1 AND work_item_assignee.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssigneesByAssigneeWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, assignee, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignee, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByAssigneeWorkItemID/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByAssigneeWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAssigneeByWorkItemIDAssignee retrieves a row from 'extra_schema.work_item_assignee' as a ExtraSchemaWorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func ExtraSchemaWorkItemAssigneeByWorkItemIDAssignee(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, assignee ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAssigneeSelectConfigOption) (*ExtraSchemaWorkItemAssignee, error) {
	c := &ExtraSchemaWorkItemAssigneeSelectConfig{joins: ExtraSchemaWorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL)
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
	 FROM extra_schema.work_item_assignee %s 
	 WHERE work_item_assignee.work_item_id = $1 AND work_item_assignee.assignee = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssigneeByWorkItemIDAssignee */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, assignee)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, assignee}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assignee/WorkItemAssigneeByWorkItemIDAssignee/db.Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	eswia, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}

	return &eswia, nil
}

// ExtraSchemaWorkItemAssigneesByWorkItemID retrieves a row from 'extra_schema.work_item_assignee' as a ExtraSchemaWorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func ExtraSchemaWorkItemAssigneesByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAssigneeSelectConfigOption) ([]ExtraSchemaWorkItemAssignee, error) {
	c := &ExtraSchemaWorkItemAssigneeSelectConfig{joins: ExtraSchemaWorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL)
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
	 FROM extra_schema.work_item_assignee %s 
	 WHERE work_item_assignee.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssigneesByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAssigneesByAssignee retrieves a row from 'extra_schema.work_item_assignee' as a ExtraSchemaWorkItemAssignee.
//
// Generated from index 'work_item_assignee_pkey'.
func ExtraSchemaWorkItemAssigneesByAssignee(ctx context.Context, db DB, assignee ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAssigneeSelectConfigOption) ([]ExtraSchemaWorkItemAssignee, error) {
	c := &ExtraSchemaWorkItemAssigneeSelectConfig{joins: ExtraSchemaWorkItemAssigneeJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableWorkItemsGroupBySQL)
	}

	if c.joins.Assignees {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssigneeTableAssigneesSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssigneeTableAssigneesJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssigneeTableAssigneesGroupBySQL)
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
	 FROM extra_schema.work_item_assignee %s 
	 WHERE work_item_assignee.assignee = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssigneesByAssignee */\n" + sqlstr

	// run
	// logf(sqlstr, assignee)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignee}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/Query: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignee])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignee/WorkItemAssigneeByWorkItemIDAssignee/pgx.CollectRows: %w", &XoError{Entity: "Work item assignee", Err: err}))
	}
	return res, nil
}

// FKUser_Assignee returns the User associated with the ExtraSchemaWorkItemAssignee's (Assignee).
//
// Generated from foreign key 'work_item_assignee_assignee_fkey'.
func (eswia *ExtraSchemaWorkItemAssignee) FKUser_Assignee(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, eswia.Assignee)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the ExtraSchemaWorkItemAssignee's (WorkItemID).
//
// Generated from foreign key 'work_item_assignee_work_item_id_fkey'.
func (eswia *ExtraSchemaWorkItemAssignee) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	return ExtraSchemaWorkItemByWorkItemID(ctx, db, eswia.WorkItemID)
}
