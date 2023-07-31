package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// TimeEntry represents a row from 'public.time_entries'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type TimeEntry struct {
	TimeEntryID     int64     `json:"timeEntryID" db:"time_entry_id" required:"true"` // time_entry_id
	WorkItemID      *int64    `json:"workItemID" db:"work_item_id"`                   // work_item_id
	ActivityID      int       `json:"activityID" db:"activity_id" required:"true"`    // activity_id
	TeamID          *int      `json:"teamID" db:"team_id"`                            // team_id
	UserID          uuid.UUID `json:"userID" db:"user_id" required:"true"`            // user_id
	Comment         string    `json:"comment" db:"comment" required:"true"`           // comment
	Start           time.Time `json:"start" db:"start" required:"true"`               // start
	DurationMinutes *int      `json:"durationMinutes" db:"duration_minutes"`          // duration_minutes

	ActivityJoin *Activity `json:"-" db:"activity_activity_id" openapi-go:"ignore"`   // O2O activities (generated from M2O)
	TeamJoin     *Team     `json:"-" db:"team_team_id" openapi-go:"ignore"`           // O2O teams (generated from M2O)
	UserJoin     *User     `json:"-" db:"user_user_id" openapi-go:"ignore"`           // O2O users (generated from M2O)
	WorkItemJoin *WorkItem `json:"-" db:"work_item_work_item_id" openapi-go:"ignore"` // O2O work_items (generated from M2O)

}

// TimeEntryCreateParams represents insert params for 'public.time_entries'.
type TimeEntryCreateParams struct {
	WorkItemID      *int64    `json:"workItemID"`                 // work_item_id
	ActivityID      int       `json:"activityID" required:"true"` // activity_id
	TeamID          *int      `json:"teamID"`                     // team_id
	UserID          uuid.UUID `json:"userID" required:"true"`     // user_id
	Comment         string    `json:"comment" required:"true"`    // comment
	Start           time.Time `json:"start" required:"true"`      // start
	DurationMinutes *int      `json:"durationMinutes"`            // duration_minutes
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

// TimeEntryUpdateParams represents update params for 'public.time_entries'.
type TimeEntryUpdateParams struct {
	WorkItemID      **int64    `json:"workItemID"`                 // work_item_id
	ActivityID      *int       `json:"activityID" required:"true"` // activity_id
	TeamID          **int      `json:"teamID"`                     // team_id
	UserID          *uuid.UUID `json:"userID" required:"true"`     // user_id
	Comment         *string    `json:"comment" required:"true"`    // comment
	Start           *time.Time `json:"start" required:"true"`      // start
	DurationMinutes **int      `json:"durationMinutes"`            // duration_minutes
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
	filters map[string][]any
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

type TimeEntryOrderBy string

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
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
		}
	}
}

type TimeEntryJoins struct {
	Activity bool // O2O activities
	Team     bool // O2O teams
	User     bool // O2O users
	WorkItem bool // O2O work_items
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

// WithTimeEntryFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithTimeEntryFilters(filters map[string][]any) TimeEntrySelectConfigOption {
	return func(s *TimeEntrySelectConfig) {
		s.filters = filters
	}
}

const timeEntryTableActivityJoinSQL = `-- O2O join generated from "time_entries_activity_id_fkey (Generated from M2O)"
left join activities as _time_entries_activity_id on _time_entries_activity_id.activity_id = time_entries.activity_id
`

const timeEntryTableActivitySelectSQL = `(case when _time_entries_activity_id.activity_id is not null then row(_time_entries_activity_id.*) end) as activity_activity_id`

const timeEntryTableActivityGroupBySQL = `_time_entries_activity_id.activity_id,
      _time_entries_activity_id.activity_id,
	time_entries.time_entry_id`

const timeEntryTableTeamJoinSQL = `-- O2O join generated from "time_entries_team_id_fkey (Generated from M2O)"
left join teams as _time_entries_team_id on _time_entries_team_id.team_id = time_entries.team_id
`

const timeEntryTableTeamSelectSQL = `(case when _time_entries_team_id.team_id is not null then row(_time_entries_team_id.*) end) as team_team_id`

const timeEntryTableTeamGroupBySQL = `_time_entries_team_id.team_id,
      _time_entries_team_id.team_id,
	time_entries.time_entry_id`

const timeEntryTableUserJoinSQL = `-- O2O join generated from "time_entries_user_id_fkey (Generated from M2O)"
left join users as _time_entries_user_id on _time_entries_user_id.user_id = time_entries.user_id
`

const timeEntryTableUserSelectSQL = `(case when _time_entries_user_id.user_id is not null then row(_time_entries_user_id.*) end) as user_user_id`

const timeEntryTableUserGroupBySQL = `_time_entries_user_id.user_id,
      _time_entries_user_id.user_id,
	time_entries.time_entry_id`

const timeEntryTableWorkItemJoinSQL = `-- O2O join generated from "time_entries_work_item_id_fkey (Generated from M2O)"
left join work_items as _time_entries_work_item_id on _time_entries_work_item_id.work_item_id = time_entries.work_item_id
`

const timeEntryTableWorkItemSelectSQL = `(case when _time_entries_work_item_id.work_item_id is not null then row(_time_entries_work_item_id.*) end) as work_item_work_item_id`

const timeEntryTableWorkItemGroupBySQL = `_time_entries_work_item_id.work_item_id,
      _time_entries_work_item_id.work_item_id,
	time_entries.time_entry_id`

// Insert inserts the TimeEntry to the database.
func (te *TimeEntry) Insert(ctx context.Context, db DB) (*TimeEntry, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.time_entries (
	work_item_id, activity_id, team_id, user_id, comment, start, duration_minutes
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7
	) RETURNING * `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)

	rows, err := db.Query(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Insert/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	newte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}

	*te = newte

	return te, nil
}

// Update updates a TimeEntry in the database.
func (te *TimeEntry) Update(ctx context.Context, db DB) (*TimeEntry, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.time_entries SET 
	work_item_id = $1, activity_id = $2, team_id = $3, user_id = $4, comment = $5, start = $6, duration_minutes = $7 
	WHERE time_entry_id = $8 
	RETURNING * `
	// run
	logf(sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)

	rows, err := db.Query(ctx, sqlstr, te.WorkItemID, te.ActivityID, te.TeamID, te.UserID, te.Comment, te.Start, te.DurationMinutes, te.TimeEntryID)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Update/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	newte, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	*te = newte

	return te, nil
}

// Upsert upserts a TimeEntry in the database.
// Requires appropriate PK(s) to be set beforehand.
func (te *TimeEntry) Upsert(ctx context.Context, db DB, params *TimeEntryCreateParams) (*TimeEntry, error) {
	var err error

	te.WorkItemID = params.WorkItemID
	te.ActivityID = params.ActivityID
	te.TeamID = params.TeamID
	te.UserID = params.UserID
	te.Comment = params.Comment
	te.Start = params.Start
	te.DurationMinutes = params.DurationMinutes

	te, err = te.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Time entry", Err: err})
			}
			te, err = te.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Time entry", Err: err})
			}
		}
	}

	return te, err
}

// Delete deletes the TimeEntry from the database.
func (te *TimeEntry) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.time_entries 
	WHERE time_entry_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, te.TimeEntryID); err != nil {
		return logerror(err)
	}
	return nil
}

// TimeEntryPaginatedByTimeEntryIDAsc returns a cursor-paginated list of TimeEntry in Asc order.
func TimeEntryPaginatedByTimeEntryIDAsc(ctx context.Context, db DB, timeEntryID int64, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Activity {
		selectClauses = append(selectClauses, timeEntryTableActivitySelectSQL)
		joinClauses = append(joinClauses, timeEntryTableActivityJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableActivityGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, timeEntryTableTeamSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableTeamGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, timeEntryTableUserSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableUserJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, timeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	time_entries.time_entry_id,
	time_entries.work_item_id,
	time_entries.activity_id,
	time_entries.team_id,
	time_entries.user_id,
	time_entries.comment,
	time_entries.start,
	time_entries.duration_minutes %s 
	 FROM public.time_entries %s 
	 WHERE time_entries.time_entry_id > $1
	 %s   %s 
  ORDER BY 
		time_entry_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* TimeEntryPaginatedByTimeEntryIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{timeEntryID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/Asc/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	return res, nil
}

// TimeEntryPaginatedByTimeEntryIDDesc returns a cursor-paginated list of TimeEntry in Desc order.
func TimeEntryPaginatedByTimeEntryIDDesc(ctx context.Context, db DB, timeEntryID int64, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Activity {
		selectClauses = append(selectClauses, timeEntryTableActivitySelectSQL)
		joinClauses = append(joinClauses, timeEntryTableActivityJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableActivityGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, timeEntryTableTeamSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableTeamGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, timeEntryTableUserSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableUserJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, timeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	time_entries.time_entry_id,
	time_entries.work_item_id,
	time_entries.activity_id,
	time_entries.team_id,
	time_entries.user_id,
	time_entries.comment,
	time_entries.start,
	time_entries.duration_minutes %s 
	 FROM public.time_entries %s 
	 WHERE time_entries.time_entry_id < $1
	 %s   %s 
  ORDER BY 
		time_entry_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* TimeEntryPaginatedByTimeEntryIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{timeEntryID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/Desc/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	return res, nil
}

// TimeEntryByTimeEntryID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_pkey'.
func TimeEntryByTimeEntryID(ctx context.Context, db DB, timeEntryID int64, opts ...TimeEntrySelectConfigOption) (*TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Activity {
		selectClauses = append(selectClauses, timeEntryTableActivitySelectSQL)
		joinClauses = append(joinClauses, timeEntryTableActivityJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableActivityGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, timeEntryTableTeamSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableTeamGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, timeEntryTableUserSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableUserJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, timeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	time_entries.time_entry_id,
	time_entries.work_item_id,
	time_entries.activity_id,
	time_entries.team_id,
	time_entries.user_id,
	time_entries.comment,
	time_entries.start,
	time_entries.duration_minutes %s 
	 FROM public.time_entries %s 
	 WHERE time_entries.time_entry_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* TimeEntryByTimeEntryID */\n" + sqlstr

	// run
	// logf(sqlstr, timeEntryID)
	rows, err := db.Query(ctx, sqlstr, append([]any{timeEntryID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/db.Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	te, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("time_entries/TimeEntryByTimeEntryID/pgx.CollectOneRow: %w", &XoError{Entity: "Time entry", Err: err}))
	}

	return &te, nil
}

// TimeEntriesByUserIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_user_id_team_id_idx'.
func TimeEntriesByUserIDTeamID(ctx context.Context, db DB, userID uuid.UUID, teamID *int, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Activity {
		selectClauses = append(selectClauses, timeEntryTableActivitySelectSQL)
		joinClauses = append(joinClauses, timeEntryTableActivityJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableActivityGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, timeEntryTableTeamSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableTeamGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, timeEntryTableUserSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableUserJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, timeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	time_entries.time_entry_id,
	time_entries.work_item_id,
	time_entries.activity_id,
	time_entries.team_id,
	time_entries.user_id,
	time_entries.comment,
	time_entries.start,
	time_entries.duration_minutes %s 
	 FROM public.time_entries %s 
	 WHERE time_entries.user_id = $1 AND time_entries.team_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* TimeEntriesByUserIDTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, userID, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID, teamID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByUserIDTeamID/Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByUserIDTeamID/pgx.CollectRows: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	return res, nil
}

// TimeEntriesByWorkItemIDTeamID retrieves a row from 'public.time_entries' as a TimeEntry.
//
// Generated from index 'time_entries_work_item_id_team_id_idx'.
func TimeEntriesByWorkItemIDTeamID(ctx context.Context, db DB, workItemID *int64, teamID *int, opts ...TimeEntrySelectConfigOption) ([]TimeEntry, error) {
	c := &TimeEntrySelectConfig{joins: TimeEntryJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Activity {
		selectClauses = append(selectClauses, timeEntryTableActivitySelectSQL)
		joinClauses = append(joinClauses, timeEntryTableActivityJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableActivityGroupBySQL)
	}

	if c.joins.Team {
		selectClauses = append(selectClauses, timeEntryTableTeamSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableTeamJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableTeamGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, timeEntryTableUserSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableUserJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, timeEntryTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, timeEntryTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, timeEntryTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	time_entries.time_entry_id,
	time_entries.work_item_id,
	time_entries.activity_id,
	time_entries.team_id,
	time_entries.user_id,
	time_entries.comment,
	time_entries.start,
	time_entries.duration_minutes %s 
	 FROM public.time_entries %s 
	 WHERE time_entries.work_item_id = $1 AND time_entries.team_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* TimeEntriesByWorkItemIDTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, teamID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByWorkItemIDTeamID/Query: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[TimeEntry])
	if err != nil {
		return nil, logerror(fmt.Errorf("TimeEntry/TimeEntriesByWorkItemIDTeamID/pgx.CollectRows: %w", &XoError{Entity: "Time entry", Err: err}))
	}
	return res, nil
}

// FKActivity_ActivityID returns the Activity associated with the TimeEntry's (ActivityID).
//
// Generated from foreign key 'time_entries_activity_id_fkey'.
func (te *TimeEntry) FKActivity_ActivityID(ctx context.Context, db DB) (*Activity, error) {
	return ActivityByActivityID(ctx, db, te.ActivityID)
}

// FKTeam_TeamID returns the Team associated with the TimeEntry's (TeamID).
//
// Generated from foreign key 'time_entries_team_id_fkey'.
func (te *TimeEntry) FKTeam_TeamID(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, *te.TeamID)
}

// FKUser_UserID returns the User associated with the TimeEntry's (UserID).
//
// Generated from foreign key 'time_entries_user_id_fkey'.
func (te *TimeEntry) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, te.UserID)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the TimeEntry's (WorkItemID).
//
// Generated from foreign key 'time_entries_work_item_id_fkey'.
func (te *TimeEntry) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*WorkItem, error) {
	return WorkItemByWorkItemID(ctx, db, *te.WorkItemID)
}
