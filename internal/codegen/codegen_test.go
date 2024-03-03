package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeSpec(t *testing.T) {
	t.Parallel()

	const baseDir = "testdata/analyze_specs"

	cases := []struct {
		name        string
		file        string
		errContains string
	}{
		{
			"valid",
			"valid.yaml",
			``,
		},
		{
			"invalid_operationid",
			"invalid_operationid.yaml",
			fmt.Sprintf(`path "/pet/ConflictEndpointPet": method "GET": operationId "Conflict-Endpoint-Pet" does not match pattern %q`, OperationIDRE.String()),
		},
		{
			"missing_operationid",
			"missing_operationid.yaml",
			`path "/pet/ConflictEndpointPet": method "GET": operationId is required for codegen`,
		},
		{
			"more_than_one_tag",
			"more_than_one_tag.yaml",
			`path "/pet/ConflictEndpointPet": method "GET": at most one tag is permitted for codegen`,
		},
		{
			"invalid_tag",
			"invalid_tag.yaml",
			fmt.Sprintf(`path "/pet/ConflictEndpointPet": method "GET": tag must be a valid filename with pattern %q`, validFilenameRE.String()),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var stderr bytes.Buffer

			og := New(&stderr, path.Join(baseDir, tc.file), "", "")

			err := og.analyzeSpec()
			if tc.errContains != "" {
				assert.ErrorContains(t, err, tc.errContains)
			} else if err != nil {
				t.Fatalf("err: %s\nstderr: %s\n", err, &stderr)
			}
		})
	}
}

func TestGetHandlersMethods(t *testing.T) {
	t.Parallel()

	source := `
package rest

type StrictHandlers struct{}

func (h *StrictHandlers) MyFunction1(c *gin.Context) {}

func (h *StrictHandlers) MyFunction2(c *gin.Context) {}

func (h *StrictHandlers) UnrelatedFunction(c *gin.Context) {}

func MyFunction3(c *gin.Context) {}

func (h *SomeOtherStruct) NotHandlersMethod() {}
`

	file, err := parseAST(strings.NewReader(source))
	require.NoError(t, err)

	functions := getHandlersMethods(file)
	assert.ElementsMatch(t, []any{"MyFunction1", "MyFunction2", "UnrelatedFunction"}, functions)
}

func TestEnsureFunctionMethods_MisplacedMethod(t *testing.T) {
	t.Parallel()

	type tagHandlerFile struct {
		// leave empty to skip file creation
		content    string
		newContent string
	}

	type tag = string
	const foo = tag("foo")
	const bar = tag("bar")
	const barCamelCase = tag("barCamelCase")

	type handlerFiles map[tag]tagHandlerFile

	type testCase struct {
		name        string
		operations  map[tag][]string
		files       handlerFiles
		errContains []string
	}

	tests := []testCase{
		{
			name:       "swap handlers to correct file",
			operations: map[tag][]string{foo: {"Foo", "Bar"}, bar: {"Baz"}},
			files: handlerFiles{
				foo: {
					content: `package rest

func (h *StrictHandlers) Foo()             {}
func (h *StrictHandlers) UnrelatedMethod() {}
`,
					newContent: `package rest

func (h *StrictHandlers) Foo()             {}
func (h *StrictHandlers) UnrelatedMethod() {}
func (h *StrictHandlers) Bar()             {}
`,
				},
				bar: {
					content: `package rest

func (h *StrictHandlers) Baz() {}
func (h *StrictHandlers) Bar() {}
`,
					newContent: `package rest

func (h *StrictHandlers) Baz() {}
`,
				},
			},
		},
		{
			name:       "no changes",
			operations: map[tag][]string{foo: {"Foo", "Bar"}, bar: {"Baz"}},
			files: func() handlerFiles {
				fooContent := `package rest

				func (h *StrictHandlers) Foo() {}
				func (h *StrictHandlers) Bar() {}
`
				barContent := `package rest

				func (h *StrictHandlers) Baz() {}
`

				return handlerFiles{
					foo: {content: fooContent, newContent: fooContent},
					bar: {content: barContent, newContent: barContent},
				}
			}(),
		},
		{
			name:       "misplaced with no correct file created yet gets transferred retaining method body",
			operations: map[tag][]string{foo: {"Foo"}, bar: {"Bar"}},
			files: func() handlerFiles {
				return handlerFiles{
					foo: {
						content: `package rest

func (h *StrictHandlers) Foo() {}
func (h *StrictHandlers) Bar() {
	a := 1
	_ = a
}
`,
						newContent: `package rest

func (h *StrictHandlers) Foo() {}
`,
					},
					bar: {
						content: "", // won't create
						newContent: `package rest

func (h *StrictHandlers) Bar() {
	a := 1
	_ = a
}
`,
					},
				}
			}(),
		},
		// we now fix these automatically via server interface implementation
		// {
		// 	name:        "missing method in existing file",
		// 	errContains: []string{`missing function method for operation ID "Bar" in api_foo.go`},
		// 	operations:  map[tag][]string{foo: {"Foo", "Bar"}},
		// 	files: func() handlerFiles {
		// 		content := `package rest

		// 		func (h *StrictHandlers) Foo() {}
		// 		`

		// 		return handlerFiles{
		// 			foo: {content: content, newContent: content},
		// 		}
		// 	}(),
		// },
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tmpDir := t.TempDir()

			for tag, file := range tc.files {
				if file.content == "" {
					continue
				}
				apiFilePath := filepath.Join(tmpDir, fmt.Sprintf("api_%s.go", tag))

				f, err := os.Create(apiFilePath)
				require.NoError(t, err, "Failed to create test file")
				defer f.Close()

				_, err = f.WriteString(file.content)
				require.NoError(t, err, "Failed to write test file")
			}

			o := New(&bytes.Buffer{}, "", "", tmpDir)
			o.operations = tc.operations

			err := o.ensureHandlerMethodsExist()
			if len(tc.errContains) > 0 {
				for _, e := range tc.errContains {
					require.ErrorContains(t, err, e)
				}
			} else {
				require.NoError(t, err)
			}

			for tag, file := range tc.files {
				apiFilePath := filepath.Join(tmpDir, fmt.Sprintf("api_%s.go", tag))

				c, err := os.ReadFile(apiFilePath)
				require.NoError(t, err, "Failed to read test file %s", apiFilePath)

				if diff := cmp.Diff(file.newContent, string(c)); diff != "" {
					t.Errorf("file.newContent mismatch (tag=%s) (-want +got):\n%s", tag, diff)
				}
			}
		})
	}
}

func TestRemoveAndAppendHandlersMethod(t *testing.T) {
	t.Parallel()

	sourceCode := `
package main

type StrictHandlers struct{}

func (h *StrictHandlers) GetSomething() {
	return "GetSomething"
}

func (h *StrictHandlers) UpdateSomething() {
	return "UpdateSomething"
}

func main() {}
`
	file, err := parseAST(strings.NewReader(sourceCode))
	require.NoError(t, err)

	methodNameToRemove := "GetSomething"

	secondFile, err := parseAST(strings.NewReader(`
package main

func main() {}
`))
	require.NoError(t, err)

	removeAndAppendHandlersMethod(file, secondFile, methodNameToRemove)

	for _, decl := range file.Decls {
		if fd, ok := decl.(*dst.FuncDecl); ok {
			if fd.Name.Name == methodNameToRemove {
				t.Errorf("StrictHandlers method '%s' still found in the first file after removal.", methodNameToRemove)
				break
			}
		}
	}

	var found bool
	for _, decl := range secondFile.Decls {
		if fd, ok := decl.(*dst.FuncDecl); ok {
			if fd.Name.Name == methodNameToRemove {
				found = true
				break
			}
		}
	}
	var buf bytes.Buffer
	if err := decorator.Fprint(&buf, secondFile); err != nil {
		fmt.Println("Error printing AST:", err)
		return
	}

	fmt.Printf("buf.String(): %v\n", buf.String())

	if !found {
		t.Errorf("StrictHandlers method '%s' not found in the second file after appending.", methodNameToRemove)
	}
}
