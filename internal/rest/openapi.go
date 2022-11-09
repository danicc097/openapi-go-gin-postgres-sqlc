package rest

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

// ReadOpenAPI parses and validates an OpenAPI by filename and returns it.
func ReadOpenAPI(path string) (*openapi3.T, error) {
	schemaBlob, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("openapi spec: %w", err)
	}
	sl := openapi3.NewLoader()

	openapi, err := sl.LoadFromData(schemaBlob)
	if err != nil {
		return nil, fmt.Errorf("openapi spec: %w", err)
	}

	if err = openapi.Validate(sl.Context); err != nil {
		return nil, fmt.Errorf("openapi validation: %w", err)
	}

	return openapi, nil
}
