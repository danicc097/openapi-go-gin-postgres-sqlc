package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"golang.org/x/exp/maps"
	"golang.org/x/tools/go/packages"
)

//nolint:gochecknoglobals
var (
	privateOnly, publicOnly bool

	structsCmd    = flag.NewFlagSet("find-structs", flag.ExitOnError)
	interfacesCmd = flag.NewFlagSet("find-interfaces", flag.ExitOnError)

	subcommands = map[string]*flag.FlagSet{
		structsCmd.Name():    structsCmd,
		interfacesCmd.Name(): interfacesCmd,
	}

	wg sync.WaitGroup
)

func parseStructs(filepath string, resultCh chan<- []string, errCh chan<- error) {
	defer wg.Done()

	var sts []string

	loadConfig := &packages.Config{
		Fset: token.NewFileSet(),
		Mode: loadMode,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			if strings.Contains(filename, filepath) {
				fmt.Printf("parsing file: %v\n", filename)

				const mode = parser.AllErrors | parser.ParseComments
				return parser.ParseFile(fset, filename, src, mode)
			}

			return nil, nil
		},
	}

	pkgs, err := packages.Load(loadConfig, "file="+filepath)
	if err != nil {
		errCh <- err
		return
	}

	for _, pkg := range pkgs {
		for _, syn := range pkg.Syntax {
			for _, dec := range syn.Decls {
				if gen, ok := dec.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
					for _, spec := range gen.Specs {
						if ts, ok := spec.(*ast.TypeSpec); ok {
							obj, ok := pkg.TypesInfo.Defs[ts.Name]
							if !ok {
								continue
							}
							if _, ok := obj.Type().Underlying().(*types.Struct); ok {
								if (obj.Exported() && privateOnly) || (!obj.Exported() && publicOnly) {
									continue
								}

								sts = append(sts, obj.Name())
							}
						}
					}
				}
			}
		}
	}

	resultCh <- sts
}

const loadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedDeps |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedSyntax |
	packages.NeedTypesInfo

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
		resultCh := make(chan []string)
		errCh := make(chan error)

		for _, pattern := range cmd.Args() {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				log.Fatal(err)
			}
			for _, filename := range matches {
				fileInfo, err := os.Stat(filename)
				if err != nil {
					log.Fatalf("os.Stat: %s", err)
				}

				if fileInfo.IsDir() {
					err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if !info.IsDir() && filepath.Ext(path) == ".go" && filepath.Base(path) != "_test.go" {
							wg.Add(1)
							go parseStructs(path, resultCh, errCh)
						}

						return nil
					})
					if err != nil {
						log.Fatal(err)
					}
				} else {
					wg.Add(1)
					go parseStructs(filename, resultCh, errCh)
				}

				go func() {
					wg.Wait()
					close(resultCh)
					close(errCh)
				}()

				structs := map[string]struct{}{}
				for sts := range resultCh {
					for _, st := range sts {
						structs[st] = struct{}{}
					}
				}

				if len(errCh) > 0 {
					err := <-errCh
					log.Fatal(err)
				}

				sts := maps.Keys(structs)
				sort.Slice(sts, func(i, j int) bool {
					return sts[i] < sts[j]
				})
				fmt.Printf("sts: %v\n", sts)
				os.Exit(0)
			}
		}
	}
}
