package postgresql_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
)

func TestAdvisoryLock(t *testing.T) {
	t.Parallel()

	t.Run("Locking twice in same session", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(testPool, lockID)
		defer lock.ReleaseConn()
		require.NoError(t, err)

		lock2, err := postgresql.NewAdvisoryLock(testPool, lockID)
		defer lock2.ReleaseConn()
		require.NoError(t, err)

		acquired, err := lock.TryLock(t.Context())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		acquiredTwice, err := lock.TryLock(t.Context())
		require.NoError(t, err)
		assert.True(t, acquiredTwice)

		locked := lock.Release()
		require.True(t, locked)

		acquired, err = lock2.TryLock(t.Context())
		require.NoError(t, err)
		require.False(t, acquired, "Should have failed to acquire lock after only one release")

		locked = lock.Release()
		require.False(t, locked)

		acquired, err = lock2.TryLock(t.Context())
		require.NoError(t, err)
		require.True(t, acquired, "Should have acquired lock after second release")

		lock.Release()  // TODO: check if upon return to pgxpool its gone
		lock2.Release() // TODO: check if upon return to pgxpool its gone
	})

	/**
	 *
	 * TODO: test ReleaseConn
	 */

	t.Run("Wait for release in concurrent calls", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(testPool, lockID)
		defer lock.ReleaseConn()
		require.NoError(t, err)

		lockOwner, err := postgresql.NewAdvisoryLock(testPool, lockID)
		defer lockOwner.ReleaseConn()
		require.NoError(t, err)

		acquired, err := lockOwner.TryLock(t.Context())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		acquired, err = lock.TryLock(t.Context())
		require.NoError(t, err)
		require.False(t, acquired)

		locked := lockOwner.Release()
		require.NoError(t, err)
		require.False(t, locked)

		err = lock.WaitForRelease(100, 50*time.Millisecond)
		require.NoError(t, err)

		lockAcquiredAfterWait, err := lock.TryLock(t.Context())
		require.NoError(t, err)
		require.True(t, lockAcquiredAfterWait)

		lock.Release()      // TODO: check if upon return to pgxpool its gone
		lockOwner.Release() // TODO: check if upon return to pgxpool its gone
	})
}
