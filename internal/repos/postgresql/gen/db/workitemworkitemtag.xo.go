package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// WorkItemWorkItemTag represents a row from 'public.work_item_work_item_tag'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItemWorkItemTag struct {
	WorkItemTagID int   `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID" db:"work_item_id" required:"true"`        // work_item_id

	WorkItemWorkItemTagsJoin *[]WorkItemTag `json:"-" db:"work_item_work_item_tag_work_item_tags" openapi-go:"ignore"` // M2M work_item_work_item_tag
	WorkItemTagWorkItemsJoin *[]WorkItem    `json:"-" db:"work_item_work_item_tag_work_items" openapi-go:"ignore"`     // M2M work_item_work_item_tag

}

// WorkItemWorkItemTagCreateParams represents insert params for 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagCreateParams struct {
	WorkItemTagID int   `json:"workItemTagID" required:"true"` // work_item_tag_id
	WorkItemID    int64 `json:"workItemID" required:"true"`    // work_item_id
}

// CreateWorkItemWorkItemTag creates a new WorkItemWorkItemTag in the database with the given params.
func CreateWorkItemWorkItemTag(ctx context.Context, db DB, params *WorkItemWorkItemTagCreateParams) (*WorkItemWorkItemTag, error) {
	wiwit := &WorkItemWorkItemTag{
		WorkItemTagID: params.WorkItemTagID,
		WorkItemID:    params.WorkItemID,
	}

	return wiwit.Insert(ctx, db)
}

// WorkItemWorkItemTagUpdateParams represents update params for 'public.work_item_work_item_tag'.
type WorkItemWorkItemTagUpdateParams struct {
	WorkItemTagID *int   `json:"workItemTagID" required:"true"` // work_item_tag_id
	WorkItemID    *int64 `json:"workItemID" required:"true"`    // work_item_id
}

// SetUpdateParams updates public.work_item_work_item_tag struct fields with the specified params.
func (wiwit *WorkItemWorkItemTag) SetUpdateParams(params *WorkItemWorkItemTagUpdateParams) {
	if params.WorkItemTagID != nil {
		wiwit.WorkItemTagID = *params.WorkItemTagID
	}
	if params.WorkItemID != nil {
		wiwit.WorkItemID = *params.WorkItemID
	}
}

type WorkItemWorkItemTagSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemWorkItemTagJoins
	filters map[string][]any
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
	WorkItemTags         bool // M2M work_item_work_item_tag
	WorkItemsWorkItemTag bool // M2M work_item_work_item_tag
}

// WithWorkItemWorkItemTagJoin joins with the given tables.
func WithWorkItemWorkItemTagJoin(joins WorkItemWorkItemTagJoins) WorkItemWorkItemTagSelectConfigOption {
	return func(s *WorkItemWorkItemTagSelectConfig) {
		s.joins = WorkItemWorkItemTagJoins{
			WorkItemTags:         s.joins.WorkItemTags || joins.WorkItemTags,
			WorkItemsWorkItemTag: s.joins.WorkItemsWorkItemTag || joins.WorkItemsWorkItemTag,
		}
	}
}

// WithWorkItemWorkItemTagFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
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
) as joined_work_item_work_item_tag_work_item_tags on joined_work_item_work_item_tag_work_item_tags.work_item_work_item_tag_work_item_id = work_item_work_item_tag.work_item_tag_id
`

const workItemWorkItemTagTableWorkItemTagsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_item_tags.__work_item_tags
		)) filter (where joined_work_item_work_item_tag_work_item_tags.__work_item_tags_work_item_tag_id is not null), '{}') as work_item_work_item_tag_work_item_tags`

const workItemWorkItemTagTableWorkItemTagsGroupBySQL = `work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_id`

const workItemWorkItemTagTableWorkItemsWorkItemTagJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
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
) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_work_item_tag.work_item_id
`

const workItemWorkItemTagTableWorkItemsWorkItemTagSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items_work_item_id is not null), '{}') as work_item_work_item_tag_work_items`

const workItemWorkItemTagTableWorkItemsWorkItemTagGroupBySQL = `work_item_work_item_tag.work_item_id, work_item_work_item_tag.work_item_tag_id, work_item_work_item_tag.work_item_id`

// Insert inserts the WorkItemWorkItemTag to the database.
func (wiwit *WorkItemWorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemWorkItemTag, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_work_item_tag (` +
		`work_item_tag_id, work_item_id` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/db.Query: %w", err))
	}
	newwiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Insert/pgx.CollectOneRow: %w", err))
	}
	*wiwit = newwiwit

	return wiwit, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the WorkItemWorkItemTag from the database.
func (wiwit *WorkItemWorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_work_item_tag ` +
		`WHERE work_item_tag_id = $1 AND work_item_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wiwit.WorkItemTagID, wiwit.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemIDAsc returns a cursor-paginated list of WorkItemWorkItemTag in Asc order.
func WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemIDAsc(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, ",\n") + " "
	}
	joins := ""
	if len(joinClauses) > 0 {
		joins = ", " + strings.Join(joinClauses, ",\n") + " "
	}
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = ", " + strings.Join(groupByClauses, ",\n") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id %s `+
		`FROM public.work_item_work_item_tag %s `+
		` WHERE work_item_work_item_tag.work_item_tag_id > $1 AND work_item_work_item_tag.work_item_id > $2`+
		` %s  GROUP BY work_item_work_item_tag.work_item_tag_id, 
work_item_work_item_tag.work_item_id 
 %s 
 ORDER BY 
		work_item_tag_id Asc ,
		work_item_id Asc `, filters, selects, joins, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemIDDesc returns a cursor-paginated list of WorkItemWorkItemTag in Desc order.
func WorkItemWorkItemTagPaginatedByWorkItemTagIDWorkItemIDDesc(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	if c.joins.WorkItemTags {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemTagsSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemTagsJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemTagsGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemWorkItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemWorkItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemWorkItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, ",\n") + " "
	}
	joins := ""
	if len(joinClauses) > 0 {
		joins = ", " + strings.Join(joinClauses, ",\n") + " "
	}
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = ", " + strings.Join(groupByClauses, ",\n") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id %s `+
		`FROM public.work_item_work_item_tag %s `+
		` WHERE work_item_work_item_tag.work_item_tag_id < $1 AND work_item_work_item_tag.work_item_id < $2`+
		` %s  GROUP BY work_item_work_item_tag.work_item_tag_id, 
work_item_work_item_tag.work_item_id 
 %s 
 ORDER BY 
		work_item_tag_id Desc ,
		work_item_id Desc `, filters, selects, joins, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagByWorkItemIDWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagByWorkItemIDWorkItemTagID(ctx context.Context, db DB, workItemID int64, workItemTagID int, opts ...WorkItemWorkItemTagSelectConfigOption) (*WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id `+
		`FROM public.work_item_work_item_tag `+
		``+
		` WHERE work_item_work_item_tag.work_item_id = $1 AND work_item_work_item_tag.work_item_tag_id = $2`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, workItemTagID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/db.Query: %w", err))
	}
	wiwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_work_item_tag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectOneRow: %w", err))
	}

	return &wiwit, nil
}

// WorkItemWorkItemTagsByWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id `+
		`FROM public.work_item_work_item_tag `+
		``+
		` WHERE work_item_work_item_tag.work_item_id = $1`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_pkey'.
func WorkItemWorkItemTagsByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id `+
		`FROM public.work_item_work_item_tag `+
		``+
		` WHERE work_item_work_item_tag.work_item_tag_id = $1`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemIDWorkItemTagID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemWorkItemTagsByWorkItemTagIDWorkItemID retrieves a row from 'public.work_item_work_item_tag' as a WorkItemWorkItemTag.
//
// Generated from index 'work_item_work_item_tag_work_item_tag_id_work_item_id_idx'.
func WorkItemWorkItemTagsByWorkItemTagIDWorkItemID(ctx context.Context, db DB, workItemTagID int, workItemID int64, opts ...WorkItemWorkItemTagSelectConfigOption) ([]WorkItemWorkItemTag, error) {
	c := &WorkItemWorkItemTagSelectConfig{joins: WorkItemWorkItemTagJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`work_item_work_item_tag.work_item_tag_id,
work_item_work_item_tag.work_item_id `+
		`FROM public.work_item_work_item_tag `+
		``+
		` WHERE work_item_work_item_tag.work_item_tag_id = $1 AND work_item_work_item_tag.work_item_id = $2`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemTagID, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID, workItemID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemWorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemWorkItemTag/WorkItemWorkItemTagByWorkItemTagIDWorkItemID/pgx.CollectRows: %w", err))
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
