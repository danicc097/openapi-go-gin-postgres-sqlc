package e2e_test

import (
	"os"
	"path/filepath"
	"testing"
)

// getStderr returns the contents of stderr.txt in dir.
func getStderr(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "stderr.txt")

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		blob, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		return string(blob)
	}

	return ""
}
