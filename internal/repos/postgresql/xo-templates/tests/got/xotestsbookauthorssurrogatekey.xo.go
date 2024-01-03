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

// XoTestsBookAuthorsSurrogateKey represents a row from 'xo_tests.book_authors_surrogate_key'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsBookAuthorsSurrogateKey struct {
	BookAuthorsSurrogateKeyID XoTestsBookAuthorsSurrogateKeyID `json:"bookAuthorsSurrogateKeyID" db:"book_authors_surrogate_key_id" required:"true" nullable:"false"` // book_authors_surrogate_key_id
	BookID                    XoTestsBookID                    `json:"bookID" db:"book_id" required:"true" nullable:"false"`                                          // book_id
	AuthorID                  XoTestsUserID                    `json:"authorID" db:"author_id" required:"true" nullable:"false"`                                      // author_id
	Pseudonym                 *string                          `json:"pseudonym" db:"pseudonym"`                                                                      // pseudonym

	AuthorBooksJoin *[]Book__BASK_XoTestsBookAuthorsSurrogateKey `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"`   // M2M book_authors_surrogate_key
	BookAuthorsJoin *[]User__BASK_XoTestsBookAuthorsSurrogateKey `json:"-" db:"book_authors_surrogate_key_authors" openapi-go:"ignore"` // M2M book_authors_surrogate_key
}

// XoTestsBookAuthorsSurrogateKeyCreateParams represents insert params for 'xo_tests.book_authors_surrogate_key'.
type XoTestsBookAuthorsSurrogateKeyCreateParams struct {
	AuthorID  XoTestsUserID `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    XoTestsBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string       `json:"pseudonym"`                                 // pseudonym
}

type XoTestsBookAuthorsSurrogateKeyID int

// CreateXoTestsBookAuthorsSurrogateKey creates a new XoTestsBookAuthorsSurrogateKey in the database with the given params.
func CreateXoTestsBookAuthorsSurrogateKey(ctx context.Context, db DB, params *XoTestsBookAuthorsSurrogateKeyCreateParams) (*XoTestsBookAuthorsSurrogateKey, error) {
	xtbask := &XoTestsBookAuthorsSurrogateKey{
		AuthorID:  params.AuthorID,
		BookID:    params.BookID,
		Pseudonym: params.Pseudonym,
	}

	return xtbask.Insert(ctx, db)
}

// XoTestsBookAuthorsSurrogateKeyUpdateParams represents update params for 'xo_tests.book_authors_surrogate_key'.
type XoTestsBookAuthorsSurrogateKeyUpdateParams struct {
	AuthorID  *XoTestsUserID `json:"authorID" nullable:"false"` // author_id
	BookID    *XoTestsBookID `json:"bookID" nullable:"false"`   // book_id
	Pseudonym **string       `json:"pseudonym"`                 // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors_surrogate_key struct fields with the specified params.
func (xtbask *XoTestsBookAuthorsSurrogateKey) SetUpdateParams(params *XoTestsBookAuthorsSurrogateKeyUpdateParams) {
	if params.AuthorID != nil {
		xtbask.AuthorID = *params.AuthorID
	}
	if params.BookID != nil {
		xtbask.BookID = *params.BookID
	}
	if params.Pseudonym != nil {
		xtbask.Pseudonym = *params.Pseudonym
	}
}

type XoTestsBookAuthorsSurrogateKeySelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsBookAuthorsSurrogateKeyJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookAuthorsSurrogateKeySelectConfigOption func(*XoTestsBookAuthorsSurrogateKeySelectConfig)

// WithXoTestsBookAuthorsSurrogateKeyLimit limits row selection.
func WithXoTestsBookAuthorsSurrogateKeyLimit(limit int) XoTestsBookAuthorsSurrogateKeySelectConfigOption {
	return func(s *XoTestsBookAuthorsSurrogateKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsBookAuthorsSurrogateKeyOrderBy string

type XoTestsBookAuthorsSurrogateKeyJoins struct {
	BooksAuthor bool // M2M book_authors_surrogate_key
	AuthorsBook bool // M2M book_authors_surrogate_key
}

// WithXoTestsBookAuthorsSurrogateKeyJoin joins with the given tables.
func WithXoTestsBookAuthorsSurrogateKeyJoin(joins XoTestsBookAuthorsSurrogateKeyJoins) XoTestsBookAuthorsSurrogateKeySelectConfigOption {
	return func(s *XoTestsBookAuthorsSurrogateKeySelectConfig) {
		s.joins = XoTestsBookAuthorsSurrogateKeyJoins{
			BooksAuthor: s.joins.BooksAuthor || joins.BooksAuthor,
			AuthorsBook: s.joins.AuthorsBook || joins.AuthorsBook,
		}
	}
}

// Book__BASK_XoTestsBookAuthorsSurrogateKey represents a M2M join against "xo_tests.book_authors_surrogate_key"
type Book__BASK_XoTestsBookAuthorsSurrogateKey struct {
	Book      XoTestsBook `json:"book" db:"books" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BASK_XoTestsBookAuthorsSurrogateKey represents a M2M join against "xo_tests.book_authors_surrogate_key"
type User__BASK_XoTestsBookAuthorsSurrogateKey struct {
	User      XoTestsUser `json:"user" db:"users" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithXoTestsBookAuthorsSurrogateKeyFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookAuthorsSurrogateKeyFilters(filters map[string][]any) XoTestsBookAuthorsSurrogateKeySelectConfigOption {
	return func(s *XoTestsBookAuthorsSurrogateKeySelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookAuthorsSurrogateKeyHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// // filter a given aggregate of assigned users to return results where at least one of them has id of userId
//
//	filters := map[string][]any{
//		"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsBookAuthorsSurrogateKeyHavingClause(conditions map[string][]any) XoTestsBookAuthorsSurrogateKeySelectConfigOption {
	return func(s *XoTestsBookAuthorsSurrogateKeySelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
left join (
	select
		book_authors_surrogate_key.author_id as book_authors_surrogate_key_author_id
		, book_authors_surrogate_key.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_authors_surrogate_key
	join xo_tests.books on books.book_id = book_authors_surrogate_key.book_id
	group by
		book_authors_surrogate_key_author_id
		, books.book_id
		, pseudonym
) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = book_authors_surrogate_key.author_id
`

const xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books_book_id is not null), '{}') as book_authors_surrogate_key_books`

const xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL = `book_authors_surrogate_key.author_id, book_authors_surrogate_key.book_authors_surrogate_key_id`

const xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
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
) as joined_book_authors_surrogate_key_authors on joined_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = book_authors_surrogate_key.book_id
`

const xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') as book_authors_surrogate_key_authors`

const xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL = `book_authors_surrogate_key.book_id, book_authors_surrogate_key.book_authors_surrogate_key_id`

// Insert inserts the XoTestsBookAuthorsSurrogateKey to the database.
func (xtbask *XoTestsBookAuthorsSurrogateKey) Insert(ctx context.Context, db DB) (*XoTestsBookAuthorsSurrogateKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.book_authors_surrogate_key (
	author_id, book_id, pseudonym
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtbask.AuthorID, xtbask.BookID, xtbask.Pseudonym)

	rows, err := db.Query(ctx, sqlstr, xtbask.AuthorID, xtbask.BookID, xtbask.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Insert/db.Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	newxtbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}

	*xtbask = newxtbask

	return xtbask, nil
}

// Update updates a XoTestsBookAuthorsSurrogateKey in the database.
func (xtbask *XoTestsBookAuthorsSurrogateKey) Update(ctx context.Context, db DB) (*XoTestsBookAuthorsSurrogateKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors_surrogate_key SET 
	author_id = $1, book_id = $2, pseudonym = $3 
	WHERE book_authors_surrogate_key_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, xtbask.AuthorID, xtbask.BookID, xtbask.Pseudonym, xtbask.BookAuthorsSurrogateKeyID)

	rows, err := db.Query(ctx, sqlstr, xtbask.AuthorID, xtbask.BookID, xtbask.Pseudonym, xtbask.BookAuthorsSurrogateKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Update/db.Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	newxtbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	*xtbask = newxtbask

	return xtbask, nil
}

// Upsert upserts a XoTestsBookAuthorsSurrogateKey in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtbask *XoTestsBookAuthorsSurrogateKey) Upsert(ctx context.Context, db DB, params *XoTestsBookAuthorsSurrogateKeyCreateParams) (*XoTestsBookAuthorsSurrogateKey, error) {
	var err error

	xtbask.AuthorID = params.AuthorID
	xtbask.BookID = params.BookID
	xtbask.Pseudonym = params.Pseudonym

	xtbask, err = xtbask.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book authors surrogate key", Err: err})
			}
			xtbask, err = xtbask.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book authors surrogate key", Err: err})
			}
		}
	}

	return xtbask, err
}

// Delete deletes the XoTestsBookAuthorsSurrogateKey from the database.
func (xtbask *XoTestsBookAuthorsSurrogateKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.book_authors_surrogate_key 
	WHERE book_authors_surrogate_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtbask.BookAuthorsSurrogateKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyID returns a cursor-paginated list of XoTestsBookAuthorsSurrogateKey.
func XoTestsBookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyID(ctx context.Context, db DB, bookAuthorsSurrogateKeyID XoTestsBookAuthorsSurrogateKeyID, direction models.Direction, opts ...XoTestsBookAuthorsSurrogateKeySelectConfigOption) ([]XoTestsBookAuthorsSurrogateKey, error) {
	c := &XoTestsBookAuthorsSurrogateKeySelectConfig{joins: XoTestsBookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
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
	book_authors_surrogate_key.author_id,
	book_authors_surrogate_key.book_authors_surrogate_key_id,
	book_authors_surrogate_key.book_id,
	book_authors_surrogate_key.pseudonym %s 
	 FROM xo_tests.book_authors_surrogate_key %s 
	 WHERE book_authors_surrogate_key.book_authors_surrogate_key_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		book_authors_surrogate_key_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookAuthorsSurrogateKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Paginated/db.Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorsSurrogateKeyByBookIDAuthorID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a XoTestsBookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func XoTestsBookAuthorsSurrogateKeyByBookIDAuthorID(ctx context.Context, db DB, bookID XoTestsBookID, authorID XoTestsUserID, opts ...XoTestsBookAuthorsSurrogateKeySelectConfigOption) (*XoTestsBookAuthorsSurrogateKey, error) {
	c := &XoTestsBookAuthorsSurrogateKeySelectConfig{joins: XoTestsBookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
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
	book_authors_surrogate_key.author_id,
	book_authors_surrogate_key.book_authors_surrogate_key_id,
	book_authors_surrogate_key.book_id,
	book_authors_surrogate_key.pseudonym %s 
	 FROM xo_tests.book_authors_surrogate_key %s 
	 WHERE book_authors_surrogate_key.book_id = $1 AND book_authors_surrogate_key.author_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsSurrogateKeyByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookIDAuthorID/db.Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	xtbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectOneRow: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}

	return &xtbask, nil
}

// XoTestsBookAuthorsSurrogateKeysByBookID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a XoTestsBookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func XoTestsBookAuthorsSurrogateKeysByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookAuthorsSurrogateKeySelectConfigOption) ([]XoTestsBookAuthorsSurrogateKey, error) {
	c := &XoTestsBookAuthorsSurrogateKeySelectConfig{joins: XoTestsBookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
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
	book_authors_surrogate_key.author_id,
	book_authors_surrogate_key.book_authors_surrogate_key_id,
	book_authors_surrogate_key.book_id,
	book_authors_surrogate_key.pseudonym %s 
	 FROM xo_tests.book_authors_surrogate_key %s 
	 WHERE book_authors_surrogate_key.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsSurrogateKeysByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorsSurrogateKeysByAuthorID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a XoTestsBookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func XoTestsBookAuthorsSurrogateKeysByAuthorID(ctx context.Context, db DB, authorID XoTestsUserID, opts ...XoTestsBookAuthorsSurrogateKeySelectConfigOption) ([]XoTestsBookAuthorsSurrogateKey, error) {
	c := &XoTestsBookAuthorsSurrogateKeySelectConfig{joins: XoTestsBookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
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
	book_authors_surrogate_key.author_id,
	book_authors_surrogate_key.book_authors_surrogate_key_id,
	book_authors_surrogate_key.book_id,
	book_authors_surrogate_key.pseudonym %s 
	 FROM xo_tests.book_authors_surrogate_key %s 
	 WHERE book_authors_surrogate_key.author_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsSurrogateKeysByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a XoTestsBookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_pkey'.
func XoTestsBookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID(ctx context.Context, db DB, bookAuthorsSurrogateKeyID XoTestsBookAuthorsSurrogateKeyID, opts ...XoTestsBookAuthorsSurrogateKeySelectConfigOption) (*XoTestsBookAuthorsSurrogateKey, error) {
	c := &XoTestsBookAuthorsSurrogateKeySelectConfig{joins: XoTestsBookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
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
	book_authors_surrogate_key.author_id,
	book_authors_surrogate_key.book_authors_surrogate_key_id,
	book_authors_surrogate_key.book_id,
	book_authors_surrogate_key.pseudonym %s 
	 FROM xo_tests.book_authors_surrogate_key %s 
	 WHERE book_authors_surrogate_key.book_authors_surrogate_key_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID */\n" + sqlstr

	// run
	// logf(sqlstr, bookAuthorsSurrogateKeyID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookAuthorsSurrogateKeyID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/db.Query: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}
	xtbask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/pgx.CollectOneRow: %w", &XoError{Entity: "Book authors surrogate key", Err: err}))
	}

	return &xtbask, nil
}

// FKUser_AuthorID returns the User associated with the XoTestsBookAuthorsSurrogateKey's (AuthorID).
//
// Generated from foreign key 'book_authors_surrogate_key_author_id_fkey'.
func (xtbask *XoTestsBookAuthorsSurrogateKey) FKUser_AuthorID(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtbask.AuthorID)
}

// FKBook_BookID returns the Book associated with the XoTestsBookAuthorsSurrogateKey's (BookID).
//
// Generated from foreign key 'book_authors_surrogate_key_book_id_fkey'.
func (xtbask *XoTestsBookAuthorsSurrogateKey) FKBook_BookID(ctx context.Context, db DB) (*XoTestsBook, error) {
	return XoTestsBookByBookID(ctx, db, xtbask.BookID)
}
