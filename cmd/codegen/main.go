package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/codegen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// nolint: gochecknoglobals
var (
	env, spec, opIDAuthPath string
	stderr                  bytes.Buffer

	validateSpecCmd = flag.NewFlagSet("validate-spec", flag.ExitOnError)
	preCmd          = flag.NewFlagSet("pre", flag.ExitOnError)

	subcommands = map[string]*flag.FlagSet{
		validateSpecCmd.Name(): validateSpecCmd,
		preCmd.Name():          preCmd,
	}
)

func main() {
	for _, fs := range subcommands {
		fs.StringVar(&opIDAuthPath, "op-id-auth", "", "JSON file with authorization information per operation ID")
		fs.StringVar(&env, "env", ".env", "Environment Variables filename")
		fs.StringVar(&spec, "spec", "openapi.yaml", "OpenAPI specification")
	}

	cmd, ok := subcommands[os.Args[1]]
	if !ok {
		validateSpecCmd.Usage()
		preCmd.Usage()

		return
	}

	cmd.Parse(os.Args[2:])

	cg := codegen.New(&stderr, spec, opIDAuthPath, "internal/rest")

	switch os.Args[1] {
	case "validate-spec":
		if err := cg.ValidateProjectSpec(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}
		os.Exit(0)
	case "pre":
		if err := envvar.Load(env); err != nil {
			log.Fatalf("envvar.Load: %s\n", err)
		}

		if opIDAuthPath == "" {
			log.Fatal("op-id-auth flag is required")
		}

		if err := cg.Generate(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}

		if err := cg.EnsureCorrectMethodsPerTag(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}
	}
}
