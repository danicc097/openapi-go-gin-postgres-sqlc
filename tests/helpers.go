package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// Returns the directory of the file this function lives in.
func GetFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(b))
	return dir
}

// Setup runs necessary pre-testing commands. See project bin for details.
func Setup() {
	os.Setenv("POSTGRES_DB", "postgres_test")
	os.Setenv("IS_TESTING", "1")
	rootDir := path.Join(GetFileRuntimeDirectory(), "..")

	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok {
		appEnv = "dev"
	}
	if err := envvar.Load(path.Join(rootDir, ".env."+appEnv)); err != nil {
		fmt.Fprintf(os.Stderr, "envvar.Load: %s\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = rootDir
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}

// GetStderr returns the contents of stderr.txt in dir.
func GetStderr(t *testing.T, dir string) string {
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
