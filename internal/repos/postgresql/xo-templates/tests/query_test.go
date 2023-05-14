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
 * order bys
 *
 * test M2M when its 1 pk and 2 fks, and with extra info ()
 *
	also test join table name clash for O2O constraint too:
	name clash probably needs to be detected between constraints, check M2M-M2O and M2O-O2O
	at the same time

	TODO xo maybe should allow custom untyped filters for paginated queries.
	e.g. workitems paginated by created at while filtering TeamID is X
	filters []string and just append those, while warning for obvious sql injection risks
	this way we could have any filtering logic, `team_id in {team1, team2}`, etc.
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
	assert.Equal(t, u.UserAPIKeyJoin.UserID, userID)

	uak, err := db.UserAPIKeyByUserID(ctx, testPool, userID, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	assert.NoError(t, err)
	assert.Equal(t, uak.UserJoin.UserID, userID)
	assert.Equal(t, uak.UserID, userID)
}
