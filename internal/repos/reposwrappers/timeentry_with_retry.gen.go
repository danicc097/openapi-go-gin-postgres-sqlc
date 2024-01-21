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

// TimeEntryWithRetry implements repos.TimeEntry interface instrumented with retries
type TimeEntryWithRetry struct {
	repos.TimeEntry
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewTimeEntryWithRetry returns TimeEntryWithRetry
func NewTimeEntryWithRetry(base repos.TimeEntry, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) TimeEntryWithRetry {
	return TimeEntryWithRetry{
		TimeEntry:      base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements repos.TimeEntry
func (_d TimeEntryWithRetry) ByID(ctx context.Context, d db.DBTX, id db.TimeEntryID, opts ...db.TimeEntrySelectConfigOption) (tp1 *db.TimeEntry, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT TimeEntryWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	tp1, err = _d.TimeEntry.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to TimeEntryWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		tp1, err = _d.TimeEntry.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryByID")
	}
	return
}

// Create implements repos.TimeEntry
func (_d TimeEntryWithRetry) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (tp1 *db.TimeEntry, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT TimeEntryWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	tp1, err = _d.TimeEntry.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryCreate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to TimeEntryWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		tp1, err = _d.TimeEntry.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryCreate")
	}
	return
}

// Delete implements repos.TimeEntry
func (_d TimeEntryWithRetry) Delete(ctx context.Context, d db.DBTX, id db.TimeEntryID) (tp1 *db.TimeEntry, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT TimeEntryWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	tp1, err = _d.TimeEntry.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryDelete")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to TimeEntryWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		tp1, err = _d.TimeEntry.Delete(ctx, d, id)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryDelete")
	}
	return
}

// Update implements repos.TimeEntry
func (_d TimeEntryWithRetry) Update(ctx context.Context, d db.DBTX, id db.TimeEntryID, params *db.TimeEntryUpdateParams) (tp1 *db.TimeEntry, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT TimeEntryWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	} else if p, ok := d.(*pgxpool.Pool); ok {
		_d.logger.Infof("p.Stat(): %v\n", p.Stat())
	}
	tp1, err = _d.TimeEntry.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryUpdate")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to TimeEntryWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		tp1, err = _d.TimeEntry.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT TimeEntryWithRetryUpdate")
	}
	return
}
