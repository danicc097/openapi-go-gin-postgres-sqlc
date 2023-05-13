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

// User represents a row from 'xo_tests.users'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type User struct {
	UserID    uuid.UUID  `json:"userID" db:"user_id" required:"true"`       // user_id
	Name      string     `json:"name" db:"name" required:"true"`            // name
	APIKeyID  *int       `json:"apiKeyID" db:"api_key_id" required:"true"`  // api_key_id
	CreatedAt time.Time  `json:"createdAt" db:"created_at" required:"true"` // created_at
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at" required:"true"` // deleted_at

	AuthorBooksJoin           *[]Book__BA_User   `json:"-" db:"book_authors_books" openapi-go:"ignore"`               // M2M book_authors
	AuthorBooksJoinBASK       *[]Book__BASK_User `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"` // M2M book_authors_surrogate_key
	ReviewerBookReviewsJoin   *[]BookReview      `json:"-" db:"book_reviews" openapi-go:"ignore"`                     // M2O users
	SellerBooksJoin           *[]Book            `json:"-" db:"book_sellers_books" openapi-go:"ignore"`               // M2M book_sellers
	ReceiverNotificationsJoin *[]Notification    `json:"-" db:"notifications_receiver" openapi-go:"ignore"`           // M2O users
	SenderNotificationsJoin   *[]Notification    `json:"-" db:"notifications_sender" openapi-go:"ignore"`             // M2O users
	UserJoin                  *UserAPIKey        `json:"-" db:"user_api_key_user_id" openapi-go:"ignore"`             // O2O user_api_keys (inferred)
}

// UserCreateParams represents insert params for 'xo_tests.users'.
type UserCreateParams struct {
	Name     string `json:"name" required:"true"`     // name
	APIKeyID *int   `json:"apiKeyID" required:"true"` // api_key_id
}

// CreateUser creates a new User in the database with the given params.
func CreateUser(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	u := &User{
		Name:     params.Name,
		APIKeyID: params.APIKeyID,
	}

	return u.Insert(ctx, db)
}

// UserUpdateParams represents update params for 'xo_tests.users'
type UserUpdateParams struct {
	Name     *string `json:"name" required:"true"`     // name
	APIKeyID **int   `json:"apiKeyID" required:"true"` // api_key_id
}

// SetUpdateParams updates xo_tests.users struct fields with the specified params.
func (u *User) SetUpdateParams(params *UserUpdateParams) {
	if params.Name != nil {
		u.Name = *params.Name
	}
	if params.APIKeyID != nil {
		u.APIKeyID = *params.APIKeyID
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
	BooksAuthor           bool // M2M book_authors
	BooksAuthorBooks      bool // M2M book_authors_surrogate_key
	BookReviews           bool // M2O book_reviews
	BooksSeller           bool // M2M book_sellers
	NotificationsReceiver bool // M2O notifications
	NotificationsSender   bool // M2O notifications
	UserAPIKey            bool // O2O user_api_keys
}

// WithUserJoin joins with the given tables.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = UserJoins{
			BooksAuthor:           s.joins.BooksAuthor || joins.BooksAuthor,
			BooksAuthorBooks:      s.joins.BooksAuthorBooks || joins.BooksAuthorBooks,
			BookReviews:           s.joins.BookReviews || joins.BookReviews,
			BooksSeller:           s.joins.BooksSeller || joins.BooksSeller,
			NotificationsReceiver: s.joins.NotificationsReceiver || joins.NotificationsReceiver,
			NotificationsSender:   s.joins.NotificationsSender || joins.NotificationsSender,
			UserAPIKey:            s.joins.UserAPIKey || joins.UserAPIKey,
		}
	}
}

// Book__BA_User represents a M2M join against "xo_tests.book_authors"
type Book__BA_User struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// Book__BASK_User represents a M2M join against "xo_tests.book_authors_surrogate_key"
type Book__BASK_User struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) (*User, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.users (` +
		`name, api_key_id, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, u.Name, u.APIKeyID, u.DeletedAt)

	rows, err := db.Query(ctx, sqlstr, u.Name, u.APIKeyID, u.DeletedAt)
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
	sqlstr := `UPDATE xo_tests.users SET ` +
		`name = $1, api_key_id = $2, deleted_at = $3 ` +
		`WHERE user_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, u.Name, u.APIKeyID, u.CreatedAt, u.DeletedAt, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.Name, u.APIKeyID, u.DeletedAt, u.UserID)
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
	u.APIKeyID = params.APIKeyID

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
	sqlstr := `DELETE FROM xo_tests.users ` +
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
	sqlstr := `UPDATE xo_tests.users ` +
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
users.api_key_id,
users.created_at,
users.deleted_at,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books is not null), '{}') end) as book_authors_books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books is not null), '{}') end) as book_authors_surrogate_key_books,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books,
(case when $5::boolean = true then COALESCE(joined_notifications_receiver.notifications, '{}') end) as notifications_receiver,
(case when $6::boolean = true then COALESCE(joined_notifications_sender.notifications, '{}') end) as notifications_sender,
(case when $7::boolean = true and _user_api_keys_user_ids.user_id is not null then row(_user_api_keys_user_ids.*) end) as user_api_key_user_id `+
		`FROM xo_tests.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = users.user_id

-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_id as book_authors_surrogate_key_author_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_id
    group by
			book_authors_surrogate_key_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id
-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = users.user_id

-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        receiver) joined_notifications_receiver on joined_notifications_receiver.notifications_user_id = users.user_id
-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        sender) joined_notifications_sender on joined_notifications_sender.notifications_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey(O2O inferred)"
left join xo_tests.user_api_keys as _user_api_keys_user_ids on _user_api_keys_user_ids.user_id = users.user_id`+
		` WHERE users.created_at > $8  AND users.deleted_at is %s  GROUP BY users.user_id, users.user_id, 
users.user_id, users.user_id, 
joined_book_reviews.book_reviews, users.user_id, 
users.user_id, users.user_id, 
joined_notifications_receiver.notifications, users.user_id, 
joined_notifications_sender.notifications, users.user_id, 
_user_api_keys_user_ids.user_id,
      _user_api_keys_user_ids.user_api_key_id,
	users.user_id  ORDER BY 
		created_at DESC`, c.deletedAt)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.BooksAuthor, c.joins.BooksAuthorBooks, c.joins.BookReviews, c.joins.BooksSeller, c.joins.NotificationsReceiver, c.joins.NotificationsSender, c.joins.UserAPIKey, createdAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserByCreatedAt retrieves a row from 'xo_tests.users' as a User.
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
users.api_key_id,
users.created_at,
users.deleted_at,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books is not null), '{}') end) as book_authors_books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books is not null), '{}') end) as book_authors_surrogate_key_books,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books,
(case when $5::boolean = true then COALESCE(joined_notifications_receiver.notifications, '{}') end) as notifications_receiver,
(case when $6::boolean = true then COALESCE(joined_notifications_sender.notifications, '{}') end) as notifications_sender,
(case when $7::boolean = true and _user_api_keys_user_ids.user_id is not null then row(_user_api_keys_user_ids.*) end) as user_api_key_user_id `+
		`FROM xo_tests.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = users.user_id

-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_id as book_authors_surrogate_key_author_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_id
    group by
			book_authors_surrogate_key_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id
-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = users.user_id

-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        receiver) joined_notifications_receiver on joined_notifications_receiver.notifications_user_id = users.user_id
-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        sender) joined_notifications_sender on joined_notifications_sender.notifications_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey(O2O inferred)"
left join xo_tests.user_api_keys as _user_api_keys_user_ids on _user_api_keys_user_ids.user_id = users.user_id`+
		` WHERE users.created_at = $8  AND users.deleted_at is %s   GROUP BY users.user_id, users.user_id, 
users.user_id, users.user_id, 
joined_book_reviews.book_reviews, users.user_id, 
users.user_id, users.user_id, 
joined_notifications_receiver.notifications, users.user_id, 
joined_notifications_sender.notifications, users.user_id, 
_user_api_keys_user_ids.user_id,
      _user_api_keys_user_ids.user_api_key_id,
	users.user_id `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.BooksAuthor, c.joins.BooksAuthorBooks, c.joins.BookReviews, c.joins.BooksSeller, c.joins.NotificationsReceiver, c.joins.NotificationsSender, c.joins.UserAPIKey, createdAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", err))
	}

	return &u, nil
}

// UserByUserID retrieves a row from 'xo_tests.users' as a User.
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
users.api_key_id,
users.created_at,
users.deleted_at,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books is not null), '{}') end) as book_authors_books,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books is not null), '{}') end) as book_authors_surrogate_key_books,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books is not null), '{}') end) as book_sellers_books,
(case when $5::boolean = true then COALESCE(joined_notifications_receiver.notifications, '{}') end) as notifications_receiver,
(case when $6::boolean = true then COALESCE(joined_notifications_sender.notifications, '{}') end) as notifications_sender,
(case when $7::boolean = true and _user_api_keys_user_ids.user_id is not null then row(_user_api_keys_user_ids.*) end) as user_api_key_user_id `+
		`FROM xo_tests.users `+
		`-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
			book_authors.author_id as book_authors_author_id
			, book_authors.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors
    join xo_tests.books on books.book_id = book_authors.book_id
    group by
			book_authors_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = users.user_id

-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
left join (
	select
			book_authors_surrogate_key.author_id as book_authors_surrogate_key_author_id
			, book_authors_surrogate_key.pseudonym as pseudonym
			, row(books.*) as __books
		from
			xo_tests.book_authors_surrogate_key
    join xo_tests.books on books.book_id = book_authors_surrogate_key.book_id
    group by
			book_authors_surrogate_key_author_id
			, books.book_id
			, pseudonym
  ) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = users.user_id

-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        reviewer) joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id
-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
			book_sellers.seller as book_sellers_seller
			, row(books.*) as __books
		from
			xo_tests.book_sellers
    join xo_tests.books on books.book_id = book_sellers.book_id
    group by
			book_sellers_seller
			, books.book_id
  ) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = users.user_id

-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        receiver) joined_notifications_receiver on joined_notifications_receiver.notifications_user_id = users.user_id
-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        sender) joined_notifications_sender on joined_notifications_sender.notifications_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey(O2O inferred)"
left join xo_tests.user_api_keys as _user_api_keys_user_ids on _user_api_keys_user_ids.user_id = users.user_id`+
		` WHERE users.user_id = $8  AND users.deleted_at is %s   GROUP BY users.user_id, users.user_id, 
users.user_id, users.user_id, 
joined_book_reviews.book_reviews, users.user_id, 
users.user_id, users.user_id, 
joined_notifications_receiver.notifications, users.user_id, 
joined_notifications_sender.notifications, users.user_id, 
_user_api_keys_user_ids.user_id,
      _user_api_keys_user_ids.user_api_key_id,
	users.user_id `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.BooksAuthor, c.joins.BooksAuthorBooks, c.joins.BookReviews, c.joins.BooksSeller, c.joins.NotificationsReceiver, c.joins.NotificationsSender, c.joins.UserAPIKey, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", err))
	}

	return &u, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the User's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (u *User) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *u.APIKeyID)
}
