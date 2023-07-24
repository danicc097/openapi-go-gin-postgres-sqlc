package codegen

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

var (
	handlersFileTagRE = regexp.MustCompile("api_(.*).go")
	OperationIDRE     = regexp.MustCompile("^[a-zA-Z0-9]*$")
	validFilenameRE   = regexp.MustCompile("^[a-zA-Z0-9_-]+$")
)

func isValidFilename(s string) bool {
	return validFilenameRE.MatchString(s)
}

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

type CodeGen struct {
	stderr       io.Writer
	specPath     string
	opIDAuthPath string
	handlersPath string
	operations   map[string][]string
}

// New returns a new internal code generator.
func New(stderr io.Writer, specPath, opIDAuthPath, handlersPath string) *CodeGen {
	operations := make(map[string][]string)

	return &CodeGen{
		stderr:       stderr,
		specPath:     specPath,
		opIDAuthPath: opIDAuthPath,
		operations:   operations,
		handlersPath: handlersPath,
	}
}

// validateSpec validates an OpenAPI 3.0 specification.
func (o *CodeGen) validateSpec() error {
	_, err := rest.ReadOpenAPI(o.specPath)
	if err != nil {
		return err
	}

	return nil
}

func (o *CodeGen) EnsureCorrectMethodsPerTag() error {
	if err := o.ensureFunctionMethods(); err != nil {
		return fmt.Errorf("tag methods: %w", err)
	}

	return nil
}

func (o *CodeGen) ValidateProjectSpec() error {
	if err := o.validateSpec(); err != nil {
		return fmt.Errorf("validate spec: %w", err)
	}

	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	return nil
}

func (o *CodeGen) Generate() error {
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
func (o *CodeGen) generateOpIDAuthMiddlewares() error {
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

	fname := path.Join(o.handlersPath, "api_auth_middlewares.gen.go")

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
func (o *CodeGen) analyzeSpec() error {
	var errors []string

	schemaBlob, err := os.ReadFile(o.specPath)
	if err != nil {
		errors = append(errors, fmt.Errorf("error opening schema file: %w", err).Error())
	}

	sl := openapi3.NewLoader()

	openapi, err := sl.LoadFromData(schemaBlob)
	if err != nil {
		errors = append(errors, fmt.Errorf("error loading openapi spec: %w", err).Error())
	}

	if err = openapi.Validate(sl.Context); err != nil {
		errors = append(errors, fmt.Errorf("error validating openapi spec: %w", err).Error())
	}

	for path, pi := range openapi.Paths {
		ops := pi.Operations()
		for method, v := range ops {
			if v.OperationID == "" {
				errors = append(errors, fmt.Errorf("path %q: method %q: operationId is required for codegen", path, method).Error())
			}

			if !OperationIDRE.MatchString(v.OperationID) {
				errors = append(errors, fmt.Errorf("path %q: method %q: operationId %q does not match pattern %q",
					path, method, v.OperationID, OperationIDRE.String()).Error())
			}

			if len(v.Tags) > 1 {
				errors = append(errors, fmt.Errorf("path %q: method %q: at most one tag is permitted for codegen", path, method).Error())
			}

			t := "default"
			if len(v.Tags) > 0 {
				t = v.Tags[0]
			}

			if !isValidFilename(t) {
				errors = append(errors, fmt.Errorf("path %q: method %q: tag must be a valid filename with pattern %q", path, method, validFilenameRE.String()).Error())
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

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
}

// generateOpIDs fills in a template with all operation IDs to a dest.
func (o *CodeGen) generateOpIDs() error {
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

	fname := path.Join(o.handlersPath, "operation_ids.gen.go")

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o660)
	if err != nil {
		return fmt.Errorf("could not open %s: %w", fname, err)
	}

	if _, err = f.Write(src); err != nil {
		return fmt.Errorf("could not write opId template: %w", err)
	}

	return nil
}

func parseAST(reader io.Reader) (*dst.File, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	if string(fileContents) == "" {
		return nil, fmt.Errorf("empty file")
	}

	fset := token.NewFileSet()
	file, err := decorator.ParseFile(fset, "", fileContents, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("decorator.ParseFile: %w", err)
	}

	return file, nil
}

func getHandlersMethods(file *dst.File) []string {
	var functions []string

	dst.Inspect(file, func(n dst.Node) bool {
		switch x := n.(type) {
		case *dst.FuncDecl:
			if x.Recv == nil {
				return false
			}
			if len(x.Recv.List) == 1 {
				recvType, ok := x.Recv.List[0].Type.(*dst.StarExpr)
				if ok {
					ident, ok := recvType.X.(*dst.Ident)
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

// removeAndAppendHandlersMethod removes a Handlers method with the given name from its source file
// and appends it to a target file.
func removeAndAppendHandlersMethod(src, target *dst.File, opID string) {
	for i, decl := range src.Decls {
		if fd, ok := decl.(*dst.FuncDecl); ok {
			if fd.Name.Name == opID {
				target.Decls = append(target.Decls, src.Decls[i])
				src.Decls = append(src.Decls[:i], src.Decls[i+1:]...)

				break
			}
		}
	}
}

// ensureFunctionMethods parses the AST of each api_<lowercase of tag>.go file and
// ensure it contains function methods for each operation ID.
func (o *CodeGen) ensureFunctionMethods() error {
	tagFiles, err := filepath.Glob(filepath.Join(o.handlersPath, "api_*.go"))
	if err != nil {
		return fmt.Errorf("failed to find api_<tag>.go files: %w", err)
	}

	var errs []string

	for _, tagFile := range tagFiles {
		matches := handlersFileTagRE.FindStringSubmatch(filepath.Base(tagFile))
		if len(matches) < 2 {
			return fmt.Errorf("failed to extract tag from file name: %s", tagFile)
		}
		tag := matches[1]

		apiFileContent, err := os.ReadFile(tagFile)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", tagFile, err)
		}

		file, err := parseAST(bytes.NewReader(apiFileContent))
		if err != nil {
			return fmt.Errorf("failed to parse file %s: %w", tagFile, err)
		}

		// operation ids are preprocessed to pascal case
		functions := getHandlersMethods(file)

		var restOfOpIDs []string
	tag:
		for opTag, opIDs := range o.operations {
			if opTag == tag {
				continue tag
			}
			restOfOpIDs = append(restOfOpIDs, opIDs...)
		}

	fn:
		for _, opID := range functions {
			if contains(restOfOpIDs, opID) {
				correspondingTag := o.findTagByOpID(opID)

				content, err := os.ReadFile(filepath.Join(o.handlersPath, fmt.Sprintf("api_%s.go", correspondingTag)))
				if err != nil {
					errs = append(errs, fmt.Sprintf("misplaced method for operation ID %q - should be in api_%s.go (file does not exist)", opID, correspondingTag))

					break fn
				}

				correctFile, err := parseAST(bytes.NewReader(content))
				if err != nil {
					return fmt.Errorf("failed to parse file %s: %w", tagFile, err)
				}

				removeAndAppendHandlersMethod(file, correctFile, opID)

				return nil
			}
		}

		for _, opID := range o.operations[tag] {
			if !contains(functions, opID) {
				errs = append(errs, fmt.Sprintf("missing function method for operation ID %q", opID))
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}

	return nil
}

// findTagByOpID returns the corresponding tag for a given operation ID.
func (o *CodeGen) findTagByOpID(opID string) string {
	for tag, opIDs := range o.operations {
		if contains(opIDs, opID) {
			return tag
		}
	}
	return ""
}
