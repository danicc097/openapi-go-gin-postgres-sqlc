package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgx/v5"
)

// WorkItemWorkItemTag represents a row from 'public.work_item_work_item_tag'.
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
type WorkItemWorkItemTag struct {
	WorkItemTagID WorkItemTagID `json:"workItemTagID" db:"work_item_tag_id" required:"true" nullable:"false"` // work_item_tag_id
	WorkItemID    WorkItemID    `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`        // work_item_id

	WorkItemTagsJoin *[]WorkItemTag `json:"-" db:"work_item_work_item_tag_work_item_tags" openapi-go:"ignore"` // M2M work_item_work_item_tag
	WorkItemsJoin    *[]WorkItem    `json:"-" db:"work_item_work_item_tag_work_items" openapi-go:"ignore"`     // M2M work_item_work_item_tag

}

// WorkItemWorkItemTagCreateParams represents insert params for 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagCreateParams struct {
	WorkItemID    WorkItemID    `json:"workItemID" required:"true" nullable:"false"`    // work_item_id
	WorkItemTagID WorkItemTagID `json:"workItemTagID" required:"true" nullable:"false"` // work_item_tag_id
}

// WorkItemWorkItemTagParams represents common params for both insert and update of 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagParams interface {
	GetWorkItemID() *WorkItemID
	GetWorkItemTagID() *WorkItemTagID
}

func (p WorkItemWorkItemTagCreateParams) GetWorkItemID() *WorkItemID {
	x := p.WorkItemID
	return &x
}
func (p WorkItemWorkItemTagUpdateParams) GetWorkItemID() *WorkItemID {
	return p.WorkItemID
}

func (p WorkItemWorkItemTagCreateParams) GetWorkItemTagID() *WorkItemTagID {
	x := p.WorkItemTagID
	return &x
}
func (p WorkItemWorkItemTagUpdateParams) GetWorkItemTagID() *WorkItemTagID {
	return p.WorkItemTagID
}

// CreateWorkItemWorkItemTag creates a new WorkItemWorkItemTag in the database with the given params.
func CreateWorkItemWorkItemTag(ctx context.Context, db DB, params *WorkItemWorkItemTagCreateParams) (*WorkItemWorkItemTag, error) {
	wiwit := &WorkItemWorkItemTag{
		WorkItemID:    params.WorkItemID,
		WorkItemTagID: params.WorkItemTagID,
	}

	return wiwit.Insert(ctx, db)
}

type WorkItemWorkItemTagSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   WorkItemWorkItemTagJoins
	filters map[string][]any
	having  map[string][]any
}
type WorkItemWorkItemTagSelectConfigOption func(*WorkItemWorkItemTagSelectConfig)

// WithWorkItemWorkItemTagLimit limits row selection.
func WithWorkItemWorkItemTagLimit(limit int) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemWorkItemTagOrderBy string

const ()

type WorkItemWorkItemTagJoins struct {
	WorkItemTags bool `json:"workItemTags" required:"true" nullable:"false"` // M2M work_item_work_item_tag
	WorkItems    bool `json:"workItems" required:"true" nullable:"false"`    // M2M work_item_work_item_tag
}

// WithWorkItemWorkItemTagJoin joins with the given tables.
func WithWorkItemWorkItemTagJoin(joins WorkItemWorkItemTagJoins) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.joins = WorkItemWorkItemTagJoins{
			WorkItemTags: s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItems:    s.joins.WorkItems || joins.WorkItems,
		}
	}
}

// WithWorkItemWorkItemTagFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemWorkItemTagFilters(filters map[string][]any) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.filters = filters
	}
}

// WithWorkItemWorkItemTagHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithWorkItemWorkItemTagHavingClause(conditions map[string][]any) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.having = conditions
	}
}

const workItemWorkItemTagTableWorkItemTagsJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
	select
		work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
		, work_item_tags.work_item_tag_id as __work_item_tags_work_item_tag_id
		, row(work_item_tags.*) as __work_item_tags
	from
		work_item_work_item_tag
	join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
	group by
		work_item_work_item_tag_work_item_id
		, work_item_tags.work_item_tag_id
) as xo_join_work_item_work_item_tag_work_item_tags on xo_join_work_item_work_item_tag_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id
`

const workItemWorkItemTagTableWorkItemTagsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_work_item_tag_work_item_tags.__work_item_tags
		)) filter (where xo_join_work_item_work_item_tag_work_item_tags.__work_item_tags_work_item_tag_id is not null), '{}') as work_item_work_item_tag_work_item_tags`

const workItemWorkItemTagTableWorkItemTagsGroupBySQL = `work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_id`

const workItemWorkItemTagTableWorkItemsJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
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
) as xo_join_work_item_work_item_tag_work_items on xo_join_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
`

const workItemWorkItemTagTableWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_work_item_tag_work_items.__work_items
		)) filter (where xo_join_work_item_work_item_tag_work_items.__work_items_work_item_id is not null), '{}') as work_item_work_item_tag_work_items`

const workItemWorkItemTagTableWorkItemsGroupBySQL = `work_item_work_item_tag.work_item_id, work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_id`

// WorkItemWorkItemTagUpdateParams represents update params for 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagUpdateParams struct {
	WorkItemID    *WorkItemID    `json:"workItemID" nullable:"false"`    // work_item_id
	WorkItemTagID *WorkItemTagID `json:"workItemTagID" nullable:"false"` // work_item_tag_id
}

// SetUpdateParams updates public.work_item_work_item_tag struct fields with the specified params.
func (wiwit *WorkItemWorkItemTag) SetUpdateParams(params *WorkItemWorkItemTagUpdateParams) {
	if params.WorkItemID != nil {
		wiwit.WorkItemID = *params.WorkItemID
	}
	if params.WorkItemTagID != nil {
		wiwit.WorkItemTagID = *params.WorkItemTagID
	}
}

// Insert inserts the WorkItemWorkItemTag to the database.
func (wiwit *WorkItemWorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemWorkItemTag, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_work_item_tag (
	work_item_id, work_item_tag_id
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, wiwit.WorkItemID, wiwit.WorkItemTagID)
	rows, err := db.Query(ctx, sqlstr, wiwit.WorkItemID, wiwit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/db.Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	newwiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	*wiwit = newwiwit

	return wiwit, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key or generated fields

// Delete deletes the WorkItemWorkItemTag from the database.
func (wiwit *WorkItemWorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_work_item_tag 
	WHERE work_item_tag_id = $1 AND work_item_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemID returns a cursor-paginated list of WorkItemWorkItemTag.
func WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID WorkItemTagID, workItemID WorkItemID, direction models.Direction, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsGroupBySQL)
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

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_work_item_tag.work_item_id,
	work_item_work_item_tag.work_item_tag_id %s 
	 FROM public.work_item_work_item_tag %s 
	 WHERE work_item_work_item_tag.work_item_tag_id %s $1 AND work_item_work_item_tag.work_item_id %s $2
	 %s   %s 
  %s 
  ORDER BY 
		work_item_tag_id %s  ,
		work_item_id %s `, selects, joins, operator, operator, filters, groupbys, havingClause, direction, direction)
	sqlstr += c.limit
	sqlstr = "/* WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/db.Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemWorkItemTagByWorkItemIDWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagByWorkItemIDWorkItemTagID(ctx context.Context, db DB, workItemID WorkItemID, workItemTagID WorkItemTagID, opts ...WorkItemWorkItemTagSelectConfigOption) (*WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsGroupBySQL)
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
	work_item_work_item_tag.work_item_id,
	work_item_work_item_tag.work_item_tag_id %s 
	 FROM public.work_item_work_item_tag %s 
	 WHERE work_item_work_item_tag.work_item_id = $1 AND work_item_work_item_tag.work_item_tag_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemWorkItemTagByWorkItemIDWorkItemTagID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, workItemTagID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/db.Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	wiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}

	return &wiwit, nil
}

// WorkItemWorkItemTagsByWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemID(ctx context.Context, db DB, workItemID WorkItemID, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsGroupBySQL)
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
	work_item_work_item_tag.work_item_id,
	work_item_work_item_tag.work_item_tag_id %s 
	 FROM public.work_item_work_item_tag %s 
	 WHERE work_item_work_item_tag.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemWorkItemTagsByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemTagID(ctx context.Context, db DB, workItemTagID WorkItemTagID, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsGroupBySQL)
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
	work_item_work_item_tag.work_item_id,
	work_item_work_item_tag.work_item_tag_id %s 
	 FROM public.work_item_work_item_tag %s 
	 WHERE work_item_work_item_tag.work_item_tag_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemWorkItemTagsByWorkItemTagID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagIDWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_work_item_tag_id_work_item_id_idx'.
func WorkItemWorkItemTagsByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID WorkItemTagID, workItemID WorkItemID, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItems {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsGroupBySQL)
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
	work_item_work_item_tag.work_item_id,
	work_item_work_item_tag.work_item_tag_id %s 
	 FROM public.work_item_work_item_tag %s 
	 WHERE work_item_work_item_tag.work_item_tag_id = $1 AND work_item_work_item_tag.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemWorkItemTagsByWorkItemTagIDWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTagID, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/Query: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item work item tag", Err: err}))
	}
	return res, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemWorkItemTag's (WorkItemID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wiwit.WorkItemID)
}

// FKWorkItemTag_WorkItemTagID returns the WorkItemTag associated with the WorkItemWorkItemTag's (WorkItemTagID).
//
// Generated from foreign key 'work_item_work_item_tag_work_item_tag_id_fkey'.
func (wiwit *WorkItemWorkItemTag) FKWorkItemTag_WorkItemTagID(ctx context.Context, db DB) (*WorkItemTag, error) {
	return WorkItemTagByWorkItemTagID(ctx, db, wiwit.WorkItemTagID)
}
