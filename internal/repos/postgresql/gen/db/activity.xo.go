package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Activity represents a row from 'public.activities'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Activity struct {
	ActivityID   int    `json:"activityID" db:"activity_id"`     // activity_id
	ProjectID    int    `json:"projectID" db:"project_id"`       // project_id
	Name         string `json:"name" db:"name"`                  // name
	Description  string `json:"description" db:"description"`    // description
	IsProductive bool   `json:"isProductive" db:"is_productive"` // is_productive

	TimeEntries *[]TimeEntry `json:"timeEntries" db:"time_entries"` // O2M
	// xo fields
	_exists, _deleted bool
}

type ActivitySelectConfig struct {
	limit   string
	orderBy string
	joins   ActivityJoins
}
type ActivitySelectConfigOption func(*ActivitySelectConfig)

// WithActivityLimit limits row selection.
func WithActivityLimit(limit int) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type ActivityOrderBy = string

const ()

type ActivityJoins struct {
	TimeEntries bool
}

// WithActivityJoin joins with the given tables.
func WithActivityJoin(joins ActivityJoins) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Activity exists in the database.
func (a *Activity) Exists() bool {
	return a._exists
}

// Deleted returns true when the Activity has been marked for deletion from
// the database.
func (a *Activity) Deleted() bool {
	return a._deleted
}

// Insert inserts the Activity to the database.

func (a *Activity) Insert(ctx context.Context, db DB) (*Activity, error) {
	switch {
	case a._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.activities (` +
		`project_id, name, description, is_productive` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING * `
	// run
	logf(sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive)

	rows, err := db.Query(ctx, sqlstr, a.ActivityID, a.ProjectID, a.Name, a.Description, a.IsProductive)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/db.Query: %w", err))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/pgx.CollectOneRow: %w", err))
	}
	newa._exists = true
	a = &newa

	return a, nil
}

// Update updates a Activity in the database.
func (a *Activity) Update(ctx context.Context, db DB) (*Activity, error) {
	switch {
	case !a._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.activities SET ` +
		`project_id = $1, name = $2, description = $3, is_productive = $4 ` +
		`WHERE activity_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive, a.ActivityID)

	rows, err := db.Query(ctx, sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive, a.ActivityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/db.Query: %w", err))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/pgx.CollectOneRow: %w", err))
	}
	newa._exists = true
	a = &newa

	return a, nil
}

// Save saves the Activity to the database.
func (a *Activity) Save(ctx context.Context, db DB) (*Activity, error) {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Upsert performs an upsert for Activity.
func (a *Activity) Upsert(ctx context.Context, db DB) error {
	switch {
	case a._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.activities (` +
		`activity_id, project_id, name, description, is_productive` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (activity_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, is_productive = EXCLUDED.is_productive  `
	// run
	logf(sqlstr, a.ActivityID, a.ProjectID, a.Name, a.Description, a.IsProductive)
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID, a.ProjectID, a.Name, a.Description, a.IsProductive); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Delete deletes the Activity from the database.
func (a *Activity) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.activities ` +
		`WHERE activity_id = $1 `
	// run
	logf(sqlstr, a.ActivityID)
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID); err != nil {
		return logerror(err)
	}
	// set deleted
	a._deleted = true
	return nil
}

// ActivityByNameProjectID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivityByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true then array_agg(joined_time_entries.time_entries) filter (where joined_teams.teams is not null) end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2M join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.name = $2 AND activities.project_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/db.Query: %w", err))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/pgx.CollectOneRow: %w", err))
	}
	a._exists = true
	return &a, nil
}

// ActivityByActivityID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_pkey'.
func ActivityByActivityID(ctx context.Context, db DB, activityID int, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true then array_agg(joined_time_entries.time_entries) filter (where joined_teams.teams is not null) end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2M join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.activity_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, activityID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, activityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/db.Query: %w", err))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/pgx.CollectOneRow: %w", err))
	}
	a._exists = true
	return &a, nil
}

// FKProject_ProjectID returns the Project associated with the Activity's (ProjectID).
//
// Generated from foreign key 'activities_project_id_fkey'.
func (a *Activity) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, a.ProjectID)
}
