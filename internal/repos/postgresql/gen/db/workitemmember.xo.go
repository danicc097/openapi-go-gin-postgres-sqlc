package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

// WorkItemMember represents a row from 'public.work_item_member'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type WorkItemMember struct {
	WorkItemID int64               `json:"workItemID" db:"work_item_id" required:"true"`                           // work_item_id
	Member     uuid.UUID           `json:"member" db:"member" required:"true"`                                     // member
	Role       models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role

	WorkItemsJoin *[]WorkItem              `json:"-" db:"work_items" openapi-go:"ignore"` // M2M
	MembersJoin   *[]WorkItemMember_Member `json:"-" db:"members" openapi-go:"ignore"`    // M2M

}

// WorkItemMemberCreateParams represents insert params for 'public.work_item_member'
type WorkItemMemberCreateParams struct {
	WorkItemID int64               `json:"workItemID" required:"true"`                                   // work_item_id
	Member     uuid.UUID           `json:"member" required:"true"`                                       // member
	Role       models.WorkItemRole `json:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role
}

// CreateWorkItemMember creates a new WorkItemMember in the database with the given params.
func CreateWorkItemMember(ctx context.Context, db DB, params *WorkItemMemberCreateParams) (*WorkItemMember, error) {
	wim := &WorkItemMember{
		WorkItemID: params.WorkItemID,
		Member:     params.Member,
		Role:       params.Role,
	}

	return wim.Insert(ctx, db)
}

// WorkItemMemberUpdateParams represents update params for 'public.work_item_member'
type WorkItemMemberUpdateParams struct {
	WorkItemID *int64               `json:"workItemID" required:"true"`                                   // work_item_id
	Member     *uuid.UUID           `json:"member" required:"true"`                                       // member
	Role       *models.WorkItemRole `json:"role" required:"true" ref:"#/components/schemas/WorkItemRole"` // role
}

// SetUpdateParams updates public.work_item_member struct fields with the specified params.
func (wim *WorkItemMember) SetUpdateParams(params *WorkItemMemberUpdateParams) {
	if params.WorkItemID != nil {
		wim.WorkItemID = *params.WorkItemID
	}
	if params.Member != nil {
		wim.Member = *params.Member
	}
	if params.Role != nil {
		wim.Role = *params.Role
	}
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemMemberOrderBy = string

const ()

type WorkItemMemberJoins struct {
	WorkItems bool
	Members   bool
}

// WithWorkItemMemberJoin joins with the given tables.
func WithWorkItemMemberJoin(joins WorkItemMemberJoins) WorkItemMemberSelectConfigOption {
	return func(s *WorkItemMemberSelectConfig) {
		s.joins = WorkItemMemberJoins{
			WorkItems: s.joins.WorkItems || joins.WorkItems,
			Members:   s.joins.Members || joins.Members,
		}
	}
}

type WorkItemMember_Member struct {
	User User                `json:"user" db:"users"`
	Role models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole"`
}

// Insert inserts the WorkItemMember to the database.
func (wim *WorkItemMember) Insert(ctx context.Context, db DB) (*WorkItemMember, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.work_item_member (` +
		`work_item_id, member, role` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` RETURNING * `
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
	*wim = newwim

	return wim, nil
}

// Update updates a WorkItemMember in the database.
func (wim *WorkItemMember) Update(ctx context.Context, db DB) (*WorkItemMember, error) {
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
	*wim = newwim

	return wim, nil
}

// Upsert performs an upsert for WorkItemMember.
func (wim *WorkItemMember) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.work_item_member (` +
		`work_item_id, member, role` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (work_item_id, member) DO ` +
		`UPDATE SET ` +
		`role = EXCLUDED.role ` +
		` RETURNING * `
	// run
	logf(sqlstr, wim.WorkItemID, wim.Member, wim.Role)
	if _, err := db.Exec(ctx, sqlstr, wim.WorkItemID, wim.Member, wim.Role); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the WorkItemMember from the database.
func (wim *WorkItemMember) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.work_item_member ` +
		`WHERE work_item_id = $1 AND member = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wim.WorkItemID, wim.Member); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemMemberPaginatedByWorkItemIDMember returns a cursor-paginated list of WorkItemMember.
func WorkItemMemberPaginatedByWorkItemIDMember(ctx context.Context, db DB, workItemID int64, member uuid.UUID, opts ...WorkItemMemberSelectConfigOption) ([]WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role,
(case when $1::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $2::boolean = true then COALESCE(joined_members.__users, '{}') end) as members ` +
		`FROM public.work_item_member ` +
		`-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
			work_item_member.member as work_item_member_member
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_member
    	join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = work_item_member.member

-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_item_member.work_item_id
` +
		` WHERE work_item_member.work_item_id > $3 AND work_item_member.member > $4` +
		` ORDER BY 
		work_item_id DESC ,
		member DESC `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, workItemID, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemMembersByMemberWorkItemID retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_member_work_item_id_idx'.
func WorkItemMembersByMemberWorkItemID(ctx context.Context, db DB, member uuid.UUID, workItemID int64, opts ...WorkItemMemberSelectConfigOption) ([]WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role,
(case when $1::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $2::boolean = true then COALESCE(joined_members.__users, '{}') end) as members ` +
		`FROM public.work_item_member ` +
		`-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
			work_item_member.member as work_item_member_member
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_member
    	join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = work_item_member.member

-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_item_member.work_item_id
` +
		` WHERE work_item_member.member = $3 AND work_item_member.work_item_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, member, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItems, c.joins.Members, member, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByMemberWorkItemID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByMemberWorkItemID/pgx.CollectRows: %w", err))
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
work_item_member.role,
(case when $1::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $2::boolean = true then COALESCE(joined_members.__users, '{}') end) as members ` +
		`FROM public.work_item_member ` +
		`-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
			work_item_member.member as work_item_member_member
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_member
    	join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = work_item_member.member

-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_item_member.work_item_id
` +
		` WHERE work_item_member.work_item_id = $3 AND work_item_member.member = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID, member)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItems, c.joins.Members, workItemID, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_member/WorkItemMemberByWorkItemIDMember/db.Query: %w", err))
	}
	wim, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_member/WorkItemMemberByWorkItemIDMember/pgx.CollectOneRow: %w", err))
	}

	return &wim, nil
}

// WorkItemMembersByWorkItemID retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_pkey'.
func WorkItemMembersByWorkItemID(ctx context.Context, db DB, workItemID int64, opts ...WorkItemMemberSelectConfigOption) ([]WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role,
(case when $1::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $2::boolean = true then COALESCE(joined_members.__users, '{}') end) as members ` +
		`FROM public.work_item_member ` +
		`-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
			work_item_member.member as work_item_member_member
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_member
    	join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = work_item_member.member

-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_item_member.work_item_id
` +
		` WHERE work_item_member.work_item_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItems, c.joins.Members, workItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByWorkItemIDMember/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByWorkItemIDMember/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemMembersByMember retrieves a row from 'public.work_item_member' as a WorkItemMember.
//
// Generated from index 'work_item_member_pkey'.
func WorkItemMembersByMember(ctx context.Context, db DB, member uuid.UUID, opts ...WorkItemMemberSelectConfigOption) ([]WorkItemMember, error) {
	c := &WorkItemMemberSelectConfig{joins: WorkItemMemberJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_member.work_item_id,
work_item_member.member,
work_item_member.role,
(case when $1::boolean = true then COALESCE(joined_work_items.__work_items, '{}') end) as work_items,
(case when $2::boolean = true then COALESCE(joined_members.__users, '{}') end) as members ` +
		`FROM public.work_item_member ` +
		`-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
			work_item_member.member as work_item_member_member
			, array_agg(work_items.*) filter (where work_items.* is not null) as __work_items
		from work_item_member
    	join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = work_item_member.member

-- M2M join generated from "work_item_member_member_fkey"
left join (
	select
			work_item_member.work_item_id as work_item_member_work_item_id
			, work_item_member.role as role
			, array_agg(users.*) filter (where users.* is not null) as __users
		from work_item_member
    	join users on users.user_id = work_item_member.member
    group by work_item_member_work_item_id
			, role
  ) as joined_members on joined_members.work_item_member_work_item_id = work_item_member.work_item_id
` +
		` WHERE work_item_member.member = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, member)
	rows, err := db.Query(ctx, sqlstr, c.joins.WorkItems, c.joins.Members, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByWorkItemIDMember/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemMember])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemMember/WorkItemMemberByWorkItemIDMember/pgx.CollectRows: %w", err))
	}
	return res, nil
}
