package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Movie represents a row from 'public.movies'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type Movie struct {
	MovieID  int    `json:"movieID" db:"movie_id" required:"true"`  // movie_id
	Title    string `json:"title" db:"title" required:"true"`       // title
	Year     int    `json:"year" db:"year" required:"true"`         // year
	Synopsis string `json:"synopsis" db:"synopsis" required:"true"` // synopsis

}

// MovieCreateParams represents insert params for 'public.movies'
type MovieCreateParams struct {
	Title    string `json:"title" required:"true"`    // title
	Year     int    `json:"year" required:"true"`     // year
	Synopsis string `json:"synopsis" required:"true"` // synopsis
}

// CreateMovie creates a new Movie in the database with the given params.
func CreateMovie(ctx context.Context, db DB, params *MovieCreateParams) (*Movie, error) {
	m := &Movie{
		Title:    params.Title,
		Year:     params.Year,
		Synopsis: params.Synopsis,
	}

	return m.Insert(ctx, db)
}

// MovieUpdateParams represents update params for 'public.movies'
type MovieUpdateParams struct {
	Title    *string `json:"title" required:"true"`    // title
	Year     *int    `json:"year" required:"true"`     // year
	Synopsis *string `json:"synopsis" required:"true"` // synopsis
}

// SetUpdateParams updates public.movies struct fields with the specified params.
func (m *Movie) SetUpdateParams(params *MovieUpdateParams) {
	if params.Title != nil {
		m.Title = *params.Title
	}
	if params.Year != nil {
		m.Year = *params.Year
	}
	if params.Synopsis != nil {
		m.Synopsis = *params.Synopsis
	}
}

type MovieSelectConfig struct {
	limit   string
	orderBy string
	joins   MovieJoins
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

type MovieOrderBy = string

const ()

type MovieJoins struct {
}

// WithMovieJoin joins with the given tables.
func WithMovieJoin(joins MovieJoins) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.joins = MovieJoins{}
	}
}

// Insert inserts the Movie to the database.
func (m *Movie) Insert(ctx context.Context, db DB) (*Movie, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.movies (` +
		`title, year, synopsis` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, m.Title, m.Year, m.Synopsis)

	rows, err := db.Query(ctx, sqlstr, m.Title, m.Year, m.Synopsis)
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Insert/db.Query: %w", err))
	}
	newm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Insert/pgx.CollectOneRow: %w", err))
	}

	*m = newm

	return m, nil
}

// Update updates a Movie in the database.
func (m *Movie) Update(ctx context.Context, db DB) (*Movie, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.movies SET ` +
		`title = $1, year = $2, synopsis = $3 ` +
		`WHERE movie_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, m.Title, m.Year, m.Synopsis, m.MovieID)

	rows, err := db.Query(ctx, sqlstr, m.Title, m.Year, m.Synopsis, m.MovieID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Update/db.Query: %w", err))
	}
	newm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("Movie/Update/pgx.CollectOneRow: %w", err))
	}
	*m = newm

	return m, nil
}

// Upsert performs an upsert for Movie.
func (m *Movie) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.movies (` +
		`movie_id, title, year, synopsis` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (movie_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, year = EXCLUDED.year, synopsis = EXCLUDED.synopsis ` +
		` RETURNING * `
	// run
	logf(sqlstr, m.MovieID, m.Title, m.Year, m.Synopsis)
	if _, err := db.Exec(ctx, sqlstr, m.MovieID, m.Title, m.Year, m.Synopsis); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the Movie from the database.
func (m *Movie) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.movies ` +
		`WHERE movie_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, m.MovieID); err != nil {
		return logerror(err)
	}
	return nil
}

// MovieByMovieID retrieves a row from 'public.movies' as a Movie.
//
// Generated from index 'movies_pkey'.
func MovieByMovieID(ctx context.Context, db DB, movieID int, opts ...MovieSelectConfigOption) (*Movie, error) {
	c := &MovieSelectConfig{joins: MovieJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`movies.movie_id,
movies.title,
movies.year,
movies.synopsis ` +
		`FROM public.movies ` +
		`` +
		` WHERE movies.movie_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, movieID)
	rows, err := db.Query(ctx, sqlstr, movieID)
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/db.Query: %w", err))
	}
	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/pgx.CollectOneRow: %w", err))
	}

	return &m, nil
}
