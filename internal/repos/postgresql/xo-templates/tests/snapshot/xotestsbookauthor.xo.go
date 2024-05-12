// Code generated by xo. DO NOT EDIT.

//lint:ignore

package got

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsBookAuthor represents a row from 'xo_tests.book_authors'.
type XoTestsBookAuthor struct {
	BookID    XoTestsBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"`     // book_id
	AuthorID  XoTestsUserID `json:"authorID" db:"author_id" required:"true" nullable:"false"` // author_id
	Pseudonym *string       `json:"pseudonym" db:"pseudonym"`                                 // pseudonym

	BooksJoin   *[]XoTestsBookAuthorM2MBookBA   `json:"-" db:"book_authors_books"`   // M2M book_authors
	AuthorsJoin *[]XoTestsBookAuthorM2MAuthorBA `json:"-" db:"book_authors_authors"` // M2M book_authors
}

// XoTestsBookAuthorCreateParams represents insert params for 'xo_tests.book_authors'.
type XoTestsBookAuthorCreateParams struct {
	AuthorID  XoTestsUserID `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    XoTestsBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string       `json:"pseudonym"`                                 // pseudonym
}

// XoTestsBookAuthorParams represents common params for both insert and update of 'xo_tests.book_authors'.
type XoTestsBookAuthorParams interface {
	GetAuthorID() *XoTestsUserID
	GetBookID() *XoTestsBookID
	GetPseudonym() *string
}

func (p XoTestsBookAuthorCreateParams) GetAuthorID() *XoTestsUserID {
	x := p.AuthorID
	return &x
}

func (p XoTestsBookAuthorUpdateParams) GetAuthorID() *XoTestsUserID {
	return p.AuthorID
}

func (p XoTestsBookAuthorCreateParams) GetBookID() *XoTestsBookID {
	x := p.BookID
	return &x
}

func (p XoTestsBookAuthorUpdateParams) GetBookID() *XoTestsBookID {
	return p.BookID
}

func (p XoTestsBookAuthorCreateParams) GetPseudonym() *string {
	return p.Pseudonym
}

func (p XoTestsBookAuthorUpdateParams) GetPseudonym() *string {
	if p.Pseudonym != nil {
		return *p.Pseudonym
	}
	return nil
}

// CreateXoTestsBookAuthor creates a new XoTestsBookAuthor in the database with the given params.
func CreateXoTestsBookAuthor(ctx context.Context, db DB, params *XoTestsBookAuthorCreateParams) (*XoTestsBookAuthor, error) {
	xtba := &XoTestsBookAuthor{
		AuthorID:  params.AuthorID,
		BookID:    params.BookID,
		Pseudonym: params.Pseudonym,
	}

	return xtba.Insert(ctx, db)
}

type XoTestsBookAuthorSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   XoTestsBookAuthorJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookAuthorSelectConfigOption func(*XoTestsBookAuthorSelectConfig)

// WithXoTestsBookAuthorLimit limits row selection.
func WithXoTestsBookAuthorLimit(limit int) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsBookAuthorOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsBookAuthorOrderBy(rows map[string]*Direction) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsBookAuthor]
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

type XoTestsBookAuthorJoins struct {
	Books   bool `json:"books" required:"true" nullable:"false"`   // M2M book_authors
	Authors bool `json:"authors" required:"true" nullable:"false"` // M2M book_authors
}

// WithXoTestsBookAuthorJoin joins with the given tables.
func WithXoTestsBookAuthorJoin(joins XoTestsBookAuthorJoins) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.joins = XoTestsBookAuthorJoins{
			Books:   s.joins.Books || joins.Books,
			Authors: s.joins.Authors || joins.Authors,
		}
	}
}

// XoTestsBookAuthorM2MBookBA represents a M2M join against "xo_tests.book_authors"
type XoTestsBookAuthorM2MBookBA struct {
	Book      XoTestsBook `json:"book" db:"books" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// XoTestsBookAuthorM2MAuthorBA represents a M2M join against "xo_tests.book_authors"
type XoTestsBookAuthorM2MAuthorBA struct {
	User      XoTestsUser `json:"user" db:"users" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithXoTestsBookAuthorFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookAuthorFilters(filters map[string][]any) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookAuthorHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsBookAuthorHavingClause(conditions map[string][]any) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookAuthorTableBooksJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
		book_authors.author_id as book_authors_author_id
		, book_authors.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_authors
	join xo_tests.books on books.book_id = book_authors.book_id
	group by
		book_authors_author_id
		, books.book_id
		, pseudonym
) as xo_join_book_authors_books on xo_join_book_authors_books.book_authors_author_id = book_authors.author_id
`

const xoTestsBookAuthorTableBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_books.__books
		, xo_join_book_authors_books.pseudonym
		)) filter (where xo_join_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const xoTestsBookAuthorTableBooksGroupBySQL = `book_authors.author_id, book_authors.book_id, book_authors.author_id`

const xoTestsBookAuthorTableAuthorsJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
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
) as xo_join_book_authors_authors on xo_join_book_authors_authors.book_authors_book_id = book_authors.book_id
`

const xoTestsBookAuthorTableAuthorsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_authors.__users
		, xo_join_book_authors_authors.pseudonym
		)) filter (where xo_join_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const xoTestsBookAuthorTableAuthorsGroupBySQL = `book_authors.book_id, book_authors.book_id, book_authors.author_id`

// XoTestsBookAuthorUpdateParams represents update params for 'xo_tests.book_authors'.
type XoTestsBookAuthorUpdateParams struct {
	AuthorID  *XoTestsUserID `json:"authorID" nullable:"false"` // author_id
	BookID    *XoTestsBookID `json:"bookID" nullable:"false"`   // book_id
	Pseudonym **string       `json:"pseudonym"`                 // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors struct fields with the specified params.
func (xtba *XoTestsBookAuthor) SetUpdateParams(params *XoTestsBookAuthorUpdateParams) {
	if params.AuthorID != nil {
		xtba.AuthorID = *params.AuthorID
	}
	if params.BookID != nil {
		xtba.BookID = *params.BookID
	}
	if params.Pseudonym != nil {
		xtba.Pseudonym = *params.Pseudonym
	}
}

// Insert inserts the XoTestsBookAuthor to the database.
func (xtba *XoTestsBookAuthor) Insert(ctx context.Context, db DB) (*XoTestsBookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_authors (
	author_id, book_id, pseudonym
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtba.AuthorID, xtba.BookID, xtba.Pseudonym)
	rows, err := db.Query(ctx, sqlstr, xtba.AuthorID, xtba.BookID, xtba.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Insert/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newxtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*xtba = newxtba

	return xtba, nil
}

// Update updates a XoTestsBookAuthor in the database.
func (xtba *XoTestsBookAuthor) Update(ctx context.Context, db DB) (*XoTestsBookAuthor, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors SET 
	pseudonym = $1 
	WHERE book_id = $2  AND author_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, xtba.Pseudonym, xtba.BookID, xtba.AuthorID)

	rows, err := db.Query(ctx, sqlstr, xtba.Pseudonym, xtba.BookID, xtba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Update/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newxtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*xtba = newxtba

	return xtba, nil
}

// Upsert upserts a XoTestsBookAuthor in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtba *XoTestsBookAuthor) Upsert(ctx context.Context, db DB, params *XoTestsBookAuthorCreateParams) (*XoTestsBookAuthor, error) {
	var err error

	xtba.AuthorID = params.AuthorID
	xtba.BookID = params.BookID
	xtba.Pseudonym = params.Pseudonym

	xtba, err = xtba.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertXoTestsBookAuthor/Insert: %w", &XoError{Entity: "Book author", Err: err})
			}
			xtba, err = xtba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertXoTestsBookAuthor/Update: %w", &XoError{Entity: "Book author", Err: err})
			}
		}
	}

	return xtba, err
}

// Delete deletes the XoTestsBookAuthor from the database.
func (xtba *XoTestsBookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_authors 
	WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtba.BookID, xtba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookAuthorPaginated returns a cursor-paginated list of XoTestsBookAuthor.
// At least one cursor is required.
func XoTestsBookAuthorPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...XoTestsBookAuthorSelectConfigOption) ([]XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{
		joins:   XoTestsBookAuthorJoins{},
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
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsBookAuthor][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Paginated/cursor: %w", &XoError{Entity: "Book author", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("book_authors.%s %s $i", field.Db, op)] = []any{*cursor.Value}
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
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Paginated/orderBy: %w", &XoError{Entity: "Book author", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Paginated/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorByBookIDAuthorID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID XoTestsBookID, authorID XoTestsUserID, opts ...XoTestsBookAuthorSelectConfigOption) (*XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1 AND book_authors.author_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	xtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}

	return &xtba, nil
}

// XoTestsBookAuthorsByBookID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorsByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookAuthorSelectConfigOption) ([]XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorsByAuthorID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorsByAuthorID(ctx context.Context, db DB, authorID XoTestsUserID, opts ...XoTestsBookAuthorSelectConfigOption) ([]XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.author_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// FKUser_AuthorID returns the User associated with the XoTestsBookAuthor's (AuthorID).
//
// Generated from foreign key 'book_authors_author_id_fkey'.
func (xtba *XoTestsBookAuthor) FKUser_AuthorID(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtba.AuthorID)
}

// FKBook_BookID returns the Book associated with the XoTestsBookAuthor's (BookID).
//
// Generated from foreign key 'book_authors_book_id_fkey'.
func (xtba *XoTestsBookAuthor) FKBook_BookID(ctx context.Context, db DB) (*XoTestsBook, error) {
	return XoTestsBookByBookID(ctx, db, xtba.BookID)
}
