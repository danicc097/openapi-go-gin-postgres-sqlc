package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Team represents a row from 'public.teams'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type Team struct {
	TeamID      int       `json:"teamID" db:"team_id" required:"true"`          // team_id
	ProjectID   int       `json:"projectID" db:"project_id" required:"true"`    // project_id
	Name        string    `json:"name" db:"name" required:"true"`               // name
	Description string    `json:"description" db:"description" required:"true"` // description
	CreatedAt   time.Time `json:"createdAt" db:"created_at" required:"true"`    // created_at
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at" required:"true"`    // updated_at

	ProjectJoin     *Project     `json:"-" db:"project" openapi-go:"ignore"`      // O2O (generated from M2O)
	TimeEntriesJoin *[]TimeEntry `json:"-" db:"time_entries" openapi-go:"ignore"` // M2O
	UsersJoin       *[]User      `json:"-" db:"users" openapi-go:"ignore"`        // M2M
	WorkItemJoin    *WorkItem    `json:"-" db:"work_item" openapi-go:"ignore"`    // O2O (inferred)

}

// TeamCreateParams represents insert params for 'public.teams'.
type TeamCreateParams struct {
	ProjectID   int    `json:"projectID" required:"true"`   // project_id
	Name        string `json:"name" required:"true"`        // name
	Description string `json:"description" required:"true"` // description
}

// CreateTeam creates a new Team in the database with the given params.
func CreateTeam(ctx context.Context, db DB, params *TeamCreateParams) (*Team, error) {
	t := &Team{
		ProjectID:   params.ProjectID,
		Name:        params.Name,
		Description: params.Description,
	}

	return t.Insert(ctx, db)
}

// TeamUpdateParams represents update params for 'public.teams'
type TeamUpdateParams struct {
	ProjectID   *int    `json:"projectID" required:"true"`   // project_id
	Name        *string `json:"name" required:"true"`        // name
	Description *string `json:"description" required:"true"` // description
}

// SetUpdateParams updates public.teams struct fields with the specified params.
func (t *Team) SetUpdateParams(params *TeamUpdateParams) {
	if params.ProjectID != nil {
		t.ProjectID = *params.ProjectID
	}
	if params.Name != nil {
		t.Name = *params.Name
	}
	if params.Description != nil {
		t.Description = *params.Description
	}
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
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
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type TeamJoins struct {
	Project     bool
	TimeEntries bool
	Users       bool
	WorkItem    bool
}

// WithTeamJoin joins with the given tables.
func WithTeamJoin(joins TeamJoins) TeamSelectConfigOption {
	return func(s *TeamSelectConfig) {
		s.joins = TeamJoins{
			Project:     s.joins.Project || joins.Project,
			TimeEntries: s.joins.TimeEntries || joins.TimeEntries,
			Users:       s.joins.Users || joins.Users,
			WorkItem:    s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// Insert inserts the Team to the database.
func (t *Team) Insert(ctx context.Context, db DB) (*Team, error) {
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

	*t = newt

	return t, nil
}

// Update updates a Team in the database.
func (t *Team) Update(ctx context.Context, db DB) (*Team, error) {
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
	*t = newt

	return t, nil
}

// Upsert upserts a Team in the database.
// Requires appropiate PK(s) to be set beforehand.
func (t *Team) Upsert(ctx context.Context, db DB, params *TeamCreateParams) (*Team, error) {
	var err error

	t.ProjectID = params.ProjectID
	t.Name = params.Name
	t.Description = params.Description

	t, err = t.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			t, err = t.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return t, err
}

// Delete deletes the Team from the database.
func (t *Team) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.teams ` +
		`WHERE team_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, t.TeamID); err != nil {
		return logerror(err)
	}
	return nil
}

// TeamPaginatedByTeamID returns a cursor-paginated list of Team.
func TeamPaginatedByTeamID(ctx context.Context, db DB, teamID int, opts ...TeamSelectConfigOption) ([]Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`teams.team_id,
teams.project_id,
teams.name,
teams.description,
teams.created_at,
teams.updated_at,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.team_id > $5 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// TeamPaginatedByProjectID returns a cursor-paginated list of Team.
func TeamPaginatedByProjectID(ctx context.Context, db DB, projectID int, opts ...TeamSelectConfigOption) ([]Team, error) {
	c := &TeamSelectConfig{joins: TeamJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`teams.team_id,
teams.project_id,
teams.name,
teams.description,
teams.created_at,
teams.updated_at,
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.project_id > $5 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.name = $5 AND teams.project_id = $6 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, c.joins.Users, c.joins.WorkItem, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/db.Query: %w", err))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByNameProjectID/pgx.CollectOneRow: %w", err))
	}

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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.name = $5 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, c.joins.Users, c.joins.WorkItem, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/pgx.CollectRows: %w", err))
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.project_id = $5 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, c.joins.Users, c.joins.WorkItem, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("Team/TeamByNameProjectID/pgx.CollectRows: %w", err))
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
(case when $1::boolean = true and projects.project_id is not null then row(projects.*) end) as project,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $3::boolean = true then ARRAY_AGG((
		joined_users.__users
		)) end) as users,
(case when $4::boolean = true and work_items.team_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.teams ` +
		`-- O2O join generated from "teams_project_id_fkey (Generated from M2O)"
left join projects on projects.project_id = teams.project_id
-- M2O join generated from "time_entries_team_id_fkey"
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
			, row(users.*) as __users
		from
			user_team
    join users on users.user_id = user_team.user_id
    group by
			user_team_team_id
			, users.user_id
  ) as joined_users on joined_users.user_team_team_id = teams.team_id

-- O2O join generated from "work_items_team_id_fkey(O2O inferred)"
left join work_items on work_items.team_id = teams.team_id` +
		` WHERE teams.team_id = $5 GROUP BY projects.project_id, projects.project_id, teams.team_id, 
joined_time_entries.time_entries, teams.team_id, 
teams.team_id, teams.team_id, 
work_items.team_id, work_items.work_item_id, teams.team_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, c.joins.Users, c.joins.WorkItem, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/db.Query: %w", err))
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Team])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/pgx.CollectOneRow: %w", err))
	}

	return &t, nil
}

// FKProject_ProjectID returns the Project associated with the Team's (ProjectID).
//
// Generated from foreign key 'teams_project_id_fkey'.
func (t *Team) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, t.ProjectID)
}
