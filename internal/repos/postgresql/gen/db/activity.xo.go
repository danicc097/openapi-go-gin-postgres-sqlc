package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Activity represents a row from 'public.activities'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Activity struct {
	ActivityID   ActivityID `json:"activityID" db:"activity_id" required:"true" nullable:"false"`     // activity_id
	ProjectID    ProjectID  `json:"projectID" db:"project_id" required:"true" nullable:"false"`       // project_id
	Name         string     `json:"name" db:"name" required:"true" nullable:"false"`                  // name
	Description  string     `json:"description" db:"description" required:"true" nullable:"false"`    // description
	IsProductive bool       `json:"isProductive" db:"is_productive" required:"true" nullable:"false"` // is_productive
	DeletedAt    *time.Time `json:"deletedAt" db:"deleted_at"`                                        // deleted_at

	ProjectJoin             *Project     `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)
	ActivityTimeEntriesJoin *[]TimeEntry `json:"-" db:"time_entries" openapi-go:"ignore"`       // M2O activities

}

// ActivityCreateParams represents insert params for 'public.activities'.
type ActivityCreateParams struct {
	Description  string    `json:"description" required:"true" nullable:"false"`  // description
	IsProductive bool      `json:"isProductive" required:"true" nullable:"false"` // is_productive
	Name         string    `json:"name" required:"true" nullable:"false"`         // name
	ProjectID    ProjectID `json:"-" openapi-go:"ignore"`                         // project_id
}

type ActivityID int

// CreateActivity creates a new Activity in the database with the given params.
func CreateActivity(ctx context.Context, db DB, params *ActivityCreateParams) (*Activity, error) {
	a := &Activity{
		Description:  params.Description,
		IsProductive: params.IsProductive,
		Name:         params.Name,
		ProjectID:    params.ProjectID,
	}

	return a.Insert(ctx, db)
}

type ActivitySelectConfig struct {
	limit   string
	orderBy string
	joins   ActivityJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
}
type ActivitySelectConfigOption func(*ActivitySelectConfig)

// WithActivityLimit limits row selection.
func WithActivityLimit(limit int) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedActivityOnly limits result to records marked as deleted.
func WithDeletedActivityOnly() ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.deletedAt = " not null "
	}
}

type ActivityOrderBy string

const (
	ActivityDeletedAtDescNullsFirst ActivityOrderBy = " deleted_at DESC NULLS FIRST "
	ActivityDeletedAtDescNullsLast  ActivityOrderBy = " deleted_at DESC NULLS LAST "
	ActivityDeletedAtAscNullsFirst  ActivityOrderBy = " deleted_at ASC NULLS FIRST "
	ActivityDeletedAtAscNullsLast   ActivityOrderBy = " deleted_at ASC NULLS LAST "
)

// WithActivityOrderBy orders results by the given columns.
func WithActivityOrderBy(rows ...ActivityOrderBy) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
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

type ActivityJoins struct {
	Project     bool // O2O projects
	TimeEntries bool // M2O time_entries
}

// WithActivityJoin joins with the given tables.
func WithActivityJoin(joins ActivityJoins) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.joins = ActivityJoins{
			Project:     s.joins.Project || joins.Project,
			TimeEntries: s.joins.TimeEntries || joins.TimeEntries,
		}
	}
}

// WithActivityFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithActivityFilters(filters map[string][]any) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.filters = filters
	}
}

// WithActivityHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithActivityHavingClause(conditions map[string][]any) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.having = conditions
	}
}

const activityTableProjectJoinSQL = `-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
`

const activityTableProjectSelectSQL = `(case when _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id`

const activityTableProjectGroupBySQL = `_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id`

const activityTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id
) as xo_join_time_entries on xo_join_time_entries.time_entries_activity_id = activities.activity_id
`

const activityTableTimeEntriesSelectSQL = `COALESCE(xo_join_time_entries.time_entries, '{}') as time_entries`

const activityTableTimeEntriesGroupBySQL = `xo_join_time_entries.time_entries, activities.activity_id`

// ActivityUpdateParams represents update params for 'public.activities'.
type ActivityUpdateParams struct {
	Description  *string    `json:"description" nullable:"false"`  // description
	IsProductive *bool      `json:"isProductive" nullable:"false"` // is_productive
	Name         *string    `json:"name" nullable:"false"`         // name
	ProjectID    *ProjectID `json:"-" openapi-go:"ignore"`         // project_id
}

// SetUpdateParams updates public.activities struct fields with the specified params.
func (a *Activity) SetUpdateParams(params *ActivityUpdateParams) {
	if params.Description != nil {
		a.Description = *params.Description
	}
	if params.IsProductive != nil {
		a.IsProductive = *params.IsProductive
	}
	if params.Name != nil {
		a.Name = *params.Name
	}
	if params.ProjectID != nil {
		a.ProjectID = *params.ProjectID
	}
}

// Insert inserts the Activity to the database.
func (a *Activity) Insert(ctx context.Context, db DB) (*Activity, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.activities (
	deleted_at, description, is_productive, name, project_id
	) VALUES (
	$1, $2, $3, $4, $5
	) RETURNING * `
	// run
	logf(sqlstr, a.DeletedAt, a.Description, a.IsProductive, a.Name, a.ProjectID)

	rows, err := db.Query(ctx, sqlstr, a.DeletedAt, a.Description, a.IsProductive, a.Name, a.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Activity", Err: err}))
	}

	*a = newa

	return a, nil
}

// Update updates a Activity in the database.
func (a *Activity) Update(ctx context.Context, db DB) (*Activity, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.activities SET
	deleted_at = $1, description = $2, is_productive = $3, name = $4, project_id = $5
	WHERE activity_id = $6
	RETURNING * `
	// run
	logf(sqlstr, a.DeletedAt, a.Description, a.IsProductive, a.Name, a.ProjectID, a.ActivityID)

	rows, err := db.Query(ctx, sqlstr, a.DeletedAt, a.Description, a.IsProductive, a.Name, a.ProjectID, a.ActivityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Activity", Err: err}))
	}
	*a = newa

	return a, nil
}

// Upsert upserts a Activity in the database.
// Requires appropriate PK(s) to be set beforehand.
func (a *Activity) Upsert(ctx context.Context, db DB, params *ActivityCreateParams) (*Activity, error) {
	var err error

	a.Description = params.Description
	a.IsProductive = params.IsProductive
	a.Name = params.Name
	a.ProjectID = params.ProjectID

	a, err = a.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Activity", Err: err})
			}
			a, err = a.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Activity", Err: err})
			}
		}
	}

	return a, err
}

// Delete deletes the Activity from the database.
func (a *Activity) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.activities
	WHERE activity_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the Activity from the database via 'deleted_at'.
func (a *Activity) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.activities
	SET deleted_at = NOW()
	WHERE activity_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID); err != nil {
		return logerror(err)
	}
	// set deleted
	a.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted Activity from the database.
func (a *Activity) Restore(ctx context.Context, db DB) (*Activity, error) {
	a.DeletedAt = nil
	newa, err := a.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Restore/pgx.CollectRows: %w", &XoError{Entity: "Activity", Err: err}))
	}
	return newa, nil
}

// ActivityPaginatedByActivityID returns a cursor-paginated list of Activity.
func ActivityPaginatedByActivityID(ctx context.Context, db DB, activityID ActivityID, direction models.Direction, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.activity_id %s $1
	 %s   AND activities.deleted_at is %s  %s
  %s
  ORDER BY
		activity_id %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ActivityPaginatedByActivityID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{activityID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Activity", Err: err}))
	}
	return res, nil
}

// ActivityPaginatedByProjectID returns a cursor-paginated list of Activity.
func ActivityPaginatedByProjectID(ctx context.Context, db DB, projectID ProjectID, direction models.Direction, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.project_id %s $1
	 %s   AND activities.deleted_at is %s  %s
  %s
  ORDER BY
		project_id %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ActivityPaginatedByProjectID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Activity", Err: err}))
	}
	return res, nil
}

// ActivityByNameProjectID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivityByNameProjectID(ctx context.Context, db DB, name string, projectID ProjectID, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.name = $1 AND activities.project_id = $2
	 %s   AND activities.deleted_at is %s  %s
  %s
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ActivityByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/pgx.CollectOneRow: %w", &XoError{Entity: "Activity", Err: err}))
	}

	return &a, nil
}

// ActivitiesByName retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivitiesByName(ctx context.Context, db DB, name string, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.name = $1
	 %s   AND activities.deleted_at is %s  %s
  %s
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ActivitiesByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Activity", Err: err}))
	}
	return res, nil
}

// ActivitiesByProjectID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivitiesByProjectID(ctx context.Context, db DB, projectID ProjectID, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.project_id = $1
	 %s   AND activities.deleted_at is %s  %s
  %s
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ActivitiesByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/pgx.CollectRows: %w", &XoError{Entity: "Activity", Err: err}))
	}
	return res, nil
}

// ActivityByActivityID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_pkey'.
func ActivityByActivityID(ctx context.Context, db DB, activityID ActivityID, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{deletedAt: " null ", joins: ActivityJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, activityTableProjectSelectSQL)
		joinClauses = append(joinClauses, activityTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, activityTableProjectGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, activityTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, activityTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, activityTableTimeEntriesGroupBySQL)
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
	activities.activity_id,
	activities.deleted_at,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id %s
	 FROM public.activities %s
	 WHERE activities.activity_id = $1
	 %s   AND activities.deleted_at is %s  %s
  %s
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ActivityByActivityID */\n" + sqlstr

	// run
	// logf(sqlstr, activityID)
	rows, err := db.Query(ctx, sqlstr, append([]any{activityID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/db.Query: %w", &XoError{Entity: "Activity", Err: err}))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/pgx.CollectOneRow: %w", &XoError{Entity: "Activity", Err: err}))
	}

	return &a, nil
}

// FKProject_ProjectID returns the Project associated with the Activity's (ProjectID).
//
// Generated from foreign key 'activities_project_id_fkey'.
func (a *Activity) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, a.ProjectID)
}
