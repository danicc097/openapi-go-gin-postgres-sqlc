package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// WorkItemMember represents a row from 'public.work_item_member'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type WorkItemMember struct {
	WorkItemID int64        `json:"workItemID" db:"work_item_id" required:"true"` // work_item_id
	Member     uuid.UUID    `json:"member" db:"member" required:"true"`           // member
	Role       WorkItemRole `json:"role" db:"role" required:"true"`               // role

	// xo fields
	_exists, _deleted bool
}

type WorkItemMemberSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemMemberJoins
}
type WorkItemMemberSelectConfigOption func(*WorkItemMemberSelectConfig)

// WithWorkItemMemberLimit limits row selection.
func WithWorkItemMemberLimit(limit int) WorkItemMemberSelectConfigOption {
	return func(s *WorkItemMemberSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemMemberOrderBy = string

const ()

type WorkItemMemberJoins struct {
}

// WithWorkItemMemberJoin joins with the given tables.
func WithWorkItemMemberJoin(joins WorkItemMemberJoins) WorkItemMemberSelectConfigOption {
	return func(s *WorkItemMemberSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemMember exists in the database.
func (wim *WorkItemMember) Exists() bool {
	return wim._exists
}

// Deleted returns true when the WorkItemMember has been marked for deletion from
// the database.
func (wim *WorkItemMember) Deleted() bool {
	return wim._deleted
}

// Insert inserts the WorkItemMember to the database.

func (wim *WorkItemMember) Insert(ctx context.Context, db DB) (*WorkItemMember, error) {
	switch {
	case wim._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wim._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_member (` +
		`work_item_id, member, role` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) `
	// run
	logf(sqlstr, wim.WorkItemID, wim.Member, wim.Role)
	rows, err := db.Query(ctx, sqlstr, wim.WorkItemID, wim.Member, wim.Role)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Insert/db.Query: %w", err))
	}
	newwim, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Insert/pgx.CollectOneRow: %w", err))
	}
	newwim._exists = true
	*wim = newwim

	return wim, nil
}

// Update updates a WorkItemMember in the database.
func (wim *WorkItemMember) Update(ctx context.Context, db DB) (*WorkItemMember, error) {
	switch {
	case !wim._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wim._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_member SET ` +
		`role = $1 ` +
		`WHERE work_item_id = $2  AND member = $3 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wim.Role, wim.WorkItemID, wim.Member)

	rows, err := db.Query(ctx, sqlstr, wim.Role, wim.WorkItemID, wim.Member)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Update/db.Query: %w", err))
	}
	newwim, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Update/pgx.CollectOneRow: %w", err))
	}
	newwim._exists = true
	*wim = newwim

	return wim, nil
}

// Save saves the WorkItemMember to the database.
func (wim *WorkItemMember) Save(ctx context.Context, db DB) (*WorkItemMember, error) {
	if wim.Exists() {
		return wim.Update(ctx, db)
	}
	return wim.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemMember.
func (wim *WorkItemMember) Upsert(ctx context.Context, db DB) error {
	switch {
	case wim._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_member (` +
		`work_item_id, member, role` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (work_item_id, member) DO ` +
		`UPDATE SET ` +
		`role = EXCLUDED.role  `
	// run
	logf(sqlstr, wim.WorkItemID, wim.Member, wim.Role)
	if _, err := db.Exec(ctx, sqlstr, wim.WorkItemID, wim.Member, wim.Role); err != nil {
		return logerror(err)
	}
	// set exists
	wim._exists = true
	return nil
}

// Delete deletes the WorkItemMember from the database.
func (wim *WorkItemMember) Delete(ctx context.Context, db DB) error {
	switch {
	case !wim._exists: // doesn't exist
		return nil
	case wim._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_member ` +
		`WHERE work_item_id = $1 AND member = $2 `
	// run
	logf(sqlstr, wim.WorkItemID, wim.Member)
	if _, err := db.Exec(ctx, sqlstr, wim.WorkItemID, wim.Member); err != nil {
		return logerror(err)
	}
	// set deleted
	wim._deleted = true
	return nil
}

// WorkItemMemberByMemberWorkItemID retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_member_work_item_id_idx'.
func WorkItemMemberByMemberWorkItemID(ctx context.Context, db DB, member uuid.UUID, workItemID int64, opts ...WorkItemMemberSelectConfigOption) ([]*WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role ` +
		`FROM public.work_item_member ` +
		`` +
		` WHERE work_item_member.member = $1 AND work_item_member.work_item_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, member, workItemID)
	rows, err := db.Query(ctx, sqlstr, member, workItemID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[*WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemMemberByWorkItemIDMember retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_pkey'.
func WorkItemMemberByWorkItemIDMember(ctx context.Context, db DB, workItemID int64, member uuid.UUID, opts ...WorkItemMemberSelectConfigOption) (*WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role ` +
		`FROM public.work_item_member ` +
		`` +
		` WHERE work_item_member.work_item_id = $1 AND work_item_member.member = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID, member)
	rows, err := db.Query(ctx, sqlstr, workItemID, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_member/WorkItemMemberByWorkItemIDMember/db.Query: %w", err))
	}
	wim, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_member/WorkItemMemberByWorkItemIDMember/pgx.CollectOneRow: %w", err))
	}
	wim._exists = true
	return &wim, nil
}

// FKUser_Member returns the User associated with the WorkItemMember's (Member).
//
// Generated from foreign key 'work_item_member_member_fkey'.
func (wim *WorkItemMember) FKUser_Member(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wim.Member)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemMember's (WorkItemID).
//
// Generated from foreign key 'work_item_member_work_item_id_fkey'.
func (wim *WorkItemMember) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wim.WorkItemID)
}
