

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// ExtraSchemaUser represents a row from 'extra_schema.users'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//   - private to exclude a field from JSON.
//   - not-required to make a schema field not required.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaUser struct {
	UserID    ExtraSchemaUserID `json:"userID" db:"user_id" required:"true" nullable:"false"`       // user_id
	Name      string            `json:"name" db:"name" required:"true" nullable:"false"`            // name
	APIKeyID  *UserAPIKeyID     `json:"apiKeyID" db:"api_key_id"`                                   // api_key_id
	CreatedAt time.Time         `json:"createdAt" db:"created_at" required:"true" nullable:"false"` // created_at
	DeletedAt *time.Time        `json:"deletedAt" db:"deleted_at"`                                  // deleted_at

	AuthorBooksJoin           *[]Book__BA_User       `json:"-" db:"book_authors_books" openapi-go:"ignore"`                 // M2M book_authors
	AuthorBooksJoinBASK       *[]Book__BASK_User     `json:"-" db:"book_authors_surrogate_key_books" openapi-go:"ignore"`   // M2M book_authors_surrogate_key
	ReviewerBookReviewsJoin   *[]BookReview          `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O users
	SellerBooksJoin           *[]Book                `json:"-" db:"book_sellers_books" openapi-go:"ignore"`                 // M2M book_sellers
	ReceiverNotificationsJoin *[]Notification        `json:"-" db:"notifications_receiver" openapi-go:"ignore"`             // M2O users
	SenderNotificationsJoin   *[]Notification        `json:"-" db:"notifications_sender" openapi-go:"ignore"`               // M2O users
	APIKeyJoin                *UserAPIKey            `json:"-" db:"user_api_key_api_key_id" openapi-go:"ignore"`            // O2O user_api_keys (inferred)
	APIKeyJoinAKI             *UserAPIKey            `json:"-" db:"user_api_key_api_key_id" openapi-go:"ignore"`            // O2O user_api_keys (inferred)
	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_User `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"` // M2M work_item_assigned_user

}

// ExtraSchemaUserCreateParams represents insert params for 'extra_schema.users'.
type ExtraSchemaUserCreateParams struct {
	APIKeyID *UserAPIKeyID `json:"apiKeyID"`                              // api_key_id
	Name     string        `json:"name" required:"true" nullable:"false"` // name
}

type ExtraSchemaUserID struct {
	uuid.UUID
}

func NewExtraSchemaUserID(id uuid.UUID) ExtraSchemaUserID {
	return ExtraSchemaUserID{
		UUID: id,
	}
}

// CreateExtraSchemaUser creates a new ExtraSchemaUser in the database with the given params.
func CreateExtraSchemaUser(ctx context.Context, db DB, params *ExtraSchemaUserCreateParams) (*ExtraSchemaUser, error) {
	esu := &ExtraSchemaUser{
		APIKeyID: params.APIKeyID,
		Name:     params.Name,
	}

	return esu.Insert(ctx, db)
}

// ExtraSchemaUserUpdateParams represents update params for 'extra_schema.users'.
type ExtraSchemaUserUpdateParams struct {
	APIKeyID **UserAPIKeyID `json:"apiKeyID"`              // api_key_id
	Name     *string        `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates extra_schema.users struct fields with the specified params.
func (esu *ExtraSchemaUser) SetUpdateParams(params *ExtraSchemaUserUpdateParams) {
	if params.APIKeyID != nil {
		esu.APIKeyID = *params.APIKeyID
	}
	if params.Name != nil {
		esu.Name = *params.Name
	}
}

type ExtraSchemaUserSelectConfig struct {
	limit     string
	orderBy   string
	joins     ExtraSchemaUserJoins
	filters   map[string][]any
	deletedAt string
}
type ExtraSchemaUserSelectConfigOption func(*ExtraSchemaUserSelectConfig)

// WithExtraSchemaUserLimit limits row selection.
func WithExtraSchemaUserLimit(limit int) ExtraSchemaUserSelectConfigOption {
	return func(s *ExtraSchemaUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedExtraSchemaUserOnly limits result to records marked as deleted.
func WithDeletedExtraSchemaUserOnly() ExtraSchemaUserSelectConfigOption {
	return func(s *ExtraSchemaUserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type ExtraSchemaUserOrderBy string

const (
	ExtraSchemaUserCreatedAtDescNullsFirst ExtraSchemaUserOrderBy = " created_at DESC NULLS FIRST "
	ExtraSchemaUserCreatedAtDescNullsLast  ExtraSchemaUserOrderBy = " created_at DESC NULLS LAST "
	ExtraSchemaUserCreatedAtAscNullsFirst  ExtraSchemaUserOrderBy = " created_at ASC NULLS FIRST "
	ExtraSchemaUserCreatedAtAscNullsLast   ExtraSchemaUserOrderBy = " created_at ASC NULLS LAST "
	ExtraSchemaUserDeletedAtDescNullsFirst ExtraSchemaUserOrderBy = " deleted_at DESC NULLS FIRST "
	ExtraSchemaUserDeletedAtDescNullsLast  ExtraSchemaUserOrderBy = " deleted_at DESC NULLS LAST "
	ExtraSchemaUserDeletedAtAscNullsFirst  ExtraSchemaUserOrderBy = " deleted_at ASC NULLS FIRST "
	ExtraSchemaUserDeletedAtAscNullsLast   ExtraSchemaUserOrderBy = " deleted_at ASC NULLS LAST "
)

// WithExtraSchemaUserOrderBy orders results by the given columns.
func WithExtraSchemaUserOrderBy(rows ...ExtraSchemaUserOrderBy) ExtraSchemaUserSelectConfigOption {
	return func(s *ExtraSchemaUserSelectConfig) {
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

type ExtraSchemaUserJoins struct {
	BooksAuthor           bool // M2M book_authors
	BooksAuthorBooks      bool // M2M book_authors_surrogate_key
	BookReviews           bool // M2O book_reviews
	BooksSeller           bool // M2M book_sellers
	NotificationsReceiver bool // M2O notifications
	NotificationsSender   bool // M2O notifications
	UserAPIKey            bool // O2O user_api_keys
	UserAPIKeyUserAPIKeys bool // O2O user_api_keys
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
}

// WithExtraSchemaUserJoin joins with the given tables.
func WithExtraSchemaUserJoin(joins ExtraSchemaUserJoins) ExtraSchemaUserSelectConfigOption {
	return func(s *ExtraSchemaUserSelectConfig) {
		s.joins = ExtraSchemaUserJoins{
			BooksAuthor:           s.joins.BooksAuthor || joins.BooksAuthor,
			BooksAuthorBooks:      s.joins.BooksAuthorBooks || joins.BooksAuthorBooks,
			BookReviews:           s.joins.BookReviews || joins.BookReviews,
			BooksSeller:           s.joins.BooksSeller || joins.BooksSeller,
			NotificationsReceiver: s.joins.NotificationsReceiver || joins.NotificationsReceiver,
			NotificationsSender:   s.joins.NotificationsSender || joins.NotificationsSender,
			UserAPIKey:            s.joins.UserAPIKey || joins.UserAPIKey,
			UserAPIKeyUserAPIKeys: s.joins.UserAPIKeyUserAPIKeys || joins.UserAPIKeyUserAPIKeys,
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
		}
	}
}

// Book__BA_ExtraSchemaUser represents a M2M join against "extra_schema.book_authors"
type Book__BA_ExtraSchemaUser struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// Book__BASK_ExtraSchemaUser represents a M2M join against "extra_schema.book_authors_surrogate_key"
type Book__BASK_ExtraSchemaUser struct {
	Book      Book    `json:"book" db:"books" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WorkItem__WIAU_ExtraSchemaUser represents a M2M join against "extra_schema.work_item_assigned_user"
type WorkItem__WIAU_ExtraSchemaUser struct {
	WorkItem WorkItem         `json:"workItem" db:"work_items" required:"true"`
	Role     NullWorkItemRole `json:"role" db:"role" required:"true" `
}

// WithExtraSchemaUserFilters adds the given filters, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaUserFilters(filters map[string][]any) ExtraSchemaUserSelectConfigOption {
	return func(s *ExtraSchemaUserSelectConfig) {
		s.filters = filters
	}
}

const extraSchemaUserTableBooksAuthorJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
		book_authors.author_id as book_authors_author_id
		, book_authors.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		extra_schema.book_authors
	join extra_schema.books on books.book_id = book_authors.book_id
	group by
		book_authors_author_id
		, books.book_id
		, pseudonym
) as joined_book_authors_books on joined_book_authors_books.book_authors_author_id = users.user_id
`

const extraSchemaUserTableBooksAuthorSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_books.__books
		, joined_book_authors_books.pseudonym
		)) filter (where joined_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const extraSchemaUserTableBooksAuthorGroupBySQL = `users.user_id, users.user_id`

const extraSchemaUserTableBooksAuthorBooksJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_book_id_fkey"
left join (
	select
		book_authors_surrogate_key.author_id as book_authors_surrogate_key_author_id
		, book_authors_surrogate_key.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		extra_schema.book_authors_surrogate_key
	join extra_schema.books on books.book_id = book_authors_surrogate_key.book_id
	group by
		book_authors_surrogate_key_author_id
		, books.book_id
		, pseudonym
) as joined_book_authors_surrogate_key_books on joined_book_authors_surrogate_key_books.book_authors_surrogate_key_author_id = users.user_id
`

const extraSchemaUserTableBooksAuthorBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_books.__books
		, joined_book_authors_surrogate_key_books.pseudonym
		)) filter (where joined_book_authors_surrogate_key_books.__books_book_id is not null), '{}') as book_authors_surrogate_key_books`

const extraSchemaUserTableBooksAuthorBooksGroupBySQL = `users.user_id, users.user_id`

const extraSchemaUserTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_reviewer_fkey"
left join (
  select
  reviewer as book_reviews_user_id
    , array_agg(book_reviews.*) as book_reviews
  from
    extra_schema.book_reviews
  group by
        reviewer
) as joined_book_reviews on joined_book_reviews.book_reviews_user_id = users.user_id
`

const extraSchemaUserTableBookReviewsSelectSQL = `COALESCE(joined_book_reviews.book_reviews, '{}') as book_reviews`

const extraSchemaUserTableBookReviewsGroupBySQL = `joined_book_reviews.book_reviews, users.user_id`

const extraSchemaUserTableBooksSellerJoinSQL = `-- M2M join generated from "book_sellers_book_id_fkey"
left join (
	select
		book_sellers.seller as book_sellers_seller
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		extra_schema.book_sellers
	join extra_schema.books on books.book_id = book_sellers.book_id
	group by
		book_sellers_seller
		, books.book_id
) as joined_book_sellers_books on joined_book_sellers_books.book_sellers_seller = users.user_id
`

const extraSchemaUserTableBooksSellerSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_books.__books
		)) filter (where joined_book_sellers_books.__books_book_id is not null), '{}') as book_sellers_books`

const extraSchemaUserTableBooksSellerGroupBySQL = `users.user_id, users.user_id`

const extraSchemaUserTableNotificationsReceiverJoinSQL = `-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    extra_schema.notifications
  group by
        receiver
) as joined_notifications_receiver on joined_notifications_receiver.notifications_user_id = users.user_id
`

const extraSchemaUserTableNotificationsReceiverSelectSQL = `COALESCE(joined_notifications_receiver.notifications, '{}') as notifications_receiver`

const extraSchemaUserTableNotificationsReceiverGroupBySQL = `joined_notifications_receiver.notifications, users.user_id`

const extraSchemaUserTableNotificationsSenderJoinSQL = `-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , array_agg(notifications.*) as notifications
  from
    extra_schema.notifications
  group by
        sender
) as joined_notifications_sender on joined_notifications_sender.notifications_user_id = users.user_id
`

const extraSchemaUserTableNotificationsSenderSelectSQL = `COALESCE(joined_notifications_sender.notifications, '{}') as notifications_sender`

const extraSchemaUserTableNotificationsSenderGroupBySQL = `joined_notifications_sender.notifications, users.user_id`

const extraSchemaUserTableUserAPIKeyJoinSQL = `-- O2O join generated from "users_api_key_id_fkey (inferred)"
left join extra_schema.user_api_keys as _users_api_key_id on _users_api_key_id.user_api_key_id = users.api_key_id
`

const extraSchemaUserTableUserAPIKeySelectSQL = `(case when _users_api_key_id.user_api_key_id is not null then row(_users_api_key_id.*) end) as user_api_key_api_key_id`

const extraSchemaUserTableUserAPIKeyGroupBySQL = `_users_api_key_id.user_api_key_id,
      _users_api_key_id.user_api_key_id,
	users.user_id`

const extraSchemaUserTableUserAPIKeyUserAPIKeysJoinSQL = `-- O2O join generated from "users_api_key_id_fkey (inferred)"
left join extra_schema.user_api_keys as _users_api_key_id on _users_api_key_id.user_api_key_id = users.api_key_id
`

const extraSchemaUserTableUserAPIKeyUserAPIKeysSelectSQL = `(case when _users_api_key_id.user_api_key_id is not null then row(_users_api_key_id.*) end) as user_api_key_api_key_id`

const extraSchemaUserTableUserAPIKeyUserAPIKeysGroupBySQL = `_users_api_key_id.user_api_key_id,
      _users_api_key_id.user_api_key_id,
	users.user_id`

const extraSchemaUserTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
		work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
		, work_item_assigned_user.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		extra_schema.work_item_assigned_user
	join extra_schema.work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
	group by
		work_item_assigned_user_assigned_user
		, work_items.work_item_id
		, role
) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = users.user_id
`

const extraSchemaUserTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const extraSchemaUserTableWorkItemsAssignedUserGroupBySQL = `users.user_id, users.user_id`

// Insert inserts the ExtraSchemaUser to the database.
func (esu *ExtraSchemaUser) Insert(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.users (
	api_key_id, deleted_at, name
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, esu.APIKeyID, esu.DeletedAt, esu.Name)

	rows, err := db.Query(ctx, sqlstr, esu.APIKeyID, esu.DeletedAt, esu.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Insert/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newesu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	*esu = newesu

	return esu, nil
}

// Update updates a ExtraSchemaUser in the database.
func (esu *ExtraSchemaUser) Update(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.users SET 
	api_key_id = $1, deleted_at = $2, name = $3 
	WHERE user_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, esu.APIKeyID, esu.CreatedAt, esu.DeletedAt, esu.Name, esu.UserID)

	rows, err := db.Query(ctx, sqlstr, esu.APIKeyID, esu.DeletedAt, esu.Name, esu.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Update/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newesu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}
	*esu = newesu

	return esu, nil
}

// Upsert upserts a ExtraSchemaUser in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esu *ExtraSchemaUser) Upsert(ctx context.Context, db DB, params *ExtraSchemaUserCreateParams) (*ExtraSchemaUser, error) {
	var err error

	esu.APIKeyID = params.APIKeyID
	esu.Name = params.Name

	esu, err = esu.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User", Err: err})
			}
			esu, err = esu.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User", Err: err})
			}
		}
	}

	return esu, err
}

// Delete deletes the ExtraSchemaUser from the database.
func (esu *ExtraSchemaUser) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.users 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esu.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the ExtraSchemaUser from the database via 'deleted_at'.
func (esu *ExtraSchemaUser) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE extra_schema.users 
	SET deleted_at = NOW() 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esu.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	esu.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted ExtraSchemaUser from the database.
func (esu *ExtraSchemaUser) Restore(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	esu.DeletedAt = nil
	newesu, err := esu.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Restore/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return newesu, nil
}

// ExtraSchemaUserPaginatedByCreatedAt returns a cursor-paginated list of ExtraSchemaUser.
func ExtraSchemaUserPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, direction models.Direction, opts ...ExtraSchemaUserSelectConfigOption) ([]ExtraSchemaUser, error) {
	c := &ExtraSchemaUserSelectConfig{deletedAt: " null ", joins: ExtraSchemaUserJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.UserAPIKeyUserAPIKeys {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableWorkItemsAssignedUserGroupBySQL)
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
	 FROM extra_schema.users %s 
	 WHERE users.created_at %s $1
	 %s   AND users.deleted_at is %s  %s 
  ORDER BY 
		created_at %s `, selects, joins, operator, filters, c.deletedAt, groupbys, direction)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserPaginatedByCreatedAt */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Paginated/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaUser/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// ExtraSchemaUserByCreatedAt retrieves a row from 'extra_schema.users' as a ExtraSchemaUser.
//
// Generated from index 'users_created_at_key'.
func ExtraSchemaUserByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...ExtraSchemaUserSelectConfigOption) (*ExtraSchemaUser, error) {
	c := &ExtraSchemaUserSelectConfig{deletedAt: " null ", joins: ExtraSchemaUserJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.UserAPIKeyUserAPIKeys {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableWorkItemsAssignedUserGroupBySQL)
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
	 FROM extra_schema.users %s 
	 WHERE users.created_at = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	esu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &esu, nil
}

// ExtraSchemaUserByName retrieves a row from 'extra_schema.users' as a ExtraSchemaUser.
//
// Generated from index 'users_name_key'.
func ExtraSchemaUserByName(ctx context.Context, db DB, name string, opts ...ExtraSchemaUserSelectConfigOption) (*ExtraSchemaUser, error) {
	c := &ExtraSchemaUserSelectConfig{deletedAt: " null ", joins: ExtraSchemaUserJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.UserAPIKeyUserAPIKeys {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableWorkItemsAssignedUserGroupBySQL)
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
	 FROM extra_schema.users %s 
	 WHERE users.name = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	esu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByName/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &esu, nil
}

// ExtraSchemaUserByUserID retrieves a row from 'extra_schema.users' as a ExtraSchemaUser.
//
// Generated from index 'users_pkey'.
func ExtraSchemaUserByUserID(ctx context.Context, db DB, userID ExtraSchemaUserID, opts ...ExtraSchemaUserSelectConfigOption) (*ExtraSchemaUser, error) {
	c := &ExtraSchemaUserSelectConfig{deletedAt: " null ", joins: ExtraSchemaUserJoins{}, filters: make(map[string][]any)}

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
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorGroupBySQL)
	}

	if c.joins.BooksAuthorBooks {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksAuthorBooksSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksAuthorBooksJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksAuthorBooksGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaUserTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBookReviewsGroupBySQL)
	}

	if c.joins.BooksSeller {
		selectClauses = append(selectClauses, extraSchemaUserTableBooksSellerSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableBooksSellerJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableBooksSellerGroupBySQL)
	}

	if c.joins.NotificationsReceiver {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsReceiverSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsReceiverJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsReceiverGroupBySQL)
	}

	if c.joins.NotificationsSender {
		selectClauses = append(selectClauses, extraSchemaUserTableNotificationsSenderSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableNotificationsSenderJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableNotificationsSenderGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyGroupBySQL)
	}

	if c.joins.UserAPIKeyUserAPIKeys {
		selectClauses = append(selectClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableUserAPIKeyUserAPIKeysGroupBySQL)
	}

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaUserTableWorkItemsAssignedUserGroupBySQL)
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
	 FROM extra_schema.users %s 
	 WHERE users.user_id = $1
	 %s   AND users.deleted_at is %s  %s 
`, selects, joins, filters, c.deletedAt, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaUserByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	esu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &esu, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the ExtraSchemaUser's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (esu *ExtraSchemaUser) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *esu.APIKeyID)
}
