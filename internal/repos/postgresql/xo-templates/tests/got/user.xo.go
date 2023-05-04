package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// User represents a row from 'public.users'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type User struct {
	UserID uuid.UUID `json:"userID" db:"user_id" required:"true"` // user_id
	Name   string    `json:"name" db:"name" required:"true"`      // name

	BooksJoin       *[]Book       `json:"-" db:"books" openapi-go:"ignore"`        // M2M
	BookReviewsJoin *[]BookReview `json:"-" db:"book_reviews" openapi-go:"ignore"` // M2O
}

// UserCreateParams represents insert params for 'public.users'
type UserCreateParams struct {
	Name string `json:"name" required:"true"` // name
}

// CreateUser creates a new User in the database with the given params.
func CreateUser(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	u := &User{
		Name: params.Name,
	}

	return u.Insert(ctx, db)
}

// UserUpdateParams represents update params for 'public.users'
type UserUpdateParams struct {
	Name *string `json:"name" required:"true"` // name
}

// SetUpdateParams updates public.users struct fields with the specified params.
func (u *User) SetUpdateParams(params *UserUpdateParams) {
	if params.Name != nil {
		u.Name = *params.Name
	}
}

type UserSelectConfig struct {
	limit   string
	orderBy string
	joins   UserJoins
}
type UserSelectConfigOption func(*UserSelectConfig)

// WithUserLimit limits row selection.
func WithUserLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserOrderBy = string

type UserJoins struct {
	Books       bool
	BookReviews bool
}

// WithUserJoin joins with the given tables.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = UserJoins{
			Books:       s.joins.Books || joins.Books,
			BookReviews: s.joins.BookReviews || joins.BookReviews,
		}
	}
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) (*User, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.users (` +
		`name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING * `
	// run
	logf(sqlstr, u.Name)

	rows, err := db.Query(ctx, sqlstr, u.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/db.Query: %w", err))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/pgx.CollectOneRow: %w", err))
	}

	*u = newu

	return u, nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) (*User, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.users SET ` +
		`name = $1 ` +
		`WHERE user_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, u.Name, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.Name, u.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/db.Query: %w", err))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/pgx.CollectOneRow: %w", err))
	}
	*u = newu

	return u, nil
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.users (` +
		`user_id, name` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (user_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name ` +
		` RETURNING * `
	// run
	logf(sqlstr, u.UserID, u.Name)
	if _, err := db.Exec(ctx, sqlstr, u.UserID, u.Name); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.users ` +
		`WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`users.user_id,
users.name,
(case when $1::boolean = true then COALESCE(joined_books.__books, '{}') end) as books,
(case when $2::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews ` +
		`FROM public.users ` +
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, array_agg(books.*) filter (where books.* is not null) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by book_authors_author_id
  ) as joined_books on joined_books.book_authors_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id` +
		` WHERE users.user_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Books, c.joins.BookReviews, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", err))
	}

	return &u, nil
}
