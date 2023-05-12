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

// BookAuthorsSurrogateKey represents a row from 'xo_tests.book_authors_surrogate_key'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type BookAuthorsSurrogateKey struct {
	BookAuthorsSurrogateKeyID int       `json:"bookAuthorsSurrogateKeyID" db:"book_authors_surrogate_key_id" required:"true"` // book_authors_surrogate_key_id
	BookSurrID                int       `json:"bookSurrID" db:"book_surr_id" required:"true"`                                 // book_surr_id
	AuthorSurrID              uuid.UUID `json:"authorSurrID" db:"author_surr_id" required:"true"`                             // author_surr_id
	Pseudonym                 *string   `json:"pseudonym" db:"pseudonym" required:"true"`                                     // pseudonym

	BookSurrsJoin   *[]BookAuthorsSurrogateKey_BookSurr   `json:"-" db:"book_authors_surrogate_key_book_surrs" openapi-go:"ignore"`   // M2M
	AuthorSurrsJoin *[]BookAuthorsSurrogateKey_AuthorSurr `json:"-" db:"book_authors_surrogate_key_author_surrs" openapi-go:"ignore"` // M2M
}

// BookAuthorsSurrogateKeyCreateParams represents insert params for 'xo_tests.book_authors_surrogate_key'.
type BookAuthorsSurrogateKeyCreateParams struct {
	BookSurrID   int       `json:"bookSurrID" required:"true"`   // book_surr_id
	AuthorSurrID uuid.UUID `json:"authorSurrID" required:"true"` // author_surr_id
	Pseudonym    *string   `json:"pseudonym" required:"true"`    // pseudonym
}

// CreateBookAuthorsSurrogateKey creates a new BookAuthorsSurrogateKey in the database with the given params.
func CreateBookAuthorsSurrogateKey(ctx context.Context, db DB, params *BookAuthorsSurrogateKeyCreateParams) (*BookAuthorsSurrogateKey, error) {
	bask := &BookAuthorsSurrogateKey{
		BookSurrID:   params.BookSurrID,
		AuthorSurrID: params.AuthorSurrID,
		Pseudonym:    params.Pseudonym,
	}

	return bask.Insert(ctx, db)
}

// BookAuthorsSurrogateKeyUpdateParams represents update params for 'xo_tests.book_authors_surrogate_key'
type BookAuthorsSurrogateKeyUpdateParams struct {
	BookSurrID   *int       `json:"bookSurrID" required:"true"`   // book_surr_id
	AuthorSurrID *uuid.UUID `json:"authorSurrID" required:"true"` // author_surr_id
	Pseudonym    **string   `json:"pseudonym" required:"true"`    // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors_surrogate_key struct fields with the specified params.
func (bask *BookAuthorsSurrogateKey) SetUpdateParams(params *BookAuthorsSurrogateKeyUpdateParams) {
	if params.BookSurrID != nil {
		bask.BookSurrID = *params.BookSurrID
	}
	if params.AuthorSurrID != nil {
		bask.AuthorSurrID = *params.AuthorSurrID
	}
	if params.Pseudonym != nil {
		bask.Pseudonym = *params.Pseudonym
	}
}

type BookAuthorsSurrogateKeySelectConfig struct {
	limit   string
	orderBy string
	joins   BookAuthorsSurrogateKeyJoins
}
type BookAuthorsSurrogateKeySelectConfigOption func(*BookAuthorsSurrogateKeySelectConfig)

// WithBookAuthorsSurrogateKeyLimit limits row selection.
func WithBookAuthorsSurrogateKeyLimit(limit int) BookAuthorsSurrogateKeySelectConfigOption {
	return func(s *BookAuthorsSurrogateKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type BookAuthorsSurrogateKeyOrderBy = string

type BookAuthorsSurrogateKeyJoins struct {
	BookSurrs   bool
	AuthorSurrs bool
}

// WithBookAuthorsSurrogateKeyJoin joins with the given tables.
func WithBookAuthorsSurrogateKeyJoin(joins BookAuthorsSurrogateKeyJoins) BookAuthorsSurrogateKeySelectConfigOption {
	return func(s *BookAuthorsSurrogateKeySelectConfig) {
		s.joins = BookAuthorsSurrogateKeyJoins{
			BookSurrs:   s.joins.BookSurrs || joins.BookSurrs,
			AuthorSurrs: s.joins.AuthorSurrs || joins.AuthorSurrs,
		}
	}
}

// BookAuthorsSurrogateKey_BookSurr represents a M2M join against "xo_tests.book_authors_surrogate_key"
type BookAuthorsSurrogateKey_BookSurr struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true"`
}

// BookAuthorsSurrogateKey_AuthorSurr represents a M2M join against "xo_tests.book_authors_surrogate_key"
type BookAuthorsSurrogateKey_AuthorSurr struct {
	User      User    `json:"user" db:"users" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true"`
}

// Insert inserts the BookAuthorsSurrogateKey to the database.
func (bask *BookAuthorsSurrogateKey) Insert(ctx context.Context, db DB) (*BookAuthorsSurrogateKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.book_authors_surrogate_key (` +
		`book_surr_id, author_surr_id, pseudonym` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, bask.BookSurrID, bask.AuthorSurrID, bask.Pseudonym)

	rows, err := db.Query(ctx, sqlstr, bask.BookSurrID, bask.AuthorSurrID, bask.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Insert/db.Query: %w", err))
	}
	newbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Insert/pgx.CollectOneRow: %w", err))
	}

	*bask = newbask

	return bask, nil
}

// Update updates a BookAuthorsSurrogateKey in the database.
func (bask *BookAuthorsSurrogateKey) Update(ctx context.Context, db DB) (*BookAuthorsSurrogateKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors_surrogate_key SET ` +
		`book_surr_id = $1, author_surr_id = $2, pseudonym = $3 ` +
		`WHERE book_authors_surrogate_key_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, bask.BookSurrID, bask.AuthorSurrID, bask.Pseudonym, bask.BookAuthorsSurrogateKeyID)

	rows, err := db.Query(ctx, sqlstr, bask.BookSurrID, bask.AuthorSurrID, bask.Pseudonym, bask.BookAuthorsSurrogateKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Update/db.Query: %w", err))
	}
	newbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Update/pgx.CollectOneRow: %w", err))
	}
	*bask = newbask

	return bask, nil
}

// Upsert upserts a BookAuthorsSurrogateKey in the database.
// Requires appropiate PK(s) to be set beforehand.
func (bask *BookAuthorsSurrogateKey) Upsert(ctx context.Context, db DB, params *BookAuthorsSurrogateKeyCreateParams) (*BookAuthorsSurrogateKey, error) {
	var err error

	bask.BookSurrID = params.BookSurrID
	bask.AuthorSurrID = params.AuthorSurrID
	bask.Pseudonym = params.Pseudonym

	bask, err = bask.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			bask, err = bask.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return bask, err
}

// Delete deletes the BookAuthorsSurrogateKey from the database.
func (bask *BookAuthorsSurrogateKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.book_authors_surrogate_key ` +
		`WHERE book_authors_surrogate_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, bask.BookAuthorsSurrogateKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyID returns a cursor-paginated list of BookAuthorsSurrogateKey.
func BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyID(ctx context.Context, db DB, bookAuthorsSurrogateKeyID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_surr_id,
book_authors_surrogate_key.author_surr_id,
book_authors_surrogate_key.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_book_surrs.__books
		, joined_book_authors_surrogate_key_book_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_book_surrs.__books is not null), '{}') end) as book_authors_surrogate_key_book_surrs,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_author_surrs.__users
		, joined_book_authors_surrogate_key_author_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_author_surrs.__users is not null), '{}') end) as book_authors_surrogate_key_author_surrs ` +
		`FROM xo_tests.book_authors_surrogate_key ` +
		`-- M2M join generated from "book_authors_surrogate_key_book_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_surr_id as book_authors_surrogate_key_author_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_surr_id
    group by
			book_authors_surrogate_key_author_surr_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_book_surrs on joined_book_authors_surrogate_key_book_surrs.book_authors_surrogate_key_author_surr_id = book_authors_surrogate_key.author_surr_id

-- M2M join generated from "book_authors_surrogate_key_author_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.book_surr_id as book_authors_surrogate_key_book_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.users on users.user_id = book_authors_surrogate_key.author_surr_id
    group by
			book_authors_surrogate_key_book_surr_id
			, users.user_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_author_surrs on joined_book_authors_surrogate_key_author_surrs.book_authors_surrogate_key_book_surr_id = book_authors_surrogate_key.book_surr_id
` +
		` WHERE book_authors_surrogate_key.book_authors_surrogate_key_id > $3 GROUP BY book_authors_surrogate_key.author_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, bookAuthorsSurrogateKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_surr_id_author_surr_id_key'.
func BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID(ctx context.Context, db DB, bookSurrID int, authorSurrID uuid.UUID, opts ...BookAuthorsSurrogateKeySelectConfigOption) (*BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_surr_id,
book_authors_surrogate_key.author_surr_id,
book_authors_surrogate_key.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_book_surrs.__books
		, joined_book_authors_surrogate_key_book_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_book_surrs.__books is not null), '{}') end) as book_authors_surrogate_key_book_surrs,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_author_surrs.__users
		, joined_book_authors_surrogate_key_author_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_author_surrs.__users is not null), '{}') end) as book_authors_surrogate_key_author_surrs ` +
		`FROM xo_tests.book_authors_surrogate_key ` +
		`-- M2M join generated from "book_authors_surrogate_key_book_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_surr_id as book_authors_surrogate_key_author_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_surr_id
    group by
			book_authors_surrogate_key_author_surr_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_book_surrs on joined_book_authors_surrogate_key_book_surrs.book_authors_surrogate_key_author_surr_id = book_authors_surrogate_key.author_surr_id

-- M2M join generated from "book_authors_surrogate_key_author_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.book_surr_id as book_authors_surrogate_key_book_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.users on users.user_id = book_authors_surrogate_key.author_surr_id
    group by
			book_authors_surrogate_key_book_surr_id
			, users.user_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_author_surrs on joined_book_authors_surrogate_key_author_surrs.book_authors_surrogate_key_book_surr_id = book_authors_surrogate_key.book_surr_id
` +
		` WHERE book_authors_surrogate_key.book_surr_id = $3 AND book_authors_surrogate_key.author_surr_id = $4 GROUP BY book_authors_surrogate_key.author_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookSurrID, authorSurrID)
	rows, err := db.Query(ctx, sqlstr, c.joins.BookSurrs, c.joins.AuthorSurrs, bookSurrID, authorSurrID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/db.Query: %w", err))
	}
	bask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/pgx.CollectOneRow: %w", err))
	}

	return &bask, nil
}

// BookAuthorsSurrogateKeysByBookSurrID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_surr_id_author_surr_id_key'.
func BookAuthorsSurrogateKeysByBookSurrID(ctx context.Context, db DB, bookSurrID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_surr_id,
book_authors_surrogate_key.author_surr_id,
book_authors_surrogate_key.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_book_surrs.__books
		, joined_book_authors_surrogate_key_book_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_book_surrs.__books is not null), '{}') end) as book_authors_surrogate_key_book_surrs,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_author_surrs.__users
		, joined_book_authors_surrogate_key_author_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_author_surrs.__users is not null), '{}') end) as book_authors_surrogate_key_author_surrs ` +
		`FROM xo_tests.book_authors_surrogate_key ` +
		`-- M2M join generated from "book_authors_surrogate_key_book_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_surr_id as book_authors_surrogate_key_author_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_surr_id
    group by
			book_authors_surrogate_key_author_surr_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_book_surrs on joined_book_authors_surrogate_key_book_surrs.book_authors_surrogate_key_author_surr_id = book_authors_surrogate_key.author_surr_id

-- M2M join generated from "book_authors_surrogate_key_author_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.book_surr_id as book_authors_surrogate_key_book_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.users on users.user_id = book_authors_surrogate_key.author_surr_id
    group by
			book_authors_surrogate_key_book_surr_id
			, users.user_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_author_surrs on joined_book_authors_surrogate_key_author_surrs.book_authors_surrogate_key_book_surr_id = book_authors_surrogate_key.book_surr_id
` +
		` WHERE book_authors_surrogate_key.book_surr_id = $3 GROUP BY book_authors_surrogate_key.author_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookSurrID)
	rows, err := db.Query(ctx, sqlstr, c.joins.BookSurrs, c.joins.AuthorSurrs, bookSurrID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeysByAuthorSurrID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_surr_id_author_surr_id_key'.
func BookAuthorsSurrogateKeysByAuthorSurrID(ctx context.Context, db DB, authorSurrID uuid.UUID, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_surr_id,
book_authors_surrogate_key.author_surr_id,
book_authors_surrogate_key.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_book_surrs.__books
		, joined_book_authors_surrogate_key_book_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_book_surrs.__books is not null), '{}') end) as book_authors_surrogate_key_book_surrs,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_author_surrs.__users
		, joined_book_authors_surrogate_key_author_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_author_surrs.__users is not null), '{}') end) as book_authors_surrogate_key_author_surrs ` +
		`FROM xo_tests.book_authors_surrogate_key ` +
		`-- M2M join generated from "book_authors_surrogate_key_book_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_surr_id as book_authors_surrogate_key_author_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_surr_id
    group by
			book_authors_surrogate_key_author_surr_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_book_surrs on joined_book_authors_surrogate_key_book_surrs.book_authors_surrogate_key_author_surr_id = book_authors_surrogate_key.author_surr_id

-- M2M join generated from "book_authors_surrogate_key_author_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.book_surr_id as book_authors_surrogate_key_book_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.users on users.user_id = book_authors_surrogate_key.author_surr_id
    group by
			book_authors_surrogate_key_book_surr_id
			, users.user_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_author_surrs on joined_book_authors_surrogate_key_author_surrs.book_authors_surrogate_key_book_surr_id = book_authors_surrogate_key.book_surr_id
` +
		` WHERE book_authors_surrogate_key.author_surr_id = $3 GROUP BY book_authors_surrogate_key.author_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, authorSurrID)
	rows, err := db.Query(ctx, sqlstr, c.joins.BookSurrs, c.joins.AuthorSurrs, authorSurrID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookSurrIDAuthorSurrID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_pkey'.
func BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID(ctx context.Context, db DB, bookAuthorsSurrogateKeyID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) (*BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_surr_id,
book_authors_surrogate_key.author_surr_id,
book_authors_surrogate_key.pseudonym,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_book_surrs.__books
		, joined_book_authors_surrogate_key_book_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_book_surrs.__books is not null), '{}') end) as book_authors_surrogate_key_book_surrs,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_authors_surrogate_key_author_surrs.__users
		, joined_book_authors_surrogate_key_author_surrs.pseudonym
		)) filter (where joined_book_authors_surrogate_key_author_surrs.__users is not null), '{}') end) as book_authors_surrogate_key_author_surrs ` +
		`FROM xo_tests.book_authors_surrogate_key ` +
		`-- M2M join generated from "book_authors_surrogate_key_book_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_surr_id as book_authors_surrogate_key_author_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_surr_id
    group by
			book_authors_surrogate_key_author_surr_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_book_surrs on joined_book_authors_surrogate_key_book_surrs.book_authors_surrogate_key_author_surr_id = book_authors_surrogate_key.author_surr_id

-- M2M join generated from "book_authors_surrogate_key_author_surr_id_fkey"
left join (
	select
			book_authors_surrogate_key.book_surr_id as book_authors_surrogate_key_book_surr_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(users.*) as __users
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.users on users.user_id = book_authors_surrogate_key.author_surr_id
    group by
			book_authors_surrogate_key_book_surr_id
			, users.user_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_author_surrs on joined_book_authors_surrogate_key_author_surrs.book_authors_surrogate_key_book_surr_id = book_authors_surrogate_key.book_surr_id
` +
		` WHERE book_authors_surrogate_key.book_authors_surrogate_key_id = $3 GROUP BY book_authors_surrogate_key.author_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_surr_id, book_authors_surrogate_key.book_authors_surrogate_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookAuthorsSurrogateKeyID)
	rows, err := db.Query(ctx, sqlstr, c.joins.BookSurrs, c.joins.AuthorSurrs, bookAuthorsSurrogateKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/db.Query: %w", err))
	}
	bask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/pgx.CollectOneRow: %w", err))
	}

	return &bask, nil
}

// FKUser_AuthorSurrID returns the User associated with the BookAuthorsSurrogateKey's (AuthorSurrID).
//
// Generated from foreign key 'book_authors_surrogate_key_author_surr_id_fkey'.
func (bask *BookAuthorsSurrogateKey) FKUser_AuthorSurrID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, bask.AuthorSurrID)
}

// FKBook_BookSurrID returns the Book associated with the BookAuthorsSurrogateKey's (BookSurrID).
//
// Generated from foreign key 'book_authors_surrogate_key_book_surr_id_fkey'.
func (bask *BookAuthorsSurrogateKey) FKBook_BookSurrID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, bask.BookSurrID)
}