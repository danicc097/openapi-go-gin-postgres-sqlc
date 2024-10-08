// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// NotificationWithTimeout implements _sourceRepos.Notification interface instrumented with timeouts
type NotificationWithTimeout struct {
	_sourceRepos.Notification
	config NotificationWithTimeoutConfig
}

type NotificationWithTimeoutConfig struct {
	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	LatestNotificationsTimeout time.Duration

	PaginatedUserNotificationsTimeout time.Duration
}

// NewNotificationWithTimeout returns NotificationWithTimeout
func NewNotificationWithTimeout(base _sourceRepos.Notification, config NotificationWithTimeoutConfig) NotificationWithTimeout {
	return NotificationWithTimeout{
		Notification: base,
		config:       config,
	}
}

// Create implements _sourceRepos.Notification
func (_d NotificationWithTimeout) Create(ctx context.Context, d models.DBTX, params *models.NotificationCreateParams) (up1 *models.UserNotification, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.Notification.Create(ctx, d, params)
}

// Delete implements _sourceRepos.Notification
func (_d NotificationWithTimeout) Delete(ctx context.Context, d models.DBTX, id models.NotificationID) (np1 *models.Notification, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.Notification.Delete(ctx, d, id)
}

// LatestNotifications implements _sourceRepos.Notification
func (_d NotificationWithTimeout) LatestNotifications(ctx context.Context, d models.DBTX, params *models.GetUserNotificationsParams) (ga1 []models.GetUserNotificationsRow, err error) {
	var cancelFunc func()
	if _d.config.LatestNotificationsTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.LatestNotificationsTimeout)
		defer cancelFunc()
	}
	return _d.Notification.LatestNotifications(ctx, d, params)
}

// PaginatedUserNotifications implements _sourceRepos.Notification
func (_d NotificationWithTimeout) PaginatedUserNotifications(ctx context.Context, d models.DBTX, userID models.UserID, params models.GetPaginatedNotificationsParams) (ua1 []models.UserNotification, err error) {
	var cancelFunc func()
	if _d.config.PaginatedUserNotificationsTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.PaginatedUserNotificationsTimeout)
		defer cancelFunc()
	}
	return _d.Notification.PaginatedUserNotifications(ctx, d, userID, params)
}
