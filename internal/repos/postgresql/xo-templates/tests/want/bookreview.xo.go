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
	BookID   *int      `json:"bookID" db:"book_id" required:"true"`    // book_id
	Reviewer uuid.UUID `json:"reviewer" db:"reviewer" required:"true"` // reviewer

	BookJoin *Book `json:"-" db:"book" openapi-go:"ignore"` // O2O (inferred O2O - modify via `cardinality:` column comment)
}

// BookReviewCreateParams represents insert params for 'public.book_reviews'
type BookReviewCreateParams struct {
	BookID   *int      `json:"bookID"`   // book_id
	Reviewer uuid.UUID `json:"reviewer"` // reviewer
}

// BookReviewUpdateParams represents update params for 'public.book_reviews'
type BookReviewUpdateParams struct {
	BookID   **int      `json:"bookID"`   // book_id
	Reviewer *uuid.UUID `json:"reviewer"` // reviewer
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

const ()

type BookReviewJoins struct {
	Book bool
}

// WithBookReviewJoin joins with the given tables.
func WithBookReviewJoin(joins BookReviewJoins) BookReviewSelectConfigOption {
	return func(s *BookReviewSelectConfig) {
		s.joins = joins
	}
}

// BookReviewByReviewerBookID retrieves a row from 'public.book_reviews' as a BookReview.
//
// Generated from index 'book_reviews_reviewer_book_id_key'.
func BookReviewByReviewerBookID(ctx context.Context, db DB, reviewer uuid.UUID, bookID *int, opts ...BookReviewSelectConfigOption) (*BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey"
left join books on books.book_id = book_reviews.book_id` +
		` WHERE book_reviews.reviewer = $2 AND book_reviews.book_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, reviewer, bookID)
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
		`book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey"
left join books on books.book_id = book_reviews.book_id` +
		` WHERE book_reviews.reviewer = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, reviewer)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, reviewer)
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
func BookReviewsByBookID(ctx context.Context, db DB, bookID *int, opts ...BookReviewSelectConfigOption) ([]BookReview, error) {
	c := &BookReviewSelectConfig{joins: BookReviewJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_reviews.book_id,
book_reviews.reviewer,
(case when $1::boolean = true and books.book_id is not null then row(books.*) end) as book ` +
		`FROM public.book_reviews ` +
		`-- O2O join generated from "book_reviews_book_id_fkey"
left join books on books.book_id = book_reviews.book_id` +
		` WHERE book_reviews.book_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Book, bookID)
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

// FKBook_BookID returns the Book associated with the BookReview's (BookID).
//
// Generated from foreign key 'book_reviews_book_id_fkey'.
func (br *BookReview) FKBook_BookID(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, *br.BookID)
}

// FKUser_Reviewer returns the User associated with the BookReview's (Reviewer).
//
// Generated from foreign key 'book_reviews_reviewer_fkey'.
func (br *BookReview) FKUser_Reviewer(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, br.Reviewer)
}
