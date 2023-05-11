package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// BookAuthor represents a row from 'xo_tests.book_authors'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type BookAuthor struct {
	BookID    int       `json:"bookID" db:"book_id" required:"true"`      // book_id
	AuthorID  uuid.UUID `json:"authorID" db:"author_id" required:"true"`  // author_id
	Pseudonym *string   `json:"pseudonym" db:"pseudonym" required:"true"` // pseudonym

	BooksJoin   *[]BookAuthor_Book   `json:"-" db:"books" openapi-go:"ignore"`   // M2M
	AuthorsJoin *[]BookAuthor_Author `json:"-" db:"authors" openapi-go:"ignore"` // M2M
}

// BookAuthorCreateParams represents insert params for 'xo_tests.book_authors'.
type BookAuthorCreateParams struct {
	BookID    int       `json:"bookID" required:"true"`    // book_id
	AuthorID  uuid.UUID `json:"authorID" required:"true"`  // author_id
	Pseudonym *string   `json:"pseudonym" required:"true"` // pseudonym
}

// CreateBookAuthor creates a new BookAuthor in the database with the given params.
func CreateBookAuthor(ctx context.Context, db DB, params *BookAuthorCreateParams) (*BookAuthor, error) {
	ba := &BookAuthor{
		BookID:    params.BookID,
		AuthorID:  params.AuthorID,
		Pseudonym: params.Pseudonym,
	}

	return ba.Insert(ctx, db)
}

// BookAuthorUpdateParams represents update params for 'xo_tests.book_authors'
type BookAuthorUpdateParams struct {
	BookID    *int       `json:"bookID" required:"true"`    // book_id
	AuthorID  *uuid.UUID `json:"authorID" required:"true"`  // author_id
	Pseudonym **string   `json:"pseudonym" required:"true"` // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors struct fields with the specified params.
func (ba *BookAuthor) SetUpdateParams(params *BookAuthorUpdateParams) {
	if params.BookID != nil {
		ba.BookID = *params.BookID
	}
	if params.AuthorID != nil {
		ba.AuthorID = *params.AuthorID
	}
	if params.Pseudonym != nil {
		ba.Pseudonym = *params.Pseudonym
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
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type BookAuthorOrderBy = string

type BookAuthorJoins struct {
	Books   bool
	Authors bool
}

// WithBookAuthorJoin joins with the given tables.
func WithBookAuthorJoin(joins BookAuthorJoins) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.joins = BookAuthorJoins{
			Books:   s.joins.Books || joins.Books,
			Authors: s.joins.Authors || joins.Authors,
		}
	}
}

type BookAuthor_Book struct {
	Book      Book    `json:"book" db:"books"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true"`
}

type BookAuthor_Author struct {
	User      User    `json:"user" db:"users"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true"`
}

// Insert inserts the BookAuthor to the database.
func (ba *BookAuthor) Insert(ctx context.Context, db DB) (*BookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_authors (` +
		`book_id, author_id, pseudonym` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, ba.BookID, ba.AuthorID, ba.Pseudonym)
	rows, err := db.Query(ctx, sqlstr, ba.BookID, ba.AuthorID, ba.Pseudonym)
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

// Update updates a BookAuthor in the database.
func (ba *BookAuthor) Update(ctx context.Context, db DB) (*BookAuthor, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors SET ` +
		`pseudonym = $1 ` +
		`WHERE book_id = $2  AND author_id = $3 ` +
		`RETURNING * `
	// run
	logf(sqlstr, ba.Pseudonym, ba.BookID, ba.AuthorID)

	rows, err := db.Query(ctx, sqlstr, ba.Pseudonym, ba.BookID, ba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Update/db.Query: %w", err))
	}
	newba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Update/pgx.CollectOneRow: %w", err))
	}
	*ba = newba

	return ba, nil
}

// Upsert upserts a BookAuthor in the database.
// Requires appropiate PK(s) to be set beforehand.
func (ba *BookAuthor) Upsert(ctx context.Context, db DB, params *BookAuthorCreateParams) (*BookAuthor, error) {
	var err error

	ba.BookID = params.BookID
	ba.AuthorID = params.AuthorID
	ba.Pseudonym = params.Pseudonym

	ba, err = ba.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			ba, err = ba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return ba, err
}

// Delete deletes the BookAuthor from the database.
func (ba *BookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_authors ` +
		`WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ba.BookID, ba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookAuthorByBookIDAuthorID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
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
book_authors.author_id,
book_authors.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_books.__books
		, joined_books.pseudonym
		)) filter (where joined_books.__books is not null), '{}') end) as books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_authors.__users
		, joined_authors.pseudonym
		)) filter (where joined_authors.__users is not null), '{}') end) as authors ` +
		`FROM xo_tests.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, book_authors.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors
    join xo_tests.users on users.user_id = book_authors.author_id
    group by
			book_authors_book_id
			, users.user_id
			, pseudonym
  ) as joined_authors on joined_authors.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.book_id = $3 AND book_authors.author_id = $4 GROUP BY book_authors.author_id, book_authors.book_id, book_authors.author_id, 
book_authors.book_id, book_authors.book_id, book_authors.author_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Books, c.joins.Authors, bookID, authorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", err))
	}
	ba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", err))
	}

	return &ba, nil
}

// BookAuthorsByBookID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
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
book_authors.author_id,
book_authors.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_books.__books
		, joined_books.pseudonym
		)) filter (where joined_books.__books is not null), '{}') end) as books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_authors.__users
		, joined_authors.pseudonym
		)) filter (where joined_authors.__users is not null), '{}') end) as authors ` +
		`FROM xo_tests.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, book_authors.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors
    join xo_tests.users on users.user_id = book_authors.author_id
    group by
			book_authors_book_id
			, users.user_id
			, pseudonym
  ) as joined_authors on joined_authors.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.book_id = $3 GROUP BY book_authors.author_id, book_authors.book_id, book_authors.author_id, 
book_authors.book_id, book_authors.book_id, book_authors.author_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Books, c.joins.Authors, bookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsByAuthorID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
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
book_authors.author_id,
book_authors.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_books.__books
		, joined_books.pseudonym
		)) filter (where joined_books.__books is not null), '{}') end) as books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_authors.__users
		, joined_authors.pseudonym
		)) filter (where joined_authors.__users is not null), '{}') end) as authors ` +
		`FROM xo_tests.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, book_authors.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors
    join xo_tests.users on users.user_id = book_authors.author_id
    group by
			book_authors_book_id
			, users.user_id
			, pseudonym
  ) as joined_authors on joined_authors.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.author_id = $3 GROUP BY book_authors.author_id, book_authors.book_id, book_authors.author_id, 
book_authors.book_id, book_authors.book_id, book_authors.author_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Books, c.joins.Authors, authorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKUser_AuthorID returns the User associated with the BookAuthor's (AuthorID).
//
// Generated from foreign key 'book_authors_author_id_fkey'.
func (ba *BookAuthor) FKUser_AuthorID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ba.AuthorID)
}

// FKBook_BookID returns the Book associated with the BookAuthor's (BookID).
//
// Generated from foreign key 'book_authors_book_id_fkey'.
func (ba *BookAuthor) FKBook_BookID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, ba.BookID)
}
