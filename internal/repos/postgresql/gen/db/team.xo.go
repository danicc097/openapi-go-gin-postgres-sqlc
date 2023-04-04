package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Team represents a row from 'public.teams'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Team struct {
	TeamID      int       `json:"teamID" db:"team_id"`          // team_id
	ProjectID   int       `json:"projectID" db:"project_id"`    // project_id
	Name        string    `json:"name" db:"name"`               // name
	Description string    `json:"description" db:"description"` // description
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`    // created_at
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`    // updated_at

	TimeEntries *[]TimeEntry `json:"time_entries" db:"time_entries"` // O2M
	Users       *[]User      `json:"users" db:"users"`               // M2M
	// xo fields
	_exists, _deleted bool
}

type TeamSelectConfig struct {
	limit   string
	orderBy string
	joins   TeamJoins
}
type TeamSelectConfigOption func(*TeamSelectConfig)

// WithTeamLimit limits row selection.
func WithTeamLimit(limit int) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type TeamOrderBy = string

const (
	TeamCreatedAtDescNullsFirst TeamOrderBy = " created_at DESC NULLS FIRST "
	TeamCreatedAtDescNullsLast  TeamOrderBy = " created_at DESC NULLS LAST "
	TeamCreatedAtAscNullsFirst  TeamOrderBy = " created_at ASC NULLS FIRST "
	TeamCreatedAtAscNullsLast   TeamOrderBy = " created_at ASC NULLS LAST "
	TeamUpdatedAtDescNullsFirst TeamOrderBy = " updated_at DESC NULLS FIRST "
	TeamUpdatedAtDescNullsLast  TeamOrderBy = " updated_at DESC NULLS LAST "
	TeamUpdatedAtAscNullsFirst  TeamOrderBy = " updated_at ASC NULLS FIRST "
	TeamUpdatedAtAscNullsLast   TeamOrderBy = " updated_at ASC NULLS LAST "
)

// WithTeamOrderBy orders results by the given columns.
func WithTeamOrderBy(rows ...TeamOrderBy) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type TeamJoins struct {
	TimeEntries bool
	Users       bool
}

// WithTeamJoin orders results by the given columns.
func WithTeamJoin(joins TeamJoins) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Team exists in the database.
func (t *Team) Exists() bool {
	return t._exists
}

// Deleted returns true when the Team has been marked for deletion from
// the database.
func (t *Team) Deleted() bool {
	return t._deleted
}

// Insert inserts the Team to the database.
func (t *Team) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.teams (` +
		`project_id, name, description` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING team_id, created_at, updated_at `
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description)
	if err := db.QueryRow(ctx, sqlstr, t.ProjectID, t.Name, t.Description).Scan(&t.TeamID, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Update updates a Team in the database.
func (t *Team) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.teams SET ` +
		`project_id = $1, name = $2, description = $3 ` +
		`WHERE team_id = $4 ` +
		`RETURNING team_id, created_at, updated_at `
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description, t.CreatedAt, t.UpdatedAt, t.TeamID)
	if err := db.QueryRow(ctx, sqlstr, t.ProjectID, t.Name, t.Description, t.TeamID).Scan(&t.TeamID, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Team to the database.
func (t *Team) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Team.
func (t *Team) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.teams (` +
		`team_id, project_id, name, description` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (team_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description  `
	// run
	logf(sqlstr, t.TeamID, t.ProjectID, t.Name, t.Description)
	if _, err := db.Exec(ctx, sqlstr, t.TeamID, t.ProjectID, t.Name, t.Description); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Team from the database.
func (t *Team) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.teams ` +
		`WHERE team_id = $1 `
	// run
	logf(sqlstr, t.TeamID)
	if _, err := db.Exec(ctx, sqlstr, t.TeamID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TeamByNameProjectID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...TeamSelectConfigOption) (*Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`teams.team_id,
teams.project_id,
teams.name,
teams.description,
teams.created_at,
teams.updated_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_users.users end)::jsonb as users ` +
		`FROM public.teams ` +
		`-- O2M join generated from "time_entries_team_id_fkey"
left join (
  select
  team_id as time_entries_team_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        team_id) joined_time_entries on joined_time_entries.time_entries_team_id = teams.team_id
-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
		team_id as users_team_id
		, array_agg(users.*) as users
	from
		user_team
		join users using (user_id)
	where
		team_id in (
			select
				team_id
			from
				user_team
			where
				user_id = any (
					select
						user_id
					from
						users))
			group by
				team_id) joined_users on joined_users.users_team_id = teams.team_id` +
		` WHERE teams.name = $3 AND teams.project_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name, projectID)
	t := Team{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, name, projectID).Scan(&t.TeamID, &t.ProjectID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.TimeEntries, &t.Users); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// TeamByTeamID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_pkey'.
func TeamByTeamID(ctx context.Context, db DB, teamID int, opts ...TeamSelectConfigOption) (*Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`teams.team_id,
teams.project_id,
teams.name,
teams.description,
teams.created_at,
teams.updated_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_users.users end)::jsonb as users ` +
		`FROM public.teams ` +
		`-- O2M join generated from "time_entries_team_id_fkey"
left join (
  select
  team_id as time_entries_team_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        team_id) joined_time_entries on joined_time_entries.time_entries_team_id = teams.team_id
-- M2M join generated from "user_team_user_id_fkey"
left join (
	select
		team_id as users_team_id
		, array_agg(users.*) as users
	from
		user_team
		join users using (user_id)
	where
		team_id in (
			select
				team_id
			from
				user_team
			where
				user_id = any (
					select
						user_id
					from
						users))
			group by
				team_id) joined_users on joined_users.users_team_id = teams.team_id` +
		` WHERE teams.team_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, teamID)
	t := Team{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, teamID).Scan(&t.TeamID, &t.ProjectID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.TimeEntries, &t.Users); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// FKProject_ProjectID returns the Project associated with the Team's (ProjectID).
//
// Generated from foreign key 'teams_project_id_fkey'.
func (t *Team) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, t.ProjectID)
}
