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

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/deepmap/oapi-codegen/v2/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputFile is the filename to output.
	OutputFile string `yaml:"output,omitempty"`
}

//go:embed oapi-templates
var templates embed.FS

func main() {
	log.SetFlags(0)
	var cfgPath, modelsPkg, structsStr string
	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.StringVar(&modelsPkg, "models-pkg", "models", "package containing models")
	flag.StringVar(&structsStr, "structs", "structs", "list of struct names to use in place of generated oapi-codegen ones")
	flag.Parse()
	if cfgPath == "" {
		log.Fatal("--config is required")
	}
	if flag.NArg() < 1 {
		log.Fatal("Please specify a path to an OpenAPI 3.0 spec file")
	}
	structs := strings.Split(structsStr, ",")

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
	output, err := generate(spec, cfg.Configuration, templates, modelsPkg, structs)
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

func generate(spec *openapi3.T, config codegen.Configuration, templates embed.FS, modelsPkg string, structs []string) (string, error) {
	var err error
	config, err = addTemplateOverrides(config, templates)
	if err != nil {
		return "", err
	}
	// include other template functions, if any
	templateFunctions := template.FuncMap{
		"models_pkg": func() string {
			return modelsPkg + "."
		},
		"is_rest_struct": func(st string) bool {
			return slices.Contains(structs, st)
		},
	}
	for k, v := range templateFunctions {
		codegen.TemplateFunctions[k] = v
	}

	return codegen.Generate(spec, config)
}

func addTemplateOverrides(config codegen.Configuration, templates embed.FS) (codegen.Configuration, error) {
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
	config.OutputOptions.UserTemplates = overrides

	return config, err
}
