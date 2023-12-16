package db

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

// ExtraSchemaBookAuthor represents a row from 'extra_schema.book_authors'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaBookAuthor struct {
	BookID    ExtraSchemaBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"`     // book_id
	AuthorID  ExtraSchemaUserID `json:"authorID" db:"author_id" required:"true" nullable:"false"` // author_id
	Pseudonym *string           `json:"pseudonym" db:"pseudonym"`                                 // pseudonym

	AuthorBooksJoin *[]Book__BA_ExtraSchemaBookAuthor `json:"-" db:"book_authors_books" openapi-go:"ignore"`   // M2M book_authors
	BookAuthorsJoin *[]User__BA_ExtraSchemaBookAuthor `json:"-" db:"book_authors_authors" openapi-go:"ignore"` // M2M book_authors

}

// ExtraSchemaBookAuthorCreateParams represents insert params for 'extra_schema.book_authors'.
type ExtraSchemaBookAuthorCreateParams struct {
	AuthorID  ExtraSchemaUserID `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    ExtraSchemaBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string           `json:"pseudonym"`                                 // pseudonym
}

// CreateExtraSchemaBookAuthor creates a new ExtraSchemaBookAuthor in the database with the given params.
func CreateExtraSchemaBookAuthor(ctx context.Context, db DB, params *ExtraSchemaBookAuthorCreateParams) (*ExtraSchemaBookAuthor, error) {
	esba := &ExtraSchemaBookAuthor{
		AuthorID:  params.AuthorID,
		BookID:    params.BookID,
		Pseudonym: params.Pseudonym,
	}

	return esba.Insert(ctx, db)
}

// ExtraSchemaBookAuthorUpdateParams represents update params for 'extra_schema.book_authors'.
type ExtraSchemaBookAuthorUpdateParams struct {
	AuthorID  *ExtraSchemaUserID `json:"authorID" nullable:"false"` // author_id
	BookID    *ExtraSchemaBookID `json:"bookID" nullable:"false"`   // book_id
	Pseudonym **string           `json:"pseudonym"`                 // pseudonym
}

// SetUpdateParams updates extra_schema.book_authors struct fields with the specified params.
func (esba *ExtraSchemaBookAuthor) SetUpdateParams(params *ExtraSchemaBookAuthorUpdateParams) {
	if params.AuthorID != nil {
		esba.AuthorID = *params.AuthorID
	}
	if params.BookID != nil {
		esba.BookID = *params.BookID
	}
	if params.Pseudonym != nil {
		esba.Pseudonym = *params.Pseudonym
	}
}

type ExtraSchemaBookAuthorSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaBookAuthorJoins
	filters map[string][]any
}
type ExtraSchemaBookAuthorSelectConfigOption func(*ExtraSchemaBookAuthorSelectConfig)

// WithExtraSchemaBookAuthorLimit limits row selection.
func WithExtraSchemaBookAuthorLimit(limit int) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaBookAuthorOrderBy string

const ()

type ExtraSchemaBookAuthorJoins struct {
	BooksAuthor bool // M2M book_authors
	AuthorsBook bool // M2M book_authors
}

// WithExtraSchemaBookAuthorJoin joins with the given tables.
func WithExtraSchemaBookAuthorJoin(joins ExtraSchemaBookAuthorJoins) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		s.joins = ExtraSchemaBookAuthorJoins{
			BooksAuthor: s.joins.BooksAuthor || joins.BooksAuthor,
			AuthorsBook: s.joins.AuthorsBook || joins.AuthorsBook,
		}
	}
}

// Book__BA_ExtraSchemaBookAuthor represents a M2M join against "extra_schema.book_authors"
type Book__BA_ExtraSchemaBookAuthor struct {
	Book      ExtraSchemaBook `json:"book" db:"books" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BA_ExtraSchemaBookAuthor represents a M2M join against "extra_schema.book_authors"
type User__BA_ExtraSchemaBookAuthor struct {
	User      ExtraSchemaUser `json:"user" db:"users" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithExtraSchemaBookAuthorFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaBookAuthorFilters(filters map[string][]any) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		s.filters = filters
	}
}

const extraSchemaBookAuthorTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
		book_authors.author_id as book_authors_author_id
		, book_authors.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		extra_schema.book_authors
	join extra_schema.books on books.book_id = book_authors.book_id
	group by
		book_authors_author_id
		, books.book_id
		, pseudonym
) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = book_authors.author_id
`

const extraSchemaBookAuthorTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const extraSchemaBookAuthorTableBooksAuthorGroupBySQL = `book_authors.author_id, book_authors.book_id, book_authors.author_id`

const extraSchemaBookAuthorTableAuthorsBookJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
		book_authors.book_id as book_authors_book_id
		, book_authors.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.book_authors
	join extra_schema.users on users.user_id = book_authors.author_id
	group by
		book_authors_book_id
		, users.user_id
		, pseudonym
) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = book_authors.book_id
`

const extraSchemaBookAuthorTableAuthorsBookSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const extraSchemaBookAuthorTableAuthorsBookGroupBySQL = `book_authors.book_id, book_authors.book_id, book_authors.author_id`

// Insert inserts the ExtraSchemaBookAuthor to the database.
func (esba *ExtraSchemaBookAuthor) Insert(ctx context.Context, db DB) (*ExtraSchemaBookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.book_authors (
	author_id, book_id, pseudonym
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, esba.AuthorID, esba.BookID, esba.Pseudonym)
	rows, err := db.Query(ctx, sqlstr, esba.AuthorID, esba.BookID, esba.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Insert/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newesba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*esba = newesba

	return esba, nil
}

// Update updates a ExtraSchemaBookAuthor in the database.
func (esba *ExtraSchemaBookAuthor) Update(ctx context.Context, db DB) (*ExtraSchemaBookAuthor, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.book_authors SET 
	pseudonym = $1 
	WHERE book_id = $2  AND author_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, esba.Pseudonym, esba.BookID, esba.AuthorID)

	rows, err := db.Query(ctx, sqlstr, esba.Pseudonym, esba.BookID, esba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Update/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newesba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*esba = newesba

	return esba, nil
}

// Upsert upserts a ExtraSchemaBookAuthor in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esba *ExtraSchemaBookAuthor) Upsert(ctx context.Context, db DB, params *ExtraSchemaBookAuthorCreateParams) (*ExtraSchemaBookAuthor, error) {
	var err error

	esba.AuthorID = params.AuthorID
	esba.BookID = params.BookID
	esba.Pseudonym = params.Pseudonym

	esba, err = esba.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book author", Err: err})
			}
			esba, err = esba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book author", Err: err})
			}
		}
	}

	return esba, err
}

// Delete deletes the ExtraSchemaBookAuthor from the database.
func (esba *ExtraSchemaBookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM extra_schema.book_authors 
	WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esba.BookID, esba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaBookAuthorByBookIDAuthorID retrieves a row from 'extra_schema.book_authors' as a ExtraSchemaBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func ExtraSchemaBookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID ExtraSchemaBookID, authorID ExtraSchemaUserID, opts ...ExtraSchemaBookAuthorSelectConfigOption) (*ExtraSchemaBookAuthor, error) {
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsBookGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.book_id = $1 AND book_authors.author_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	esba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}

	return &esba, nil
}

// ExtraSchemaBookAuthorsByBookID retrieves a row from 'extra_schema.book_authors' as a ExtraSchemaBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func ExtraSchemaBookAuthorsByBookID(ctx context.Context, db DB, bookID ExtraSchemaBookID, opts ...ExtraSchemaBookAuthorSelectConfigOption) ([]ExtraSchemaBookAuthor, error) {
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsBookGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.book_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// ExtraSchemaBookAuthorsByAuthorID retrieves a row from 'extra_schema.book_authors' as a ExtraSchemaBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func ExtraSchemaBookAuthorsByAuthorID(ctx context.Context, db DB, authorID ExtraSchemaUserID, opts ...ExtraSchemaBookAuthorSelectConfigOption) ([]ExtraSchemaBookAuthor, error) {
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsBookGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.author_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorsByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// FKUser_AuthorID returns the User associated with the ExtraSchemaBookAuthor's (AuthorID).
//
// Generated from foreign key 'book_authors_author_id_fkey'.
func (esba *ExtraSchemaBookAuthor) FKUser_AuthorID(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, esba.AuthorID)
}

// FKBook_BookID returns the Book associated with the ExtraSchemaBookAuthor's (BookID).
//
// Generated from foreign key 'book_authors_book_id_fkey'.
func (esba *ExtraSchemaBookAuthor) FKBook_BookID(ctx context.Context, db DB) (*ExtraSchemaBook, error) {
	return ExtraSchemaBookByBookID(ctx, db, esba.BookID)
}
