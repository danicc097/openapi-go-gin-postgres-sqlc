package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type Notification struct {
	logger *zap.SugaredLogger
	nrepo  repos.Notification
}

// NewNotification returns a new Notification service.
func NewNotification(logger *zap.SugaredLogger, nrepo repos.Notification) *Notification {
	return &Notification{
		logger: logger,
		nrepo:  nrepo,
	}
}

// LatestUserNotifications gets a notification by ID.
func (n *Notification) LatestUserNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error) {
	defer newOTelSpan(ctx, "").End()

	notification, err := n.nrepo.LatestUserNotifications(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("nrepo.LatestUserNotifications: %w", err)
	}

	return notification, nil
}

// Create creates a new notification.
func (n *Notification) Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (*db.Notification, error) {
	defer newOTelSpan(ctx, "").End()

	notification, err := n.nrepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("nrepo.Create: %w", err)
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
