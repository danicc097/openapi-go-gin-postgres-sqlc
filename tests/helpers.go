package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// Returns the directory of the file this function lives in.
func GetFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(b))
	return dir
}

// Setup runs necessary pre-testing commands. See project bin for details.
// FIXME can only run once at any time, and for any test (package,file,or test function). All packages will call this fn in their own TestMain
// since go test only tests one package at a time, there is no such thing as "before/after all the tests in subdirectories/subpackages"
// so this has to be done externally (e.g. flock might do the job to prevent the same bash function having more than one running instance).
func Setup() {
	os.Setenv("POSTGRES_DB", "postgres_test")
	os.Setenv("IS_TESTING", "1")
	rootDir := path.Join(GetFileRuntimeDirectory(), "..")

	if err := envvar.Load(path.Join(rootDir, ".env.dev")); err != nil {
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

	cmd = exec.Command(
		"bin/project",
		"test.backend-setup",
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
