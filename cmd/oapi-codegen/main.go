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
	"unicode"

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
	Mode             string `yaml:"mode,omitempty"`
	ExcludeRestTypes bool   `yaml:"exclude-rest-types,omitempty"`
	// TestClient defines whether the generated code is a client for testing purposes.
	TestClient             bool `yaml:"test-client,omitempty"`
	SkipDiscriminatorUtils bool `yaml:"skip-discriminator-utils,omitempty"`
	IsRestServerGen        bool `yaml:"is-rest-server-gen,omitempty"`
}

//go:embed oapi-templates
var templates embed.FS

func main() {
	log.SetFlags(0)
	var cfgPath, modelsPkg, typesStr, serverTypesStr, specRestTypesStr string
	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.StringVar(&modelsPkg, "models-pkg", "models", "package containing models")
	flag.StringVar(&typesStr, "types", "", "list of type names to use in place of generated oapi-codegen ones")
	flag.StringVar(&serverTypesStr, "server-types", "", "list of types to use in server generation instead of generated types package")
	flag.StringVar(&specRestTypesStr, "spec-rest-types", "", "list of types that are meant for rest package but defined in spec only (requiring models import)")
	flag.Parse()
	if cfgPath == "" {
		log.Fatal("--config is required")
	}
	if flag.NArg() < 1 {
		log.Fatal("Please specify a path to an OpenAPI 3.0 spec file")
	}
	types := strings.Split(typesStr, ",")
	serverTypes := strings.Split(serverTypesStr, ",")
	specRestTypes := strings.Split(specRestTypesStr, ",")

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
	output, err := generate(
		spec,
		cfg,
		templates,
		modelsPkg,
		types,
		specRestTypes,
		serverTypes,
	)
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

func generate(spec *openapi3.T, config configuration, templates embed.FS, modelsPkg string, types, specRestTypes, serverTypes []string) (string, error) {
	var err error
	config, err = addTemplateOverrides(config, templates)
	if err != nil {
		return "", err
	}
	// include other template functions, if any
	templateFunctions := template.FuncMap{
		"is_sse_endpoint": func(opID string) bool {
			if !config.TestClient {
				return false // for prod client use a dedicated sse client
			}
			for _, p := range spec.Paths.Map() {
				for _, op := range p.Operations() {
					if op.OperationID == opID {
						for _, res := range op.Responses.Map() {
							// as per spec
							if c := res.Value.Content.Get("text/event-stream"); c != nil {
								return true
							}
						}
					}
				}
			}

			return false
		},
		"is_rest_server_gen": func() bool {
			return config.IsRestServerGen
		},
		"skip_discriminator_utils": func() bool {
			return config.SkipDiscriminatorUtils
		},
		"is_test_client": func() bool {
			return config.TestClient
		},
		"exclude_rest_types": func() bool {
			return config.ExcludeRestTypes
		},
		"models_pkg": func() string {
			return modelsPkg + "."
		},
		"is_db_struct": func(t string) bool {
			return strings.HasPrefix(t, "Db") && unicode.IsUpper([]rune(t)[2])
		},
		"is_spec_rest_type": func(t string) bool {
			stName := strings.TrimPrefix(t, "externalRef0.")
			if slices.Contains(specRestTypes, stName) {
				return true
			}

			return false
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
		"is_rest_type": func(s string) bool {
			stName := strings.TrimPrefix(strings.ReplaceAll(s, "ExternalRef0", ""), "externalRef0.")
			for _, typ := range append(types, serverTypes...) {
				if stName == typ {
					return true
				}
			}

			return false
		},
		"gen_mode": func() string {
			return config.Mode
		},
		"rest_type": func(s string) string {
			stName := strings.TrimPrefix(strings.ReplaceAll(s, "ExternalRef0", ""), "externalRef0.")

			// rest type not defined in rest/models.go -> use generated type
			if !config.TestClient && slices.Contains(specRestTypes, stName) {
				return "models." + stName
			}

			// to allow for easier tests where we dont have to populate field by field
			if config.TestClient && slices.Contains(types, stName) {
				return "rest." + stName
			}

			return stName
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
