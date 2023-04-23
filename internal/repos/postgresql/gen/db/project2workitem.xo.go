package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// Project2WorkItem represents a row from 'public.project_2_work_items'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Project2WorkItem struct {
	WorkItemID            int64      `json:"workItemID" db:"work_item_id" required:"true"`                         // work_item_id
	CustomDateForProject2 *time.Time `json:"customDateForProject2" db:"custom_date_for_project_2" required:"true"` // custom_date_for_project_2

	workItem *WorkItem `json:"-" db:"work_item"` // O2O
	// xo fields
	_exists, _deleted bool
}

func (s *Project2WorkItem) WorkItem() *WorkItem {
	return s.workItem
}

func (s *Project2WorkItem) SetWorkItem(f *WorkItem) {
	s.workItem = f
}

// Project2WorkItemCreateParams represents insert params for 'public.project_2_work_items'
type Project2WorkItemCreateParams struct {
	WorkItemID            int64      `json:"workItemID"`            // work_item_id
	CustomDateForProject2 *time.Time `json:"customDateForProject2"` // custom_date_for_project_2
}

// Project2WorkItemUpdateParams represents update params for 'public.project_2_work_items'
type Project2WorkItemUpdateParams struct {
	CustomDateForProject2 **time.Time `json:"customDateForProject2"` // custom_date_for_project_2
}

type Project2WorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   Project2WorkItemJoins
}

type Project2WorkItemSelectConfigOption func(*Project2WorkItemSelectConfig)

// WithProject2WorkItemLimit limits row selection.
func WithProject2WorkItemLimit(limit int) Project2WorkItemSelectConfigOption {
	return func(s *Project2WorkItemSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type Project2WorkItemOrderBy = string

const (
	Project2WorkItemCustomDateForProject2DescNullsFirst Project2WorkItemOrderBy = " custom_date_for_project_2 DESC NULLS FIRST "
	Project2WorkItemCustomDateForProject2DescNullsLast  Project2WorkItemOrderBy = " custom_date_for_project_2 DESC NULLS LAST "
	Project2WorkItemCustomDateForProject2AscNullsFirst  Project2WorkItemOrderBy = " custom_date_for_project_2 ASC NULLS FIRST "
	Project2WorkItemCustomDateForProject2AscNullsLast   Project2WorkItemOrderBy = " custom_date_for_project_2 ASC NULLS LAST "
)

// WithProject2WorkItemOrderBy orders results by the given columns.
func WithProject2WorkItemOrderBy(rows ...Project2WorkItemOrderBy) Project2WorkItemSelectConfigOption {
	return func(s *Project2WorkItemSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type Project2WorkItemJoins struct {
	WorkItem bool
}

// WithProject2WorkItemJoin joins with the given tables.
func WithProject2WorkItemJoin(joins Project2WorkItemJoins) Project2WorkItemSelectConfigOption {
	return func(s *Project2WorkItemSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the Project2WorkItem to the database.
func (pi *Project2WorkItem) Insert(ctx context.Context, db DB) (*Project2WorkItem, error) {
	switch {
	case pi._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case pi._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.project_2_work_items (` +
		`work_item_id, custom_date_for_project_2` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, pi.WorkItemID, pi.CustomDateForProject2)
	rows, err := db.Query(ctx, sqlstr, pi.WorkItemID, pi.CustomDateForProject2)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project2WorkItem/Insert/db.Query: %w", err))
	}
	newpi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project2WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project2WorkItem/Insert/pgx.CollectOneRow: %w", err))
	}
	newpi._exists = true
	*pi = newpi

	return pi, nil
}

// Update updates a Project2WorkItem in the database.
func (pi *Project2WorkItem) Update(ctx context.Context, db DB) (*Project2WorkItem, error) {
	switch {
	case !pi._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case pi._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.project_2_work_items SET ` +
		`custom_date_for_project_2 = $1 ` +
		`WHERE work_item_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, pi.CustomDateForProject2, pi.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, pi.CustomDateForProject2, pi.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project2WorkItem/Update/db.Query: %w", err))
	}
	newpi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project2WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project2WorkItem/Update/pgx.CollectOneRow: %w", err))
	}
	newpi._exists = true
	*pi = newpi

	return pi, nil
}

// Save saves the Project2WorkItem to the database.
func (pi *Project2WorkItem) Save(ctx context.Context, db DB) (*Project2WorkItem, error) {
	if pi._exists {
		return pi.Update(ctx, db)
	}
	return pi.Insert(ctx, db)
}

// Upsert performs an upsert for Project2WorkItem.
func (pi *Project2WorkItem) Upsert(ctx context.Context, db DB) error {
	switch {
	case pi._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.project_2_work_items (` +
		`work_item_id, custom_date_for_project_2` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (work_item_id) DO ` +
		`UPDATE SET ` +
		`custom_date_for_project_2 = EXCLUDED.custom_date_for_project_2 ` +
		` RETURNING * `
	// run
	logf(sqlstr, pi.WorkItemID, pi.CustomDateForProject2)
	if _, err := db.Exec(ctx, sqlstr, pi.WorkItemID, pi.CustomDateForProject2); err != nil {
		return logerror(err)
	}
	// set exists
	pi._exists = true
	return nil
}

// Delete deletes the Project2WorkItem from the database.
func (pi *Project2WorkItem) Delete(ctx context.Context, db DB) error {
	switch {
	case !pi._exists: // doesn't exist
		return nil
	case pi._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.project_2_work_items ` +
		`WHERE work_item_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, pi.WorkItemID); err != nil {
		return logerror(err)
	}
	// set deleted
	pi._deleted = true
	return nil
}

// Project2WorkItemByWorkItemID retrieves a row from 'public.project_2_work_items' as a Project2WorkItem.
//
// Generated from index 'project_2_work_items_pkey'.
func Project2WorkItemByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...Project2WorkItemSelectConfigOption) (*Project2WorkItem, error) {
	c := &Project2WorkItemSelectConfig{joins: Project2WorkItemJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`project_2_work_items.work_item_id,
project_2_work_items.custom_date_for_project_2,
(case when $1::boolean = true then row(work_items.*) end) as work_item ` +
		`FROM public.project_2_work_items ` +
		`-- O2O join generated from "project_2_work_items_work_item_id_fkey"
left join work_items on work_items.work_item_id = project_2_work_items.work_item_id` +
		` WHERE project_2_work_items.work_item_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItem, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("project_2_work_items/Project2WorkItemByWorkItemID/db.Query: %w", err))
	}
	pi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project2WorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("project_2_work_items/Project2WorkItemByWorkItemID/pgx.CollectOneRow: %w", err))
	}
	pi._exists = true
	return &pi, nil
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the Project2WorkItem's (WorkItemID).
//
// Generated from foreign key 'project_2_work_items_work_item_id_fkey'.
func (pi *Project2WorkItem) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, pi.WorkItemID)
}
