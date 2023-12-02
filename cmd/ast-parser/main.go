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
	"strings"

	"golang.org/x/tools/go/packages"
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

const loadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedDeps |
	packages.NeedTypes |
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
		loadConfig := &packages.Config{
			Fset: token.NewFileSet(),
			Mode: loadMode,
			ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
				if strings.Contains(filename, cmd.Args()[0]) {
					fmt.Printf("parsing file: %v\n", filename)

					// default behavior. could speed up even more when parsing a directory by ignoring function bodies, etc.
					const mode = parser.AllErrors | parser.ParseComments
					return parser.ParseFile(fset, filename, src, mode)
				}

				return nil, nil
			},
		}
		pkgs, err := packages.Load(loadConfig, "file="+cmd.Args()[0])
		if err != nil {
			panic(err)
		}

		for _, pkg := range pkgs {
			for _, syn := range pkg.Syntax {
				for _, dec := range syn.Decls {
					fmt.Printf("dec: %v\n", dec)
					if gen, ok := dec.(*ast.GenDecl); ok && gen != nil && gen.Tok == token.TYPE {
						// print doc comment of the type
						if gen.Doc != nil {
							fmt.Println(gen.Doc.List[0])
						}
						for _, spec := range gen.Specs {
							if ts, ok := spec.(*ast.TypeSpec); ok {
								obj, ok := pkg.TypesInfo.Defs[ts.Name]
								if !ok {
									continue
								}
								typeName, ok := obj.(*types.TypeName)
								if !ok {
									continue
								}
								named, ok := typeName.Type().(*types.Named)
								if !ok {
									continue
								}
								// print the full name of the type
								fmt.Println(named)

								_, ok = named.Underlying().(*types.Struct)
								if !ok {
									continue
								}

								fmt.Printf("s.String(): %v\n", named.Obj().Name())
							}
						}
					}
				}
			}
		}
	}
}
