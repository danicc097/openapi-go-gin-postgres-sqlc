package codegen

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
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
	"github.com/iancoleman/strcase"
	"github.com/kenshaw/snaker"

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

type operationIDMethod struct {
	// file path to handlers file
	handlersFile    string
	methodSignature string
	comment         string
}

type CodeGen struct {
	stderr                            io.Writer
	specPath                          string
	opIDAuthPath                      string
	handlersPath                      string
	operations                        map[string][]string
	missingOperationIDImplementations map[string]struct{}
	serverInterfaceMethods            map[string]operationIDMethod
}

// New returns a new internal code generator.
func New(stderr io.Writer, specPath, opIDAuthPath, handlersPath string) *CodeGen {
	operations := make(map[string][]string)

	return &CodeGen{
		stderr:                            stderr,
		specPath:                          specPath,
		opIDAuthPath:                      opIDAuthPath,
		operations:                        operations,
		handlersPath:                      handlersPath,
		missingOperationIDImplementations: map[string]struct{}{},
		serverInterfaceMethods:            map[string]operationIDMethod{},
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
	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	if err := o.ensureHandlerMethodsExist(); err != nil {
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

func (o *CodeGen) ImplementServer() error {
	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	o.getServerInterfaceMethods()

	// reuses logic to extract missing methods
	if err := o.ensureHandlerMethodsExist(); err != nil {
		return fmt.Errorf("tag methods: %w", err)
	}

	if err := o.implementServerInterfaceMethods(); err != nil {
		return fmt.Errorf("implement server interface methods: %w", err)
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

	for path, pi := range openapi.Paths.Map() {
		ops := pi.Operations()
		for method, op := range ops {
			if op.OperationID == "" {
				errors = append(errors, fmt.Errorf("path %q: method %q: operationId is required for codegen", path, method).Error())
			}

			if !OperationIDRE.MatchString(op.OperationID) {
				errors = append(errors, fmt.Errorf("path %q: method %q: operationId %q does not match pattern %q",
					path, method, op.OperationID, OperationIDRE.String()).Error())
			}

			if len(op.Tags) > 1 {
				errors = append(errors, fmt.Errorf("path %q: method %q: at most one tag is permitted for codegen", path, method).Error())
			}

			t := "default"
			if len(op.Tags) > 0 {
				t = op.Tags[0]
			}

			if !isValidFilename(t) {
				errors = append(errors, fmt.Errorf("path %q: method %q: tag must be a valid filename with pattern %q", path, method, validFilenameRE.String()).Error())
			}

			o.operations[t] = append(o.operations[t], op.OperationID)
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
					if ok && ident.Name == "StrictHandlers" {
						functions = append(functions, x.Name.Name)
					}
				}
			}
		}

		return true
	})

	return functions
}

// removeAndAppendHandlersMethod removes a StrictHandlers method with the given name from its source file
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

// ensureHandlerMethodsExist parses the AST of each api_<lowercase of tag>.go file and
// ensure it contains function methods for each operation ID.
func (o *CodeGen) ensureHandlerMethodsExist() error {
	var errs []string

	for tag := range o.operations {
		snakeTag := strcase.ToSnake(tag)
		handlersPath := filepath.Join(o.handlersPath, fmt.Sprintf("api_%s.go", snakeTag))
		if _, err := os.Stat(handlersPath); err != nil {
			errs = append(errs, fmt.Sprintf("missing file %s for new tag %q", handlersPath, tag))
		}
	}

	tagFilePaths, err := filepath.Glob(filepath.Join(o.handlersPath, "api_*.go"))
	if err != nil {
		return fmt.Errorf("failed to find api_<tag>.go files: %w", err)
	}

	for _, tagFilePath := range tagFilePaths {
		if strings.HasSuffix(tagFilePath, "_test.go") {
			continue
		}
		matches := handlersFileTagRE.FindStringSubmatch(filepath.Base(tagFilePath))
		if len(matches) < 2 {
			return fmt.Errorf("failed to extract tag from file name: %s", tagFilePath)
		}
		tag := strcase.ToLowerCamel(matches[1])

		apiFileContent, err := os.ReadFile(tagFilePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", tagFilePath, err)
		}

		file, err := parseAST(bytes.NewReader(apiFileContent))
		if err != nil {
			return fmt.Errorf("failed to parse file %s: %w", tagFilePath, err)
		}

		// operation ids are preprocessed to pascal case
		functions := getHandlersMethods(file)

		restOfOpIDs := o.getOperationIDDifference(tag)

	fn:
		for _, opID := range functions {
			if contains(restOfOpIDs, opID) {
				correspondingTag := o.findTagByOpID(opID)
				snakeTag := snaker.CamelToSnake(correspondingTag)

				correctFilePath := filepath.Join(o.handlersPath, fmt.Sprintf("api_%s.go", snakeTag))
				content, err := os.ReadFile(correctFilePath)
				if err != nil {
					errs = append(errs, fmt.Sprintf("misplaced method for operation ID %q - should be in api_%s.go (file does not exist)", opID, snakeTag))

					break fn
				}

				correctFile, err := parseAST(bytes.NewReader(content))
				if err != nil {
					return fmt.Errorf("failed to parse file %s: %w", tagFilePath, err)
				}

				if path.Base(tagFilePath) != path.Base(correctFilePath) {
					fmt.Printf("Moving handler %q to correct file (%s -> %s)\n", opID, path.Base(tagFilePath), path.Base(correctFilePath))
					removeAndAppendHandlersMethod(file, correctFile, opID)
				}

				err = writeAST(correctFilePath, correctFile)
				if err != nil {
					return fmt.Errorf("failed to write AST to %s: %w", correctFilePath, err)
				}

				err = writeAST(tagFilePath, file)
				if err != nil {
					return fmt.Errorf("failed to write AST to %s: %w", tagFilePath, err)
				}

				return nil
			}
		}

		for _, opID := range o.operations[tag] {
			if !contains(functions, opID) {
				// NOTE: not worth syncing all opIDs to keep signature changes in sync, we get decent enough errors
				// to handle bad interface implementations, such as
				// have CreateTeam(*gin.Context, models.Project)
				// want CreateTeam(*gin.Context, models.Project, int)
				o.missingOperationIDImplementations[opID] = struct{}{}
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}

	return nil
}

func (o *CodeGen) implementServerInterfaceMethods() error {
	for opID := range o.missingOperationIDImplementations {
		fmt.Printf("Implementing missing server interface method for operation ID: %v\n", opID)

		m := o.serverInterfaceMethods[opID]

		methodStr := fmt.Sprintf(`
		func (h *StrictHandlers) %s%s {
			c.JSON(http.StatusNotImplemented, "not implemented")

			return nil, nil
		}
	`, strcase.ToCamel(opID), m.methodSignature)

		// we are only here if the method is missing, so just append it
		f, err := os.OpenFile(m.handlersFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return fmt.Errorf("could not open file: %s", err)
		}
		defer f.Close()
		if _, err := f.WriteString(methodStr); err != nil {
			return fmt.Errorf("could not write to file: %s", err)
		}
	}

	return nil
}

// getServerInterfaceMethods returns the generated server interface methods
// indexed by operation id.
func (o *CodeGen) getServerInterfaceMethods() map[string]operationIDMethod {
	src, err := os.ReadFile("internal/rest/openapi_server.gen.go")
	if err != nil {
		log.Fatalf("Error reading server interface file: %s", err)
	}

	// `import-mapping` oapi config generates unnamed imports
	src = bytes.ReplaceAll(src, []byte("externalRef0"), []byte("models"))

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "src.go", string(src), parser.AllErrors|parser.ParseComments)
	if err != nil {
		log.Fatalf("Error parsing server interface file: %s", err)
	}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok || typeSpec.Name.String() != "StrictServerInterface" {
				continue
			}

			for _, method := range interfaceType.Methods.List {
				operationID := method.Names[0].Name
				funcType, ok := method.Type.(*ast.FuncType)
				if !ok {
					continue
				}
				// params := extractParameters(funcType)
				// returns := extractReturns(funcType)
				correspondingTag := o.findTagByOpID(operationID)
				snakeTag := strcase.ToSnake(correspondingTag)

				var buf bytes.Buffer
				printer.Fprint(&buf, token.NewFileSet(), funcType)

				o.serverInterfaceMethods[operationID] = operationIDMethod{
					handlersFile:    filepath.Join(o.handlersPath, fmt.Sprintf("api_%s.go", snakeTag)),
					methodSignature: buf.String()[4:], // exclude func
					comment:         method.Doc.Text(),
				}
			}
		}
	}

	return o.serverInterfaceMethods
}

func extractParameters(ft *ast.FuncType) string {
	var params []string
	for _, field := range ft.Params.List {
		for _, name := range field.Names {
			params = append(params, fmt.Sprintf("%s %s", name.Name, exprToString(field.Type)))
		}
	}
	return strings.Join(params, ", ")
}

// FIXME: does nothing.
func _extractReturns(ft *ast.FuncType) string {
	var returns []string
	if ft.Results != nil {
		for _, field := range ft.Results.List {
			for _, name := range field.Names {
				returns = append(returns, fmt.Sprintf("(%s %s)", name.Name, exprToString(field.Type)))
			}
		}
	}
	return strings.Join(returns, ", ")
}

// exprToString converts an AST expression to its string representation.
func exprToString(expr ast.Expr) string {
	var buf strings.Builder
	printer.Fprint(&buf, token.NewFileSet(), expr)
	return buf.String()
}

// getOperationIDDifference returns the difference of all operation IDs
// and those associated to a given tag.
func (o *CodeGen) getOperationIDDifference(tag string) []string {
	var restOfOpIDs []string

	for opTag, opIDs := range o.operations {
		if opTag == tag {
			continue
		}
		restOfOpIDs = append(restOfOpIDs, opIDs...)
	}

	return restOfOpIDs
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

func writeAST(filePath string, file *dst.File) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	if err := decorator.Fprint(f, file); err != nil {
		return fmt.Errorf("failed to write AST to file %s: %w", filePath, err)
	}

	return nil
}
