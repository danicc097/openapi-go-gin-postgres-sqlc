package got

// Code generated by xo. DO NOT EDIT.

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

// BookAuthor represents a row from 'xo_tests.book_authors'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type BookAuthor struct {
	BookID    BookID  `json:"bookID" db:"book_id" required:"true" nullable:"false"`     // book_id
	AuthorID  UserID  `json:"authorID" db:"author_id" required:"true" nullable:"false"` // author_id
	Pseudonym *string `json:"pseudonym" db:"pseudonym"`                                 // pseudonym

	AuthorBooksJoin *[]Book__BA_BookAuthor `json:"-" db:"book_authors_books" openapi-go:"ignore"`   // M2M book_authors
	BookAuthorsJoin *[]User__BA_BookAuthor `json:"-" db:"book_authors_authors" openapi-go:"ignore"` // M2M book_authors
}

// BookAuthorCreateParams represents insert params for 'xo_tests.book_authors'.
type BookAuthorCreateParams struct {
	AuthorID  UserID  `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    BookID  `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string `json:"pseudonym"`                                 // pseudonym
}

// CreateBookAuthor creates a new BookAuthor in the database with the given params.
func CreateBookAuthor(ctx context.Context, db DB, params *BookAuthorCreateParams) (*BookAuthor, error) {
	ba := &BookAuthor{
		AuthorID:  params.AuthorID,
		BookID:    params.BookID,
		Pseudonym: params.Pseudonym,
	}

	return ba.Insert(ctx, db)
}

// BookAuthorUpdateParams represents update params for 'xo_tests.book_authors'.
type BookAuthorUpdateParams struct {
	AuthorID  *UserID  `json:"authorID" nullable:"false"` // author_id
	BookID    *BookID  `json:"bookID" nullable:"false"`   // book_id
	Pseudonym **string `json:"pseudonym"`                 // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors struct fields with the specified params.
func (ba *BookAuthor) SetUpdateParams(params *BookAuthorUpdateParams) {
	if params.AuthorID != nil {
		ba.AuthorID = *params.AuthorID
	}
	if params.BookID != nil {
		ba.BookID = *params.BookID
	}
	if params.Pseudonym != nil {
		ba.Pseudonym = *params.Pseudonym
	}
}

type BookAuthorSelectConfig struct {
	limit   string
	orderBy string
	joins   BookAuthorJoins
	filters map[string][]any
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

type BookAuthorOrderBy string

type BookAuthorJoins struct {
	BooksAuthor bool // M2M book_authors
	AuthorsBook bool // M2M book_authors
}

// WithBookAuthorJoin joins with the given tables.
func WithBookAuthorJoin(joins BookAuthorJoins) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.joins = BookAuthorJoins{
			BooksAuthor: s.joins.BooksAuthor || joins.BooksAuthor,
			AuthorsBook: s.joins.AuthorsBook || joins.AuthorsBook,
		}
	}
}

// Book__BA_BookAuthor represents a M2M join against "xo_tests.book_authors"
type Book__BA_BookAuthor struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BA_BookAuthor represents a M2M join against "xo_tests.book_authors"
type User__BA_BookAuthor struct {
	User      User    `json:"user" db:"users" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithBookAuthorFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithBookAuthorFilters(filters map[string][]any) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.filters = filters
	}
}

const bookAuthorTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
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
) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = book_authors.author_id
`

const bookAuthorTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const bookAuthorTableBooksAuthorGroupBySQL = `book_authors.author_id, book_authors.book_id, book_authors.author_id`

const bookAuthorTableAuthorsBookJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
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
) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = book_authors.book_id
`

const bookAuthorTableAuthorsBookSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const bookAuthorTableAuthorsBookGroupBySQL = `book_authors.book_id, book_authors.book_id, book_authors.author_id`

// Insert inserts the BookAuthor to the database.
func (ba *BookAuthor) Insert(ctx context.Context, db DB) (*BookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_authors (
	author_id, book_id, pseudonym
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, ba.AuthorID, ba.BookID, ba.Pseudonym)
	rows, err := db.Query(ctx, sqlstr, ba.AuthorID, ba.BookID, ba.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Insert/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*ba = newba

	return ba, nil
}

// Update updates a BookAuthor in the database.
func (ba *BookAuthor) Update(ctx context.Context, db DB) (*BookAuthor, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors SET 
	pseudonym = $1 
	WHERE book_id = $2  AND author_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, ba.Pseudonym, ba.BookID, ba.AuthorID)

	rows, err := db.Query(ctx, sqlstr, ba.Pseudonym, ba.BookID, ba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Update/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*ba = newba

	return ba, nil
}

// Upsert upserts a BookAuthor in the database.
// Requires appropriate PK(s) to be set beforehand.
func (ba *BookAuthor) Upsert(ctx context.Context, db DB, params *BookAuthorCreateParams) (*BookAuthor, error) {
	var err error

	ba.AuthorID = params.AuthorID
	ba.BookID = params.BookID
	ba.Pseudonym = params.Pseudonym

	ba, err = ba.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book author", Err: err})
			}
			ba, err = ba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book author", Err: err})
			}
		}
	}

	return ba, err
}

// Delete deletes the BookAuthor from the database.
func (ba *BookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_authors 
	WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, ba.BookID, ba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookAuthorByBookIDAuthorID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID BookAuthorID, authorID BookAuthorID, opts ...BookAuthorSelectConfigOption) (*BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}, filters: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, bookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, bookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableAuthorsBookGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1 AND book_authors.author_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookAuthorByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	ba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}

	return &ba, nil
}

// BookAuthorsByBookID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorsByBookID(ctx context.Context, db DB, bookID BookAuthorID, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}, filters: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, bookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, bookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableAuthorsBookGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookAuthorsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// BookAuthorsByAuthorID retrieves a row from 'xo_tests.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_pkey'.
func BookAuthorsByAuthorID(ctx context.Context, db DB, authorID BookAuthorID, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}, filters: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, bookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, bookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, bookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorTableAuthorsBookGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.author_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* BookAuthorsByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
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
