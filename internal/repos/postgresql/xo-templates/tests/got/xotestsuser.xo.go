package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

// XoTestsUser represents a row from 'xo_tests.users'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type XoTestsUser struct {
	UserID    XoTestsUserID        `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id
	Name      string               `json:"name" db:"name" required:"true" nullable:"false"`            // name
	APIKeyID  *XoTestsUserAPIKeyID `json:"apiKeyID" db:"api_key_id"`                                   // api_key_id
	CreatedAt time.Time            `json:"createdAt" db:"created_at" required:"true" nullable:"false"` // created_at
	DeletedAt *time.Time           `json:"deletedAt" db:"deleted_at"`                                  // deleted_at

	AuthorBooksJoin           *[]Book__BA_XoTestsUser       `json:"-" db:"book_authors_books" openapi-go:"ignore"`                 // M2M book_authors
	AuthorBooksJoinBASK       *[]Book__BASK_XoTestsUser     `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"`   // M2M book_authors_surrogate_key
	ReviewerBookReviewsJoin   *[]XoTestsBookReview          `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O users
	SellerBooksJoin           *[]XoTestsBook                `json:"-" db:"book_sellers_books" openapi-go:"ignore"`                 // M2M book_sellers
	ReceiverNotificationsJoin *[]XoTestsNotification        `json:"-" db:"notifications_receiver" openapi-go:"ignore"`             // M2O users
	SenderNotificationsJoin   *[]XoTestsNotification        `json:"-" db:"notifications_sender" openapi-go:"ignore"`               // M2O users
	APIKeyJoin                *XoTestsUserAPIKey            `json:"-" db:"user_api_key_api_key_id" openapi-go:"ignore"`            // O2O user_api_keys (inferred)
	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_XoTestsUser `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"` // M2M work_item_assigned_user
}

// XoTestsUserCreateParams represents insert params for 'xo_tests.users'.
type XoTestsUserCreateParams struct {
	APIKeyID *XoTestsUserAPIKeyID `json:"apiKeyID"`                              // api_key_id
	Name     string               `json:"name" required:"true" nullable:"false"` // name
}

type XoTestsUserID struct {
	uuid.UUID
}

func NewXoTestsUserID(id uuid.UUID) XoTestsUserID {
	return XoTestsUserID{
		UUID: id,
	}
}

// CreateXoTestsUser creates a new XoTestsUser in the database with the given params.
func CreateXoTestsUser(ctx context.Context, db DB, params *XoTestsUserCreateParams) (*XoTestsUser, error) {
	xtu := &XoTestsUser{
		APIKeyID: params.APIKeyID,
		Name:     params.Name,
	}

	return xtu.Insert(ctx, db)
}

// XoTestsUserUpdateParams represents update params for 'xo_tests.users'.
type XoTestsUserUpdateParams struct {
	APIKeyID **XoTestsUserAPIKeyID `json:"apiKeyID"`              // api_key_id
	Name     *string               `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.users struct fields with the specified params.
func (xtu *XoTestsUser) SetUpdateParams(params *XoTestsUserUpdateParams) {
	if params.APIKeyID != nil {
		xtu.APIKeyID = *params.APIKeyID
	}
	if params.Name != nil {
		xtu.Name = *params.Name
	}
}

type XoTestsUserSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsUserJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
}
type XoTestsUserSelectConfigOption func(*XoTestsUserSelectConfig)

// WithXoTestsUserLimit limits row selection.
func WithXoTestsUserLimit(limit int) XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedXoTestsUserOnly limits result to records marked as deleted.
func WithDeletedXoTestsUserOnly() XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type XoTestsUserOrderBy string

const (
	XoTestsUserCreatedAtDescNullsFirst XoTestsUserOrderBy = " created_at DESC NULLS FIRST "
	XoTestsUserCreatedAtDescNullsLast  XoTestsUserOrderBy = " created_at DESC NULLS LAST "
	XoTestsUserCreatedAtAscNullsFirst  XoTestsUserOrderBy = " created_at ASC NULLS FIRST "
	XoTestsUserCreatedAtAscNullsLast   XoTestsUserOrderBy = " created_at ASC NULLS LAST "
	XoTestsUserDeletedAtDescNullsFirst XoTestsUserOrderBy = " deleted_at DESC NULLS FIRST "
	XoTestsUserDeletedAtDescNullsLast  XoTestsUserOrderBy = " deleted_at DESC NULLS LAST "
	XoTestsUserDeletedAtAscNullsFirst  XoTestsUserOrderBy = " deleted_at ASC NULLS FIRST "
	XoTestsUserDeletedAtAscNullsLast   XoTestsUserOrderBy = " deleted_at ASC NULLS LAST "
)

// WithXoTestsUserOrderBy orders results by the given columns.
func WithXoTestsUserOrderBy(rows ...XoTestsUserOrderBy) XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
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

type XoTestsUserJoins struct {
	BooksAuthor           bool // M2M book_authors
	BooksAuthorBooks      bool // M2M book_authors_surrogate_key
	BookReviews           bool // M2O book_reviews
	BooksSeller           bool // M2M book_sellers
	NotificationsReceiver bool // M2O notifications
	NotificationsSender   bool // M2O notifications
	UserAPIKey            bool // O2O user_api_keys
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
}

// WithXoTestsUserJoin joins with the given tables.
func WithXoTestsUserJoin(joins XoTestsUserJoins) XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
		s.joins = XoTestsUserJoins{
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

// Book__BA_XoTestsUser represents a M2M join against "xo_tests.book_authors"
type Book__BA_XoTestsUser struct {
	Book      XoTestsBook `json:"book" db:"books" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// Book__BASK_XoTestsUser represents a M2M join against "xo_tests.book_authors_surrogate_key"
type Book__BASK_XoTestsUser struct {
	Book      XoTestsBook `json:"book" db:"books" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WorkItem__WIAU_XoTestsUser represents a M2M join against "xo_tests.work_item_assigned_user"
type WorkItem__WIAU_XoTestsUser struct {
	WorkItem XoTestsWorkItem      `json:"workItem" db:"work_items" required:"true"`
	Role     *XoTestsWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithXoTestsUserFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsUserFilters(filters map[string][]any) XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// // filter a given aggregate of assigned users to return results where at least one of them has id of userId
//
//	filters := map[string][]any{
//		"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsUserHavingClause(conditions map[string][]any) XoTestsUserSelectConfigOption {
	return func(s *XoTestsUserSelectConfig) {
		s.having = conditions
	}
}

const xoTestsUserTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
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

const xoTestsUserTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const xoTestsUserTableBooksAuthorGroupBySQL = `users.user_id, users.user_id`

const xoTestsUserTableBooksAuthorBooksJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
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

const xoTestsUserTableBooksAuthorBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books_book_id is not null), '{}') as book_authors_surrogate_key_books`

const xoTestsUserTableBooksAuthorBooksGroupBySQL = `users.user_id, users.user_id`

const xoTestsUserTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_reviewer_fkey"
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

const xoTestsUserTableBookReviewsSelectSQL = `COALESCE(joined_book_reviews.book_reviews, '{}') as book_reviews`

const xoTestsUserTableBookReviewsGroupBySQL = `joined_book_reviews.book_reviews, users.user_id`

const xoTestsUserTableBooksSellerJoinSQL = `-- M2M join generated from "book_sellers_book_id_fkey"
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

const xoTestsUserTableBooksSellerSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books_book_id is not null), '{}') as book_sellers_books`

const xoTestsUserTableBooksSellerGroupBySQL = `users.user_id, users.user_id`

const xoTestsUserTableNotificationsReceiverJoinSQL = `-- M2O join generated from "notifications_receiver_fkey"
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

const xoTestsUserTableNotificationsReceiverSelectSQL = `COALESCE(joined_notifications_receiver.notifications, '{}') as notifications_receiver`

const xoTestsUserTableNotificationsReceiverGroupBySQL = `joined_notifications_receiver.notifications, users.user_id`

const xoTestsUserTableNotificationsSenderJoinSQL = `-- M2O join generated from "notifications_sender_fkey"
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

const xoTestsUserTableNotificationsSenderSelectSQL = `COALESCE(joined_notifications_sender.notifications, '{}') as notifications_sender`

const xoTestsUserTableNotificationsSenderGroupBySQL = `joined_notifications_sender.notifications, users.user_id`

const xoTestsUserTableUserAPIKeyJoinSQL = `-- O2O join generated from "users_api_key_id_fkey (inferred)"
left join xo_tests.user_api_keys as _users_api_key_id on _users_api_key_id.user_api_key_id = users.api_key_id
`

const xoTestsUserTableUserAPIKeySelectSQL = `(case when _users_api_key_id.user_api_key_id is not null then row(_users_api_key_id.*) end) as user_api_key_api_key_id`

const xoTestsUserTableUserAPIKeyGroupBySQL = `_users_api_key_id.user_api_key_id,
      _users_api_key_id.user_api_key_id,
	users.user_id`

const xoTestsUserTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
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

const xoTestsUserTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const xoTestsUserTableWorkItemsAssignedUserGroupBySQL = `users.user_id, users.user_id`

// Insert inserts the XoTestsUser to the database.
func (xtu *XoTestsUser) Insert(ctx context.Context, db DB) (*XoTestsUser, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.users (
	api_key_id, deleted_at, name
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtu.APIKeyID, xtu.DeletedAt, xtu.Name)

	rows, err := db.Query(ctx, sqlstr, xtu.APIKeyID, xtu.DeletedAt, xtu.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Insert/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newxtu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	*xtu = newxtu

	return xtu, nil
}

// Update updates a XoTestsUser in the database.
func (xtu *XoTestsUser) Update(ctx context.Context, db DB) (*XoTestsUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.users SET 
	api_key_id = $1, deleted_at = $2, name = $3 
	WHERE user_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, xtu.APIKeyID, xtu.CreatedAt, xtu.DeletedAt, xtu.Name, xtu.UserID)

	rows, err := db.Query(ctx, sqlstr, xtu.APIKeyID, xtu.DeletedAt, xtu.Name, xtu.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Update/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newxtu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}
	*xtu = newxtu

	return xtu, nil
}

// Upsert upserts a XoTestsUser in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtu *XoTestsUser) Upsert(ctx context.Context, db DB, params *XoTestsUserCreateParams) (*XoTestsUser, error) {
	var err error

	xtu.APIKeyID = params.APIKeyID
	xtu.Name = params.Name

	xtu, err = xtu.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User", Err: err})
			}
			xtu, err = xtu.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User", Err: err})
			}
		}
	}

	return xtu, err
}

// Delete deletes the XoTestsUser from the database.
func (xtu *XoTestsUser) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.users 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtu.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the XoTestsUser from the database via 'deleted_at'.
func (xtu *XoTestsUser) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE xo_tests.users 
	SET deleted_at = NOW() 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtu.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	xtu.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted XoTestsUser from the database.
func (xtu *XoTestsUser) Restore(ctx context.Context, db DB) (*XoTestsUser, error) {
	xtu.DeletedAt = nil
	newxtu, err := xtu.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Restore/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return newxtu, nil
}

// XoTestsUserPaginatedByCreatedAt returns a cursor-paginated list of XoTestsUser.
func XoTestsUserPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, direction models.Direction, opts ...XoTestsUserSelectConfigOption) ([]XoTestsUser, error) {
	c := &XoTestsUserSelectConfig{deletedAt: " null ", joins: XoTestsUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, xoTestsUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, xoTestsUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableWorkItemsAssignedUserGroupBySQL)
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
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.name,
	users.user_id %s 
	 FROM xo_tests.users %s 
	 WHERE users.created_at %s $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
  ORDER BY 
		created_at %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserPaginatedByCreatedAt */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// XoTestsUserByCreatedAt retrieves a row from 'xo_tests.users' as a XoTestsUser.
//
// Generated from index 'users_created_at_key'.
func XoTestsUserByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...XoTestsUserSelectConfigOption) (*XoTestsUser, error) {
	c := &XoTestsUserSelectConfig{deletedAt: " null ", joins: XoTestsUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, xoTestsUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, xoTestsUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableWorkItemsAssignedUserGroupBySQL)
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
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	xtu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &xtu, nil
}

// XoTestsUserByName retrieves a row from 'xo_tests.users' as a XoTestsUser.
//
// Generated from index 'users_name_key'.
func XoTestsUserByName(ctx context.Context, db DB, name string, opts ...XoTestsUserSelectConfigOption) (*XoTestsUser, error) {
	c := &XoTestsUserSelectConfig{deletedAt: " null ", joins: XoTestsUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, xoTestsUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, xoTestsUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableWorkItemsAssignedUserGroupBySQL)
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
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	xtu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &xtu, nil
}

// XoTestsUserByUserID retrieves a row from 'xo_tests.users' as a XoTestsUser.
//
// Generated from index 'users_pkey'.
func XoTestsUserByUserID(ctx context.Context, db DB, userID XoTestsUserID, opts ...XoTestsUserSelectConfigOption) (*XoTestsUser, error) {
	c := &XoTestsUserSelectConfig{deletedAt: " null ", joins: XoTestsUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.BooksAuthor {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, xoTestsUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, xoTestsUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, xoTestsUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, xoTestsUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, xoTestsUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsUserTableWorkItemsAssignedUserGroupBySQL)
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
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsUserByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	xtu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &xtu, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the XoTestsUser's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (xtu *XoTestsUser) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*XoTestsUserAPIKey, error) {
	return XoTestsUserAPIKeyByUserAPIKeyID(ctx, db, *xtu.APIKeyID)
}
