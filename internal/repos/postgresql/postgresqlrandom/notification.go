package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func NotificationCreateParams(receiverRank *int, sender models.UserID, receiver *models.UserID, notificationType models.NotificationType) *models.NotificationCreateParams {
	return &models.NotificationCreateParams{
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
