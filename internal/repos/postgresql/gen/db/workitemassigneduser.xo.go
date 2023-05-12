package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

// WorkItemAssignedUser represents a row from 'public.work_item_assigned_user'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type WorkItemAssignedUser struct {
	WorkItemID   int64               `json:"workItemID" db:"work_item_id" required:"true"`                           // work_item_id
	AssignedUser uuid.UUID           `json:"assignedUser" db:"assigned_user" required:"true"`                        // assigned_user
	Role         models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role

	AssignedUserWorkItemsJoinWIAU *[]WorkItem__WIAU_WorkItemAssignedUser `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"`     // M2M work_item_assigned_user
	WorkItemAssignedUsersJoinWIAU *[]User__WIAU_WorkItemAssignedUser     `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user

}

// WorkItemAssignedUserCreateParams represents insert params for 'public.work_item_assigned_user'.
type WorkItemAssignedUserCreateParams struct {
	WorkItemID   int64               `json:"workItemID" required:"true"`                                   // work_item_id
	AssignedUser uuid.UUID           `json:"assignedUser" required:"true"`                                 // assigned_user
	Role         models.WorkItemRole `json:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role
}

// CreateWorkItemAssignedUser creates a new WorkItemAssignedUser in the database with the given params.
func CreateWorkItemAssignedUser(ctx context.Context, db DB, params *WorkItemAssignedUserCreateParams) (*WorkItemAssignedUser, error) {
	wiau := &WorkItemAssignedUser{
		WorkItemID:   params.WorkItemID,
		AssignedUser: params.AssignedUser,
		Role:         params.Role,
	}

	return wiau.Insert(ctx, db)
}

// WorkItemAssignedUserUpdateParams represents update params for 'public.work_item_assigned_user'
type WorkItemAssignedUserUpdateParams struct {
	WorkItemID   *int64               `json:"workItemID" required:"true"`                                   // work_item_id
	AssignedUser *uuid.UUID           `json:"assignedUser" required:"true"`                                 // assigned_user
	Role         *models.WorkItemRole `json:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role
}

// SetUpdateParams updates public.work_item_assigned_user struct fields with the specified params.
func (wiau *WorkItemAssignedUser) SetUpdateParams(params *WorkItemAssignedUserUpdateParams) {
	if params.WorkItemID != nil {
		wiau.WorkItemID = *params.WorkItemID
	}
	if params.AssignedUser != nil {
		wiau.AssignedUser = *params.AssignedUser
	}
	if params.Role != nil {
		wiau.Role = *params.Role
	}
}

type WorkItemAssignedUserSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemAssignedUserJoins
}
type WorkItemAssignedUserSelectConfigOption func(*WorkItemAssignedUserSelectConfig)

// WithWorkItemAssignedUserLimit limits row selection.
func WithWorkItemAssignedUserLimit(limit int) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemAssignedUserOrderBy = string

const ()

type WorkItemAssignedUserJoins struct {
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
	AssignedUsers         bool // M2M work_item_assigned_user
}

// WithWorkItemAssignedUserJoin joins with the given tables.
func WithWorkItemAssignedUserJoin(joins WorkItemAssignedUserJoins) WorkItemAssignedUserSelectConfigOption {
	return func(s *WorkItemAssignedUserSelectConfig) {
		s.joins = WorkItemAssignedUserJoins{
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
			AssignedUsers:         s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// WorkItem__WIAU_WorkItemAssignedUser represents a M2M join against "public.work_item_assigned_user"
type WorkItem__WIAU_WorkItemAssignedUser struct {
	WorkItem WorkItem            `json:"workItem" db:"work_items" required:"true"`
	Role     models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"`
}

// User__WIAU_WorkItemAssignedUser represents a M2M join against "public.work_item_assigned_user"
type User__WIAU_WorkItemAssignedUser struct {
	User User                `json:"user" db:"users" required:"true"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"`
}

// Insert inserts the WorkItemAssignedUser to the database.
func (wiau *WorkItemAssignedUser) Insert(ctx context.Context, db DB) (*WorkItemAssignedUser, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_assigned_user (` +
		`work_item_id, assigned_user, role` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, wiau.WorkItemID, wiau.AssignedUser, wiau.Role)
	rows, err := db.Query(ctx, sqlstr, wiau.WorkItemID, wiau.AssignedUser, wiau.Role)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Insert/db.Query: %w", err))
	}
	newwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Insert/pgx.CollectOneRow: %w", err))
	}
	*wiau = newwiau

	return wiau, nil
}

// Update updates a WorkItemAssignedUser in the database.
func (wiau *WorkItemAssignedUser) Update(ctx context.Context, db DB) (*WorkItemAssignedUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_assigned_user SET ` +
		`role = $1 ` +
		`WHERE work_item_id = $2  AND assigned_user = $3 ` +
		`RETURNING * `
	// run
	logf(sqlstr, wiau.Role, wiau.WorkItemID, wiau.AssignedUser)

	rows, err := db.Query(ctx, sqlstr, wiau.Role, wiau.WorkItemID, wiau.AssignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Update/db.Query: %w", err))
	}
	newwiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/Update/pgx.CollectOneRow: %w", err))
	}
	*wiau = newwiau

	return wiau, nil
}

// Upsert upserts a WorkItemAssignedUser in the database.
// Requires appropiate PK(s) to be set beforehand.
func (wiau *WorkItemAssignedUser) Upsert(ctx context.Context, db DB, params *WorkItemAssignedUserCreateParams) (*WorkItemAssignedUser, error) {
	var err error

	wiau.WorkItemID = params.WorkItemID
	wiau.AssignedUser = params.AssignedUser
	wiau.Role = params.Role

	wiau, err = wiau.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			wiau, err = wiau.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return wiau, err
}

// Delete deletes the WorkItemAssignedUser from the database.
func (wiau *WorkItemAssignedUser) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_assigned_user ` +
		`WHERE work_item_id = $1 AND assigned_user = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wiau.WorkItemID, wiau.AssignedUser); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemAssignedUsersByAssignedUserWorkItemID retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_assigned_user_work_item_id_idx'.
func WorkItemAssignedUsersByAssignedUserWorkItemID(ctx context.Context, db DB, assignedUser uuid.UUID, workItemID int64, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_assigned_user.work_item_id,
work_item_assigned_user.assigned_user,
work_item_assigned_user.role,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items is not null), '{}') end) as work_item_assigned_user_work_items,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') end) as work_item_assigned_user_assigned_users ` +
		`FROM public.work_item_assigned_user ` +
		`-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
			work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
			, work_item_assigned_user.role as role
			, row(work_items.*) as __work_items
		from
			work_item_assigned_user
    join work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
    group by
			work_item_assigned_user_assigned_user
			, work_items.work_item_id
			, role
  ) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user

-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
			work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
			, work_item_assigned_user.role as role
			, row(users.*) as __users
		from
			work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
			work_item_assigned_user_work_item_id
			, users.user_id
			, role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
` +
		` WHERE work_item_assigned_user.assigned_user = $3 AND work_item_assigned_user.work_item_id = $4 GROUP BY work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user, 
work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, assignedUser, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemsAssignedUser, c.joins.AssignedUsers, assignedUser, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemAssignedUserByWorkItemIDAssignedUser retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUserByWorkItemIDAssignedUser(ctx context.Context, db DB, workItemID int64, assignedUser uuid.UUID, opts ...WorkItemAssignedUserSelectConfigOption) (*WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_assigned_user.work_item_id,
work_item_assigned_user.assigned_user,
work_item_assigned_user.role,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items is not null), '{}') end) as work_item_assigned_user_work_items,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') end) as work_item_assigned_user_assigned_users ` +
		`FROM public.work_item_assigned_user ` +
		`-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
			work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
			, work_item_assigned_user.role as role
			, row(work_items.*) as __work_items
		from
			work_item_assigned_user
    join work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
    group by
			work_item_assigned_user_assigned_user
			, work_items.work_item_id
			, role
  ) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user

-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
			work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
			, work_item_assigned_user.role as role
			, row(users.*) as __users
		from
			work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
			work_item_assigned_user_work_item_id
			, users.user_id
			, role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
` +
		` WHERE work_item_assigned_user.work_item_id = $3 AND work_item_assigned_user.assigned_user = $4 GROUP BY work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user, 
work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID, assignedUser)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemsAssignedUser, c.joins.AssignedUsers, workItemID, assignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/db.Query: %w", err))
	}
	wiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectOneRow: %w", err))
	}

	return &wiau, nil
}

// WorkItemAssignedUsersByWorkItemID retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUsersByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_assigned_user.work_item_id,
work_item_assigned_user.assigned_user,
work_item_assigned_user.role,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items is not null), '{}') end) as work_item_assigned_user_work_items,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') end) as work_item_assigned_user_assigned_users ` +
		`FROM public.work_item_assigned_user ` +
		`-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
			work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
			, work_item_assigned_user.role as role
			, row(work_items.*) as __work_items
		from
			work_item_assigned_user
    join work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
    group by
			work_item_assigned_user_assigned_user
			, work_items.work_item_id
			, role
  ) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user

-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
			work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
			, work_item_assigned_user.role as role
			, row(users.*) as __users
		from
			work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
			work_item_assigned_user_work_item_id
			, users.user_id
			, role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
` +
		` WHERE work_item_assigned_user.work_item_id = $3 GROUP BY work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user, 
work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemsAssignedUser, c.joins.AssignedUsers, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemAssignedUsersByAssignedUser retrieves a row from 'public.work_item_assigned_user' as a WorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func WorkItemAssignedUsersByAssignedUser(ctx context.Context, db DB, assignedUser uuid.UUID, opts ...WorkItemAssignedUserSelectConfigOption) ([]WorkItemAssignedUser, error) {
	c := &WorkItemAssignedUserSelectConfig{joins: WorkItemAssignedUserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_assigned_user.work_item_id,
work_item_assigned_user.assigned_user,
work_item_assigned_user.role,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items is not null), '{}') end) as work_item_assigned_user_work_items,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') end) as work_item_assigned_user_assigned_users ` +
		`FROM public.work_item_assigned_user ` +
		`-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
			work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
			, work_item_assigned_user.role as role
			, row(work_items.*) as __work_items
		from
			work_item_assigned_user
    join work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
    group by
			work_item_assigned_user_assigned_user
			, work_items.work_item_id
			, role
  ) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user

-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
			work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
			, work_item_assigned_user.role as role
			, row(users.*) as __users
		from
			work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
			work_item_assigned_user_work_item_id
			, users.user_id
			, role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
` +
		` WHERE work_item_assigned_user.assigned_user = $3 GROUP BY work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user, 
work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, assignedUser)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItemsAssignedUser, c.joins.AssignedUsers, assignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKUser_AssignedUser returns the User associated with the WorkItemAssignedUser's (AssignedUser).
//
// Generated from foreign key 'work_item_assigned_user_assigned_user_fkey'.
func (wiau *WorkItemAssignedUser) FKUser_AssignedUser(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, wiau.AssignedUser)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the WorkItemAssignedUser's (WorkItemID).
//
// Generated from foreign key 'work_item_assigned_user_work_item_id_fkey'.
func (wiau *WorkItemAssignedUser) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, wiau.WorkItemID)
}
