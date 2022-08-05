package postgen

import (
	"bytes"
	"io"
	"text/template"
)

type Handler struct {
	OperationId string
	Comment     string
}

func GenNotImpHandlers(handlers []Handler, dest io.Writer) {
	t := template.Must(template.New("").Parse(`package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
{{range .Handlers}}
// {{.Comment}}
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
