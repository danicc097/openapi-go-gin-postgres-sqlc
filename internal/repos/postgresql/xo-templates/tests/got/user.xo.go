package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// User represents a row from 'public.users'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type User struct {
	UserID    uuid.UUID  `json:"userID" db:"user_id" required:"true"`       // user_id
	Name      string     `json:"name" db:"name" required:"true"`            // name
	CreatedAt time.Time  `json:"createdAt" db:"created_at" required:"true"` // created_at
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at" required:"true"` // updated_at
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at" required:"true"` // deleted_at

	BooksJoin       *[]Book       `json:"-" db:"books" openapi-go:"ignore"`        // M2M
	BookReviewsJoin *[]BookReview `json:"-" db:"book_reviews" openapi-go:"ignore"` // M2O
}

// UserCreateParams represents insert params for 'public.users'.
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
	limit     string
	orderBy   string
	joins     UserJoins
	deletedAt string
}
type UserSelectConfigOption func(*UserSelectConfig)

// WithUserLimit limits row selection.
func WithUserLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedUserOnly limits result to records marked as deleted.
func WithDeletedUserOnly() UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type UserOrderBy = string

const (
	UserCreatedAtDescNullsFirst UserOrderBy = " created_at DESC NULLS FIRST "
	UserCreatedAtDescNullsLast  UserOrderBy = " created_at DESC NULLS LAST "
	UserCreatedAtAscNullsFirst  UserOrderBy = " created_at ASC NULLS FIRST "
	UserCreatedAtAscNullsLast   UserOrderBy = " created_at ASC NULLS LAST "
	UserUpdatedAtDescNullsFirst UserOrderBy = " updated_at DESC NULLS FIRST "
	UserUpdatedAtDescNullsLast  UserOrderBy = " updated_at DESC NULLS LAST "
	UserUpdatedAtAscNullsFirst  UserOrderBy = " updated_at ASC NULLS FIRST "
	UserUpdatedAtAscNullsLast   UserOrderBy = " updated_at ASC NULLS LAST "
	UserDeletedAtDescNullsFirst UserOrderBy = " deleted_at DESC NULLS FIRST "
	UserDeletedAtDescNullsLast  UserOrderBy = " deleted_at DESC NULLS LAST "
	UserDeletedAtAscNullsFirst  UserOrderBy = " deleted_at ASC NULLS FIRST "
	UserDeletedAtAscNullsLast   UserOrderBy = " deleted_at ASC NULLS LAST "
)

// WithUserOrderBy orders results by the given columns.
func WithUserOrderBy(rows ...UserOrderBy) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

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
		`name, deleted_at` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING * `
	// run
	logf(sqlstr, u.Name, u.DeletedAt)

	rows, err := db.Query(ctx, sqlstr, u.Name, u.DeletedAt)
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
		`name = $1, deleted_at = $2 ` +
		`WHERE user_id = $3 ` +
		`RETURNING * `
	// run
	logf(sqlstr, u.Name, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.Name, u.DeletedAt, u.UserID)
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

// Upsert upserts a User in the database.
// Requires appropiate PK(s) to be set beforehand.
func (u *User) Upsert(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	var err error

	u.Name = params.Name

	u, err = u.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			u, err = u.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return u, err
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

// SoftDelete soft deletes the User from the database via 'deleted_at'.
func (u *User) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.users ` +
		`SET deleted_at = NOW() ` +
		`WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted User from the database.
func (u *User) Restore(ctx context.Context, db DB) (*User, error) {
	u.DeletedAt = nil
	newu, err := u.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Restore/pgx.CollectRows: %w", err))
	}
	return newu, nil
}

// UserPaginatedByCreatedAt returns a cursor-paginated list of User.
func UserPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.name,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then ARRAY_AGG((
		joined_books.__books
		)) end) as books,
(case when $2::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews `+
		`FROM public.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, row(books.*) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
  ) as joined_books on joined_books.book_authors_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id`+
		` WHERE users.created_at > $3  AND users.deleted_at is %s  ORDER BY 
		created_at DESC`, c.deletedAt)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, createdAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserByCreatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_created_at_key'.
func UserByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.name,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then ARRAY_AGG((
		joined_books.__books
		)) end) as books,
(case when $2::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews `+
		`FROM public.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, row(books.*) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
  ) as joined_books on joined_books.book_authors_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id`+
		` WHERE users.created_at = $3  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.Books, c.joins.BookReviews, createdAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", err))
	}

	return &u, nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.name,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then ARRAY_AGG((
		joined_books.__books
		)) end) as books,
(case when $2::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews `+
		`FROM public.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, row(books.*) as __books
		from book_authors
    	join books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
  ) as joined_books on joined_books.book_authors_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id`+
		` WHERE users.user_id = $3  AND users.deleted_at is %s `, c.deletedAt)
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
