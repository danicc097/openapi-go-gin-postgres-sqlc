package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/exp/maps"
)

var (
	privateOnly, publicOnly bool

	structsCmd    = flag.NewFlagSet("find-structs", flag.ExitOnError)
	interfacesCmd = flag.NewFlagSet("find-interfaces", flag.ExitOnError)

	subcommands = map[string]*flag.FlagSet{
		structsCmd.Name():    structsCmd,
		interfacesCmd.Name(): interfacesCmd,
	}
)

func mustParse(fset *token.FileSet, input string) *ast.File {
	f, err := parser.ParseFile(fset, "", input, 0)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func parseStructs(path string) ([]string, error) {
	var sts []string

	fset := token.NewFileSet()
	inputBlob, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("os.ReadFile: %s", err)
	}
	astFile := mustParse(fset, string(inputBlob))
	info := types.Info{
		// Types: make(map[ast.Expr]types.TypeAndValue),
		Defs: make(map[*ast.Ident]types.Object),
		// Uses:  make(map[*ast.Ident]types.Object),
	}

	// cfg := &packages.Config{
	// 	Mode: packages.NeedTypes | packages.NeedSyntax,
	// }
	// _, err = packages.Load(cfg, "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// FIXME: super slow due to imports in file itself but works.
	conf := types.Config{
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
		Importer:                 importer.ForCompiler(fset, "source", nil),
	}
	_, err = conf.Check("rest", fset, []*ast.File{astFile}, &info)
	if err != nil {
		log.Fatal(err)
	}
	for ident, obj := range info.Defs {
		if tn, ok := obj.(*types.TypeName); ok {
			if _, ok := tn.Type().Underlying().(*types.Struct); ok {
				if (ident.IsExported() && privateOnly) || (!ident.IsExported() && publicOnly) {
					continue
				}
				sts = append(sts, ident.Name)
				// fmt.Printf("  Defined at: %s\n", fset.Position(ident.Pos()))
			}
		}
	}

	return sts, nil
}

func main() {
	structsCmd.BoolVar(&privateOnly, "private-only", false, "Find private structs only")
	structsCmd.BoolVar(&publicOnly, "public-only", false, "Find public structs only")

	cmd, ok := subcommands[os.Args[1]]
	if !ok {
		for _, fs := range subcommands {
			fs.Usage()
		}

		return
	}

	err := cmd.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}

	// go build -o bin/build/ast-parser cmd/ast-parser/main.go; ast-parser find-structs internal/rest/models.go
	switch os.Args[1] {
	case "find-interfaces":
		os.Exit(0)
	case "find-structs":
		structs := map[string]struct{}{}
		for _, filename := range cmd.Args() {
			err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && filepath.Ext(path) == ".go" && filepath.Base(path) != "_test.go" {
					sts, err := parseStructs(path)
					if err != nil {
						return fmt.Errorf("parseStructs: %w", err)
					}
					for _, st := range sts {
						structs[st] = struct{}{}
					}
				}

				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		sts := maps.Keys(structs)
		sort.Slice(sts, func(i, j int) bool {
			return sts[i] < sts[j]
		})
		fmt.Printf("sts: %v\n", sts)
		os.Exit(0)
	}
}
