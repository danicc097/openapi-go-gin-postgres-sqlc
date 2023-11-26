package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type Notification struct {
	logger   *zap.SugaredLogger
	repos    *repos.Repos
	authzsvc *Authorization
	usvc     *User
}

type NotificationCreateParamsBase struct {
	Body   string    `json:"body"   nullable:"false" required:"true"`
	Labels []string  `json:"labels" nullable:"false" required:"true"`
	Link   *string   `json:"link"`
	Sender db.UserID `json:"sender" nullable:"false" required:"true"`
	Title  string    `json:"title"  nullable:"false" required:"true"`
}

type PersonalNotificationCreateParams struct {
	NotificationCreateParamsBase
	Receiver db.UserID `json:"receiver"`
}

type GlobalNotificationCreateParams struct {
	NotificationCreateParamsBase
	ReceiverRole models.Role `json:"receiverRole"`
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

// LatestUserNotifications gets user notifications ordered by creation date.
func (n *Notification) LatestUserNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error) {
	defer newOTelSpan().Build(ctx).End()

	notification, err := n.repos.Notification.LatestUserNotifications(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.LatestUserNotifications: %w", err)
	}

	return notification, nil
}

// Create creates a new notification.
func (n *Notification) CreatePersonalNotification(ctx context.Context, d db.DBTX, params *PersonalNotificationCreateParams) (*db.Notification, error) {
	defer newOTelSpan().Build(ctx).End()

	notification, err := n.repos.Notification.Create(ctx, d, &db.NotificationCreateParams{
		Body:             params.Body,
		Labels:           params.Labels,
		Link:             params.Link,
		Receiver:         &params.Receiver,
		NotificationType: db.NotificationTypePersonal,
		Title:            params.Title,
		Sender:           params.Sender,
	})
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.Create: %w", err)
	}

	return notification, nil
}

func (n *Notification) CreateGlobalNotification(ctx context.Context, d db.DBTX, params *GlobalNotificationCreateParams) (*db.Notification, error) {
	defer newOTelSpan().Build(ctx).End()

	role := n.authzsvc.RoleByName(params.ReceiverRole)
	superAdmin, err := n.usvc.ByEmail(ctx, d, internal.Config.SuperAdmin.DefaultEmail)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodePrivate, "could not get admin user: %s", err)
	}

	notification, err := n.repos.Notification.Create(ctx, d, &db.NotificationCreateParams{
		Body:             params.Body,
		Labels:           params.Labels,
		Link:             params.Link,
		ReceiverRank:     &role.Rank,
		NotificationType: db.NotificationTypeGlobal,
		Title:            params.Title,
		Sender:           superAdmin.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("repos.Notification.Create: %w", err)
	}

	return notification, nil
}

// TODO: latest joins on sender and returns id, username.
// paginated notification should have orderby id (somehow missing when its pk) and each index query should also have
// ...By..._PaginatedBy<cursor_cols> queries (will always be unique fns).
// this way we could paginate e.g. on notification id DESC where user_id = <current_user_id>
// using UserNotificationsByUserID_PaginatedByNotificationID
// instead of repeatedly calling FKUser_Sender, or generating an adhoc query,
// we could have a cache of users client-side, so if sender's user_id is unknown we just GET /users/ with a bunch of ids, return all at once and be done with it.
// or better yet server side cache for X hours of `users map[uuid]User` for both sender and receiver joins, since contextual info will barely change and ids certainly won't change.
// MarkAsRead,
