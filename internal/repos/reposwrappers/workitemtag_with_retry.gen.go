// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry-repo.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// WorkItemTagWithRetry implements repos.WorkItemTag interface instrumented with retries
type WorkItemTagWithRetry struct {
	repos.WorkItemTag
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewWorkItemTagWithRetry returns WorkItemTagWithRetry
func NewWorkItemTagWithRetry(base repos.WorkItemTag, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) WorkItemTagWithRetry {
	return WorkItemTagWithRetry{
		WorkItemTag:    base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) ByID(ctx context.Context, d db.DBTX, id db.WorkItemTagID, opts ...db.WorkItemTagSelectConfigOption) (wp1 *db.WorkItemTag, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTagWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	wp1, err = _d.WorkItemTag.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTagWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemTag.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryByID")
	}
	return
}

// ByName implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.WorkItemTagSelectConfigOption) (wp1 *db.WorkItemTag, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTagWithRetryByName")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	wp1, err = _d.WorkItemTag.ByName(ctx, d, name, projectID, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryByName")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTagWithRetryByName"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemTag.ByName(ctx, d, name, projectID, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryByName")
	}
	return
}

// Create implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Create(ctx context.Context, d db.DBTX, params *db.WorkItemTagCreateParams) (wp1 *db.WorkItemTag, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTagWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	wp1, err = _d.WorkItemTag.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryCreate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTagWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemTag.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryCreate")
	}
	return
}

// Delete implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Delete(ctx context.Context, d db.DBTX, id db.WorkItemTagID) (wp1 *db.WorkItemTag, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTagWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	wp1, err = _d.WorkItemTag.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryDelete")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTagWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemTag.Delete(ctx, d, id)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryDelete")
	}
	return
}

// Update implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Update(ctx context.Context, d db.DBTX, id db.WorkItemTagID, params *db.WorkItemTagUpdateParams) (wp1 *db.WorkItemTag, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTagWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	wp1, err = _d.WorkItemTag.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryUpdate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTagWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemTag.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemTagWithRetryUpdate")
	}
	return
}
