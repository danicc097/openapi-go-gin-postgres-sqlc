package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// TimeEntry represents a row from 'public.time_entries'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type TimeEntry struct {
	TimeEntryID     int64     `json:"timeEntryID" db:"time_entry_id" required:"true"`        // time_entry_id
	WorkItemID      *int64    `json:"workItemID" db:"work_item_id" required:"true"`          // work_item_id
	ActivityID      int       `json:"activityID" db:"activity_id" required:"true"`           // activity_id
	TeamID          *int      `json:"teamID" db:"team_id" required:"true"`                   // team_id
	UserID          uuid.UUID `json:"userID" db:"user_id" required:"true"`                   // user_id
	Comment         string    `json:"comment" db:"comment" required:"true"`                  // comment
	Start           time.Time `json:"start" db:"start" required:"true"`                      // start
	DurationMinutes *int      `json:"durationMinutes" db:"duration_minutes" required:"true"` // duration_minutes

	ActivityJoin *Activity `json:"-" db:"activity" openapi-go:"ignore"`  // O2O (generated from M2O)
	TeamJoin     *Team     `json:"-" db:"team" openapi-go:"ignore"`      // O2O (generated from M2O)
	UserJoin     *User     `json:"-" db:"user" openapi-go:"ignore"`      // O2O (generated from M2O)
	WorkItemJoin *WorkItem `json:"-" db:"work_item" openapi-go:"ignore"` // O2O (generated from M2O)

}

// TimeEntryCreateParams represents insert params for 'public.time_entries'
type TimeEntryCreateParams struct {
	WorkItemID      *int64    `json:"workItemID" required:"true"`      // work_item_id
	ActivityID      int       `json:"activityID" required:"true"`      // activity_id
	TeamID          *int      `json:"teamID" required:"true"`          // team_id
	UserID          uuid.UUID `json:"userID" required:"true"`          // user_id
	Comment         string    `json:"comment" required:"true"`         // comment
	Start           time.Time `json:"start" required:"true"`           // start
	DurationMinutes *int      `json:"durationMinutes" required:"true"` // duration_minutes
}

// CreateTimeEntry creates a new TimeEntry in the database with the given params.
func CreateTimeEntry(ctx context.Context, db DB, params *TimeEntryCreateParams) (*TimeEntry, error) {
	te := &TimeEntry{
		WorkItemID:      params.WorkItemID,
		ActivityID:      params.ActivityID,
		TeamID:          params.TeamID,
		UserID:          params.UserID,
		Comment:         params.Comment,
		Start:           params.Start,
		DurationMinutes: params.DurationMinutes,
	}

	return te.Insert(ctx, db)
}

// TimeEntryUpdateParams represents update params for 'public.time_entries'
type TimeEntryUpdateParams struct {
	WorkItemID      **int64    `json:"workItemID" required:"true"`      // work_item_id
	ActivityID      *int       `json:"activityID" required:"true"`      // activity_id
	TeamID          **int      `json:"teamID" required:"true"`          // team_id
	UserID          *uuid.UUID `json:"userID" required:"true"`          // user_id
	Comment         *string    `json:"comment" required:"true"`         // comment
	Start           *time.Time `json:"start" required:"true"`           // start
	DurationMinutes **int      `json:"durationMinutes" required:"true"` // duration_minutes
}

// SetUpdateParams updates public.time_entries struct fields with the specified params.
func (te *TimeEntry) SetUpdateParams(params *TimeEntryUpdateParams) {
	if params.WorkItemID != nil {
		te.WorkItemID = *params.WorkItemID
	}
	if params.ActivityID != nil {
		te.ActivityID = *params.ActivityID
	}
	if params.TeamID != nil {
		te.TeamID = *params.TeamID
	}
	if params.UserID != nil {
		te.UserID = *params.UserID
	}
	if params.Comment != nil {
		te.Comment = *params.Comment
	}
	if params.Start != nil {
		te.Start = *params.Start
	}
	if params.DurationMinutes != nil {
		te.DurationMinutes = *params.DurationMinutes
	}
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
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
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type TimeEntryJoins struct {
	Activity bool
	Team     bool
	User     bool
	WorkItem bool
}

// WithTimeEntryJoin joins with the given tables.
func WithTimeEntryJoin(joins TimeEntryJoins) TimeEntrySelectConfigOption {
	return func(s *TimeEntrySelectConfig) {
		s.joins = TimeEntryJoins{

			Activity: s.joins.Activity || joins.Activity,
			Team:     s.joins.Team || joins.Team,
			User:     s.joins.User || joins.User,
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// Insert inserts the TimeEntry to the database.
func (te *TimeEntry) Insert(ctx context.Context, db DB) (*TimeEntry, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.time_entries (` +
		`work_item_id, activity_id, team_id, user_id, comment, start, duration_minutes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING * `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)

	rows, err := db.Query(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Insert/db.Query: %w", err))
	}
	newte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Insert/pgx.CollectOneRow: %w", err))
	}

	*te = newte

	return te, nil
}

// Update updates a TimeEntry in the database.
func (te *TimeEntry) Update(ctx context.Context, db DB) (*TimeEntry, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.time_entries SET ` +
		`work_item_id = $1, activity_id = $2, team_id = $3, user_id = $4, comment = $5, start = $6, duration_minutes = $7 ` +
		`WHERE time_entry_id = $8 ` +
		`RETURNING * `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)

	rows, err := db.Query(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Update/db.Query: %w", err))
	}
	newte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Update/pgx.CollectOneRow: %w", err))
	}
	*te = newte

	return te, nil
}

// Upsert performs an upsert for TimeEntry.
func (te *TimeEntry) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.time_entries (` +
		`time_entry_id, work_item_id, activity_id, team_id, user_id, comment, start, duration_minutes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`)` +
		` ON CONFLICT (time_entry_id) DO ` +
		`UPDATE SET ` +
		`work_item_id = EXCLUDED.work_item_id, activity_id = EXCLUDED.activity_id, team_id = EXCLUDED.team_id, user_id = EXCLUDED.user_id, comment = EXCLUDED.comment, start = EXCLUDED.start, duration_minutes = EXCLUDED.duration_minutes ` +
		` RETURNING * `
	// run
	logf(sqlstr, te.TimeEntryID, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)
	if _, err := db.Exec(ctx, sqlstr, te.TimeEntryID, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the TimeEntry from the database.
func (te *TimeEntry) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.time_entries ` +
		`WHERE time_entry_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, te.TimeEntryID); err != nil {
		return logerror(err)
	}
	return nil
}

// PaginatedTimeEntryByTimeEntryID returns a cursor-paginated list of TimeEntry.
func (te *TimeEntry) PaginatedTimeEntryByTimeEntryID(ctx context.Context, db DB) ([]TimeEntry, error) {
	sqlstr := `SELECT ` +
		`time_entries.time_entry_id,
time_entries.work_item_id,
time_entries.activity_id,
time_entries.team_id,
time_entries.user_id,
time_entries.comment,
time_entries.start,
time_entries.duration_minutes,
(case when $1::boolean = true and activities.activity_id is not null then row(activities.*) end) as activity,
(case when $2::boolean = true and teams.team_id is not null then row(teams.*) end) as team,
(case when $3::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $4::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.time_entries ` +
		`-- O2O join generated from "time_entries_activity_id_fkey (Generated from M2O)"
left join activities on activities.activity_id = time_entries.activity_id
-- O2O join generated from "time_entries_team_id_fkey (Generated from M2O)"
left join teams on teams.team_id = time_entries.team_id
-- O2O join generated from "time_entries_user_id_fkey (Generated from M2O)"
left join users on users.user_id = time_entries.user_id
-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = time_entries.work_item_id` +
		` WHERE time_entries.time_entry_id > $5 `
	// run

	rows, err := db.Query(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
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
time_entries.duration_minutes,
(case when $1::boolean = true and activities.activity_id is not null then row(activities.*) end) as activity,
(case when $2::boolean = true and teams.team_id is not null then row(teams.*) end) as team,
(case when $3::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $4::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.time_entries ` +
		`-- O2O join generated from "time_entries_activity_id_fkey (Generated from M2O)"
left join activities on activities.activity_id = time_entries.activity_id
-- O2O join generated from "time_entries_team_id_fkey (Generated from M2O)"
left join teams on teams.team_id = time_entries.team_id
-- O2O join generated from "time_entries_user_id_fkey (Generated from M2O)"
left join users on users.user_id = time_entries.user_id
-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = time_entries.work_item_id` +
		` WHERE time_entries.time_entry_id = $5 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, timeEntryID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activity, c.joins.Team, c.joins.User, c.joins.WorkItem, timeEntryID)
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/db.Query: %w", err))
	}
	te, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/pgx.CollectOneRow: %w", err))
	}

	return &te, nil
}

// TimeEntriesByUserIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_user_id_team_id_idx'.
func TimeEntriesByUserIDTeamID(ctx context.Context, db DB, userID uuid.UUID, teamID *int, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
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
time_entries.duration_minutes,
(case when $1::boolean = true and activities.activity_id is not null then row(activities.*) end) as activity,
(case when $2::boolean = true and teams.team_id is not null then row(teams.*) end) as team,
(case when $3::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $4::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.time_entries ` +
		`-- O2O join generated from "time_entries_activity_id_fkey (Generated from M2O)"
left join activities on activities.activity_id = time_entries.activity_id
-- O2O join generated from "time_entries_team_id_fkey (Generated from M2O)"
left join teams on teams.team_id = time_entries.team_id
-- O2O join generated from "time_entries_user_id_fkey (Generated from M2O)"
left join users on users.user_id = time_entries.user_id
-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = time_entries.work_item_id` +
		` WHERE time_entries.user_id = $5 AND time_entries.team_id = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activity, c.joins.Team, c.joins.User, c.joins.WorkItem, userID, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByUserIDTeamID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByUserIDTeamID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// TimeEntriesByWorkItemIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_work_item_id_team_id_idx'.
func TimeEntriesByWorkItemIDTeamID(ctx context.Context, db DB, workItemID *int64, teamID *int, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
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
time_entries.duration_minutes,
(case when $1::boolean = true and activities.activity_id is not null then row(activities.*) end) as activity,
(case when $2::boolean = true and teams.team_id is not null then row(teams.*) end) as team,
(case when $3::boolean = true and users.user_id is not null then row(users.*) end) as user,
(case when $4::boolean = true and work_items.work_item_id is not null then row(work_items.*) end) as work_item ` +
		`FROM public.time_entries ` +
		`-- O2O join generated from "time_entries_activity_id_fkey (Generated from M2O)"
left join activities on activities.activity_id = time_entries.activity_id
-- O2O join generated from "time_entries_team_id_fkey (Generated from M2O)"
left join teams on teams.team_id = time_entries.team_id
-- O2O join generated from "time_entries_user_id_fkey (Generated from M2O)"
left join users on users.user_id = time_entries.user_id
-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join work_items on work_items.work_item_id = time_entries.work_item_id` +
		` WHERE time_entries.work_item_id = $5 AND time_entries.team_id = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, workItemID, teamID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activity, c.joins.Team, c.joins.User, c.joins.WorkItem, workItemID, teamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByWorkItemIDTeamID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByWorkItemIDTeamID/pgx.CollectRows: %w", err))
	}
	return res, nil
}
