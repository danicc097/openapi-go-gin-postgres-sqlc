package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdvisoryLock(t *testing.T) {
	t.Parallel()

	t.Run("Locking twice in same session", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		defer lock.ReleaseConn()
		require.NoError(t, err)

		lock2, err := postgresql.NewAdvisoryLock(pool, lockID)
		defer lock2.ReleaseConn()
		require.NoError(t, err)

		acquired, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		acquiredTwice, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.True(t, acquiredTwice)

		locked := lock.Release(context.Background())
		require.True(t, locked)

		acquired, err = lock2.TryLock(context.Background())
		require.NoError(t, err)
		require.False(t, acquired, "Should have failed to acquire lock after only one release")

		locked = lock.Release(context.Background())
		require.False(t, locked)

		acquired, err = lock2.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Should have acquired lock after second release")

		lock.ReleaseConn()
		lock2.ReleaseConn()
	})

	/**
	 *
	 * TODO: test ReleaseConn
	 */

	t.Run("Wait for release in concurrent calls", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		defer lock.ReleaseConn()
		require.NoError(t, err)

		lockOwner, err := postgresql.NewAdvisoryLock(pool, lockID)
		defer lockOwner.ReleaseConn()
		require.NoError(t, err)

		acquired, err := lockOwner.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		acquired, err = lock.TryLock(context.Background())
		require.NoError(t, err)
		require.False(t, acquired)

		locked := lockOwner.Release(context.Background())
		require.NoError(t, err)
		require.False(t, locked)

		err = lock.WaitForRelease(100, 50*time.Millisecond)
		require.NoError(t, err)

		lockAcquiredAfterWait, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, lockAcquiredAfterWait)

		lock.ReleaseConn()
		lockOwner.ReleaseConn()
	})
}
