package pregen

import (
	"fmt"
	"io"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"gopkg.in/yaml.v3"
)

type PreGen struct {
	stderr   io.Writer
	specPath string
}

// New returns a new pre-generator.
func New(stderr io.Writer, specPath string) *PreGen {
	return &PreGen{
		stderr:   stderr,
		specPath: specPath,
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

	return nil
}
