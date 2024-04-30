// Code generated by xo. DO NOT EDIT.

//lint:ignore

package got

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsTimeEntry represents a row from 'xo_tests.time_entries'.
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
type XoTestsTimeEntry struct {
	TimeEntryID XoTestsTimeEntryID `json:"timeEntryID" db:"time_entry_id" required:"true" nullable:"false"` // time_entry_id
	WorkItemID  *XoTestsWorkItemID `json:"workItemID" db:"work_item_id"`                                    // work_item_id
	Start       time.Time          `json:"start" db:"start" required:"true" nullable:"false"`               // start

	WorkItemJoin *XoTestsWorkItem `json:"-" db:"work_item_work_item_id"` // O2O work_items (generated from M2O)
}

// XoTestsTimeEntryCreateParams represents insert params for 'xo_tests.time_entries'.
type XoTestsTimeEntryCreateParams struct {
	Start      time.Time          `json:"start" required:"true" nullable:"false"` // start
	WorkItemID *XoTestsWorkItemID `json:"workItemID"`                             // work_item_id
}

// XoTestsTimeEntryParams represents common params for both insert and update of 'xo_tests.time_entries'.
type XoTestsTimeEntryParams interface {
	GetStart() *time.Time
	GetWorkItemID() *XoTestsWorkItemID
}

func (p XoTestsTimeEntryCreateParams) GetStart() *time.Time {
	x := p.Start
	return &x
}

func (p XoTestsTimeEntryUpdateParams) GetStart() *time.Time {
	return p.Start
}

func (p XoTestsTimeEntryCreateParams) GetWorkItemID() *XoTestsWorkItemID {
	return p.WorkItemID
}

func (p XoTestsTimeEntryUpdateParams) GetWorkItemID() *XoTestsWorkItemID {
	if p.WorkItemID != nil {
		return *p.WorkItemID
	}
	return nil
}

type XoTestsTimeEntryID int

// CreateXoTestsTimeEntry creates a new XoTestsTimeEntry in the database with the given params.
func CreateXoTestsTimeEntry(ctx context.Context, db DB, params *XoTestsTimeEntryCreateParams) (*XoTestsTimeEntry, error) {
	xtte := &XoTestsTimeEntry{
		Start:      params.Start,
		WorkItemID: params.WorkItemID,
	}

	return xtte.Insert(ctx, db)
}

type XoTestsTimeEntrySelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   XoTestsTimeEntryJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsTimeEntrySelectConfigOption func(*XoTestsTimeEntrySelectConfig)

// WithXoTestsTimeEntryLimit limits row selection.
func WithXoTestsTimeEntryLimit(limit int) XoTestsTimeEntrySelectConfigOption {
	return func(s *XoTestsTimeEntrySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsTimeEntryOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsTimeEntryOrderBy(rows map[string]*models.Direction) XoTestsTimeEntrySelectConfigOption {
	return func(s *XoTestsTimeEntrySelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsTimeEntry]
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

type XoTestsTimeEntryJoins struct {
	WorkItem bool `json:"workItem" required:"true" nullable:"false"` // O2O work_items
}

// WithXoTestsTimeEntryJoin joins with the given tables.
func WithXoTestsTimeEntryJoin(joins XoTestsTimeEntryJoins) XoTestsTimeEntrySelectConfigOption {
	return func(s *XoTestsTimeEntrySelectConfig) {
		s.joins = XoTestsTimeEntryJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithXoTestsTimeEntryFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsTimeEntryFilters(filters map[string][]any) XoTestsTimeEntrySelectConfigOption {
	return func(s *XoTestsTimeEntrySelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsTimeEntryHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsTimeEntryHavingClause(conditions map[string][]any) XoTestsTimeEntrySelectConfigOption {
	return func(s *XoTestsTimeEntrySelectConfig) {
		s.having = conditions
	}
}

const xoTestsTimeEntryTableWorkItemJoinSQL = `-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join xo_tests.work_items as _time_entries_work_item_id on _time_entries_work_item_id.work_item_id = time_entries.work_item_id
`

const xoTestsTimeEntryTableWorkItemSelectSQL = `(case when _time_entries_work_item_id.work_item_id is not null then row(_time_entries_work_item_id.*) end) as work_item_work_item_id`

const xoTestsTimeEntryTableWorkItemGroupBySQL = `_time_entries_work_item_id.work_item_id,
      _time_entries_work_item_id.work_item_id,
	time_entries.time_entry_id`

// XoTestsTimeEntryUpdateParams represents update params for 'xo_tests.time_entries'.
type XoTestsTimeEntryUpdateParams struct {
	Start      *time.Time          `json:"start" nullable:"false"` // start
	WorkItemID **XoTestsWorkItemID `json:"workItemID"`             // work_item_id
}

// SetUpdateParams updates xo_tests.time_entries struct fields with the specified params.
func (xtte *XoTestsTimeEntry) SetUpdateParams(params *XoTestsTimeEntryUpdateParams) {
	if params.Start != nil {
		xtte.Start = *params.Start
	}
	if params.WorkItemID != nil {
		xtte.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the XoTestsTimeEntry to the database.
func (xtte *XoTestsTimeEntry) Insert(ctx context.Context, db DB) (*XoTestsTimeEntry, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.time_entries (
	start, work_item_id
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, xtte.Start, xtte.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, xtte.Start, xtte.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Insert/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	newxtte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}

	*xtte = newxtte

	return xtte, nil
}

// Update updates a XoTestsTimeEntry in the database.
func (xtte *XoTestsTimeEntry) Update(ctx context.Context, db DB) (*XoTestsTimeEntry, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.time_entries SET 
	start = $1, work_item_id = $2 
	WHERE time_entry_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, xtte.Start, xtte.WorkItemID, xtte.TimeEntryID)

	rows, err := db.Query(ctx, sqlstr, xtte.Start, xtte.WorkItemID, xtte.TimeEntryID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Update/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	newxtte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	*xtte = newxtte

	return xtte, nil
}

// Upsert upserts a XoTestsTimeEntry in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtte *XoTestsTimeEntry) Upsert(ctx context.Context, db DB, params *XoTestsTimeEntryCreateParams) (*XoTestsTimeEntry, error) {
	var err error

	xtte.Start = params.Start
	xtte.WorkItemID = params.WorkItemID

	xtte, err = xtte.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertXoTestsTimeEntry/Insert: %w", &XoError{Entity: "Time entry", Err: err})
			}
			xtte, err = xtte.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertXoTestsTimeEntry/Update: %w", &XoError{Entity: "Time entry", Err: err})
			}
		}
	}

	return xtte, err
}

// Delete deletes the XoTestsTimeEntry from the database.
func (xtte *XoTestsTimeEntry) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.time_entries 
	WHERE time_entry_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtte.TimeEntryID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsTimeEntryPaginated returns a cursor-paginated list of XoTestsTimeEntry.
// At least one cursor is required.
func XoTestsTimeEntryPaginated(ctx context.Context, db DB, cursor models.PaginationCursor, opts ...XoTestsTimeEntrySelectConfigOption) ([]XoTestsTimeEntry, error) {
	c := &XoTestsTimeEntrySelectConfig{
		joins:   XoTestsTimeEntryJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsTimeEntry][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Paginated/cursor: %w", &XoError{Entity: "Time entry", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == models.DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("time_entries.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Paginated/orderBy: %w", &XoError{Entity: "Time entry", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, xoTestsTimeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsTimeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsTimeEntryTableWorkItemGroupBySQL)
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
	time_entries.start,
	time_entries.time_entry_id,
	time_entries.work_item_id %s 
	 FROM xo_tests.time_entries %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsTimeEntryPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Paginated/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsTimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTimeEntry/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	return res, nil
}

// XoTestsTimeEntryByTimeEntryID retrieves a row from 'xo_tests.time_entries' as a XoTestsTimeEntry.
//
// Generated from index 'time_entries_pkey'.
func XoTestsTimeEntryByTimeEntryID(ctx context.Context, db DB, timeEntryID XoTestsTimeEntryID, opts ...XoTestsTimeEntrySelectConfigOption) (*XoTestsTimeEntry, error) {
	c := &XoTestsTimeEntrySelectConfig{joins: XoTestsTimeEntryJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, xoTestsTimeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsTimeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsTimeEntryTableWorkItemGroupBySQL)
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
	time_entries.start,
	time_entries.time_entry_id,
	time_entries.work_item_id %s 
	 FROM xo_tests.time_entries %s 
	 WHERE time_entries.time_entry_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsTimeEntryByTimeEntryID */\n" + sqlstr

	// run
	// logf(sqlstr, timeEntryID)
	rows, err := db.Query(ctx, sqlstr, append([]any{timeEntryID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	xtte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}

	return &xtte, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the XoTestsTimeEntry's (WorkItemID).
//
// Generated from foreign key 'time_entries_work_item_id_fkey'.
func (xtte *XoTestsTimeEntry) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	return XoTestsWorkItemByWorkItemID(ctx, db, *xtte.WorkItemID)
}
