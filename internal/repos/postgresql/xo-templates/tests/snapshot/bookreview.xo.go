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

// BookReview represents a row from 'xo_tests.book_reviews'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type BookReview struct {
	BookReviewID int       `json:"bookReviewID" db:"book_review_id" required:"true"` // book_review_id
	BookID       int       `json:"bookID" db:"book_id" required:"true"`              // book_id
	Reviewer     uuid.UUID `json:"reviewer" db:"reviewer" required:"true"`           // reviewer

	BookJoin     *Book `json:"-" db:"book_book_id" openapi-go:"ignore"`  // O2O books (generated from M2O)
	ReviewerJoin *User `json:"-" db:"user_reviewer" openapi-go:"ignore"` // O2O users (generated from M2O)
}

// BookReviewCreateParams represents insert params for 'xo_tests.book_reviews'.
type BookReviewCreateParams struct {
	BookID   int       `json:"bookID" required:"true"`   // book_id
	Reviewer uuid.UUID `json:"reviewer" required:"true"` // reviewer
}

// CreateBookReview creates a new BookReview in the database with the given params.
func CreateBookReview(ctx context.Context, db DB, params *BookReviewCreateParams) (*BookReview, error) {
	br := &BookReview{
		BookID:   params.BookID,
		Reviewer: params.Reviewer,
	}

	return br.Insert(ctx, db)
}

// BookReviewUpdateParams represents update params for 'xo_tests.book_reviews'
type BookReviewUpdateParams struct {
	BookID   *int       `json:"bookID" required:"true"`   // book_id
	Reviewer *uuid.UUID `json:"reviewer" required:"true"` // reviewer
}

// SetUpdateParams updates xo_tests.book_reviews struct fields with the specified params.
func (br *BookReview) SetUpdateParams(params *BookReviewUpdateParams) {
	if params.BookID != nil {
		br.BookID = *params.BookID
	}
	if params.Reviewer != nil {
		br.Reviewer = *params.Reviewer
	}
}

type BookReviewSelectConfig struct {
	limit   string
	orderBy string
	joins   BookReviewJoins
	filters map[string][]any
}
type BookReviewSelectConfigOption func(*BookReviewSelectConfig)

// WithBookReviewLimit limits row selection.
func WithBookReviewLimit(limit int) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type BookReviewOrderBy = string

type BookReviewJoins struct {
	Book bool // O2O books
	User bool // O2O users
}

// WithBookReviewJoin joins with the given tables.
func WithBookReviewJoin(joins BookReviewJoins) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		s.joins = BookReviewJoins{
			Book: s.joins.Book || joins.Book,
			User: s.joins.User || joins.User,
		}
	}
}

// WithBookReviewFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithBookReviewFilters(filters map[string][]any) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the BookReview to the database.
func (br *BookReview) Insert(ctx context.Context, db DB) (*BookReview, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.book_reviews (` +
		`book_id, reviewer` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING * `
	// run
	logf(sqlstr, br.BookID, br.Reviewer)

	rows, err := db.Query(ctx, sqlstr, br.BookID, br.Reviewer)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Insert/db.Query: %w", err))
	}
	newbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Insert/pgx.CollectOneRow: %w", err))
	}

	*br = newbr

	return br, nil
}

// Update updates a BookReview in the database.
func (br *BookReview) Update(ctx context.Context, db DB) (*BookReview, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_reviews SET ` +
		`book_id = $1, reviewer = $2 ` +
		`WHERE book_review_id = $3 ` +
		`RETURNING * `
	// run
	logf(sqlstr, br.BookID, br.Reviewer, br.BookReviewID)

	rows, err := db.Query(ctx, sqlstr, br.BookID, br.Reviewer, br.BookReviewID)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Update/db.Query: %w", err))
	}
	newbr, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Update/pgx.CollectOneRow: %w", err))
	}
	*br = newbr

	return br, nil
}

// Upsert upserts a BookReview in the database.
// Requires appropiate PK(s) to be set beforehand.
func (br *BookReview) Upsert(ctx context.Context, db DB, params *BookReviewCreateParams) (*BookReview, error) {
	var err error

	br.BookID = params.BookID
	br.Reviewer = params.Reviewer

	br, err = br.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			br, err = br.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return br, err
}

// Delete deletes the BookReview from the database.
func (br *BookReview) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.book_reviews ` +
		`WHERE book_review_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, br.BookReviewID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookReviewPaginatedByBookReviewIDAsc returns a cursor-paginated list of BookReview in Asc order.
func BookReviewPaginatedByBookReviewIDAsc(ctx context.Context, db DB, bookReviewID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_review_id > $3`+
		` %s  GROUP BY book_reviews.book_review_id, 
book_reviews.book_id, 
book_reviews.reviewer, 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id ORDER BY 
		book_review_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookReviewID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewPaginatedByBookIDAsc returns a cursor-paginated list of BookReview in Asc order.
func BookReviewPaginatedByBookIDAsc(ctx context.Context, db DB, bookID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_id > $3`+
		` %s  GROUP BY book_reviews.book_review_id, 
book_reviews.book_id, 
book_reviews.reviewer, 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id ORDER BY 
		book_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewPaginatedByBookReviewIDDesc returns a cursor-paginated list of BookReview in Desc order.
func BookReviewPaginatedByBookReviewIDDesc(ctx context.Context, db DB, bookReviewID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_review_id < $3`+
		` %s  GROUP BY book_reviews.book_review_id, 
book_reviews.book_id, 
book_reviews.reviewer, 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id ORDER BY 
		book_review_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookReviewID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewPaginatedByBookIDDesc returns a cursor-paginated list of BookReview in Desc order.
func BookReviewPaginatedByBookIDDesc(ctx context.Context, db DB, bookID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_id < $3`+
		` %s  GROUP BY book_reviews.book_review_id, 
book_reviews.book_id, 
book_reviews.reviewer, 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id ORDER BY 
		book_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewByBookReviewID retrieves a row from 'xo_tests.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_pkey'.
func BookReviewByBookReviewID(ctx context.Context, db DB, bookReviewID int, opts ...BookReviewSelectConfigOption) (*BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_review_id = $3`+
		` %s  GROUP BY 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookReviewID)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookReviewID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/db.Query: %w", err))
	}
	br, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/pgx.CollectOneRow: %w", err))
	}

	return &br, nil
}

// BookReviewByReviewerBookID retrieves a row from 'xo_tests.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewByReviewerBookID(ctx context.Context, db DB, reviewer uuid.UUID, bookID int, opts ...BookReviewSelectConfigOption) (*BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 4
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.reviewer = $3 AND book_reviews.book_id = $4`+
		` %s  GROUP BY 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, reviewer, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/db.Query: %w", err))
	}
	br, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/pgx.CollectOneRow: %w", err))
	}

	return &br, nil
}

// BookReviewsByReviewer retrieves a row from 'xo_tests.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewsByReviewer(ctx context.Context, db DB, reviewer uuid.UUID, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.reviewer = $3`+
		` %s  GROUP BY 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, reviewer}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/BookReviewByReviewerBookID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewsByBookID retrieves a row from 'xo_tests.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewsByBookID(ctx context.Context, db DB, bookID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 3
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
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and _book_reviews_book_id.book_id is not null then row(_book_reviews_book_id.*) end) as book_book_id,
(case when $2::boolean = true and _book_reviews_reviewer.user_id is not null then row(_book_reviews_reviewer.*) end) as user_reviewer `+
		`FROM xo_tests.book_reviews `+
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join xo_tests.books as _book_reviews_book_id on _book_reviews_book_id.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join xo_tests.users as _book_reviews_reviewer on _book_reviews_reviewer.user_id = book_reviews.reviewer`+
		` WHERE book_reviews.book_id = $3`+
		` %s  GROUP BY 
_book_reviews_book_id.book_id,
      _book_reviews_book_id.book_id,
	book_reviews.book_review_id, 
_book_reviews_reviewer.user_id,
      _book_reviews_reviewer.user_id,
	book_reviews.book_review_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.Book, c.joins.User, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/BookReviewByReviewerBookID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("BookReview/BookReviewByReviewerBookID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// FKBook_BookID returns the Book associated with the BookReview's (BookID).
//
// Generated from foreign key 'book_reviews_book_id_fkey'.
func (br *BookReview) FKBook_BookID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, br.BookID)
}

// FKUser_Reviewer returns the User associated with the BookReview's (Reviewer).
//
// Generated from foreign key 'book_reviews_reviewer_fkey'.
func (br *BookReview) FKUser_Reviewer(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, br.Reviewer)
}
