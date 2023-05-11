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
 *
 * - PK is FK tests like demoworkitems->workitemid
 * - vert partitioned columns -> user_api_keys inferred O2O
 */

func TestM2M(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{Books: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoin, 1)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{Books: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoin, 2)
}

func TestM2O(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	userID := uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")

	n, err := db.NotificationsBySender(ctx, testPool, userID, db.WithNotificationJoin(db.NotificationJoins{UserSender: true}))
	assert.NoError(t, err)
	assert.Len(t, n, 2)
	assert.Equal(t, n[0].UserJoinSender.UserID, userID)

	u, err := db.UserByUserID(ctx, testPool, userID, db.WithUserJoin(db.UserJoins{NotificationsSender: true}))
	assert.NoError(t, err)
	assert.Len(t, *u.NotificationsJoinSender, 2)
	assert.Equal(t, n[0].UserJoinSender.UserID, userID)
}
