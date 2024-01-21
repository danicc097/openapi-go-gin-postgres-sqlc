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
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// ProjectWithRetry implements repos.Project interface instrumented with retries
type ProjectWithRetry struct {
	repos.Project
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewProjectWithRetry returns ProjectWithRetry
func NewProjectWithRetry(base repos.Project, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) ProjectWithRetry {
	return ProjectWithRetry{
		Project:        base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements repos.Project
func (_d ProjectWithRetry) ByID(ctx context.Context, d db.DBTX, id db.ProjectID, opts ...db.ProjectSelectConfigOption) (pp1 *db.Project, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT ProjectWithRetryByID")
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
	pp1, err = _d.Project.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryByID")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to ProjectWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		pp1, err = _d.Project.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryByID")
	}
	return
}

// ByName implements repos.Project
func (_d ProjectWithRetry) ByName(ctx context.Context, d db.DBTX, name models.Project, opts ...db.ProjectSelectConfigOption) (pp1 *db.Project, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT ProjectWithRetryByName")
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
	pp1, err = _d.Project.ByName(ctx, d, name, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryByName")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to ProjectWithRetryByName"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		pp1, err = _d.Project.ByName(ctx, d, name, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryByName")
	}
	return
}

// IsTeamInProject implements repos.Project
func (_d ProjectWithRetry) IsTeamInProject(ctx context.Context, d db.DBTX, arg db.IsTeamInProjectParams) (b1 bool, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT ProjectWithRetryIsTeamInProject")
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
	b1, err = _d.Project.IsTeamInProject(ctx, d, arg)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryIsTeamInProject")
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
			if _, err = tx.Exec(ctx, "ROLLBACK to ProjectWithRetryIsTeamInProject"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		b1, err = _d.Project.IsTeamInProject(ctx, d, arg)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT ProjectWithRetryIsTeamInProject")
	}
	return
}
