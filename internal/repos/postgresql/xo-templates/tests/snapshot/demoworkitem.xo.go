package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// DemoWorkItem represents a row from 'xo_tests.demo_work_items'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type DemoWorkItem struct {
	WorkItemID int64 `json:"workItemID" db:"work_item_id" required:"true"` // work_item_id
	Checked    bool  `json:"checked" db:"checked" required:"true"`         // checked

	WorkItemJoin *WorkItem `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"` // O2O work_items (inferred)
}

// DemoWorkItemCreateParams represents insert params for 'xo_tests.demo_work_items'.
type DemoWorkItemCreateParams struct {
	WorkItemID int64 `json:"workItemID" required:"true"` // work_item_id
	Checked    bool  `json:"checked" required:"true"`    // checked
}

// CreateDemoWorkItem creates a new DemoWorkItem in the database with the given params.
func CreateDemoWorkItem(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	dwi := &DemoWorkItem{
		WorkItemID: params.WorkItemID,
		Checked:    params.Checked,
	}

	return dwi.Insert(ctx, db)
}

// DemoWorkItemUpdateParams represents update params for 'xo_tests.demo_work_items'
type DemoWorkItemUpdateParams struct {
	Checked *bool `json:"checked" required:"true"` // checked
}

// SetUpdateParams updates xo_tests.demo_work_items struct fields with the specified params.
func (dwi *DemoWorkItem) SetUpdateParams(params *DemoWorkItemUpdateParams) {
	if params.Checked != nil {
		dwi.Checked = *params.Checked
	}
}

type DemoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   DemoWorkItemJoins
	filters map[string][]any
}
type DemoWorkItemSelectConfigOption func(*DemoWorkItemSelectConfig)

// WithDemoWorkItemLimit limits row selection.
func WithDemoWorkItemLimit(limit int) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type DemoWorkItemOrderBy = string

type DemoWorkItemJoins struct {
	WorkItem bool // O2O work_items
}

// WithDemoWorkItemJoin joins with the given tables.
func WithDemoWorkItemJoin(joins DemoWorkItemJoins) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.joins = DemoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithDemoWorkItemFilters adds the given filters, which may be parameterized.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`col.created_at > $i AND
//		col.created_at < $i`: {time.Now().Add(-24 * time.Hour), time.Now().Add(24 * time.Hour)},
//	}
func WithDemoWorkItemFilters(filters map[string][]any) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the DemoWorkItem to the database.
func (dwi *DemoWorkItem) Insert(ctx context.Context, db DB) (*DemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.demo_work_items (` +
		`work_item_id, checked` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, dwi.WorkItemID, dwi.Checked)
	rows, err := db.Query(ctx, sqlstr, dwi.WorkItemID, dwi.Checked)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Insert/db.Query: %w", err))
	}
	newdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Insert/pgx.CollectOneRow: %w", err))
	}
	*dwi = newdwi

	return dwi, nil
}

// Update updates a DemoWorkItem in the database.
func (dwi *DemoWorkItem) Update(ctx context.Context, db DB) (*DemoWorkItem, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.demo_work_items SET ` +
		`checked = $1 ` +
		`WHERE work_item_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, dwi.Checked, dwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dwi.Checked, dwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Update/db.Query: %w", err))
	}
	newdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Update/pgx.CollectOneRow: %w", err))
	}
	*dwi = newdwi

	return dwi, nil
}

// Upsert upserts a DemoWorkItem in the database.
// Requires appropiate PK(s) to be set beforehand.
func (dwi *DemoWorkItem) Upsert(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	var err error

	dwi.WorkItemID = params.WorkItemID
	dwi.Checked = params.Checked

	dwi, err = dwi.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			dwi, err = dwi.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return dwi, err
}

// Delete deletes the DemoWorkItem from the database.
func (dwi *DemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.demo_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// DemoWorkItemPaginatedByWorkItemIDAsc returns a cursor-paginated list of DemoWorkItem in Asc order.
func DemoWorkItemPaginatedByWorkItemIDAsc(ctx context.Context, db DB, workItemID int64, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.checked,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id ` +
		`FROM xo_tests.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.work_item_id > $2 GROUP BY demo_work_items.work_item_id, 
demo_work_items.checked, 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id ORDER BY 
		work_item_id Asc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// DemoWorkItemPaginatedByWorkItemIDDesc returns a cursor-paginated list of DemoWorkItem in Desc order.
func DemoWorkItemPaginatedByWorkItemIDDesc(ctx context.Context, db DB, workItemID int64, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.checked,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id ` +
		`FROM xo_tests.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.work_item_id < $2 GROUP BY demo_work_items.work_item_id, 
demo_work_items.checked, 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id ORDER BY 
		work_item_id Desc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// DemoWorkItemByWorkItemID retrieves a row from 'xo_tests.demo_work_items' as a DemoWorkItem.
//
// Generated from index 'demo_work_items_pkey'.
func DemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...DemoWorkItemSelectConfigOption) (*DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.checked,
(case when $1::boolean = true and _work_items_work_item_id.work_item_id is not null then row(_work_items_work_item_id.*) end) as work_item_work_item_id ` +
		`FROM xo_tests.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join xo_tests.work_items as _work_items_work_item_id on _work_items_work_item_id.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.work_item_id = $2 GROUP BY 
_work_items_work_item_id.work_item_id,
      _work_items_work_item_id.work_item_id,
	demo_work_items.work_item_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/db.Query: %w", err))
	}
	dwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", err))
	}

	return &dwi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_work_items_work_item_id_fkey'.
func (dwi *DemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dwi.WorkItemID)
}
