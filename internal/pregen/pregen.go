package pregen

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	// internalformat "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"gopkg.in/yaml.v3"
)

//go:embed templates
var templateFiles embed.FS

type PreGen struct {
	stderr       io.Writer
	specPath     string
	opIDAuthPath string
}

// New returns a new pre-generator.
func New(stderr io.Writer, specPath string, opIDAuthPath string) *PreGen {
	return &PreGen{
		stderr:       stderr,
		specPath:     specPath,
		opIDAuthPath: opIDAuthPath,
	}
}

// analyzeSpec ensures specific rules for codegen are met and extracts necessary data.
func (o *PreGen) analyzeSpec() error {
	var spec yaml.Node

	schemaBlob, err := os.ReadFile(o.specPath)
	if err != nil {
		return fmt.Errorf("error opening schema file: %w", err)
	}

	if err = yaml.Unmarshal([]byte(schemaBlob), &spec); err != nil {
		return fmt.Errorf("error unmarshalling schema: %w", err)
	}

	_, err = yaml.Marshal(&spec)
	if err != nil {
		return fmt.Errorf("error marshalling schema: %w", err)
	}

	return nil
}

// validateSpec validates an OpenAPI 3.0 specification.
func (o *PreGen) validateSpec() error {
	_, err := rest.ReadOpenAPI(o.specPath)
	if err != nil {
		return err
	}

	return nil
}

func (o *PreGen) Generate() error {
	if err := o.validateSpec(); err != nil {
		return fmt.Errorf("validate spec: %w", err)
	}

	if err := o.analyzeSpec(); err != nil {
		return fmt.Errorf("analyze spec: %w", err)
	}

	if err := internal.GenerateConfigTemplate(); err != nil {
		return fmt.Errorf("GenerateConfigTemplate: %w", err)
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
