package zzz

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type Handler struct {
	OperationId string
	Comment     string
	Origin      string
}

func (h Handler) String() string {
	return fmt.Sprintf("%v", h.OperationId)
}

// GenerateHandlers fills in a template with the
// given route handlers to a dest.
func GenerateHandlers(handlers []Handler, dest io.Writer) {
	t := template.Must(template.New("").Parse(`package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
{{range .Handlers}}
// {{.Comment}}
// Origin: {{.Origin}}
func {{.OperationId}}(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
{{end}}`))

	buf := &bytes.Buffer{}

	params := map[string]interface{}{
		"Handlers": handlers,
	}

	if err := t.Execute(buf, params); err != nil {
		panic(err)
	}

	dest.Write(buf.Bytes())
}

// ParseHandlers returns a map with all the handler functions
// found in all files matching pattern.
func ParseHandlers(pattern string) map[string]Handler {
	funcs := make(map[string]Handler)
	fset := token.NewFileSet()

	paths, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for _, p := range paths {
		blob, err := os.ReadFile(p)
		if err != nil {
			panic(err)
		}

		f, err := parser.ParseFile(fset, "", string(blob), parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for _, d := range f.Decls {
			if fn, isFn := d.(*ast.FuncDecl); isFn {
				funcs[fn.Name.Name] = Handler{
					OperationId: fn.Name.Name,
					Comment:     strings.TrimSpace(fn.Doc.Text()),
					Origin:      path.Base(p),
				}
			}
		}
	}

	return funcs
}
