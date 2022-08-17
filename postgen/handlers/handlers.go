package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}

	return false
}

type Conf struct {
	// CurrentHandlersDir is the directory with edited generated files.
	CurrentHandlersDir string
	// GenHandlersDir is the directory with raw generated files for a given spec.
	GenHandlersDir string
	// OutHandlersDir is the directory to store merged files,
	// which may differ from CurrentHandlersDir.
	OutHandlersDir string
	// OutServicesDir is the directory to store new default services.
	OutServicesDir string
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

// generateService fills in a template with a default service struct to a dest.
func generateService(tag string, dest io.Writer) {
	t := template.Must(template.New("").Parse(`package services

type {{.Tag}} struct {
}
`))

	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Tag": strings.Title(tag),
	}

	if err := t.Execute(buf, params); err != nil {
		panic(err)
	}

	dest.Write(buf.Bytes())
}

// analyzeHandlers returns all necessary merging information about handlers, indexed
// by directory and tag.
func analyzeHandlers(conf Conf, basenames []string) map[string]map[string]HandlerFile {
	handlers := make(map[string]map[string]HandlerFile)

	dirs := []string{conf.GenHandlersDir, conf.CurrentHandlersDir}
	for _, dir := range dirs {
		handlers[dir] = make(map[string]HandlerFile)

		for _, basename := range basenames {
			file := path.Join(dir, basename)

			blob, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}

			f, err := decorator.Parse(blob)
			if err != nil {
				panic(err)
			}

			reg := regexp.MustCompile("api_(.*).go")
			tag := strings.Title(reg.FindStringSubmatch(basename)[1])

			mm := inspectStruct(f, tag)
			rr := inspectNodes(f, tag)
			routes := extractRoutes(rr)
			hf := HandlerFile{
				F:          f,
				Methods:    mm,
				RoutesNode: rr,
				Routes:     routes,
			}

			handlers[dir][tag] = hf
		}
	}

	findClashingMethodNames(basenames, handlers, conf)

	return handlers
}

// findClashingMethodNames ensures no previous methods that are not
// handlers conflict with a newly generated operation id.
func findClashingMethodNames(basenames []string, handlers map[string]map[string]HandlerFile, conf Conf) {
	clashes := []string{}

	for _, basename := range basenames {
		reg := regexp.MustCompile("api_(.*).go")
		tag := strings.Title(reg.FindStringSubmatch(basename)[1])

		for opId := range handlers[conf.GenHandlersDir][tag].Routes {
			fmt.Printf("[%s] opId: %s\n", tag, opId)
			_, rok := handlers[conf.CurrentHandlersDir][tag].Routes[opId]
			_, mok := handlers[conf.CurrentHandlersDir][tag].Methods[opId]

			if !rok && mok {
				clashes = append(clashes, tag+"->"+opId)
			}
		}
	}

	if len(clashes) > 0 {
		fmt.Fprintf(os.Stderr, `
Error: conflicting method names
%s
Please rename either the affected method or operation id.
`, clashes)
		os.Exit(1)
	}
}

// getCommonBasenames returns api filename (tag) intersections in current and raw gen dirs,
// and copies the remaining files to the out dir without further analysis.
func getCommonBasenames(conf Conf) (out []string) {
	k := 0
	currentBasenames := getApiBasenames(conf.CurrentHandlersDir)
	genBasenames := getApiBasenames(conf.GenHandlersDir)

	os.MkdirAll(conf.OutHandlersDir, 0777)

	for _, genBasename := range genBasenames {
		if contains(currentBasenames, genBasename) {
			genBasenames[k] = genBasename
			k++

			continue
		}

		genBlob, err := os.ReadFile(path.Join(conf.GenHandlersDir, genBasename))
		if err != nil {
			panic(err)
		}

		os.WriteFile(path.Join(conf.OutHandlersDir, genBasename), genBlob, 0666)
	}

	genBasenames = genBasenames[:k]

	for _, currentBasename := range currentBasenames {
		if contains(genBasenames, currentBasename) {
			out = append(out, currentBasename)

			continue
		}

		currentBlob, err := os.ReadFile(path.Join(conf.CurrentHandlersDir, currentBasename))
		if err != nil {
			panic(err)
		}

		os.WriteFile(path.Join(conf.OutHandlersDir, currentBasename), currentBlob, 0666)
	}

	return out
}

// replaceNodes replaces handler file nodes accordingly.
func replaceNodes(f *dst.File, genHf, currentHf HandlerFile, tag string, opId string) {
	dstutil.Apply(f, nil, func(c *dstutil.Cursor) bool {
		fn, isFn := c.Parent().(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		m := fn.Name.String()

		if rok && identok && ident.Name == tag && m == "Register" {
			fn.Body.List[0] = genHf.RoutesNode
		}
		if rok && identok && ident.Name == tag && m == "middlewares" {
			fn.Body = currentHf.Methods["middlewares"].Decl.Body
		}

		return true
	})
}

func generateMergedFiles(handlers map[string]map[string]HandlerFile, conf Conf) {
	for tag, currentHF := range handlers[conf.CurrentHandlersDir] {
		//nolint: forcetypeassert
		outF := dst.Clone(currentHF.F).(*dst.File)
		gh := handlers[conf.GenHandlersDir][tag]
		fmt.Println(format.Underlined + format.Blue + tag + format.Off)

		for _, cv := range gh.Routes {
			fmt.Printf("gen r.Name: %s\n", cv.Name)
		}

		// get generated operation ids as list
		gkk := make([]string, len(handlers[conf.GenHandlersDir][tag].Methods))
		i := 0

		for gk := range handlers[conf.GenHandlersDir][tag].Methods {
			gkk[i] = gk
			i++
		}

		sort.Slice(gkk, func(i, j int) bool {
			return gkk[i] < gkk[j]
		})

		genHF := handlers[conf.GenHandlersDir][tag]

		for _, gk := range gkk {
			if _, ok := currentHF.Methods[gk]; !ok {
				fmt.Printf("method %s not in current %s - appending generated method.\n", gk, tag)
				outF.Decls = append(outF.Decls, handlers[conf.GenHandlersDir][tag].Methods[gk].Decl)

				continue
			}

			replaceNodes(outF, genHF, currentHF, tag, gk)
		}

		buf := &bytes.Buffer{}

		f, err := os.Create(path.Join(conf.OutHandlersDir, "api_"+strings.ToLower(tag)+".go"))
		if err != nil {
			panic(err)
		}

		if err := decorator.Fprint(buf, outF); err != nil {
			panic(err)
		}

		f.Write(buf.Bytes())
	}

	// generate default service if not exists
	for tag := range handlers[conf.GenHandlersDir] {
		s := path.Join(conf.OutServicesDir, strings.ToLower(tag)+".go")
		if _, err := os.Stat(s); errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(s)
			if err != nil {
				panic(err)
			}

			generateService(tag, f)
		}
	}
}

type Method struct {
	// Name is the method identifier.
	Name string
	// Decl is the method declaration node.
	Decl *dst.FuncDecl
}

type HandlerFile struct {
	// F is the node of the file.
	F *dst.File
	// Methods represents all methods in the generated struct indexed by method name.
	Methods map[string]Method
	// RoutesNode represents the complete routes assignment node.
	RoutesNode *dst.AssignStmt
	// Routes represents convenient extracted fields from route elements
	// indexed by operation id.
	Routes map[string]Route
}

func getApiBasenames(src string) []string {
	out := []string{}
	// glob uses https://pkg.go.dev/path#Match patterns
	paths, err := filepath.Glob(path.Join(src, "api_*.go"))
	if err != nil {
		panic(err)
	}

	for _, p := range paths {
		if !strings.HasSuffix(p, "_test.go") {
			out = append(out, path.Base(p))
		}
	}

	return out
}

func applyFunc(c *dstutil.Cursor) bool {
	node := c.Node()

	switch n := node.(type) {
	case (*dst.FuncDecl):
		fmt.Printf("\n---------------\n")
		// dst.Print(n)
		dst.Print(n.Body)
	}

	return true
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
			opId  string
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
					if lit, islit := kv.Value.(*dst.BasicLit); islit {
						opId, _ = strconv.Unquote(lit.Value)
						route.Name = opId
					}
					// case "Middlewares":
					// 	route.Middlewares = kv.Value
				}

				out[opId] = route
			}
		}
	}

	return out
}

// inspectNodes extracts the routes slice assignment node and middlewares method body for tag.
func inspectNodes(f dst.Node, tag string) *dst.AssignStmt {
	routesNode := &dst.AssignStmt{}

	dst.Inspect(f, func(n dst.Node) bool {
		fn, isFn := n.(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		if rok && identok && ident.Name == tag {
			m := fn.Name.String()

			if as, isas := fn.Body.List[0].(*dst.AssignStmt); isas && m == "Register" {
				routesNode, _ = dst.Clone(as).(*dst.AssignStmt)
			}
		}

		return true
	})

	return routesNode
}

func inspectStruct(f dst.Node, tag string) map[string]Method {
	out := make(map[string]Method)

	dst.Inspect(f, func(n dst.Node) bool {
		fn, isFn := n.(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		if rok && identok && ident.Name == tag {
			out[fn.Name.String()] = Method{
				Name: fn.Name.String(),
				Decl: fn,
			}
		}

		return true
	})

	return out
}
