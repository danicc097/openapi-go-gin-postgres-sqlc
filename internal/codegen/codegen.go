package codegen

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	// internalformat "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var (
	handlerRegex     = regexp.MustCompile("api_(.*).go")
	OperationIDRegex = regexp.MustCompile("^[a-zA-Z0-9]*$")
)

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}

	return false
}

func ensureUnique[T comparable](s []T) error {
	set := make(map[T]struct{})
	for _, element := range s {
		if _, ok := set[element]; ok {
			return fmt.Errorf("element %T not unique", element)
		}
		set[element] = struct{}{}
	}

	return nil
}

//go:embed templates
var templateFiles embed.FS

type PreGen struct {
	stderr       io.Writer
	specPath     string
	opIDAuthPath string
	operations   map[string][]string
}

// New returns a new internal code generator.
func New(stderr io.Writer, specPath string, opIDAuthPath string) *PreGen {
	operations := make(map[string][]string)

	return &PreGen{
		stderr:       stderr,
		specPath:     specPath,
		opIDAuthPath: opIDAuthPath,
		operations:   operations,
	}
}

// validateSpec validates an OpenAPI 3.0 specification.
func (o *PreGen) validateSpec() error {
	_, err := rest.ReadOpenAPI(o.specPath)
	if err != nil {
		return err
	}

	return nil
}

func (o *PreGen) ValidateProjectSpec() error {
	if err := o.validateSpec(); err != nil {
		return fmt.Errorf("validate spec: %w", err)
	}

	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	return nil
}

func (o *PreGen) Generate() error {
	// if err := o.validateSpec(); err != nil {
	// 	return fmt.Errorf("validate spec: %w", err)
	// }

	// if err := internal.GenerateConfigTemplate(); err != nil {
	// 	return fmt.Errorf("internal.GenerateConfigTemplate: %w", err)
	// }

	// necessary before generation
	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	if err := o.generateOpIDs(); err != nil {
		return fmt.Errorf("generateOpIDs: %w", err)
	}

	if err := o.generateOpIDAuthMiddlewares(); err != nil {
		return fmt.Errorf("generateOpIDAuthMiddlewares: %w", err)
	}

	return nil
}

type AuthInfo struct {
	Scopes                 []string `json:"scopes"`
	Role                   string   `json:"role"`
	RequiresAuthentication bool     `json:"requiresAuthentication"`
}

type opIDAuthInfo = map[rest.OperationID]AuthInfo

// generateOpIDAuthMiddlewares generates middlewares based on role and scopes restrictions on an operation ID.
func (o *PreGen) generateOpIDAuthMiddlewares() error {
	opIDAuthInfos := make(opIDAuthInfo)

	opIDAuthInfoBlob, err := os.ReadFile(o.opIDAuthPath)
	if err != nil {
		return fmt.Errorf("opIDAuthInfo: %w", err)
	}
	if err := json.Unmarshal(opIDAuthInfoBlob, &opIDAuthInfos); err != nil {
		return fmt.Errorf("opIDAuthInfo json.Unmarshal: %w", err)
	}
	// internalformat.PrintJSON(opIDAuthInfos)

	funcs := template.FuncMap{
		"stringsJoin": strings.Join,
		"stringsJoinSlice": func(elems []string, prefix string, suffix string, sep string) string {
			for i, e := range elems {
				elems[i] = prefix + e + suffix
			}

			return strings.Join(elems, sep)
		},
	}

	tmpl := "templates/api_auth_middlewares.tmpl"
	name := path.Base(tmpl)
	t := template.Must(template.New(name).Funcs(funcs).ParseFS(templateFiles, tmpl))
	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Operations": opIDAuthInfos,
	}

	if err := t.Execute(buf, params); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	// internalformat.PrintJSON(buf.String())

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("could not format opID auth middlewares: %w", err)
	}

	fname := "internal/rest/api_auth_middlewares.gen.go"

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o660)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", fname, err)
	}

	if _, err = f.Write(src); err != nil {
		return fmt.Errorf("could not write opID auth middlewares: %w", err)
	}

	return nil
}

// analyzeSpec ensures specific rules for codegen are met and extracts necessary data.
func (o *PreGen) analyzeSpec() error {
	schemaBlob, err := os.ReadFile(o.specPath)
	if err != nil {
		return fmt.Errorf("error opening schema file: %w", err)
	}

	sl := openapi3.NewLoader()

	openapi, err := sl.LoadFromData(schemaBlob)
	if err != nil {
		return fmt.Errorf("error loading openapi spec: %w", err)
	}

	if err = openapi.Validate(sl.Context); err != nil {
		return fmt.Errorf("error validating openapi spec: %w", err)
	}

	for path, pi := range openapi.Paths {
		ops := pi.Operations()
		for method, v := range ops {
			if v.OperationID == "" {
				return fmt.Errorf("path %q: method %q: operationId is required for codegen", path, method)
			}

			if !OperationIDRegex.MatchString(v.OperationID) {
				return fmt.Errorf("path %q: method %q: operationId %q does not match pattern %q",
					path, method, v.OperationID, OperationIDRegex.String())
			}

			if len(v.Tags) > 1 {
				return fmt.Errorf("path %q: method %q: at most one tag is permitted for codegen", path, method)
			}

			t := "default"
			if len(v.Tags) > 0 {
				t = strings.ToLower(v.Tags[0])
			}

			o.operations[t] = append(o.operations[t], v.OperationID)
		}
		for t, opIDs := range o.operations {
			sort.Slice(opIDs, func(i, j int) bool {
				return opIDs[i] < opIDs[j]
			})
			o.operations[t] = opIDs
		}
	}

	return nil
}

// generateOpIDs fills in a template with all operation IDs to a dest.
func (o *PreGen) generateOpIDs() error {
	funcs := template.FuncMap{
		// "stringsJoin": func(elems []string, prefix string, suffix string, sep string) string {
		// 	for i, e := range elems {
		// 		elems[i] = prefix + e + suffix
		// 	}

		// 	return strings.Join(elems, sep)
		// },
	}

	tmpl := "templates/operation_ids.tmpl"
	name := path.Base(tmpl)
	t := template.Must(template.New(name).Funcs(funcs).ParseFS(templateFiles, tmpl))

	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Operations": o.operations,
	}

	if err := t.Execute(buf, params); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("could not format opId template: %w", err)
	}

	fname := path.Join("internal/rest/operation_ids.gen.go")

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o660)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", fname, err)
	}

	if _, err = f.Write(src); err != nil {
		return fmt.Errorf("could not write opId template: %w", err)
	}

	return nil
}

// EnsureFunctionsForOperationIDs checks if each operation ID has a corresponding function in the Go AST.

func parseAST(reader io.Reader) (*ast.File, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", fileContents, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	return file, nil
}

func getHandlersMethods(file *ast.File) []string {
	var functions []string

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv == nil {
				return false
			}
			if len(x.Recv.List) == 1 {
				recvType, ok := x.Recv.List[0].Type.(*ast.StarExpr)
				if ok {
					ident, ok := recvType.X.(*ast.Ident)
					if ok && ident.Name == "Handlers" {
						functions = append(functions, x.Name.Name)
					}
				}
			}
		}

		return true
	})

	return functions
}
