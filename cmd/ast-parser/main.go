/**
 * ast-parser allows for reliable extraction of structs and interfaces in the given files or directories.
 * For simple cases where structs are defined directly with a type definition (not a generic instantiation),
 * see bash helper go-utils.
 */
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

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
	findStructsFlagSet      flagSet = "find-structs"
	findInterfacesFlagSet   flagSet = "find-interfaces"
	deleteRedeclaredFlagSet flagSet = "delete-redeclared"
	verifyNoImportFlagSet   flagSet = "verify-no-import"
)

//nolint:gochecknoglobals
var (
	privateOnly, publicOnly, excludeGenerics, createGenericsInstanceMap bool
	importsStr                                                          string

	structsCmd          = flag.NewFlagSet(string(findStructsFlagSet), flag.ExitOnError)
	interfacesCmd       = flag.NewFlagSet(string(findInterfacesFlagSet), flag.ExitOnError)
	deleteRedeclaredCmd = flag.NewFlagSet(string(deleteRedeclaredFlagSet), flag.ExitOnError)
	verifyNoImportCmd   = flag.NewFlagSet(string(verifyNoImportFlagSet), flag.ExitOnError)

	subcommands = map[string]*flag.FlagSet{
		structsCmd.Name():          structsCmd,
		interfacesCmd.Name():       interfacesCmd,
		deleteRedeclaredCmd.Name(): deleteRedeclaredCmd,
		verifyNoImportCmd.Name():   verifyNoImportCmd,
	}

	wg sync.WaitGroup

	reflectTypeMap = &ReflectTypeMap{Map: map[string]map[string]string{}}
)

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

func verifyNoImport(path string, imports []string, errCh chan<- error) {
	defer wg.Done()

	for _, importPath := range imports {
		if _, found := importedPkgs.pkgs[importPath]; found {
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

func parseStructs(filepath string, resultCh chan<- []string, errCh chan<- error) {
	defer wg.Done()

	var sts []string

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
							if _, ok := obj.Type().Underlying().(*types.Struct); ok {
								structName := obj.Name()

								if isGeneric && excludeGenerics {
									fmt.Fprintf(os.Stderr, "Skipping generic struct %s\n", structName)

									continue
								}
								if (obj.Exported() && privateOnly) || (!obj.Exported() && publicOnly) {
									continue
								}

								if isGenericInstance(obj.Type().String()) {
									reflectTypeName := getReflectionType(obj.Type().String())
									reflectTypeMap.Add(structName, reflectTypeName)
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
	packages.NeedDeps | // necessary for resolving structs from package imports
	packages.NeedTypes | // necessary to get position information later, which contains filename that we can use to match against filepath
	packages.NeedSyntax |
	packages.NeedTypesInfo

type Pkgs struct {
	pkgs           map[string]struct{}
	mu             sync.Mutex
	redeclarations []string
}

var importedPkgs = Pkgs{
	pkgs:           make(map[string]struct{}),
	redeclarations: []string{},
	mu:             sync.Mutex{},
}

func (p *Pkgs) addPkg(pkg string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.pkgs[pkg] = struct{}{}
}

func (p *Pkgs) addRedeclaration(rd string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.redeclarations = append(p.redeclarations, rd)
}

func (p *Pkgs) String() string {
	return fmt.Sprintf("%s", p.pkgs)
}

func main() {
	structsCmd.BoolVar(&createGenericsInstanceMap, "create-generics-map", false, "Returns a JSON encoded map of structs to their generics reflection type name instead of a list of structs found")
	structsCmd.BoolVar(&excludeGenerics, "exclude-generics", false, "Find non generic structs only")
	structsCmd.BoolVar(&privateOnly, "private-only", false, "Find private structs only")
	structsCmd.BoolVar(&publicOnly, "public-only", false, "Find public structs only")
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

	// go build -o bin/build/ast-parser cmd/ast-parser/main.go; ast-parser find-structs internal/rest/models.go
	switch flag := os.Args[1]; flagSet(flag) {
	case deleteRedeclaredFlagSet:
	case findStructsFlagSet, findInterfacesFlagSet, verifyNoImportFlagSet:
		resultCh := make(chan []string)
		errCh := make(chan error)

		for _, pattern := range cmd.Args() {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				log.Fatal(err)
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
							case findStructsFlagSet:
								go parseStructs(path, resultCh, errCh)
							case verifyNoImportFlagSet:

								go func() {
									loadPackages(path)
									imports := strings.Split(importsStr, ",")
									verifyNoImport(path, imports, errCh)
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
					case findStructsFlagSet:
						go parseStructs(filename, resultCh, errCh)
					case verifyNoImportFlagSet:
						go func() {
							loadPackages(filename)
							imports := strings.Split(importsStr, ",")
							verifyNoImport(filename, imports, errCh)
						}()
					}
				}

				go func() {
					wg.Wait()
					close(resultCh)
					close(errCh)
				}()

				items := map[string]struct{}{}

				if flagSet(flag) == findStructsFlagSet || flagSet(flag) == findInterfacesFlagSet {
					for res := range resultCh {
						println("here 1")
						for _, st := range res {
							items[st] = struct{}{}
						}
					}
				}

				fmt.Printf("importedPkgs.redeclarations: %v\n", importedPkgs.redeclarations)
				errs := []error{}
				for err := range errCh {
					errs = append(errs, err)
				}
				if len(errs) > 0 {
					fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", errors.Join(errs...)))
					os.Exit(1)
				}

				if createGenericsInstanceMap {
					bytes, _ := json.MarshalIndent(reflectTypeMap, "", "  ")
					fmt.Println(string(bytes))
					os.Exit(0)
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

func loadPackages(filename string) {
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
				matches := re.FindStringSubmatch(terr.Msg)
				if len(matches) > 0 {
					fmt.Println("Redeclaration:", matches[1])
					importedPkgs.addRedeclaration(matches[1])
				}
			}
		}
		for _, ip := range p.Imports {
			importedPkgs.addPkg(ip.PkgPath)
		}
	}
}
