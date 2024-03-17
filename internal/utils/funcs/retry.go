package funcs

import (
	"time"
)

// Retry runs the given function until maximum retries are reached or it succeeds
// and returns whether it succeeded or not.
func Retry(
	fn func() error,
	maxRetry int, startBackoff, maxBackoff time.Duration,
) bool {
	for attempt := 0; ; attempt++ {
		if err := fn(); err == nil {
			return true
		}

		if attempt == maxRetry-1 {
			return false
		}

		time.Sleep(startBackoff)
		if startBackoff < maxBackoff {
			startBackoff *= 2
		}
	}
}
