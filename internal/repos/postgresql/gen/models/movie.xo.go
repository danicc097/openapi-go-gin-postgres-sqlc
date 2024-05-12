// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Movie represents a row from 'public.movies'.
type Movie struct {
	MovieID  MovieID `json:"movieID" db:"movie_id" required:"true" nullable:"false"`  // movie_id
	Title    string  `json:"title" db:"title" required:"true" nullable:"false"`       // title
	Year     int     `json:"year" db:"year" required:"true" nullable:"false"`         // year
	Synopsis string  `json:"synopsis" db:"synopsis" required:"true" nullable:"false"` // synopsis

}

// MovieCreateParams represents insert params for 'public.movies'.
type MovieCreateParams struct {
	Synopsis string `json:"synopsis" required:"true" nullable:"false"` // synopsis
	Title    string `json:"title" required:"true" nullable:"false"`    // title
	Year     int    `json:"year" required:"true" nullable:"false"`     // year
}

// MovieParams represents common params for both insert and update of 'public.movies'.
type MovieParams interface {
	GetSynopsis() *string
	GetTitle() *string
	GetYear() *int
}

func (p MovieCreateParams) GetSynopsis() *string {
	x := p.Synopsis
	return &x
}
func (p MovieUpdateParams) GetSynopsis() *string {
	return p.Synopsis
}

func (p MovieCreateParams) GetTitle() *string {
	x := p.Title
	return &x
}
func (p MovieUpdateParams) GetTitle() *string {
	return p.Title
}

func (p MovieCreateParams) GetYear() *int {
	x := p.Year
	return &x
}
func (p MovieUpdateParams) GetYear() *int {
	return p.Year
}

type MovieID int

// CreateMovie creates a new Movie in the database with the given params.
func CreateMovie(ctx context.Context, db DB, params *MovieCreateParams) (*Movie, error) {
	m := &Movie{
		Synopsis: params.Synopsis,
		Title:    params.Title,
		Year:     params.Year,
	}

	return m.Insert(ctx, db)
}

type MovieSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   MovieJoins
	filters map[string][]any
	having  map[string][]any
}
type MovieSelectConfigOption func(*MovieSelectConfig)

// WithMovieLimit limits row selection.
func WithMovieLimit(limit int) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithMovieOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithMovieOrderBy(rows map[string]*Direction) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		te := EntityFields[TableEntityMovie]
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

type MovieJoins struct {
}

// WithMovieJoin joins with the given tables.
func WithMovieJoin(joins MovieJoins) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.joins = MovieJoins{}
	}
}

// WithMovieFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithMovieFilters(filters map[string][]any) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.filters = filters
	}
}

// WithMovieHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithMovieHavingClause(conditions map[string][]any) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.having = conditions
	}
}

// MovieUpdateParams represents update params for 'public.movies'.
type MovieUpdateParams struct {
	Synopsis *string `json:"synopsis" nullable:"false"` // synopsis
	Title    *string `json:"title" nullable:"false"`    // title
	Year     *int    `json:"year" nullable:"false"`     // year
}

// SetUpdateParams updates public.movies struct fields with the specified params.
func (m *Movie) SetUpdateParams(params *MovieUpdateParams) {
	if params.Synopsis != nil {
		m.Synopsis = *params.Synopsis
	}
	if params.Title != nil {
		m.Title = *params.Title
	}
	if params.Year != nil {
		m.Year = *params.Year
	}
}

// Insert inserts the Movie to the database.
func (m *Movie) Insert(ctx context.Context, db DB) (*Movie, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.movies (
	synopsis, title, year
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, m.Synopsis, m.Title, m.Year)

	rows, err := db.Query(ctx, sqlstr, m.Synopsis, m.Title, m.Year)
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Insert/db.Query: %w", &XoError{Entity: "Movie", Err: err}))
	}
	newm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Movie", Err: err}))
	}

	*m = newm

	return m, nil
}

// Update updates a Movie in the database.
func (m *Movie) Update(ctx context.Context, db DB) (*Movie, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.movies SET 
	synopsis = $1, title = $2, year = $3 
	WHERE movie_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, m.Synopsis, m.Title, m.Year, m.MovieID)

	rows, err := db.Query(ctx, sqlstr, m.Synopsis, m.Title, m.Year, m.MovieID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Update/db.Query: %w", &XoError{Entity: "Movie", Err: err}))
	}
	newm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Movie", Err: err}))
	}
	*m = newm

	return m, nil
}

// Upsert upserts a Movie in the database.
// Requires appropriate PK(s) to be set beforehand.
func (m *Movie) Upsert(ctx context.Context, db DB, params *MovieCreateParams) (*Movie, error) {
	var err error

	m.Synopsis = params.Synopsis
	m.Title = params.Title
	m.Year = params.Year

	m, err = m.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertMovie/Insert: %w", &XoError{Entity: "Movie", Err: err})
			}
			m, err = m.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertMovie/Update: %w", &XoError{Entity: "Movie", Err: err})
			}
		}
	}

	return m, err
}

// Delete deletes the Movie from the database.
func (m *Movie) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.movies 
	WHERE movie_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, m.MovieID); err != nil {
		return logerror(err)
	}
	return nil
}

// MoviePaginated returns a cursor-paginated list of Movie.
// At least one cursor is required.
func MoviePaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...MovieSelectConfigOption) ([]Movie, error) {
	c := &MovieSelectConfig{joins: MovieJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {

		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := EntityFields[TableEntityMovie][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("Movie/Paginated/cursor: %w", &XoError{Entity: "Movie", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("movies.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("Movie/Paginated/orderBy: %w", &XoError{Entity: "Movie", Err: fmt.Errorf("at least one sorted column is required")}))
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
	movies.movie_id,
	movies.synopsis,
	movies.title,
	movies.year %s 
	 FROM public.movies %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* MoviePaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Paginated/db.Query: %w", &XoError{Entity: "Movie", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Movie", Err: err}))
	}
	return res, nil
}

// MovieByMovieID retrieves a row from 'public.movies' as a Movie.
//
// Generated from index 'movies_pkey'.
func MovieByMovieID(ctx context.Context, db DB, movieID MovieID, opts ...MovieSelectConfigOption) (*Movie, error) {
	c := &MovieSelectConfig{joins: MovieJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	movies.movie_id,
	movies.synopsis,
	movies.title,
	movies.year %s 
	 FROM public.movies %s 
	 WHERE movies.movie_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* MovieByMovieID */\n" + sqlstr

	// run
	// logf(sqlstr, movieID)
	rows, err := db.Query(ctx, sqlstr, append([]any{movieID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/db.Query: %w", &XoError{Entity: "Movie", Err: err}))
	}
	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/pgx.CollectOneRow: %w", &XoError{Entity: "Movie", Err: err}))
	}

	return &m, nil
}
