package servicetestutil

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

type CreateNotificationParams struct {
	Receiver     *db.UserID
	ReceiverRole *models.Role
}

// CreatePersonalNotification creates a new notification with the given configuration.
func (ff *FixtureFactory) CreatePersonalNotification(ctx context.Context, params CreateNotificationParams) (*db.Notification, error) {
	admin, err := ff.CreateUser(ctx, CreateUserParams{Role: models.RoleAdmin})
	if err != nil {
		return nil, fmt.Errorf("ff.CreateUser: %w", err)
	}
	n, err := ff.svc.Notification.CreateNotification(ctx, ff.db, &services.NotificationCreateParams{
		NotificationCreateParams: db.NotificationCreateParams{
			Body:             testutil.RandomString(10),
			Labels:           []string{"label " + string(testutil.RandomInt(1, 9999))},
			Link:             pointers.New("https://somelink"),
			Title:            testutil.RandomNameIdentifier(0, "-"),
			Sender:           admin.User.UserID,
			NotificationType: db.NotificationTypePersonal,
			Receiver:         params.Receiver,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("ff.svc.Notification.CreateNotification: %w", err)
	}

	return n, nil
}

// CreateGlobalNotification creates a new notification with the given configuration.
func (ff *FixtureFactory) CreateGlobalNotification(ctx context.Context, params CreateNotificationParams) (*db.Notification, error) {
	admin, err := ff.CreateUser(ctx, CreateUserParams{Role: models.RoleAdmin})
	if err != nil {
		return nil, fmt.Errorf("ff.CreateUser: %w", err)
	}
	n, err := ff.svc.Notification.CreateNotification(ctx, ff.db, &services.NotificationCreateParams{
		NotificationCreateParams: db.NotificationCreateParams{
			Body:             testutil.RandomString(10),
			Labels:           []string{"label " + string(testutil.RandomInt(1, 9999))},
			Link:             pointers.New("https://somelink"),
			Title:            testutil.RandomNameIdentifier(0, "-"),
			Sender:           admin.User.UserID,
			NotificationType: db.NotificationTypeGlobal,
		},
		ReceiverRole: params.ReceiverRole,
	})
	if err != nil {
		return nil, fmt.Errorf("ff.svc.Notification.CreateNotification: %w", err)
	}

	return n, nil
}
