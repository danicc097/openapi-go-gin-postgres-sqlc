// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry-repo.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// NotificationWithRetry implements repos.Notification interface instrumented with retries
type NotificationWithRetry struct {
	repos.Notification
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewNotificationWithRetry returns NotificationWithRetry
func NewNotificationWithRetry(base repos.Notification, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) NotificationWithRetry {
	return NotificationWithRetry{
		Notification:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// Create implements repos.Notification
func (_d NotificationWithRetry) Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (up1 *db.UserNotification, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT NotificationWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.Notification.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryCreate")
		}
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to NotificationWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.Notification.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryCreate")
	}
	return
}

// Delete implements repos.Notification
func (_d NotificationWithRetry) Delete(ctx context.Context, d db.DBTX, id db.NotificationID) (np1 *db.Notification, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT NotificationWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	np1, err = _d.Notification.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryDelete")
		}
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to NotificationWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		np1, err = _d.Notification.Delete(ctx, d, id)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryDelete")
	}
	return
}

// LatestNotifications implements repos.Notification
func (_d NotificationWithRetry) LatestNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) (ga1 []db.GetUserNotificationsRow, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT NotificationWithRetryLatestNotifications")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ga1, err = _d.Notification.LatestNotifications(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryLatestNotifications")
		}
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to NotificationWithRetryLatestNotifications"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ga1, err = _d.Notification.LatestNotifications(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryLatestNotifications")
	}
	return
}

// PaginatedNotifications implements repos.Notification
func (_d NotificationWithRetry) PaginatedNotifications(ctx context.Context, d db.DBTX, userID db.UserID, params models.GetPaginatedNotificationsParams) (ua1 []db.UserNotification, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT NotificationWithRetryPaginatedNotifications")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ua1, err = _d.Notification.PaginatedNotifications(ctx, d, userID, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryPaginatedNotifications")
		}
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to NotificationWithRetryPaginatedNotifications"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ua1, err = _d.Notification.PaginatedNotifications(ctx, d, userID, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT NotificationWithRetryPaginatedNotifications")
	}
	return
}
