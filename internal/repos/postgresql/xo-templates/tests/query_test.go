package tests

import (
	"context"
	"testing"
	"time"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/xo-templates/tests/got"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/**
* TODO: test extensively:
*
* - order bys
* - index queries
* - join table name clash for O2O constraint too:
    name clash probably needs to be detected between constraints, check M2M-M2O and M2O-O2O
    at the same time
*/

func TestCursorPagination_Timestamp(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ee, err := db.PagElementPaginatedByCreatedAtDesc(ctx, testPool, time.Now().Add(-(24+1)*time.Hour), db.WithPagElementLimit(1), db.WithPagElementJoin(db.PagElementJoins{}))
	assert.NoError(t, err)
	assert.Len(t, ee, 1)
	assert.Equal(t, ee[0].Name, "element -2 days")

	ee, err = db.PagElementPaginatedByCreatedAtDesc(ctx, testPool, ee[0].CreatedAt, db.WithPagElementLimit(2))
	assert.NoError(t, err)
	assert.Len(t, ee, 2)
	assert.Equal(t, ee[0].Name, "element -3 days")
	assert.Equal(t, ee[1].Name, "element -4 days")
}

func Test_Filters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ee, err := db.PagElementPaginatedByCreatedAtDesc(ctx, testPool, time.Now().Add(-(24+1)*time.Hour), db.WithPagElementLimit(1), db.WithPagElementJoin(db.PagElementJoins{}))
	assert.NoError(t, err)
	assert.Len(t, ee, 1)
	assert.Equal(t, ee[0].Name, "element -2 days")

	ee, err = db.PagElementPaginatedByCreatedAtDesc(ctx, testPool, ee[0].CreatedAt, db.WithPagElementLimit(2))
	assert.NoError(t, err)
	assert.Len(t, ee, 2)
	assert.Equal(t, ee[0].Name, "element -3 days")
	assert.Equal(t, ee[1].Name, "element -4 days")
}

func TestM2M_TwoFKsAndExtraColumns(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{BooksAuthor: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.AuthorBooksJoin, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))
	assert.NoError(t, err)
	assert.Nil(t, u.AuthorBooksJoin)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{BooksAuthor: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.AuthorBooksJoin, 2)
	for _, b := range *u.AuthorBooksJoin {
		if b.Book.BookID == 1 {
			assert.Equal(t, "not Jane Smith", *b.Pseudonym)
		}
	}
}

func TestM2M_SurrogatePK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{BooksAuthorBooks: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.AuthorBooksJoinBASK, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))
	assert.NoError(t, err)
	assert.Nil(t, u.AuthorBooksJoinBASK)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{BooksAuthorBooks: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.AuthorBooksJoinBASK, 2)
	for _, b := range *u.AuthorBooksJoinBASK {
		if b.Book.BookID == 1 {
			assert.Equal(t, *b.Pseudonym, "not Jane Smith")
		}
	}
}

func TestM2M_TwoFKs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{BooksSeller: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.SellerBooksJoin, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"))
	assert.NoError(t, err)
	assert.Nil(t, u.SellerBooksJoin)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8c67f1f9-2be4-4b1a-a49b-b7a10a60c53a"), db.WithUserJoin(db.UserJoins{BooksSeller: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.SellerBooksJoin, 1)
	assert.Equal(t, (*u.SellerBooksJoin)[0].BookID, 1)
}

func TestM2O(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userID := uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")

	u, err := db.UserByUserID(ctx, testPool, userID, db.WithUserJoin(db.UserJoins{NotificationsSender: true, NotificationsReceiver: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.ReceiverNotificationsJoin, 1)
	assert.Len(t, *u.SenderNotificationsJoin, 2)

	n, err := db.NotificationsBySender(ctx, testPool, userID, db.WithNotificationJoin(db.NotificationJoins{UserSender: true, UserReceiver: true}))
	assert.NoError(t, err)
	assert.Len(t, n, 2)
	assert.Equal(t, n[0].SenderJoin.UserID, userID)
}

func TestO2OInferred_PKisFK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	workitemID := int64(1)

	dwi, err := db.DemoWorkItemByWorkItemID(ctx, testPool, workitemID, db.WithDemoWorkItemJoin(db.DemoWorkItemJoins{WorkItem: true}))
	assert.NoError(t, err)
	assert.Equal(t, dwi.WorkItemID, workitemID)
	assert.Equal(t, dwi.WorkItemJoin.WorkItemID, workitemID)

	wi, err := db.WorkItemByWorkItemID(ctx, testPool, workitemID, db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true}))
	assert.NoError(t, err)
	assert.Equal(t, wi.DemoWorkItemJoin.WorkItemID, workitemID)
	assert.Equal(t, wi.WorkItemID, workitemID)
}

func TestO2OInferred_VerticallyPartitioned(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")

	u, err := db.UserByUserID(ctx, testPool, userID, db.WithUserJoin(db.UserJoins{UserAPIKey: true}))
	assert.NoError(t, err)
	assert.Equal(t, u.APIKeyJoin.UserID, userID)

	uak, err := db.UserAPIKeyByUserID(ctx, testPool, userID, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	assert.NoError(t, err)
	assert.Equal(t, uak.UserJoin.UserID, userID)
	assert.Equal(t, uak.UserID, userID)
}

func TestCustomFilters(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	uu, err := db.UserPaginatedByCreatedAtAsc(ctx, testPool, time.Now().Add(-999*time.Hour),
		db.WithUserJoin(db.UserJoins{UserAPIKey: true, BooksAuthor: true}),
		db.WithUserFilters(map[string][]any{
			"xo_tests.users.name = any ($i)":       {[]string{"Jane Smith"}}, // unique
			"NOT (xo_tests.users.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
			`(xo_tests.users.created_at > $i OR
		true = $i)`: {time.Now().Add(-24 * time.Hour), true},
		}))
	assert.NoError(t, err)
	assert.Len(t, uu, 1)
	assert.NotNil(t, uu[0].AuthorBooksJoin)
	assert.Len(t, *uu[0].AuthorBooksJoin, 2)
}

func TestCRUD_UniqueIndex(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u1, err := db.CreateUser(ctx, testPool, &db.UserCreateParams{Name: "test_user_1"})
	assert.NoError(t, err)
	assert.Equal(t, "test_user_1", u1.Name)
	u2, err := db.CreateUser(ctx, testPool, &db.UserCreateParams{Name: "test_user_2"})
	assert.NoError(t, err)

	u1.Name = "test_user_1_update"
	u1, err = u1.Update(ctx, testPool)
	assert.NoError(t, err)
	assert.Equal(t, "test_user_1_update", u1.Name)

	// test hard delete
	err = u1.Delete(ctx, testPool)
	assert.NoError(t, err)

	_, err = db.UserByName(ctx, testPool, u1.Name)
	assert.ErrorContains(t, err, errNoRows)

	// test soft delete and restore
	err = u2.SoftDelete(ctx, testPool)
	assert.NoError(t, err)
	assert.NotNil(t, u2.DeletedAt)

	_, err = db.UserByName(ctx, testPool, u2.Name) // default deleted_at null
	assert.ErrorContains(t, err, errNoRows)

	deletedUser, err := db.UserByName(ctx, testPool, u2.Name, db.WithDeletedUserOnly())
	assert.NoError(t, err)
	assert.Equal(t, u2.Name, deletedUser.Name)
	assert.NotNil(t, deletedUser.DeletedAt)

	restoredUser, err := deletedUser.Restore(ctx, testPool)
	assert.NoError(t, err)
	assert.Nil(t, restoredUser.DeletedAt)
	assert.Equal(t, u2.Name, restoredUser.Name)

	// TODO test same things with nonunique index too
}
