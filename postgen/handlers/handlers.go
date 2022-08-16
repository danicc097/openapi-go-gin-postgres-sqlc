package main

import (
	"bytes"
	"fmt"
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
	// CurrentDir is the directory with edited generated files.
	CurrentDir string
	// GenDir is the directory with raw generated files for a given spec.
	GenDir string
	// OutDir is the directory to store merged files.
	OutDir string
}

/*
	Read:
	https://go.dev/src/go/ast
	https://pkg.go.dev/go/ast
	https://github.com/dave/dst/blob/master/dstutil/rewrite_test.go
	https://developers.mattermost.com/blog/instrumenting-go-code-via-ast-2/

	Rationale:
	We dont want users to manually add handlers and routes to handlers/*
	If we let them, we wouldnt know at plain sight what was in the spec and what wasnt
	and parsing will become a bigger mess.
	Users can still add new methods to the struct, but the routes slice in
	Register will be overridden, only retaining certain properties, currently Middlewares.
	If we need a new route that cant be defined in the spec, e.g. fileserver,
	we purposely want that out of the generated handler struct,
	so its clear that its outside the spec.
	It can still remain in handlers/* as long as its not api_*(!_test).go, e.g. fileserver.go
	and it can still follow the same handlers struct pattern for all we care, it wont be touched.
	IMPORTANT: if a method already exists in current but has no routes item (meaning
	its probably some handler helper method created afterwards) then panic and alert
	the user to rename. it shouldve been unexported or a function in the first place anyway.

*/
func main() {
	var (
		baseDir = "testdata"
		conf    = Conf{
			CurrentDir: path.Join(baseDir, "merge_changes/current"),
			GenDir:     path.Join(baseDir, "merge_changes/internal/gen"),
			OutDir:     path.Join(baseDir, "merge_changes/got")}
	)

	// TODO add a method in current that is not a handler and conflicts with a new method from gen -> should panic and prompt to rename.

	// TODO refactor for clearness to https://stackoverflow.com/questions/52120488/what-is-the-most-efficient-way-to-get-the-intersection-and-exclusions-from-two-a
	cb := getCommonBasenames(conf)
	handlers := analyzeHandlers(conf, cb)

	generateMergedFiles(handlers, conf)
}

// analyzeHandlers returns all necessary merging information about handlers, indexed
// by directory and tag.
func analyzeHandlers(conf Conf, basenames []string) map[string]map[string]HandlerFile {
	handlers := make(map[string]map[string]HandlerFile)

	dirs := []string{conf.GenDir, conf.CurrentDir}
	for _, dir := range dirs {
		handlers[dir] = make(map[string]HandlerFile)

		for _, basename := range basenames {
			file := path.Join(dir, basename)

			c, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}

			f, err := decorator.Parse(c)
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

	/* TODO:
	for opId, _ in handlers[conf.GenDir][tag].Routes:
		if handlers[conf.CurrentDir][tag].Routes[opId] does not exist
		and handlers[conf.CurrentDir][tag].Methods[opId] exists
		then we have a clash in current method and should be renamed (panic)
	*/

	return handlers
}

// getCommonBasenames returns api filename (tag) intersections in current and raw gen dirs,
// and copies the remaining files to the out dir without further analysis.
func getCommonBasenames(conf Conf) (out []string) {
	k := 0
	currentBasenames := getApiBasenames(conf.CurrentDir)
	genBasenames := getApiBasenames(conf.GenDir)

	os.MkdirAll(conf.OutDir, 0777)

	for _, genBasename := range genBasenames {
		if contains(currentBasenames, genBasename) {
			genBasenames[k] = genBasename
			k++

			continue
		}

		genContent, err := os.ReadFile(path.Join(conf.GenDir, genBasename))
		if err != nil {
			panic(err)
		}

		os.WriteFile(path.Join(conf.OutDir, genBasename), genContent, 0666)
	}

	genBasenames = genBasenames[:k]

	for _, currentBasename := range currentBasenames {
		if contains(genBasenames, currentBasename) {
			out = append(out, currentBasename)

			continue
		}

		currentContent, err := os.ReadFile(path.Join(conf.CurrentDir, currentBasename))
		if err != nil {
			panic(err)
		}

		os.WriteFile(path.Join(conf.OutDir, currentBasename), currentContent, 0666)
	}

	return out
}

// replaceRoute replaces a routes slice element node in Register() for an operation id.
func replaceRoute(f *dst.File, hf, hfUpdate HandlerFile, tag string, opId string) {
	dstutil.Apply(f, nil, func(c *dstutil.Cursor) bool {
		fn, isFn := c.Parent().(*dst.FuncDecl)
		if !isFn || fn.Recv == nil || len(fn.Recv.List) != 1 {
			return true // keep traversing children
		}
		r, rok := fn.Recv.List[0].Type.(*dst.StarExpr)
		ident, identok := r.X.(*dst.Ident)
		m := fn.Name.String()

		if rok && identok && ident.Name == tag && m == "Register" {
			fn.Body.List[0] = hf.RoutesNode
		}
		if rok && identok && ident.Name == tag && m == "middlewares" {
			fn.Body = hfUpdate.Methods["middlewares"].Decl.Body
		}

		return true
	})
}

func generateMergedFiles(handlers map[string]map[string]HandlerFile, conf Conf) {
	for tag, currentHF := range handlers[conf.CurrentDir] {
		//nolint: forcetypeassert
		outF := dst.Clone(currentHF.F).(*dst.File)
		gh := handlers[conf.GenDir][tag]
		fmt.Println(format.Underlined + format.Blue + tag + format.Off)

		for _, cv := range gh.Routes {
			fmt.Printf("gen r.Name: %s\n", cv.Name)
		}

		// get generated operation ids as list
		gkk := make([]string, len(handlers[conf.GenDir][tag].Routes))
		i := 0

		for gk := range handlers[conf.GenDir][tag].Routes {
			gkk[i] = gk
			i++
		}

		sort.Slice(gkk, func(i, j int) bool {
			return gkk[i] < gkk[j]
		})

		genHF := handlers[conf.GenDir][tag]

		for _, gk := range gkk {
			if _, ok := currentHF.Routes[gk]; !ok {
				fmt.Printf("op id %s not in current->%s, adding method.\n", gk, tag)
				outF.Decls = append(outF.Decls, handlers[conf.GenDir][tag].Methods[gk].Decl)

				continue
			}

			replaceRoute(outF, genHF, currentHF, tag, gk)
		}

		buf := &bytes.Buffer{}

		f, err := os.Create(path.Join(conf.OutDir, "api_"+strings.ToLower(tag)+".go"))
		if err != nil {
			panic(err)
		}

		if err := decorator.Fprint(buf, outF); err != nil {
			panic(err)
		}

		f.Write(buf.Bytes())
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
