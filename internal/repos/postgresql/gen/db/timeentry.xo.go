package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// TimeEntry represents a row from 'public.time_entries'.
type TimeEntry struct {
	TimeEntryID     int64     `json:"time_entry_id" db:"time_entry_id"`       // time_entry_id
	WorkItemID      *int64    `json:"work_item_id" db:"work_item_id"`         // work_item_id
	ActivityID      int       `json:"activity_id" db:"activity_id"`           // activity_id
	TeamID          *int      `json:"team_id" db:"team_id"`                   // team_id
	UserID          uuid.UUID `json:"user_id" db:"user_id"`                   // user_id
	Comment         string    `json:"comment" db:"comment"`                   // comment
	Start           time.Time `json:"start" db:"start"`                       // start
	DurationMinutes *int      `json:"duration_minutes" db:"duration_minutes"` // duration_minutes

	// xo fields
	_exists, _deleted bool
}

type TimeEntrySelectConfig struct {
	limit   string
	orderBy string
	joins   TimeEntryJoins
}
type TimeEntrySelectConfigOption func(*TimeEntrySelectConfig)

// WithTimeEntryLimit limits row selection.
func WithTimeEntryLimit(limit int) TimeEntrySelectConfigOption {
	return func(s *TimeEntrySelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type TimeEntryOrderBy = string

const (
	TimeEntryStartDescNullsFirst TimeEntryOrderBy = " start DESC NULLS FIRST "
	TimeEntryStartDescNullsLast  TimeEntryOrderBy = " start DESC NULLS LAST "
	TimeEntryStartAscNullsFirst  TimeEntryOrderBy = " start ASC NULLS FIRST "
	TimeEntryStartAscNullsLast   TimeEntryOrderBy = " start ASC NULLS LAST "
)

// WithTimeEntryOrderBy orders results by the given columns.
func WithTimeEntryOrderBy(rows ...TimeEntryOrderBy) TimeEntrySelectConfigOption {
	return func(s *TimeEntrySelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type TimeEntryJoins struct{}

// WithTimeEntryJoin orders results by the given columns.
func WithTimeEntryJoin(joins TimeEntryJoins) TimeEntrySelectConfigOption {
	return func(s *TimeEntrySelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the TimeEntry exists in the database.
func (te *TimeEntry) Exists() bool {
	return te._exists
}

// Deleted returns true when the TimeEntry has been marked for deletion from
// the database.
func (te *TimeEntry) Deleted() bool {
	return te._deleted
}

// Insert inserts the TimeEntry to the database.
func (te *TimeEntry) Insert(ctx context.Context, db DB) error {
	switch {
	case te._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case te._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.time_entries (` +
		`work_item_id, activity_id, team_id, user_id, comment, start, duration_minutes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING time_entry_id `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)
	if err := db.QueryRow(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes).Scan(&te.TimeEntryID); err != nil {
		return logerror(err)
	}
	// set exists
	te._exists = true
	return nil
}

// Update updates a TimeEntry in the database.
func (te *TimeEntry) Update(ctx context.Context, db DB) error {
	switch {
	case !te._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case te._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.time_entries SET ` +
		`work_item_id = $1, activity_id = $2, team_id = $3, user_id = $4, comment = $5, start = $6, duration_minutes = $7 ` +
		`WHERE time_entry_id = $8 ` +
		`RETURNING time_entry_id `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)
	if err := db.QueryRow(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID).Scan(&te.TimeEntryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the TimeEntry to the database.
func (te *TimeEntry) Save(ctx context.Context, db DB) error {
	if te.Exists() {
		return te.Update(ctx, db)
	}
	return te.Insert(ctx, db)
}

// Upsert performs an upsert for TimeEntry.
func (te *TimeEntry) Upsert(ctx context.Context, db DB) error {
	switch {
	case te._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.time_entries (` +
		`time_entry_id, work_item_id, activity_id, team_id, user_id, comment, start, duration_minutes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`)` +
		` ON CONFLICT (time_entry_id) DO ` +
		`UPDATE SET ` +
		`work_item_id = EXCLUDED.work_item_id, activity_id = EXCLUDED.activity_id, team_id = EXCLUDED.team_id, user_id = EXCLUDED.user_id, comment = EXCLUDED.comment, start = EXCLUDED.start, duration_minutes = EXCLUDED.duration_minutes  `
	// run
	logf(sqlstr, te.TimeEntryID, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)
	if _, err := db.Exec(ctx, sqlstr, te.TimeEntryID, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes); err != nil {
		return logerror(err)
	}
	// set exists
	te._exists = true
	return nil
}

// Delete deletes the TimeEntry from the database.
func (te *TimeEntry) Delete(ctx context.Context, db DB) error {
	switch {
	case !te._exists: // doesn't exist
		return nil
	case te._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.time_entries ` +
		`WHERE time_entry_id = $1 `
	// run
	logf(sqlstr, te.TimeEntryID)
	if _, err := db.Exec(ctx, sqlstr, te.TimeEntryID); err != nil {
		return logerror(err)
	}
	// set deleted
	te._deleted = true
	return nil
}

// TimeEntryByTimeEntryID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_pkey'.
func TimeEntryByTimeEntryID(ctx context.Context, db DB, timeEntryID int64, opts ...TimeEntrySelectConfigOption) (*TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`time_entries.time_entry_id,
time_entries.work_item_id,
time_entries.activity_id,
time_entries.team_id,
time_entries.user_id,
time_entries.comment,
time_entries.start,
time_entries.duration_minutes ` +
		`FROM public.time_entries ` +
		`` +
		` WHERE time_entries.time_entry_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, timeEntryID)
	te := TimeEntry{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, timeEntryID).Scan(&te.TimeEntryID, &te.WorkItemID, &te.ActivityID, &te.TeamID, &te.UserID, &te.Comment, &te.Start, &te.DurationMinutes); err != nil {
		return nil, logerror(err)
	}
	return &te, nil
}

// TimeEntriesByUserIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_user_id_team_id_idx'.
func TimeEntriesByUserIDTeamID(ctx context.Context, db DB, userID uuid.UUID, teamID *int, opts ...TimeEntrySelectConfigOption) ([]*TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`time_entries.time_entry_id,
time_entries.work_item_id,
time_entries.activity_id,
time_entries.team_id,
time_entries.user_id,
time_entries.comment,
time_entries.start,
time_entries.duration_minutes ` +
		`FROM public.time_entries ` +
		`` +
		` WHERE time_entries.user_id = $1 AND time_entries.team_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userID, teamID)
	rows, err := db.Query(ctx, sqlstr, userID, teamID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*TimeEntry
	for rows.Next() {
		te := TimeEntry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&te.TimeEntryID, &te.WorkItemID, &te.ActivityID, &te.TeamID, &te.UserID, &te.Comment, &te.Start, &te.DurationMinutes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &te)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// TimeEntriesByWorkItemIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_work_item_id_team_id_idx'.
func TimeEntriesByWorkItemIDTeamID(ctx context.Context, db DB, workItemID *int64, teamID *int, opts ...TimeEntrySelectConfigOption) ([]*TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`time_entries.time_entry_id,
time_entries.work_item_id,
time_entries.activity_id,
time_entries.team_id,
time_entries.user_id,
time_entries.comment,
time_entries.start,
time_entries.duration_minutes ` +
		`FROM public.time_entries ` +
		`` +
		` WHERE time_entries.work_item_id = $1 AND time_entries.team_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemID, teamID)
	rows, err := db.Query(ctx, sqlstr, workItemID, teamID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*TimeEntry
	for rows.Next() {
		te := TimeEntry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&te.TimeEntryID, &te.WorkItemID, &te.ActivityID, &te.TeamID, &te.UserID, &te.Comment, &te.Start, &te.DurationMinutes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &te)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// FKActivity returns the Activity associated with the TimeEntry's (ActivityID).
//
// Generated from foreign key 'time_entries_activity_id_fkey'.
func (te *TimeEntry) FKActivity(ctx context.Context, db DB) (*Activity, error) {
	return ActivityByActivityID(ctx, db, te.ActivityID)
}

// FKTeam returns the Team associated with the TimeEntry's (TeamID).
//
// Generated from foreign key 'time_entries_team_id_fkey'.
func (te *TimeEntry) FKTeam(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, *te.TeamID)
}

// FKUser returns the User associated with the TimeEntry's (UserID).
//
// Generated from foreign key 'time_entries_user_id_fkey'.
func (te *TimeEntry) FKUser(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, te.UserID)
}

// FKWorkItem returns the WorkItem associated with the TimeEntry's (WorkItemID).
//
// Generated from foreign key 'time_entries_work_item_id_fkey'.
func (te *TimeEntry) FKWorkItem(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, *te.WorkItemID)
}
