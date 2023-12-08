package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/codegen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/codegen/symbols"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// nolint: gochecknoglobals
var (
	env, spec, opIDAuthPath, structNamesList string
	stderr                                   bytes.Buffer

	implementServerCmd = flag.NewFlagSet("implement-server", flag.ExitOnError)
	validateSpecCmd    = flag.NewFlagSet("validate-spec", flag.ExitOnError)
	preCmd             = flag.NewFlagSet("pre", flag.ExitOnError)
	genSchemaCmd       = flag.NewFlagSet("gen-schema", flag.ExitOnError)

	subcommands = map[string]*flag.FlagSet{
		validateSpecCmd.Name():    validateSpecCmd,
		preCmd.Name():             preCmd,
		genSchemaCmd.Name():       genSchemaCmd,
		implementServerCmd.Name(): implementServerCmd,
	}
)

func main() {
	for _, fs := range []*flag.FlagSet{validateSpecCmd, preCmd} {
		fs.StringVar(&opIDAuthPath, "op-id-auth", "", "JSON file with authorization information per operation ID")
		fs.StringVar(&env, "env", ".env", "Environment Variables filename")
		fs.StringVar(&spec, "spec", "openapi.yaml", "OpenAPI specification")
	}

	genSchemaCmd.StringVar(&structNamesList, "struct-names", "", "comma-delimited db package struct names to generate an OpenAPI schema for")

	cmd, ok := subcommands[os.Args[1]]
	if !ok {
		for _, fs := range subcommands {
			fs.Usage()
		}

		return
	}

	cmd.Parse(os.Args[2:])

	codeGen := codegen.New(&stderr, spec, opIDAuthPath, "internal/rest")

	switch os.Args[1] {
	case "gen-schema":
		i := interp.New(interp.Options{
			GoPath: ".",
			// SourcecodeFilesystem: sourceFS{path: "src/main/vendor/" + target, fs: src},
			Unrestricted: true,
		})
		if err := i.Use(stdlib.Symbols); err != nil {
			panic(err)
		}
		if err := i.Use(symbols.Symbols); err != nil {
			panic(err)
		}

		if _, err := i.EvalPath("internal/codegen/structs.gen.go"); err != nil {
			panic(err)
		}
		if _, err := i.EvalPath("internal/codegen/gen_schema.go"); err != nil {
			panic(err)
		}
		structNames := strings.Split(structNamesList, ",")
		for i := range structNames {
			structNames[i] = strings.TrimSpace(structNames[i])
		}

		v, err := i.Eval("codegen.GenerateSpecSchemas")
		if err != nil {
			panic(err)
		}
		genSpecSchemas := v.Interface().(func([]string) string)
		r := genSpecSchemas(structNames)
		println(r)
		os.Exit(0)
	case "implement-server":
		if err := codeGen.ImplementServer(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}
		os.Exit(0)
	case "validate-spec":
		if err := codeGen.ValidateProjectSpec(); err != nil {
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

		if err := codeGen.Generate(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}

		if err := codeGen.EnsureCorrectMethodsPerTag(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, stderr.String())
			os.Exit(1)
		}
	}
}
