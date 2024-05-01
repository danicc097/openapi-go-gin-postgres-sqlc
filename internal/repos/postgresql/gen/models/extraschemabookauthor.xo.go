// Code generated by xo. DO NOT EDIT.

//lint:ignore

package models

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
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaBookAuthor struct {
	BookID    ExtraSchemaBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"`     // book_id
	AuthorID  ExtraSchemaUserID `json:"authorID" db:"author_id" required:"true" nullable:"false"` // author_id
	Pseudonym *string           `json:"pseudonym" db:"pseudonym"`                                 // pseudonym

	BooksJoin   *[]ExtraSchemaBookAuthorM2MBookBA   `json:"-" db:"book_authors_books"`   // M2M book_authors
	AuthorsJoin *[]ExtraSchemaBookAuthorM2MAuthorBA `json:"-" db:"book_authors_authors"` // M2M book_authors

}

// ExtraSchemaBookAuthorCreateParams represents insert params for 'extra_schema.book_authors'.
type ExtraSchemaBookAuthorCreateParams struct {
	AuthorID  ExtraSchemaUserID `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    ExtraSchemaBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string           `json:"pseudonym"`                                 // pseudonym
}

// ExtraSchemaBookAuthorParams represents common params for both insert and update of 'extra_schema.book_authors'.
type ExtraSchemaBookAuthorParams interface {
	GetAuthorID() *ExtraSchemaUserID
	GetBookID() *ExtraSchemaBookID
	GetPseudonym() *string
}

func (p ExtraSchemaBookAuthorCreateParams) GetAuthorID() *ExtraSchemaUserID {
	x := p.AuthorID
	return &x
}
func (p ExtraSchemaBookAuthorUpdateParams) GetAuthorID() *ExtraSchemaUserID {
	return p.AuthorID
}

func (p ExtraSchemaBookAuthorCreateParams) GetBookID() *ExtraSchemaBookID {
	x := p.BookID
	return &x
}
func (p ExtraSchemaBookAuthorUpdateParams) GetBookID() *ExtraSchemaBookID {
	return p.BookID
}

func (p ExtraSchemaBookAuthorCreateParams) GetPseudonym() *string {
	return p.Pseudonym
}
func (p ExtraSchemaBookAuthorUpdateParams) GetPseudonym() *string {
	if p.Pseudonym != nil {
		return *p.Pseudonym
	}
	return nil
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

type ExtraSchemaBookAuthorSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   ExtraSchemaBookAuthorJoins
	filters map[string][]any
	having  map[string][]any
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

// WithExtraSchemaBookAuthorOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaBookAuthorOrderBy(rows map[string]*Direction) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaBookAuthor]
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

type ExtraSchemaBookAuthorJoins struct {
	Books   bool `json:"books" required:"true" nullable:"false"`   // M2M book_authors
	Authors bool `json:"authors" required:"true" nullable:"false"` // M2M book_authors
}

// WithExtraSchemaBookAuthorJoin joins with the given tables.
func WithExtraSchemaBookAuthorJoin(joins ExtraSchemaBookAuthorJoins) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		s.joins = ExtraSchemaBookAuthorJoins{
			Books:   s.joins.Books || joins.Books,
			Authors: s.joins.Authors || joins.Authors,
		}
	}
}

// ExtraSchemaBookAuthorM2MBookBA represents a M2M join against "extra_schema.book_authors"
type ExtraSchemaBookAuthorM2MBookBA struct {
	Book      ExtraSchemaBook `json:"book" db:"books" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// ExtraSchemaBookAuthorM2MAuthorBA represents a M2M join against "extra_schema.book_authors"
type ExtraSchemaBookAuthorM2MAuthorBA struct {
	User      ExtraSchemaUser `json:"user" db:"users" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithExtraSchemaBookAuthorFilters adds the given WHERE clause conditions, which can be dynamically parameterized
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

// WithExtraSchemaBookAuthorHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaBookAuthorHavingClause(conditions map[string][]any) ExtraSchemaBookAuthorSelectConfigOption {
	return func(s *ExtraSchemaBookAuthorSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaBookAuthorTableBooksJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
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
) as xo_join_book_authors_books on xo_join_book_authors_books.book_authors_author_id = book_authors.author_id
`

const extraSchemaBookAuthorTableBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_books.__books
		, xo_join_book_authors_books.pseudonym
		)) filter (where xo_join_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const extraSchemaBookAuthorTableBooksGroupBySQL = `book_authors.author_id, book_authors.book_id, book_authors.author_id`

const extraSchemaBookAuthorTableAuthorsJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
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
) as xo_join_book_authors_authors on xo_join_book_authors_authors.book_authors_book_id = book_authors.book_id
`

const extraSchemaBookAuthorTableAuthorsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_authors.__users
		, xo_join_book_authors_authors.pseudonym
		)) filter (where xo_join_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const extraSchemaBookAuthorTableAuthorsGroupBySQL = `book_authors.book_id, book_authors.book_id, book_authors.author_id`

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
				return nil, fmt.Errorf("UpsertExtraSchemaBookAuthor/Insert: %w", &XoError{Entity: "Book author", Err: err})
			}
			esba, err = esba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaBookAuthor/Update: %w", &XoError{Entity: "Book author", Err: err})
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

// ExtraSchemaBookAuthorPaginated returns a cursor-paginated list of ExtraSchemaBookAuthor.
// At least one cursor is required.
func ExtraSchemaBookAuthorPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...ExtraSchemaBookAuthorSelectConfigOption) ([]ExtraSchemaBookAuthor, error) {
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{},
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
	field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaBookAuthor][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Paginated/cursor: %w", &XoError{Entity: "Book author", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
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
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Paginated/orderBy: %w", &XoError{Entity: "Book author", Err: fmt.Errorf("at least one sorted column is required")}))
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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Paginated/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookAuthor/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// ExtraSchemaBookAuthorByBookIDAuthorID retrieves a row from 'extra_schema.book_authors' as a ExtraSchemaBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func ExtraSchemaBookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID ExtraSchemaBookID, authorID ExtraSchemaUserID, opts ...ExtraSchemaBookAuthorSelectConfigOption) (*ExtraSchemaBookAuthor, error) {
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.book_id = $1 AND book_authors.author_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, append(filterParams, havingParams...)...)...)
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
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
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
	c := &ExtraSchemaBookAuthorSelectConfig{joins: ExtraSchemaBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, extraSchemaBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookAuthorTableAuthorsGroupBySQL)
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
	 FROM extra_schema.book_authors %s 
	 WHERE book_authors.author_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookAuthorsByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, append(filterParams, havingParams...)...)...)
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
