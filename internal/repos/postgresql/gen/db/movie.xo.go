package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Movie represents a row from 'public.movies'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Movie struct {
	MovieID  int    `json:"movieID" db:"movie_id" required:"true"`  // movie_id
	Title    string `json:"title" db:"title" required:"true"`       // title
	Year     int    `json:"year" db:"year" required:"true"`         // year
	Synopsis string `json:"synopsis" db:"synopsis" required:"true"` // synopsis

	// xo fields
	_exists, _deleted bool
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type MovieOrderBy = string

const ()

type MovieJoins struct {
}

// WithMovieJoin joins with the given tables.
func WithMovieJoin(joins MovieJoins) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Movie exists in the database.
func (m *Movie) Exists() bool {
	return m._exists
}

// Deleted returns true when the Movie has been marked for deletion from
// the database.
func (m *Movie) Deleted() bool {
	return m._deleted
}

// Insert inserts the Movie to the database.

func (m *Movie) Insert(ctx context.Context, db DB) (*Movie, error) {
	switch {
	case m._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case m._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
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
	newm._exists = true
	*m = newm

	return m, nil
}

// Update updates a Movie in the database.
func (m *Movie) Update(ctx context.Context, db DB) (*Movie, error) {
	switch {
	case !m._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case m._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
	newm._exists = true
	*m = newm

	return m, nil
}

// Save saves the Movie to the database.
func (m *Movie) Save(ctx context.Context, db DB) (*Movie, error) {
	if m.Exists() {
		return m.Update(ctx, db)
	}
	return m.Insert(ctx, db)
}

// Upsert performs an upsert for Movie.
func (m *Movie) Upsert(ctx context.Context, db DB) error {
	switch {
	case m._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.movies (` +
		`movie_id, title, year, synopsis` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (movie_id) DO ` +
		`UPDATE SET ` +
		`title = EXCLUDED.title, year = EXCLUDED.year, synopsis = EXCLUDED.synopsis  `
	// run
	logf(sqlstr, m.MovieID, m.Title, m.Year, m.Synopsis)
	if _, err := db.Exec(ctx, sqlstr, m.MovieID, m.Title, m.Year, m.Synopsis); err != nil {
		return logerror(err)
	}
	// set exists
	m._exists = true
	return nil
}

// Delete deletes the Movie from the database.
func (m *Movie) Delete(ctx context.Context, db DB) error {
	switch {
	case !m._exists: // doesn't exist
		return nil
	case m._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.movies ` +
		`WHERE movie_id = $1 `
	// run
	logf(sqlstr, m.MovieID)
	if _, err := db.Exec(ctx, sqlstr, m.MovieID); err != nil {
		return logerror(err)
	}
	// set deleted
	m._deleted = true
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
	logf(sqlstr, movieID)
	rows, err := db.Query(ctx, sqlstr, movieID)
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/db.Query: %w", err))
	}
	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Movie])
	if err != nil {
		return nil, logerror(fmt.Errorf("movies/MovieByMovieID/pgx.CollectOneRow: %w", err))
	}
	m._exists = true
	return &m, nil
}
