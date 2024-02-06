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
	conn   *pgxpool.Conn
	pool   *pgxpool.Pool
	LockID int

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

	row := al.conn.QueryRow(ctx, `SELECT pg_try_advisory_lock($1)`, al.LockID)
	if err := row.Scan(&lockSuccess); err != nil {
		return false, fmt.Errorf("lock query: %w", err)
	}

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

// WaitForRelease waits for the advisory lock to be released by another session.
// Returns an error if the wait times out.
func (al *AdvisoryLock) WaitForRelease(retryCount int, d time.Duration) error {
	if err := al.ensureConnAcquired(); err != nil {
		return fmt.Errorf("conn: %w", err)
	}

	for i := 0; i < retryCount; i++ {
		if !al.IsLocked() {
			return nil
		}

		time.Sleep(d)
	}

	return fmt.Errorf("timeout waiting for lock release with objid %d", al.LockID)
}

func (al *AdvisoryLock) IsLocked() bool {
	var lockExists bool

	row := al.pool.QueryRow(context.Background(), checkLockQuery, al.LockID)
	if err := row.Scan(&lockExists); err != nil {
		fmt.Printf("lock check error: %v\n", err)

		return true
	}

	return lockExists
}

// Release releases a single advisory lock.
// Returns whether unlocking was successful or not (doesn't own lock, failed...).
// Note that stacked lock requests require the same number of Release calls.
func (al *AdvisoryLock) Release() bool {
	ctx := context.Background()
	if al.conn == nil {
		// ReleaseConn was called beforehand and conn back in the pool,
		// so there's no way to release its locks, if any
		// TODO: test if pgx really doesn't clean them up, by acquiring then calling ReleaseConn
		// and then attempt TryLock with a new instance
		return false
	}

	if _, err := al.conn.Exec(ctx, `SELECT pg_advisory_unlock($1)`, al.LockID); err != nil {
		return false
	}

	var unlockSuccess bool
	row := al.conn.QueryRow(ctx, checkLockQuery, al.LockID)
	if err := row.Scan(&unlockSuccess); err != nil {
		return false
	}

	return unlockSuccess
}

// ReleaseConn releases the connection to the pool.
// It will not guarantee lock release.
func (al *AdvisoryLock) ReleaseConn() {
	if al.conn == nil {
		return
	}
	al.conn.Release()
	al.conn = nil
}
