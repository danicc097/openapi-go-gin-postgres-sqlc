package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// User represents a row from 'xo_tests.users'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type User struct {
	UserID    UserID        `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id
	Name      string        `json:"name" db:"name" required:"true" nullable:"false"`            // name
	APIKeyID  *UserAPIKeyID `json:"apiKeyID" db:"api_key_id"`                                   // api_key_id
	CreatedAt time.Time     `json:"createdAt" db:"created_at" required:"true" nullable:"false"` // created_at
	DeletedAt *time.Time    `json:"deletedAt" db:"deleted_at"`                                  // deleted_at

	AuthorBooksJoin           *[]Book__BA_User       `json:"-" db:"book_authors_books" openapi-go:"ignore"`                 // M2M book_authors
	AuthorBooksJoinBASK       *[]Book__BASK_User     `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"`   // M2M book_authors_surrogate_key
	ReviewerBookReviewsJoin   *[]BookReview          `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O users
	SellerBooksJoin           *[]Book                `json:"-" db:"book_sellers_books" openapi-go:"ignore"`                 // M2M book_sellers
	ReceiverNotificationsJoin *[]Notification        `json:"-" db:"notifications_receiver" openapi-go:"ignore"`             // M2O users
	SenderNotificationsJoin   *[]Notification        `json:"-" db:"notifications_sender" openapi-go:"ignore"`               // M2O users
	APIKeyJoin                *UserAPIKey            `json:"-" db:"user_api_key_api_key_id" openapi-go:"ignore"`            // O2O user_api_keys (inferred)
	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_User `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// UserCreateParams represents insert params for 'xo_tests.users'.
type UserCreateParams struct {
	APIKeyID *UserAPIKeyID `json:"apiKeyID"`                              // api_key_id
	Name     string        `json:"name" required:"true" nullable:"false"` // name
}

type UserID struct {
	uuid.UUID
} // user_id

// CreateUser creates a new User in the database with the given params.
func CreateUser(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	u := &User{
		APIKeyID: params.APIKeyID,
		Name:     params.Name,
	}

	return u.Insert(ctx, db)
}

// UserUpdateParams represents update params for 'xo_tests.users'.
type UserUpdateParams struct {
	APIKeyID **UserAPIKeyID `json:"apiKeyID"`              // api_key_id
	Name     *string        `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.users struct fields with the specified params.
func (u *User) SetUpdateParams(params *UserUpdateParams) {
	if params.APIKeyID != nil {
		u.APIKeyID = *params.APIKeyID
	}
	if params.Name != nil {
		u.Name = *params.Name
	}
}

type UserSelectConfig struct {
	limit     string
	orderBy   string
	joins     UserJoins
	filters   map[string][]any
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

type UserOrderBy string

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
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
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
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
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
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
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

// WorkItem__WIAU_User represents a M2M join against "xo_tests.work_item_assigned_user"
type WorkItem__WIAU_User struct {
	WorkItem WorkItem         `json:"workItem" db:"work_items" required:"true"`
	Role     NullWorkItemRole `json:"role" db:"role" required:"true" `
}

// WithUserFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserFilters(filters map[string][]any) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.filters = filters
	}
}

const userTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
		book_authors.author_id as book_authors_author_id
		, book_authors.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_authors
	join xo_tests.books on books.book_id = book_authors.book_id
	group by
		book_authors_author_id
		, books.book_id
		, pseudonym
) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = users.user_id
`

const userTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const userTableBooksAuthorGroupBySQL = `users.user_id, users.user_id`

const userTableBooksAuthorBooksJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
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
) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = users.user_id
`

const userTableBooksAuthorBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books_book_id is not null), '{}') as book_authors_surrogate_key_books`

const userTableBooksAuthorBooksGroupBySQL = `users.user_id, users.user_id`

const userTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        reviewer
) as joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id
`

const userTableBookReviewsSelectSQL = `COALESCE(joined_book_reviews.book_reviews, '{}') as book_reviews`

const userTableBookReviewsGroupBySQL = `joined_book_reviews.book_reviews, users.user_id`

const userTableBooksSellerJoinSQL = `-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
		book_sellers.seller as book_sellers_seller
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_sellers
	join xo_tests.books on books.book_id = book_sellers.book_id
	group by
		book_sellers_seller
		, books.book_id
) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = users.user_id
`

const userTableBooksSellerSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books_book_id is not null), '{}') as book_sellers_books`

const userTableBooksSellerGroupBySQL = `users.user_id, users.user_id`

const userTableNotificationsReceiverJoinSQL = `-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        receiver
) as joined_notifications_receiver on joined_notifications_receiver.notifications_user_id = users.user_id
`

const userTableNotificationsReceiverSelectSQL = `COALESCE(joined_notifications_receiver.notifications, '{}') as notifications_receiver`

const userTableNotificationsReceiverGroupBySQL = `joined_notifications_receiver.notifications, users.user_id`

const userTableNotificationsSenderJoinSQL = `-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    xo_tests.notifications
  group by
        sender
) as joined_notifications_sender on joined_notifications_sender.notifications_user_id = users.user_id
`

const userTableNotificationsSenderSelectSQL = `COALESCE(joined_notifications_sender.notifications, '{}') as notifications_sender`

const userTableNotificationsSenderGroupBySQL = `joined_notifications_sender.notifications, users.user_id`

const userTableUserAPIKeyJoinSQL = `-- O2O join generated from "users_api_key_id_fkey (inferred)"
left join xo_tests.user_api_keys as _users_api_key_id on _users_api_key_id.user_api_key_id = users.api_key_id
`

const userTableUserAPIKeySelectSQL = `(case when _users_api_key_id.user_api_key_id is not null then row(_users_api_key_id.*) end) as user_api_key_api_key_id`

const userTableUserAPIKeyGroupBySQL = `_users_api_key_id.user_api_key_id,
      _users_api_key_id.user_api_key_id,
	users.user_id`

const userTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
		work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
		, work_item_assigned_user.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		xo_tests.work_item_assigned_user
	join xo_tests.work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
	group by
		work_item_assigned_user_assigned_user
		, work_items.work_item_id
		, role
) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = users.user_id
`

const userTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const userTableWorkItemsAssignedUserGroupBySQL = `users.user_id, users.user_id`

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) (*User, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.users (
	api_key_id, deleted_at, name
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, u.APIKeyID, u.DeletedAt, u.Name)

	rows, err := db.Query(ctx, sqlstr, u.APIKeyID, u.DeletedAt, u.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	*u = newu

	return u, nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) (*User, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.users SET 
	api_key_id = $1, deleted_at = $2, name = $3 
	WHERE user_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, u.APIKeyID, u.CreatedAt, u.DeletedAt, u.Name, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.APIKeyID, u.DeletedAt, u.Name, u.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}
	*u = newu

	return u, nil
}

// Upsert upserts a User in the database.
// Requires appropriate PK(s) to be set beforehand.
func (u *User) Upsert(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	var err error

	u.APIKeyID = params.APIKeyID
	u.Name = params.Name

	u, err = u.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User", Err: err})
			}
			u, err = u.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User", Err: err})
			}
		}
	}

	return u, err
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.users 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the User from the database via 'deleted_at'.
func (u *User) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE xo_tests.users 
	SET deleted_at = NOW() 
	WHERE user_id = $1 `
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
		return nil, logerror(fmt.Errorf("User/Restore/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return newu, nil
}

// UserPaginatedByCreatedAtAsc returns a cursor-paginated list of User in Asc order.
func UserPaginatedByCreatedAtAsc(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, userTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, userTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, userTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, userTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, userTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, userTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, userTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, userTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, userTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, userTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.created_at > $1
	 %s   AND users.deleted_at is %s  %s 
  ORDER BY 
		created_at Asc`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.limit
	sqlstr = "/* UserPaginatedByCreatedAtAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/Asc/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/Asc/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// UserPaginatedByCreatedAtDesc returns a cursor-paginated list of User in Desc order.
func UserPaginatedByCreatedAtDesc(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, userTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, userTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, userTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, userTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, userTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, userTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, userTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, userTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, userTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, userTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.created_at < $1
	 %s   AND users.deleted_at is %s  %s 
  ORDER BY 
		created_at Desc`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.limit
	sqlstr = "/* UserPaginatedByCreatedAtDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/Desc/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/Desc/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// UserByCreatedAt retrieves a row from 'xo_tests.users' as a User.
//
// Generated from index 'users_created_at_key'.
func UserByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, userTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, userTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, userTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, userTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, userTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, userTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, userTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, userTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, userTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, userTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.created_at = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UserByName retrieves a row from 'xo_tests.users' as a User.
//
// Generated from index 'users_name_key'.
func UserByName(ctx context.Context, db DB, name string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, userTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, userTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, userTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, userTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, userTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, userTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, userTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, userTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, userTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, userTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.name = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UserByUserID retrieves a row from 'xo_tests.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID UserID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any)}

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

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, userTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, userTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, userTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, userTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, userTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, userTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, userTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, userTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, userTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, userTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, userTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, userTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, userTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, userTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.user_id = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the User's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (u *User) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *u.APIKeyID)
}
