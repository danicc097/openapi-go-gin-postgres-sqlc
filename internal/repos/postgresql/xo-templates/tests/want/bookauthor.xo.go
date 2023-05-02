package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookAuthor represents a row from 'public.book_authors'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type BookAuthor struct {
	BookID   int       `json:"bookID" db:"book_id" required:"true"`     // book_id
	AuthorID uuid.UUID `json:"authorID" db:"author_id" required:"true"` // author_id
}

// BookAuthorCreateParams represents insert params for 'public.book_authors'
type BookAuthorCreateParams struct {
	BookID   int       `json:"bookID"`   // book_id
	AuthorID uuid.UUID `json:"authorID"` // author_id
}

// CreateBookAuthor creates a new BookAuthor in the database with the given params.
func CreateBookAuthor(ctx context.Context, db DB, params *BookAuthorCreateParams) (*BookAuthor, error) {
	ba := &BookAuthor{
		BookID:   params.BookID,
		AuthorID: params.AuthorID,
	}

	return ba.Insert(ctx, db)
}

// BookAuthorUpdateParams represents update params for 'public.book_authors'
type BookAuthorUpdateParams struct {
	BookID   *int       `json:"bookID"`   // book_id
	AuthorID *uuid.UUID `json:"authorID"` // author_id
}

// SetUpdateParams updates public.book_authors struct fields with the specified params.
func (ba *BookAuthor) SetUpdateParams(params *BookAuthorUpdateParams) {
	if params.BookID != nil {
		ba.BookID = *params.BookID
	}
	if params.AuthorID != nil {
		ba.AuthorID = *params.AuthorID
	}
}

type BookAuthorSelectConfig struct {
	limit   string
	orderBy string
	joins   BookAuthorJoins
}
type BookAuthorSelectConfigOption func(*BookAuthorSelectConfig)

// WithBookAuthorLimit limits row selection.
func WithBookAuthorLimit(limit int) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type BookAuthorOrderBy = string

type BookAuthorJoins struct{}

// WithBookAuthorJoin joins with the given tables.
func WithBookAuthorJoin(joins BookAuthorJoins) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the BookAuthor to the database.
func (ba *BookAuthor) Insert(ctx context.Context, db DB) (*BookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.book_authors (` +
		`book_id, author_id` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, ba.BookID, ba.AuthorID)
	rows, err := db.Query(ctx, sqlstr, ba.BookID, ba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Insert/db.Query: %w", err))
	}
	newba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Insert/pgx.CollectOneRow: %w", err))
	}
	*ba = newba

	return ba, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the BookAuthor from the database.
func (ba *BookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.book_authors ` +
		`WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ba.BookID, ba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookAuthorByBookIDAuthorID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID int, authorID uuid.UUID, opts ...BookAuthorSelectConfigOption) (*BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.book_id = $1 AND book_authors.author_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, bookID, authorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", err))
	}
	ba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", err))
	}

	return &ba, nil
}

// BookAuthorsByBookID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorsByBookID(ctx context.Context, db DB, bookID int, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.book_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, bookID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsByAuthorID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorsByAuthorID(ctx context.Context, db DB, authorID uuid.UUID, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.author_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, authorID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}
