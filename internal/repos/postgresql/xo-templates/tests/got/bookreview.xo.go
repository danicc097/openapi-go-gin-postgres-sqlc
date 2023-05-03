package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookReview represents a row from 'public.book_reviews'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type BookReview struct {
	BookReviewID int       `json:"bookReviewID" db:"book_review_id" required:"true"` // book_review_id
	BookID       int       `json:"bookID" db:"book_id" required:"true"`              // book_id
	Reviewer     uuid.UUID `json:"reviewer" db:"reviewer" required:"true"`           // reviewer

	BookJoin *Book `json:"-" db:"book" openapi-go:"ignore"` // O2O (generated from M2O)
	UserJoin *User `json:"-" db:"user" openapi-go:"ignore"` // O2O (generated from M2O)
}

// BookReviewCreateParams represents insert params for 'public.book_reviews'
type BookReviewCreateParams struct {
	BookID   int       `json:"bookID"`   // book_id
	Reviewer uuid.UUID `json:"reviewer"` // reviewer
}

// CreateBookReview creates a new BookReview in the database with the given params.
func CreateBookReview(ctx context.Context, db DB, params *BookReviewCreateParams) (*BookReview, error) {
	br := &BookReview{
		BookID:   params.BookID,
		Reviewer: params.Reviewer,
	}

	return br.Insert(ctx, db)
}

// BookReviewUpdateParams represents update params for 'public.book_reviews'
type BookReviewUpdateParams struct {
	BookID   *int       `json:"bookID"`   // book_id
	Reviewer *uuid.UUID `json:"reviewer"` // reviewer
}

// SetUpdateParams updates public.book_reviews struct fields with the specified params.
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
}
type BookReviewSelectConfigOption func(*BookReviewSelectConfig)

// WithBookReviewLimit limits row selection.
func WithBookReviewLimit(limit int) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type BookReviewOrderBy = string

type BookReviewJoins struct {
	Book bool
	User bool
}

// WithBookReviewJoin joins with the given tables.
func WithBookReviewJoin(joins BookReviewJoins) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the BookReview to the database.
func (br *BookReview) Insert(ctx context.Context, db DB) (*BookReview, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.book_reviews (` +
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
	sqlstr := `UPDATE public.book_reviews SET ` +
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

// Upsert performs an upsert for BookReview.
func (br *BookReview) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.book_reviews (` +
		`book_review_id, book_id, reviewer` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (book_review_id) DO ` +
		`UPDATE SET ` +
		`book_id = EXCLUDED.book_id, reviewer = EXCLUDED.reviewer ` +
		` RETURNING * `
	// run
	logf(sqlstr, br.BookReviewID, br.BookID, br.Reviewer)
	if _, err := db.Exec(ctx, sqlstr, br.BookReviewID, br.BookID, br.Reviewer); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the BookReview from the database.
func (br *BookReview) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.book_reviews ` +
		`WHERE book_review_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, br.BookReviewID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookReviewByBookReviewID retrieves a row from 'public.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_pkey'.
func BookReviewByBookReviewID(ctx context.Context, db DB, bookReviewID int, opts ...BookReviewSelectConfigOption) (*BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join books on books.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join users on users.user_id = book_reviews.reviewer` +
		` WHERE book_reviews.book_review_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookReviewID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, c.joins.User, bookReviewID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/db.Query: %w", err))
	}
	br, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByBookReviewID/pgx.CollectOneRow: %w", err))
	}

	return &br, nil
}

// BookReviewByReviewerBookID retrieves a row from 'public.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewByReviewerBookID(ctx context.Context, db DB, reviewer uuid.UUID, bookID int, opts ...BookReviewSelectConfigOption) (*BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join books on books.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join users on users.user_id = book_reviews.reviewer` +
		` WHERE book_reviews.reviewer = $3 AND book_reviews.book_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, c.joins.User, reviewer, bookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/db.Query: %w", err))
	}
	br, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_reviews/BookReviewByReviewerBookID/pgx.CollectOneRow: %w", err))
	}

	return &br, nil
}

// BookReviewsByReviewer retrieves a row from 'public.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewsByReviewer(ctx context.Context, db DB, reviewer uuid.UUID, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join books on books.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join users on users.user_id = book_reviews.reviewer` +
		` WHERE book_reviews.reviewer = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, c.joins.User, reviewer)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookReviewsByBookID retrieves a row from 'public.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewsByBookID(ctx context.Context, db DB, bookID int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_review_id,
book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey (Generated from M2O)"
left join books on books.book_id = book_reviews.book_id
-- O2O join generated from "book_reviews_reviewer_fkey (Generated from M2O)"
left join users on users.user_id = book_reviews.reviewer` +
		` WHERE book_reviews.book_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, c.joins.User, bookID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookReview])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}
