package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// Team represents a row from 'public.teams'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Team struct {
	TeamID      int       `json:"teamID" db:"team_id" required:"true"`          // team_id
	ProjectID   int       `json:"projectID" db:"project_id" required:"true"`    // project_id
	Name        string    `json:"name" db:"name" required:"true"`               // name
	Description string    `json:"description" db:"description" required:"true"` // description
	CreatedAt   time.Time `json:"createdAt" db:"created_at" required:"true"`    // created_at
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at" required:"true"`    // updated_at

	TimeEntries *[]TimeEntry `json:"timeEntries" db:"time_entries" openapi-go:"ignore"` // O2M
	Users       *[]User      `json:"users" db:"users" openapi-go:"ignore"`              // M2M
	// xo fields
	_exists, _deleted bool
}

// TeamCreateParams represents insert params for 'public.teams'
type TeamCreateParams struct {
	ProjectID   int    `json:"projectID"`   // project_id
	Name        string `json:"name"`        // name
	Description string `json:"description"` // description
}

// TeamUpdateParams represents update params for 'public.teams'
type TeamUpdateParams struct {
	ProjectID   *int    `json:"projectID"`   // project_id
	Name        *string `json:"name"`        // name
	Description *string `json:"description"` // description
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

// WithTeamJoin joins with the given tables.
func WithTeamJoin(joins TeamJoins) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the Team to the database.
func (t *Team) Insert(ctx context.Context, db DB) (*Team, error) {
	switch {
	case t._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.teams (` +
		`project_id, name, description` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description)

	rows, err := db.Query(ctx, sqlstr, t.ProjectID, t.Name, t.Description)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Insert/db.Query: %w", err))
	}
	newt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Insert/pgx.CollectOneRow: %w", err))
	}
	newt._exists = true
	*t = newt

	return t, nil
}

// Update updates a Team in the database.
func (t *Team) Update(ctx context.Context, db DB) (*Team, error) {
	switch {
	case !t._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.teams SET ` +
		`project_id = $1, name = $2, description = $3 ` +
		`WHERE team_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description, t.CreatedAt, t.UpdatedAt, t.TeamID)

	rows, err := db.Query(ctx, sqlstr, t.ProjectID, t.Name, t.Description, t.TeamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Update/db.Query: %w", err))
	}
	newt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Update/pgx.CollectOneRow: %w", err))
	}
	newt._exists = true
	*t = newt

	return t, nil
}

// Save saves the Team to the database.
func (t *Team) Save(ctx context.Context, db DB) (*Team, error) {
	if t._exists {
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
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description ` +
		` RETURNING * `
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
(case when $1::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $2::boolean = true then COALESCE(joined_users.__users, '{}') end) as users ` +
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
		user_team.team_id as user_team_team_id
		, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id
` +
		` WHERE teams.name = $3 AND teams.project_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/db.Query: %w", err))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/pgx.CollectOneRow: %w", err))
	}
	t._exists = true
	return &t, nil
}

// TeamsByName retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamsByName(ctx context.Context, db DB, name string, opts ...TeamSelectConfigOption) ([]Team, error) {
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
(case when $1::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $2::boolean = true then COALESCE(joined_users.__users, '{}') end) as users ` +
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
		user_team.team_id as user_team_team_id
		, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id
` +
		` WHERE teams.name = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// TeamsByProjectID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_project_id_key'.
func TeamsByProjectID(ctx context.Context, db DB, projectID int, opts ...TeamSelectConfigOption) ([]Team, error) {
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
(case when $1::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $2::boolean = true then COALESCE(joined_users.__users, '{}') end) as users ` +
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
		user_team.team_id as user_team_team_id
		, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id
` +
		` WHERE teams.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, projectID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
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
(case when $1::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $2::boolean = true then COALESCE(joined_users.__users, '{}') end) as users ` +
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
		user_team.team_id as user_team_team_id
		, array_agg(users.*) filter (where users.* is not null) as __users
		from user_team
    join users on users.user_id = user_team.user_id
    group by user_team_team_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id
` +
		` WHERE teams.team_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Users, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/db.Query: %w", err))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/pgx.CollectOneRow: %w", err))
	}
	t._exists = true
	return &t, nil
}

// FKProject_ProjectID returns the Project associated with the Team's (ProjectID).
//
// Generated from foreign key 'teams_project_id_fkey'.
func (t *Team) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, t.ProjectID)
}
