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

	t.Run("Locking and releasing in same instance", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)

		acquired, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		acquiredTwice, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.True(t, acquiredTwice)

		err = lock.Release(context.Background())
		require.NoError(t, err)
		assert.False(t, lock.HasLock)

		// should not need two Release calls since we ignore consecutive lock calls
		// and also spam call unlock if needed
		acquiredAfterRelease, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.True(t, acquiredAfterRelease, "Failed to acquire lock after release")
	})

	t.Run("Wait for release in concurrent calls", func(t *testing.T) {
		t.Parallel()

		// for test count>1 must be unique...
		lockID := testutil.RandomInt(124342232, 999945323)

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)

		lockOwner, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)

		acquired, err := lockOwner.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		go func() {
			time.Sleep(200 * time.Millisecond)
			err := lockOwner.Release(context.Background())
			require.NoError(t, err)
			require.False(t, lockOwner.HasLock)
		}()

		time.Sleep(50 * time.Millisecond)
		acquired, err = lock.TryLock(context.Background())
		require.False(t, acquired)
		require.NoError(t, err)

		require.False(t, lock.HasLock)
		err = lock.WaitForRelease(context.Background(), 100, 50*time.Millisecond)
		require.NoError(t, err)

		lockAcquiredAfterWait, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, lockAcquiredAfterWait)
	})
}
