// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// NotificationWithRetry implements repos.Notification interface instrumented with retries
type NotificationWithRetry struct {
	repos.Notification
	_retryCount    int
	_retryInterval time.Duration
}

// NewNotificationWithRetry returns NotificationWithRetry
func NewNotificationWithRetry(base repos.Notification, retryCount int, retryInterval time.Duration) NotificationWithRetry {
	return NotificationWithRetry{
		Notification:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// Create implements repos.Notification
func (_d NotificationWithRetry) Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (np1 *db.Notification, err error) {
	np1, err = _d.Notification.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		np1, err = _d.Notification.Create(ctx, d, params)
	}
	return
}

// Delete implements repos.Notification
func (_d NotificationWithRetry) Delete(ctx context.Context, d db.DBTX, id db.NotificationID) (np1 *db.Notification, err error) {
	np1, err = _d.Notification.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		np1, err = _d.Notification.Delete(ctx, d, id)
	}
	return
}

// LatestNotifications implements repos.Notification
func (_d NotificationWithRetry) LatestNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) (ga1 []db.GetUserNotificationsRow, err error) {
	ga1, err = _d.Notification.LatestNotifications(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		ga1, err = _d.Notification.LatestNotifications(ctx, d, params)
	}
	return
}
