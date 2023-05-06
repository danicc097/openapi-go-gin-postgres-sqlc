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

	BooksJoin   *[]Book              `json:"-" db:"books" openapi-go:"ignore"`   // M2M
	AuthorsJoin *[]BookAuthor_Author `json:"-" db:"authors" openapi-go:"ignore"` // M2M
}

// BookAuthorCreateParams represents insert params for 'public.book_authors'
type BookAuthorCreateParams struct {
	BookID   int       `json:"bookID" required:"true"`   // book_id
	AuthorID uuid.UUID `json:"authorID" required:"true"` // author_id
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
	BookID   *int       `json:"bookID" required:"true"`   // book_id
	AuthorID *uuid.UUID `json:"authorID" required:"true"` // author_id
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

type BookAuthor_Author struct {
	User User `json:"user" db:"users"`
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

// BookAuthorPaginatedByBookIDAuthorID returns a cursor-paginated list of BookAuthor.
func BookAuthorPaginatedByBookIDAuthorID(ctx context.Context, db DB, bookID int, authorID uuid.UUID, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id,
(case when $1::boolean = true then COALESCE(joined_books.__books, '{}') end) as books,
(case when $2::boolean = true then COALESCE(joined_author_ids.__author_ids, '{}') end) as author_ids ` +
		`FROM public.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, array_agg(books.*) filter (where books.* is not null) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by book_authors_author_id
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, array_agg(users.*) filter (where users.* is not null) as __author_ids
		from book_authors
    	join users on users.user_id = book_authors.author_id
    group by book_authors_book_id
  ) as joined_author_ids on joined_author_ids.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.book_id > $3 AND book_authors.author_id > $4` +
		` ORDER BY 
		book_id DESC ,
		author_id DESC `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, bookID, authorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
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
book_authors.author_id,
(case when $1::boolean = true then COALESCE(joined_books.__books, '{}') end) as books,
(case when $2::boolean = true then COALESCE(joined_author_ids.__author_ids, '{}') end) as author_ids ` +
		`FROM public.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, array_agg(books.*) filter (where books.* is not null) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by book_authors_author_id
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, array_agg(users.*) filter (where users.* is not null) as __author_ids
		from book_authors
    	join users on users.user_id = book_authors.author_id
    group by book_authors_book_id
  ) as joined_author_ids on joined_author_ids.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.book_id = $3 AND book_authors.author_id = $4 `
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
book_authors.author_id,
(case when $1::boolean = true then COALESCE(joined_books.__books, '{}') end) as books,
(case when $2::boolean = true then COALESCE(joined_author_ids.__author_ids, '{}') end) as author_ids ` +
		`FROM public.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, array_agg(books.*) filter (where books.* is not null) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by book_authors_author_id
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, array_agg(users.*) filter (where users.* is not null) as __author_ids
		from book_authors
    	join users on users.user_id = book_authors.author_id
    group by book_authors_book_id
  ) as joined_author_ids on joined_author_ids.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.book_id = $3 `
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
book_authors.author_id,
(case when $1::boolean = true then COALESCE(joined_books.__books, '{}') end) as books,
(case when $2::boolean = true then COALESCE(joined_author_ids.__author_ids, '{}') end) as author_ids ` +
		`FROM public.book_authors ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, array_agg(books.*) filter (where books.* is not null) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by book_authors_author_id
  ) as joined_books on joined_books.book_authors_author_id = book_authors.author_id

-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
			book_authors.book_id as book_authors_book_id
			, array_agg(users.*) filter (where users.* is not null) as __author_ids
		from book_authors
    	join users on users.user_id = book_authors.author_id
    group by book_authors_book_id
  ) as joined_author_ids on joined_author_ids.book_authors_book_id = book_authors.book_id
` +
		` WHERE book_authors.author_id = $3 `
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
