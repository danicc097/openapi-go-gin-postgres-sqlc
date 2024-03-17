package main

// https://github.com/deepmap/oapi-codegen/pull/707

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"os"
	"path"
	"slices"
	"strings"
	"text/template"

	"github.com/deepmap/oapi-codegen/v2/pkg/codegen"
	"github.com/deepmap/oapi-codegen/v2/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputFile is the filename to output.
	OutputFile       string `yaml:"output,omitempty"`
	ExcludeRestTypes bool   `yaml:"exclude-rest-types,omitempty"`
	// TestClient defines whether the generated code is a client for testing purposes.
	TestClient bool `yaml:"test-client,omitempty"`
}

//go:embed oapi-templates
var templates embed.FS

func main() {
	log.SetFlags(0)
	var cfgPath, modelsPkg, typesStr, serverTypesStr string
	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.StringVar(&modelsPkg, "models-pkg", "models", "package containing models")
	flag.StringVar(&typesStr, "types", "", "list of type names to use in place of generated oapi-codegen ones")
	flag.StringVar(&serverTypesStr, "server-types", "", "list of types to use in server generation instead of generated types package")
	flag.Parse()
	if cfgPath == "" {
		log.Fatal("--config is required")
	}
	if flag.NArg() < 1 {
		log.Fatal("Please specify a path to an OpenAPI 3.0 spec file")
	}
	types := strings.Split(typesStr, ",")
	serverTypes := strings.Split(serverTypesStr, ",")

	// loading specification
	input := flag.Arg(0)
	spec, err := util.LoadSwagger(input)
	if err != nil {
		log.Fatalf("error loading openapi specification: %v", err)
	}

	// will fail on separated yamls
	// err = spec.Validate(context.Background())
	// if err != nil {
	// 	log.Fatalf("error validating openapi specification: %v", err)
	// }

	// loading configuration
	cfgdata, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	var cfg configuration
	err = yaml.Unmarshal(cfgdata, &cfg)
	if err != nil {
		log.Fatalf("error unmarshaling config %v", err)
	}

	// generating output
	output, err := generate(spec, cfg, templates, modelsPkg, types, serverTypes)
	if err != nil {
		log.Fatalf("error generating code: %v", err)
	}

	// writing output to file
	outFile, err := os.Create(cfg.OutputFile)
	if err != nil {
		log.Fatalf("error creating output file: %v", err)
	}
	_, err = outFile.WriteString(output)
	if err != nil {
		log.Fatalf("error writing output file: %v", err)
	}
	outFile.Close()
}

func generate(spec *openapi3.T, config configuration, templates embed.FS, modelsPkg string, types, serverTypes []string) (string, error) {
	var err error
	config, err = addTemplateOverrides(config, templates)
	if err != nil {
		return "", err
	}
	// include other template functions, if any
	templateFunctions := template.FuncMap{
		"is_test_client": func() bool {
			return config.TestClient
		},
		"exclude_rest_types": func() bool {
			return config.ExcludeRestTypes
		},
		"models_pkg": func() string {
			return modelsPkg + "."
		},
		"should_exclude_type": func(t string) bool {
			stName := strings.TrimPrefix(t, "externalRef0.")
			if slices.Contains(serverTypes, stName) {
				return false
			}
			for _, typ := range types {
				if stName == typ {
					return true
				}
			}

			return false
		},
		"is_rest_type": func(t string) bool {
			stName := strings.TrimPrefix(t, "externalRef0.")
			for _, typ := range append(types, serverTypes...) {
				if stName == typ {
					return true
				}
			}

			return false
		},
		"rest_type": func(s string) string {
			return strings.TrimPrefix(strings.ReplaceAll(s, "ExternalRef0", ""), "externalRef0.")
		},
		"camel": strcase.ToCamel,
	}
	for k, v := range templateFunctions {
		codegen.TemplateFunctions[k] = v
	}

	return codegen.Generate(spec, config.Configuration)
}

func addTemplateOverrides(config configuration, templates embed.FS) (configuration, error) {
	overrides := config.OutputOptions.UserTemplates
	if overrides == nil {
		overrides = make(map[string]string)
	}
	err := fs.WalkDir(templates, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if err != nil {
				return err
			}
			f, err := templates.ReadFile(p)
			if err != nil {
				return err
			}
			name := strings.TrimSuffix(p, path.Ext(p)) + ".tmpl"
			name = strings.Join(strings.Split(name, "/")[1:], "/")
			overrides[name] = string(f)
		}

		return nil
	})
	config.Configuration.OutputOptions.UserTemplates = overrides

	return config, err
}
