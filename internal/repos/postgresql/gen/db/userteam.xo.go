package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UserTeam represents a row from 'public.user_team'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type UserTeam struct {
	TeamID int       `json:"teamID" db:"team_id" required:"true"` // team_id
	UserID uuid.UUID `json:"userID" db:"user_id" required:"true"` // user_id

	// xo fields
	_exists, _deleted bool
}

// UserTeamCreateParams represents insert params for 'public.user_team'
type UserTeamCreateParams struct {
	TeamID int       `json:"teamID"` // team_id
	UserID uuid.UUID `json:"userID"` // user_id
}

// UserTeamUpdateParams represents update params for 'public.user_team'
type UserTeamUpdateParams struct {
	TeamID *int       `json:"teamID"` // team_id
	UserID *uuid.UUID `json:"userID"` // user_id
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserTeamOrderBy = string

const ()

type UserTeamJoins struct {
}

// WithUserTeamJoin joins with the given tables.
func WithUserTeamJoin(joins UserTeamJoins) UserTeamSelectConfigOption {
	return func(s *UserTeamSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the UserTeam exists in the database.
func (ut *UserTeam) Exists() bool {
	return ut._exists
}

// Deleted returns true when the UserTeam has been marked for deletion from
// the database.
func (ut *UserTeam) Deleted() bool {
	return ut._deleted
}

// Insert inserts the UserTeam to the database.
func (ut *UserTeam) Insert(ctx context.Context, db DB) (*UserTeam, error) {
	switch {
	case ut._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ut._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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
	newut._exists = true
	*ut = newut

	return ut, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the UserTeam from the database.
func (ut *UserTeam) Delete(ctx context.Context, db DB) error {
	switch {
	case !ut._exists: // doesn't exist
		return nil
	case ut._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	sqlstr := `DELETE FROM public.user_team ` +
		`WHERE team_id = $1 AND user_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ut.TeamID, ut.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	ut._deleted = true
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
user_team.user_id ` +
		`FROM public.user_team ` +
		`` +
		` WHERE user_team.user_id = $1 AND user_team.team_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID, teamID)
	rows, err := db.Query(ctx, sqlstr, userID, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByUserIDTeamID/db.Query: %w", err))
	}
	ut, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_team/UserTeamByUserIDTeamID/pgx.CollectOneRow: %w", err))
	}
	ut._exists = true
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
user_team.user_id ` +
		`FROM public.user_team ` +
		`` +
		` WHERE user_team.team_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, teamID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
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
user_team.user_id ` +
		`FROM public.user_team ` +
		`` +
		` WHERE user_team.team_id = $1 AND user_team.user_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID, userID)
	rows, err := db.Query(ctx, sqlstr, teamID, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
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
user_team.user_id ` +
		`FROM public.user_team ` +
		`` +
		` WHERE user_team.user_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKTeam_TeamID returns the Team associated with the UserTeam's (TeamID).
//
// Generated from foreign key 'user_team_team_id_fkey'.
func (ut *UserTeam) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, ut.TeamID)
}

// FKUser_UserID returns the User associated with the UserTeam's (UserID).
//
// Generated from foreign key 'user_team_user_id_fkey'.
func (ut *UserTeam) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ut.UserID)
}
