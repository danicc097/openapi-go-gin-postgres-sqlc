
// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"github.com/lib/pq"
	"github.com/lib/pq/hstore"

	"github.com/google/uuid"

)


// ExtraSchemaWorkItem represents a row from 'extra_schema.work_items'.
// Change properties via SQL column comments, joined with " && ":
//     - "properties":<p1>,<p2>,...
//         -- private: exclude a field from JSON.
//         -- not-required: make a schema field not required.
//         -- hidden: exclude field from OpenAPI generation.
//         -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//            i.e. no joins will be generated.
//         -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//     - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//     - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//     - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaWorkItem struct {
	WorkItemID ExtraSchemaWorkItemID `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"` // work_item_id
	Title *string `json:"title" db:"title"` // title
	Description *string `json:"description" db:"description"` // description

	DemoWorkItemJoin *ExtraSchemaDemoWorkItem `json:"-" db:"demo_work_item_work_item_id"` // O2O demo_work_items (inferred)
	AdminsJoin *[]ExtraSchemaUser `json:"-" db:"work_item_admin_admins"` // M2M work_item_admin
	AssigneesJoin *[]ExtraSchemaWorkItemM2MAssigneeWIA `json:"-" db:"work_item_assignee_assignees"` // M2M work_item_assignee

}



// ExtraSchemaWorkItemCreateParams represents insert params for 'extra_schema.work_items'.
type ExtraSchemaWorkItemCreateParams struct {
	Description *string `json:"description"` // description
	Title *string `json:"title"` // title
}

// ExtraSchemaWorkItemParams represents common params for both insert and update of 'extra_schema.work_items'.
type ExtraSchemaWorkItemParams interface {
GetDescription() *string 
GetTitle() *string 
}


			func (p ExtraSchemaWorkItemCreateParams) GetDescription() *string {
				return p.Description
			}
			func (p ExtraSchemaWorkItemUpdateParams) GetDescription() *string {
				if p.Description != nil {
					return *p.Description
				}
				return nil
			}
			

			func (p ExtraSchemaWorkItemCreateParams) GetTitle() *string {
				return p.Title
			}
			func (p ExtraSchemaWorkItemUpdateParams) GetTitle() *string {
				if p.Title != nil {
					return *p.Title
				}
				return nil
			}
			

type ExtraSchemaWorkItemID int

// CreateExtraSchemaWorkItem creates a new ExtraSchemaWorkItem in the database with the given params.
func CreateExtraSchemaWorkItem(ctx context.Context, db DB, params *ExtraSchemaWorkItemCreateParams) (*ExtraSchemaWorkItem, error) {
  eswi := &ExtraSchemaWorkItem{
	Description: params.Description,
	Title: params.Title,
}

  return eswi.Insert(ctx, db)
}

	type ExtraSchemaWorkItemSelectConfig struct {
		limit       string
		orderBy     map[string]Direction
		joins       ExtraSchemaWorkItemJoins
		filters     map[string][]any
		having     map[string][]any
		
	}
	type ExtraSchemaWorkItemSelectConfigOption func(*ExtraSchemaWorkItemSelectConfig)

	// WithExtraSchemaWorkItemLimit limits row selection.
	func WithExtraSchemaWorkItemLimit(limit int) ExtraSchemaWorkItemSelectConfigOption {
		return func(s *ExtraSchemaWorkItemSelectConfig) {
			if limit > 0 {
				s.limit = fmt.Sprintf(" limit %d ", limit)
			}
		}
	}
// WithExtraSchemaWorkItemOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaWorkItemOrderBy(rows map[string]*Direction) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaWorkItem]
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
	type ExtraSchemaWorkItemJoins struct {
DemoWorkItem bool `json:"demoWorkItem" required:"true" nullable:"false"` // O2O demo_work_items
Admins bool `json:"admins" required:"true" nullable:"false"` // M2M work_item_admin
Assignees bool `json:"assignees" required:"true" nullable:"false"` // M2M work_item_assignee
}

	// WithExtraSchemaWorkItemJoin joins with the given tables.
func WithExtraSchemaWorkItemJoin(joins ExtraSchemaWorkItemJoins) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.joins = ExtraSchemaWorkItemJoins{
			DemoWorkItem:  s.joins.DemoWorkItem || joins.DemoWorkItem,
		Admins:  s.joins.Admins || joins.Admins,
		Assignees:  s.joins.Assignees || joins.Assignees,

		}
	}
}
// ExtraSchemaWorkItemM2MAssigneeWIA represents a M2M join against "extra_schema.work_item_assignee"
type ExtraSchemaWorkItemM2MAssigneeWIA struct {
	User ExtraSchemaUser `json:"user" db:"users" required:"true"`
	Role *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}
	
// WithExtraSchemaWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//filters := map[string][]any{
//	"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//	`(col.created_at > $i OR 
//	col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//}
func WithExtraSchemaWorkItemFilters(filters map[string][]any) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaWorkItemHavingClause(conditions map[string][]any) ExtraSchemaWorkItemSelectConfigOption {
	return func(s *ExtraSchemaWorkItemSelectConfig) {
		s.having = conditions
	}
}
	const extraSchemaWorkItemTableDemoWorkItemJoinSQL = `-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join extra_schema.demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableDemoWorkItemSelectSQL = `(case when _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id`

const extraSchemaWorkItemTableDemoWorkItemGroupBySQL = `_demo_work_items_work_item_id.work_item_id,
	work_items.work_item_id`

const extraSchemaWorkItemTableAdminsJoinSQL = `-- M2M join generated from "work_item_admin_admin_fkey"
left join (
	select
		work_item_admin.work_item_id as work_item_admin_work_item_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.work_item_admin
	join extra_schema.users on users.user_id = work_item_admin.admin
	group by
		work_item_admin_work_item_id
		, users.user_id
) as xo_join_work_item_admin_admins on xo_join_work_item_admin_admins.work_item_admin_work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableAdminsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_admin_admins.__users
		)) filter (where xo_join_work_item_admin_admins.__users_user_id is not null), '{}') as work_item_admin_admins`

const extraSchemaWorkItemTableAdminsGroupBySQL = `work_items.work_item_id, work_items.work_item_id`

const extraSchemaWorkItemTableAssigneesJoinSQL = `-- M2M join generated from "work_item_assignee_assignee_fkey"
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
) as xo_join_work_item_assignee_assignees on xo_join_work_item_assignee_assignees.work_item_assignee_work_item_id = work_items.work_item_id
`

const extraSchemaWorkItemTableAssigneesSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_assignees.__users
		, xo_join_work_item_assignee_assignees.role
		)) filter (where xo_join_work_item_assignee_assignees.__users_user_id is not null), '{}') as work_item_assignee_assignees`

const extraSchemaWorkItemTableAssigneesGroupBySQL = `work_items.work_item_id, work_items.work_item_id`





// ExtraSchemaWorkItemUpdateParams represents update params for 'extra_schema.work_items'.
type ExtraSchemaWorkItemUpdateParams struct {
	Description **string `json:"description"` // description
	Title **string `json:"title"` // title
}

// SetUpdateParams updates extra_schema.work_items struct fields with the specified params.
func (eswi *ExtraSchemaWorkItem) SetUpdateParams(params *ExtraSchemaWorkItemUpdateParams) {
if params.Description != nil {
	eswi.Description = *params.Description
}
if params.Title != nil {
	eswi.Title = *params.Title
}
}

// Insert inserts the ExtraSchemaWorkItem to the database.
func (eswi *ExtraSchemaWorkItem) Insert(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.work_items (
	description, title
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, eswi.Description, eswi.Title)

	rows, err := db.Query(ctx, sqlstr, eswi.Description, eswi.Title)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Insert/db.Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	neweswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err }))
	}

  *eswi = neweswi

	return eswi, nil
}


// Update updates a ExtraSchemaWorkItem in the database.
func (eswi *ExtraSchemaWorkItem) Update(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error)  {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.work_items SET 
	description = $1, title = $2 
	WHERE work_item_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, eswi.Description, eswi.Title, eswi.WorkItemID)

  rows, err := db.Query(ctx, sqlstr, eswi.Description, eswi.Title, eswi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Update/db.Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	neweswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err }))
	}
  *eswi = neweswi

	return eswi, nil
}


// Upsert upserts a ExtraSchemaWorkItem in the database.
// Requires appropriate PK(s) to be set beforehand.
func (eswi *ExtraSchemaWorkItem) Upsert(ctx context.Context, db DB, params *ExtraSchemaWorkItemCreateParams) (*ExtraSchemaWorkItem, error)  {
	var err error

  	eswi.Description = params.Description
	eswi.Title = params.Title


  eswi, err = eswi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
			  return nil, fmt.Errorf("UpsertExtraSchemaWorkItem/Insert: %w", &XoError{Entity: "Work item", Err: err })
			}
		  eswi, err = eswi.Update(ctx, db)
      if err != nil {
			  return nil, fmt.Errorf("UpsertExtraSchemaWorkItem/Update: %w", &XoError{Entity: "Work item", Err: err })
      }
		}
	}

  return eswi, err
}

// Delete deletes the ExtraSchemaWorkItem from the database.
func (eswi *ExtraSchemaWorkItem) Delete(ctx context.Context, db DB) (error) {
// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.work_items 
	WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, eswi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}





// ExtraSchemaWorkItemPaginated returns a cursor-paginated list of ExtraSchemaWorkItem.
// At least one cursor is required.
func ExtraSchemaWorkItemPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{},
		filters: make(map[string][]any), 
		having: make(map[string][]any),
		orderBy: make(map[string]Direction),
}

	for _, o := range opts {
		o(c)
	}

  if cursor.Value == nil {
    
    return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
  }
  field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaWorkItem][cursor.Column]
  if !ok {
    return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/cursor: %w", &XoError{Entity: "Work item", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
  }

  op := "<"
  if cursor.Direction == DirectionAsc {
    op = ">"
  }
  c.filters[fmt.Sprintf("work_items.%s %s $i", field.Db, op)] = []any{*cursor.Value}
  c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

  paramStart := 0 // all filters will come from the user
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
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
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/orderBy: %w", &XoError{Entity: "Work item", Err: fmt.Errorf("at least one sorted column is required")}))
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
		
			if c.joins.DemoWorkItem {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
			}
			
			if c.joins.Admins {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAdminsSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAdminsJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAdminsGroupBySQL)
			}
			
			if c.joins.Assignees {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAssigneesSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAssigneesJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAssigneesGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM extra_schema.work_items %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
  sqlstr = "/* ExtraSchemaWorkItemPaginated */\n"+sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err }))
	}
	return res, nil
}


// ExtraSchemaWorkItems retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index '[xo] base filter query'.
func ExtraSchemaWorkItems(ctx context.Context, db DB, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{},filters: make(map[string][]any), having: make(map[string][]any),
}

	for _, o := range opts {
		o(c)
	}

  paramStart := 0
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND "+strings.Join(filterClauses, " AND ")+" "
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
		
			if c.joins.DemoWorkItem {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
			}
			
			if c.joins.Admins {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAdminsSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAdminsJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAdminsGroupBySQL)
			}
			
			if c.joins.Assignees {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAssigneesSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAssigneesJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAssigneesGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM extra_schema.work_items %s 
	 WHERE true
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
  sqlstr = "/* ExtraSchemaWorkItems */\n"+sqlstr

	// run
	// logf(sqlstr, )
	rows, err := db.Query(ctx, sqlstr, append([]any{}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByDescription/Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	defer rows.Close()
	// process
  
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err }))
	}
	return res, nil
}


// ExtraSchemaWorkItemByWorkItemID retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index 'work_items_pkey'.
func ExtraSchemaWorkItemByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemSelectConfigOption) (*ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{},filters: make(map[string][]any), having: make(map[string][]any),
}

	for _, o := range opts {
		o(c)
	}

  paramStart := 1
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND "+strings.Join(filterClauses, " AND ")+" "
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
		
			if c.joins.DemoWorkItem {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
			}
			
			if c.joins.Admins {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAdminsSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAdminsJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAdminsGroupBySQL)
			}
			
			if c.joins.Assignees {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAssigneesSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAssigneesJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAssigneesGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM extra_schema.work_items %s 
	 WHERE work_items.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
  sqlstr = "/* ExtraSchemaWorkItemByWorkItemID */\n"+sqlstr

	// run
	// logf(sqlstr, workItemID)
  rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	eswi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_items/WorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item", Err: err }))
	}
	

	return &eswi, nil
}


// ExtraSchemaWorkItemsByTitle retrieves a row from 'extra_schema.work_items' as a ExtraSchemaWorkItem.
//
// Generated from index 'work_items_title_description_idx1'.
func ExtraSchemaWorkItemsByTitle(ctx context.Context, db DB, title *string, opts ...ExtraSchemaWorkItemSelectConfigOption) ([]ExtraSchemaWorkItem, error) {
	c := &ExtraSchemaWorkItemSelectConfig{joins: ExtraSchemaWorkItemJoins{},filters: make(map[string][]any), having: make(map[string][]any),
}

	for _, o := range opts {
		o(c)
	}

  paramStart := 1
	nth := func ()  string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i"){
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND "+strings.Join(filterClauses, " AND ")+" "
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
		
			if c.joins.DemoWorkItem {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableDemoWorkItemSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableDemoWorkItemJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableDemoWorkItemGroupBySQL)
			}
			
			if c.joins.Admins {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAdminsSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAdminsJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAdminsGroupBySQL)
			}
			
			if c.joins.Assignees {
				selectClauses = append(selectClauses, extraSchemaWorkItemTableAssigneesSelectSQL)
				joinClauses = append(joinClauses, extraSchemaWorkItemTableAssigneesJoinSQL)
				groupByClauses = append(groupByClauses, extraSchemaWorkItemTableAssigneesGroupBySQL)
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
	work_items.description,
	work_items.title,
	work_items.work_item_id %s 
	 FROM extra_schema.work_items %s 
	 WHERE work_items.title = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
  sqlstr = "/* ExtraSchemaWorkItemsByTitle */\n"+sqlstr

	// run
	// logf(sqlstr, title)
	rows, err := db.Query(ctx, sqlstr, append([]any{title}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByTitleDescription/Query: %w", &XoError{Entity: "Work item", Err: err }))
	}
	defer rows.Close()
	// process
  
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItem/WorkItemsByTitleDescription/pgx.CollectRows: %w", &XoError{Entity: "Work item", Err: err }))
	}
	return res, nil
}

