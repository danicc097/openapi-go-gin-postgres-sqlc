package pregen

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type PreGen struct {
	stderr io.Writer
	spec   string
}

// New returns a new pre-generator.
func New(stderr io.Writer, spec string) *PreGen {
	return &PreGen{
		stderr: stderr,
		spec:   spec,
	}
}

// analyzeSpec ensures specific rules for codegen are met and extracts necessary data.
func (o *PreGen) analyzeSpec() error {
	var spec yaml.Node

	schemaBlob, err := os.ReadFile(o.spec)
	if err != nil {
		return fmt.Errorf("error opening schema file: %w", err)
	}

	if err = yaml.Unmarshal([]byte(schemaBlob), &spec); err != nil {
		return fmt.Errorf("error unmarshalling schema: %w", err)
	}

	output, err := yaml.Marshal(&spec)
	if err != nil {
		return fmt.Errorf("error marshalling schema: %w", err)
	}
	fmt.Println(string(output))

	return nil
}

func (o *PreGen) Generate() error {
	err := o.analyzeSpec()
	if err != nil {
		return fmt.Errorf("invalid spec: %w", err)
	}

	return nil
}
