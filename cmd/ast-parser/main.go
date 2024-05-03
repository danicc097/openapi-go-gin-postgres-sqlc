/**
 * ast-parser allows for reliable extraction of structs and interfaces in the given files or directories.
 * For simple cases where structs are defined directly with a type definition (not a generic instantiation),
 * see bash helper go-utils.
 */
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
	"golang.org/x/exp/maps"
	"golang.org/x/tools/go/packages"
)

type ReflectTypeMap struct {
	sync.Mutex
	Map map[string]map[string]string
}

func (m *ReflectTypeMap) Add(structName string, reflectType ReflectionType) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Map[reflectType.Pkg]; !ok {
		m.Map[reflectType.Pkg] = map[string]string{}
	}

	m.Map[reflectType.Pkg][reflectType.Name] = structName
}

func (m *ReflectTypeMap) MarshalJSON() ([]byte, error) {
	m.Lock()
	defer m.Unlock()

	return json.Marshal(m.Map)
}

type flagSet string

const (
	findStructsFlagSet    flagSet = "find-structs"
	findTypesFlagSet      flagSet = "find-types"
	findInterfacesFlagSet flagSet = "find-interfaces"
	findRedeclaredFlagSet flagSet = "find-redeclared"
	verifyNoImportFlagSet flagSet = "verify-no-import"
)

//nolint:gochecknoglobals
var (
	privateOnly, publicOnly, excludeGenerics, createGenericsInstanceMap, deleteRedeclared bool
	importsStr                                                                            string

	structsCmd        = flag.NewFlagSet(string(findStructsFlagSet), flag.ExitOnError)
	typesCmd          = flag.NewFlagSet(string(findTypesFlagSet), flag.ExitOnError)
	interfacesCmd     = flag.NewFlagSet(string(findInterfacesFlagSet), flag.ExitOnError)
	findRedeclaredCmd = flag.NewFlagSet(string(findRedeclaredFlagSet), flag.ExitOnError) // ast-parser find-redeclared internal/rest/models.spec.go [--delete]
	verifyNoImportCmd = flag.NewFlagSet(string(verifyNoImportFlagSet), flag.ExitOnError) // ast-parser verify-no-import --imports "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models" internal/rest/models.spec.go

	subcommands = map[string]*flag.FlagSet{
		typesCmd.Name():          typesCmd,
		structsCmd.Name():        structsCmd,
		interfacesCmd.Name():     interfacesCmd,
		findRedeclaredCmd.Name(): findRedeclaredCmd,
		verifyNoImportCmd.Name(): verifyNoImportCmd,
	}

	wg sync.WaitGroup

	reflectTypeMap = &ReflectTypeMap{Map: map[string]map[string]string{}}
)

func main() {
	for _, cmd := range []*flag.FlagSet{structsCmd, interfacesCmd, typesCmd} {
		cmd.BoolVar(&excludeGenerics, "exclude-generics", false, "Find non generic types only")
		cmd.BoolVar(&privateOnly, "private-only", false, "Find private types only")
		cmd.BoolVar(&publicOnly, "public-only", false, "Find public types only")
	}

	structsCmd.BoolVar(&createGenericsInstanceMap, "create-generics-map", false, "Returns a JSON encoded map of structs to their generics reflection type name instead of a list of structs found")
	findRedeclaredCmd.BoolVar(&deleteRedeclared, "delete", false, "Delete the nodes that declare duplicated types")
	verifyNoImportCmd.StringVar(&importsStr, "imports", "", "Comma-separated list of imports to verify")

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

	if createGenericsInstanceMap {
		excludeGenerics = true
	}

	// go build -o bin/build/ast-parser cmd/ast-parser/main.go; ast-parser find-structs internal/rest/models.spec.go
	switch flag := os.Args[1]; flagSet(flag) {
	case findStructsFlagSet, findInterfacesFlagSet, verifyNoImportFlagSet, findRedeclaredFlagSet, findTypesFlagSet:
		itemsCh := make(chan []string)
		errCh := make(chan error)

		for _, pattern := range cmd.Args() {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				log.Fatalf("filepath.Glob: %s", err)
			}

			// validation
			for _, filename := range matches {
				var err error
				fileInfo, err := os.Stat(filename)
				if err != nil {
					log.Fatalf("os.Stat: %s", err)
				}
				if fileInfo.IsDir() && deleteRedeclared {
					log.Fatal("cannot delete redeclared types for multiple files at once") // not worth it. call independently
				}
			}

			for _, filename := range matches {
				var err error
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
							switch flagSet(flag) {
							case findTypesFlagSet:
								go func() {
									defer wg.Done()
									parseTypes(parseModeAll, path, itemsCh, errCh)
								}()
							case findStructsFlagSet:
								go func() {
									defer wg.Done()
									parseTypes(parseModeStruct, path, itemsCh, errCh)
								}()
							case verifyNoImportFlagSet:
								go func() {
									defer wg.Done()
									loadPackages(path)
									imports := strings.Split(importsStr, ",")
									verifyNoImport(path, imports, errCh)
								}()
							case findRedeclaredFlagSet:
								go func() {
									defer wg.Done()
									typeErrors := loadPackages(path)
									itemsCh <- typeErrors.redeclarations
								}()
							}
						}

						return nil
					})
					if err != nil {
						log.Fatal(err)
					}
				} else {
					wg.Add(1)
					switch flagSet(flag) {
					case findTypesFlagSet:
						go func() {
							defer wg.Done()
							parseTypes(parseModeAll, filename, itemsCh, errCh)
						}()
					case findStructsFlagSet:
						go func() {
							defer wg.Done()
							parseTypes(parseModeStruct, filename, itemsCh, errCh)
						}()
					case verifyNoImportFlagSet:
						go func() {
							defer wg.Done()
							loadPackages(filename)
							imports := strings.Split(importsStr, ",")
							verifyNoImport(filename, imports, errCh)
						}()
					case findRedeclaredFlagSet:
						go func() {
							defer wg.Done()
							typeErrors := loadPackages(filename)
							itemsCh <- typeErrors.redeclarations
						}()
					}
				}

				go func() {
					wg.Wait()
					close(itemsCh)
					close(errCh)
				}()

				items := map[string]struct{}{}

				if flagSet(flag) == findStructsFlagSet ||
					flagSet(flag) == findTypesFlagSet ||
					flagSet(flag) == findInterfacesFlagSet ||
					flagSet(flag) == findRedeclaredFlagSet {
					for res := range itemsCh {
						for _, st := range res {
							items[st] = struct{}{}
						}
					}
				}

				verifyNoErrorsOrExit(errCh)

				if deleteRedeclared {
					paths := []string{"internal/rest/openapi_server.gen.go"} // "internal/rest/openapi_types.gen.go",
					for _, path := range paths {
						src, err := os.ReadFile(path)
						if err != nil {
							fmt.Printf("[WARNING] Could not read %s: %v\n", path, err)
							continue
						}
						fileAST, err := decorator.Parse(src)
						if err != nil {
							log.Fatalf("Error parsing file: %v", err)
						}
						for typeName := range items {
							fileAST = deleteNodesFromAST(fileAST, typeName)
						}

						fmt.Printf("deleting duplicate rest models in %s...\n", path)
						err = writeAST(path, fileAST)
						if err != nil {
							log.Fatalf("Failed to write modified AST to file: %v", err)
						}
					}
					os.Exit(0)
				}

				if createGenericsInstanceMap {
					bytes, _ := json.MarshalIndent(reflectTypeMap, "", "  ")
					fmt.Println(string(bytes))
					os.Exit(0)
				}

				// default to printing search results if any
				sortedItems := maps.Keys(items)
				sort.Slice(sortedItems, func(i, j int) bool {
					return sortedItems[i] < sortedItems[j]
				})
				if len(sortedItems) > 0 {
					fmt.Println(strings.Join(sortedItems, "\n"))
				}
			}
		}
	}
}

func verifyNoErrorsOrExit(errCh chan error) {
	errs := []error{}
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", errors.Join(errs...))
		os.Exit(1)
	}
}

func isGenericInstance(s string) bool {
	r := regexp.MustCompile(`\.?(.*\[.*)`)

	return len(r.FindStringSubmatch(s)) > 0
}

type ReflectionType struct {
	Name string
	Pkg  string
}

// getReflectionType converts a generic instance type.Type string to its reflect.Type.Name.
func getReflectionType(s string) ReflectionType {
	r := regexp.MustCompile(`.*/([a-zA-Z_]{1}[a-zA-Z0-9_]*)*.*\.(.*\[.*)`) // greedily match longest until last dot

	matches := r.FindStringSubmatch(s)
	if len(matches) > 0 {
		return ReflectionType{
			Name: matches[2],
			Pkg:  matches[1],
		}
	}

	return ReflectionType{Name: s}
}

// FIXME: get list of imports from parsing file directly. packages.Load gives everything.
func verifyNoImport(path string, imports []string, errCh chan<- error) {
	fileImports, err := getImports(path)
	if err != nil {
		errCh <- fmt.Errorf("could not getImports in %s: %s", path, err)
		return
	}
	for _, importPath := range imports {
		if slices.Contains(fileImports, importPath) {
			errCh <- fmt.Errorf("restricted import detected in %s: %s", path, importPath)
			return
		}
	}
}

var loadConfig = &packages.Config{
	Fset: token.NewFileSet(),
	Mode: loadMode,
	// large packages still slow
	ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
		// IMPORTANT: we need to parser.ParseFile every file and package.

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

type parseMode string

const (
	parseModeStruct parseMode = "struct"
	parseModeAll    parseMode = "all"
)

func parseTypes(mode parseMode, filepath string, resultCh chan<- []string, errCh chan<- error) {
	var typs []string

	pkgs, err := packages.Load(loadConfig, "file="+filepath)
	if err != nil {
		errCh <- err
		return
	}

	for _, pkg := range pkgs {
		for _, syn := range pkg.Syntax {
			for _, dec := range syn.Decls {
				if gen, ok := dec.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
					position := pkg.Fset.Position(gen.Pos())
					// if we want to parse a complete package, pass glob as cli arg
					if !strings.Contains(position.Filename, filepath) {
						continue
					}
					for _, spec := range gen.Specs {
						if ts, ok := spec.(*ast.TypeSpec); ok {
							isGeneric := ts.TypeParams != nil
							obj, ok := pkg.TypesInfo.Defs[ts.Name]
							if !ok {
								continue
							}

							typeName := obj.Name()

							if isGeneric && excludeGenerics {
								fmt.Fprintf(os.Stderr, "Skipping generic type %s\n", typeName)

								continue
							}

							if (obj.Exported() && privateOnly) || (!obj.Exported() && publicOnly) {
								continue
							}
							switch mode {
							case parseModeAll:
								typs = append(typs, typeName)
							case parseModeStruct:
								if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
									continue
								}
								if isGenericInstance(obj.Type().String()) {
									reflectTypeName := getReflectionType(obj.Type().String())
									reflectTypeMap.Add(typeName, reflectTypeName)
								}

								typs = append(typs, typeName)
							default:
								errCh <- fmt.Errorf("unknown parse mode: %s", mode)
								return
							}
						}
					}
				}
			}
		}
	}

	resultCh <- typs
}

const loadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedDeps | // necessary for resolving structs from package imports
	packages.NeedTypes | // necessary to get position information later, which contains filename that we can use to match against filepath
	packages.NeedSyntax |
	packages.NeedTypesInfo

type Pkgs struct {
	pkgs map[string]struct{}
	mu   sync.Mutex
}

var importedPkgs = &Pkgs{
	pkgs: map[string]struct{}{},
	mu:   sync.Mutex{},
}

func (p *Pkgs) addPkg(pkg string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.pkgs[pkg] = struct{}{}
}

func (p *Pkgs) String() string {
	return fmt.Sprintf("%s", p.pkgs)
}

type typeErrors struct {
	redeclarations []string
}

func loadPackages(filename string) typeErrors {
	res := typeErrors{}
	cfg := *loadConfig
	// bare minimum to get imports per file. speeds up considerably
	cfg.Mode = packages.NeedName |
		packages.NeedFiles |
		packages.NeedImports |
		packages.NeedTypes // TODO: to type check when we use remove duplicate delcarations
	pp, err := packages.Load(&cfg, "file="+filename)
	if err != nil {
		log.Fatalf("failed to load package: %v", err)
	}

	re := regexp.MustCompile(`^(.*)\sredeclared in this block.*`)

	for _, p := range pp {
		for _, terr := range p.TypeErrors {
			// internal codes: https://go.dev/src/internal/types/errors/codes.go
			// see https://tip.golang.org/src/internal/types/errors/codes_test.go
			// but for now must stick to regex.

			if e := terr.Msg; strings.Contains(e, "redeclared in this block") {
				// in our case for now we can just hardcode internal/rest/openapi_types.gen.go since its the only use case
				// position := p.Fset.Position(terr.Pos)
				// fmt.Printf("position: %v\n", position)
				matches := re.FindStringSubmatch(terr.Msg)
				if len(matches) > 0 {
					res.redeclarations = append(res.redeclarations, matches[1])
				}
			}
		}
		for _, ip := range p.Imports {
			importedPkgs.addPkg(ip.PkgPath)
		}
	}

	return res
}

// Function to delete nodes from the AST.
func deleteNodesFromAST(file *dst.File, typeNameToDelete string) *dst.File {
	dstutil.Apply(file, func(c *dstutil.Cursor) bool {
		n := c.Node()
		if genDecl, ok := n.(*dst.GenDecl); ok && genDecl.Tok == token.TYPE {
			var specsToDelete []dst.Spec
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*dst.TypeSpec); ok && typeSpec.Name.Name == typeNameToDelete {
					// genDecl.Decs.Start.Append(fmt.Sprintf("/* Removed duplicated type %s */", typeNameToDelete))
					specsToDelete = append(specsToDelete, spec)
				}
			}

			// Remove the specsToDelete from the GenDecl
			genDecl.Specs = removeSpecs(genDecl.Specs, specsToDelete...)

			// Check if there are remaining specs in the GenDecl
			if len(genDecl.Specs) > 0 {
				c.Replace(genDecl)
			} else {
				c.Delete()
			}
		}

		return true
	}, nil)

	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), file)

	return file
}

// Function to remove specs from a list of specs.
func removeSpecs(specs []dst.Spec, specsToRemove ...dst.Spec) []dst.Spec {
	var newSpecs []dst.Spec
	for _, spec := range specs {
		var found bool
		for _, specToRemove := range specsToRemove {
			if spec == specToRemove {
				found = true
				break
			}
		}
		if !found {
			newSpecs = append(newSpecs, spec)
		}
	}
	return newSpecs
}

func writeAST(filePath string, file *dst.File) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	if err := decorator.Fprint(f, file); err != nil {
		return fmt.Errorf("failed to write AST to %s: %w", filePath, err)
	}

	return nil
}

// getImports returns a slice of strings representing the imported packages in the specified file.
func getImports(filePath string) ([]string, error) {
	// Set up the file set and parse the file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}

	// Extract imported packages from the file's AST
	var imports []string
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			for _, spec := range genDecl.Specs {
				if importSpec, ok := spec.(*ast.ImportSpec); ok {
					path, err := strconv.Unquote(importSpec.Path.Value)
					if err != nil {
						return nil, err
					}
					imports = append(imports, path)
				}
			}
		}
	}

	return imports, nil
}
