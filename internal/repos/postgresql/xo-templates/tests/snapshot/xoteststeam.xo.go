// Code generated by xo. DO NOT EDIT.

//lint:ignore

package got

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsTeam represents a row from 'xo_tests.teams'.
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
type XoTestsTeam struct {
	TeamID XoTestsTeamID `json:"teamID" db:"team_id" required:"true" nullable:"false"` // team_id
	Name   string        `json:"name" db:"name" required:"true" nullable:"false"`      // name
}

// XoTestsTeamCreateParams represents insert params for 'xo_tests.teams'.
type XoTestsTeamCreateParams struct {
	Name string `json:"name" required:"true" nullable:"false"` // name
}

// XoTestsTeamParams represents common params for both insert and update of 'xo_tests.teams'.
type XoTestsTeamParams interface {
	GetName() *string
}

func (p XoTestsTeamCreateParams) GetName() *string {
	x := p.Name
	return &x
}

func (p XoTestsTeamUpdateParams) GetName() *string {
	return p.Name
}

type XoTestsTeamID int

// CreateXoTestsTeam creates a new XoTestsTeam in the database with the given params.
func CreateXoTestsTeam(ctx context.Context, db DB, params *XoTestsTeamCreateParams) (*XoTestsTeam, error) {
	xtt := &XoTestsTeam{
		Name: params.Name,
	}

	return xtt.Insert(ctx, db)
}

type XoTestsTeamSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   XoTestsTeamJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsTeamSelectConfigOption func(*XoTestsTeamSelectConfig)

// WithXoTestsTeamLimit limits row selection.
func WithXoTestsTeamLimit(limit int) XoTestsTeamSelectConfigOption {
	return func(s *XoTestsTeamSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsTeamOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsTeamOrderBy(rows map[string]*models.Direction) XoTestsTeamSelectConfigOption {
	return func(s *XoTestsTeamSelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsTeam]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type XoTestsTeamJoins struct{}

// WithXoTestsTeamJoin joins with the given tables.
func WithXoTestsTeamJoin(joins XoTestsTeamJoins) XoTestsTeamSelectConfigOption {
	return func(s *XoTestsTeamSelectConfig) {
		s.joins = XoTestsTeamJoins{}
	}
}

// WithXoTestsTeamFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsTeamFilters(filters map[string][]any) XoTestsTeamSelectConfigOption {
	return func(s *XoTestsTeamSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsTeamHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsTeamHavingClause(conditions map[string][]any) XoTestsTeamSelectConfigOption {
	return func(s *XoTestsTeamSelectConfig) {
		s.having = conditions
	}
}

// XoTestsTeamUpdateParams represents update params for 'xo_tests.teams'.
type XoTestsTeamUpdateParams struct {
	Name *string `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.teams struct fields with the specified params.
func (xtt *XoTestsTeam) SetUpdateParams(params *XoTestsTeamUpdateParams) {
	if params.Name != nil {
		xtt.Name = *params.Name
	}
}

// Insert inserts the XoTestsTeam to the database.
func (xtt *XoTestsTeam) Insert(ctx context.Context, db DB) (*XoTestsTeam, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.teams (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, xtt.Name)

	rows, err := db.Query(ctx, sqlstr, xtt.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Insert/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	newxtt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}

	*xtt = newxtt

	return xtt, nil
}

// Update updates a XoTestsTeam in the database.
func (xtt *XoTestsTeam) Update(ctx context.Context, db DB) (*XoTestsTeam, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.teams SET 
	name = $1 
	WHERE team_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, xtt.Name, xtt.TeamID)

	rows, err := db.Query(ctx, sqlstr, xtt.Name, xtt.TeamID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Update/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	newxtt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}
	*xtt = newxtt

	return xtt, nil
}

// Upsert upserts a XoTestsTeam in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtt *XoTestsTeam) Upsert(ctx context.Context, db DB, params *XoTestsTeamCreateParams) (*XoTestsTeam, error) {
	var err error

	xtt.Name = params.Name

	xtt, err = xtt.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertXoTestsTeam/Insert: %w", &XoError{Entity: "Team", Err: err})
			}
			xtt, err = xtt.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertXoTestsTeam/Update: %w", &XoError{Entity: "Team", Err: err})
			}
		}
	}

	return xtt, err
}

// Delete deletes the XoTestsTeam from the database.
func (xtt *XoTestsTeam) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.teams 
	WHERE team_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtt.TeamID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsTeamPaginated returns a cursor-paginated list of XoTestsTeam.
// At least one cursor is required.
func XoTestsTeamPaginated(ctx context.Context, db DB, cursor models.PaginationCursor, opts ...XoTestsTeamSelectConfigOption) ([]XoTestsTeam, error) {
	c := &XoTestsTeamSelectConfig{
		joins:   XoTestsTeamJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsTeam][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Paginated/cursor: %w", &XoError{Entity: "Team", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == models.DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("teams.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

	paramStart := 0 // all filters will come from the user
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
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
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

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Paginated/orderBy: %w", &XoError{Entity: "Team", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	teams.name,
	teams.team_id %s 
	 FROM xo_tests.teams %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsTeamPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Paginated/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsTeam/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Team", Err: err}))
	}
	return res, nil
}

// XoTestsTeamByTeamID retrieves a row from 'xo_tests.teams' as a XoTestsTeam.
//
// Generated from index 'teams_pkey'.
func XoTestsTeamByTeamID(ctx context.Context, db DB, teamID XoTestsTeamID, opts ...XoTestsTeamSelectConfigOption) (*XoTestsTeam, error) {
	c := &XoTestsTeamSelectConfig{joins: XoTestsTeamJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	teams.name,
	teams.team_id %s 
	 FROM xo_tests.teams %s 
	 WHERE teams.team_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsTeamByTeamID */\n" + sqlstr

	// run
	// logf(sqlstr, teamID)
	rows, err := db.Query(ctx, sqlstr, append([]any{teamID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/db.Query: %w", &XoError{Entity: "Team", Err: err}))
	}
	xtt, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsTeam])
	if err != nil {
		return nil, logerror(fmt.Errorf("teams/TeamByTeamID/pgx.CollectOneRow: %w", &XoError{Entity: "Team", Err: err}))
	}

	return &xtt, nil
}
