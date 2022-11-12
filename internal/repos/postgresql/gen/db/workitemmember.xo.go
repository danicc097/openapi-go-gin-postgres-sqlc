package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// WorkItemMember represents a row from 'public.work_item_member'.
type WorkItemMember struct {
	WorkItemID int64     `json:"work_item_id" db:"work_item_id"` // work_item_id
	Member     uuid.UUID `json:"member" db:"member"`             // member

	// xo fields
	_exists, _deleted bool
}

type WorkItemMemberSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemMemberJoins
}

type WorkItemMemberSelectConfigOption func(*WorkItemMemberSelectConfig)

// WorkItemMemberWithLimit limits row selection.
func WorkItemMemberWithLimit(limit int) WorkItemMemberSelectConfigOption {
	return func(s *WorkItemMemberSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemMemberOrderBy = string

type WorkItemMemberJoins struct{}

// WorkItemMemberWithJoin orders results by the given columns.
func WorkItemMemberWithJoin(joins WorkItemMemberJoins) WorkItemMemberSelectConfigOption {
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
func (wim *WorkItemMember) Insert(ctx context.Context, db DB) error {
	switch {
	case wim._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wim._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_member (` +
		`work_item_id, member` +
		`) VALUES (` +
		`$1, $2` +
		`) `
	// run
	logf(sqlstr, wim.WorkItemID, wim.Member)
	if _, err := db.Exec(ctx, sqlstr, wim.WorkItemID, wim.Member); err != nil {
		return logerror(err)
	}
	// set exists
	wim._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

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
	c := &WorkItemMemberSelectConfig{
		joins: WorkItemMemberJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member ` +
		`FROM public.work_item_member ` +
		`` +
		` WHERE member = $1 AND work_item_id = $2 `
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
	var res []*WorkItemMember
	for rows.Next() {
		wim := WorkItemMember{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&wim.WorkItemID, &wim.Member); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &wim)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// WorkItemMemberByWorkItemIDMember retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_pkey'.
func WorkItemMemberByWorkItemIDMember(ctx context.Context, db DB, workItemID int64, member uuid.UUID, opts ...WorkItemMemberSelectConfigOption) (*WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{
		joins: WorkItemMemberJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member ` +
		`FROM public.work_item_member ` +
		`` +
		` WHERE work_item_id = $1 AND member = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID, member)
	wim := WorkItemMember{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, workItemID, member).Scan(&wim.WorkItemID, &wim.Member); err != nil {
		return nil, logerror(err)
	}
	return &wim, nil
}

// User returns the User associated with the WorkItemMember's (Member).
//
// Generated from foreign key 'work_item_member_member_fkey'.
func (wim *WorkItemMember) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wim.Member)
}

// WorkItem returns the WorkItem associated with the WorkItemMember's (WorkItemID).
//
// Generated from foreign key 'work_item_member_work_item_id_fkey'.
func (wim *WorkItemMember) WorkItem(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wim.WorkItemID)
}
