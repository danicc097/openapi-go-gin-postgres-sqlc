package postgen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type Handler struct {
	OperationId string
	Comment     string
}

func (h Handler) String() string {
	return fmt.Sprintf("%v", h.OperationId)
}

// GenerateHandlers generates fills in a template with the
// given route handlers to a dest.
func GenerateHandlers(handlers []Handler, dest io.Writer) {
	t := template.Must(template.New("").Parse(`package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
{{range .Handlers}}
{{.Comment}}
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

// ParseHandlers returns a map with all the functions found in
// all files matching pattern.
func ParseHandlers(pattern string) map[string]Handler {
	funcs := make(map[string]Handler)
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		f, err := parser.ParseFile(fset, "", string(content), 0)
		if err != nil {
			panic(err)
		}

		// TODO does not return function comments (empty str)
		for _, d := range f.Decls {
			if fn, isFn := d.(*ast.FuncDecl); isFn {
				funcs[fn.Name.Name] = Handler{OperationId: fn.Name.Name, Comment: fn.Doc.Text()}
			}
		}
	}
	return funcs
}
