package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// BookAuthorsSurrogateKey represents a row from 'xo_tests.book_authors_surrogate_key'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type BookAuthorsSurrogateKey struct {
	BookAuthorsSurrogateKeyID int       `json:"bookAuthorsSurrogateKeyID" db:"book_authors_surrogate_key_id" required:"true"` // book_authors_surrogate_key_id
	BookID                    int       `json:"bookID" db:"book_id" required:"true"`                                          // book_id
	AuthorID                  uuid.UUID `json:"authorID" db:"author_id" required:"true"`                                      // author_id
	Pseudonym                 *string   `json:"pseudonym" db:"pseudonym" required:"true"`                                     // pseudonym

	AuthorBooksJoin *[]Book__BASK_BookAuthorsSurrogateKey `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"`   // M2M book_authors_surrogate_key
	BookAuthorsJoin *[]User__BASK_BookAuthorsSurrogateKey `json:"-" db:"book_authors_surrogate_key_authors" openapi-go:"ignore"` // M2M book_authors_surrogate_key
}

// BookAuthorsSurrogateKeyCreateParams represents insert params for 'xo_tests.book_authors_surrogate_key'.
type BookAuthorsSurrogateKeyCreateParams struct {
	BookID    int       `json:"bookID" required:"true"`    // book_id
	AuthorID  uuid.UUID `json:"authorID" required:"true"`  // author_id
	Pseudonym *string   `json:"pseudonym" required:"true"` // pseudonym
}

// CreateBookAuthorsSurrogateKey creates a new BookAuthorsSurrogateKey in the database with the given params.
func CreateBookAuthorsSurrogateKey(ctx context.Context, db DB, params *BookAuthorsSurrogateKeyCreateParams) (*BookAuthorsSurrogateKey, error) {
	bask := &BookAuthorsSurrogateKey{
		BookID:    params.BookID,
		AuthorID:  params.AuthorID,
		Pseudonym: params.Pseudonym,
	}

	return bask.Insert(ctx, db)
}

// BookAuthorsSurrogateKeyUpdateParams represents update params for 'xo_tests.book_authors_surrogate_key'.
type BookAuthorsSurrogateKeyUpdateParams struct {
	BookID    *int       `json:"bookID" required:"true"`    // book_id
	AuthorID  *uuid.UUID `json:"authorID" required:"true"`  // author_id
	Pseudonym **string   `json:"pseudonym" required:"true"` // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors_surrogate_key struct fields with the specified params.
func (bask *BookAuthorsSurrogateKey) SetUpdateParams(params *BookAuthorsSurrogateKeyUpdateParams) {
	if params.BookID != nil {
		bask.BookID = *params.BookID
	}
	if params.AuthorID != nil {
		bask.AuthorID = *params.AuthorID
	}
	if params.Pseudonym != nil {
		bask.Pseudonym = *params.Pseudonym
	}
}

type BookAuthorsSurrogateKeySelectConfig struct {
	limit   string
	orderBy string
	joins   BookAuthorsSurrogateKeyJoins
	filters map[string][]any
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

type BookAuthorsSurrogateKeyOrderBy string

type BookAuthorsSurrogateKeyJoins struct {
	BooksAuthor bool // M2M book_authors_surrogate_key
	AuthorsBook bool // M2M book_authors_surrogate_key
}

// WithBookAuthorsSurrogateKeyJoin joins with the given tables.
func WithBookAuthorsSurrogateKeyJoin(joins BookAuthorsSurrogateKeyJoins) BookAuthorsSurrogateKeySelectConfigOption {
	return func(s *BookAuthorsSurrogateKeySelectConfig) {
		s.joins = BookAuthorsSurrogateKeyJoins{
			BooksAuthor: s.joins.BooksAuthor || joins.BooksAuthor,
			AuthorsBook: s.joins.AuthorsBook || joins.AuthorsBook,
		}
	}
}

// Book__BASK_BookAuthorsSurrogateKey represents a M2M join against "xo_tests.book_authors_surrogate_key"
type Book__BASK_BookAuthorsSurrogateKey struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BASK_BookAuthorsSurrogateKey represents a M2M join against "xo_tests.book_authors_surrogate_key"
type User__BASK_BookAuthorsSurrogateKey struct {
	User      User    `json:"user" db:"users" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithBookAuthorsSurrogateKeyFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithBookAuthorsSurrogateKeyFilters(filters map[string][]any) BookAuthorsSurrogateKeySelectConfigOption {
	return func(s *BookAuthorsSurrogateKeySelectConfig) {
		s.filters = filters
	}
}

const bookAuthorsSurrogateKeyTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
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

const bookAuthorsSurrogateKeyTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books_book_id is not null), '{}') as book_authors_surrogate_key_books`

const bookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL = `book_authors_surrogate_key.author_id, book_authors_surrogate_key.book_authors_surrogate_key_id`

const bookAuthorsSurrogateKeyTableAuthorsBookJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
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

const bookAuthorsSurrogateKeyTableAuthorsBookSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') as book_authors_surrogate_key_authors`

const bookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL = `book_authors_surrogate_key.book_id, book_authors_surrogate_key.book_authors_surrogate_key_id`

// Insert inserts the BookAuthorsSurrogateKey to the database.
func (bask *BookAuthorsSurrogateKey) Insert(ctx context.Context, db DB) (*BookAuthorsSurrogateKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.book_authors_surrogate_key (` +
		`book_id, author_id, pseudonym` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, bask.BookID, bask.AuthorID, bask.Pseudonym)

	rows, err := db.Query(ctx, sqlstr, bask.BookID, bask.AuthorID, bask.Pseudonym)
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
		`book_id = $1, author_id = $2, pseudonym = $3 ` +
		`WHERE book_authors_surrogate_key_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, bask.BookID, bask.AuthorID, bask.Pseudonym, bask.BookAuthorsSurrogateKeyID)

	rows, err := db.Query(ctx, sqlstr, bask.BookID, bask.AuthorID, bask.Pseudonym, bask.BookAuthorsSurrogateKeyID)
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

	bask.BookID = params.BookID
	bask.AuthorID = params.AuthorID
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

// BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyIDAsc returns a cursor-paginated list of BookAuthorsSurrogateKey in Asc order.
func BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyIDAsc(ctx context.Context, db DB, bookAuthorsSurrogateKeyID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, bookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, bookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, bookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, bookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, ",\n") + " "
	}
	joins := ""
	if len(joinClauses) > 0 {
		joins = ", " + strings.Join(joinClauses, ",\n") + " "
	}
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = ", " + strings.Join(groupByClauses, ",\n") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym %s `+
		`FROM xo_tests.book_authors_surrogate_key %s `+
		` WHERE book_authors_surrogate_key.book_authors_surrogate_key_id > $1`+
		` %s  GROUP BY book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_id, 
book_authors_surrogate_key.author_id, 
book_authors_surrogate_key.pseudonym 
 %s 
 ORDER BY 
		book_authors_surrogate_key_id Asc `, filters, selects, joins, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookAuthorsSurrogateKeyID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyIDDesc returns a cursor-paginated list of BookAuthorsSurrogateKey in Desc order.
func BookAuthorsSurrogateKeyPaginatedByBookAuthorsSurrogateKeyIDDesc(ctx context.Context, db DB, bookAuthorsSurrogateKeyID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, bookAuthorsSurrogateKeyTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, bookAuthorsSurrogateKeyTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorsSurrogateKeyTableBooksAuthorGroupBySQL)
	}

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, bookAuthorsSurrogateKeyTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, bookAuthorsSurrogateKeyTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, bookAuthorsSurrogateKeyTableAuthorsBookGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, ",\n") + " "
	}
	joins := ""
	if len(joinClauses) > 0 {
		joins = ", " + strings.Join(joinClauses, ",\n") + " "
	}
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = ", " + strings.Join(groupByClauses, ",\n") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym %s `+
		`FROM xo_tests.book_authors_surrogate_key %s `+
		` WHERE book_authors_surrogate_key.book_authors_surrogate_key_id < $1`+
		` %s  GROUP BY book_authors_surrogate_key.book_authors_surrogate_key_id, 
book_authors_surrogate_key.book_id, 
book_authors_surrogate_key.author_id, 
book_authors_surrogate_key.pseudonym 
 %s 
 ORDER BY 
		book_authors_surrogate_key_id Desc `, filters, selects, joins, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookAuthorsSurrogateKeyID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeyByBookIDAuthorID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func BookAuthorsSurrogateKeyByBookIDAuthorID(ctx context.Context, db DB, bookID int, authorID uuid.UUID, opts ...BookAuthorsSurrogateKeySelectConfigOption) (*BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym `+
		`FROM xo_tests.book_authors_surrogate_key `+
		``+
		` WHERE book_authors_surrogate_key.book_id = $1 AND book_authors_surrogate_key.author_id = $2`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookIDAuthorID/db.Query: %w", err))
	}
	bask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectOneRow: %w", err))
	}

	return &bask, nil
}

// BookAuthorsSurrogateKeysByBookID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func BookAuthorsSurrogateKeysByBookID(ctx context.Context, db DB, bookID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym `+
		`FROM xo_tests.book_authors_surrogate_key `+
		``+
		` WHERE book_authors_surrogate_key.book_id = $1`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeysByAuthorID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_book_id_author_id_key'.
func BookAuthorsSurrogateKeysByAuthorID(ctx context.Context, db DB, authorID uuid.UUID, opts ...BookAuthorsSurrogateKeySelectConfigOption) ([]BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym `+
		`FROM xo_tests.book_authors_surrogate_key `+
		``+
		` WHERE book_authors_surrogate_key.author_id = $1`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookAuthorsSurrogateKey/BookAuthorsSurrogateKeyByBookIDAuthorID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID retrieves a row from 'xo_tests.book_authors_surrogate_key' as a BookAuthorsSurrogateKey.
//
// Generated from index 'book_authors_surrogate_key_pkey'.
func BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID(ctx context.Context, db DB, bookAuthorsSurrogateKeyID int, opts ...BookAuthorsSurrogateKeySelectConfigOption) (*BookAuthorsSurrogateKey, error) {
	c := &BookAuthorsSurrogateKeySelectConfig{joins: BookAuthorsSurrogateKeyJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`book_authors_surrogate_key.book_authors_surrogate_key_id,
book_authors_surrogate_key.book_id,
book_authors_surrogate_key.author_id,
book_authors_surrogate_key.pseudonym `+
		`FROM xo_tests.book_authors_surrogate_key `+
		``+
		` WHERE book_authors_surrogate_key.book_authors_surrogate_key_id = $1`+
		` %s  `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookAuthorsSurrogateKeyID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookAuthorsSurrogateKeyID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/db.Query: %w", err))
	}
	bask, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthorsSurrogateKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors_surrogate_key/BookAuthorsSurrogateKeyByBookAuthorsSurrogateKeyID/pgx.CollectOneRow: %w", err))
	}

	return &bask, nil
}

// FKUser_AuthorID returns the User associated with the BookAuthorsSurrogateKey's (AuthorID).
//
// Generated from foreign key 'book_authors_surrogate_key_author_id_fkey'.
func (bask *BookAuthorsSurrogateKey) FKUser_AuthorID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, bask.AuthorID)
}

// FKBook_BookID returns the Book associated with the BookAuthorsSurrogateKey's (BookID).
//
// Generated from foreign key 'book_authors_surrogate_key_book_id_fkey'.
func (bask *BookAuthorsSurrogateKey) FKBook_BookID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, bask.BookID)
}
