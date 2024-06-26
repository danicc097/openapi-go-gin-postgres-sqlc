// Code generated by xo. DO NOT EDIT.

//lint:ignore

package got

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// XoTestsBookSeller represents a row from 'xo_tests.book_sellers'.
type XoTestsBookSeller struct {
	BookID XoTestsBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"` // book_id
	Seller XoTestsUserID `json:"seller" db:"seller" required:"true" nullable:"false"`  // seller

	SellersJoin *[]XoTestsUser `json:"-" db:"book_sellers_sellers"` // M2M book_sellers
	BooksJoin   *[]XoTestsBook `json:"-" db:"book_sellers_books"`   // M2M book_sellers
}

// XoTestsBookSellerCreateParams represents insert params for 'xo_tests.book_sellers'.
type XoTestsBookSellerCreateParams struct {
	BookID XoTestsBookID `json:"bookID" required:"true" nullable:"false"` // book_id
	Seller XoTestsUserID `json:"seller" required:"true" nullable:"false"` // seller
}

// XoTestsBookSellerParams represents common params for both insert and update of 'xo_tests.book_sellers'.
type XoTestsBookSellerParams interface {
	GetBookID() *XoTestsBookID
	GetSeller() *XoTestsUserID
}

func (p XoTestsBookSellerCreateParams) GetBookID() *XoTestsBookID {
	x := p.BookID
	return &x
}

func (p XoTestsBookSellerUpdateParams) GetBookID() *XoTestsBookID {
	return p.BookID
}

func (p XoTestsBookSellerCreateParams) GetSeller() *XoTestsUserID {
	x := p.Seller
	return &x
}

func (p XoTestsBookSellerUpdateParams) GetSeller() *XoTestsUserID {
	return p.Seller
}

// CreateXoTestsBookSeller creates a new XoTestsBookSeller in the database with the given params.
func CreateXoTestsBookSeller(ctx context.Context, db DB, params *XoTestsBookSellerCreateParams) (*XoTestsBookSeller, error) {
	xtbs := &XoTestsBookSeller{
		BookID: params.BookID,
		Seller: params.Seller,
	}

	return xtbs.Insert(ctx, db)
}

type XoTestsBookSellerSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   XoTestsBookSellerJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookSellerSelectConfigOption func(*XoTestsBookSellerSelectConfig)

// WithXoTestsBookSellerLimit limits row selection.
func WithXoTestsBookSellerLimit(limit int) XoTestsBookSellerSelectConfigOption {
	return func(s *XoTestsBookSellerSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsBookSellerOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsBookSellerOrderBy(rows map[string]*Direction) XoTestsBookSellerSelectConfigOption {
	return func(s *XoTestsBookSellerSelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsBookSeller]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type XoTestsBookSellerJoins struct {
	Sellers bool `json:"sellers" required:"true" nullable:"false"` // M2M book_sellers
	Books   bool `json:"books" required:"true" nullable:"false"`   // M2M book_sellers
}

// WithXoTestsBookSellerJoin joins with the given tables.
func WithXoTestsBookSellerJoin(joins XoTestsBookSellerJoins) XoTestsBookSellerSelectConfigOption {
	return func(s *XoTestsBookSellerSelectConfig) {
		s.joins = XoTestsBookSellerJoins{
			Sellers: s.joins.Sellers || joins.Sellers,
			Books:   s.joins.Books || joins.Books,
		}
	}
}

// WithXoTestsBookSellerFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookSellerFilters(filters map[string][]any) XoTestsBookSellerSelectConfigOption {
	return func(s *XoTestsBookSellerSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookSellerHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsBookSellerHavingClause(conditions map[string][]any) XoTestsBookSellerSelectConfigOption {
	return func(s *XoTestsBookSellerSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookSellerTableSellersJoinSQL = `-- M2M join generated from "book_sellers_seller_fkey"
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
) as xo_join_book_sellers_sellers on xo_join_book_sellers_sellers.book_sellers_book_id = book_sellers.book_id
`

const xoTestsBookSellerTableSellersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_sellers_sellers.__users
		)) filter (where xo_join_book_sellers_sellers.__users_user_id is not null), '{}') as book_sellers_sellers`

const xoTestsBookSellerTableSellersGroupBySQL = `book_sellers.book_id, book_sellers.book_id, book_sellers.seller`

const xoTestsBookSellerTableBooksJoinSQL = `-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
		book_sellers.seller as book_sellers_seller
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_sellers
	join xo_tests.books on books.book_id = book_sellers.book_id
	group by
		book_sellers_seller
		, books.book_id
) as xo_join_book_sellers_books on xo_join_book_sellers_books.book_sellers_seller = book_sellers.seller
`

const xoTestsBookSellerTableBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_sellers_books.__books
		)) filter (where xo_join_book_sellers_books.__books_book_id is not null), '{}') as book_sellers_books`

const xoTestsBookSellerTableBooksGroupBySQL = `book_sellers.seller, book_sellers.book_id, book_sellers.seller`

// XoTestsBookSellerUpdateParams represents update params for 'xo_tests.book_sellers'.
type XoTestsBookSellerUpdateParams struct {
	BookID *XoTestsBookID `json:"bookID" nullable:"false"` // book_id
	Seller *XoTestsUserID `json:"seller" nullable:"false"` // seller
}

// SetUpdateParams updates xo_tests.book_sellers struct fields with the specified params.
func (xtbs *XoTestsBookSeller) SetUpdateParams(params *XoTestsBookSellerUpdateParams) {
	if params.BookID != nil {
		xtbs.BookID = *params.BookID
	}
	if params.Seller != nil {
		xtbs.Seller = *params.Seller
	}
}

// Insert inserts the XoTestsBookSeller to the database.
func (xtbs *XoTestsBookSeller) Insert(ctx context.Context, db DB) (*XoTestsBookSeller, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_sellers (
	book_id, seller
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtbs.BookID, xtbs.Seller)
	rows, err := db.Query(ctx, sqlstr, xtbs.BookID, xtbs.Seller)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Insert/db.Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	newxtbs, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	*xtbs = newxtbs

	return xtbs, nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key or generated fields

// Delete deletes the XoTestsBookSeller from the database.
func (xtbs *XoTestsBookSeller) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_sellers 
	WHERE book_id = $1 AND seller = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtbs.BookID, xtbs.Seller); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookSellerPaginated returns a cursor-paginated list of XoTestsBookSeller.
// At least one cursor is required.
func XoTestsBookSellerPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...XoTestsBookSellerSelectConfigOption) ([]XoTestsBookSeller, error) {
	c := &XoTestsBookSellerSelectConfig{
		joins:   XoTestsBookSellerJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsBookSeller][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Paginated/cursor: %w", &XoError{Entity: "Book seller", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("book_sellers.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

	paramStart := 0 // all filters will come from the user
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
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
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

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Paginated/orderBy: %w", &XoError{Entity: "Book seller", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableSellersGroupBySQL)
	}

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookSellerTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableBooksGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM xo_tests.book_sellers %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookSellerPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Paginated/db.Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// XoTestsBookSellersByBookIDSeller retrieves a row from 'xo_tests.book_sellers' as a XoTestsBookSeller.
//
// Generated from index 'book_sellers_book_id_seller_idx'.
func XoTestsBookSellersByBookIDSeller(ctx context.Context, db DB, bookID XoTestsBookID, seller XoTestsUserID, opts ...XoTestsBookSellerSelectConfigOption) ([]XoTestsBookSeller, error) {
	c := &XoTestsBookSellerSelectConfig{joins: XoTestsBookSellerJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableSellersGroupBySQL)
	}

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookSellerTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableBooksGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM xo_tests.book_sellers %s 
	 WHERE book_sellers.book_id = $1 AND book_sellers.seller = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookSellersByBookIDSeller */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, seller)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, seller}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellersByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellersByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// XoTestsBookSellersByBookID retrieves a row from 'xo_tests.book_sellers' as a XoTestsBookSeller.
//
// Generated from index 'book_sellers_pkey'.
func XoTestsBookSellersByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookSellerSelectConfigOption) ([]XoTestsBookSeller, error) {
	c := &XoTestsBookSellerSelectConfig{joins: XoTestsBookSellerJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableSellersGroupBySQL)
	}

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookSellerTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableBooksGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM xo_tests.book_sellers %s 
	 WHERE book_sellers.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookSellersByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellerByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// XoTestsBookSellersBySeller retrieves a row from 'xo_tests.book_sellers' as a XoTestsBookSeller.
//
// Generated from index 'book_sellers_pkey'.
func XoTestsBookSellersBySeller(ctx context.Context, db DB, seller XoTestsUserID, opts ...XoTestsBookSellerSelectConfigOption) ([]XoTestsBookSeller, error) {
	c := &XoTestsBookSellerSelectConfig{joins: XoTestsBookSellerJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableSellersGroupBySQL)
	}

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookSellerTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableBooksGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM xo_tests.book_sellers %s 
	 WHERE book_sellers.seller = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookSellersBySeller */\n" + sqlstr

	// run
	// logf(sqlstr, seller)
	rows, err := db.Query(ctx, sqlstr, append([]any{seller}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellerByBookIDSeller/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellerByBookIDSeller/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// XoTestsBookSellersBySellerBookID retrieves a row from 'xo_tests.book_sellers' as a XoTestsBookSeller.
//
// Generated from index 'book_sellers_seller_book_id_idx'.
func XoTestsBookSellersBySellerBookID(ctx context.Context, db DB, seller XoTestsUserID, bookID XoTestsBookID, opts ...XoTestsBookSellerSelectConfigOption) ([]XoTestsBookSeller, error) {
	c := &XoTestsBookSellerSelectConfig{joins: XoTestsBookSellerJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookSellerTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableSellersGroupBySQL)
	}

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookSellerTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookSellerTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookSellerTableBooksGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	book_sellers.book_id,
	book_sellers.seller %s 
	 FROM xo_tests.book_sellers %s 
	 WHERE book_sellers.seller = $1 AND book_sellers.book_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookSellersBySellerBookID */\n" + sqlstr

	// run
	// logf(sqlstr, seller, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{seller, bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellersBySellerBookID/Query: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookSeller])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookSeller/BookSellersBySellerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book seller", Err: err}))
	}
	return res, nil
}

// FKBook_BookID returns the Book associated with the XoTestsBookSeller's (BookID).
//
// Generated from foreign key 'book_sellers_book_id_fkey'.
func (xtbs *XoTestsBookSeller) FKBook_BookID(ctx context.Context, db DB) (*XoTestsBook, error) {
	return XoTestsBookByBookID(ctx, db, xtbs.BookID)
}

// FKUser_Seller returns the User associated with the XoTestsBookSeller's (Seller).
//
// Generated from foreign key 'book_sellers_seller_fkey'.
func (xtbs *XoTestsBookSeller) FKUser_Seller(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtbs.Seller)
}
