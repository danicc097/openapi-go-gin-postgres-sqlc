package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// DemoTwoWorkItem represents a row from 'public.demo_two_work_items'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type DemoTwoWorkItem struct {
	WorkItemID            int64      `json:"workItemID" db:"work_item_id" required:"true"`                         // work_item_id
	CustomDateForProject2 *time.Time `json:"customDateForProject2" db:"custom_date_for_project_2" required:"true"` // custom_date_for_project_2

	WorkItemJoin *WorkItem `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"` // O2O work_items (inferred)

}

// DemoTwoWorkItemCreateParams represents insert params for 'public.demo_two_work_items'.
type DemoTwoWorkItemCreateParams struct {
	WorkItemID            int64      `json:"workItemID" required:"true"`            // work_item_id
	CustomDateForProject2 *time.Time `json:"customDateForProject2" required:"true"` // custom_date_for_project_2
}

// CreateDemoTwoWorkItem creates a new DemoTwoWorkItem in the database with the given params.
func CreateDemoTwoWorkItem(ctx context.Context, db DB, params *DemoTwoWorkItemCreateParams) (*DemoTwoWorkItem, error) {
	dtwi := &DemoTwoWorkItem{
		WorkItemID:            params.WorkItemID,
		CustomDateForProject2: params.CustomDateForProject2,
	}

	return dtwi.Insert(ctx, db)
}

// DemoTwoWorkItemUpdateParams represents update params for 'public.demo_two_work_items'
type DemoTwoWorkItemUpdateParams struct {
	CustomDateForProject2 **time.Time `json:"customDateForProject2" required:"true"` // custom_date_for_project_2
}

// SetUpdateParams updates public.demo_two_work_items struct fields with the specified params.
func (dtwi *DemoTwoWorkItem) SetUpdateParams(params *DemoTwoWorkItemUpdateParams) {
	if params.CustomDateForProject2 != nil {
		dtwi.CustomDateForProject2 = *params.CustomDateForProject2
	}
}

type DemoTwoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   DemoTwoWorkItemJoins
	filters map[string][]any
}
type DemoTwoWorkItemSelectConfigOption func(*DemoTwoWorkItemSelectConfig)

// WithDemoTwoWorkItemLimit limits row selection.
func WithDemoTwoWorkItemLimit(limit int) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type DemoTwoWorkItemOrderBy = string

const (
	DemoTwoWorkItemCustomDateForProject2DescNullsFirst DemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS FIRST "
	DemoTwoWorkItemCustomDateForProject2DescNullsLast  DemoTwoWorkItemOrderBy = " custom_date_for_project_2 DESC NULLS LAST "
	DemoTwoWorkItemCustomDateForProject2AscNullsFirst  DemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS FIRST "
	DemoTwoWorkItemCustomDateForProject2AscNullsLast   DemoTwoWorkItemOrderBy = " custom_date_for_project_2 ASC NULLS LAST "
)

// WithDemoTwoWorkItemOrderBy orders results by the given columns.
func WithDemoTwoWorkItemOrderBy(rows ...DemoTwoWorkItemOrderBy) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type DemoTwoWorkItemJoins struct {
	WorkItem bool // O2O work_items
}

// WithDemoTwoWorkItemJoin joins with the given tables.
func WithDemoTwoWorkItemJoin(joins DemoTwoWorkItemJoins) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		s.joins = DemoTwoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithDemoTwoWorkItemFilters adds the given filters, which may be parameterized.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`col.created_at > $i AND
//		col.created_at < $i`: {time.Now().Add(-24 * time.Hour), time.Now().Add(24 * time.Hour)},
//	}
func WithDemoTwoWorkItemFilters(filters map[string][]any) DemoTwoWorkItemSelectConfigOption {
	return func(s *DemoTwoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the DemoTwoWorkItem to the database.
func (dtwi *DemoTwoWorkItem) Insert(ctx context.Context, db DB) (*DemoTwoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.demo_two_work_items (` +
		`work_item_id, custom_date_for_project_2` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, dtwi.WorkItemID, dtwi.CustomDateForProject2)
	rows, err := db.Query(ctx, sqlstr, dtwi.WorkItemID, dtwi.CustomDateForProject2)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Insert/db.Query: %w", err))
	}
	newdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Insert/pgx.CollectOneRow: %w", err))
	}
	*dtwi = newdtwi

	return dtwi, nil
}

// Update updates a DemoTwoWorkItem in the database.
func (dtwi *DemoTwoWorkItem) Update(ctx context.Context, db DB) (*DemoTwoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.demo_two_work_items SET ` +
		`custom_date_for_project_2 = $1 ` +
		`WHERE work_item_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dtwi.CustomDateForProject2, dtwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Update/db.Query: %w", err))
	}
	newdtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Update/pgx.CollectOneRow: %w", err))
	}
	*dtwi = newdtwi

	return dtwi, nil
}

// Upsert upserts a DemoTwoWorkItem in the database.
// Requires appropiate PK(s) to be set beforehand.
func (dtwi *DemoTwoWorkItem) Upsert(ctx context.Context, db DB, params *DemoTwoWorkItemCreateParams) (*DemoTwoWorkItem, error) {
	var err error

	dtwi.WorkItemID = params.WorkItemID
	dtwi.CustomDateForProject2 = params.CustomDateForProject2

	dtwi, err = dtwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			dtwi, err = dtwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return dtwi, err
}

// Delete deletes the DemoTwoWorkItem from the database.
func (dtwi *DemoTwoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_two_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dtwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// DemoTwoWorkItemPaginatedByWorkItemIDAsc returns a cursor-paginated list of DemoTwoWorkItem in Asc order.
func DemoTwoWorkItemPaginatedByWorkItemIDAsc(ctx context.Context, db DB, workItemID int64, opts ...DemoTwoWorkItemSelectConfigOption) ([]DemoTwoWorkItem, error) {
	c := &DemoTwoWorkItemSelectConfig{joins: DemoTwoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`demo_two_work_items.work_item_id,
demo_two_work_items.custom_date_for_project_2,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id `+
		`FROM public.demo_two_work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_two_work_items.work_item_id`+
		` WHERE demo_two_work_items.work_item_id > $2`+
		` %s  GROUP BY demo_two_work_items.work_item_id, 
demo_two_work_items.custom_date_for_project_2, 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_two_work_items.work_item_id ORDER BY 
		work_item_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// DemoTwoWorkItemPaginatedByWorkItemIDDesc returns a cursor-paginated list of DemoTwoWorkItem in Desc order.
func DemoTwoWorkItemPaginatedByWorkItemIDDesc(ctx context.Context, db DB, workItemID int64, opts ...DemoTwoWorkItemSelectConfigOption) ([]DemoTwoWorkItem, error) {
	c := &DemoTwoWorkItemSelectConfig{joins: DemoTwoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`demo_two_work_items.work_item_id,
demo_two_work_items.custom_date_for_project_2,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id `+
		`FROM public.demo_two_work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_two_work_items.work_item_id`+
		` WHERE demo_two_work_items.work_item_id < $2`+
		` %s  GROUP BY demo_two_work_items.work_item_id, 
demo_two_work_items.custom_date_for_project_2, 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_two_work_items.work_item_id ORDER BY 
		work_item_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoTwoWorkItem/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// DemoTwoWorkItemByWorkItemID retrieves a row from 'public.demo_two_work_items' as a DemoTwoWorkItem.
//
// Generated from index 'demo_two_work_items_pkey'.
func DemoTwoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...DemoTwoWorkItemSelectConfigOption) (*DemoTwoWorkItem, error) {
	c := &DemoTwoWorkItemSelectConfig{joins: DemoTwoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	filters := ""

	sqlstr := fmt.Sprintf(`SELECT `+
		`demo_two_work_items.work_item_id,
demo_two_work_items.custom_date_for_project_2,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id `+
		`FROM public.demo_two_work_items `+
		`-- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_two_work_items.work_item_id`+
		` WHERE demo_two_work_items.work_item_id = $2`+
		` %s  GROUP BY 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_two_work_items.work_item_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_two_work_items/DemoTwoWorkItemByWorkItemID/db.Query: %w", err))
	}
	dtwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoTwoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_two_work_items/DemoTwoWorkItemByWorkItemID/pgx.CollectOneRow: %w", err))
	}

	return &dtwi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoTwoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_two_work_items_work_item_id_fkey'.
func (dtwi *DemoTwoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dtwi.WorkItemID)
}
