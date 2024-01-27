package postgresql

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const checkLockQuery = `
SELECT EXISTS (
		SELECT 1
		FROM pg_locks
		JOIN pg_stat_activity USING (pid)
		WHERE locktype = 'advisory' AND objid = $1
) AS lock_acquired;
`

// AdvisoryLock represents an advisory lock.
// See https://www.postgresql.org/docs/current/explicit-locking.html#ADVISORY-LOCKS
type AdvisoryLock struct {
	conn    *pgxpool.Conn
	pool    *pgxpool.Pool
	LockID  int
	HasLock bool

	mu sync.Mutex
}

// NewAdvisoryLock creates a new AdvisoryLock.
func NewAdvisoryLock(pool *pgxpool.Pool, lockID int) (*AdvisoryLock, error) {
	return &AdvisoryLock{
		pool:   pool,
		LockID: lockID,
	}, nil
}

// TryLock tries to acquire the advisory lock.
// Returns whether the lock was acquired and any error.
func (al *AdvisoryLock) TryLock(ctx context.Context) (bool, error) {
	al.mu.Lock()
	defer al.mu.Unlock()

	if err := al.ensureConnAcquired(); err != nil {
		return false, fmt.Errorf("conn: %w", err)
	}

	var lockSuccess bool
	if al.HasLock {
		// prevent multiple calls to pg_try_advisory_lock.
		// if it succeeded n times, we would have had to unlock it n times too to release it.
		return true, nil
	}

	row := al.conn.QueryRow(ctx, `SELECT pg_try_advisory_lock($1)`, al.LockID)
	if err := row.Scan(&lockSuccess); err != nil {
		return false, fmt.Errorf("lock query: %w", err)
	}
	al.HasLock = lockSuccess

	return lockSuccess, nil
}

func (al *AdvisoryLock) ensureConnAcquired() error {
	if al.conn == nil {
		conn, err := al.pool.Acquire(context.Background())
		if err != nil {
			return fmt.Errorf("could not acquire connection: %w", err)
		}
		al.conn = conn
	}

	return nil
}

// WaitForRelease waits for the advisory lock to be released by another process.
// Returns an error if the wait times out.
func (al *AdvisoryLock) WaitForRelease(ctx context.Context, retryCount int, d time.Duration) error {
	if err := al.ensureConnAcquired(); err != nil {
		return fmt.Errorf("conn: %w", err)
	}

	for i := 0; i < retryCount; i++ {
		lockExists := true

		row := al.conn.QueryRow(ctx, checkLockQuery, al.LockID)
		if err := row.Scan(&lockExists); err != nil {
			return fmt.Errorf("query: %w", err)
		}

		if !lockExists {
			return nil
		}

		time.Sleep(d)
	}

	return fmt.Errorf("timeout waiting for lock release with objid %d", al.LockID)
}

// Release releases the advisory lock and the acquired connection.
func (al *AdvisoryLock) Release(ctx context.Context) error {
	if al.conn == nil {
		// AdvisoryLock.Release was called beforehand or was never locked
		return nil
	}

	locked := true // assume it was locked at least check once

	// it won't unlock on the first call if it is was locked multiple times by the same owner
	for i := 0; i < 10 && locked; i++ {
		if _, err := al.conn.Exec(ctx, `SELECT pg_advisory_unlock($1)`, al.LockID); err != nil {
			return fmt.Errorf("lock query: %w", err)
		}

		row := al.conn.QueryRow(ctx, checkLockQuery, al.LockID)
		if err := row.Scan(&locked); err != nil {
			return fmt.Errorf("query: %w", err)
		}

		time.Sleep(200 * time.Millisecond)
	}
	if locked {
		return fmt.Errorf("could not release lock with id %d", al.LockID)
	}
	al.HasLock = false
	// conn.Release just returns it to the pool, lock would still be there.
	al.conn.Release()
	al.conn = nil

	return nil
}
