package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdvisoryLock(t *testing.T) {
	t.Parallel()

	t.Run("Locking and releasing", func(t *testing.T) {
		t.Parallel()

		lock, err := postgresql.NewAdvisoryLock(pool, 12343212)
		require.NoError(t, err)

		acquired, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		time.Sleep(3000 * time.Millisecond)
		acquiredTwice, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.False(t, acquiredTwice, "Lock was acquired again unexpectedly")

		err = lock.Release(context.Background())
		require.NoError(t, err)

		acquiredAfterRelease, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.True(t, acquiredAfterRelease, "Failed to acquire lock after release")
	})

	t.Run("Wait for release", func(t *testing.T) {
		t.Parallel()

		lockID := 223432112

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)

		lock2, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)

		acquired, err := lock2.TryLock(context.Background())
		require.NoError(t, err)
		require.True(t, acquired, "Could not acquire lock for the first time")

		err = lock.WaitForRelease(context.Background())
		require.NoError(t, err)

		lockAcquiredAfterWait, err := lock.TryLock(context.Background())
		require.NoError(t, err)
		assert.True(t, lockAcquiredAfterWait, "Failed to acquire lock after waiting")
	})

	t.Run("Connection closed after release", func(t *testing.T) {
		t.Parallel()

		lockID := 2234312112

		lock, err := postgresql.NewAdvisoryLock(pool, lockID)
		require.NoError(t, err)
		require.NotNil(t, lock)

		err = lock.Release(context.Background())
		require.NoError(t, err)

		require.Panics(t, func() { lock.TryLock(context.Background()) })
	})
}
