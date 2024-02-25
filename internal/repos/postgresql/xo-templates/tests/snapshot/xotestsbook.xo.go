package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsBook represents a row from 'xo_tests.books'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsBook struct {
	BookID XoTestsBookID `db:"book_id" json:"bookID" nullable:"false" required:"true"` // book_id
	Name   string        `db:"name"    json:"name"   nullable:"false" required:"true"`      // name

	AuthorsJoin     *[]User__BA_XoTestsBook   `db:"book_authors_authors"               json:"-" openapi-go:"ignore"`               // M2M book_authors
	AuthorsBASKJoin *[]User__BASK_XoTestsBook `db:"book_authors_surrogate_key_authors" json:"-" openapi-go:"ignore"` // M2M book_authors_surrogate_key
	BookReviewsJoin *[]XoTestsBookReview      `db:"book_reviews"                       json:"-" openapi-go:"ignore"`                       // M2O books
	SellersJoin     *[]XoTestsUser            `db:"book_sellers_sellers"               json:"-" openapi-go:"ignore"`               // M2M book_sellers
}

// XoTestsBookCreateParams represents insert params for 'xo_tests.books'.
type XoTestsBookCreateParams struct {
	Name string `json:"name" nullable:"false" required:"true"` // name
}

type XoTestsBookID int

// CreateXoTestsBook creates a new XoTestsBook in the database with the given params.
func CreateXoTestsBook(ctx context.Context, db DB, params *XoTestsBookCreateParams) (*XoTestsBook, error) {
	xtb := &XoTestsBook{
		Name: params.Name,
	}

	return xtb.Insert(ctx, db)
}

type XoTestsBookSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsBookJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookSelectConfigOption func(*XoTestsBookSelectConfig)

// WithXoTestsBookLimit limits row selection.
func WithXoTestsBookLimit(limit int) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsBookOrderBy string

type XoTestsBookJoins struct {
	Authors     bool // M2M book_authors
	AuthorsBASK bool // M2M book_authors_surrogate_key
	BookReviews bool // M2O book_reviews
	Sellers     bool // M2M book_sellers
}

// WithXoTestsBookJoin joins with the given tables.
func WithXoTestsBookJoin(joins XoTestsBookJoins) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.joins = XoTestsBookJoins{
			Authors:     s.joins.Authors || joins.Authors,
			AuthorsBASK: s.joins.AuthorsBASK || joins.AuthorsBASK,
			BookReviews: s.joins.BookReviews || joins.BookReviews,
			Sellers:     s.joins.Sellers || joins.Sellers,
		}
	}
}

// User__BA_XoTestsBook represents a M2M join against "xo_tests.book_authors".
type User__BA_XoTestsBook struct {
	User      XoTestsUser `db:"users"     json:"user"      required:"true"`
	Pseudonym *string     `db:"pseudonym" json:"pseudonym" required:"true"`
}

// User__BASK_XoTestsBook represents a M2M join against "xo_tests.book_authors_surrogate_key".
type User__BASK_XoTestsBook struct {
	User      XoTestsUser `db:"users"     json:"user"      required:"true"`
	Pseudonym *string     `db:"pseudonym" json:"pseudonym" required:"true"`
}

// WithXoTestsBookFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookFilters(filters map[string][]any) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsBookHavingClause(conditions map[string][]any) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookTableAuthorsJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
		book_authors.book_id as book_authors_book_id
		, book_authors.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_authors
	join xo_tests.users on users.user_id = book_authors.author_id
	group by
		book_authors_book_id
		, users.user_id
		, pseudonym
) as xo_join_book_authors_authors on xo_join_book_authors_authors.book_authors_book_id = books.book_id
`

const xoTestsBookTableAuthorsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_authors.__users
		, xo_join_book_authors_authors.pseudonym
		)) filter (where xo_join_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const xoTestsBookTableAuthorsGroupBySQL = `books.book_id, books.book_id`

const xoTestsBookTableAuthorsBASKJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
left join (
	select
		book_authors_surrogate_key.book_id as book_authors_surrogate_key_book_id
		, book_authors_surrogate_key.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_authors_surrogate_key
	join xo_tests.users on users.user_id = book_authors_surrogate_key.author_id
	group by
		book_authors_surrogate_key_book_id
		, users.user_id
		, pseudonym
) as xo_join_book_authors_surrogate_key_authors on xo_join_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id
`

const xoTestsBookTableAuthorsBASKSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_surrogate_key_authors.__users
		, xo_join_book_authors_surrogate_key_authors.pseudonym
		)) filter (where xo_join_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') as book_authors_surrogate_key_authors`

const xoTestsBookTableAuthorsBASKGroupBySQL = `books.book_id, books.book_id`

const xoTestsBookTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        book_id
) as xo_join_book_reviews on xo_join_book_reviews.book_reviews_book_id = books.book_id
`

const xoTestsBookTableBookReviewsSelectSQL = `COALESCE(xo_join_book_reviews.book_reviews, '{}') as book_reviews`

const xoTestsBookTableBookReviewsGroupBySQL = `xo_join_book_reviews.book_reviews, books.book_id`

const xoTestsBookTableSellersJoinSQL = `-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
		book_sellers.book_id as book_sellers_book_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_sellers
	join xo_tests.users on users.user_id = book_sellers.seller
	group by
		book_sellers_book_id
		, users.user_id
) as xo_join_book_sellers_sellers on xo_join_book_sellers_sellers.book_sellers_book_id = books.book_id
`

const xoTestsBookTableSellersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_sellers_sellers.__users
		)) filter (where xo_join_book_sellers_sellers.__users_user_id is not null), '{}') as book_sellers_sellers`

const xoTestsBookTableSellersGroupBySQL = `books.book_id, books.book_id`

// XoTestsBookUpdateParams represents update params for 'xo_tests.books'.
type XoTestsBookUpdateParams struct {
	Name *string `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.books struct fields with the specified params.
func (xtb *XoTestsBook) SetUpdateParams(params *XoTestsBookUpdateParams) {
	if params.Name != nil {
		xtb.Name = *params.Name
	}
}

// Insert inserts the XoTestsBook to the database.
func (xtb *XoTestsBook) Insert(ctx context.Context, db DB) (*XoTestsBook, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.books (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, xtb.Name)

	rows, err := db.Query(ctx, sqlstr, xtb.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Insert/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newxtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	*xtb = newxtb

	return xtb, nil
}

// Update updates a XoTestsBook in the database.
func (xtb *XoTestsBook) Update(ctx context.Context, db DB) (*XoTestsBook, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.books SET 
	name = $1 
	WHERE book_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, xtb.Name, xtb.BookID)

	rows, err := db.Query(ctx, sqlstr, xtb.Name, xtb.BookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Update/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newxtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}
	*xtb = newxtb

	return xtb, nil
}

// Upsert upserts a XoTestsBook in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtb *XoTestsBook) Upsert(ctx context.Context, db DB, params *XoTestsBookCreateParams) (*XoTestsBook, error) {
	var err error

	xtb.Name = params.Name

	xtb, err = xtb.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book", Err: err})
			}
			xtb, err = xtb.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book", Err: err})
			}
		}
	}

	return xtb, err
}

// Delete deletes the XoTestsBook from the database.
func (xtb *XoTestsBook) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.books 
	WHERE book_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtb.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookPaginatedByBookID returns a cursor-paginated list of XoTestsBook.
func XoTestsBookPaginatedByBookID(ctx context.Context, db DB, bookID XoTestsBookID, direction models.Direction, opts ...XoTestsBookSelectConfigOption) ([]XoTestsBook, error) {
	c := &XoTestsBookSelectConfig{joins: XoTestsBookJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsGroupBySQL)
	}

	if c.joins.AuthorsBASK {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsBASKSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsBASKJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsBASKGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableSellersGroupBySQL)
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

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	books.book_id,
	books.name %s 
	 FROM xo_tests.books %s 
	 WHERE books.book_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		book_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookPaginatedByBookID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book", Err: err}))
	}
	return res, nil
}

// XoTestsBookByBookID retrieves a row from 'xo_tests.books' as a XoTestsBook.
//
// Generated from index 'books_pkey'.
func XoTestsBookByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookSelectConfigOption) (*XoTestsBook, error) {
	c := &XoTestsBookSelectConfig{joins: XoTestsBookJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsGroupBySQL)
	}

	if c.joins.AuthorsBASK {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsBASKSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsBASKJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsBASKGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableSellersGroupBySQL)
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
	books.book_id,
	books.name %s 
	 FROM xo_tests.books %s 
	 WHERE books.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	xtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	return &xtb, nil
}
