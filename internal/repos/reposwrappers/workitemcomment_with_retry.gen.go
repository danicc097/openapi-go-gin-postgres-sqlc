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

// WorkItemCommentWithRetry implements repos.WorkItemComment interface instrumented with retries
type WorkItemCommentWithRetry struct {
	repos.WorkItemComment
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewWorkItemCommentWithRetry returns WorkItemCommentWithRetry
func NewWorkItemCommentWithRetry(base repos.WorkItemComment, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) WorkItemCommentWithRetry {
	return WorkItemCommentWithRetry{
		WorkItemComment: base,
		_retryCount:     retryCount,
		_retryInterval:  retryInterval,
		logger:          logger,
	}
}

// ByID implements repos.WorkItemComment
func (_d WorkItemCommentWithRetry) ByID(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, opts ...db.WorkItemCommentSelectConfigOption) (wp1 *db.WorkItemComment, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemCommentWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("IdleConns: %v\n", p.Stat().IdleConns())
		_d.logger.Infof("AcquiredConns: %v\n", p.Stat().AcquiredConns())
		_d.logger.Infof("ConstructingConns: %v\n", p.Stat().ConstructingConns())
		_d.logger.Infof("TotalConns: %v\n", p.Stat().TotalConns())
	}
	wp1, err = _d.WorkItemComment.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemCommentWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemComment.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryByID")
	}
	return
}

// Create implements repos.WorkItemComment
func (_d WorkItemCommentWithRetry) Create(ctx context.Context, d db.DBTX, params *db.WorkItemCommentCreateParams) (wp1 *db.WorkItemComment, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemCommentWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("IdleConns: %v\n", p.Stat().IdleConns())
		_d.logger.Infof("AcquiredConns: %v\n", p.Stat().AcquiredConns())
		_d.logger.Infof("ConstructingConns: %v\n", p.Stat().ConstructingConns())
		_d.logger.Infof("TotalConns: %v\n", p.Stat().TotalConns())
	}
	wp1, err = _d.WorkItemComment.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryCreate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemCommentWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemComment.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryCreate")
	}
	return
}

// Delete implements repos.WorkItemComment
func (_d WorkItemCommentWithRetry) Delete(ctx context.Context, d db.DBTX, id db.WorkItemCommentID) (wp1 *db.WorkItemComment, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemCommentWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("IdleConns: %v\n", p.Stat().IdleConns())
		_d.logger.Infof("AcquiredConns: %v\n", p.Stat().AcquiredConns())
		_d.logger.Infof("ConstructingConns: %v\n", p.Stat().ConstructingConns())
		_d.logger.Infof("TotalConns: %v\n", p.Stat().TotalConns())
	}
	wp1, err = _d.WorkItemComment.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryDelete")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemCommentWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemComment.Delete(ctx, d, id)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryDelete")
	}
	return
}

// Update implements repos.WorkItemComment
func (_d WorkItemCommentWithRetry) Update(ctx context.Context, d db.DBTX, id db.WorkItemCommentID, params *db.WorkItemCommentUpdateParams) (wp1 *db.WorkItemComment, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemCommentWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("IdleConns: %v\n", p.Stat().IdleConns())
		_d.logger.Infof("AcquiredConns: %v\n", p.Stat().AcquiredConns())
		_d.logger.Infof("ConstructingConns: %v\n", p.Stat().ConstructingConns())
		_d.logger.Infof("TotalConns: %v\n", p.Stat().TotalConns())
	}
	wp1, err = _d.WorkItemComment.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryUpdate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemCommentWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemComment.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT WorkItemCommentWithRetryUpdate")
	}
	return
}
