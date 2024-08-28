// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry-repo.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"fmt"
	"time"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// DemoWorkItemWithRetry implements _sourceRepos.DemoWorkItem interface instrumented with retries
type DemoWorkItemWithRetry struct {
	_sourceRepos.DemoWorkItem
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewDemoWorkItemWithRetry returns DemoWorkItemWithRetry
func NewDemoWorkItemWithRetry(base _sourceRepos.DemoWorkItem, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) DemoWorkItemWithRetry {
	return DemoWorkItemWithRetry{
		DemoWorkItem:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements _sourceRepos.DemoWorkItem
func (_d DemoWorkItemWithRetry) ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (wp1 *models.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoWorkItemWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoWorkItem.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoWorkItemWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoWorkItem.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryByID")
	}
	return
}

// Create implements _sourceRepos.DemoWorkItem
func (_d DemoWorkItemWithRetry) Create(ctx context.Context, d models.DBTX, params _sourceRepos.DemoWorkItemCreateParams) (wp1 *models.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoWorkItemWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoWorkItem.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryCreate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoWorkItemWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoWorkItem.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryCreate")
	}
	return
}

// Paginated implements _sourceRepos.DemoWorkItem
func (_d DemoWorkItemWithRetry) Paginated(ctx context.Context, d models.DBTX, cursor models.WorkItemID, opts ...models.CacheDemoWorkItemSelectConfigOption) (ca1 []models.CacheDemoWorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoWorkItemWithRetryPaginated")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ca1, err = _d.DemoWorkItem.Paginated(ctx, d, cursor, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryPaginated")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoWorkItemWithRetryPaginated"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ca1, err = _d.DemoWorkItem.Paginated(ctx, d, cursor, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryPaginated")
	}
	return
}

// Update implements _sourceRepos.DemoWorkItem
func (_d DemoWorkItemWithRetry) Update(ctx context.Context, d models.DBTX, id models.WorkItemID, params _sourceRepos.DemoWorkItemUpdateParams) (wp1 *models.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT DemoWorkItemWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.DemoWorkItem.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryUpdate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to DemoWorkItemWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.DemoWorkItem.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT DemoWorkItemWithRetryUpdate")
	}
	return
}
