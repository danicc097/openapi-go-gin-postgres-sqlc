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

// DemoWorkItem represents a row from 'public.demo_work_items'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type DemoWorkItem struct {
	WorkItemID    int64     `json:"workItemID" db:"work_item_id" required:"true"`       // work_item_id
	Ref           string    `json:"ref" db:"ref" required:"true"`                       // ref
	Line          string    `json:"line" db:"line" required:"true"`                     // line
	LastMessageAt time.Time `json:"lastMessageAt" db:"last_message_at" required:"true"` // last_message_at
	Reopened      bool      `json:"reopened" db:"reopened" required:"true"`             // reopened

	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O

}

// DemoWorkItemCreateParams represents insert params for 'public.demo_work_items'
type DemoWorkItemCreateParams struct {
	WorkItemID    int64     `json:"workItemID" required:"true"`    // work_item_id
	Ref           string    `json:"ref" required:"true"`           // ref
	Line          string    `json:"line" required:"true"`          // line
	LastMessageAt time.Time `json:"lastMessageAt" required:"true"` // last_message_at
	Reopened      bool      `json:"reopened" required:"true"`      // reopened
}

// CreateDemoWorkItem creates a new DemoWorkItem in the database with the given params.
func CreateDemoWorkItem(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	dwi := &DemoWorkItem{
		WorkItemID:    params.WorkItemID,
		Ref:           params.Ref,
		Line:          params.Line,
		LastMessageAt: params.LastMessageAt,
		Reopened:      params.Reopened,
	}

	return dwi.Insert(ctx, db)
}

// UpsertDemoWorkItem upserts a DemoWorkItem in the database with the given params.
func UpsertDemoWorkItem(ctx context.Context, db DB, params *DemoWorkItemCreateParams) (*DemoWorkItem, error) {
	var err error
	dwi := &DemoWorkItem{
		WorkItemID:    params.WorkItemID,
		Ref:           params.Ref,
		Line:          params.Line,
		LastMessageAt: params.LastMessageAt,
		Reopened:      params.Reopened,
	}

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

	return dwi, nil
}

// DemoWorkItemUpdateParams represents update params for 'public.demo_work_items'
type DemoWorkItemUpdateParams struct {
	Ref           *string    `json:"ref" required:"true"`           // ref
	Line          *string    `json:"line" required:"true"`          // line
	LastMessageAt *time.Time `json:"lastMessageAt" required:"true"` // last_message_at
	Reopened      *bool      `json:"reopened" required:"true"`      // reopened
}

// SetUpdateParams updates public.demo_work_items struct fields with the specified params.
func (dwi *DemoWorkItem) SetUpdateParams(params *DemoWorkItemUpdateParams) {
	if params.Ref != nil {
		dwi.Ref = *params.Ref
	}
	if params.Line != nil {
		dwi.Line = *params.Line
	}
	if params.LastMessageAt != nil {
		dwi.LastMessageAt = *params.LastMessageAt
	}
	if params.Reopened != nil {
		dwi.Reopened = *params.Reopened
	}
}

type DemoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   DemoWorkItemJoins
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

const (
	DemoWorkItemLastMessageAtDescNullsFirst DemoWorkItemOrderBy = " last_message_at DESC NULLS FIRST "
	DemoWorkItemLastMessageAtDescNullsLast  DemoWorkItemOrderBy = " last_message_at DESC NULLS LAST "
	DemoWorkItemLastMessageAtAscNullsFirst  DemoWorkItemOrderBy = " last_message_at ASC NULLS FIRST "
	DemoWorkItemLastMessageAtAscNullsLast   DemoWorkItemOrderBy = " last_message_at ASC NULLS LAST "
)

// WithDemoWorkItemOrderBy orders results by the given columns.
func WithDemoWorkItemOrderBy(rows ...DemoWorkItemOrderBy) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type DemoWorkItemJoins struct {
	WorkItem bool
}

// WithDemoWorkItemJoin joins with the given tables.
func WithDemoWorkItemJoin(joins DemoWorkItemJoins) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.joins = DemoWorkItemJoins{
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// Insert inserts the DemoWorkItem to the database.
func (dwi *DemoWorkItem) Insert(ctx context.Context, db DB) (*DemoWorkItem, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.demo_work_items (` +
		`work_item_id, ref, line, last_message_at, reopened` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, dwi.WorkItemID, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened)
	rows, err := db.Query(ctx, sqlstr, dwi.WorkItemID, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened)
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
	sqlstr := `UPDATE public.demo_work_items SET ` +
		`ref = $1, line = $2, last_message_at = $3, reopened = $4 ` +
		`WHERE work_item_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened, dwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened, dwi.WorkItemID)
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

// Delete deletes the DemoWorkItem from the database.
func (dwi *DemoWorkItem) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dwi.WorkItemID); err != nil {
		return logerror(err)
	}
	return nil
}

// DemoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of DemoWorkItem.
func DemoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.ref,
demo_work_items.line,
demo_work_items.last_message_at,
demo_work_items.reopened,
(case when $1::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.work_item_id > $2` +
		` ORDER BY 
		work_item_id DESC `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// DemoWorkItemByWorkItemID retrieves a row from 'public.demo_work_items' as a DemoWorkItem.
//
// Generated from index 'demo_work_items_pkey'.
func DemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...DemoWorkItemSelectConfigOption) (*DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.ref,
demo_work_items.line,
demo_work_items.last_message_at,
demo_work_items.reopened,
(case when $1::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.work_item_id = $2 `
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

// DemoWorkItemsByRefLine retrieves a row from 'public.demo_work_items' as a DemoWorkItem.
//
// Generated from index 'demo_work_items_ref_line_idx'.
func DemoWorkItemsByRefLine(ctx context.Context, db DB, ref string, line string, opts ...DemoWorkItemSelectConfigOption) ([]DemoWorkItem, error) {
	c := &DemoWorkItemSelectConfig{joins: DemoWorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`demo_work_items.work_item_id,
demo_work_items.ref,
demo_work_items.line,
demo_work_items.last_message_at,
demo_work_items.reopened,
(case when $1::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.demo_work_items ` +
		`-- O2O join generated from "demo_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = demo_work_items.work_item_id` +
		` WHERE demo_work_items.ref = $2 AND demo_work_items.line = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, ref, line)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, ref, line)
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/DemoWorkItemsByRefLine/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("DemoWorkItem/DemoWorkItemsByRefLine/pgx.CollectRows: %w", err))
	}
	return res, nil
}
