package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func RandomNotificationCreateParams(t *testing.T, receiverRank *int, sender db.UserID, receiver *db.UserID, notificationType db.NotificationType) *db.NotificationCreateParams {
	t.Helper()

	return &db.NotificationCreateParams{
		Title:            testutil.RandomNameIdentifier(3, " "),
		Body:             testutil.RandomString(6),
		Labels:           []string{testutil.RandomString(6)},
		Link:             pointers.New("https://" + testutil.RandomString(6)),
		ReceiverRank:     receiverRank,
		Sender:           sender,
		Receiver:         receiver,
		NotificationType: notificationType,
	}
}
