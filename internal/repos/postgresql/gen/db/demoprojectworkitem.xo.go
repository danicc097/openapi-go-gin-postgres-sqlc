package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// DemoProjectWorkItem represents a row from 'public.demo_project_work_items'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type DemoProjectWorkItem struct {
	WorkItemID    int64     `json:"workItemID" db:"work_item_id"`       // work_item_id
	Ref           string    `json:"ref" db:"ref"`                       // ref
	Line          string    `json:"line" db:"line"`                     // line
	LastMessageAt time.Time `json:"lastMessageAt" db:"last_message_at"` // last_message_at
	Reopened      bool      `json:"reopened" db:"reopened"`             // reopened

	WorkItem *WorkItem `json:"workItem" db:"work_item"` // O2O
	// xo fields
	_exists, _deleted bool
}

type DemoProjectWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   DemoProjectWorkItemJoins
}
type DemoProjectWorkItemSelectConfigOption func(*DemoProjectWorkItemSelectConfig)

// WithDemoProjectWorkItemLimit limits row selection.
func WithDemoProjectWorkItemLimit(limit int) DemoProjectWorkItemSelectConfigOption {
	return func(s *DemoProjectWorkItemSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type DemoProjectWorkItemOrderBy = string

const (
	DemoProjectWorkItemLastMessageAtDescNullsFirst DemoProjectWorkItemOrderBy = " last_message_at DESC NULLS FIRST "
	DemoProjectWorkItemLastMessageAtDescNullsLast  DemoProjectWorkItemOrderBy = " last_message_at DESC NULLS LAST "
	DemoProjectWorkItemLastMessageAtAscNullsFirst  DemoProjectWorkItemOrderBy = " last_message_at ASC NULLS FIRST "
	DemoProjectWorkItemLastMessageAtAscNullsLast   DemoProjectWorkItemOrderBy = " last_message_at ASC NULLS LAST "
)

// WithDemoProjectWorkItemOrderBy orders results by the given columns.
func WithDemoProjectWorkItemOrderBy(rows ...DemoProjectWorkItemOrderBy) DemoProjectWorkItemSelectConfigOption {
	return func(s *DemoProjectWorkItemSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type DemoProjectWorkItemJoins struct {
	WorkItem bool
}

// WithDemoProjectWorkItemJoin orders results by the given columns.
func WithDemoProjectWorkItemJoin(joins DemoProjectWorkItemJoins) DemoProjectWorkItemSelectConfigOption {
	return func(s *DemoProjectWorkItemSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the DemoProjectWorkItem exists in the database.
func (dpwi *DemoProjectWorkItem) Exists() bool {
	return dpwi._exists
}

// Deleted returns true when the DemoProjectWorkItem has been marked for deletion from
// the database.
func (dpwi *DemoProjectWorkItem) Deleted() bool {
	return dpwi._deleted
}

// Insert inserts the DemoProjectWorkItem to the database.
/* TODO insert may generate rows. use Query instead of exec */
func (dpwi *DemoProjectWorkItem) Insert(ctx context.Context, db DB) (*DemoProjectWorkItem, error) {
	switch {
	case dpwi._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dpwi._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.demo_project_work_items (` +
		`work_item_id, ref, line, last_message_at, reopened` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) `
	// run
	logf(sqlstr, dpwi.WorkItemID, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened)
	rows, err := db.Query(ctx, sqlstr, dpwi.WorkItemID, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened)
	if err != nil {
		return nil, logerror(fmt.Errorf("db.Query: %w", err))
	}
	newdpwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoProjectWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectOneRow: %w", err))
	}
	newdpwi._exists = true
	dpwi = &newdpwi

	return dpwi, nil
}

// Update updates a DemoProjectWorkItem in the database.
func (dpwi *DemoProjectWorkItem) Update(ctx context.Context, db DB) (*DemoProjectWorkItem, error) {
	switch {
	case !dpwi._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dpwi._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.demo_project_work_items SET ` +
		`ref = $1, line = $2, last_message_at = $3, reopened = $4 ` +
		`WHERE work_item_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened, dpwi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened, dpwi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("db.Query: %w", err))
	}
	newdpwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoProjectWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectOneRow: %w", err))
	}
	newdpwi._exists = true
	dpwi = &newdpwi

	return dpwi, nil
}

// Save saves the DemoProjectWorkItem to the database.
func (dpwi *DemoProjectWorkItem) Save(ctx context.Context, db DB) (*DemoProjectWorkItem, error) {
	if dpwi.Exists() {
		return dpwi.Update(ctx, db)
	}
	return dpwi.Insert(ctx, db)
}

// Upsert performs an upsert for DemoProjectWorkItem.
func (dpwi *DemoProjectWorkItem) Upsert(ctx context.Context, db DB) error {
	switch {
	case dpwi._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.demo_project_work_items (` +
		`work_item_id, ref, line, last_message_at, reopened` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`ref = EXCLUDED.ref, line = EXCLUDED.line, last_message_at = EXCLUDED.last_message_at, reopened = EXCLUDED.reopened  `
	// run
	logf(sqlstr, dpwi.WorkItemID, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened)
	if _, err := db.Exec(ctx, sqlstr, dpwi.WorkItemID, dpwi.Ref, dpwi.Line, dpwi.LastMessageAt, dpwi.Reopened); err != nil {
		return logerror(err)
	}
	// set exists
	dpwi._exists = true
	return nil
}

// Delete deletes the DemoProjectWorkItem from the database.
func (dpwi *DemoProjectWorkItem) Delete(ctx context.Context, db DB) error {
	switch {
	case !dpwi._exists: // doesn't exist
		return nil
	case dpwi._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.demo_project_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	logf(sqlstr, dpwi.WorkItemID)
	if _, err := db.Exec(ctx, sqlstr, dpwi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	dpwi._deleted = true
	return nil
}

// DemoProjectWorkItemByWorkItemID retrieves a row from 'public.demo_project_work_items' as a DemoProjectWorkItem.
//
// Generated from index 'demo_project_work_items_pkey'.
func DemoProjectWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...DemoProjectWorkItemSelectConfigOption) (*DemoProjectWorkItem, error) {
	c := &DemoProjectWorkItemSelectConfig{joins: DemoProjectWorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`demo_project_work_items.work_item_id,
demo_project_work_items.ref,
demo_project_work_items.line,
demo_project_work_items.last_message_at,
demo_project_work_items.reopened,
(case when $1::boolean = true then row(work_items.*) end) as work_item ` +
		`FROM public.demo_project_work_items ` +
		`-- O2O join generated from "demo_project_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = demo_project_work_items.work_item_id` +
		` WHERE demo_project_work_items.work_item_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("db.Query: %w", err))
	}
	dpwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[DemoProjectWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectOneRow: %w", err))
	}
	dpwi._exists = true
	return &dpwi, nil
}

// DemoProjectWorkItemsByRefLine retrieves a row from 'public.demo_project_work_items' as a DemoProjectWorkItem.
//
// Generated from index 'demo_project_work_items_ref_line_idx'.
func DemoProjectWorkItemsByRefLine(ctx context.Context, db DB, ref string, line string, opts ...DemoProjectWorkItemSelectConfigOption) ([]*DemoProjectWorkItem, error) {
	c := &DemoProjectWorkItemSelectConfig{joins: DemoProjectWorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`demo_project_work_items.work_item_id,
demo_project_work_items.ref,
demo_project_work_items.line,
demo_project_work_items.last_message_at,
demo_project_work_items.reopened,
(case when $1::boolean = true then row(work_items.*) end) as work_item ` +
		`FROM public.demo_project_work_items ` +
		`-- O2O join generated from "demo_project_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = demo_project_work_items.work_item_id` +
		` WHERE demo_project_work_items.ref = $2 AND demo_project_work_items.line = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, ref, line)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, ref, line)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DemoProjectWorkItem
	for rows.Next() {
		dpwi := DemoProjectWorkItem{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dpwi.WorkItemID, &dpwi.Ref, &dpwi.Line, &dpwi.LastMessageAt, &dpwi.Reopened); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dpwi)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the DemoProjectWorkItem's (WorkItemID).
//
// Generated from foreign key 'demo_project_work_items_work_item_id_fkey'.
func (dpwi *DemoProjectWorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, dpwi.WorkItemID)
}
