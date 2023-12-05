package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"go.uber.org/zap"
)

type Notification struct {
	logger   *zap.SugaredLogger
	repos    *repos.Repos
	authzsvc *Authorization
	usvc     *User
}

type NotificationCreateParams struct {
	db.NotificationCreateParams
	ReceiverRole *models.Role `json:"receiverRole"`
}

// NewNotification returns a new Notification service.
func NewNotification(logger *zap.SugaredLogger, repos *repos.Repos) *Notification {
	usvc := NewUser(logger, repos)
	authzsvc, err := NewAuthorization(logger)
	if err != nil {
		panic(fmt.Sprintf("NewAuthorization: %v", err))
	}

	return &Notification{
		logger:   logger,
		repos:    repos,
		authzsvc: authzsvc,
		usvc:     usvc,
	}
}

// LatestNotifications gets user notifications ordered by creation date.
func (n *Notification) LatestNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error) {
	defer newOTelSpan().Build(ctx).End()

	notification, err := n.repos.Notification.LatestNotifications(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.LatestNotifications: %w", err)
	}

	return notification, nil
}

// PaginatedNotifications gets user notifications by cursor.
func (n *Notification) PaginatedNotifications(ctx context.Context, d db.DBTX, userID db.UserID, params models.GetPaginatedNotificationsParams) ([]db.UserNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	notifications, err := n.repos.Notification.PaginatedNotifications(ctx, d, userID, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.PaginatedNotifications: %w", err)
	}

	return notifications, nil
}

// Create creates a new notification. In case of global notifications returns a single one from the fan out.
func (n *Notification) CreateNotification(ctx context.Context, d db.DBTX, params *NotificationCreateParams) (*db.UserNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	switch params.NotificationType {
	case db.NotificationTypeGlobal:
		if params.ReceiverRole == nil {
			return nil, internal.NewErrorWithLocf(models.ErrorCodeInvalidArgument, []string{"receiverRole"}, "minimum receiver role is not set")
		}
		params.NotificationCreateParams.ReceiverRank = pointers.New(n.authzsvc.RoleByName(*params.ReceiverRole).Rank)
		// let sender be whatever was set, no need to be superadmin
	case db.NotificationTypePersonal:
		if params.Receiver == nil {
			return nil, internal.NewErrorWithLocf(models.ErrorCodeInvalidArgument, []string{"receiver"}, "receiver is not set")
		}
	}

	notification, err := n.repos.Notification.Create(ctx, d, &params.NotificationCreateParams)
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.Create: %w", err)
	}

	return notification, nil
}
