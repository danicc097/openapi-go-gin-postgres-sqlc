package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// ActivityPublic represents fields that may be exposed from 'public.activities'
// and embedded in other response models.
type ActivityPublic struct {
	ActivityID   int    `json:"activityID"`   // activity_id
	Name         string `json:"name"`         // name
	Description  string `json:"description"`  // description
	IsProductive bool   `json:"isProductive"` // is_productive

	TimeEntries *[]TimeEntryPublic `json:"timeEntries"` // O2M
}

// Activity represents a row from 'public.activities'.
type Activity struct {
	ActivityID   int    `json:"activity_id" db:"activity_id" openapi-json:"activityID"`       // activity_id
	Name         string `json:"name" db:"name" openapi-json:"name"`                           // name
	Description  string `json:"description" db:"description" openapi-json:"description"`      // description
	IsProductive bool   `json:"is_productive" db:"is_productive" openapi-json:"isProductive"` // is_productive

	TimeEntries *[]TimeEntry `json:"time_entries" db:"time_entries" openapi-json:"timeEntries"` // O2M
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

type ActivityJoins struct {
	TimeEntries bool
}

// WithActivityJoin orders results by the given columns.
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
func (a *Activity) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.activities (` +
		`name, description, is_productive` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING activity_id `
	// run
	logf(sqlstr, a.Name, a.Description, a.IsProductive)
	if err := db.QueryRow(ctx, sqlstr, a.Name, a.Description, a.IsProductive).Scan(&a.ActivityID); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Update updates a Activity in the database.
func (a *Activity) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.activities SET ` +
		`name = $1, description = $2, is_productive = $3 ` +
		`WHERE activity_id = $4 ` +
		`RETURNING activity_id `
	// run
	logf(sqlstr, a.Name, a.Description, a.IsProductive, a.ActivityID)
	if err := db.QueryRow(ctx, sqlstr, a.Name, a.Description, a.IsProductive, a.ActivityID).Scan(&a.ActivityID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Activity to the database.
func (a *Activity) Save(ctx context.Context, db DB) error {
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
		`activity_id, name, description, is_productive` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (activity_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description, is_productive = EXCLUDED.is_productive  `
	// run
	logf(sqlstr, a.ActivityID, a.Name, a.Description, a.IsProductive)
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID, a.Name, a.Description, a.IsProductive); err != nil {
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

// ActivityByName retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_key'.
func ActivityByName(ctx context.Context, db DB, name string, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries ` +
		`FROM public.activities ` +
		`-- O2M join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.name = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name)
	a := Activity{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, name).Scan(&a.ActivityID, &a.Name, &a.Description, &a.IsProductive, &a.TimeEntries); err != nil {
		return nil, logerror(err)
	}
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
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries ` +
		`FROM public.activities ` +
		`-- O2M join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.activity_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, activityID)
	a := Activity{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, activityID).Scan(&a.ActivityID, &a.Name, &a.Description, &a.IsProductive, &a.TimeEntries); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}
