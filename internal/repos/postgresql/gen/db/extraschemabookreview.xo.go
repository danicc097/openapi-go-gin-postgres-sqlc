package db

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

// ExtraSchemaBookReview represents a row from 'extra_schema.book_reviews'.
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
type ExtraSchemaBookReview struct {
	BookReviewID ExtraSchemaBookReviewID `json:"bookReviewID" db:"book_review_id" required:"true" nullable:"false"` // book_review_id
	BookID       ExtraSchemaBookID       `json:"bookID" db:"book_id" required:"true" nullable:"false"`              // book_id
	Reviewer     ExtraSchemaUserID       `json:"reviewer" db:"reviewer" required:"true" nullable:"false"`           // reviewer

	BookJoin *ExtraSchemaBook `json:"-" db:"book_book_id" openapi-go:"ignore"`  // O2O books (generated from M2O)
	UserJoin *ExtraSchemaUser `json:"-" db:"user_reviewer" openapi-go:"ignore"` // O2O users (generated from M2O)

}

// ExtraSchemaBookReviewCreateParams represents insert params for 'extra_schema.book_reviews'.
type ExtraSchemaBookReviewCreateParams struct {
	BookID   ExtraSchemaBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Reviewer ExtraSchemaUserID `json:"reviewer" required:"true" nullable:"false"` // reviewer
}

// ExtraSchemaBookReviewParams represents common params for both insert and update of 'extra_schema.book_reviews'.
type ExtraSchemaBookReviewParams interface {
	GetBookID() *ExtraSchemaBookID
	GetReviewer() *ExtraSchemaUserID
}

func (p ExtraSchemaBookReviewCreateParams) GetBookID() *ExtraSchemaBookID {
	x := p.BookID
	return &x
}
func (p ExtraSchemaBookReviewUpdateParams) GetBookID() *ExtraSchemaBookID {
	return p.BookID
}

func (p ExtraSchemaBookReviewCreateParams) GetReviewer() *ExtraSchemaUserID {
	x := p.Reviewer
	return &x
}
func (p ExtraSchemaBookReviewUpdateParams) GetReviewer() *ExtraSchemaUserID {
	return p.Reviewer
}

type ExtraSchemaBookReviewID int

// CreateExtraSchemaBookReview creates a new ExtraSchemaBookReview in the database with the given params.
func CreateExtraSchemaBookReview(ctx context.Context, db DB, params *ExtraSchemaBookReviewCreateParams) (*ExtraSchemaBookReview, error) {
	esbr := &ExtraSchemaBookReview{
		BookID:   params.BookID,
		Reviewer: params.Reviewer,
	}

	return esbr.Insert(ctx, db)
}

type ExtraSchemaBookReviewSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   ExtraSchemaBookReviewJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaBookReviewSelectConfigOption func(*ExtraSchemaBookReviewSelectConfig)

// WithExtraSchemaBookReviewLimit limits row selection.
func WithExtraSchemaBookReviewLimit(limit int) ExtraSchemaBookReviewSelectConfigOption {
	return func(s *ExtraSchemaBookReviewSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithExtraSchemaBookReviewOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithExtraSchemaBookReviewOrderBy(rows map[string]*models.Direction) ExtraSchemaBookReviewSelectConfigOption {
	return func(s *ExtraSchemaBookReviewSelectConfig) {
		te := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaBookReview]
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

type ExtraSchemaBookReviewJoins struct {
	Book bool `json:"book" required:"true" nullable:"false"` // O2O books
	User bool `json:"user" required:"true" nullable:"false"` // O2O users
}

// WithExtraSchemaBookReviewJoin joins with the given tables.
func WithExtraSchemaBookReviewJoin(joins ExtraSchemaBookReviewJoins) ExtraSchemaBookReviewSelectConfigOption {
	return func(s *ExtraSchemaBookReviewSelectConfig) {
		s.joins = ExtraSchemaBookReviewJoins{
			Book: s.joins.Book || joins.Book,
			User: s.joins.User || joins.User,
		}
	}
}

// WithExtraSchemaBookReviewFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaBookReviewFilters(filters map[string][]any) ExtraSchemaBookReviewSelectConfigOption {
	return func(s *ExtraSchemaBookReviewSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaBookReviewHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithExtraSchemaBookReviewHavingClause(conditions map[string][]any) ExtraSchemaBookReviewSelectConfigOption {
	return func(s *ExtraSchemaBookReviewSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaBookReviewTableBookJoinSQL = `-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join extra_schema.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
`

const extraSchemaBookReviewTableBookSelectSQL = `(case when _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id`

const extraSchemaBookReviewTableBookGroupBySQL = `_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id`

const extraSchemaBookReviewTableUserJoinSQL = `-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join extra_schema.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer
`

const extraSchemaBookReviewTableUserSelectSQL = `(case when _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer`

const extraSchemaBookReviewTableUserGroupBySQL = `_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id`

// ExtraSchemaBookReviewUpdateParams represents update params for 'extra_schema.book_reviews'.
type ExtraSchemaBookReviewUpdateParams struct {
	BookID   *ExtraSchemaBookID `json:"bookID" nullable:"false"`   // book_id
	Reviewer *ExtraSchemaUserID `json:"reviewer" nullable:"false"` // reviewer
}

// SetUpdateParams updates extra_schema.book_reviews struct fields with the specified params.
func (esbr *ExtraSchemaBookReview) SetUpdateParams(params *ExtraSchemaBookReviewUpdateParams) {
	if params.BookID != nil {
		esbr.BookID = *params.BookID
	}
	if params.Reviewer != nil {
		esbr.Reviewer = *params.Reviewer
	}
}

// Insert inserts the ExtraSchemaBookReview to the database.
func (esbr *ExtraSchemaBookReview) Insert(ctx context.Context, db DB) (*ExtraSchemaBookReview, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.book_reviews (
	book_id, reviewer
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, esbr.BookID, esbr.Reviewer)

	rows, err := db.Query(ctx, sqlstr, esbr.BookID, esbr.Reviewer)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Insert/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	newesbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	*esbr = newesbr

	return esbr, nil
}

// Update updates a ExtraSchemaBookReview in the database.
func (esbr *ExtraSchemaBookReview) Update(ctx context.Context, db DB) (*ExtraSchemaBookReview, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.book_reviews SET 
	book_id = $1, reviewer = $2 
	WHERE book_review_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, esbr.BookID, esbr.Reviewer, esbr.BookReviewID)

	rows, err := db.Query(ctx, sqlstr, esbr.BookID, esbr.Reviewer, esbr.BookReviewID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Update/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	newesbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}
	*esbr = newesbr

	return esbr, nil
}

// Upsert upserts a ExtraSchemaBookReview in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esbr *ExtraSchemaBookReview) Upsert(ctx context.Context, db DB, params *ExtraSchemaBookReviewCreateParams) (*ExtraSchemaBookReview, error) {
	var err error

	esbr.BookID = params.BookID
	esbr.Reviewer = params.Reviewer

	esbr, err = esbr.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertExtraSchemaBookReview/Insert: %w", &XoError{Entity: "Book review", Err: err})
			}
			esbr, err = esbr.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertExtraSchemaBookReview/Update: %w", &XoError{Entity: "Book review", Err: err})
			}
		}
	}

	return esbr, err
}

// Delete deletes the ExtraSchemaBookReview from the database.
func (esbr *ExtraSchemaBookReview) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.book_reviews 
	WHERE book_review_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esbr.BookReviewID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaBookReviewPaginated returns a cursor-paginated list of ExtraSchemaBookReview.
// At least one cursor is required.
func ExtraSchemaBookReviewPaginated(ctx context.Context, db DB, cursors models.PaginationCursors, opts ...ExtraSchemaBookReviewSelectConfigOption) ([]ExtraSchemaBookReview, error) {
	c := &ExtraSchemaBookReviewSelectConfig{joins: ExtraSchemaBookReviewJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	for _, cursor := range cursors {
		if cursor.Value == nil {

			return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
		}
		field, ok := ExtraSchemaEntityFields[ExtraSchemaTableEntityExtraSchemaBookReview][cursor.Column]
		if !ok {
			return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Paginated/cursor: %w", &XoError{Entity: "Book review", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
		}

		op := "<"
		if cursor.Direction == models.DirectionAsc {
			op = ">"
		}
		c.filters[fmt.Sprintf("book_reviews.%s %s $i", field.Db, op)] = []any{*cursor.Value}
		c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts
	}

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
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Paginated/orderBy: %w", &XoError{Entity: "Book review", Err: fmt.Errorf("at least one sorted column is required")}))
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

	if c.joins.Book {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM extra_schema.book_reviews %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookReviewPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Paginated/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// ExtraSchemaBookReviewByBookReviewID retrieves a row from 'extra_schema.book_reviews' as a ExtraSchemaBookReview.
//
// Generated from index 'book_reviews_pkey'.
func ExtraSchemaBookReviewByBookReviewID(ctx context.Context, db DB, bookReviewID ExtraSchemaBookReviewID, opts ...ExtraSchemaBookReviewSelectConfigOption) (*ExtraSchemaBookReview, error) {
	c := &ExtraSchemaBookReviewSelectConfig{joins: ExtraSchemaBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM extra_schema.book_reviews %s 
	 WHERE book_reviews.book_review_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookReviewByBookReviewID */\n" + sqlstr

	// run
	// logf(sqlstr, bookReviewID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookReviewID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	esbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	return &esbr, nil
}

// ExtraSchemaBookReviewByReviewerBookID retrieves a row from 'extra_schema.book_reviews' as a ExtraSchemaBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func ExtraSchemaBookReviewByReviewerBookID(ctx context.Context, db DB, reviewer ExtraSchemaUserID, bookID ExtraSchemaBookID, opts ...ExtraSchemaBookReviewSelectConfigOption) (*ExtraSchemaBookReview, error) {
	c := &ExtraSchemaBookReviewSelectConfig{joins: ExtraSchemaBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM extra_schema.book_reviews %s 
	 WHERE book_reviews.reviewer = $1 AND book_reviews.book_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookReviewByReviewerBookID */\n" + sqlstr

	// run
	// logf(sqlstr, reviewer, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{reviewer, bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	esbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	return &esbr, nil
}

// ExtraSchemaBookReviewsByReviewer retrieves a row from 'extra_schema.book_reviews' as a ExtraSchemaBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func ExtraSchemaBookReviewsByReviewer(ctx context.Context, db DB, reviewer ExtraSchemaUserID, opts ...ExtraSchemaBookReviewSelectConfigOption) ([]ExtraSchemaBookReview, error) {
	c := &ExtraSchemaBookReviewSelectConfig{joins: ExtraSchemaBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM extra_schema.book_reviews %s 
	 WHERE book_reviews.reviewer = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookReviewsByReviewer */\n" + sqlstr

	// run
	// logf(sqlstr, reviewer)
	rows, err := db.Query(ctx, sqlstr, append([]any{reviewer}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/BookReviewByReviewerBookID/Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// ExtraSchemaBookReviewsByBookID retrieves a row from 'extra_schema.book_reviews' as a ExtraSchemaBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func ExtraSchemaBookReviewsByBookID(ctx context.Context, db DB, bookID ExtraSchemaBookID, opts ...ExtraSchemaBookReviewSelectConfigOption) ([]ExtraSchemaBookReview, error) {
	c := &ExtraSchemaBookReviewSelectConfig{joins: ExtraSchemaBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, extraSchemaBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM extra_schema.book_reviews %s 
	 WHERE book_reviews.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookReviewsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/BookReviewByReviewerBookID/Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// FKBook_BookID returns the Book associated with the ExtraSchemaBookReview's (BookID).
//
// Generated from foreign key 'book_reviews_book_id_fkey'.
func (esbr *ExtraSchemaBookReview) FKBook_BookID(ctx context.Context, db DB) (*ExtraSchemaBook, error) {
	return ExtraSchemaBookByBookID(ctx, db, esbr.BookID)
}

// FKUser_Reviewer returns the User associated with the ExtraSchemaBookReview's (Reviewer).
//
// Generated from foreign key 'book_reviews_reviewer_fkey'.
func (esbr *ExtraSchemaBookReview) FKUser_Reviewer(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, esbr.Reviewer)
}
