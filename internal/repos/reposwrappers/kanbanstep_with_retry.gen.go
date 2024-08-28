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

// KanbanStepWithRetry implements _sourceRepos.KanbanStep interface instrumented with retries
type KanbanStepWithRetry struct {
	_sourceRepos.KanbanStep
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewKanbanStepWithRetry returns KanbanStepWithRetry
func NewKanbanStepWithRetry(base _sourceRepos.KanbanStep, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) KanbanStepWithRetry {
	return KanbanStepWithRetry{
		KanbanStep:     base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements _sourceRepos.KanbanStep
func (_d KanbanStepWithRetry) ByID(ctx context.Context, d models.DBTX, id models.KanbanStepID, opts ...models.KanbanStepSelectConfigOption) (kp1 *models.KanbanStep, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT KanbanStepWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	kp1, err = _d.KanbanStep.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT KanbanStepWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to KanbanStepWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		kp1, err = _d.KanbanStep.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT KanbanStepWithRetryByID")
	}
	return
}

// ByProject implements _sourceRepos.KanbanStep
func (_d KanbanStepWithRetry) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.KanbanStepSelectConfigOption) (ka1 []models.KanbanStep, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT KanbanStepWithRetryByProject")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ka1, err = _d.KanbanStep.ByProject(ctx, d, projectID, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT KanbanStepWithRetryByProject")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to KanbanStepWithRetryByProject"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ka1, err = _d.KanbanStep.ByProject(ctx, d, projectID, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT KanbanStepWithRetryByProject")
	}
	return
}
