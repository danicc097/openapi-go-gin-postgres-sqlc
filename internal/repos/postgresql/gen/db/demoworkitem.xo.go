package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// DemoWorkItem represents a row from 'public.demo_work_items'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type DemoWorkItem struct {
	WorkItemID    int64     `json:"workItemID" db:"work_item_id" required:"true"`       // work_item_id
	Ref           string    `json:"ref" db:"ref" required:"true"`                       // ref
	Line          string    `json:"line" db:"line" required:"true"`                     // line
	LastMessageAt time.Time `json:"lastMessageAt" db:"last_message_at" required:"true"` // last_message_at
	Reopened      bool      `json:"reopened" db:"reopened" required:"true"`             // reopened

	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O
	// xo fields
	_exists, _deleted bool
}

// DemoWorkItemCreateParams represents insert params for 'public.demo_work_items'
type DemoWorkItemCreateParams struct {
	WorkItemID    int64     `json:"workItemID"`    // work_item_id
	Ref           string    `json:"ref"`           // ref
	Line          string    `json:"line"`          // line
	LastMessageAt time.Time `json:"lastMessageAt"` // last_message_at
	Reopened      bool      `json:"reopened"`      // reopened
}

// DemoWorkItemUpdateParams represents update params for 'public.demo_work_items'
type DemoWorkItemUpdateParams struct {
	Ref           *string    `json:"ref"`           // ref
	Line          *string    `json:"line"`          // line
	LastMessageAt *time.Time `json:"lastMessageAt"` // last_message_at
	Reopened      *bool      `json:"reopened"`      // reopened
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
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
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type DemoWorkItemJoins struct {
	WorkItem bool
}

// WithDemoWorkItemJoin joins with the given tables.
func WithDemoWorkItemJoin(joins DemoWorkItemJoins) DemoWorkItemSelectConfigOption {
	return func(s *DemoWorkItemSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the DemoWorkItem to the database.
func (dwi *DemoWorkItem) Insert(ctx context.Context, db DB) (*DemoWorkItem, error) {
	switch {
	case dwi._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dwi._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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
	newdwi._exists = true
	*dwi = newdwi

	return dwi, nil
}

// Update updates a DemoWorkItem in the database.
func (dwi *DemoWorkItem) Update(ctx context.Context, db DB) (*DemoWorkItem, error) {
	switch {
	case !dwi._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dwi._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
	newdwi._exists = true
	*dwi = newdwi

	return dwi, nil
}

// Save saves the DemoWorkItem to the database.
func (dwi *DemoWorkItem) Save(ctx context.Context, db DB) (*DemoWorkItem, error) {
	if dwi._exists {
		return dwi.Update(ctx, db)
	}
	return dwi.Insert(ctx, db)
}

// Upsert performs an upsert for DemoWorkItem.
func (dwi *DemoWorkItem) Upsert(ctx context.Context, db DB) error {
	switch {
	case dwi._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.demo_work_items (` +
		`work_item_id, ref, line, last_message_at, reopened` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`ref = EXCLUDED.ref, line = EXCLUDED.line, last_message_at = EXCLUDED.last_message_at, reopened = EXCLUDED.reopened ` +
		` RETURNING * `
	// run
	logf(sqlstr, dwi.WorkItemID, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened)
	if _, err := db.Exec(ctx, sqlstr, dwi.WorkItemID, dwi.Ref, dwi.Line, dwi.LastMessageAt, dwi.Reopened); err != nil {
		return logerror(err)
	}
	// set exists
	dwi._exists = true
	return nil
}

// Delete deletes the DemoWorkItem from the database.
func (dwi *DemoWorkItem) Delete(ctx context.Context, db DB) error {
	switch {
	case !dwi._exists: // doesn't exist
		return nil
	case dwi._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, dwi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	dwi._deleted = true
	return nil
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
	dwi._exists = true

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
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[DemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_work_items_work_item_id_fkey'.
func (dwi *DemoWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dwi.WorkItemID)
}
