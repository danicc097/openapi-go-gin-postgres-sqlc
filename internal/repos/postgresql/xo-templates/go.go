//go:build xotpl

package xotemplates

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/danicc097/xo/v5/loader"
	xo "github.com/danicc097/xo/v5/types"
	"github.com/kenshaw/inflector"
	"github.com/kenshaw/snaker"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

const (
	Off = "\033[0m"

	Red    = "\033[31m"
	R      = Red
	Green  = "\033[32m"
	G      = Green
	Yellow = "\033[33m"
	Y      = Yellow
	Blue   = "\033[34m"
	B      = Blue
	Purple = "\033[35m"
	P      = Purple
	Cyan   = "\033[36m"
	C      = Cyan

	Bold       = "\033[1m"
	Dim        = "\033[2m"
	Italic     = "\033[3m"
	Underlined = "\033[4m"
	Blinking   = "\033[5m"
	Reverse    = "\033[6m"
	Invisible  = "\033[7m"
)

// TODO configurable
var excludedIndexTypes = []string{"gin_trgm_ops"}

type cardinality string

const (
	M2M cardinality = "M2M"
	M2O cardinality = "M2O"
	O2O cardinality = "O2O"
)

type annotation = string

// table column custom properties via SQL column comments.
const (
	annotationJoinOperator       = " && "
	annotationAssignmentOperator = ":"

	cardinalityAnnot annotation = `"cardinality"`
	// custom properties for code generation
	propertiesAnnot annotation = `"properties"`
	// literal Go type to override with
	typeAnnot annotation = `"type"`
	// literal Go struct tags to be appended
	tagsAnnot annotation = `"tags"`

	propertiesJoinOperator = ","
	// propertyRefsIgnore generates a field whose constraints are ignored by referenced table,
	// ie no joins will be generated.
	propertyRefsIgnore = "refs-ignore"
	// target column will generate the same M2O and M2M join fields the ref column has
	propertyShareRefConstraints = "share-ref-constraints"
	// propertyJSONPrivate sets a json:"-" tag.
	propertyJSONPrivate = "private"
	// propertyOpenAPINotRequired marks schema field as not required
	propertyOpenAPINotRequired = "not-required"
	// propertyOpenAPIHidden makes schema generator skip over field in db (Create|Update)Params
	// Useful when the field doesn't need to be set in body or we want to replace it with a more user-friendly field
	// that gets converted to the db field internally
	propertyOpenAPIHidden = "hidden"

	// example: "properties":private,another-property && "type":ProjectName && "tags":pattern: ^[\.a-zA-Z0-9_-]+$
)

// to not have to analyze everything for convertConstraints
var cardinalityRE = regexp.MustCompile(string(cardinalityAnnot) + annotationAssignmentOperator + "([A-Za-z0-9_-]*)")

func columnCommentToAnnotations(c cardinality) bool {
	if c != "" && c != M2M && c != M2O && c != O2O {
		return false
	}
	return true
}

func validCardinality(c cardinality) bool {
	if c != "" && c != M2M && c != M2O && c != O2O {
		return false
	}
	return true
}

func formatJSON(obj any) string {
	bytes, _ := json.MarshalIndent(obj, "  ", "  ")
	return string(bytes)
}

var ErrNoSingle = errors.New("in query exec mode, the --single or -S must be provided")

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func removeEmptyStrings(arr []string) []string {
	result := make([]string, 0)
	for _, str := range arr {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func toAcronym(input string) string {
	acronym := ""

	// Check if input is in snake_case or camelCase/PascalCase
	if strings.Contains(input, "_") {
		// Snake_case input
		words := strings.Split(input, "_")
		for _, word := range words {
			if len(word) > 0 {
				acronym += string(word[0])
			}
		}
	} else {
		// requires CamelCase or PascalCase input
		for i, char := range input {
			if unicode.IsUpper(char) {
				// Check if it's not the first character and the previous character is not uppercase
				if i > 0 && !unicode.IsUpper(rune(input[i-1])) {
					acronym += string(char)
				}
			}
		}
	}

	if acronym == "" {
		acronym = strings.ToUpper(string(input[len(input)-1]))
	}

	return strings.ToUpper(acronym)
}

func uniqueSort(slice []string) []string {
	uniqueMap := make(map[string]bool)

	for _, item := range slice {
		uniqueMap[item] = true
	}

	uniqueSlice := make([]string, 0, len(uniqueMap))
	for item := range uniqueMap {
		uniqueSlice = append(uniqueSlice, item)
	}

	sort.Strings(uniqueSlice)

	return uniqueSlice
}

// Init registers the template.
func Init(ctx context.Context, f func(xo.TemplateType)) error {
	knownTypes := map[string]bool{
		"bool":            true,
		"string":          true,
		"byte":            true,
		"rune":            true,
		"int":             true,
		"int16":           true,
		"int32":           true,
		"int64":           true,
		"uint":            true,
		"uint8":           true,
		"uint16":          true,
		"uint32":          true,
		"uint64":          true,
		"float32":         true,
		"float64":         true,
		"[]bool":          true,
		"[][]byte":        true,
		"[]float64":       true,
		"[]float32":       true,
		"[]int64":         true,
		"[]int32":         true,
		"[]string":        true,
		"[]byte":          true,
		"pq.BoolArray":    true,
		"pq.ByteArray":    true,
		"pq.Float64Array": true,
		"pq.Float32Array": true,
		"pq.Int64Array":   true,
		"pq.Int32Array":   true,
		"pq.StringArray":  true,
		"pq.GenericArray": true,
	}
	shorts := map[string]string{
		"bool":            "b",
		"string":          "s",
		"byte":            "b",
		"rune":            "r",
		"int":             "i",
		"int16":           "i",
		"int32":           "i",
		"int64":           "i",
		"uint":            "u",
		"uint8":           "u",
		"uint16":          "u",
		"uint32":          "u",
		"uint64":          "u",
		"float32":         "f",
		"float64":         "f",
		"[]bool":          "a",
		"[][]byte":        "a",
		"[]float64":       "a",
		"[]float32":       "a",
		"[]int64":         "a",
		"[]int32":         "a",
		"[]string":        "a",
		"[]byte":          "a",
		"pq.BoolArray":    "a",
		"pq.ByteArray":    "a",
		"pq.Float64Array": "a",
		"pq.Float32Array": "a",
		"pq.Int64Array":   "a",
		"pq.Int32Array":   "a",
		"pq.StringArray":  "a",
		"pq.GenericArray": "a",
	}
	f(xo.TemplateType{
		Modes: []string{"query", "schema"},
		Flags: []xo.Flag{
			{
				ContextKey: AppendKey,
				Type:       "bool",
				Desc:       "enable append mode",
				Short:      "a",
				Aliases:    []string{"append"},
			},
			{
				ContextKey: NotFirstKey,
				Type:       "bool",
				Desc:       "disable package file (ie. not first generated file)",
				Short:      "2",
				Default:    "false",
			},
			{
				ContextKey: Int32Key,
				Type:       "string",
				Desc:       "int32 type",
				Default:    "int",
			},
			{
				ContextKey: Uint32Key,
				Type:       "string",
				Desc:       "uint32 type",
				Default:    "uint",
			},
			{
				ContextKey: ArrayModeKey,
				Type:       "string",
				Desc:       "array type mode (postgres only)",
				Enums:      []string{"stdlib", "pq"},
			},
			{
				ContextKey: PkgKey,
				Type:       "string",
				Desc:       "package name",
			},
			{
				ContextKey: TagKey,
				Type:       "[]string",
				Desc:       "build tags",
			},
			{
				ContextKey: ImportKey,
				Type:       "[]string",
				Desc:       "package imports",
			},
			{
				ContextKey: UUIDKey,
				Type:       "string",
				Desc:       "uuid type package",
				Default:    "github.com/google/uuid",
			},
			{
				ContextKey: CustomKey,
				Type:       "string",
				Desc:       "package name for custom types",
			},
			{
				ContextKey: ConflictKey,
				Type:       "string",
				Desc:       "name conflict suffix",
				Default:    "Val",
			},
			{
				ContextKey: InitialismKey,
				Type:       "[]string",
				Desc:       "add initialism (e.g. ID, API, URI, ...)",
			},
			{
				ContextKey: EscKey,
				Type:       "[]string",
				Desc:       "escape fields",
				Default:    "none",
				Enums:      []string{"none", "schema", "table", "column", "all"},
			},
			{
				ContextKey: FieldTagKey,
				Type:       "string",
				Desc:       "field tag",
				Short:      "g",
				Default: `json:"{{ if or (.ignoreJSON) (.hidden) }}-{{ else }}{{ camel .field.GoName }}{{end}}"
{{- if not .skipExtraTags }} db:"{{ .field.SQLName -}}"
{{- end }}
{{- if and (.required) (not .hidden)}} required:"true"
{{- end }}
{{- if and (not .nullable) (not .hidden) }} nullable:"false"
{{- end }}
{{- if and (.field.OpenAPISchema) (not .hidden) }} ref:"#/components/schemas/{{ .field.OpenAPISchema }}"
{{- end }}
{{- .field.ExtraTags }}`,
			},
			{
				ContextKey: PublicFieldTagKey,
				Type:       "string",
				Desc:       "public field tag",
				Default:    `json:"{{ camel .GoName }}" required:"true"`,
			},
			{
				ContextKey: PrivateFieldTagKey,
				Type:       "string",
				Desc:       "private field tag",
				Default:    `json:"-"`,
			},
			{
				ContextKey: ContextKey,
				Type:       "string",
				Desc:       "context mode",
				Default:    "only",
				Enums:      []string{"disable", "both", "only"},
			},
			{
				ContextKey: InjectKey,
				Type:       "string",
				Desc:       "insert code into generated file headers",
				Default:    "",
			},
			{
				ContextKey: InjectFileKey,
				Type:       "string",
				Desc:       "insert code into generated file headers from a file",
				Default:    "",
			},
			{
				ContextKey: LegacyKey,
				Type:       "bool",
				Desc:       "enables legacy v1 template funcs",
				Default:    "false",
			},
		},
		Funcs: func(ctx context.Context, _ string) (template.FuncMap, error) {
			funcs, err := NewFuncs(ctx)
			if err != nil {
				return nil, err
			}
			if Legacy(ctx) {
				addLegacyFuncs(ctx, funcs)
			}
			return funcs, nil
		},
		NewContext: func(ctx context.Context, _ string) context.Context {
			ctx = context.WithValue(ctx, KnownTypesKey, knownTypes)
			ctx = context.WithValue(ctx, ShortsKey, shorts)
			return ctx
		},
		Order: func(ctx context.Context, mode string) []string {
			base := []string{"header", "db", "extra"}
			switch mode {
			case "query":
				return append(base, "typedef", "query")
			case "schema":
				return append(base, "enum", "proc", "typedef", "query", "index", "foreignkey")
			}
			return nil
		},
		Pre: func(ctx context.Context, mode string, set *xo.Set, out fs.FS, emit func(xo.Template)) error {
			if err := addInitialisms(ctx); err != nil {
				return err
			}
			files, err := fileNames(ctx, mode, set)
			if err != nil {
				return err
			}
			// If -2 is provided, skip package template outputs as requested.
			// If -a is provided, skip to avoid duplicating the template.
			_, _, schemaOpt := xo.DriverDbSchema(ctx) // cli arg
			// TODO: we should merge public and custom schemas generation in one with comma separated list.
			// so we can generate f.entities properly
			if !NotFirst(ctx) && !Append(ctx) {
				tables := make(Tables)
				for _, schema := range set.Schemas {
					tcc := append(schema.Tables, schema.Views...)
					tcc = append(tcc, schema.MatViews...)
					for _, tc := range tcc {
						table, err := convertTable(ctx, tc)
						if err != nil {
							return err
						}
						tables[table.SQLName] = table
					}
				}
				if schemaOpt == "public" {
					emit(xo.Template{
						Partial: "db",
						Dest:    "db.xo.go",
					})
					// If --single is provided, don't generate header for db.xo.go.
					if xo.Single(ctx) == "" {
						files["db.xo.go"] = true
					}

					emit(xo.Template{
						Partial: "extra",
						Dest:    "extra.xo.go",
						Data: struct {
							Tables any
							Schema string
						}{Tables: tables, Schema: schemaOpt},
					})
					// If --single is provided, don't generate header for db.xo.go.
					if xo.Single(ctx) == "" {
						files["extra.xo.go"] = true
					}
				}

				if schemaOpt != "public" {
					emit(xo.Template{
						Partial: "extra",
						Dest:    schemaOpt + "_" + "extra.xo.go",
						Data: struct {
							Tables any
							Schema string
						}{Tables: tables, Schema: schemaOpt},
					})
				}

				// If --single is provided, don't generate header for db.xo.go.
				if xo.Single(ctx) == "" {
					files[schemaOpt+"_"+"extra.xo.go"] = true
				}
			}

			if Append(ctx) {
				for filename := range files {
					f, err := out.Open(filename)
					switch {
					case errors.Is(err, os.ErrNotExist):
						continue
					case err != nil:
						return err
					}
					defer f.Close()
					data, err := io.ReadAll(f)
					if err != nil {
						return err
					}
					emit(xo.Template{
						Src:     "{{.Data}}",
						Partial: "header", // ordered first
						Data:    string(data),
						Dest:    filename,
					})
					delete(files, filename)
				}
			}
			for filename := range files {
				emit(xo.Template{
					Partial: "header",
					Dest:    filename,
				})
			}
			return nil
		},
		Process: func(ctx context.Context, mode string, set *xo.Set, emit func(xo.Template)) error {
			if mode == "query" {
				for _, query := range set.Queries {
					if err := emitQuery(ctx, query, emit); err != nil {
						return err
					}
				}
			} else {
				for _, schema := range set.Schemas {
					if err := emitSchema(ctx, schema, emit); err != nil {
						return err
					}
				}
				// TODO: for frontend usage. will also duplicate structs for
				// internal use
				// content := []byte("a\n")

				// err := os.WriteFile("entityFields.json", content, 0644)
				// if err != nil {
				// 	fmt.Println("Error:", err)
				// }
			}
			return nil
		},
		Post: func(ctx context.Context, mode string, files map[string][]byte, emit func(string, []byte)) error {
			for file, content := range files {
				// Run goimports.
				buf, err := imports.Process("", content, nil)
				if err != nil {
					return fmt.Errorf("%s:%w", file, err)
				}
				// Run gofumpt.
				formatted, err := format.Source(buf, format.Options{
					ExtraRules: true,
				})
				if err != nil {
					return err
				}
				emit(file, formatted)
			}
			return nil
		},
	})
	return nil
}

// fileNames returns a list of file names that will be generated by the
// template based on the parameters and schema.
func fileNames(ctx context.Context, mode string, set *xo.Set) (map[string]bool, error) {
	// In single mode, only the specified file be generated.
	singleFile := xo.Single(ctx)
	if singleFile != "" {
		return map[string]bool{
			singleFile: true,
		}, nil
	}
	// Otherwise, infer filenames from set.
	files := make(map[string]bool)
	_, _, schemaOpt := xo.DriverDbSchema(ctx) // cli arg

	addFile := func(filename, schema string) {
		// Filenames are always lowercase.
		var prefix string
		if schema != "public" {
			// don't refactor to use schema + "_" + "name"
			// since it will clash struct names regardless whether
			// its cache.my_table or public.cache_my_table.
			prefix = camel(schema)
		}
		filename = strings.ToLower(prefix + filename)
		files[filename+ext] = true
	}
	switch mode {
	case "schema":
		for _, schema := range set.Schemas {
			for _, e := range schema.Enums {
				// NOTE: we will generate enums and tables from other schemas in the same manner as sqlc
				// for full compat out of the box
				if e.Schema == "public" && schemaOpt != "public" {
					continue // will generate all other schemas alongside public, do not emit again
				}

				addFile(camelExport(e.Name), schema.Name)
			}
			for _, p := range schema.Procs {
				goName := camelExport(p.Name)
				if p.Type == "function" {
					addFile("sf_"+goName, schema.Name)
				} else {
					addFile("sp_"+goName, schema.Name)
				}
			}
			for _, t := range schema.Tables {
				addFile(camelExport(singularize(t.Name)), schema.Name)
			}
			for _, v := range schema.Views {
				addFile(camelExport(singularize(v.Name)), schema.Name)
			}
			for _, v := range schema.MatViews {
				addFile(camelExport(singularize(v.Name)), schema.Name)
			}
		}
	case "query":
		for _, query := range set.Queries {
			addFile(query.Type, "")
			if query.Exec {
				// Single mode is handled at the start of the function but it
				// must be used for Exec queries.
				return nil, ErrNoSingle
			}
		}
	default:
		panic("unknown mode: " + mode)
	}
	return files, nil
}

// emitQuery emits the query.
func emitQuery(ctx context.Context, query xo.Query, emit func(xo.Template)) error {
	var table Table
	// build type if needed
	if !query.Exec {
		var err error
		if table, err = buildQueryType(ctx, query); err != nil {
			return err
		}
	}
	// emit type definition
	if !query.Exec && !query.Flat && !Append(ctx) {
		emit(xo.Template{
			Partial:  "typedef",
			Dest:     strings.ToLower(table.GoName) + ext,
			SortType: query.Type,
			SortName: query.Name,
			Data: struct {
				Table       any
				Constraints any
				Schema      string
			}{Table: table, Constraints: []Constraint{}, Schema: table.Schema},
		})
	}
	// build query params
	var params []QueryParam
	for _, param := range query.Params {
		params = append(params, QueryParam{
			Name:        param.Name,
			Type:        param.Type.Type,
			Interpolate: param.Interpolate,
			Join:        param.Join,
		})
	}
	// emit query
	emit(xo.Template{
		Partial:  "query",
		Dest:     strings.ToLower(table.GoName) + ext,
		SortType: query.Type,
		SortName: query.Name,
		Data: Query{
			Name:        buildQueryName(query),
			Query:       query.Query,
			Comments:    query.Comments,
			Params:      params,
			One:         query.Exec || query.Flat || query.One,
			Flat:        query.Flat,
			Exec:        query.Exec,
			Interpolate: query.Interpolate,
			Type:        table,
			Comment:     query.Comment,
		},
	})
	return nil
}

func buildQueryType(ctx context.Context, query xo.Query) (Table, error) {
	tf := camelExport
	if query.Flat {
		tf = camel
	}
	var fields []Field
	for _, z := range query.Fields {
		f, err := convertField(ctx, tf, z)
		if err != nil {
			return Table{}, err
		}
		// dont use convertField; the types are already provided by the user
		if query.ManualFields {
			f = Field{
				GoName:  z.Name,
				SQLName: snake(z.Name),
				Type:    z.Type.Type,
			}
		}
		fields = append(fields, f)
	}
	sqlName := snake(query.Type)
	return Table{
		GoName:  query.Type,
		SQLName: sqlName,
		Fields:  fields,
		Comment: query.TypeComment,
	}, nil
}

// buildQueryName builds a name for the query.
func buildQueryName(query xo.Query) string {
	if query.Name != "" {
		return query.Name
	}
	// generate name if not specified
	name := query.Type
	if !query.One {
		name = inflector.Pluralize(name)
	}
	// add params
	if len(query.Params) == 0 {
		name = "Get" + name
	} else {
		name += "By"
		for _, p := range query.Params {
			name += camelExport(p.Name)
		}
	}
	return name
}

// Tables contains Table indexed by SQLName
type Tables map[string]Table

// emitSchema emits the xo schema for the template set.
func emitSchema(ctx context.Context, schema xo.Schema, emit func(xo.Template)) error {
	// emit tables
	tcc := append(schema.Tables, schema.Views...)
	tcc = append(tcc, schema.MatViews...)

	// will need access to all tables beforehand for indexes, struct generation...

	tables := make(Tables)
	for _, tc := range tcc {
		table, err := convertTable(ctx, tc)
		if err != nil {
			return err
		}
		tables[table.SQLName] = table
	}

	constraints, err := convertConstraints(ctx, schema.Constraints, tables)
	if err != nil {
		return err
	}
	_, _, schemaOpt := xo.DriverDbSchema(ctx) // cli arg

	for _, e := range schema.Enums {
		if e.Schema == "public" && schemaOpt != "public" {
			continue // will generate all other schemas alongside public, do not emit again
		}

		enum := convertEnum(ctx, e)
		emit(xo.Template{
			Partial:  "enum",
			Dest:     strings.ToLower(enum.GoName) + ext,
			SortName: enum.GoName,
			Data:     enum,
		})
	}
	// build procs
	overloadMap := make(map[string][]Proc)
	// procOrder ensures procs are always emitted in alphabetic order for
	// consistency in single mode
	var procOrder []string
	for _, p := range schema.Procs {
		var err error
		if procOrder, err = convertProc(ctx, overloadMap, procOrder, p); err != nil {
			return err
		}
	}
	// emit procs
	for _, name := range procOrder {
		procs := overloadMap[name]
		prefix := "sp_"
		if procs[0].Type == "function" {
			prefix = "sf_"
		}
		// Set flag to change name to their overloaded versions if needed.
		for i := range procs {
			procs[i].Overloaded = len(procs) > 1
		}
		emit(xo.Template{
			Dest:     prefix + strings.ToLower(name) + ext,
			Partial:  "procs",
			SortName: prefix + name,
			Data:     procs,
		})
	}

	// IMPORTANT: can't use map[string]*Table - it messes up k and t for some reason
	// for k, t := range tables {
	// 	fmt.Printf("k: %v - table: %v\n", k, t.SQLName)
	// }

	for _, tc := range tcc {
		table, err := convertTable(ctx, tc)
		if err != nil {
			return err
		}
		emit(xo.Template{
			Dest:     strings.ToLower(table.GoName) + ext,
			Partial:  "typedef",
			SortType: table.Type,
			SortName: table.GoName,
			Data: struct {
				Table       any
				Constraints any
				Tables      any
				Schema      string
			}{Table: table, Constraints: constraints, Tables: tables, Schema: table.Schema},
		})

		// emit indexes
		var emittedIndexes []string
		for _, i := range tc.Indexes {
			index, err := convertIndex(ctx, table, i)
			if err != nil {
				return err
			}

			newFields, base := removeExcludedIndexTypes(index, excludedIndexTypes)
			if newFields != nil {
				index.Fields = newFields
			}
			if base {
				index.SQLName = "[xo] base filter query"
			}

			idxIdentifier := extractIndexIdentifier(index)
			if contains(emittedIndexes, idxIdentifier) {
				continue
			}

			// emit normal index
			emit(xo.Template{
				Dest:     strings.ToLower(table.GoName) + ext,
				Partial:  "index",
				SortType: table.Type,
				SortName: index.SQLName,
				Data: struct {
					Index       any
					Constraints any
					Tables      any
				}{Index: index, Constraints: constraints, Tables: tables},
			})
			emittedIndexes = append(emittedIndexes, extractIndexIdentifier(index))
		}

		// emit additional indexes in a second run so they don't interfere with "real" ones
		for _, i := range tc.Indexes {
			index, err := convertIndex(ctx, table, i)
			if err != nil {
				return err
			}

			newFields, base := removeExcludedIndexTypes(index, excludedIndexTypes)
			if newFields != nil {
				index.Fields = newFields
			}
			if base {
				index.SQLName = "[xo] base filter query"
			}

			if index.IsUnique && len(index.Fields) > 1 {
				// patch each index and constraints and emit queries with a subset of index fields
				index.IsUnique = false
				for _, f := range index.Fields {
					index.Fields = []Field{f}

					if _, after, ok := strings.Cut(index.Definition, " WHERE "); ok { // index def is normalized in db
						if strings.Contains(after, f.SQLName) {
							// log.Default().Printf("%s: index filter contains field: %s", table.GoName, f.SQLName)
						}
					}

					idxIdentifier := extractIndexIdentifier(index)
					if contains(emittedIndexes, idxIdentifier) {
						continue // most likely a dedicated index already exists
					}

					emit(xo.Template{
						Dest:     strings.ToLower(table.GoName) + ext,
						Partial:  "index",
						SortType: table.Type,
						SortName: index.SQLName,
						Data: struct {
							Index       any
							Constraints any
							Tables      any
						}{Index: index, Constraints: constraints, Tables: tables},
					})
					emittedIndexes = append(emittedIndexes, idxIdentifier)
				}
			}

		}

		// emit fkeys
		for _, fk := range tc.ForeignKeys {
			fkey, err := convertFKey(ctx, table, fk)
			if err != nil {
				return err
			}
			emit(xo.Template{
				Dest:     strings.ToLower(table.GoName) + ext,
				Partial:  "foreignkey",
				SortType: table.Type,
				SortName: fkey.SQLName,
				Data:     fkey,
			})
		}
	}

	// TODO: returns map[string]TableFilter indexed by table.GoName
	// where
	type JoinFilter struct {
		SQLJoin string `json:"sqlJoin"` // xo_join_*
	}
	type FieldFilter struct {
		SQLName string `json:"sqlName"`
		Type    string `json:"type"` // date|datetime|string|number or array of them
	}
	type TableFilter struct {
		// only implement if needed
		// indexed by join field db key. if frontend passes entries, selectOption
		// will reflect and set where db tag matches
		// Joins  map[string]JoinFilter `json:"joins"`
		Fields map[string]FieldFilter // indexed by field json key
	}

	return nil
}

func removeExcludedIndexTypes(index Index, excludedIndexTypes []string) ([]Field, bool) {
	base := false
	excludedColumnNames := extractExcludedColumnNames(index.Definition, excludedIndexTypes)
	if len(excludedColumnNames) == len(index.Fields) {
		fmt.Println("skipping index where all fields are excluded index types: ", index.Definition)
		return []Field{}, true
	}
	if len(excludedColumnNames) > 0 {
		fmt.Println("patching index containing excluded index types: ", index.Definition)
		return patchIndexFields(index.Fields, excludedColumnNames), false
	}

	return nil, base
}

func extractExcludedColumnNames(definition string, excludedIndexTypes []string) []string {
	var excludedColumnNames []string

	re := regexp.MustCompile(`INDEX .*? USING .*?\((?P<columns>[\w\s.,]+)`)
	match := re.FindStringSubmatch(definition)
	subexpNames := re.SubexpNames()

	for i, name := range subexpNames {
		if name == "columns" && len(match) > i {
			idxColumns := strings.Split(match[i], ",")
			for _, idxColumn := range idxColumns {
				column, idxTypePath, _ := strings.Cut(strings.TrimSpace(idxColumn), " ")
				pp := strings.Split(idxTypePath, ".")
				indexTypeName := strings.TrimSpace(pp[len(pp)-1])
				if contains(excludedIndexTypes, indexTypeName) {
					excludedColumnNames = append(excludedColumnNames, column)
				}
			}
		}
	}

	return excludedColumnNames
}

func patchIndexFields(fields []Field, excludedColumnNames []string) []Field {
	var patchedFields []Field

	fmt.Printf("excludedColumnNames: %v\n", excludedColumnNames)

	for _, field := range fields {
		includeField := true
		for _, columnName := range excludedColumnNames {
			if field.SQLName == columnName {
				includeField = false
				break
			}
		}
		if includeField {
			patchedFields = append(patchedFields, field)
		}
	}

	return patchedFields
}

// extractIndexIdentifier generates a unique identifier for patched index generation.
func extractIndexIdentifier(i Index) string {
	excludedColumnNames := extractExcludedColumnNames(i.Definition, excludedIndexTypes)

	var fields []string
	for _, field := range i.Fields {
		if contains(excludedColumnNames, field.SQLName) {
			continue
		}
		fields = append(fields, field.GoName)
	}

	if _, after, ok := strings.Cut(i.Definition, " WHERE "); ok { // index def is normalized in db
		fields = append(fields, after)
	}

	return strings.Join(fields, "-")
}

// convertEnum converts a xo.Enum.
func convertEnum(ctx context.Context, e xo.Enum) Enum {
	var vals []EnumValue
	goName := camelExport(e.Name)
	var prefix string
	_, _, schemaOpt := xo.DriverDbSchema(ctx) // cli arg
	if schemaOpt != "public" {
		prefix = camelExport(schemaOpt)
	}

	for _, v := range e.Values {
		name := camelExport(strings.ToLower(v.Name))
		if strings.HasSuffix(name, goName) && goName != name {
			name = strings.TrimSuffix(name, goName)
		}
		vals = append(vals, EnumValue{
			GoName:     name, // no prefix here, enough just for goname
			SQLName:    v.Name,
			ConstValue: fmt.Sprintf(`"%s"`, v.Name),
		})
	}

	return Enum{
		GoName:       prefix + goName,
		GoNamePrefix: prefix, // for template gen
		SQLName:      e.Name,
		Values:       vals,
	}
}

// convertProc converts a xo.Proc.
func convertProc(ctx context.Context, overloadMap map[string][]Proc, order []string, p xo.Proc) ([]string, error) {
	_, _, schema := xo.DriverDbSchema(ctx)
	proc := Proc{
		Type:      p.Type,
		GoName:    camelExport(p.Name),
		SQLName:   p.Name,
		Signature: fmt.Sprintf("%s.%s", schema, p.Name),
		Void:      p.Void,
	}
	// proc params
	var types []string
	for _, z := range p.Params {
		f, err := convertField(ctx, camel, z)
		if err != nil {
			return nil, err
		}
		proc.Params = append(proc.Params, f)
		types = append(types, z.Type.Type)
	}
	// add to signature, generate name
	proc.Signature += "(" + strings.Join(types, ", ") + ")"
	proc.OverloadedName = overloadedName(types, proc)
	types = nil
	// proc return
	for _, z := range p.Returns {
		f, err := convertField(ctx, camel, z)
		if err != nil {
			return nil, err
		}
		proc.Returns = append(proc.Returns, f)
		types = append(types, z.Type.Type)
	}
	// append signature
	if !p.Void {
		format := " (%s)"
		if len(p.Returns) == 1 {
			format = " %s"
		}
		proc.Signature += fmt.Sprintf(format, strings.Join(types, ", "))
	}
	// add proc
	procs, ok := overloadMap[proc.GoName]
	if !ok {
		order = append(order, proc.GoName)
	}
	overloadMap[proc.GoName] = append(procs, proc)
	return order, nil
}

// convertTable converts a xo.Table to a Table.
func convertTable(ctx context.Context, t xo.Table) (Table, error) {
	var cols, pkCols, generatedCols, ignoredCols []Field
	_, _, schema := xo.DriverDbSchema(ctx)

	for _, z := range t.Columns {
		f, err := convertField(ctx, camelExport, z)
		if err != nil {
			return Table{}, err
		}
		cols = append(cols, f)
		if z.IsPrimary {
			pkCols = append(pkCols, f)
		}
		if f.IsGenerated {
			generatedCols = append(generatedCols, f)
		}
		if f.IsIgnored {
			ignoredCols = append(ignoredCols, f)
		}
	}

	// custom manual override
	manual := false
	for _, pk := range pkCols {
		if !pk.IsGenerated {
			manual = true
			break
		}
	}
	var prefix string
	_, _, schemaOpt := xo.DriverDbSchema(ctx) // cli arg
	if schemaOpt != "public" {
		prefix = camelExport(schemaOpt)
	}

	table := Table{
		GoName:      prefix + camelExport(singularize(t.Name)),
		SQLName:     t.Name,
		Fields:      cols,
		PrimaryKeys: pkCols,
		Generated:   generatedCols,
		Ignored:     ignoredCols,
		Manual:      manual && t.Manual,
		Type:        t.Type,
		Schema:      schema,
	}

	// conversion requires Table
	var fkeys []TableForeignKey
	for _, fk := range t.ForeignKeys {
		fkey, err := convertFKey(ctx, table, fk)
		if err != nil {
			return Table{}, fmt.Errorf("could not convert to fk: %w", err)
		}
		fkFields := make([]string, len(fkey.Fields))
		for i, fkField := range fkey.Fields {
			fkFields[i] = fkField.SQLName
		}
		fkRefFields := make([]string, len(fkey.Fields))
		for i, fkField := range fkey.RefFields {
			fkRefFields[i] = fkField.SQLName
		}

		tfk := TableForeignKey{
			FieldNames: fkFields,
			RefTable:   fk.RefTable,
			RefColumns: fkRefFields,
		}

		fkeys = append(fkeys, tfk)
	}

	table.ForeignKeys = fkeys

	return table, nil
}

func convertIndex(ctx context.Context, t Table, i xo.Index) (Index, error) {
	var fields []Field
	for _, z := range i.Fields {
		f, err := convertField(ctx, camelExport, z)
		if err != nil {
			return Index{}, err
		}
		fields = append(fields, f)
	}
	return Index{
		SQLName:    i.Name,
		Func:       camelExport(i.Func),
		Table:      t,
		Fields:     fields,
		IsUnique:   i.IsUnique,
		IsPrimary:  i.IsPrimary,
		Definition: i.IndexDefinition,
	}, nil
}

func convertConstraints(ctx context.Context, constraints []xo.Constraint, tables Tables) ([]Constraint, error) {
	var cc []Constraint // will create additional dummy constraints for automatic O2O
cc_label:
	for _, constraint := range constraints {
		var card cardinality
		cards := cardinalityRE.FindStringSubmatch(constraint.ColumnComment)
		if len(cards) > 0 {
			card = cardinality(strings.ToUpper(cards[1]))

			if !validCardinality(card) {
				return []Constraint{}, fmt.Errorf("invalid cardinality: %s", card)
			}
		}

		if card != "" {
			switch constraint.Type {
			case "foreign_key":
				fmt.Printf("%-48s | %-12s | %s | %-45s <- %s\n", constraint.Name, constraint.Type, card, constraint.TableName+"."+constraint.ColumnName, constraint.RefTableName+"."+constraint.RefColumnName)
			case "primary_key":
				// may be O2O (PK is FK) or M2M (PKs on lookup table)
				fmt.Printf("%-48s | %-12s | %s | %-45s <- %s\n", constraint.Name, constraint.Type, card, constraint.TableName+"."+constraint.ColumnName, constraint.RefTableName+"."+constraint.RefColumnName)
			case "unique":
				fmt.Printf("%-48s | %-12s | %s | %s \n", constraint.Name, constraint.Type, card, constraint.RefTableName+"."+constraint.RefColumnName)
			}
		}

		var joinTableClash bool
		for _, c := range constraints {
			if c.Type != "foreign_key" {
				continue
			}
			if c.Name == constraint.Name {
				continue // itself
			}
			var ccard cardinality
			ccards := cardinalityRE.FindStringSubmatch(constraint.ColumnComment)
			if len(ccards) > 0 {
				ccard = cardinality(strings.ToUpper(ccards[1]))

				if !validCardinality(ccard) {
					return []Constraint{}, fmt.Errorf("invalid cardinality: %s", ccard)
				}
			}
		outer:
			// TODO need dual M2O-M2M and O2M-O2O checks with reversed ref checks. else generated O2Os from M2Os might clash and we wont know
			// TODO proper pre or postprocessing to simply generated names (remove suffix if not really needed, etc.)
			switch card {
			case M2M:
				if c.ColumnName == constraint.ColumnName && c.RefTableName == constraint.RefTableName && c.RefColumnName == constraint.RefColumnName && ccard == M2M {
					joinTableClash = true
					break outer
				}
			case M2O:
				if c.TableName == constraint.TableName && c.RefTableName == constraint.RefTableName && c.RefColumnName == constraint.RefColumnName && ccard == M2O {
					joinTableClash = true
					break outer
				}
			case O2O:
				// NOTE: probably needs more checks
				if c.TableName == constraint.TableName && c.RefTableName == constraint.RefTableName && c.RefColumnName == constraint.RefColumnName && ccard == O2O {
					joinTableClash = true
					break outer
				}
			}
		}

		// assume it's O2O. Can be overridden at any time
		if constraint.Type == "foreign_key" && card == "" {
			// FIXME generate constraint only if fk fields len = 1
			// and check if field is unique or not
			// ignore duplicate joins generated for partitioned columns to new tables, joined by helper keys, e.g. api_key_id
			for _, seenConstraint := range cc {
				if sameConstraint(seenConstraint, constraint) && seenConstraint.Cardinality == card {
					continue cc_label
				}
			}

			annotations, err := parseAnnotations(constraint.ColumnComment)
			if err != nil {
				panic(fmt.Sprintf("parseAnnotations: %v", err))
			}

			properties := extractPropertiesAnnotation(annotations[propertiesAnnot])

			ignoreConstraints := contains(properties, propertyRefsIgnore)
			if ignoreConstraints {
				continue
			}

			// dummy constraint to automatically create join in FK reference table
			cc = append(cc, Constraint{
				Type:             constraint.Type,
				Cardinality:      O2O,
				Name:             constraint.Name + " (inferred)",
				RefTableName:     constraint.RefTableName,
				TableName:        constraint.TableName,
				RefColumnName:    constraint.RefColumnName,
				RefColumnComment: constraint.RefColumnComment,
				ColumnName:       constraint.ColumnName,
				ColumnComment:    constraint.ColumnComment,
				JoinTableClash:   joinTableClash,
				IsInferredO2O:    true,
			})

			t := tables[constraint.TableName]
			// fmt.Printf("%s: t.PrimaryKeys: %v\n", constraint.TableName, t.PrimaryKeys)
			// fmt.Printf("%s: t.ForeignKeys: %v\n", constraint.TableName, t.ForeignKeys)
			// rt := tables[constraint.RefTableName]
			// fmt.Printf("%s (ref): rt.PrimaryKeys: %v\n", constraint.RefTableName, rt.PrimaryKeys)
			// fmt.Printf("%s (ref): rt.ForeignKeys: %v\n", constraint.RefTableName, rt.ForeignKeys)
			// println(".....")

			var f Field
			for _, tf := range t.PrimaryKeys {
				if tf.SQLName == constraint.RefColumnName {
					f = tf
				}
			}
			// need to check RefTable PKs since this should get called when generating for a
			// table that has *referenced* O2O where PK is FK. e.g. work_item gen -> we see demo_work_item has work_item_id PK that is FK.
			// viceversa we don't care as it's a regular PK.
			af := analyzeField(t, f)
			// FIXME: check should just be looping all constraints and if there exists
			// two contraints one with type == "primary_key" and other "foreign_key" for
			// the same table and column only then RefPKisFK: true
			// this makes no sense
			if af.PKisFK != nil {
				fmt.Printf("%s.%s is a single foreign and primary key in O2O\n", constraint.TableName, constraint.RefColumnName)
				cc = append(cc, Constraint{
					Type:             constraint.Type,
					Cardinality:      O2O,
					Name:             constraint.Name + "(O2O inferred - PK is FK)",
					RefTableName:     constraint.TableName,
					TableName:        constraint.RefTableName,
					RefColumnName:    constraint.ColumnName,
					RefColumnComment: constraint.ColumnComment,
					ColumnName:       constraint.RefColumnName,
					ColumnComment:    constraint.RefColumnComment,
					JoinTableClash:   joinTableClash,
					IsInferredO2O:    true,
					PKisFK:           true,
				})
			}

			continue
		}

		if card == "O2O" {
			for _, seenConstraint := range cc {
				if seenConstraint.TableName == constraint.TableName &&
					seenConstraint.RefTableName == constraint.RefTableName &&
					seenConstraint.ColumnName == constraint.ColumnName &&
					seenConstraint.RefColumnName == constraint.RefColumnName &&
					seenConstraint.Type == constraint.Type &&
					seenConstraint.Cardinality == card {
					continue cc_label
				}
			}
			// create a dummy referenced constraint
			cc = append(cc, Constraint{
				Type:             constraint.Type,
				Cardinality:      O2O,
				Name:             constraint.Name + "(O2O reference)",
				RefTableName:     constraint.TableName,
				TableName:        constraint.RefTableName,
				RefColumnName:    constraint.ColumnName,
				RefColumnComment: constraint.ColumnComment,
				ColumnName:       constraint.RefColumnName,
				ColumnComment:    constraint.RefColumnComment,
				JoinTableClash:   joinTableClash,
			})
		}

		if card == "M2O" {
			/**
			 *
			 */
			for _, seenConstraint := range cc {
				if seenConstraint.TableName == constraint.TableName &&
					seenConstraint.RefTableName == constraint.RefTableName &&
					seenConstraint.ColumnName == constraint.ColumnName &&
					seenConstraint.RefColumnName == constraint.RefColumnName &&
					seenConstraint.Type == constraint.Type && seenConstraint.Cardinality == card {
					continue cc_label
				}
			}

			cc = append(cc, Constraint{
				Type:                  constraint.Type,
				Cardinality:           O2O,
				Name:                  constraint.Name + " (Generated from M2O)",
				TableName:             constraint.TableName,
				RefTableName:          constraint.RefTableName,
				ColumnName:            constraint.ColumnName,
				ColumnComment:         constraint.ColumnComment,
				RefColumnName:         constraint.RefColumnName,
				RefColumnComment:      constraint.RefColumnComment,
				JoinTableClash:        joinTableClash,
				IsGeneratedO2OFromM2O: true,
			})
		}

		cc = append(cc, Constraint{
			Type:             constraint.Type,
			Cardinality:      card, // cardinality comments only needed on FK columns, never base tables
			Name:             constraint.Name,
			TableName:        constraint.TableName,
			RefTableName:     constraint.RefTableName,
			ColumnName:       constraint.ColumnName,
			ColumnComment:    constraint.ColumnComment,
			RefColumnName:    constraint.RefColumnName,
			RefColumnComment: constraint.RefColumnComment,
			JoinTableClash:   joinTableClash,
		})
	}

	// check for name future struct field name and type clashes just once here at startup
	// for _, t := range tables {
	// 	t.SQLName
	// 	for _, constraint := range constraints {
	// 		}
	// }

	return cc, nil
}

func convertFKey(ctx context.Context, t Table, fk xo.ForeignKey) (ForeignKey, error) {
	var fields, refFields []Field
	// convert fields
	for _, f := range fk.Fields {
		field, err := convertField(ctx, camelExport, f)
		if err != nil {
			return ForeignKey{}, err
		}
		fields = append(fields, field)
	}
	// convert ref fields
	for _, f := range fk.RefFields {
		refField, err := convertField(ctx, camelExport, f)
		if err != nil {
			return ForeignKey{}, err
		}
		refFields = append(refFields, refField)
	}
	return ForeignKey{
		GoName:    camelExport(fk.Func),
		SQLName:   fk.Name,
		Table:     t,
		Fields:    fields,
		RefTable:  camelExport(singularize(fk.RefTable)),
		RefFields: refFields,
		RefFunc:   camelExport(fk.RefFunc),
	}, nil
}

func overloadedName(sqlTypes []string, proc Proc) string {
	if len(proc.Params) == 0 {
		return proc.GoName
	}
	var names []string
	// build parameters for proc.
	// if the proc's parameter has no name, use the types of the proc instead
	for i, f := range proc.Params {
		if f.SQLName == fmt.Sprintf("p%d", i) {
			names = append(names, camelExport(strings.Split(sqlTypes[i], " ")...))
			continue
		}
		names = append(names, camelExport(f.GoName))
	}
	if len(names) == 1 {
		return fmt.Sprintf("%sBy%s", proc.GoName, names[0])
	}
	front, last := strings.Join(names[:len(names)-1], ""), names[len(names)-1]
	return fmt.Sprintf("%sBy%sAnd%s", proc.GoName, front, last)
}

func parseAnnotations(comment string) (map[annotation]string, error) {
	annotations := make(map[annotation]string)
	for _, a := range strings.Split(comment, annotationJoinOperator) {
		if a == "" {
			continue
		}
		typ, val, found := strings.Cut(a, annotationAssignmentOperator)
		if !found {
			return nil, fmt.Errorf("invalid column comment annotation format: %s", a)
		}
		typ = annotation(strings.TrimSpace(typ))
		switch typ {
		case cardinalityAnnot, tagsAnnot, typeAnnot, propertiesAnnot:
			annotations[typ] = strings.TrimSpace(val)
		default:
			return nil, fmt.Errorf("invalid column comment annotation type: %s", typ)
		}
	}

	return annotations, nil
}

func convertField(ctx context.Context, tf transformFunc, f xo.Field) (Field, error) {
	typ, zero, err := goType(ctx, f.Type)
	if err != nil {
		return Field{}, err
	}
	var enumPkg, enumSchema, openAPISchema string
	if f.Type.Enum != nil {
		enumPkg = f.Type.Enum.EnumPkg
		enumSchema = f.Type.Enum.Schema
		openAPISchema = camelExport(f.Type.Enum.Name)
	}

	annotations, err := parseAnnotations(f.Comment)
	if err != nil {
		return Field{}, fmt.Errorf("parse annotations: %w", err)
	}

	properties := extractPropertiesAnnotation(annotations[propertiesAnnot])
	typeOverride := annotations[typeAnnot]
	extraTags := annotations[tagsAnnot]
	if extraTags != "" {
		extraTags = " " + extraTags
	}
	originalType := typ
	if typeOverride != "" {
		typ = typeOverride
		if strings.Count(typeOverride, ".") > 0 {
			openAPISchema = camelExport(strings.Split(typeOverride, ".")[1])
		} else {
			// db schema (same package)
			openAPISchema = camelExport(typeOverride)
		}
	}

	return Field{
		Type:           typ,
		GoName:         tf(f.Name),
		SQLName:        f.Name,
		Zero:           zero,
		IsPrimary:      f.IsPrimary,
		IsSequence:     f.IsSequence,
		IsIgnored:      f.IsIgnored,
		EnumPkg:        enumPkg,
		EnumSchema:     enumSchema,
		Comment:        f.Comment,
		IsDateOrTime:   f.IsDateOrTime,
		UnderlyingType: originalType,
		OpenAPISchema:  openAPISchema,
		ExtraTags:      extraTags,
		Properties:     properties,
		IsGenerated:    strings.Contains(f.Default, "()") || f.IsSequence || f.IsGenerated, // TODO: we have default gen_random_uuid(), clock_timestamp(), current_timestamp... not reliable
	}, nil
}

func extractPropertiesAnnotation(annotation annotation) []string {
	var properties []string
	for _, p := range strings.Split(annotation, propertiesJoinOperator) {
		properties = append(properties, strings.TrimSpace(strings.ToLower(p)))
	}

	return properties
}

func goType(ctx context.Context, typ xo.Type) (string, string, error) {
	_, _, schema := xo.DriverDbSchema(ctx)
	var f func(xo.Type, string, string, string) (string, string, error)
	switch mode := ArrayMode(ctx); mode {
	case "stdlib":
		f = loader.StdlibPostgresGoType
	case "pq", "":
		f = loader.PQPostgresGoType
	default:
		return "", "", fmt.Errorf("unknown array mode: %q", mode)
	}
	return f(typ, schema, Int32(ctx), Uint32(ctx))
}

type transformFunc func(...string) string

func snake(names ...string) string {
	return snaker.CamelToSnake(strings.Join(names, "_"))
}

func camel(names ...string) string {
	return snaker.ForceLowerCamelIdentifier(strings.Join(names, "_"))
}

// NOTE: broken func with .
func camelExport(names ...string) string {
	return snaker.ForceCamelIdentifier(strings.Join(names, "_"))
}

const ext = ".xo.go"

// can extend as required to prevent getting db info via reflection
type DbField struct {
	// Type is one of: string, number, integer, boolean, date-time
	// Arrays and objects are ignored for default filter generation
	Type string `json:"type"`
	// Db is the corresponding db column name
	Db       string `json:"db"`
	Nullable bool   `json:"nullable"`
	Public   bool   `json:"public"`
}

// Funcs is a set of template funcs.
type Funcs struct {
	driver          string
	schema          string
	schemaPrefix    string
	currentDatabase string
	nth             func(int) string
	first           bool
	pkg             string
	tags            []string
	imports         []string
	// joinTableDbTags map[string]map[string]string
	tableConstraints map[string][]Constraint
	// filters indexed by entity name and json field name.
	entityFields    map[string]map[string]DbField
	conflict        string
	custom          string
	escSchema       bool
	escTable        bool
	escColumn       bool
	fieldtag        *template.Template
	publicfieldtag  *template.Template
	privatefieldtag *template.Template
	context         string
	inject          string
	// knownTypes is the collection of known Go types.
	knownTypes map[string]bool
	// shorts is the collection of Go style short names for types, mainly
	// used for use with declaring a func receiver on a type.
	shorts map[string]string
}

// NewFuncs creates custom template funcs for the context.
func NewFuncs(ctx context.Context) (template.FuncMap, error) {
	first := !NotFirst(ctx)
	publicfieldtag, err := template.New("publicfieldtag").Funcs(template.FuncMap{"camel": camel}).Parse(PublicFieldTag(ctx))
	if err != nil {
		return nil, err
	}
	privatefieldtag, err := template.New("privatefieldtag").Funcs(template.FuncMap{"camel": camel}).Parse(PrivateFieldTag(ctx))
	if err != nil {
		return nil, err
	}
	// parse field tag template
	fieldtag, err := template.New("fieldtag").Funcs(template.FuncMap{"camel": camel}).Parse(FieldTag(ctx))
	if err != nil {
		return nil, err
	}
	// load inject
	inject := Inject(ctx)
	if s := InjectFile(ctx); s != "" {
		buf, err := ioutil.ReadFile(s)
		if err != nil {
			return nil, fmt.Errorf("unable to read file: %v", err)
		}
		inject = string(buf)
	}
	driver, sqldb, schema := xo.DriverDbSchema(ctx)
	var currentDatabase string
	err = sqldb.QueryRow("SELECT current_database()").Scan(&currentDatabase)
	if err != nil {
		panic(err)
	}
	var schemaPrefix string
	if schema != "public" {
		schemaPrefix = schema
	}
	nth, err := loader.NthParam(ctx)
	if err != nil {
		return nil, err
	}
	funcs := &Funcs{
		tableConstraints: make(map[string][]Constraint),
		entityFields:     make(map[string]map[string]DbField),
		first:            first,
		currentDatabase:  currentDatabase,
		driver:           driver,
		schema:           schema,
		schemaPrefix:     schemaPrefix,
		nth:              nth,
		pkg:              Pkg(ctx),
		tags:             Tags(ctx),
		imports:          Imports(ctx),
		conflict:         Conflict(ctx),
		custom:           Custom(ctx),
		escSchema:        Esc(ctx, "schema"),
		escTable:         Esc(ctx, "table"),
		escColumn:        Esc(ctx, "column"),
		fieldtag:         fieldtag,
		publicfieldtag:   publicfieldtag,
		privatefieldtag:  privatefieldtag,
		context:          Context(ctx),
		inject:           inject,
		knownTypes:       KnownTypes(ctx),
		shorts:           Shorts(ctx),
	}
	return funcs.FuncMap(), nil
}

func (f *Funcs) camel_export(names ...string) string {
	return snaker.ForceCamelIdentifier(strings.Join(names, "_"))
}

func (f *Funcs) camel(names ...string) string {
	return snaker.ForceLowerCamelIdentifier(strings.Join(names, "_"))
}

func (f *Funcs) sentence_case(names ...string) string {
	c := strings.Title(snaker.CamelToSnake(strings.Join(names, "_")))
	return singularize(strings.ReplaceAll(c, "_", " "))
}

// FuncMap returns the func map.
func (f *Funcs) FuncMap() template.FuncMap {
	return template.FuncMap{
		// general
		"entities":      f.entities,
		"sentence_case": f.sentence_case,
		"camel":         f.camel,
		"camel_export":  f.camel_export,
		"lowerFirst":    f.lower_first,
		"first":         f.firstfn,
		"driver":        f.driverfn,
		"schema":        f.schemafn,
		"pkg":           f.pkgfn,
		"tags":          f.tagsfn,
		"imports":       f.importsfn,
		"inject":        f.injectfn,
		// context
		"context":         f.contextfn,
		"context_both":    f.context_both,
		"context_disable": f.context_disable,
		// func opts
		"initial_opts": f.initial_opts,
		// func and query
		"func_name_context":      f.func_name_context,
		"has_deleted_at":         f.has_deleted_at,
		"func_name":              f.func_name_none,
		"func_context":           f.func_context,
		"extratypes":             f.extratypes,
		"func":                   f.func_none,
		"recv_context":           f.recv_context,
		"recv":                   f.recv_none,
		"foreign_key_context":    f.foreign_key_context,
		"foreign_key":            f.foreign_key_none,
		"db":                     f.db,
		"db_prefix":              f.db_prefix,
		"db_update":              f.db_update,
		"db_named":               f.db_named,
		"named":                  f.named,
		"generate_entity_fields": f.generate_entity_fields,
		"logf":                   f.logf,
		"logf_pkeys":             f.logf_pkeys,
		"logf_update":            f.logf_update,
		// type
		"initialize_constraints": f.initialize_constraints,
		"names":                  f.names,
		"names_all":              f.names_all,
		"names_ignore":           f.names_ignore,
		"params":                 f.params,
		"zero":                   f.zero,
		"type":                   f.typefn,
		"field":                  f.field,
		"set_field":              f.set_field,
		"sort_fields":            f.sort_fields,
		"fieldmapping":           f.fieldmapping,
		"join_fields":            f.join_fields,
		"short":                  f.short,
		// sqlstr funcs
		"querystr":                   f.querystr,
		"sqlstr":                     f.sqlstr,
		"sqlstr_index":               f.sqlstr_index,
		"sqlstr_paginated":           f.sqlstr_paginated,
		"db_paginated":               f.db_paginated,
		"cursor_columns":             f.cursor_columns,
		"func_name_context_suffixed": f.func_name_context_suffixed,
		"recv_context_suffixed":      f.recv_context_suffixed,
		"last_nth":                   f.last_nth,
		// helpers
		"combine_values":     combine_values,
		"fields_to_goname":   fields_to_goname,
		"check_name":         checkName,
		"eval":               eval,
		"add":                add,
		"table_is_updatable": table_is_updatable,
	}
}

// last_nth returns the last hardcoded nth param for sqlstr
func (f *Funcs) last_nth(v any, tables Tables, fields ...any) string {
	var extraFields []Field
	for _, field := range fields {
		switch x := field.(type) {
		case []Field:
			extraFields = append(extraFields, x...)
		case Field:
			extraFields = append(extraFields, x)
		default:
			return fmt.Sprintf("[[ UNSUPPORTED TYPE last_nth fields: %T ]]", x)
		}
	}

	switch x := v.(type) {
	case Index:
		tableName := x.Table.SQLName
		t := x.Table
		return lastNth(tableName, tables, t, append(x.Fields, extraFields...))
	case Table:
		tableName := x.SQLName
		t := x
		return lastNth(tableName, tables, t, extraFields)
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE last_nth: %T ]]", v)
	}
	// return fmt.Sprintf("[[ UNSUPPORTED TYPES last_nth: %T - %T - %s ]]", v, w, typ)
}

func lastNth(tableName string, tables Tables, t Table, fields []Field) string {
	var n int

	for range fields {
		n++
	}

	return strconv.Itoa(n)
}

func (f *Funcs) lower_first(str string) string {
	var b strings.Builder
	i := 1
	for IsUpper(string(str[i : i+1])) {
		if i == len(str)-1 {
			i = len(str) + 1
			break
		}
		i++
	}
	if i == 1 {
		i++ // first letter always lower
	}
	b.WriteString(strings.ToLower(string(str[0 : i-1])))
	b.WriteString(string(str[i-1:]))

	return b.String()
}

func combine_values(values ...string) []string {
	return values
}

func (f *Funcs) firstfn() bool {
	if f.first {
		f.first = false
		return true
	}
	return false
}

// driverfn returns true if the driver is any of the passed drivers.
func (f *Funcs) driverfn(drivers ...string) bool {
	for _, driver := range drivers {
		if f.driver == driver {
			return true
		}
	}
	return false
}

// schemafn takes a series of names and joins them with the schema name.
func (f *Funcs) schemafn(names ...string) string {
	s := f.schema
	// escape table names
	if f.escTable {
		for i, name := range names {
			names[i] = escfn(name)
		}
	}
	n := strings.Join(names, ".")
	switch {
	case s == "" && n == "":
		return ""
	case f.driver == "sqlite3" && n == "":
		return f.schema
	case f.driver == "sqlite3":
		return n
	case s != "" && n != "":
		if f.escSchema {
			s = escfn(s)
		}
		s += "."
	}
	return s + n
}

// pkgfn returns the package name.
func (f *Funcs) pkgfn() string {
	return f.pkg
}

// tagsfn returns the tags.
func (f *Funcs) tagsfn() []string {
	return f.tags
}

// importsfn returns the imports.
func (f *Funcs) importsfn() []PackageImport {
	var imports []PackageImport
	for _, s := range f.imports {
		alias, pkg := "", s
		if i := strings.Index(pkg, " "); i != -1 {
			alias, pkg = pkg[:i], strings.TrimSpace(pkg[i:])
		}
		imports = append(imports, PackageImport{
			Alias: alias,
			Pkg:   strconv.Quote(pkg),
		})
	}
	return imports
}

// contextfn returns true when the context mode is both or only.
func (f *Funcs) contextfn() bool {
	return f.context == "both" || f.context == "only"
}

// context_both returns true with the context mode is both.
func (f *Funcs) context_both() bool {
	return f.context == "both"
}

// context_disable returns true with the context mode is both.
func (f *Funcs) context_disable() bool {
	return f.context == "disable"
}

// injectfn returns the injected content provided from args.
func (f *Funcs) injectfn() string {
	return f.inject
}

// func_name_none builds a func name.
func (f *Funcs) func_name_none(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case Query:
		return x.Name
	case Table:
		return x.GoName
	case ForeignKey:
		return x.GoName
	case Proc:
		n := x.GoName
		if x.Overloaded {
			n = x.OverloadedName
		}
		return n
	case Index:
		return x.Func
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 1: %T ]]", v)
}

// has_deleted_at checks if a table has a deleted_at column.
func (f *Funcs) has_deleted_at(t Table) bool {
	for _, f := range t.Fields {
		if f.SQLName == "deleted_at" {
			return true
		}
	}
	return false
}

// func_name_context generates a name for the func.
func (f *Funcs) func_name_context(v any, suffix string) string {
	switch x := v.(type) {
	case string:
		return x + suffix
	case Query:
		return x.Name + suffix
	case Table:
		return x.GoName + suffix
	case ForeignKey:
		var fields []string
		for _, f := range x.Fields {
			fields = append(fields, f.GoName)
		}

		return "FK" + x.GoName + "_" + strings.Join(fields, "") // else clash with join fields in struct
	case Proc:
		n := x.GoName
		if x.Overloaded {
			n = x.OverloadedName
		}
		return n
	case Index:
		var fields []string
		var suffix string
		name := x.Table.GoName
		if !x.IsUnique {
			name = inflector.Pluralize(name)
		}

		for _, field := range x.Fields {
			fields = append(fields, field.GoName)
		}

		if _, after, ok := strings.Cut(x.Definition, " WHERE "); ok { // index def is normalized in db
			suffix = "_" + snaker.ForceCamelIdentifier("Where "+strings.ToLower(after))
		}

		// need custom Func to handle additional index creation instead of Func field
		// https://github.com/danicc097/xo/blob/main/cmd/schema.go#L629 which originally sets i.Func
		cond := ""
		if len(fields) > 0 {
			cond = "By" + strings.Join(fields, "")
		}

		funcName := name + cond + suffix

		return funcName
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 2: %T ]]", v)
}

// funcfn builds a func definition.
func (f *Funcs) funcfn(name string, context bool, v any, columns []Field, table Table, extraFields string) string {
	var params, returns []string
	if context {
		params = append(params, "ctx context.Context")
	}
	params = append(params, "db DB")
	switch x := v.(type) {
	case Query:
		for _, z := range x.Params {
			params = append(params, fmt.Sprintf("%s %s", z.Name, z.Type))
		}
		if extraFields != "" {
			params = append(params, extraFields)
		}
		switch {
		case x.Exec:
			returns = append(returns, "sql.Result")
		case x.Flat:
			for _, z := range x.Type.Fields {
				returns = append(returns, f.typefn(z.Type))
			}
		case x.One:
			returns = append(returns, "*"+camelExport(f.schemaPrefix)+x.Type.GoName)
		default:
			returns = append(returns, "[]*"+camelExport(f.schemaPrefix)+x.Type.GoName)
		}
	case Proc:
		params = append(params, f.params(x.Params, true, table))
		if extraFields != "" {
			params = append(params, extraFields)
		}
		if !x.Void {
			for _, ret := range x.Returns {
				returns = append(returns, f.typefn(ret.Type))
			}
		}
	case Index:
		params = append(params, f.params(x.Fields, true, table))
		if extraFields != "" {
			params = append(params, extraFields)
		}
		params = append(params, "opts ..."+x.Table.GoName+"SelectConfigOption")
		rt := x.Table.GoName
		if !x.IsUnique {
			rt = "[]" + rt
		} else {
			rt = "*" + rt
		}
		returns = append(returns, rt)
	case Table: // Paginated query
		params = append(params, f.params(columns, true, table))
		if extraFields != "" {
			params = append(params, extraFields)
		}

		params = append(params, "opts ..."+x.GoName+"SelectConfigOption")
		rt := "[]" + x.GoName

		returns = append(returns, rt)
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 3: %T ]]", v)
	}
	returns = append(returns, "error")

	p := ""
	params = removeEmptyStrings(params)
	if len(params) > 0 {
		p = strings.Join(params, ", ")
	}

	return fmt.Sprintf("func %s(%s) (%s)", name, p, strings.Join(returns, ", "))
}

// initial_opts returns base conf for select queries.
func (f *Funcs) initial_opts(v any) string {
	var tableHasDeletedAt bool
	var deletedAtNullIndexCond, deletedAtNotNullIndexCond bool
	var buf strings.Builder

	switch x := v.(type) {
	case Table:
		for _, field := range x.Fields {
			if field.SQLName == "deleted_at" {
				tableHasDeletedAt = true
			}
		}
		buf.WriteString(fmt.Sprintf(`c := &%sSelectConfig{`, x.GoName))
		if tableHasDeletedAt {
			buf.WriteString(`deletedAt: " null ",`)
		}
		buf.WriteString(fmt.Sprintf(`joins: %sJoins{},`, x.GoName))
		buf.WriteString(`
		filters: make(map[string][]any),
		having: make(map[string][]any),
		orderBy: make(map[string]Direction),
}`)
	case Index:
		for _, field := range x.Table.Fields { // table fields, not index fields
			if field.SQLName == "deleted_at" {
				tableHasDeletedAt = true
			}
		}
		if _, after, ok := strings.Cut(x.Definition, " WHERE "); ok { // index def is normalized in db
			if strings.Contains(strings.ToLower(after), "deleted_at is not null") {
				deletedAtNotNullIndexCond = true
			}
			if strings.Contains(strings.ToLower(after), "deleted_at is null") {
				deletedAtNullIndexCond = true
			}
		}

		buf.WriteString(fmt.Sprintf(`c := &%sSelectConfig{`, x.Table.GoName))

		if deletedAtNullIndexCond {
			buf.WriteString(`deletedAt: " null ",`)
		} else if deletedAtNotNullIndexCond {
			buf.WriteString(`deletedAt: " not null ",`)
		} else if tableHasDeletedAt {
			buf.WriteString(`deletedAt: " null ",`)
		}

		buf.WriteString(fmt.Sprintf(`joins: %sJoins{},`, x.Table.GoName))
		buf.WriteString(`filters: make(map[string][]any), having: make(map[string][]any),
}`)
	default:
		return ""
	}

	return buf.String()
}

// funcfn builds a type definition.
func (f *Funcs) extratypes(tGoName string, sqlname string, constraints []Constraint, t Table, tables Tables) string {
	if len(constraints) > 0 {
		// always run
		f.loadConstraints(constraints, sqlname, nil)
	}

	// -- emit ORDER BY opts

	var tableHasDeletedAt bool
	var orderbys []Field
	var cc []Constraint

	for _, z := range t.Fields {
		if z.IsDateOrTime {
			orderbys = append(orderbys, z)
		}
		if z.SQLName == "deleted_at" {
			tableHasDeletedAt = true
		}
	}
	if tablecc, ok := f.tableConstraints[t.SQLName]; ok {
		cc = tablecc
	}

	var buf strings.Builder
	var sqlstrBuf strings.Builder

	buf.WriteString(fmt.Sprintf(`
	type %[1]sSelectConfig struct {
		limit       string
		orderBy     map[string]Direction
		joins       %[1]sJoins
		filters     map[string][]any
		having     map[string][]any
		`, tGoName))
	if tableHasDeletedAt {
		buf.WriteString(`
			deletedAt   string`)
	}
	buf.WriteString(`
	}`)
	buf.WriteString(fmt.Sprintf(`
	type %[1]sSelectConfigOption func(*%[1]sSelectConfig)

	// With%[1]sLimit limits row selection.
	func With%[1]sLimit(limit int) %[1]sSelectConfigOption {
		return func(s *%[1]sSelectConfig) {
			if limit > 0 {
				s.limit = fmt.Sprintf(" limit %%d ", limit)
			}
		}
	}`, tGoName))

	if tableHasDeletedAt {
		buf.WriteString(fmt.Sprintf(`
	// WithDeleted%[1]sOnly limits result to records marked as deleted.
	func WithDeleted%[1]sOnly() %[1]sSelectConfigOption {
		return func(s *%[1]sSelectConfig) {
			s.deletedAt = " not null "
		}
	}`, tGoName))
	}

	buf.WriteString(fmt.Sprintf(`
// With%[1]sOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func With%[1]sOrderBy(rows map[string]*Direction) %[1]sSelectConfigOption {
	return func(s *%[1]sSelectConfig) {
		te := %[2]sEntityFields[%[2]sTableEntity%[1]s]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}
	`, tGoName, camelExport(f.schemaPrefix)))

	var extraStructs []string

	buf.WriteString(fmt.Sprintf("type %sJoins struct {\n", tGoName))

	var goNames []string

	funcs := template.FuncMap{
		"singularize": singularize,
	}

	for _, c := range cc {
		if c.Type != "foreign_key" {
			continue
		}

		var goName string
		notes := "// "

		switch c.Cardinality {
		case M2M:
			notes += string(c.Cardinality) + " " + c.TableName
			lookupName := strings.TrimSuffix(c.ColumnName, "_id")
			goName = camelExport(inflector.Pluralize(lookupName))

			rc := camelExport(strings.TrimSuffix(c.RefColumnName, "_id"))
			rlc := camelExport(strings.TrimSuffix(c.LookupRefColumnName, "_id"))
			lc := camelExport(strings.TrimSuffix(c.LookupColumnName, "_id"))
			col := camelExport(strings.TrimSuffix(c.ColumnName, "_id"))
			_, _, _, _ = rc, rlc, lc, col
			if rlc != lc && rlc != col {
				fmt.Printf("(rlc != lc) lc: %v rlc: %v rc: %v col: %v\n", lc, rlc, rc, col)
				goName = lc + goName
			}
			// if rc != col && rlc != col {
			// 	fmt.Printf("(rc != col) lc: %v rlc: %v rc: %v col: %v\n", lc, rlc, rc, col)
			// 	goName = col + goName
			// }
			lookupTable := tables[c.TableName]
			m2mExtraCols := getTableRegularFields(lookupTable)
			if len(m2mExtraCols) > 0 {
				originalStruct := camelExport(singularize(strings.TrimSuffix(c.RefColumnName, "_id")))
				tag := fmt.Sprintf("`json:\"%s\" db:\"%s\" required:\"true\"`", camel(originalStruct), inflector.Pluralize(strings.TrimSuffix(c.RefColumnName, "_id")))

				// create custom struct for each join with lookup table that has extra fields
				var lookupFields []string

				for _, col := range m2mExtraCols {
					tag := fmt.Sprintf("`json:\"%s\" db:\"%s\"", camel(col.GoName), col.SQLName)
					properties := extractPropertiesAnnotation(col.Annotations[propertiesAnnot])
					isPrivate := contains(properties, propertyJSONPrivate)
					if !isPrivate {
						tag = tag + ` required:"true"`
					}
					if col.OpenAPISchema != "" {
						tag = tag + ` ref:"#/components/schemas/` + col.OpenAPISchema + `"`
					}
					tag = tag + " " + col.ExtraTags + "`"
					typ := f.typefn(col.Type)
					if strings.HasPrefix(typ, "Null") && f.schemaPrefix != "public" && col.EnumSchema != "public" {
						typ = "*" + camelExport(f.schemaPrefix) + strings.TrimPrefix(typ, "Null")
					}

					lookupFields = append(lookupFields, fmt.Sprintf("%s %s %s", camelExport(col.GoName), typ, tag))
				}
				joinField := originalStruct + " " + camelExport(f.schemaPrefix) + originalStruct + " " + tag
				// typ := camelExport(singularize(c.RefTableName)) // same typ as in struct

				st := camelExport(tGoName) + "M2M" + camelExport(lookupName) + toAcronym(c.TableName)
				lookupTableSQLName := f.schema + "." + c.TableName
				docstring := fmt.Sprintf("// %s represents a M2M join against %q", st, lookupTableSQLName)

				extraStructs = append(extraStructs, (fmt.Sprintf(`
%s
type %s struct {
	%s
	%s
}
	`, docstring, st, joinField, strings.Join(lookupFields, "\n")))) // prevent name clashing
			}
		case M2O:
			if c.RefTableName != sqlname {
				continue
			}
			notes += string(c.Cardinality) + " " + c.TableName
			goName = camelExport(c.TableName)
			if c.JoinTableClash {
				goName = camelExport(c.ColumnName) + goName
			}
		case O2O:
			if c.TableName != sqlname {
				continue
			}
			notes += string(c.Cardinality) + " " + c.RefTableName
			goName = camelExport(singularize(c.RefTableName))
			lc := strings.TrimSuffix(c.ColumnName, "_id")
			if c.JoinTableClash {
				goName = goName + camelExport(lc)
			}
			// dummy created automatically to avoid this duplication
			// if c.RefTableName == sqlname {
			// 	joinName = camelExport(singularize(c.TableName))
			// }
		default:
		}
		if goName == "" {
			continue
		}
		for _, g := range goNames {
			if g == goName {
				fmt.Printf("preventing clash -- joinName: %v\n", goName)
				goName = goName + toAcronym(c.TableName)
			}
		}
		goNames = append(goNames, goName)
		var tag string

		// no need for adapter from db.*Joins to oapi-codegens' Db*Joins
		tag = fmt.Sprintf("`json:\"%s\" required:\"true\" nullable:\"false\"`", camel(goName))
		buf.WriteString(fmt.Sprintf("%s bool %s %s\n", goName, tag, notes))

		joinClause, selectClause, groupby := f.createJoinStatement(tables, c, t, funcs)
		if joinClause == "" || selectClause == "" {
			continue
		}

		// prevent clashing
		sqlstrBuf.WriteString(fmt.Sprintf("const %sTable%sJoinSQL = `%s`\n\n", f.lower_first(tGoName), goName, joinClause))
		sqlstrBuf.WriteString(fmt.Sprintf("const %sTable%sSelectSQL = `%s`\n\n", f.lower_first(tGoName), goName, selectClause))
		sqlstrBuf.WriteString(fmt.Sprintf("const %sTable%sGroupBySQL = `%s`\n\n", f.lower_first(tGoName), goName, groupby))
	}

	buf.WriteString("}\n")

	// recursive would go out of hand quickly, use go-jet or sqlc for these cases.
	buf.WriteString(fmt.Sprintf(`
	// With%[1]sJoin joins with the given tables.
func With%[1]sJoin(joins %[1]sJoins) %[1]sSelectConfigOption {
	return func(s *%[1]sSelectConfig) {
		s.joins = %[1]sJoins{
	`, tGoName))

	for _, g := range goNames {
		buf.WriteString(fmt.Sprintf("\t\t%[1]s:  s.joins.%[1]s || joins.%[1]s,\n", g))
	}
	buf.WriteString(`
		}
	}
}`)

	for _, stt := range extraStructs {
		buf.WriteString(stt)
	}

	/**
	 * ideally the having clause should be generated for each table differently, containing the mapping
	 * of join name to xo_join_... so its a matter of just reading the doc instead of diving into the selectSQL string
	 * not so trivial as it is now
	 * */
	buf.WriteString(fmt.Sprintf(`
// With%[1]sFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//filters := map[string][]any{
//	"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//	`+"`(col.created_at > $i OR \n//	col.is_closed = $i)`"+`: {time.Now().Add(-24 * time.Hour), true},
//}
func With%[1]sFilters(filters map[string][]any) %[1]sSelectConfigOption {
	return func(s *%[1]sSelectConfig) {
		s.filters = filters
	}
}

// With%[1]sHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func With%[1]sHavingClause(conditions map[string][]any) %[1]sSelectConfigOption {
	return func(s *%[1]sSelectConfig) {
		s.having = conditions
	}
}
	`, tGoName))

	buf.WriteString(sqlstrBuf.String())

	return buf.String()
}

// func_context generates a func signature for v with context determined by the
// context mode.
func (f *Funcs) func_context(v any, suffix string, columns any, w any, extraFields string) string {
	var cc []Field
	var t Table
	switch x := columns.(type) {
	case []Field:
		cc = x
	default:
		cc = []Field{}
	}
	switch x := w.(type) {
	case Table:
		t = x
	default:
		t = Table{}
	}

	return f.funcfn(f.func_name_context(v, suffix), f.contextfn(), v, cc, t, extraFields)
}

// func_none genarates a func signature for v without context.
func (f *Funcs) func_none(v any, columns any, extraFields string) string {
	var cc []Field
	switch x := columns.(type) {
	case []Field:
		cc = x
	default:
		cc = []Field{}
	}
	return f.funcfn(f.func_name_none(v), false, v, cc, Table{}, extraFields)
}

// recv builds a receiver func definition.
func (f *Funcs) recv(name string, context bool, t Table, v any) string {
	short := f.short(t)
	var p, r []string
	// determine params and return type
	if context {
		p = append(p, "ctx context.Context")
	}
	p = append(p, "db DB")
	switch x := v.(type) {
	case ForeignKey:
		r = append(r, "*"+camelExport(f.schemaPrefix)+x.RefTable)
	case string:
		if x == "Delete" || x == "SoftDelete" { // only exec
			break
		}
		if x == "Upsert" {
			p = append(p, fmt.Sprintf("params *%sCreateParams", t.GoName))
			r = append(r, "*"+t.GoName)
			break
		}
		if x == "Paginated" {
			r = append(r, "[]"+t.GoName)
			break
		}
		r = append(r, "*"+t.GoName)
	}
	r = append(r, "error")
	params := ""
	if len(p) > 0 {
		params = strings.Join(p, ", ")
	}
	return fmt.Sprintf("func (%s *%s) %s(%s) (%s)", short, t.GoName, name, params, strings.Join(r, ", "))
}

func fields_to_goname(fields []Field, sep string) string {
	var res string
	for _, s := range fields {
		res += s.GoName
	}

	return res
}

// cant explode in template
func (f *Funcs) func_name_context_suffixed(typ any, suffixes string) string {
	return f.func_name_context(typ, suffixes)
}

// cant explode in template
func (f *Funcs) recv_context_suffixed(typ any, v any, suffixes string) string {
	return f.recv_context(typ, v, suffixes)
}

// recv_context builds a receiver func definition with context determined by
// the context mode.
func (f *Funcs) recv_context(typ any, v any, suffixes string) string {
	switch x := typ.(type) {
	case Table:
		return f.recv(f.func_name_context(v, suffixes), f.contextfn(), x, v)
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 4: %T ]]", typ)
}

// recv_none builds a receiver func definition without context.
func (f *Funcs) recv_none(typ any, v any) string {
	switch x := typ.(type) {
	case Table:
		return f.recv(f.func_name_none(v), false, x, v)
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 5: %T ]]", typ)
}

func (f *Funcs) foreign_key_context(v any) string {
	var name string
	var p []string
	if f.contextfn() {
		p = append(p, "ctx")
	}
	switch x := v.(type) {
	case ForeignKey:
		name = camelExport(f.schemaPrefix) + x.RefFunc
		// add params
		p = append(p, "db", f.convertTypes(x))
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 6: %T ]]", v)
	}
	return fmt.Sprintf("%s(%s)", name, strings.Join(p, ", "))
}

func (f *Funcs) foreign_key_none(v any) string {
	var name string
	var p []string
	switch x := v.(type) {
	case ForeignKey:
		name = x.RefFunc
		p = append(p, "context.Background()", "db", f.convertTypes(x))
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 7: %T ]]", v)
	}
	return fmt.Sprintf("%s(%s)", name, strings.Join(p, ", "))
}

// db generates a db.<name>Context(ctx, sqlstr, ...)
func (f *Funcs) db(name string, v ...any) string {
	// params
	var p []any
	// for sqlc compatibility always use context
	// if f.contextfn() {
	p = append(p, "ctx")
	// }
	p = append(p, "sqlstr")
	return fmt.Sprintf("db.%s(%s)", name, f.names("", append(p, v...)...))
}

// db_prefix generates a db.<name>Context(ctx, sqlstr, <prefix>.param, ...).
// Similar to db
//
// Will skip the specific parameters based on the type provided.
func (f *Funcs) db_prefix(name string, includeGenerated bool, includeIgnored bool, vs ...any) string {
	var prefix string
	var params []any
	for i, v := range vs {
		var ignore []any
		switch x := v.(type) {
		case string:
			params = append(params, x)
		case Table:
			prefix = f.short(x.GoName) + "."
			// skip primary keys and ignored fields for insertion
			for _, field := range x.Fields {
				if (field.IsGenerated && !includeGenerated) || (field.IsIgnored && !includeIgnored) {
					ignore = append(ignore, field.GoName)
				}
			}
			p := f.names_ignore(prefix, v, ignore...)
			// p is "" when no columns are present except for primary key
			if p != "" {
				params = append(params, p)
			}
		default:
			return fmt.Sprintf("[[ UNSUPPORTED TYPE 8 (%d): %T ]]", i, v)
		}
	}
	return f.db(name, params...)
}

// db_update generates a db.<name>Context(ctx, sqlstr, regularparams,
// primaryparams)
func (f *Funcs) db_update(name string, v any) string {
	var ignore []any
	var p []string
	switch x := v.(type) {
	case Table:
		prefix := f.short(x.GoName) + "."
		for _, pk := range x.Generated {
			ignore = append(ignore, pk.GoName)
		}
		for _, pk := range x.PrimaryKeys {
			ignore = append(ignore, pk.GoName)
		}
		for _, pk := range x.Ignored {
			ignore = append(ignore, pk.GoName)
		}
		p = append(p, f.names_ignore(prefix, x, ignore...), f.names(prefix, x.PrimaryKeys))
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 9: %T ]]", v)
	}
	return f.db(name, strings.Join(p, ", "))
}

// db_paginated generates a db.<name>Context(ctx, sqlstr, params)
// query for cursor pagination by the given columns
func (f *Funcs) db_paginated(name string, t Table) string {
	return f.db(name, CursorPagination{Table: t})
}

// db_named generates a db.<name>Context(ctx, sql.Named(name, res)...)
func (f *Funcs) db_named(name string, v any) string {
	var p []string
	switch x := v.(type) {
	case Proc:
		for _, z := range x.Params {
			p = append(p, f.named(z.SQLName, z.GoName, false))
		}
		for _, z := range x.Returns {
			p = append(p, f.named(z.SQLName, "&"+z.GoName, true))
		}
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 10: %T ]]", v)
	}
	return f.db(name, strings.Join(p, ", "))
}

func (f *Funcs) entities(schema string, tables Tables) string {
	if schema == "public" {
		schema = ""
	}
	var b strings.Builder

	tt := make([]Table, len(tables))
	i := 0
	for _, t := range tables {
		tt[i] = t
		i++
	}

	sort.Slice(tt, func(i, j int) bool {
		return tt[i].GoName < tt[j].GoName
	})

	b.WriteString(fmt.Sprintf("type %sTableEntity string\n", camelExport(schema)))
	b.WriteString("const (\n")
	for _, t := range tt {
		tident := t.SQLName
		if f.schemaPrefix != "" {
			tident = f.schemaPrefix + "." + tident
		}
		b.WriteString(fmt.Sprintf("%[1]sTableEntity%[2]s %[1]sTableEntity = %[3]q \n", camelExport(schema), camelExport(t.GoName), tident))
	}
	b.WriteString(")")

	return b.String()
}

func (f *Funcs) generate_entity_fields(schema string, tables Tables) string {
	if schema == "public" {
		schema = ""
	}
	if f.currentDatabase != "gen_db" && f.currentDatabase != "xo_tests_db" {
		return ""
	}
	for _, t := range tables {
		for _, field := range t.Fields {
			f.field(field, "EntityFields", t)
		}
	}

	content, err := json.MarshalIndent(f.entityFields, "", "  ")
	if err != nil {
		panic("json.MarshalIndent: " + err.Error())
	}
	if schema == "" && f.currentDatabase == "gen_db" {
		outputPath := "entityFields.gen.json"
		if err := os.WriteFile(outputPath, content, 0o644); err != nil {
			panic("os.WriteFile: " + err.Error())
		}
	}

	return fmt.Sprintf(formatEntityFields(f.schemaPrefix, f.entityFields))
}

func formatEntityFields(schema string, entityFields map[string]map[string]DbField) string {
	var buf bytes.Buffer

	var entityTypes []string
	for entityType := range entityFields {
		entityTypes = append(entityTypes, entityType)
	}
	sort.Strings(entityTypes)

	buf.WriteString(fmt.Sprintf("var %[1]sEntityFields = map[%[1]sTableEntity]map[string]DbField {\n", camelExport(schema)))
	for _, entityType := range entityTypes {
		fields := entityFields[entityType]

		var fieldNames []string
		for fieldName := range fields {
			fieldNames = append(fieldNames, fieldName)
		}
		sort.Strings(fieldNames)

		buf.WriteString(fmt.Sprintf("\t%sTableEntity%s: {\n", camelExport(schema), camelExport(entityType)))
		for _, fieldName := range fieldNames {
			field := fields[fieldName]
			buf.WriteString(
				fmt.Sprintf("\t\t\"%s\": DbField{Type: %s, Db: \"%s\", Nullable: %t, Public: %t},\n",
					fieldName, simpleTypeToTypeEnum(field.Type), field.Db, field.Nullable, field.Public,
				))
		}
		buf.WriteString("\t},\n")
	}
	buf.WriteString("}\n")

	return buf.String()
}

func (f *Funcs) named(name, value string, out bool) string {
	switch {
	case out:
		return fmt.Sprintf("sql.Named(%q, sql.Out{Dest: %s})", name, value)
	}
	return fmt.Sprintf("sql.Named(%q, %s)", name, value)
}

func (f *Funcs) logf_pkeys(v any) string {
	p := []string{"sqlstr"}
	switch x := v.(type) {
	case Table:
		p = append(p, f.names(f.short(x.GoName)+".", x.PrimaryKeys))
	}
	return fmt.Sprintf("logf(%s)", strings.Join(p, ", "))
}

func (f *Funcs) logf(v any, ignore ...any) string {
	var ignoreNames []any
	p := []string{"sqlstr"}
	// build ignore list
	for i, x := range ignore {
		switch z := x.(type) {
		case string:
			ignoreNames = append(ignoreNames, z)
		case Field:
			ignoreNames = append(ignoreNames, z.GoName)
		case []Field:
			for _, f := range z {
				ignoreNames = append(ignoreNames, f.GoName)
			}
		default:
			return fmt.Sprintf("[[ UNSUPPORTED TYPE 11 (%d): %T ]]", i, x)
		}
	}
	// add fields
	switch x := v.(type) {
	case Table:
		p = append(p, f.names_ignore(f.short(x.GoName)+".", x, ignoreNames...))
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 12: %T ]]", v)
	}
	x := ""
	if len(p) > 0 {
		x = strings.Join(p, ", ")
	}
	return fmt.Sprintf("logf(%s)", x)
}

func (f *Funcs) logf_update(v any) string {
	var ignore []any
	p := []string{"sqlstr"}
	switch x := v.(type) {
	case Table:
		prefix := f.short(x.GoName) + "."
		for _, pk := range x.Generated {
			ignore = append(ignore, pk.GoName)
		}
		for _, pk := range x.PrimaryKeys {
			ignore = append(ignore, pk.GoName)
		}
		p = append(p, f.names_ignore(prefix, x, ignore...), f.names(prefix, x.PrimaryKeys))
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 13: %T ]]", v)
	}
	return fmt.Sprintf("logf(%s)", strings.Join(p, ", "))
}

type CursorPagination struct {
	Table  Table
	Fields []Field
}

// names generates a list of names.
func (f *Funcs) namesfn(all bool, prefix string, z ...any) string {
	var names []string
	for i, v := range z {
		switch x := v.(type) {
		case string:
			names = append(names, x)
		case Query:
			for _, p := range x.Params {
				if !all && p.Interpolate {
					continue
				}
				names = append(names, prefix+p.Name)
			}
		case []Field:
			for _, p := range x {
				var pre string
				if p.EnumSchema != "" && f.schemaPrefix != "public" {
					pre = camelExport(f.schemaPrefix)
				}
				names = append(names, prefix+pre+checkName(p.GoName))
			}
		case Proc:
			if params := f.params(x.Params, false, nil); params != "" {
				names = append(names, params)
			}
		case Index:
			names = append(names, f.params(x.Fields, false, nil))

			nn := ""
			if len(names[2:]) > 0 {
				nn = strings.Join(names[2:], ", ")
			}
			return "ctx, sqlstr, append([]any{" + nn + "}, append(filterParams, havingParams...)...)..."
		case CursorPagination:
			names = append(names, f.params(x.Fields, false, nil))

			return "ctx, sqlstr, append(filterParams, havingParams...)..."
		default:
			names = append(names, fmt.Sprintf("/* UNSUPPORTED TYPE 14 (%d): %T */", i, v))
		}
	}
	x := ""
	if len(names) > 0 {
		x = strings.Join(names, ", ")
	}
	return x
}

func (f *Funcs) initialize_constraints(t Table, constraints []Constraint) bool {
	var pkisfkc *TableForeignKey
	for _, f := range t.Fields {
		af := analyzeField(t, f)
		if af.PKisFK != nil {
			pkisfkc = af.PKisFK
		}
	}

	if _, ok := f.tableConstraints[t.SQLName]; !ok {
		if len(constraints) > 0 {
			f.loadConstraints(constraints, t.SQLName, pkisfkc)
		}
	}

	return true
}

// TODO: triplicated logic. preload everything in init for each table
func (f *Funcs) joinNames(t Table) []string {
	joinNames := []string{}

	for _, c := range f.tableConstraints[t.SQLName] {
		if c.Type != "foreign_key" {
			continue
		}
		var joinName string
		switch c.Cardinality {
		case M2M:
			lookupName := strings.TrimSuffix(c.ColumnName, "_id")
			joinName = camelExport(inflector.Pluralize(lookupName))
			rc := camelExport(strings.TrimSuffix(c.RefColumnName, "_id"))
			rlc := camelExport(strings.TrimSuffix(c.LookupRefColumnName, "_id"))
			lc := camelExport(strings.TrimSuffix(c.LookupColumnName, "_id"))
			col := camelExport(strings.TrimSuffix(c.ColumnName, "_id"))
			_, _, _, _ = rc, rlc, lc, col
			if rlc != lc && rlc != col {
				joinName = lc + joinName
			}
			// if rc != col && rlc != col {
			// 	joinName = col + joinName
			// }
		case M2O:
			if c.RefTableName != t.SQLName {
				continue
			}
			joinName = camelExport(c.TableName)
			if c.JoinTableClash {
				joinName = camelExport(c.ColumnName) + joinName
			}
		case O2O:
			if c.TableName == t.SQLName {
				joinName = camelExport(singularize(c.RefTableName))
				if c.JoinTableClash {
					joinName = joinName + camelExport(c.ColumnName)
				}
			}

		default:
		}
		if joinName == "" {
			continue
		}
		for _, name := range joinNames {
			if name == joinName {
				joinName = joinName + toAcronym(c.TableName)
			}
		}
		joinNames = append(joinNames, joinName)
	}

	return joinNames
}

// names generates a list of names (excluding certain ones such as interpolated
// names).
func (f *Funcs) names(prefix string, z ...any) string {
	return f.namesfn(false, prefix, z...)
}

// names_all generates a list of all names.
func (f *Funcs) names_all(prefix string, z ...any) string {
	return f.namesfn(true, prefix, z...)
}

// names_ignore generates a list of all names, ignoring fields that match the value in ignore.
func (f *Funcs) names_ignore(prefix string, v any, ignore ...any) string {
	m := make(map[string]bool)
	for _, v := range ignore {
		switch x := v.(type) {
		case string:
			m[x] = true
		case Field:
			m[x.GoName] = true
		}
	}

	var vals []Field
	switch x := v.(type) {
	case Table:
		for _, p := range x.Fields {
			if m[p.GoName] {
				continue
			}
			vals = append(vals, p)
		}
	case []Field:
		for _, p := range x {
			if m[p.GoName] {
				continue
			}
			vals = append(vals, p)
		}
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 15: %T ]]", v)
	}
	return f.namesfn(true, prefix, vals)
}

// querystr generates a querystr for the specified query and any accompanying
// comments.
func (f *Funcs) querystr(v any) string {
	var interpolate bool
	var query, comments []string
	switch x := v.(type) {
	case Query:
		interpolate, query, comments = x.Interpolate, x.Query, x.Comments
	default:
		return fmt.Sprintf("const sqlstr = [[ UNSUPPORTED TYPE 16: %T ]]", v)
	}
	typ := "const"
	if interpolate {
		typ = "var"
	}
	var lines []string
	for i := 0; i < len(query); i++ {
		line := "`" + query[i] + "`"
		if i != len(query)-1 {
			line += " + "
		}
		if s := strings.TrimSpace(comments[i]); s != "" {
			line += "// " + s
		}
		lines = append(lines, line)
	}
	sqlstr := stripRE.ReplaceAllString(strings.Join(lines, "\n"), " ")
	return fmt.Sprintf("%s sqlstr = %s", typ, sqlstr)
}

var stripRE = regexp.MustCompile(`\s+\+\s+` + "``")

func (f *Funcs) sqlstr(typ string, v any) string {
	var lines []string
	switch typ {
	case "soft_delete":
		lines = f.sqlstr_soft_delete(v)
	case "insert_manual":
		lines = f.sqlstr_insert_manual(v)
	case "insert":
		lines = f.sqlstr_insert(v)
	case "update":
		lines = f.sqlstr_update(v)
	case "upsert":
		lines = append(f.sqlstr_upsert(v), " RETURNING *")
	case "delete":
		lines = f.sqlstr_delete(v)
	case "proc":
		lines = f.sqlstr_proc(v)
	default:
		return fmt.Sprintf("const sqlstr = `UNKNOWN QUERY TYPE: %s`", typ)
	}
	return fmt.Sprintf("sqlstr := `%s `", strings.Join(lines, "\n\t"))
}

func isValidCursor(pk Field) bool {
	return pk.UnderlyingType == "time.Time" || pk.UnderlyingType == "int" || pk.UnderlyingType == "int64"
}

// cursor_columns returns a list of possible combinations of columns for cursor pagination
// (pk, unique field, ...).
func (f *Funcs) cursor_columns(table Table, constraints []Constraint, tables Tables) [][]Field {
	var cursorCols [][]Field
	var tableConstraints []Constraint
	if tc, ok := f.tableConstraints[table.SQLName]; ok {
		tableConstraints = tc
	}
	existingCursors := make(map[string]bool)
	allPKsAreValidCursor := len(table.PrimaryKeys) > 0
	for _, pk := range table.PrimaryKeys {
		if !isValidCursor(pk) {
			allPKsAreValidCursor = false
		}
	}
	if allPKsAreValidCursor {
		cursorCols = append(cursorCols, table.PrimaryKeys) // assume its incremental. if it's not then simply dont call it...
	}

	for _, z := range table.Fields {
		for _, c := range tableConstraints {
			if c.Type == "unique" && c.ColumnName == z.SQLName {
				if isValidCursor(z) {
					if existingCursors[z.SQLName] {
						continue
					}
					cursorCols = append(cursorCols, []Field{z})
					existingCursors[z.SQLName] = true
				}
			}
		}
	}

	return cursorCols
}

// sqlstr_paginated builds a cursor-paginated query string from columns.
func (f *Funcs) sqlstr_paginated(v any, tables Tables) string {
	var groupbys []string
	switch x := v.(type) {
	case Table:
		var fields []string

		// build table fieldnames
		for _, z := range x.Fields {
			// add current table fields
			fields = append(fields, x.SQLName+"."+f.colname(z))
		}
		// create joins for constraints
		funcs := template.FuncMap{
			"singularize": singularize,
		}

		// all fields already selected in main table need to appear
		groupbys = append(groupbys, mainGroupByFields(x, f, funcs)...)

		lines := []string{
			"SELECT ",
			strings.Join(fields, ",\n\t") + " %s ",
			" FROM " + f.schemafn(x.SQLName) + " %s ",
			" %s ",
		}

		buf := f.sqlstrBase(x)

		buf.WriteString(fmt.Sprintf("\nsqlstr := fmt.Sprintf(`%s %%s %%s %%s`, selects, joins, filters, groupByClause, havingClause, orderByClause)",
			strings.Join(lines, "\n\t"),
		))

		return buf.String()
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 26: %T ]]", v)
}

func (f *Funcs) sqlstrBase(x Table) *strings.Builder {
	buf := &strings.Builder{}

	buf.WriteString(`
		var selectClauses []string
		var joinClauses []string
		var groupByClauses []string
		`)
	for _, j := range f.joinNames(x) {
		buf.WriteString(fmt.Sprintf(`
			if c.joins.%[2]s {
				selectClauses = append(selectClauses, %[1]sTable%[2]sSelectSQL)
				joinClauses = append(joinClauses, %[1]sTable%[2]sJoinSQL)
				groupByClauses = append(groupByClauses, %[1]sTable%[2]sGroupBySQL)
			}
			`, f.lower_first(x.GoName), j))
	}
	buf.WriteString(`
		selects := ""
		if len(selectClauses) > 0 {
			selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
		}
		joins := strings.Join(joinClauses, " \n ") + " "
		groupByClause := ""
		if len(groupByClauses) > 0 {
			groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
		}
		`)

	return buf
}

func mainGroupByFields(x Table, f *Funcs, funcs template.FuncMap) []string {
	var mainGroupBys []string
	for _, z := range x.Fields {
		mainGroupBys = append(mainGroupBys, x.SQLName+"."+f.colname(z))
	}
	return mainGroupBys
}

// sqlstr_insert_base builds an INSERT query
// If not all, sequence columns are skipped.
func (f *Funcs) sqlstr_insert_base(all bool, v any) []string {
	switch x := v.(type) {
	case Table:
		// build names and values
		var n int
		var fields, vals []string
		for _, z := range x.Fields {
			if z.IsGenerated && !all || z.IsIgnored {
				continue
			}
			fields, vals = append(fields, f.colname(z)), append(vals, f.nth(n))
			n++
		}
		return []string{
			"INSERT INTO " + f.schemafn(x.SQLName) + " (",
			strings.Join(fields, ", "),
			") VALUES (",
			strings.Join(vals, ", "),
			")",
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 17: %T ]]", v)}
}

// sqlstr_insert_manual builds an INSERT query that inserts all fields.
func (f *Funcs) sqlstr_insert_manual(v any) []string {
	return append(f.sqlstr_insert_base(true, v), " RETURNING *")
}

// sqlstr_insert builds an INSERT query, skipping the sequence field with
// applicable RETURNING clause for generated primary key fields.
func (f *Funcs) sqlstr_insert(v any) []string {
	switch x := v.(type) {
	case Table:
		var generatedFields []string
		var count int
		for _, field := range x.Fields {
			if field.IsGenerated || field.IsIgnored {
				generatedFields = append(generatedFields, f.colname(field))
			} else {
				count++
			}
		}
		lines := f.sqlstr_insert_base(false, v)
		// add return clause
		switch f.driver {
		case "postgres":
			lines[len(lines)-1] += ` RETURNING *`
		}
		return lines
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 18: %T ]]", v)}
}

// sqlstr_update_base builds an UPDATE query, using primary key fields as the WHERE
// clause, adding prefix.
//
// When prefix is empty, the WHERE clause will be in the form of name = $1.
// When prefix is non-empty, the WHERE clause will be in the form of name = <PREFIX>name.
//
// Similarly, when prefix is empty, the table's name is added after UPDATE,
// otherwise it is omitted.
func (f *Funcs) sqlstr_update_base(prefix string, v any) (int, []string) {
	switch x := v.(type) {
	case Table:
		// build names and values
		var n int
		var list []string
		for _, z := range x.Fields {
			if z.IsPrimary || z.IsGenerated || z.IsIgnored {
				continue
			}
			name, param := f.colname(z), f.nth(n)
			if prefix != "" {
				param = prefix + name
			}
			list = append(list, fmt.Sprintf("%s = %s", name, param))
			n++
		}
		name := ""
		if prefix == "" {
			name = f.schemafn(x.SQLName) + " "
		}
		return n, []string{
			"UPDATE " + name + "SET ",
			strings.Join(list, ", ") + " ",
		}
	}
	return 0, []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 19: %T ]]", v)}
}

// sqlstr_update builds an UPDATE query, using primary key fields as the WHERE
// clause.
func (f *Funcs) sqlstr_update(v any) []string {
	// build pkey vals
	switch x := v.(type) {
	case Table:
		var conditions []string
		n, lines := f.sqlstr_update_base("", v)
		for i, z := range x.PrimaryKeys {
			conditions = append(conditions, fmt.Sprintf("%s = %s ", f.colname(z), f.nth(n+i)))
		}

		return append(lines, "WHERE "+strings.Join(conditions, " AND "), "RETURNING *")
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 20: %T ]]", v)}
}

func (f *Funcs) sqlstr_upsert(v any) []string {
	switch x := v.(type) {
	case Table:
		// build insert
		lines := f.sqlstr_insert_base(true, x)
		switch f.driver {
		case "postgres", "sqlite3":
			return append(lines, f.sqlstr_upsert_postgres_sqlite(x)...)
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 21 %s: %T ]]", f.driver, v)}
}

// sqlstr_upsert_postgres_sqlite builds an uspert query for postgres and sqlite
//
// INSERT (..) VALUES (..) ON CONFLICT DO UPDATE SET ...
func (f *Funcs) sqlstr_upsert_postgres_sqlite(v any) []string {
	switch x := v.(type) {
	case Table:
		// add conflict and update
		var conflicts []string
		for _, f := range x.PrimaryKeys {
			conflicts = append(conflicts, f.SQLName)
		}
		lines := []string{" ON CONFLICT (" + strings.Join(conflicts, ", ") + ") DO "}
		_, update := f.sqlstr_update_base("EXCLUDED.", v)
		return append(lines, update...)
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 22: %T ]]", v)}
}

// sqlstr_delete builds a DELETE query for the primary keys.
func (f *Funcs) sqlstr_delete(v any) []string {
	switch x := v.(type) {
	case Table:
		// names and values
		var list []string
		for i, z := range x.PrimaryKeys {
			list = append(list, fmt.Sprintf("%s = %s", f.colname(z), f.nth(i)))
		}
		return []string{
			"DELETE FROM " + f.schemafn(x.SQLName) + " ",
			"WHERE " + strings.Join(list, " AND "),
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 25: %T ]]", v)}
}

// sqlstr_soft_delete builds a soft DELETE query for the primary keys.
func (f *Funcs) sqlstr_soft_delete(v any) []string {
	switch x := v.(type) {
	case Table:
		// names and values
		var list []string
		for i, z := range x.PrimaryKeys {
			list = append(list, fmt.Sprintf("%s = %s", f.colname(z), f.nth(i)))
		}
		return []string{
			"UPDATE " + f.schemafn(x.SQLName) + " ",
			"SET deleted_at = NOW() ",
			"WHERE " + strings.Join(list, " AND "),
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 25: %T ]]", v)}
}

// M2MSelect = `(case when {{.Nth}}::boolean = true then array_agg(xo_join_{{.JoinTable}}.{{.JoinTable}}) filter (where xo_join_teams.teams is not null) end) as {{.JoinTable}}`

const (
	M2MSelect = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_{{.LookupJoinTablePKSuffix}}{{.ClashSuffix}}.__{{.LookupJoinTablePKAgg}}
		{{- range .LookupExtraCols }}
		, xo_join_{{$.LookupJoinTablePKSuffix}}{{$.ClashSuffix}}.{{ . -}}
		{{- end }}
		)) filter (where xo_join_{{.LookupJoinTablePKSuffix}}{{.ClashSuffix}}.__{{.LookupJoinTablePKAgg}}_{{.JoinTablePK}} is not null), '{}') as {{.LookupJoinTablePKSuffix}}{{.ClashSuffix}}`
	// TODO: xo_tests both empty m2o and >1 joined array
	M2OSelect = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_{{.JoinTable}}{{.ClashSuffix}}.__{{.JoinTable}})) filter (where xo_join_{{.JoinTable}}{{.ClashSuffix}}.{{.JoinTable}}_{{.JoinRefColumn}} is not null), '{}') as {{.JoinTable}}{{.ClashSuffix}}`
	// extra check needed to prevent pgx from trying to scan a record with NULL values into the ???Join struct
	O2OSelect = `(case when {{ .Alias}}_{{.JoinTableAlias}}.{{.JoinColumn}} is not null then row({{ .Alias}}_{{.JoinTableAlias}}.*) end) as {{ singularize .JoinTable}}_{{ singularize .JoinTableAlias}}`
)

const (
	M2MGroupBy = `{{.CurrentTable}}.{{.LookupRefColumn}}, {{.CurrentTablePKGroupBys}}`
	M2OGroupBy = `{{.CurrentTablePKGroupBys}}`
	O2OGroupBy = `{{ .Alias}}_{{.JoinTableAlias}}.{{.JoinColumn}},
	{{- range $g := .JoinTablePKGroupBys}}
      {{if $g}}{{$g}},{{end}}

  {{- end}}
	{{.CurrentTablePKGroupBys}}`
)

const (
	// need group by in all
	M2MJoin = `
left join (
	select
		{{.LookupTable}}.{{.LookupColumn}} as {{.LookupTable}}_{{.LookupColumn}}
		{{- range .LookupExtraCols }}
		, {{$.LookupTable}}.{{.}} as {{ . -}}
		{{- end }}
		, {{.JoinTable}}.{{.JoinTablePK}} as __{{.LookupJoinTablePKAgg}}_{{.JoinTablePK}}
		, row({{.JoinTable}}.*) as __{{.LookupJoinTablePKAgg}}
	from
		{{.Schema}}{{.LookupTable}}
	join {{.Schema}}{{.JoinTable}} on {{.JoinTable}}.{{.JoinTablePK}} = {{.LookupTable}}.{{.LookupJoinTablePK}}
	group by
		{{.LookupTable}}_{{.LookupColumn}}
		, {{.JoinTable}}.{{.JoinTablePK}}
		{{- range .LookupExtraCols }}
		, {{ . -}}
		{{- end }}
) as xo_join_{{.LookupJoinTablePKSuffix}}{{.ClashSuffix}} on xo_join_{{.LookupJoinTablePKSuffix}}{{.ClashSuffix}}.{{.LookupTable}}_{{.LookupColumn}} = {{.CurrentTable}}.{{.LookupRefColumn}}
`
	M2OJoin = `
left join (
  select
  {{.JoinColumn}} as {{.JoinTable}}_{{.JoinRefColumn}}
    , row({{.JoinTable}}.*) as __{{.JoinTable}}
  from
    {{.Schema}}{{.JoinTable}}
  group by
	  {{.JoinTable}}_{{.JoinRefColumn}}, {{.Schema}}{{.JoinTable}}.{{.JoinTablePKName}}
) as xo_join_{{.JoinTable}}{{.ClashSuffix}} on xo_join_{{.JoinTable}}{{.ClashSuffix}}.{{.JoinTable}}_{{.JoinRefColumn}} = {{.CurrentTable}}.{{.JoinRefColumn}}
`
	O2OJoin = `
left join {{.Schema}}{{.JoinTable}} as {{ .Alias}}_{{.JoinTableAlias}} on {{ .Alias}}_{{.JoinTableAlias}}.{{.JoinColumn}} = {{.CurrentTable}}.{{.JoinRefColumn}}
`
)

// sqlstr_index builds a index fields.
func (f *Funcs) sqlstr_index(v any, tables Tables) string {
	switch x := v.(type) {
	case Index:
		var filters, fields []string

		var tableHasDeletedAt bool
		for _, field := range x.Table.Fields {
			if field.SQLName == "deleted_at" {
				tableHasDeletedAt = true
			}
		}

		// build table fieldnames
		for _, z := range x.Table.Fields {
			// add current table fields
			fields = append(fields, x.Table.SQLName+"."+f.colname(z))
		}

		var n int
		// index fields
		for _, z := range x.Fields {
			filters = append(filters, fmt.Sprintf("%s.%s = %s", x.Table.SQLName, f.colname(z), f.nth(n)))
			n++
		}
		// generate filters if we are generating a subset query from multicol index
		// e.g.
		// 	create unique index on kanban_steps (project_id , name , step_order)
		// 	where
		// 		step_order is not null;

		// 	create unique index on kanban_steps (project_id , name)
		// 	where
		// 		step_order is null;
		//
		// in this case, func names need to be eg KanbanStepByName_StepOrderNotNull (via first index) and KanbanStepByName_StepOrderNull (2nd) so
		// that we dont skip generation due to `emittedIndexes`
		// a hypothetical KanbanStepByName without index conditions will not be generated since it will not use index scan without it
		// (unless a explicit name index is created obviously)
		if _, after, ok := strings.Cut(x.Definition, " WHERE "); ok { // index def is normalized in db
			filters = append(filters, after)
		}

		ff := "true"
		if len(filters) > 0 {
			ff = strings.Join(filters, " AND ")
		}

		lines := []string{
			"SELECT ",
			strings.Join(fields, ",\n\t") + " %s ",
			" FROM " + f.schemafn(x.Table.SQLName) + " %s ",
			" WHERE " + ff,
			" %s ", // deleted at, etc.
		}

		groupbyClause := " %s \n"
		havingClause := " %s \n"

		buf := f.sqlstrBase(x.Table)

		if tableHasDeletedAt {
			buf.WriteString(fmt.Sprintf("\nsqlstr := fmt.Sprintf(`%s %s %s %s`, selects, joins, filters, c.deletedAt, groupByClause, havingClause)",
				strings.Join(lines, "\n\t"),
				fmt.Sprintf(" AND %s.deleted_at is %%s", x.Table.SQLName),
				groupbyClause,
				havingClause,
			))
		} else {
			buf.WriteString(fmt.Sprintf("\nsqlstr := fmt.Sprintf(`%s %s %s`, selects, joins, filters, groupByClause, havingClause)",
				strings.Join(lines, "\n\t"),
				groupbyClause,
				havingClause,
			))
		}

		return buf.String()
	}
	return fmt.Sprintf("[[ UNSUPPORTED TYPE 26: %T ]]", v)
}

// loadConstraints saves possible joins for a table based on constraints to tableConstraints
// if pk is fk for the current table, its foreignkey is required to generate shared refs (if property flag set)
func (f *Funcs) loadConstraints(cc []Constraint, table string, pkIsFK *TableForeignKey) {
	if _, ok := f.tableConstraints[table]; ok {
		// fmt.Printf("Constraints for %s:\n%v\n", table, formatJSON(f.tableConstraints[table]))
		return // don't duplicate
	}
	mustShareRefs := false
outer:
	for _, constraint := range cc {
		// we need unique constraints for paginated query generation. instead do this check when generating joins only
		// if constraint.Type != "foreign_key" {
		// 	continue
		// }
		// TODO: need Table instead of just string, since foreign keys required for shared ref constraints
		// do not exist, therefore we cannt find the ref table just by prmary key
		// TODO: ignore seen constraints
		for _, seenConstraint := range f.tableConstraints[table] {
			if sameConstraint(seenConstraint, constraint) {
				continue outer
			}
		}
		if constraint.Cardinality == M2M && constraint.RefTableName == table {
			for _, c1 := range cc {
				if c1.TableName == constraint.TableName && c1.ColumnName != constraint.ColumnName && c1.Type == "foreign_key" {
					c1.LookupColumnName = constraint.ColumnName
					c1.LookupColumnComment = constraint.ColumnComment
					c1.LookupRefColumnName = constraint.RefColumnName
					c1.LookupRefColumnComment = constraint.RefColumnComment
					f.tableConstraints[table] = append(f.tableConstraints[table], c1)
				}
			}
		} else if constraint.Cardinality == O2O && (constraint.TableName == table || constraint.RefTableName == table) {
			f.tableConstraints[table] = append(f.tableConstraints[table], constraint)
		} else if constraint.RefTableName == table {
			f.tableConstraints[table] = append(f.tableConstraints[table], constraint)
		}

		if constraint.TableName == table {
			annotations, err := parseAnnotations(constraint.ColumnComment)
			if err != nil {
				panic(fmt.Sprintf("parseAnnotations: %v", err))
			}

			properties := extractPropertiesAnnotation(annotations[propertiesAnnot])

			shareRefConstraints := contains(properties, propertyShareRefConstraints)
			if shareRefConstraints && pkIsFK != nil {
				mustShareRefs = true
			}
		}

	}

	if mustShareRefs {
		f.loadConstraints(cc, pkIsFK.RefTable, nil)
		refConstraints := f.tableConstraints[pkIsFK.RefTable]
		var newRefConstraints []Constraint
		for _, c := range refConstraints {
			if c.Type != "foreign_key" {
				continue
			}
			newr := c
			switch newr.Cardinality {
			case M2O:
				newr.RefTableName = table
				newr.RefColumnName = c.ColumnName
			case M2M:
				// works as is
			}
			newr.Name = newr.Name + "-shared-ref-" + table
			newRefConstraints = append(newRefConstraints, newr)
		}
		f.tableConstraints[table] = append(f.tableConstraints[table], newRefConstraints...)
	}
}

// sameConstraint returns whether constraint b is equivalent to a.
// accepts Constraint or xo.Constraint.
// NOTE: xo.Constraint requires external cardinality check.
func sameConstraint(a Constraint, b any) bool {
	switch x := b.(type) {
	case xo.Constraint:
		return a.TableName == x.TableName &&
			a.RefTableName == x.RefTableName &&
			a.ColumnName == x.ColumnName &&
			a.RefColumnName == x.RefColumnName &&
			a.Type == x.Type
	case Constraint:
		return a.TableName == x.TableName &&
			a.RefTableName == x.RefTableName &&
			a.ColumnName == x.ColumnName &&
			a.RefColumnName == x.RefColumnName &&
			a.Type == x.Type &&
			a.Cardinality == x.Cardinality
	default:
		panic(fmt.Sprintf("unknown constraint type: %T", b))
	}
}

// createJoinStatement returns select queries and join statements strings
// for a given index table.
func (f *Funcs) createJoinStatement(tables Tables, c Constraint, table Table, funcs template.FuncMap) (string, string, string) {
	var joinTpl, selectTpl, groupbyTpl string
	join := &bytes.Buffer{}
	selec := &bytes.Buffer{}
	groupby := &bytes.Buffer{}
	params := make(map[string]any)
	fmt.Fprintf(join, "-- %s join generated from %q", c.Cardinality, c.Name)

	params["ClashSuffix"] = ""
	params["Schema"] = ""

	var currentTablePKGroupBys []string
	for _, pk := range table.PrimaryKeys {
		currentTablePKGroupBys = append(currentTablePKGroupBys, table.SQLName+"."+pk.SQLName)
	}
	params["CurrentTablePKGroupBys"] = strings.Join(currentTablePKGroupBys, ", ")
	params["Alias"] = "" // to prevent alias name clashes

	if table.Schema != "public" {
		params["Schema"] = table.Schema + "."
	}

	switch c.Cardinality {
	case M2M:
		joinTpl = M2MJoin
		selectTpl = M2MSelect
		groupbyTpl = M2MGroupBy

		lookupName := strings.TrimSuffix(c.ColumnName, "_id")

		params["LookupColumn"] = c.LookupColumnName
		params["JoinTable"] = c.RefTableName
		params["LookupRefColumn"] = c.LookupRefColumnName
		params["JoinTablePK"] = c.RefColumnName
		params["LookupJoinTablePK"] = c.ColumnName
		params["LookupJoinTablePKAgg"] = c.RefTableName
		params["LookupJoinTablePKSuffix"] = c.TableName + "_" + inflector.Pluralize(lookupName)
		params["CurrentTable"] = table.SQLName
		params["LookupTable"] = c.TableName
		params["LookupExtraCols"] = []string{}

		if c.JoinTableClash {
			// lc := strings.TrimSuffix(c.LookupColumn, "_id")
			// params["ClashSuffix"] = "_" + lc
		}
		lookupTable := tables[c.TableName]
		m2mExtraCols := getTableRegularFields(lookupTable)

		if len(m2mExtraCols) > 0 {
			// we are not changing lookup name, i.e. we're using the pk name, e.g. user_id instead of something like member, leader, sender...
			// don't rename to avoid confusion:
			params["LookupJoinTablePK"] = c.ColumnName
			params["LookupJoinTablePKAgg"] = inflector.Pluralize(c.ColumnName)
			// params["LookupJoinTablePKSuffix"] = inflector.Pluralize(strings.TrimSuffix(c.ColumnName, "_id"))

			colNames := make([]string, len(m2mExtraCols))
			for i, col := range m2mExtraCols {
				colNames[i] = col.SQLName
			}
			params["LookupExtraCols"] = colNames

			if len(m2mExtraCols) > 0 {
				// in this case we will create an extra struct that holds the joined table + these extra fields.
				// need to change agg name to avoid confusion
				params["LookupJoinTablePKAgg"] = params["JoinTable"]
			}
		}
	case M2O:
		// TODO: must be the same as M2M, just without lookup table join
		joinTpl = M2OJoin
		selectTpl = M2OSelect
		groupbyTpl = M2OGroupBy
		if c.RefTableName == table.SQLName {
			for _, t := range tables {
				if t.SQLName == c.TableName && t.Schema == f.schema {
					params["JoinTablePKName"] = t.PrimaryKeys[0].SQLName
				}
			}
			params["JoinColumn"] = c.ColumnName
			params["JoinTable"] = c.TableName
			params["JoinRefColumn"] = c.RefColumnName
			params["CurrentTable"] = table.SQLName

			if c.JoinTableClash {
				params["ClashSuffix"] = "_" + c.ColumnName
			}

		}

		// joinTable := tables[params["JoinTable"].(string)]
		// var joinTablePKGroupBys []string
		// for _, pk := range joinTable.PrimaryKeys {
		// 	joinTablePKGroupBys = append(joinTablePKGroupBys, params["JoinTable"].(string)+"."+pk.SQLName)
		// }
		// params["JoinTablePKGroupBys"] = strings.Join(joinTablePKGroupBys, ", ")

	case O2O:
		if c.TableName == table.SQLName {
			groupbyTpl = O2OGroupBy
			joinTpl = O2OJoin
			selectTpl = O2OSelect
			params["JoinColumn"] = c.RefColumnName
			params["JoinTable"] = c.RefTableName
			params["JoinRefColumn"] = c.ColumnName
			params["JoinTableAlias"] = c.ColumnName
			params["CurrentTable"] = c.TableName
			if c.JoinTableClash {
				params["ClashSuffix"] = "_" + c.RefColumnName
			}

			t := tables[c.RefTableName]
			var field Field
			for _, tf := range t.Fields {
				if tf.SQLName == c.RefColumnName {
					field = tf
				}
			}
			// need to check RefTable PKs since this should get called when generating for a
			// table that has *referenced* O2O where PK is FK. e.g. work_item gen -> we see demo_work_item has work_item_id PK that is FK.
			// viceversa we don't care as it's a regular PK.
			params["Alias"] = "_" + c.TableName
			af := analyzeField(t, field)
			if af.PKisFK != nil || c.PKisFK {
				params["Alias"] = "_" + c.RefTableName
				params["JoinTableAlias"] = c.ColumnName
			}

			joinTable := tables[c.RefTableName]
			var joinTablePKGroupBys []string
			for _, pk := range joinTable.PrimaryKeys {
				if af.PKisFK == nil {
					gb := params["Alias"].(string) + "_" + params["JoinTableAlias"].(string) + "." + pk.SQLName
					joinTablePKGroupBys = append(joinTablePKGroupBys, gb)
				}
			}

			params["JoinTablePKGroupBys"] = uniqueSort(joinTablePKGroupBys)

			break
		}

		// dummy created automatically to avoid this duplication
		// if c.RefTableName == table.SQLName {
		// 	joinTpl = O2OJoin
		// 	selectTpl = O2OSelect
		// 	params["JoinColumn"] = c.ColumnName
		// 	params["JoinTable"] = c.TableName
		// 	params["JoinRefColumn"] = c.RefColumnName
		// 	params["CurrentTable"] = table.SQLName
		// 	break
		// }

	default:
	}

	t := template.Must(template.New("").Option("missingkey=zero").Funcs(funcs).Parse(joinTpl))
	if err := t.Execute(join, params); err != nil {
		panic(fmt.Sprintf("could not execute join template: %s", err))
	}

	t = template.Must(template.New("").Option("missingkey=zero").Funcs(funcs).Parse(selectTpl))
	if err := t.Execute(selec, params); err != nil {
		panic(fmt.Sprintf("could not execute selec template: %s", err))
	}

	t = template.Must(template.New("").Option("missingkey=zero").Funcs(funcs).Parse(groupbyTpl))
	if err := t.Execute(groupby, params); err != nil {
		panic(fmt.Sprintf("could not execute groupby template: %s", err))
	}

	return join.String(), selec.String(), groupby.String()
}

// getTableRegularFields gets extra columns in a lookup table that are not PK or FK
func getTableRegularFields(t Table) []Field {
	ltFieldsMap := make(map[string]Field)
	for _, f := range t.Fields {
		ltFieldsMap[f.SQLName] = f
	}
	for _, tfk := range t.ForeignKeys {
		for _, fk := range tfk.FieldNames {
			delete(ltFieldsMap, fk)
		}
	}
	for _, pk := range t.PrimaryKeys {
		delete(ltFieldsMap, pk.SQLName)
	}
	m2mExtraCols := make([]Field, 0, len(ltFieldsMap))
	for _, v := range ltFieldsMap {
		m2mExtraCols = append(m2mExtraCols, v)
	}

	return m2mExtraCols
}

// sqlstr_proc builds a stored procedure call.
func (f *Funcs) sqlstr_proc(v any) []string {
	switch x := v.(type) {
	case Proc:
		if x.Type == "function" {
			return f.sqlstr_func(v)
		}
		// sql string format
		var format string
		switch f.driver {
		case "postgres":
			format = "CALL %s(%s)"
		}
		// build params list; add return fields for orcle
		l := x.Params
		var list []string
		for i := range l {
			s := f.nth(i)
			list = append(list, s)
		}
		name := f.schemafn(x.SQLName)
		return []string{
			fmt.Sprintf(format, name, strings.Join(list, ", ")),
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 27: %T ]]", v)}
}

func (f *Funcs) sqlstr_func(v any) []string {
	switch x := v.(type) {
	case Proc:
		var format string
		switch f.driver {
		case "postgres":
			format = "SELECT * FROM %s(%s)"
		}
		var list []string
		l := x.Params
		for i := range l {
			list = append(list, f.nth(i))
		}
		return []string{
			fmt.Sprintf(format, f.schemafn(x.SQLName), strings.Join(list, ", ")),
		}
	}
	return []string{fmt.Sprintf("[[ UNSUPPORTED TYPE 28: %T ]]", v)}
}

// convertTypes generates the conversions to convert the foreign key field
// types to their respective referenced field types.
func (f *Funcs) convertTypes(fkey ForeignKey) string {
	var p []string
	for i := range fkey.Fields {
		field := fkey.Fields[i]
		refField := fkey.RefFields[i]
		expr := f.short(fkey.Table) + "." + field.GoName
		// types match, can match
		if field.Type == refField.Type {
			p = append(p, expr)
			continue
		}
		// convert types
		typ, refType := field.Type, refField.Type
		if strings.HasPrefix(typ, "*") {
			_typ := typ[1:]
			expr = "*" + expr
			typ = strings.ToLower(_typ)
		}
		if strings.ToLower(refType) != typ {
			expr = refType + "(" + expr + ")"
		}
		p = append(p, expr)
	}

	x := ""
	if len(p) > 0 {
		x = strings.Join(p, ", ")
	}
	return x
}

// params converts a list of fields into their named Go parameters, skipping
// any Field with Name contained in ignore. addType will cause the go Type to
// be added after each variable name. addPrefix will cause the returned string
// to be prefixed with ", " if the generated string is not empty.
//
// Any field name encountered will be checked against goReservedNames, and will
// have its name substituted by its corresponding looked up value.
//
// Used to present a comma separated list of Go variable names for use with as
// either a Go func parameter list, or in a call to another Go func.
// (ie, ", a, b, c, ..." or ", a T1, b T2, c T3, ...").
func (f *Funcs) params(fields []Field, addType bool, w any) string {
	var vals []string
	for _, field := range fields {
		var t Table
		switch x := w.(type) {
		case Table:
			t = x
		default:
			t = Table{}
		}
		vals = append(vals, f.param(field, addType, &t))
	}
	if len(vals) == 0 {
		return ""
	}
	return strings.Join(vals, ", ")
}

func (f *Funcs) param(field Field, addType bool, table *Table) string {
	n := strings.Split(snaker.CamelToSnake(field.GoName), "_")
	s := strings.ToLower(n[0]) + field.GoName[len(n[0]):]
	// check go reserved names
	if r, ok := goReservedNames[strings.ToLower(s)]; ok {
		s = r
	}

	// add the go type
	if addType {
		if table != nil {
			af := analyzeField(*table, field)
			if af.IsFK {
				for _, c := range f.tableConstraints[table.SQLName] {
					if c.Type != "foreign_key" {
						continue
					}
					switch c.Cardinality {
					case M2M:
						if c.ColumnName == field.SQLName {
							field.Type = camelExport(f.schemaPrefix) + camelExport(singularize(c.RefTableName)) + "ID"
							break
						}
					case M2O:
						if c.TableName == table.SQLName && c.ColumnName == field.SQLName {
							field.Type = camelExport(f.schemaPrefix) + camelExport(c.RefTableName) + "ID"
							break
						}
					case O2O:
						if c.TableName == table.SQLName && c.ColumnName == field.SQLName {
							field.Type = camelExport(f.schemaPrefix) + camelExport(singularize(c.RefTableName)) + "ID"
							break
						}
					default:
					}
				}
			} else if field.IsPrimary {
				field.Type = table.GoName + "ID"
			}

			if strings.HasPrefix(field.Type, "Null") && f.schemaPrefix != "public" && field.EnumSchema != "public" {
				field.Type = "*" + strings.TrimPrefix(field.Type, "Null")
			}

			s += " " + f.typefn(field.Type)

		}
	}

	// if table.SQLName == "cache__demo_work_items" {
	// 	fmt.Printf("s: %v\n", s)
	// }

	// add to vals
	return s
}

// zero generates a zero list.
func (f *Funcs) zero(z ...any) string {
	var zeroes []string
	for i, v := range z {
		switch x := v.(type) {
		case string:
			zeroes = append(zeroes, x)
		case Table:
			for _, p := range x.Fields {
				zeroes = append(zeroes, f.zero(p))
			}
		case []Field:
			for _, p := range x {
				zeroes = append(zeroes, f.zero(p))
			}
		case Field:
			if _, ok := f.knownTypes[x.Type]; ok || x.Zero == "nil" {
				zeroes = append(zeroes, x.Zero)
				break
			}
			zeroes = append(zeroes, f.typefn(x.Type)+"{}")
		default:
			zeroes = append(zeroes, fmt.Sprintf("/* UNSUPPORTED TYPE 29 (%d): %T */", i, v))
		}
	}
	return strings.Join(zeroes, ", ")
}

// typefn generates the Go type, prefixing the custom package name if applicable.
func (f *Funcs) typefn(typ string) string {
	if strings.Contains(typ, ".") {
		return typ
	}
	var prefix string
	for strings.HasPrefix(typ, "[]") {
		typ = typ[2:]
		prefix += "[]"
	}
	if _, ok := f.knownTypes[typ]; ok || f.custom == "" {
		return prefix + typ
	}
	return prefix + f.custom + "." + typ
}

// field generates a field definition for a struct.
func (f *Funcs) field(field Field, mode string, table Table) (string, error) {
	buf := new(bytes.Buffer)
	hidden := false
	isPrivate := contains(field.Properties, propertyJSONPrivate)
	// ignoreConstraints := contains(field.Properties, propertyIgnoreConstraints)
	notRequired := contains(field.Properties, propertyOpenAPINotRequired)
	isPointer := strings.HasPrefix(field.Type, "*")
	af := analyzeField(table, field)
	skipField := field.IsGenerated || field.IsIgnored || field.SQLName == "deleted_at" //|| contains(table.ForeignKeys, field.SQLName)
	ignoreJson := isPrivate

	var skipExtraTags bool
	switch mode {
	case "CreateParams":
		hidden = contains(field.Properties, propertyOpenAPIHidden)
		if af.PKisFK != nil {
			ignoreJson = true // need for repo but unknown for request
		}
		if skipField {
			return "", nil
		}
		skipExtraTags = true
	case "UpdateParams":
		hidden = contains(field.Properties, propertyOpenAPIHidden)
		notRequired = true // PATCH, all optional
		if skipField {
			return "", nil
		}
		if af.PKisFK != nil { // e.g. workitemid in project tables. don't ever want to update it. PK is FK
			fmt.Printf("UpdateParams: skipping %q: is a single foreign and primary key in table %q\n", field.SQLName, table.SQLName)
			return "", nil
		}
		skipExtraTags = true
	case "Table":
	}

	nullable := true
	if !isPointer || strings.HasPrefix(field.Type, "map") || strings.HasPrefix(field.Type, "[]") {
		nullable = false
	}

	if err := f.fieldtag.Funcs(f.FuncMap()).Execute(buf, map[string]any{
		"field":         field,
		"ignoreJSON":    ignoreJson,
		"required":      !isPointer && !isPrivate && (!notRequired || !skipExtraTags),
		"skipExtraTags": skipExtraTags,
		"nullable":      nullable,
		"hidden":        hidden,
	}); err != nil {
		return "", err
	}
	var tag string
	if s := buf.String(); s != "" {
		tag = " `" + s + "`"
	}
	fieldType := f.typefn(field.Type)
	if field.IsPrimary {
		field.OpenAPISchema = field.Type
		if mode != "IDTypes" {
			pc := strings.Count(fieldType, "*")
			fieldType = strings.Repeat("*", pc) + table.GoName + "ID"
		}
	}

	// TODO: its not skipping the current field constraints,
	// but making sure other tables ignore constraints
	// therefore we need to have RefColumnField and LookupRefField
	// so we can keep this constraint is set or not.
	// ignoreConstraints := contains(field.Properties, propertyIgnoreConstraints)

	var constraintTyp string
	if af.IsFK {
		for _, c := range f.tableConstraints[table.SQLName] {
			if c.Type != "foreign_key" {
				continue
			}
			switch c.Cardinality {
			case M2M:
				if mode == "IDTypes" {
					return "", nil
				}
				if c.ColumnName == field.SQLName {
					constraintTyp = camelExport(f.schemaPrefix) + camelExport(singularize(c.RefTableName)) + "ID"
					break
				}
			case M2O:
				if c.RefTableName == table.SQLName && c.RefColumnName == field.SQLName {
					constraintTyp = camelExport(f.schemaPrefix) + camelExport(c.TableName) + "ID"
					break
				}
			case O2O:
				if c.TableName == table.SQLName && c.ColumnName == field.SQLName {
					constraintTyp = camelExport(f.schemaPrefix) + camelExport(singularize(c.RefTableName)) + "ID"
					break
				}

			default:
			}
		}
	}
	if constraintTyp != "" && mode != "IDTypes" {
		pc := strings.Count(fieldType, "*")
		fieldType = strings.Repeat("*", pc) + constraintTyp
	}

	if mode == "IDTypes" {
		if af.PKisFK != nil {
			return "", nil
		}
		if field.IsPrimary {
			goName := table.GoName + "ID"
			if strings.HasSuffix(fieldType, "uuid.UUID") {
				fieldType = "struct {\n	uuid.UUID \n}\n" +
					fmt.Sprintf("func New%[1]s(id uuid.UUID) %[1]s {\n return %[1]s{\n UUID: id,\n}\n }\n", goName)
			}

			return fmt.Sprintf("type %s %s\n\n", goName, fieldType), nil
		} else {
			return "", nil
		}
	}

	if mode != "IDTypes" {
		if af.PKisFK != nil {
			for _, tfk := range table.ForeignKeys {
				if len(tfk.FieldNames) == 1 && tfk.FieldNames[0] == field.SQLName {
					fieldType = camelExport(f.schemaPrefix) + camelExport(singularize(tfk.RefTable)) + "ID"
					break
				}
			}
		}
	}

	goName := field.GoName
	referencesCustomSchemaEnum := field.EnumSchema != "" && field.EnumSchema != "public"
	if strings.HasPrefix(fieldType, "Null") && f.schemaPrefix != "public" && referencesCustomSchemaEnum {
		fieldType = "*" + camelExport(f.schemaPrefix) + strings.TrimPrefix(fieldType, "Null")
		goName = camelExport(f.schemaPrefix) + goName
	} else if referencesCustomSchemaEnum {
		fieldType = camelExport(table.Schema) + fieldType // assumes no pointers
		goName = camelExport(f.schemaPrefix) + goName
		// fieldType = camelExport(f.schemaPrefix) + fieldType
	}

	if mode == "UpdateParams" {
		fieldType = "*" + fieldType // we do want **<field> and *<field>
	}

	if mode == "ParamsInterface" {
		if skipField || af.PKisFK != nil {
			return "", nil
		}
		return fmt.Sprintf("Get%[1]s() *%[2]s \n", goName, strings.TrimPrefix(fieldType, "*")), nil
	}

	// call initializer for all tables in extra.xo.go
	if mode == "EntityFields" {
		// mantine react table EMPTY and NOT EMPTY special filterModes - ie, enabling those Fields if nullable)
		entityName := camel(table.GoName)
		if _, ok := f.entityFields[entityName]; !ok {
			f.entityFields[entityName] = make(map[string]DbField)
		}
		if typ, err := goTypeToSimpleType(field.UnderlyingType); err == nil {
			f.entityFields[entityName][camel(field.GoName)] = DbField{
				Db:       field.SQLName,
				Type:     typ,
				Nullable: nullable,
				Public:   !hidden,
			}
		}
	}

	if mode == "ParamsGetter" {
		if skipField || af.PKisFK != nil {
			return "", nil
		}
		if strings.Count(fieldType, "*") == 0 {
			return fmt.Sprintf(`
			func (p %[1]sCreateParams) Get%[2]s() *%[3]s {
				x := p.%[2]s
				return &x
			}
			func (p %[1]sUpdateParams) Get%[2]s() *%[3]s {
				return p.%[2]s
			}
			`, table.GoName, goName, strings.TrimPrefix(fieldType, "*")), nil
		} else {
			return fmt.Sprintf(`
			func (p %[1]sCreateParams) Get%[2]s() *%[3]s {
				return p.%[2]s
			}
			func (p %[1]sUpdateParams) Get%[2]s() *%[3]s {
				if p.%[2]s != nil {
					return *p.%[2]s
				}
				return nil
			}
			`, table.GoName, goName, strings.TrimPrefix(fieldType, "*")), nil
		}
	}

	// TODO: if mode paramsInterface or paramsGetters just return methods, not struct fields,
	// to prevent duplicating logic above.

	return fmt.Sprintf("\t%s %s%s // %s\n", goName, fieldType, tag, field.SQLName), nil
}

func simpleTypeToTypeEnum(simpleType string) string {
	switch simpleType {
	case "date-time":
		return "ColumnSimpleTypeDateTime"
	case "integer":
		return "ColumnSimpleTypeInteger"
	case "number":
		return "ColumnSimpleTypeNumber"
	case "string":
		return "ColumnSimpleTypeString"
	case "boolean":
		return "ColumnSimpleTypeBoolean"
	case "array":
		return "ColumnSimpleTypeArray"
	case "object":
		return "ColumnSimpleTypeObject"
	default:
		return ""
	}
}

func goTypeToSimpleType(goType string) (string, error) {
	t := strings.TrimPrefix(goType, "*")
	switch {
	case t == "time.Time":
		return "date-time", nil
	case strings.HasPrefix(t, "int"):
		return "integer", nil
	case strings.HasPrefix(t, "float"):
		return "number", nil
	case t == "string" || t == "uuid.UUID":
		return "string", nil
	case t == "bool":
		return "boolean", nil
	case strings.HasPrefix(t, "[]"):
		return "array", nil
	case strings.HasPrefix(t, "map["):
		return "object", nil
	default:
		return "", fmt.Errorf("unsupported simple type: %v", t)
	}
}

func (f *Funcs) sort_fields(fields []Field) []Field {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].SQLName < fields[j].SQLName
	})

	return fields
}

// set_field generates an assignment to a struct field.
func (f *Funcs) set_field(field Field, typ string, table Table) (string, error) {
	skipField := field.IsGenerated || field.IsIgnored || field.SQLName == "deleted_at" //|| contains(table.ForeignKeys, field.SQLName)
	if skipField {
		return "", nil
	}

	af := analyzeField(table, field)
	switch typ {
	case "CreateParams":
	case "UpsertParams":
	case "UpdateParams":
		if af.PKisFK != nil { // e.g. workitemid in project tables. don't ever want to update it.
			fmt.Printf("UpdateParams: skipping %q: is a single foreign and primary key in table %q\n", field.SQLName, table.SQLName)
			return "", nil
		}
	}

	goName := field.GoName
	if field.EnumSchema != "" {
		goName = camelExport(f.schemaPrefix) + goName
	}

	switch typ {
	case "UpsertParams":
		return fmt.Sprintf("\t%[1]s.%[2]s = params.%[2]s\n", f.short(table), goName), nil
	case "CreateParams":
		return fmt.Sprintf("\t%[1]s: params.%[1]s,\n", goName), nil
	case "UpdateParams":
		return fmt.Sprintf(`if params.%[2]s != nil {
	%[1]s.%[2]s = *params.%[2]s
}
`, f.short(table), goName), nil
	}

	return "", fmt.Errorf("invalid typ: %s", typ)
}

type fieldInfo struct {
	IsSinglePK bool
	PKisFK     *TableForeignKey
	IsFK       bool
	IsSingleFK bool
	IsPK       bool
}

func analyzeField(table Table, field Field) fieldInfo {
	var isSinglePK, isFK, isSingleFK, isPK bool
	var pkisfk *TableForeignKey

	for _, tfk := range table.ForeignKeys {
		for _, f := range tfk.FieldNames {
			if field.SQLName != f {
				continue
			}
			isFK = true
			if len(tfk.FieldNames) == 1 {
				isSingleFK = true
			}
		}
	}

	for _, tpk := range table.PrimaryKeys {
		if tpk.SQLName != field.SQLName {
			continue
		}
		isPK = true
		if len(table.PrimaryKeys) == 1 {
			isSinglePK = true
		}
	}

	if isSinglePK && isSingleFK { // excluding m2m join tables with 2 primary keys that are fks
		for _, tfk := range table.ForeignKeys {
			if tfk.FieldNames[0] == table.PrimaryKeys[0].SQLName {
				tfk := tfk
				pkisfk = &tfk
				break
			}
		}
	}

	r := fieldInfo{
		IsSinglePK: isSinglePK,
		PKisFK:     pkisfk,
		IsFK:       isFK,
		IsSingleFK: isSingleFK,
		IsPK:       isPK,
	}

	return r
}

// fieldmapping generates field mappings from a struct to another.
func (f *Funcs) fieldmapping(field Field, recv string, public bool) (string, error) {
	if public {
		if contains(field.Properties, propertyJSONPrivate) {
			return "", nil
		}
	}
	goName := field.GoName
	if field.EnumSchema != "" {
		goName = camelExport(f.schemaPrefix) + goName
	}

	return fmt.Sprintf("\t%s: %s.%s,", goName, recv, goName), nil
}

// join_fields generates a struct field definition from join constraints
func (f *Funcs) join_fields(t Table, constraints []Constraint, tables Tables) (string, error) {
	var buf strings.Builder
	var goName, tag, typ string

	/**
	 * TODO generate FK joins here regardless of constraints, instead of generating FK*** functions
	 */
	// for _, tfk := range t.ForeignKeys {
	// 	fmt.Printf("tfk: %v\n", tfk)
	// 	if len(tfk.FieldNames) == 1 { // generate O2O joins automatically. Reuse O2O logic to generate queries, structs, etc.
	// 		lookupName := tfk.RefTable
	// 		goName = camelExport(singularize(strings.TrimSuffix(tfk.FieldNames[0], "_id")))
	// 		typ = camelExport(singularize(lookupName))
	// 		tag = fmt.Sprintf("`json:\"-\" db:\"%s\"`", lookupName)
	// 		buf.WriteString(fmt.Sprintf("\t%sJoin *%s %s // %s field FK \n", goName, typ, tag, goName))
	// 	}
	// }

	if len(constraints) > 0 {
		f.loadConstraints(constraints, t.SQLName, nil)
	}
	cc, ok := f.tableConstraints[t.SQLName]
	if !ok {
		return "", nil
	}
	goNames := []string{}
	for _, c := range cc {
		if c.Type != "foreign_key" {
			continue
		}
		var notes, joinName string
		// sync with extratypes
		switch c.Cardinality {
		case M2M:
			notes += " " + c.TableName
			lookupName := strings.TrimSuffix(c.ColumnName, "_id")
			joinName := c.TableName + "_" + inflector.Pluralize(lookupName)
			typ = camelExport(singularize(c.RefTableName))

			lookupTable := tables[c.TableName]
			m2mExtraCols := getTableRegularFields(lookupTable)
			if len(m2mExtraCols) > 0 {
				typ = t.GoName + "M2M" + camelExport(lookupName) + toAcronym(c.TableName)
			} else {
				typ = camelExport(f.schemaPrefix) + typ
			}

			goName = camelExport(inflector.Pluralize(lookupName))
			rc := camelExport(strings.TrimSuffix(c.RefColumnName, "_id"))
			rlc := camelExport(strings.TrimSuffix(c.LookupRefColumnName, "_id"))
			lc := camelExport(strings.TrimSuffix(c.LookupColumnName, "_id"))
			col := camelExport(strings.TrimSuffix(c.ColumnName, "_id"))
			_, _, _, _ = rc, rlc, lc, col
			if rlc != lc && rlc != col {
				// e.g. m2m join in teams table with ref lookup table column of member (instead of user id) yields just MembersJoin.
				// if we had something different than team_id in lookup table it would prefix it. but MembersJoin assumes it team member.

				goName = lc + goName
			}
			// if rc != col && rlc != col {
			// 	goName = col + goName
			// }
			for _, g := range goNames {
				if g == goName+"Join" {
					goName = goName + toAcronym(c.TableName)
				}
			}
			goName += "Join"
			tag = fmt.Sprintf("`json:\"-\" db:\"%s\"`", joinName)
			buf.WriteString(fmt.Sprintf("\t%s *[]%s %s // %s\n", goName, typ, tag, string(c.Cardinality)+notes))
		case M2O:
			if c.RefTableName != t.SQLName {
				continue
			}
			notes += " " + c.RefTableName
			joinName = inflector.Pluralize(c.TableName)
			if c.JoinTableClash {
				joinName = joinName + "_" + c.ColumnName
			}

			typ = camelExport(singularize(c.TableName))
			goName = camelExport(c.TableName)
			if c.JoinTableClash {
				goName = camelExport(c.ColumnName) + goName
			}
			for _, g := range goNames {
				if g == goName+"Join" {
					goName = goName + toAcronym(c.TableName)
				}
			}

			goName += "Join"

			typ = camelExport(f.schemaPrefix) + typ

			tag = fmt.Sprintf("`json:\"-\" db:\"%s\"`", joinName)
			buf.WriteString(fmt.Sprintf("\t%s *[]%s %s // %s\n", goName, typ, tag, string(c.Cardinality)+notes))
		case O2O:
			if c.TableName != t.SQLName {
				continue
			}
			typ = camelExport(singularize(c.RefTableName))

			notes += " " + c.RefTableName
			if c.IsInferredO2O {
				notes += " (inferred)"
			}
			if c.IsGeneratedO2OFromM2O {
				notes += " (generated from M2O)"
			}

			joinPrefix := singularize(c.RefTableName) + "_"
			joinName := joinPrefix + singularize(c.ColumnName)

			goName = camelExport(singularize(c.RefTableName))
			if c.JoinTableClash {
				goName = goName + camelExport(c.ColumnName)
			}
			for _, g := range goNames {
				if g == goName+"Join" {
					goName = goName + toAcronym(c.TableName)
				}
			}

			goName += "Join"
			typ = camelExport(f.schemaPrefix) + typ

			tag = fmt.Sprintf("`json:\"-\" db:\"%s\"`", joinName)
			buf.WriteString(fmt.Sprintf("\t%s *%s %s // %s\n", goName, typ, tag, string(c.Cardinality)+notes))
			// dummy created automatically to avoid this duplication
			// if c.RefTableName == t.SQLName {
			// 	goName = camelExport(singularize(c.TableName))
			// 	typ = goName
			// 	goName = goName + "Join"
			// 	tag = fmt.Sprintf("`json:\"-\" db:\"%s\"`", singularize(c.TableName))
			// 	buf.WriteString(fmt.Sprintf("\t%s *%s %s // %s\n", goName, typ, tag, c.Cardinality))
			// }
		default:
			continue
		}
		goNames = append(goNames, goName)
	}

	return buf.String(), nil
}

// short generates a safe Go identifier for typ. typ is first checked
// against shorts, and if not found, then the value is calculated and
// stored in the shorts for future use.
//
// A short is the concatenation of the lowercase of the first character in
// the words comprising the name. For example, "MyCustomName" will have have
// the short of "mcn".
//
// If a generated short conflicts with a Go reserved name or a name used in
// the templates, then the corresponding value in goReservedNames map will be
// used.
//
// Generated shorts that have conflicts with any scopeConflicts member will
// have nameConflictSuffix appended.
func (f *Funcs) short(v any) string {
	var n string
	switch x := v.(type) {
	case string:
		n = x
	case Table:
		n = x.GoName
	default:
		return fmt.Sprintf("[[ UNSUPPORTED TYPE 30: %T ]]", v)
	}
	// check short name map
	name, ok := f.shorts[n]
	if !ok {
		// calc the short name
		var u []string
		for _, s := range strings.Split(strings.ToLower(snaker.CamelToSnake(n)), "_") {
			if len(s) > 0 && s != "id" {
				u = append(u, s[:1])
			}
		}
		// ensure no name conflict
		name = checkName(strings.Join(u, ""))
		// store back to short name map
		f.shorts[n] = name
	}
	// append suffix if conflict exists
	if _, ok := templateReservedNames[name]; ok {
		name += f.conflict
	}
	return name
}

// colname returns the ColumnName of a field escaped if needed.
func (f *Funcs) colname(z Field) string {
	if f.escColumn {
		return escfn(z.SQLName)
	}
	return z.SQLName
}

func checkName(name string) string {
	if n, ok := goReservedNames[name]; ok {
		return n
	}
	return name
}

// escfn escapes s.
func escfn(s string) string {
	return `"` + s + `"`
}

// eval evalutates a template s against v.
func eval(v any, s string) (string, error) {
	tpl, err := template.New(fmt.Sprintf("[EVAL %q]", s)).Parse(s)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// templateReservedNames are the template reserved names.
var templateReservedNames = map[string]bool{
	// variables
	"ctx":  true,
	"db":   true,
	"err":  true,
	"log":  true,
	"logf": true,
	"res":  true,
	"rows": true,

	// packages
	"context": true,
	"csv":     true,
	"driver":  true,
	"errors":  true,
	"fmt":     true,
	"hstore":  true,
	"regexp":  true,
	"sql":     true,
	"strings": true,
	"time":    true,
	"uuid":    true,
}

// goReservedNames is a map of of go reserved names to "safe" names.
var goReservedNames = map[string]string{
	"break":       "brk",
	"case":        "cs",
	"chan":        "chn",
	"const":       "cnst",
	"continue":    "cnt",
	"default":     "def",
	"defer":       "dfr",
	"else":        "els",
	"fallthrough": "flthrough",
	"for":         "fr",
	"func":        "fn",
	"go":          "goVal",
	"goto":        "gt",
	"if":          "ifVal",
	"import":      "imp",
	"interface":   "iface",
	"map":         "mp",
	"package":     "pkg",
	"range":       "rnge",
	"return":      "ret",
	"select":      "slct",
	"struct":      "strct",
	"switch":      "swtch",
	"type":        "typ",
	"var":         "vr",
	// go types
	"error":      "e",
	"bool":       "b",
	"string":     "str",
	"byte":       "byt",
	"rune":       "r",
	"uintptr":    "uptr",
	"int":        "i",
	"int8":       "i8",
	"int16":      "i16",
	"int32":      "i32",
	"int64":      "i64",
	"uint":       "u",
	"uint8":      "u8",
	"uint16":     "u16",
	"uint32":     "u32",
	"uint64":     "u64",
	"float32":    "z",
	"float64":    "f",
	"complex64":  "c",
	"complex128": "c128",
}

// Context keys.
var (
	AppendKey          xo.ContextKey = "append"
	KnownTypesKey      xo.ContextKey = "known-types"
	ShortsKey          xo.ContextKey = "shorts"
	NotFirstKey        xo.ContextKey = "not-first"
	Int32Key           xo.ContextKey = "int32"
	Uint32Key          xo.ContextKey = "uint32"
	ArrayModeKey       xo.ContextKey = "array-mode"
	PkgKey             xo.ContextKey = "pkg"
	TagKey             xo.ContextKey = "tag"
	ImportKey          xo.ContextKey = "import"
	UUIDKey            xo.ContextKey = "uuid"
	CustomKey          xo.ContextKey = "custom"
	ConflictKey        xo.ContextKey = "conflict"
	InitialismKey      xo.ContextKey = "initialism"
	EscKey             xo.ContextKey = "esc"
	FieldTagKey        xo.ContextKey = "field-tag"
	PublicFieldTagKey  xo.ContextKey = "public-field-tag"
	PrivateFieldTagKey xo.ContextKey = "private-field-tag"
	ContextKey         xo.ContextKey = "context"
	InjectKey          xo.ContextKey = "inject"
	InjectFileKey      xo.ContextKey = "inject-file"
	LegacyKey          xo.ContextKey = "legacy"
)

// Append returns append from the context.
func Append(ctx context.Context) bool {
	b, _ := ctx.Value(AppendKey).(bool)
	return b
}

// KnownTypes returns known-types from the context.
func KnownTypes(ctx context.Context) map[string]bool {
	m, _ := ctx.Value(KnownTypesKey).(map[string]bool)
	return m
}

// Shorts retruns shorts from the context.
func Shorts(ctx context.Context) map[string]string {
	m, _ := ctx.Value(ShortsKey).(map[string]string)
	return m
}

// NotFirst returns not-first from the context.
func NotFirst(ctx context.Context) bool {
	b, _ := ctx.Value(NotFirstKey).(bool)
	return b
}

// Int32 returns int32 from the context.
func Int32(ctx context.Context) string {
	s, _ := ctx.Value(Int32Key).(string)
	return s
}

// Uint32 returns uint32 from the context.
func Uint32(ctx context.Context) string {
	s, _ := ctx.Value(Uint32Key).(string)
	return s
}

// ArrayMode returns array-mode from the context.
func ArrayMode(ctx context.Context) string {
	s, _ := ctx.Value(ArrayMode).(string)
	return s
}

// Pkg returns pkg from the context.
func Pkg(ctx context.Context) string {
	s, _ := ctx.Value(PkgKey).(string)
	if s == "" {
		s = filepath.Base(xo.Out(ctx))
	}
	return s
}

// Tags returns tags from the context.
func Tags(ctx context.Context) []string {
	v, _ := ctx.Value(TagKey).([]string)
	// build tags
	var tags []string
	for _, tag := range v {
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

// Imports returns package imports from the context.
func Imports(ctx context.Context) []string {
	v, _ := ctx.Value(ImportKey).([]string)
	// build imports
	var imports []string
	for _, s := range v {
		if s != "" {
			imports = append(imports, s)
		}
	}
	// add uuid import
	if s, _ := ctx.Value(UUIDKey).(string); s != "" {
		imports = append(imports, s)
	}

	return imports
}

// Custom returns the custom package from the context.
func Custom(ctx context.Context) string {
	s, _ := ctx.Value(CustomKey).(string)
	return s
}

// Conflict returns conflict from the context.
func Conflict(ctx context.Context) string {
	s, _ := ctx.Value(ConflictKey).(string)
	return s
}

// Esc indicates if esc should be escaped based from the context.
func Esc(ctx context.Context, esc string) bool {
	v, _ := ctx.Value(EscKey).([]string)
	return !contains(v, "none") && (contains(v, "all") || contains(v, esc))
}

// FieldTag returns field-tag from the context.
func FieldTag(ctx context.Context) string {
	s, _ := ctx.Value(FieldTagKey).(string)
	return s
}

// PublicFieldTag returns field-tag from the context.
func PublicFieldTag(ctx context.Context) string {
	s, _ := ctx.Value(PublicFieldTagKey).(string)
	return s
}

// PrivateFieldTag returns field-tag from the context.
func PrivateFieldTag(ctx context.Context) string {
	s, _ := ctx.Value(PrivateFieldTagKey).(string)
	return s
}

// Context returns context from the context.
func Context(ctx context.Context) string {
	s, _ := ctx.Value(ContextKey).(string)
	return s
}

// Inject returns inject from the context.
func Inject(ctx context.Context) string {
	s, _ := ctx.Value(InjectKey).(string)
	return s
}

// InjectFile returns inject-file from the context.
func InjectFile(ctx context.Context) string {
	s, _ := ctx.Value(InjectFileKey).(string)
	return s
}

// Legacy returns legacy from the context.
func Legacy(ctx context.Context) bool {
	b, _ := ctx.Value(LegacyKey).(bool)
	return b
}

// add returns the sum of a and b.
func add(b, a int) int {
	return a + b
}

func table_is_updatable(fields []Field) bool {
	for _, field := range fields {
		if !field.IsPrimary && !field.IsGenerated {
			return true // at least one field is updatable
		}
	}
	return false
}

// addInitialisms adds snaker initialisms from the context.
func addInitialisms(ctx context.Context) error {
	z := ctx.Value(InitialismKey)
	y, _ := z.([]string)
	var v []string
	for _, s := range y {
		if s != "" {
			v = append(v, s)
		}
	}
	return snaker.DefaultInitialisms.Add(v...)
}

// contains determines if v contains s.
func contains(v []string, s string) bool {
	for _, z := range v {
		if z == s {
			return true
		}
	}
	return false
}

// singularize singularizes s.
func singularize(s string) string {
	if i := strings.LastIndex(s, "_"); i != -1 {
		return s[:i+1] + inflector.Singularize(s[i+1:])
	}
	return inflector.Singularize(s)
}

// EnumValue is a enum value template.
type EnumValue struct {
	GoName     string
	SQLName    string
	ConstValue string
}

// Enum is a enum type template.
type Enum struct {
	GoName       string
	SQLName      string
	Values       []EnumValue
	Comment      string
	Pkg          string
	GoNamePrefix string
}

// Proc is a stored procedure template.
type Proc struct {
	Type           string
	GoName         string
	OverloadedName string
	SQLName        string
	Signature      string
	Params         []Field
	Returns        []Field
	Void           bool
	Overloaded     bool
	Comment        string
}

// Table is a type (ie, table/view/custom query) template.
// IMPORTANT: runtime out of memory... will need to optimize fields here
// (investigate why changing to []*Field didn't do anything)
type Table struct {
	Type        string
	GoName      string
	SQLName     string
	PrimaryKeys []Field
	Fields      []Field
	Manual      bool
	Comment     string
	Generated   []Field
	Ignored     []Field
	ForeignKeys []TableForeignKey
	Schema      string
}

type TableForeignKey struct {
	FieldNames []string
	RefTable   string
	RefColumns []string
}

// ForeignKey is a foreign key template.
type ForeignKey struct {
	GoName    string
	SQLName   string
	Table     Table
	Fields    []Field
	RefTable  string
	RefFields []Field
	RefFunc   string
	Comment   string
}

// Index is an index template.
type Index struct {
	SQLName    string
	Func       string
	Table      Table
	Fields     []Field
	IsUnique   bool
	IsPrimary  bool
	Comment    string
	Definition string
}

// Constraint is a table constraint.
type Constraint struct {
	// "unique" "check" "primary_key" "foreign_key"
	Type        string
	Cardinality cardinality
	// Postgres constraint name
	Name                   string
	TableName              string // table where FK is defined
	ColumnName             string
	ColumnComment          string
	RefTableName           string // table FK references
	RefColumnName          string // RefTableName column FK references
	RefColumnComment       string
	LookupColumnName       string // (M2M) lookup table column
	LookupColumnComment    string
	LookupRefColumnName    string // (M2M) referenced PK by LookupColumn
	LookupRefColumnComment string
	JoinTableClash         bool // Whether other constraints join against the same table
	IsInferredO2O          bool // Whether this constraint has been generated from a foreign key
	IsGeneratedO2OFromM2O  bool
	JoinStructFieldClash   bool // Whether 2 or more constraints of the same table have the same struct field name (and hence type as well)
	PKisFK                 bool
}

// Field is a field template.
type Field struct {
	GoName         string
	SQLName        string
	Type           string
	Zero           string
	IsPrimary      bool
	IsSequence     bool
	IsIgnored      bool
	Comment        string
	IsGenerated    bool
	EnumPkg        string
	EnumSchema     string
	TypeOverride   string
	IsDateOrTime   bool
	Properties     []string
	OpenAPISchema  string
	ExtraTags      string
	UnderlyingType string
	Annotations    map[annotation]string
}

// QueryParam is a custom query parameter template.
type QueryParam struct {
	Name        string
	Type        string
	Interpolate bool
	Join        bool
}

// Query is a custom query template.
type Query struct {
	Name        string
	Query       []string
	Comments    []string
	Params      []QueryParam
	One         bool
	Flat        bool
	Exec        bool
	Interpolate bool
	Type        Table
	Comment     string
}

// PackageImport holds information about a Go package import.
type PackageImport struct {
	Alias string
	Pkg   string
}

// String satisfies the fmt.Stringer interface.
func (v PackageImport) String() string {
	if v.Alias != "" {
		return fmt.Sprintf("%s %q", v.Alias, v.Pkg)
	}
	return fmt.Sprintf("%q", v.Pkg)
}

//--------------------------------------------------------------------------------------------
// legacy funcs

// addLegacyFuncs adds the legacy template funcs.
func addLegacyFuncs(ctx context.Context, funcs template.FuncMap) {
	nth, err := loader.NthParam(ctx)
	if err != nil {
		return
	}
	// colnames creates a list of the column names found in fields, excluding any
	// Field with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of column names, that can be used in
	// a SELECT, or UPDATE, or other SQL clause requiring an list of identifiers
	// (ie, "field_1, field_2, field_3, ...").
	funcs["colnames"] = func(fields []*Field, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + f.SQLName
			i++
		}
		return str
	}
	// colnamesmulti creates a list of the column names found in fields, excluding any
	// Field with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of column names, that can be used in
	// a SELECT, or UPDATE, or other SQL clause requiring an list of identifiers
	// (ie, "field_1, field_2, field_3, ...").
	funcs["colnamesmulti"] = func(fields []*Field, ignoreNames []*Field) string {
		ignore := map[string]bool{}
		for _, f := range ignoreNames {
			ignore[f.SQLName] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + f.SQLName
			i++
		}
		return str
	}
	// colnamesquery creates a list of the column names in fields as a query and
	// joined by sep, excluding any Field with Name contained in ignoreNames.
	//
	// Used to create a list of column names in a WHERE clause (ie, "field_1 = $1
	// AND field_2 = $2 AND ...") or in an UPDATE clause (ie, "field = $1, field =
	// $2, ...").
	funcs["colnamesquery"] = func(fields []*Field, sep string, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + sep
			}
			str = str + f.SQLName + " = " + nth(i)
			i++
		}
		return str
	}
	// colnamesquerymulti creates a list of the column names in fields as a query and
	// joined by sep, excluding any Field with Name contained in the slice of fields in ignoreNames.
	//
	// Used to create a list of column names in a WHERE clause (ie, "field_1 = $1
	// AND field_2 = $2 AND ...") or in an UPDATE clause (ie, "field = $1, field =
	// $2, ...").
	funcs["colnamesquerymulti"] = func(fields []*Field, sep string, startCount int, ignoreNames []*Field) string {
		ignore := map[string]bool{}
		for _, f := range ignoreNames {
			ignore[f.SQLName] = true
		}
		str := ""
		i := startCount
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i > startCount {
				str = str + sep
			}
			str = str + f.SQLName + " = " + nth(i)
			i++
		}
		return str
	}
	// colprefixnames creates a list of the column names found in fields with the
	// supplied prefix, excluding any Field with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of column names with a prefix. Used in
	// a SELECT, or UPDATE (ie, "t.field_1, t.field_2, t.field_3, ...").
	funcs["colprefixnames"] = func(fields []*Field, prefix string, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + prefix + "." + f.SQLName
			i++
		}
		return str
	}
	// colvals creates a list of value place holders for fields excluding any Field
	// with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of column place holders, used in a
	// SELECT or UPDATE statement (ie, "$1, $2, $3 ...").
	funcs["colvals"] = func(fields []*Field, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + nth(i)
			i++
		}
		return str
	}
	// colvalsmulti creates a list of value place holders for fields excluding any Field
	// with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of column place holders, used in a
	// SELECT or UPDATE statement (ie, "$1, $2, $3 ...").
	funcs["colvalsmulti"] = func(fields []*Field, ignoreNames []*Field) string {
		ignore := map[string]bool{}
		for _, f := range ignoreNames {
			ignore[f.SQLName] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + nth(i)
			i++
		}
		return str
	}
	// fieldnames creates a list of field names from fields of the adding the
	// provided prefix, and excluding any Field with Name contained in ignoreNames.
	//
	// Used to present a comma separated list of field names, ie in a Go statement
	// (ie, "t.Field1, t.Field2, t.Field3 ...")
	funcs["fieldnames"] = func(fields []*Field, prefix string, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + prefix + "." + f.SQLName
			i++
		}
		return str
	}
	// fieldnamesmulti creates a list of field names from fields of the adding the
	// provided prefix, and excluding any Field with the slice contained in ignoreNames.
	//
	// Used to present a comma separated list of field names, ie in a Go statement
	// (ie, "t.Field1, t.Field2, t.Field3 ...")
	funcs["fieldnamesmulti"] = func(fields []*Field, prefix string, ignoreNames []*Field) string {
		ignore := map[string]bool{}
		for _, f := range ignoreNames {
			ignore[f.SQLName] = true
		}
		str := ""
		i := 0
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			if i != 0 {
				str = str + ", "
			}
			str = str + prefix + "." + f.SQLName
			i++
		}
		return str
	}
	// colcount returns the 1-based count of fields, excluding any Field with Name
	// contained in ignoreNames.
	//
	// Used to get the count of fields, and useful for specifying the last SQL
	// parameter.
	funcs["colcount"] = func(fields []*Field, ignoreNames ...string) int {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		i := 1
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			i++
		}
		return i
	}
	// goparamlist converts a list of fields into their named Go parameters,
	// skipping any Field with Name contained in ignoreNames. addType will cause
	// the go Type to be added after each variable name. addPrefix will cause the
	// returned string to be prefixed with ", " if the generated string is not
	// empty.
	//
	// Any field name encountered will be checked against goReservedNames, and will
	// have its name substituted by its corresponding looked up value.
	//
	// Used to present a comma separated list of Go variable names for use with as
	// either a Go func parameter list, or in a call to another Go func.
	// (ie, ", a, b, c, ..." or ", a T1, b T2, c T3, ...").
	funcs["goparamlist"] = func(fields []*Field, addPrefix bool, addType bool, ignoreNames ...string) string {
		ignore := map[string]bool{}
		for _, n := range ignoreNames {
			ignore[n] = true
		}
		i := 0
		var vals []string
		for _, f := range fields {
			if ignore[f.SQLName] {
				continue
			}
			s := "v" + strconv.Itoa(i)
			if len(f.SQLName) > 0 {
				n := strings.Split(snaker.CamelToSnake(f.SQLName), "_")
				s = strings.ToLower(n[0]) + f.SQLName[len(n[0]):]
			}
			// add the go type
			if addType {
				s += " " + f.Type
			}
			// add to vals
			vals = append(vals, s)
			i++
		}
		// concat generated values
		str := strings.Join(vals, ", ")
		if addPrefix && str != "" {
			return ", " + str
		}
		return str
	}
	// convext generates the Go conversion for f in order for it to be assignable
	// to t.
	//
	//  this should be a better name, like "goconversion" or some such.
	// funcs["convext"] = func(prefix string, f *Field, t *Field) string {
	// 	expr := prefix + "." + f.SQLName
	// 	if f.Type == t.Type {
	// 		return expr
	// 	}
	// 	ft := f.Type
	// 	if strings.HasPrefix(ft, "*") {
	// 		typ := f.Type[:1]
	// 		// pending nil checks generate and return err
	// 		expr = "*" + expr
	// 		ft = strings.ToLower(typ)
	// 	}
	// 	if t.Type != ft {
	// 		expr = t.Type + "(" + expr + ")"
	// 	}
	// 	return expr
	// }
	// getstartcount returns a starting count for numbering columns in queries
	funcs["getstartcount"] = func(fields []*Field, pkFields []*Field) int {
		return len(fields) - len(pkFields)
	}
}
