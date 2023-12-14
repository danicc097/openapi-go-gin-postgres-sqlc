package schema_extra_schema

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// BookSeller represents a row from 'extra_schema.book_sellers'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type BookSeller struct {
	BookID BookID `json:"bookID" db:"book_id" required:"true" nullable:"false"` // book_id
	Seller UserID `json:"seller" db:"seller" required:"true" nullable:"false"`  // seller

	BookSellersJoin *[]User `json:"-" db:"book_sellers_sellers" openapi-go:"ignore"` // M2M book_sellers
	SellerBooksJoin *[]Book `json:"-" db:"book_sellers_books" openapi-go:"ignore"`   // M2M book_sellers

}

// BookSellerCreateParams represents insert params for 'extra_schema.book_sellers'.
type BookSellerCreateParams struct {
	BookID BookID `json:"bookID" required:"true" nullable:"false"` // book_id
	Seller UserID `json:"seller" required:"true" nullable:"false"` // seller
}

// CreateBookSeller creates a new BookSeller in the database with the given params.
func CreateBookSeller(ctx context.Context, db DB, params *BookSellerCreateParams) (*BookSeller, error) {
	bs := &BookSeller{
		BookID: params.BookID,
		Seller: params.Seller,
	}

	return bs.Insert(ctx, db)
}

// BookSellerUpdateParams represents update params for 'extra_schema.book_sellers'.
type BookSellerUpdateParams struct {
	BookID *BookID `json:"bookID" nullable:"false"` // book_id
	Seller *UserID `json:"seller" nullable:"false"` // seller
}

// SetUpdateParams updates extra_schema.book_sellers struct fields with the specified params.
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
	filters map[string][]any
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

type BookSellerOrderBy string

const ()

type BookSellerJoins struct {
	Sellers     bool // M2M book_sellers
	BooksSeller bool // M2M book_sellers
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

// WithBookSellerFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithBookSellerFilters(filters map[string][]any) BookSellerSelectConfigOption {
	return func(s *BookSellerSelectConfig) {
		s.filters = filters
	}
}

const bookSellerTableSellersJoinSQL = `-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
		book_sellers.book_id as book_sellers_book_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.book_sellers
	join extra_schema.users on users.user_id = book_sellers.seller
	group by
		book_sellers_book_id
		, users.user_id
) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = book_sellers.book_id
`

const bookSellerTableSellersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users_user_id is not null), '{}') as book_sellers_sellers`

const bookSellerTableSellersGroupBySQL = `book_sellers.book_id, book_sellers.book_id, book_sellers.seller`

const bookSellerTableBooksSellerJoinSQL = `-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
		book_sellers.seller as book_sellers_seller
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		extra_schema.book_sellers
	join extra_schema.books on books.book_id = book_sellers.book_id
	group by
		book_sellers_seller
		, books.book_id
) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = book_sellers.seller
`

const bookSellerTableBooksSellerSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books_book_id is not null), '{}') as book_sellers_books`

const bookSellerTableBooksSellerGroupBySQL = `book_sellers.seller, book_sellers.book_id, book_sellers.seller`

// Insert inserts the BookSeller to the database.
func (bs *BookSeller) Insert(ctx context.Context, db DB) (*BookSeller, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.book_sellers (
	book_id, seller
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, bs.BookID, bs.Seller)
	rows, err := db.Query(ctx, sqlstr, bs.BookID, bs.Seller)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/Insert/db.Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	newbs, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	*bs = newbs

	return bs, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the BookSeller from the database.
func (bs *BookSeller) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM extra_schema.book_sellers 
	WHERE book_id = $1 AND seller = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, bs.BookID, bs.Seller); err != nil {
		return logerror(err)
	}
	return nil
}

// BookSellersByBookIDSeller retrieves a row from 'extra_schema.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_book_id_seller_idx'.
func BookSellersByBookIDSeller(ctx context.Context, db DB, bookID BookID, seller UserID, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, bookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableSellersGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, bookSellerTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableBooksSellerGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM extra_schema.book_sellers %s 
	 WHERE book_sellers.book_id = $1 AND book_sellers.seller = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookSellersByBookIDSeller */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, seller)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, seller}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellersByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellersByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// BookSellersByBookID retrieves a row from 'extra_schema.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_pkey'.
func BookSellersByBookID(ctx context.Context, db DB, bookID BookID, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, bookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableSellersGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, bookSellerTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableBooksSellerGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM extra_schema.book_sellers %s 
	 WHERE book_sellers.book_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookSellersByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// BookSellersBySeller retrieves a row from 'extra_schema.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_pkey'.
func BookSellersBySeller(ctx context.Context, db DB, seller UserID, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, bookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableSellersGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, bookSellerTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableBooksSellerGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM extra_schema.book_sellers %s 
	 WHERE book_sellers.seller = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookSellersBySeller */\n" + sqlstr

	// run
	// logf(sqlstr, seller)
	rows, err := db.Query(ctx, sqlstr, append([]any{seller}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// BookSellersBySellerBookID retrieves a row from 'extra_schema.book_sellers' as a BookSeller.
//
// Generated from index 'book_sellers_seller_book_id_idx'.
func BookSellersBySellerBookID(ctx context.Context, db DB, seller UserID, bookID BookID, opts ...BookSellerSelectConfigOption) ([]BookSeller, error) {
	c := &BookSellerSelectConfig{joins: BookSellerJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, bookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableSellersGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, bookSellerTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, bookSellerTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, bookSellerTableBooksSellerGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM extra_schema.book_sellers %s 
	 WHERE book_sellers.seller = $1 AND book_sellers.book_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookSellersBySellerBookID */\n" + sqlstr

	// run
	// logf(sqlstr, seller, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{seller, bookID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellersBySellerBookID/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookSeller/BookSellersBySellerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
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
