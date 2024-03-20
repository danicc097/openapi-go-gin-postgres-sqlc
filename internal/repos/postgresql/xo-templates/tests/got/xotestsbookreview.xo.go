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

// XoTestsBookReview represents a row from 'xo_tests.book_reviews'.
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
type XoTestsBookReview struct {
	BookReviewID XoTestsBookReviewID `json:"bookReviewID" db:"book_review_id" required:"true" nullable:"false"` // book_review_id
	BookID       XoTestsBookID       `json:"bookID" db:"book_id" required:"true" nullable:"false"`              // book_id
	Reviewer     XoTestsUserID       `json:"reviewer" db:"reviewer" required:"true" nullable:"false"`           // reviewer

	BookJoin *XoTestsBook `json:"-" db:"book_book_id" openapi-go:"ignore"`  // O2O books (generated from M2O)
	UserJoin *XoTestsUser `json:"-" db:"user_reviewer" openapi-go:"ignore"` // O2O users (generated from M2O)
}

// XoTestsBookReviewCreateParams represents insert params for 'xo_tests.book_reviews'.
type XoTestsBookReviewCreateParams struct {
	BookID   XoTestsBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Reviewer XoTestsUserID `json:"reviewer" required:"true" nullable:"false"` // reviewer
}

// XoTestsBookReviewParams represents common params for both insert and update of 'xo_tests.book_reviews'.
type XoTestsBookReviewParams interface {
	GetBookID() *XoTestsBookID
	GetReviewer() *XoTestsUserID
}

func (p XoTestsBookReviewCreateParams) GetBookID() *XoTestsBookID {
	x := p.BookID
	return &x
}

func (p XoTestsBookReviewUpdateParams) GetBookID() *XoTestsBookID {
	return p.BookID
}

func (p XoTestsBookReviewCreateParams) GetReviewer() *XoTestsUserID {
	x := p.Reviewer
	return &x
}

func (p XoTestsBookReviewUpdateParams) GetReviewer() *XoTestsUserID {
	return p.Reviewer
}

type XoTestsBookReviewID int

// CreateXoTestsBookReview creates a new XoTestsBookReview in the database with the given params.
func CreateXoTestsBookReview(ctx context.Context, db DB, params *XoTestsBookReviewCreateParams) (*XoTestsBookReview, error) {
	xtbr := &XoTestsBookReview{
		BookID:   params.BookID,
		Reviewer: params.Reviewer,
	}

	return xtbr.Insert(ctx, db)
}

type XoTestsBookReviewSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsBookReviewJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookReviewSelectConfigOption func(*XoTestsBookReviewSelectConfig)

// WithXoTestsBookReviewLimit limits row selection.
func WithXoTestsBookReviewLimit(limit int) XoTestsBookReviewSelectConfigOption {
	return func(s *XoTestsBookReviewSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsBookReviewOrderBy string

type XoTestsBookReviewJoins struct {
	Book bool `json:"book" required:"true" nullable:"false"` // O2O books
	User bool `json:"user" required:"true" nullable:"false"` // O2O users
}

// WithXoTestsBookReviewJoin joins with the given tables.
func WithXoTestsBookReviewJoin(joins XoTestsBookReviewJoins) XoTestsBookReviewSelectConfigOption {
	return func(s *XoTestsBookReviewSelectConfig) {
		s.joins = XoTestsBookReviewJoins{
			Book: s.joins.Book || joins.Book,
			User: s.joins.User || joins.User,
		}
	}
}

// WithXoTestsBookReviewFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookReviewFilters(filters map[string][]any) XoTestsBookReviewSelectConfigOption {
	return func(s *XoTestsBookReviewSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookReviewHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsBookReviewHavingClause(conditions map[string][]any) XoTestsBookReviewSelectConfigOption {
	return func(s *XoTestsBookReviewSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookReviewTableBookJoinSQL = `-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
`

const xoTestsBookReviewTableBookSelectSQL = `(case when _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id`

const xoTestsBookReviewTableBookGroupBySQL = `_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id`

const xoTestsBookReviewTableUserJoinSQL = `-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer
`

const xoTestsBookReviewTableUserSelectSQL = `(case when _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer`

const xoTestsBookReviewTableUserGroupBySQL = `_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id`

// XoTestsBookReviewUpdateParams represents update params for 'xo_tests.book_reviews'.
type XoTestsBookReviewUpdateParams struct {
	BookID   *XoTestsBookID `json:"bookID" nullable:"false"`   // book_id
	Reviewer *XoTestsUserID `json:"reviewer" nullable:"false"` // reviewer
}

// SetUpdateParams updates xo_tests.book_reviews struct fields with the specified params.
func (xtbr *XoTestsBookReview) SetUpdateParams(params *XoTestsBookReviewUpdateParams) {
	if params.BookID != nil {
		xtbr.BookID = *params.BookID
	}
	if params.Reviewer != nil {
		xtbr.Reviewer = *params.Reviewer
	}
}

// Insert inserts the XoTestsBookReview to the database.
func (xtbr *XoTestsBookReview) Insert(ctx context.Context, db DB) (*XoTestsBookReview, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.book_reviews (
	book_id, reviewer
	) VALUES (
	$1, $2
	) RETURNING * `
	// run
	logf(sqlstr, xtbr.BookID, xtbr.Reviewer)

	rows, err := db.Query(ctx, sqlstr, xtbr.BookID, xtbr.Reviewer)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Insert/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	newxtbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	*xtbr = newxtbr

	return xtbr, nil
}

// Update updates a XoTestsBookReview in the database.
func (xtbr *XoTestsBookReview) Update(ctx context.Context, db DB) (*XoTestsBookReview, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_reviews SET 
	book_id = $1, reviewer = $2 
	WHERE book_review_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, xtbr.BookID, xtbr.Reviewer, xtbr.BookReviewID)

	rows, err := db.Query(ctx, sqlstr, xtbr.BookID, xtbr.Reviewer, xtbr.BookReviewID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Update/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	newxtbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}
	*xtbr = newxtbr

	return xtbr, nil
}

// Upsert upserts a XoTestsBookReview in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtbr *XoTestsBookReview) Upsert(ctx context.Context, db DB, params *XoTestsBookReviewCreateParams) (*XoTestsBookReview, error) {
	var err error

	xtbr.BookID = params.BookID
	xtbr.Reviewer = params.Reviewer

	xtbr, err = xtbr.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book review", Err: err})
			}
			xtbr, err = xtbr.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book review", Err: err})
			}
		}
	}

	return xtbr, err
}

// Delete deletes the XoTestsBookReview from the database.
func (xtbr *XoTestsBookReview) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.book_reviews 
	WHERE book_review_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtbr.BookReviewID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookReviewPaginatedByBookReviewID returns a cursor-paginated list of XoTestsBookReview.
func XoTestsBookReviewPaginatedByBookReviewID(ctx context.Context, db DB, bookReviewID XoTestsBookReviewID, direction models.Direction, opts ...XoTestsBookReviewSelectConfigOption) ([]XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.book_review_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		book_review_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewPaginatedByBookReviewID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookReviewID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Paginated/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// XoTestsBookReviewPaginatedByBookID returns a cursor-paginated list of XoTestsBookReview.
func XoTestsBookReviewPaginatedByBookID(ctx context.Context, db DB, bookID XoTestsBookID, direction models.Direction, opts ...XoTestsBookReviewSelectConfigOption) ([]XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.book_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		book_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewPaginatedByBookID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Paginated/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// XoTestsBookReviewByBookReviewID retrieves a row from 'xo_tests.book_reviews' as a XoTestsBookReview.
//
// Generated from index 'book_reviews_pkey'.
func XoTestsBookReviewByBookReviewID(ctx context.Context, db DB, bookReviewID XoTestsBookReviewID, opts ...XoTestsBookReviewSelectConfigOption) (*XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.book_review_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewByBookReviewID */\n" + sqlstr

	// run
	// logf(sqlstr, bookReviewID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookReviewID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	xtbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	return &xtbr, nil
}

// XoTestsBookReviewByReviewerBookID retrieves a row from 'xo_tests.book_reviews' as a XoTestsBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func XoTestsBookReviewByReviewerBookID(ctx context.Context, db DB, reviewer XoTestsUserID, bookID XoTestsBookID, opts ...XoTestsBookReviewSelectConfigOption) (*XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.reviewer = $1 AND book_reviews.book_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewByReviewerBookID */\n" + sqlstr

	// run
	// logf(sqlstr, reviewer, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{reviewer, bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/db.Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	xtbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/pgx.CollectOneRow: %w", &XoError{Entity: "Book review", Err: err}))
	}

	return &xtbr, nil
}

// XoTestsBookReviewsByReviewer retrieves a row from 'xo_tests.book_reviews' as a XoTestsBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func XoTestsBookReviewsByReviewer(ctx context.Context, db DB, reviewer XoTestsUserID, opts ...XoTestsBookReviewSelectConfigOption) ([]XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.reviewer = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewsByReviewer */\n" + sqlstr

	// run
	// logf(sqlstr, reviewer)
	rows, err := db.Query(ctx, sqlstr, append([]any{reviewer}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/BookReviewByReviewerBookID/Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// XoTestsBookReviewsByBookID retrieves a row from 'xo_tests.book_reviews' as a XoTestsBookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func XoTestsBookReviewsByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookReviewSelectConfigOption) ([]XoTestsBookReview, error) {
	c := &XoTestsBookReviewSelectConfig{joins: XoTestsBookReviewJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Book {
		selectClauses = append(selectClauses, xoTestsBookReviewTableBookSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableBookJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableBookGroupBySQL)
	}

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsBookReviewTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookReviewTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookReviewTableUserGroupBySQL)
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
	book_reviews.book_id,
	book_reviews.book_review_id,
	book_reviews.reviewer %s 
	 FROM xo_tests.book_reviews %s 
	 WHERE book_reviews.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookReviewsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/BookReviewByReviewerBookID/Query: %w", &XoError{Entity: "Book review", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", &XoError{Entity: "Book review", Err: err}))
	}
	return res, nil
}

// FKBook_BookID returns the Book associated with the XoTestsBookReview's (BookID).
//
// Generated from foreign key 'book_reviews_book_id_fkey'.
func (xtbr *XoTestsBookReview) FKBook_BookID(ctx context.Context, db DB) (*XoTestsBook, error) {
	return XoTestsBookByBookID(ctx, db, xtbr.BookID)
}

// FKUser_Reviewer returns the User associated with the XoTestsBookReview's (Reviewer).
//
// Generated from foreign key 'book_reviews_reviewer_fkey'.
func (xtbr *XoTestsBookReview) FKUser_Reviewer(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtbr.Reviewer)
}
