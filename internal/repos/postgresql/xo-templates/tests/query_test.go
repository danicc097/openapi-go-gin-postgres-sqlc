package tests

import (
	"context"
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/xo-templates/tests/got"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/**
 * TODO: test extensively:
 *
 * - pagination
 * limits
 * order bys
 *
 * test M2M when its 1 pk and 2 fks, and with extra info ()
 *
	also test join table name clash for O2O constraint too:
	name clash probably needs to be detected between constraints, check M2M-M2O and M2O-O2O
	at the same time
*/

func TestM2M_TwoFKsAndExtraColumns(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{BooksAuthor: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoinAuthor, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))
	assert.NoError(t, err)
	assert.Nil(t, u.BooksJoinAuthor)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{BooksAuthor: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoinAuthor, 2)
	for _, b := range *u.BooksJoinAuthor {
		if b.Book.BookID == 1 {
			assert.Equal(t, *b.Pseudonym, "not Jane Smith")
		}
	}
}

func TestM2M_SurrogatePK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{BookSurrs: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BookSurrsJoin, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"))
	assert.NoError(t, err)
	assert.Nil(t, u.BookSurrsJoin)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{BookSurrs: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BookSurrsJoin, 2)
	for _, b := range *u.BookSurrsJoin {
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
	assert.Len(t, *u.BooksJoinSeller, 0)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"))
	assert.NoError(t, err)
	assert.Nil(t, u.BooksJoinSeller)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("8c67f1f9-2be4-4b1a-a49b-b7a10a60c53a"), db.WithUserJoin(db.UserJoins{BooksSeller: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoinSeller, 1)
	assert.Equal(t, (*u.BooksJoinSeller)[0].BookID, 1)
}

func TestM2O(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userID := uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")

	u, err := db.UserByUserID(ctx, testPool, userID, db.WithUserJoin(db.UserJoins{NotificationsSender: true, NotificationsReceiver: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.NotificationsJoinReceiver, 1)
	assert.Len(t, *u.NotificationsJoinSender, 2)

	n, err := db.NotificationsBySender(ctx, testPool, userID, db.WithNotificationJoin(db.NotificationJoins{UserSender: true, UserReceiver: true}))
	assert.NoError(t, err)
	assert.Len(t, n, 2)
	assert.Equal(t, n[0].UserJoinSender.UserID, userID)
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
	assert.Equal(t, u.UserAPIKeyJoin.UserID, userID)

	uak, err := db.UserAPIKeyByUserID(ctx, testPool, userID, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	assert.NoError(t, err)
	assert.Equal(t, uak.UserJoin.UserID, userID)
	assert.Equal(t, uak.UserID, userID)
}
