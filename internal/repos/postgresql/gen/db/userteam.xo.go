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
	UserID uuid.UUID `json:"userID" db:"user_id" required:"true"` // user_id

	UsersJoin *[]User `json:"-" db:"users" openapi-go:"ignore"` // M2M
	TeamsJoin *[]Team `json:"-" db:"teams" openapi-go:"ignore"` // M2M

}

// UserTeamCreateParams represents insert params for 'public.user_team'.
type UserTeamCreateParams struct {
	TeamID int       `json:"teamID" required:"true"` // team_id
	UserID uuid.UUID `json:"userID" required:"true"` // user_id
}

// CreateUserTeam creates a new UserTeam in the database with the given params.
func CreateUserTeam(ctx context.Context, db DB, params *UserTeamCreateParams) (*UserTeam, error) {
	ut := &UserTeam{
		TeamID: params.TeamID,
		UserID: params.UserID,
	}

	return ut.Insert(ctx, db)
}

// UserTeamUpdateParams represents update params for 'public.user_team'
type UserTeamUpdateParams struct {
	TeamID *int       `json:"teamID" required:"true"` // team_id
	UserID *uuid.UUID `json:"userID" required:"true"` // user_id
}

// SetUpdateParams updates public.user_team struct fields with the specified params.
func (ut *UserTeam) SetUpdateParams(params *UserTeamUpdateParams) {
	if params.TeamID != nil {
		ut.TeamID = *params.TeamID
	}
	if params.UserID != nil {
		ut.UserID = *params.UserID
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
	Users bool
	Teams bool
}

// WithUserTeamJoin joins with the given tables.
func WithUserTeamJoin(joins UserTeamJoins) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.joins = UserTeamJoins{
			Users: s.joins.Users || joins.Users,
			Teams: s.joins.Teams || joins.Teams,
		}
	}
}

// Insert inserts the UserTeam to the database.
func (ut *UserTeam) Insert(ctx context.Context, db DB) (*UserTeam, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.user_team (` +
		`team_id, user_id` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, ut.TeamID, ut.UserID)
	rows, err := db.Query(ctx, sqlstr, ut.TeamID, ut.UserID)
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
		`WHERE team_id = $1 AND user_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ut.TeamID, ut.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserTeamByUserIDTeamID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_pkey'.
func UserTeamByUserIDTeamID(ctx context.Context, db DB, userID uuid.UUID, teamID int, opts ...UserTeamSelectConfigOption) (*UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.user_id,
(case when $1::boolean = true then COALESCE(joined_users.__users, '{}') end) as users,
(case when $2::boolean = true then COALESCE(joined_teams.__teams, '{}') end) as teams ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    	join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = user_team.user_id

-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) filter (where teams.* is not null) as __teams
		from user_team
    	join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = user_team.team_id
` +
		` WHERE user_team.user_id = $3 AND user_team.team_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Users, c.joins.Teams, userID, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByUserIDTeamID/db.Query: %w", err))
	}
	ut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByUserIDTeamID/pgx.CollectOneRow: %w", err))
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
user_team.user_id,
(case when $1::boolean = true then COALESCE(joined_users.__users, '{}') end) as users,
(case when $2::boolean = true then COALESCE(joined_teams.__teams, '{}') end) as teams ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    	join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = user_team.user_id

-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) filter (where teams.* is not null) as __teams
		from user_team
    	join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = user_team.team_id
` +
		` WHERE user_team.team_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Users, c.joins.Teams, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByUserIDTeamID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByUserIDTeamID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserTeamsByTeamIDUserID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_team_id_user_id_idx'.
func UserTeamsByTeamIDUserID(ctx context.Context, db DB, teamID int, userID uuid.UUID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.user_id,
(case when $1::boolean = true then COALESCE(joined_users.__users, '{}') end) as users,
(case when $2::boolean = true then COALESCE(joined_teams.__teams, '{}') end) as teams ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    	join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = user_team.user_id

-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) filter (where teams.* is not null) as __teams
		from user_team
    	join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = user_team.team_id
` +
		` WHERE user_team.team_id = $3 AND user_team.user_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Users, c.joins.Teams, teamID, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDUserID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByTeamIDUserID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserTeamsByUserID retrieves a row from 'public.user_team' as a UserTeam.
//
// Generated from index 'user_team_user_id_idx'.
func UserTeamsByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserTeamSelectConfigOption) ([]UserTeam, error) {
	c := &UserTeamSelectConfig{joins: UserTeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_team.team_id,
user_team.user_id,
(case when $1::boolean = true then COALESCE(joined_users.__users, '{}') end) as users,
(case when $2::boolean = true then COALESCE(joined_teams.__teams, '{}') end) as teams ` +
		`FROM public.user_team ` +
		`-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
			user_team.team_id as user_team_team_id
			, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    	join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = user_team.user_id

-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) filter (where teams.* is not null) as __teams
		from user_team
    	join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = user_team.team_id
` +
		` WHERE user_team.user_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Users, c.joins.Teams, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByUserID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserTeam/UserTeamByUserID/pgx.CollectRows: %w", err))
	}
	return res, nil
}
