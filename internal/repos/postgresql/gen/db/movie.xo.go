package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// Movie represents a row from 'public.movies'.
type Movie struct {
	MovieID  int    `json:"movie_id" db:"movie_id"` // movie_id
	Title    string `json:"title" db:"title"`       // title
	Year     int    `json:"year" db:"year"`         // year
	Synopsis string `json:"synopsis" db:"synopsis"` // synopsis
	// xo fields
	_exists, _deleted bool
}

type MovieSelectConfig struct {
	limit    string
	orderBy  string
	joinWith []MovieJoinBy
}

type MovieSelectConfigOption func(*MovieSelectConfig)

// MovieWithLimit limits row selection.
func MovieWithLimit(limit int) MovieSelectConfigOption {
	return func(s *MovieSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type MovieOrderBy = string

type MovieJoinBy = string

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
func (m *Movie) Insert(ctx context.Context, db DB) error {
	switch {
	case m._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case m._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.movies (` +
		`title, year, synopsis` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING movie_id `
	// run
	logf(sqlstr, m.Title, m.Year, m.Synopsis)
	if err := db.QueryRow(ctx, sqlstr, m.Title, m.Year, m.Synopsis).Scan(&m.MovieID); err != nil {
		return logerror(err)
	}
	// set exists
	m._exists = true
	return nil
}

// Update updates a Movie in the database.
func (m *Movie) Update(ctx context.Context, db DB) error {
	switch {
	case !m._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case m._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.movies SET ` +
		`title = $1, year = $2, synopsis = $3 ` +
		`WHERE movie_id = $4 `
	// run
	logf(sqlstr, m.Title, m.Year, m.Synopsis, m.MovieID)
	if _, err := db.Exec(ctx, sqlstr, m.Title, m.Year, m.Synopsis, m.MovieID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Movie to the database.
func (m *Movie) Save(ctx context.Context, db DB) error {
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
	c := &MovieSelectConfig{}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`movie_id, title, year, synopsis ` +
		`FROM public.movies ` +
		`` +
		` WHERE movie_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, movieID)
	m := Movie{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, movieID).Scan(&m.MovieID, &m.Title, &m.Year, &m.Synopsis); err != nil {
		return nil, logerror(err)
	}
	return &m, nil
}
