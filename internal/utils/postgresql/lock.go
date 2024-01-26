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
	lockID  int
	hasLock bool

	mu sync.Mutex
}

// NewAdvisoryLock creates a new AdvisoryLock.
func NewAdvisoryLock(pool *pgxpool.Pool, lockID int) (*AdvisoryLock, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}

	return &AdvisoryLock{
		conn:   conn,
		lockID: lockID,
	}, nil
}

// TryLock tries to acquire the advisory lock.
// Returns whether the lock was acquired and any error.
func (al *AdvisoryLock) TryLock(ctx context.Context) (bool, error) {
	var lockSuccess bool
	if al.hasLock {
		// prevent multiple calls to pg_try_advisory_lock.
		// if it succeeded n times, we would have had to unlock it n times too to release it.
		return true, nil
	}

	row := al.conn.QueryRow(ctx, `SELECT pg_try_advisory_lock($1)`, al.lockID)
	if err := row.Scan(&lockSuccess); err != nil {
		return false, fmt.Errorf("lock query: %w", err)
	}
	al.hasLock = lockSuccess

	return lockSuccess, nil
}

// WaitForRelease waits for the advisory lock to be released by another process.
// Returns an error if the wait times out.
func (al *AdvisoryLock) WaitForRelease(ctx context.Context) error {
	for i := 0; i < 100; i++ {
		lockExists := true

		row := al.conn.QueryRow(ctx, checkLockQuery, al.lockID)
		if err := row.Scan(&lockExists); err != nil {
			return fmt.Errorf("query: %w", err)
		}

		if !lockExists {
			return nil
		}

		time.Sleep(200 * time.Millisecond)
	}

	return fmt.Errorf("timeout waiting for lock release with objid %d", al.lockID)
}

// Release releases the advisory lock and the acquired connection.
func (al *AdvisoryLock) Release(ctx context.Context) error {
	locked := true // assume was locked

	for i := 0; i < 100 && locked; i++ {
		// sometimes it won't unlock on the first call, neither here nor in psql
		if _, err := al.conn.Exec(ctx, `SELECT pg_advisory_unlock($1)`, al.lockID); err != nil {
			return fmt.Errorf("lock query: %w", err)
		}

		row := al.conn.QueryRow(ctx, checkLockQuery, al.lockID)
		if err := row.Scan(&locked); err != nil {
			return fmt.Errorf("query: %w", err)
		}

		time.Sleep(200 * time.Millisecond)
	}
	al.hasLock = false
	al.conn.Release()

	return nil
}
