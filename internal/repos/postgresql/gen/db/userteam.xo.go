package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UserTeam represents a row from 'public.user_team'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type UserTeam struct {
	TeamID int       `json:"teamID" db:"team_id" required:"true"` // team_id
	Member uuid.UUID `json:"member" db:"member" required:"true"`  // member

	TeamsJoinMember *[]Team `json:"-" db:"teams_member" openapi-go:"ignore"` // M2M
	MembersJoin     *[]User `json:"-" db:"members" openapi-go:"ignore"`      // M2M

}

// UserTeamCreateParams represents insert params for 'public.user_team'.
type UserTeamCreateParams struct {
	TeamID int       `json:"teamID" required:"true"` // team_id
	Member uuid.UUID `json:"member" required:"true"` // member
}

// CreateUserTeam creates a new UserTeam in the database with the given params.
func CreateUserTeam(ctx context.Context, db DB, params *UserTeamCreateParams) (*UserTeam, error) {
	ut := &UserTeam{
		TeamID: params.TeamID,
		Member: params.Member,
	}

	return ut.Insert(ctx, db)
}

// UserTeamUpdateParams represents update params for 'public.user_team'
type UserTeamUpdateParams struct {
	TeamID *int       `json:"teamID" required:"true"` // team_id
	Member *uuid.UUID `json:"member" required:"true"` // member
}

// SetUpdateParams updates public.user_team struct fields with the specified params.
func (ut *UserTeam) SetUpdateParams(params *UserTeamUpdateParams) {
	if params.TeamID != nil {
		ut.TeamID = *params.TeamID
	}
	if params.Member != nil {
		ut.Member = *params.Member
	}
}

type UserTeamSelectConfig struct {
	limit   string
	orderBy string
	joins   UserTeamJoins
}
type UserTeamSelectConfigOption func(*UserTeamSelectConfig)

// WithUserTeamLimit limits row selection.
func WithUserTeamLimit(limit int) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserTeamOrderBy = string

const ()

type UserTeamJoins struct {
	TeamsMember bool
	Members     bool
}

// WithUserTeamJoin joins with the given tables.
func WithUserTeamJoin(joins UserTeamJoins) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.joins = UserTeamJoins{
			TeamsMember: s.joins.TeamsMember || joins.TeamsMember,
			Members:     s.joins.Members || joins.Members,
		}
	}
}

// Insert inserts the UserTeam to the database.
func (ut *UserTeam) Insert(ctx context.Context, db DB) (*UserTeam, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.user_team (` +
		`team_id, member` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, ut.TeamID, ut.Member)
	rows, err := db.Query(ctx, sqlstr, ut.TeamID, ut.Member)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/Insert/db.Query: %w", err))
	}
	newut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/Insert/pgx.CollectOneRow: %w", err))
	}
	*ut = newut

	return ut, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the UserTeam from the database.
func (ut *UserTeam) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.user_team ` +
		`WHERE team_id = $1 AND member = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ut.TeamID, ut.Member); err != nil {
		return logerror(err)
	}
	return nil
}

// UserTeamsByMember retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_member_idx'.
func UserTeamsByMember(ctx context.Context, db DB, member uuid.UUID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.member,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_teams_member.__teams
		)) filter (where joined_teams_member.__teams is not null), '{}') end) as teams_member,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_users.__users
		)) filter (where joined_users.__users is not null), '{}') end) as users ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.member as user_team_member
			, row(teams.*) as __teams
		from
			user_team
    join teams on teams.team_id = user_team.team_id
    group by
			user_team_member
			, teams.team_id
  ) as joined_teams_member on joined_teams_member.user_team_member = user_team.team_id

-- M2M join generated from "user_team_member_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.member
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = user_team.member
` +
		` WHERE user_team.member = $3 GROUP BY user_team.team_id, user_team.team_id, user_team.member, 
user_team.member, user_team.team_id, user_team.member `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, member)
	rows, err := db.Query(ctx, sqlstr, c.joins.TeamsMember, c.joins.Members, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMember/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMember/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserTeamByMemberTeamID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_pkey'.
func UserTeamByMemberTeamID(ctx context.Context, db DB, member uuid.UUID, teamID int, opts ...UserTeamSelectConfigOption) (*UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.member,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_teams_member.__teams
		)) filter (where joined_teams_member.__teams is not null), '{}') end) as teams_member,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_users.__users
		)) filter (where joined_users.__users is not null), '{}') end) as users ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.member as user_team_member
			, row(teams.*) as __teams
		from
			user_team
    join teams on teams.team_id = user_team.team_id
    group by
			user_team_member
			, teams.team_id
  ) as joined_teams_member on joined_teams_member.user_team_member = user_team.team_id

-- M2M join generated from "user_team_member_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.member
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = user_team.member
` +
		` WHERE user_team.member = $3 AND user_team.team_id = $4 GROUP BY user_team.team_id, user_team.team_id, user_team.member, 
user_team.member, user_team.team_id, user_team.member `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, member, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TeamsMember, c.joins.Members, member, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByMemberTeamID/db.Query: %w", err))
	}
	ut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByMemberTeamID/pgx.CollectOneRow: %w", err))
	}

	return &ut, nil
}

// UserTeamsByTeamID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_pkey'.
func UserTeamsByTeamID(ctx context.Context, db DB, teamID int, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.member,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_teams_member.__teams
		)) filter (where joined_teams_member.__teams is not null), '{}') end) as teams_member,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_users.__users
		)) filter (where joined_users.__users is not null), '{}') end) as users ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.member as user_team_member
			, row(teams.*) as __teams
		from
			user_team
    join teams on teams.team_id = user_team.team_id
    group by
			user_team_member
			, teams.team_id
  ) as joined_teams_member on joined_teams_member.user_team_member = user_team.team_id

-- M2M join generated from "user_team_member_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.member
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = user_team.member
` +
		` WHERE user_team.team_id = $3 GROUP BY user_team.team_id, user_team.team_id, user_team.member, 
user_team.member, user_team.team_id, user_team.member `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TeamsMember, c.joins.Members, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMemberTeamID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByMemberTeamID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserTeamsByTeamIDMember retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_team_id_member_idx'.
func UserTeamsByTeamIDMember(ctx context.Context, db DB, teamID int, member uuid.UUID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.member,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_teams_member.__teams
		)) filter (where joined_teams_member.__teams is not null), '{}') end) as teams_member,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_users.__users
		)) filter (where joined_users.__users is not null), '{}') end) as users ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.member as user_team_member
			, row(teams.*) as __teams
		from
			user_team
    join teams on teams.team_id = user_team.team_id
    group by
			user_team_member
			, teams.team_id
  ) as joined_teams_member on joined_teams_member.user_team_member = user_team.team_id

-- M2M join generated from "user_team_member_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.member
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = user_team.member
` +
		` WHERE user_team.team_id = $3 AND user_team.member = $4 GROUP BY user_team.team_id, user_team.team_id, user_team.member, 
user_team.member, user_team.team_id, user_team.member `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID, member)
	rows, err := db.Query(ctx, sqlstr, c.joins.TeamsMember, c.joins.Members, teamID, member)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDMember/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDMember/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKUser_Member returns the User associated with the UserTeam's (Member).
//
// Generated from foreign key 'user_team_member_fkey'.
func (ut *UserTeam) FKUser_Member(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ut.Member)
}

// FKTeam_TeamID returns the Team associated with the UserTeam's (TeamID).
//
// Generated from foreign key 'user_team_team_id_fkey'.
func (ut *UserTeam) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, ut.TeamID)
}
