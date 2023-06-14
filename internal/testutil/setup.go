package testutil

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sync"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// Returns the directory of the file this function lives in.
func GetFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(b))
	return dir
}

var setupOnce sync.Once

// Setup runs necessary pre-testing commands for a package: env vars loading, sourcing...
func Setup() {
	setupOnce.Do(func() {
		rootDir := path.Join(GetFileRuntimeDirectory(), "../..")

		appEnv := envvar.GetEnv("APP_ENV", "dev")
		if err := envvar.Load(path.Join(rootDir, ".env."+appEnv)); err != nil {
			log.Fatalf("envvar.Load: %s\n", err)
		}

		os.Setenv("POSTGRES_DB", "postgres_test")
		os.Setenv("IS_TESTING", "1") // for external scripts

		// update config with testing env vars
		if err := internal.NewAppConfig(); err != nil {
			log.Fatalf("internal.NewAppConfig: %s\n", err)
		}

		cmd := exec.Command(
			"bash", "-c",
			"source .envrc",
		)
		cmd.Dir = rootDir
		if out, err := cmd.CombinedOutput(); err != nil {
			errAndExit(out, err)
		}
	})

	// actually not really needed since variables won't be shared between package tests,
	// only in tests within the same package.
	// for !setupDone.Load() {
	// 	time.Sleep(100 * time.Millisecond)
	// }
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
