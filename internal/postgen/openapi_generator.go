package postgen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	internalformat "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Dir string
type Tag string

var handlerRegex = regexp.MustCompile("api_(.*).go")
var operationIDRegex = regexp.MustCompile("^[a-zA-Z0-9]*$")

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

/*
Rationale:
We dont want users to manually add handlers and routes to handlers/*
If we let them, we wouldnt know at plain sight what was in the spec and what wasnt
and parsing will become a bigger mess.
Users can still add new methods to the struct. In case of generated methods conflicting with existing ones, generation will stop.
If we need a new route that cant be defined in the spec, e.g. fileserver,
we purposely want that out of the generated handler struct,
so its clear that its outside the spec.
It can still remain in handlers/* as long as its not api_*(!_test).go, e.g. fileserver.go
and it can still follow the same handlers struct pattern for all we care, it wont be touched.
IMPORTANT: if a method already exists in current but has no routes item (meaning
its probably some handler helper method created afterwards) then panic and alert
the user to rename. it shouldve been unexported or a function in the first place anyway.
*/
type OpenapiGenerator struct {
	conf     *Conf
	stderr   io.Writer
	cacheDir Dir
	spec     string
}

// NewOpenapiGenerator returns a new postgen OpenapiGenerator.
func NewOpenapiGenerator(conf *Conf, stderr io.Writer, cacheDir Dir, spec string) *OpenapiGenerator {
	return &OpenapiGenerator{
		conf:     conf,
		stderr:   stderr,
		cacheDir: cacheDir,
		spec:     spec,
	}
}

func (o *OpenapiGenerator) Generate() error {
	cb, err := o.getCommonBasenames()
	if err != nil {
		return err
	}

	handlers, err := o.analyzeHandlers(cb)
	if err != nil {
		return err
	}

	err = o.generateMergedFiles(handlers)
	if err != nil {
		return err
	}

	return err
}

// analyzeHandlers returns all necessary merging information about handlers, indexed
// by directory and tag.
func (o *OpenapiGenerator) analyzeHandlers(basenames []string) (map[Dir]map[Tag]HandlerFile, error) {
	handlers := make(map[Dir]map[Tag]HandlerFile)

	dirs := []Dir{o.conf.GenHandlersDir, o.conf.CurrentHandlersDir}
	for _, dir := range dirs {
		handlers[dir] = make(map[Tag]HandlerFile)

		for _, basename := range basenames {
			p := path.Join(string(dir), basename)

			blob, err := os.ReadFile(p)
			if err != nil {
				return nil, err
			}

			file, err := decorator.Parse(blob)
			if err != nil {
				return nil, err
			}

			tag := Tag(cases.Title(language.English).String(handlerRegex.FindStringSubmatch(basename)[1]))

			mm := inspectStruct(file, tag)
			rr := inspectRegisterNode(file, tag)
			routes := extractRoutes(rr)
			hf := HandlerFile{
				F:          file,
				Methods:    mm,
				RoutesNode: rr,
				Routes:     routes,
			}

			handlers[dir][tag] = hf
		}
	}

	err := o.findClashingMethodNames(basenames, handlers)
	if err != nil {
		return nil, err
	}

	return handlers, nil
}

// findClashingMethodNames ensures no previous methods that are not
// handlers conflict with a newly generated operation id.
func (o *OpenapiGenerator) findClashingMethodNames(basenames []string, handlers map[Dir]map[Tag]HandlerFile) error {
	clashes := []string{}

	for _, basename := range basenames {
		tag := Tag(cases.Title(language.English).String(handlerRegex.FindStringSubmatch(basename)[1]))

		for opID := range handlers[o.conf.GenHandlersDir][tag].Routes {
			// fmt.Printf("[%s] opID: %s\n", tag, opID)
			_, rok := handlers[o.conf.CurrentHandlersDir][tag].Routes[opID]
			_, mok := handlers[o.conf.CurrentHandlersDir][tag].Methods[opID]

			if !rok && mok {
				clashes = append(clashes, string(tag)+"->"+opID)
			}
		}
	}

	if len(clashes) > 0 {
		fmt.Fprintf(o.stderr, `
Error: conflicting method names
%s
Please rename either the affected method or operation id.
`, clashes)

		return errors.New("method name conflict")
	}

	return nil
}

// getCommonBasenames returns api filename (tag) intersections in current and raw gen dirs,
// and copies non-intersecting files to the out dir without further analysis.
func (o *OpenapiGenerator) getCommonBasenames() ([]string, error) {
	out := []string{}
	idx := 0

	currentBasenames, err := o.getAPIBasenames(o.conf.CurrentHandlersDir)
	if err != nil {
		return nil, err
	}

	genBasenames, err := o.getAPIBasenames(o.conf.GenHandlersDir)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(string(o.conf.OutHandlersDir), 0777)
	if err != nil {
		return nil, err
	}

	for _, genBasename := range genBasenames {
		if contains(currentBasenames, genBasename) {
			genBasenames[idx] = genBasename
			idx++

			continue
		}

		genBlob, err := os.ReadFile(path.Join(string(o.conf.GenHandlersDir), genBasename))
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(path.Join(string(o.conf.OutHandlersDir), genBasename), genBlob, 0600)
		if err != nil {
			return nil, err
		}
	}

	genBasenames = genBasenames[:idx]

	for _, currentBasename := range currentBasenames {
		if contains(genBasenames, currentBasename) {
			out = append(out, currentBasename)

			continue
		}

		currentBlob, err := os.ReadFile(path.Join(string(o.conf.CurrentHandlersDir), currentBasename))
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(path.Join(string(o.conf.OutHandlersDir), currentBasename), currentBlob, 0600)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

// generateService fills in a template with a default service struct to a dest.
func generateService(tag Tag, dest io.Writer) error {
	fmt.Printf("Creating service for tag: %s", tag)

	t := template.Must(template.New("").Parse(`package services

type {{.Tag}} struct {
}

// New{{.Tag}} returns a new {{.Tag}} service.
func New{{.Tag}}() *{{.Tag}} {
	return &{{.Tag}}{}
}

`))

	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Tag": cases.Title(language.English).String(string(tag)),
	}

	if err := t.Execute(buf, params); err != nil {
		return err
	}

	_, err := dest.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// generateOpIDs fills in a template with all operation IDs to a dest.
func generateOpIDs(operations map[string][]string, dest io.Writer) error {
	t := template.Must(template.New("").Parse(`// Code generated by postgen. DO NOT EDIT.

package rest

{{range $tag, $opIDs := .Operations}}
type {{$tag}}OpID string{{end}}

const ({{range $tag, $opIDs := .Operations}}
// Operation IDs for the '{{$tag}}' tag.
{{range $opIDs}}{{.}}                {{$tag}}OpID = "{{.}}"
{{end}}{{end}}
	)

	`))

	buf := &bytes.Buffer{}

	// TODO use type for those nested maps, else its a mess
	// plus key types: Tag string (tags), HandlerFile string (filenames)

	params := map[string]interface{}{
		"Operations": operations,
	}

	if err := t.Execute(buf, params); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// TODO run gofmt on it programmatically
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting opId template: %w", err)
	}
	_, err = dest.Write(formatted)
	if err != nil {
		return fmt.Errorf("error writing opId template: %w", err)
	}

	return nil
}

// replaceNodes replaces handler file nodes accordingly.
func replaceNodes(f *dst.File, genHf, currentHf HandlerFile, tag Tag, opID string) {
	dstutil.Apply(f, nil, func(c *dstutil.Cursor) bool {
		fn, isFn := c.Parent().(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		m := fn.Name.String()

		if rok && identok && ident.Name == string(tag) && m == "Register" {
			fn.Body.List[0] = genHf.RoutesNode
		}
		if rok && identok && ident.Name == string(tag) && m == "middlewares" {
			fn.Body = currentHf.Methods["middlewares"].Decl.Body
			fn.Type = genHf.Methods["middlewares"].Decl.Type
		}

		return true
	})
}

func (o *OpenapiGenerator) generateMergedFiles(handlers map[Dir]map[Tag]HandlerFile) error {
	// -- generate typed operation ids
	// This way we get a compilation error if the
	// spec doesn't use unique op ids, and intellisense for middlewares, etc.
	schemaBlob, err := os.ReadFile(o.spec)
	if err != nil {
		return fmt.Errorf("error opening schema file: %w", err)
	}
	openapi, err := openapi3.NewLoader().LoadFromData(schemaBlob)
	if err != nil {
		return fmt.Errorf("error loading schema: %w", err)
	}

	operations := make(map[string][]string)

	for path, pi := range openapi.Paths {
		ops := pi.Operations()
		for method, v := range ops {
			if v.OperationID == "" {
				return fmt.Errorf("path %q: method %q: operationId is required for postgen", path, method)
			}
			if !operationIDRegex.MatchString(v.OperationID) {
				return fmt.Errorf("path %q: method %q: operationId %q does not match pattern %q", path, method, v.OperationID, operationIDRegex.String())
			}
			if len(v.Tags) > 1 {
				return fmt.Errorf("path %q: method %q: at most one tag is permitted for postgen", path, method)
			}

			t := "default"
			if len(v.Tags) > 0 {
				t = strings.ToLower(v.Tags[0])
			}
			operations[t] = append(operations[t], internalformat.ToLowerCamel(v.OperationID))
		}
		for t, opIDs := range operations {
			sort.Slice(opIDs, func(i, j int) bool {
				return opIDs[i] < opIDs[j]
			})
			operations[t] = opIDs
		}
	}
	// no need to check for uniqueness of operation IDs, done by openapi-generator at this point
	s := path.Join(string(o.conf.OutHandlersDir), "operation_ids.gen.go")
	f, err := os.Create(s)
	if err != nil {
		return err
	}
	err = generateOpIDs(operations, f)
	if err != nil {
		return fmt.Errorf("error generating operation IDs: %w", err)
	}

	// -- generate handler files
	for tag, currentHF := range handlers[o.conf.CurrentHandlersDir] {
		outF, ok := dst.Clone(currentHF.F).(*dst.File)
		if !ok {
			return errors.New("clone file node fail")
		}

		// get generated operation ids as list
		gkk := make([]string, len(handlers[o.conf.GenHandlersDir][tag].Methods))
		i := 0

		for gk := range handlers[o.conf.GenHandlersDir][tag].Methods {
			gkk[i] = gk
			i++
		}

		sort.Slice(gkk, func(i, j int) bool {
			return gkk[i] < gkk[j]
		})

		genHF := handlers[o.conf.GenHandlersDir][tag]

		for _, gk := range gkk {
			if _, ok := currentHF.Methods[gk]; !ok {
				fmt.Printf("method %s not in current %s - appending generated method.\n", gk, tag)
				outF.Decls = append(outF.Decls, handlers[o.conf.GenHandlersDir][tag].Methods[gk].Decl)

				continue
			}

			replaceNodes(outF, genHF, currentHF, tag, gk)
		}

		buf := &bytes.Buffer{}

		f, err := os.Create(path.Join(string(o.conf.OutHandlersDir), "api_"+strings.ToLower(string(tag))+".go"))
		if err != nil {
			return err
		}

		if err := decorator.Fprint(buf, outF); err != nil {
			return err
		}

		_, err = f.Write(buf.Bytes())
		if err != nil {
			return err
		}
	}

	// -- generate default service if not exists
	for tag := range handlers[o.conf.GenHandlersDir] {
		if tag == "Default" {
			continue
		}
		s := path.Join(string(o.conf.OutServicesDir), strings.ToLower(string(tag))+".go")
		if _, err := os.Stat(s); errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(s)
			if err != nil {
				return err
			}

			err = generateService(tag, f)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Method struct {
	// Name is the method identifier.
	Name string
	// Decl is the method declaration node.
	Decl *dst.FuncDecl
}

type HandlerFile struct {
	// F is the node of the source file.
	F *dst.File
	// Methods represents all methods in the generated struct indexed by method name.
	Methods map[string]Method
	// RoutesNode represents the routes slice assignment node.
	RoutesNode *dst.AssignStmt
	// Routes represents convenient extracted fields from route elements
	// indexed by operation id.
	Routes map[string]Route
}

func (o *OpenapiGenerator) getAPIBasenames(src Dir) ([]string, error) {
	out := []string{}
	// glob uses https://pkg.go.dev/path#Match patterns. Test files will match
	paths, err := filepath.Glob(path.Join(string(src), "api_*.go"))
	if err != nil {
		return nil, err
	}

	if len(paths) == 0 && strings.HasSuffix(string(src), "gen") {
		fmt.Printf("No files found for %s, trying cache\n", src)

		basenames, err := o.getAPIBasenames(o.cacheDir)
		if err != nil {
			return nil, err
		}

		if len(basenames) == 0 {
			fmt.Fprint(o.stderr, `
No generated files found.
Please remove the postgen *.cache directory.`)
			return nil, errors.New("no generated files")
		}

		fmt.Printf("Using cached files in %s\n", o.cacheDir)
		o.conf.GenHandlersDir = Dir(o.cacheDir)

		return basenames, nil
	}

	for _, p := range paths {
		if !strings.HasSuffix(p, "_test.go") {
			out = append(out, path.Base(p))
		}
	}

	return out, nil
}

type Route struct {
	Name        string
	Middlewares dst.Expr
}

/*
extractRoutes returns Route elements indexed by method from a routes node.
*/
func extractRoutes(rr *dst.AssignStmt) map[string]Route {
	out := make(map[string]Route)

	cl, iscl := rr.Rhs[0].(*dst.CompositeLit)
	if !iscl {
		return out
	}

	for _, r := range cl.Elts {
		var (
			opID  string
			route Route
		)

		if cl, clok := r.(*dst.CompositeLit); clok {
			for _, s := range cl.Elts {
				kv, kvok := s.(*dst.KeyValueExpr)
				if !kvok {
					continue
				}

				ident, identok := kv.Key.(*dst.Ident)
				if !identok {
					continue
				}

				switch ident.Name {
				case "Name":
					if ce, isce := kv.Value.(*dst.CallExpr); isce {
						opID = ce.Args[0].(*dst.Ident).Name
						route.Name = opID
					}
					// case "Middlewares":
					// 	route.Middlewares = kv.Value
				}

				out[opID] = route
			}
		}
	}

	return out
}

// inspectRegisterNode extracts the routes slice assignment node and middlewares method body for tag.
func inspectRegisterNode(f dst.Node, tag Tag) *dst.AssignStmt {
	routesNode := &dst.AssignStmt{}

	dst.Inspect(f, func(n dst.Node) bool {
		fn, isFn := n.(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		if rok && identok && ident.Name == string(tag) {
			m := fn.Name.String()

			if as, isas := fn.Body.List[0].(*dst.AssignStmt); isas && m == "Register" {
				routesNode, _ = dst.Clone(as).(*dst.AssignStmt)
			}
		}

		return true
	})

	return routesNode
}

// inspectStruct returns the methods of the handler struct for tag indexed by name.
func inspectStruct(f dst.Node, tag Tag) map[string]Method {
	out := make(map[string]Method)

	dst.Inspect(f, func(n dst.Node) bool {
		fn, isFn := n.(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		if rok && identok && ident.Name == string(tag) {
			out[fn.Name.String()] = Method{
				Name: fn.Name.String(),
				Decl: fn,
			}
		}

		return true
	})

	return out
}
