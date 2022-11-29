package testutil

import (
	"fmt"
	"log"
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

// Setup runs necessary pre-testing commands for a package: env vars loading, sourcing...
func Setup() {
	os.Setenv("POSTGRES_DB", "postgres_test")
	os.Setenv("IS_TESTING", "1")
	rootDir := path.Join(GetFileRuntimeDirectory(), "../..")

	appEnv := envvar.GetEnv("APP_ENV", "dev")
	if err := envvar.Load(path.Join(rootDir, ".env."+appEnv)); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
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
