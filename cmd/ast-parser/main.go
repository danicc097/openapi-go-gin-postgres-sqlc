/**
 * ast-parser allows for reliable extraction of structs and interfaces in the given files or directories.
 * For simple cases where structs are defined directly with a type definition (not a generic instantiation),
 * see bash helper go-utils.
 */
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
	privateOnly, publicOnly, excludeGenerics bool

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
		// large packages still slow
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			if !strings.Contains(filename, filepath) {
				return nil, nil
			}

			const mode = parser.AllErrors | parser.ParseComments

			file, err := parser.ParseFile(fset, filename, src, mode)
			if err != nil {
				return nil, fmt.Errorf("parser.ParseFile: %w", err)
			}

			// Skip function bodies to speed up.
			// NOTE: no improvement clearing struct fields beforehand
			for _, decl := range file.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					funcDecl.Body = nil
				}
			}

			return file, nil
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
							isGeneric := ts.TypeParams != nil
							obj, ok := pkg.TypesInfo.Defs[ts.Name]
							if !ok {
								continue
							}
							if _, ok := obj.Type().Underlying().(*types.Struct); ok {
								structName := obj.Name()
								if (obj.Exported() && privateOnly) || (!obj.Exported() && publicOnly) {
									continue
								}

								fmt.Printf("obj.Name(): %v\n", obj.Name())
								fmt.Printf("obj.Type(): %v\n", obj.Type())

								if isGeneric && excludeGenerics {
									fmt.Fprintf(os.Stderr, "Skipping generic struct %s\n", structName)
									continue
								}

								sts = append(sts, structName)
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
	packages.NeedTypesInfo

func main() {
	structsCmd.BoolVar(&excludeGenerics, "exclude-generics", false, "Find non generic structs only")
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
	switch flag := os.Args[1]; flag {
	case "find-structs", "find-interfaces":
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
							if flag == "find-structs" {
								go parseStructs(path, resultCh, errCh)
							}
						}

						return nil
					})
					if err != nil {
						log.Fatal(err)
					}
				} else {
					wg.Add(1)
					if flag == "find-structs" {
						go parseStructs(filename, resultCh, errCh)
					}
				}

				go func() {
					wg.Wait()
					close(resultCh)
					close(errCh)
				}()

				items := map[string]struct{}{}
				for res := range resultCh {
					for _, st := range res {
						items[st] = struct{}{}
					}
				}

				if len(errCh) > 0 {
					err := <-errCh
					log.Fatal(err)
				}

				sortedItems := maps.Keys(items)
				sort.Slice(sortedItems, func(i, j int) bool {
					return sortedItems[i] < sortedItems[j]
				})
				fmt.Println(strings.Join(sortedItems, "\n"))
			}
		}
	}
}
