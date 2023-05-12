package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookSeller represents a row from 'xo_tests.book_sellers'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type BookSeller struct {
	BookID int       `json:"bookID" db:"book_id" required:"true"` // book_id
	Seller uuid.UUID `json:"seller" db:"seller" required:"true"`  // seller

	SellersJoin     *[]User `json:"-" db:"book_sellers_sellers" openapi-go:"ignore"` // M2M
	BooksJoinSeller *[]Book `json:"-" db:"book_sellers_books" openapi-go:"ignore"`   // M2M
}

// BookSellerCreateParams represents insert params for 'xo_tests.book_sellers'.
type BookSellerCreateParams struct {
	BookID int       `json:"bookID" required:"true"` // book_id
	Seller uuid.UUID `json:"seller" required:"true"` // seller
}

// CreateBookSeller creates a new BookSeller in the database with the given params.
func CreateBookSeller(ctx context.Context, db DB, params *BookSellerCreateParams) (*BookSeller, error) {
	bs := &BookSeller{
		BookID: params.BookID,
		Seller: params.Seller,
	}

	return bs.Insert(ctx, db)
}

// BookSellerUpdateParams represents update params for 'xo_tests.book_sellers'
type BookSellerUpdateParams struct {
	BookID *int       `json:"bookID" required:"true"` // book_id
	Seller *uuid.UUID `json:"seller" required:"true"` // seller
}

// SetUpdateParams updates xo_tests.book_sellers struct fields with the specified params.
func (bs *BookSeller) SetUpdateParams(params *BookSellerUpdateParams) {
	if params.BookID != nil {
		bs.BookID = *params.BookID
	}
	if params.Seller != nil {
		bs.Seller = *params.Seller
	}
}

type BookSellerSelectConfig struct {
	limit   string
	orderBy string
	joins   BookSellerJoins
}
type BookSellerSelectConfigOption func(*BookSellerSelectConfig)

// WithBookSellerLimit limits row selection.
func WithBookSellerLimit(limit int) BookSellerSelectConfigOption {
	return func(s *BookSellerSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type BookSellerOrderBy = string

type BookSellerJoins struct {
	Sellers     bool
	BooksSeller bool
}

// WithBookSellerJoin joins with the given tables.
func WithBookSellerJoin(joins BookSellerJoins) BookSellerSelectConfigOption {
	return func(s *BookSellerSelectConfig) {
		s.joins = BookSellerJoins{
			Sellers:     s.joins.Sellers || joins.Sellers,
			BooksSeller: s.joins.BooksSeller || joins.BooksSeller,
		}
	}
}

// Insert inserts the BookSeller to the database.
func (bs *BookSeller) Insert(ctx context.Context, db DB) (*BookSeller, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_sellers (` +
		`book_id, seller` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, bs.BookID, bs.Seller)
	rows, err := db.Query(ctx, sqlstr, bs.BookID, bs.Seller)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/Insert/db.Query: %w", err))
	}
	newbs, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/Insert/pgx.CollectOneRow: %w", err))
	}
	*bs = newbs

	return bs, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the BookSeller from the database.
func (bs *BookSeller) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_sellers ` +
		`WHERE book_id = $1 AND seller = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, bs.BookID, bs.Seller); err != nil {
		return logerror(err)
	}
	return nil
}

// BookSellerByBookIDSeller retrieves a row from 'xo_tests.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_pkey'.
func BookSellerByBookIDSeller(ctx context.Context, db DB, bookID int, seller uuid.UUID, opts ...BookSellerSelectConfigOption) (*BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_sellers.book_id,
book_sellers.seller,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users is not null), '{}') end) as book_sellers_sellers,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books ` +
		`FROM xo_tests.book_sellers ` +
		`-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
			book_sellers.book_id as book_sellers_book_id
			, row(users.*) as __users
		from
			xo_tests.book_sellers
    join xo_tests.users on users.user_id = book_sellers.seller
    group by
			book_sellers_book_id
			, users.user_id
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = book_sellers.book_id

-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = book_sellers.seller
` +
		` WHERE book_sellers.book_id = $3 AND book_sellers.seller = $4 GROUP BY book_sellers.book_id, book_sellers.book_id, book_sellers.seller, 
book_sellers.seller, book_sellers.book_id, book_sellers.seller `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID, seller)
	rows, err := db.Query(ctx, sqlstr, c.joins.Sellers, c.joins.BooksSeller, bookID, seller)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_sellers/BookSellerByBookIDSeller/db.Query: %w", err))
	}
	bs, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_sellers/BookSellerByBookIDSeller/pgx.CollectOneRow: %w", err))
	}

	return &bs, nil
}

// BookSellersByBookID retrieves a row from 'xo_tests.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_pkey'.
func BookSellersByBookID(ctx context.Context, db DB, bookID int, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_sellers.book_id,
book_sellers.seller,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users is not null), '{}') end) as book_sellers_sellers,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books ` +
		`FROM xo_tests.book_sellers ` +
		`-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
			book_sellers.book_id as book_sellers_book_id
			, row(users.*) as __users
		from
			xo_tests.book_sellers
    join xo_tests.users on users.user_id = book_sellers.seller
    group by
			book_sellers_book_id
			, users.user_id
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = book_sellers.book_id

-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = book_sellers.seller
` +
		` WHERE book_sellers.book_id = $3 GROUP BY book_sellers.book_id, book_sellers.book_id, book_sellers.seller, 
book_sellers.seller, book_sellers.book_id, book_sellers.seller `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Sellers, c.joins.BooksSeller, bookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookSellersBySeller retrieves a row from 'xo_tests.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_pkey'.
func BookSellersBySeller(ctx context.Context, db DB, seller uuid.UUID, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_sellers.book_id,
book_sellers.seller,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users is not null), '{}') end) as book_sellers_sellers,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG((
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books ` +
		`FROM xo_tests.book_sellers ` +
		`-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
			book_sellers.book_id as book_sellers_book_id
			, row(users.*) as __users
		from
			xo_tests.book_sellers
    join xo_tests.users on users.user_id = book_sellers.seller
    group by
			book_sellers_book_id
			, users.user_id
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = book_sellers.book_id

-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = book_sellers.seller
` +
		` WHERE book_sellers.seller = $3 GROUP BY book_sellers.book_id, book_sellers.book_id, book_sellers.seller, 
book_sellers.seller, book_sellers.book_id, book_sellers.seller `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, seller)
	rows, err := db.Query(ctx, sqlstr, c.joins.Sellers, c.joins.BooksSeller, seller)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKBook_BookID returns the Book associated with the BookSeller's (BookID).
//
// Generated from foreign key 'book_sellers_book_id_fkey'.
func (bs *BookSeller) FKBook_BookID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, bs.BookID)
}

// FKUser_Seller returns the User associated with the BookSeller's (Seller).
//
// Generated from foreign key 'book_sellers_seller_fkey'.
func (bs *BookSeller) FKUser_Seller(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, bs.Seller)
}