package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/google/uuid"
)

func RandomNotificationCreateParams(t *testing.T, receiverRank *int16, sender uuid.UUID, receiver *uuid.UUID, notificationType db.NotificationType) repos.NotificationCreateParams {
	t.Helper()

	return repos.NotificationCreateParams{
		Title:            testutil.RandomNameIdentifier(3, " "),
		Body:             testutil.RandomString(6),
		Label:            testutil.RandomString(6),
		Link:             pointers.New("https://" + testutil.RandomString(6)),
		ReceiverRank:     receiverRank,
		Sender:           sender,
		Receiver:         receiver,
		NotificationType: notificationType,
	}
}
