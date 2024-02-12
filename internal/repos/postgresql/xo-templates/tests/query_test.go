//go:build !skip_xo

// Package tests is meant to be run via `project test.xo` and excluded from test runs
package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/xo-templates/tests/got"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
 *
* TODO: test extensively:
*
* - order bys
* - index queries
* - join table name clash for O2O constraint too:
    name clash probably needs to be detected between constraints, check M2M-M2O and M2O-O2O
    at the same time
* IMPORTANT: explain analyze to ensure dynamic sql query plans for joins dont do hash joins

FIXME:
- cache with type: annot and ignore-constraint should generate pagination.
- might be broken since FK ref is not in public schema
*/

func TestCursorPagination_Timestamp(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ee, err := db.XoTestsPagElementPaginatedByCreatedAt(ctx, testPool, time.Now().Add(-(24+1)*time.Hour), models.DirectionDesc, db.WithXoTestsPagElementLimit(1), db.WithXoTestsPagElementJoin(db.XoTestsPagElementJoins{}))
	require.NoError(t, err)
	require.Len(t, ee, 1)
	assert.Equal(t, "element -2 days", ee[0].Name)

	ee, err = db.XoTestsPagElementPaginatedByCreatedAt(ctx, testPool, ee[0].CreatedAt, models.DirectionDesc, db.WithXoTestsPagElementLimit(2))
	require.NoError(t, err)
	require.Len(t, ee, 2)
	assert.Equal(t, "element -3 days", ee[0].Name)
	assert.Equal(t, "element -4 days", ee[1].Name)
}

// due to created_at unique.
func createUserWithRetry(t *testing.T, params *db.XoTestsUserCreateParams) *db.XoTestsUser {
	var err error

	u, err := db.CreateXoTestsUser(context.Background(), testPool, params)

	retries := 0
	for err != nil && retries < 10 {
		u, err = db.CreateXoTestsUser(context.Background(), testPool, params)
		retries++
	}
	require.NoError(t, err)

	return u
}

func TestSharedRefConstraints(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// generated with refs-ignore,share-ref-constraints
	ee, err := db.XoTestsCacheDemoWorkItemPaginatedByWorkItemID(ctx, testPool, 0 /* should filter all */, models.DirectionAsc,
		db.WithXoTestsCacheDemoWorkItemJoin(db.XoTestsCacheDemoWorkItemJoins{AssignedUsers: true}),
	)
	require.NoError(t, err)
	require.Len(t, ee, 1)
	require.EqualValues(t, 1, ee[0].WorkItemID)
	require.NotNil(t, ee[0].WorkItemAssignedUsersJoin)
	require.Len(t, *ee[0].WorkItemAssignedUsersJoin, 2)
	require.Nil(t, ee[0].WorkItemWorkItemCommentsJoin)
}

func TestCursorPagination_HavingClause(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u1 := createUserWithRetry(t, &db.XoTestsUserCreateParams{Name: t.Name() + "_1"})
	u2 := createUserWithRetry(t, &db.XoTestsUserCreateParams{Name: t.Name() + "_2"})

	wi, err := db.CreateXoTestsWorkItem(ctx, testPool, &db.XoTestsWorkItemCreateParams{})
	require.NoError(t, err)

	_, err = db.CreateXoTestsWorkItemAssignedUser(ctx, testPool, &db.XoTestsWorkItemAssignedUserCreateParams{
		WorkItemID:   wi.WorkItemID,
		AssignedUser: u1.UserID,
		XoTestsRole:  pointers.New(db.XoTestsWorkItemRolePreparer),
	})
	require.NoError(t, err)
	_, err = db.CreateXoTestsWorkItemAssignedUser(ctx, testPool, &db.XoTestsWorkItemAssignedUserCreateParams{
		WorkItemID:   wi.WorkItemID,
		AssignedUser: u2.UserID,
		XoTestsRole:  pointers.New(db.XoTestsWorkItemRolePreparer),
	})
	require.NoError(t, err)

	ee, err := db.XoTestsWorkItemPaginatedByWorkItemID(ctx, testPool, 0 /* should filter all */, models.DirectionAsc,
		db.WithXoTestsWorkItemJoin(db.XoTestsWorkItemJoins{AssignedUsers: true}),
		db.WithXoTestsWorkItemHavingClause(map[string][]any{
			"$i = ANY(ARRAY_AGG(joined_work_item_assigned_user_assigned_users.__users_user_id))": {u1.UserID},
		}),
	)
	require.NoError(t, err)
	require.Len(t, ee, 1)
	assert.Equal(t, ee[0].WorkItemID, wi.WorkItemID)

	au := *ee[0].WorkItemAssignedUsersJoin
	found := false
	for _, u := range au {
		if u.User.UserID == u1.UserID {
			found = true
		}
	}
	require.Len(t, au, 2) // should include all users, we're just filtering for work items that contain it
	assert.True(t, found)
}

func Test_Filters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ee, err := db.XoTestsPagElementPaginatedByCreatedAt(ctx, testPool, time.Now().Add(-(24+1)*time.Hour), models.DirectionDesc, db.WithXoTestsPagElementLimit(1), db.WithXoTestsPagElementJoin(db.XoTestsPagElementJoins{}))
	require.NoError(t, err)
	require.Len(t, ee, 1)
	assert.Equal(t, ee[0].Name, "element -2 days")

	ee, err = db.XoTestsPagElementPaginatedByCreatedAt(ctx, testPool, ee[0].CreatedAt, models.DirectionDesc, db.WithXoTestsPagElementLimit(2))
	require.NoError(t, err)
	require.Len(t, ee, 2)
	assert.Equal(t, ee[0].Name, "element -3 days")
	assert.Equal(t, ee[1].Name, "element -4 days")
}

func TestTrigram_Filters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ww, err := db.XoTestsWorkItems(ctx, testPool, db.WithXoTestsWorkItemFilters(map[string][]any{"description ILIKE  '%' || $1 || '%'": {"rome"}}))
	require.NoError(t, err)
	require.Len(t, ww, 1)
	assert.Contains(t, *ww[0].Description, "Rome")
}

func TestM2M_SelectFilter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	wi, err := db.XoTestsWorkItemByWorkItemID(ctx, testPool, 1, db.WithXoTestsWorkItemJoin(db.XoTestsWorkItemJoins{AssignedUsers: true}))
	require.NoError(t, err)
	assert.NotNil(t, *wi.WorkItemAssignedUsersJoin)
	require.Len(t, *wi.WorkItemAssignedUsersJoin, 2)
	for _, member := range *wi.WorkItemAssignedUsersJoin {
		uid := db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))
		if member.User.UserID == uid {
			assert.Nil(t, member.User.DeletedAt) // ensure proper filter clause used. e.g. filter where record is not null will exclude the whole record if just one element is null, see https://github.com/danicc097/openapi-go-gin-postgres-sqlc/blob/7a9affbccc9738e728ba5532d055230f4668034c/FIXME.md#L44
			assert.True(t, db.XoTestsWorkItemRolePreparer == *member.Role)
		}
	}
}

func TestM2M_TwoFKsAndExtraColumns(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksAuthor: true}))
	require.NoError(t, err)
	require.Len(t, *u.AuthorBooksJoin, 0)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")))
	require.NoError(t, err)
	assert.Nil(t, u.AuthorBooksJoin)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksAuthor: true}))
	require.NoError(t, err)
	require.Len(t, *u.AuthorBooksJoin, 2)
	for _, b := range *u.AuthorBooksJoin {
		if b.Book.BookID == 1 {
			assert.Equal(t, "not Jane Smith", *b.Pseudonym)
		}
	}
}

func TestM2M_SurrogatePK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksAuthorBooks: true}))
	require.NoError(t, err)
	require.Len(t, *u.AuthorBooksJoinBASK, 0)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")))
	require.NoError(t, err)
	assert.Nil(t, u.AuthorBooksJoinBASK)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksAuthorBooks: true}))
	require.NoError(t, err)
	require.Len(t, *u.AuthorBooksJoinBASK, 2)
	for _, b := range *u.AuthorBooksJoinBASK {
		if b.Book.BookID == 1 {
			assert.Equal(t, *b.Pseudonym, "not Jane Smith")
		}
	}
}

func TestM2M_TwoFKs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksSeller: true}))
	require.NoError(t, err)
	require.Len(t, *u.SellerBooksJoin, 0)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71")))
	require.NoError(t, err)
	assert.Nil(t, u.SellerBooksJoin)

	u, err = db.XoTestsUserByUserID(ctx, testPool, db.NewXoTestsUserID(uuid.MustParse("8c67f1f9-2be4-4b1a-a49b-b7a10a60c53a")), db.WithXoTestsUserJoin(db.XoTestsUserJoins{BooksSeller: true}))
	require.NoError(t, err)
	require.Len(t, *u.SellerBooksJoin, 1)
	assert.EqualValues(t, (*u.SellerBooksJoin)[0].BookID, 1)
}

func TestM2O(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userID := db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))

	u, err := db.XoTestsUserByUserID(ctx, testPool, userID, db.WithXoTestsUserJoin(db.XoTestsUserJoins{NotificationsSender: true, NotificationsReceiver: true}))
	require.NoError(t, err)
	require.Len(t, *u.ReceiverNotificationsJoin, 1)
	require.Len(t, *u.SenderNotificationsJoin, 2)

	n, err := db.XoTestsNotificationsBySender(ctx, testPool, userID, db.WithXoTestsNotificationJoin(db.XoTestsNotificationJoins{UserSender: true, UserReceiver: true}))
	require.NoError(t, err)
	require.Len(t, n, 2)
	assert.Equal(t, n[0].SenderJoin.UserID, userID)
}

func TestO2OInferred_PKisFK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	workitemID := db.XoTestsWorkItemID(1)

	dwi, err := db.XoTestsDemoWorkItemByWorkItemID(ctx, testPool, workitemID, db.WithXoTestsDemoWorkItemJoin(db.XoTestsDemoWorkItemJoins{WorkItem: true}))
	require.NoError(t, err)
	assert.Equal(t, dwi.WorkItemID, workitemID)
	assert.Equal(t, dwi.WorkItemJoin.WorkItemID, workitemID)

	wi, err := db.XoTestsWorkItemByWorkItemID(ctx, testPool, workitemID, db.WithXoTestsWorkItemJoin(db.XoTestsWorkItemJoins{DemoWorkItem: true}))
	require.NoError(t, err)
	assert.Equal(t, wi.DemoWorkItemJoin.WorkItemID, workitemID)
	assert.Equal(t, wi.WorkItemID, workitemID)
}

func TestO2OInferred_VerticallyPartitioned(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := db.NewXoTestsUserID(uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))

	u, err := db.XoTestsUserByUserID(ctx, testPool, userID, db.WithXoTestsUserJoin(db.XoTestsUserJoins{UserAPIKey: true}))
	require.NoError(t, err)
	assert.Equal(t, u.APIKeyJoin.UserID, userID)

	uak, err := db.XoTestsUserAPIKeyByUserID(ctx, testPool, userID, db.WithXoTestsUserAPIKeyJoin(db.XoTestsUserAPIKeyJoins{User: true}))
	require.NoError(t, err)
	assert.Equal(t, uak.UserJoin.UserID, userID)
	assert.Equal(t, uak.UserID, userID)
}

func TestCustomFilters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	uu, err := db.XoTestsUserPaginatedByCreatedAt(ctx, testPool, time.Now().Add(-999*time.Hour), models.DirectionAsc,
		db.WithXoTestsUserJoin(db.XoTestsUserJoins{UserAPIKey: true, BooksAuthor: true}),
		db.WithXoTestsUserFilters(map[string][]any{
			"xo_tests.users.name = any ($i)":       {[]string{"Jane Smith"}}, // unique
			"NOT (xo_tests.users.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
			`(xo_tests.users.created_at > $i OR
		true = $i)`: {time.Now().Add(-24 * time.Hour), true},
		}))
	require.NoError(t, err)
	require.Len(t, uu, 1)
	assert.NotNil(t, uu[0].AuthorBooksJoin)
	require.Len(t, *uu[0].AuthorBooksJoin, 2)
}

func TestCRUD_UniqueIndex(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()

	u1 := createUserWithRetry(t, &db.XoTestsUserCreateParams{Name: "test_user_1"})
	u2 := createUserWithRetry(t, &db.XoTestsUserCreateParams{Name: "test_user_2"})

	u1.Name = "test_user_1_update"
	u1, err = u1.Update(ctx, testPool)
	require.NoError(t, err)
	assert.Equal(t, "test_user_1_update", u1.Name)

	// test hard delete
	err = u1.Delete(ctx, testPool)
	require.NoError(t, err)

	_, err = db.XoTestsUserByName(ctx, testPool, u1.Name)
	fmt.Printf("err: %v\n", err)
	assert.ErrorContains(t, err, errNoRows)

	// test soft delete and restore
	err = u2.SoftDelete(ctx, testPool)
	require.NoError(t, err)
	assert.NotNil(t, u2.DeletedAt)

	_, err = db.XoTestsUserByName(ctx, testPool, u2.Name) // default deleted_at null
	assert.ErrorContains(t, err, errNoRows)

	deletedUser, err := db.XoTestsUserByName(ctx, testPool, u2.Name, db.WithDeletedXoTestsUserOnly())
	require.NoError(t, err)
	assert.Equal(t, u2.Name, deletedUser.Name)
	assert.NotNil(t, deletedUser.DeletedAt)

	restoredUser, err := deletedUser.Restore(ctx, testPool)
	require.NoError(t, err)
	assert.Nil(t, restoredUser.DeletedAt)
	assert.Equal(t, u2.Name, restoredUser.Name)

	// TODO test same things with nonunique index too
}
