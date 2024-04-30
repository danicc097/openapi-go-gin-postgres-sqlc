// template: ../../gowrap-templates/retry-repo.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// DemoTwoWorkItemWithRetry implements repos.DemoTwoWorkItem interface instrumented with retries
type DemoTwoWorkItemWithRetry struct {
	repos.DemoTwoWorkItem
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewDemoTwoWorkItemWithRetry returns DemoTwoWorkItemWithRetry
func NewDemoTwoWorkItemWithRetry(base repos.DemoTwoWorkItem, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) DemoTwoWorkItemWithRetry {
	return DemoTwoWorkItemWithRetry{
		DemoTwoWorkItem: base,
		_retryCount:     retryCount,
		_retryInterval:  retryInterval,
		logger:          logger,
	}
}

// ByID implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithRetry) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoTwoWorkItemWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoTwoWorkItem.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoTwoWorkItemWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoTwoWorkItem.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryByID")
	}
	return
}

// Create implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithRetry) Create(ctx context.Context, d db.DBTX, params repos.DemoTwoWorkItemCreateParams) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoTwoWorkItemWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoTwoWorkItem.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryCreate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoTwoWorkItemWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoTwoWorkItem.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryCreate")
	}
	return
}

// Update implements repos.DemoTwoWorkItem
func (_d DemoTwoWorkItemWithRetry) Update(ctx context.Context, d db.DBTX, id db.WorkItemID, params repos.DemoTwoWorkItemUpdateParams) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoTwoWorkItemWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoTwoWorkItem.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryUpdate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoTwoWorkItemWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoTwoWorkItem.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoTwoWorkItemWithRetryUpdate")
	}
	return
}
