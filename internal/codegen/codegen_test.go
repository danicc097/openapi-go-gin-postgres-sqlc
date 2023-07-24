package codegen

import (
	"bytes"
	"fmt"
	"go/printer"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dave/dst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeSpec(t *testing.T) {
	t.Parallel()

	const baseDir = "testdata/analyze_specs"

	cases := []struct {
		Name        string
		File        string
		ErrContains string
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
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			var stderr bytes.Buffer

			og := New(&stderr, path.Join(baseDir, tc.File), "", "")

			err := og.analyzeSpec()
			if tc.ErrContains != "" {
				assert.ErrorContains(t, err, tc.ErrContains)
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

type Handlers struct{}

func (h *Handlers) MyFunction1(c *gin.Context) {}

func (h *Handlers) MyFunction2(c *gin.Context) {}

func (h *Handlers) UnrelatedFunction(c *gin.Context) {}

func MyFunction3(c *gin.Context) {}

func (h *SomeOtherStruct) NotHandlersMethod() {}
`

	file, err := parseAST(strings.NewReader(source))
	assert.NoError(t, err)

	functions := getHandlersMethods(file)
	assert.ElementsMatch(t, []string{"MyFunction1", "MyFunction2", "UnrelatedFunction"}, functions)
}

func TestEnsureFunctionMethods_MisplacedMethod(t *testing.T) {
	t.Parallel()

	fileContentFoo := `
package rest

type Handlers struct{}

func (h *Handlers) Foo() {}
func (h *Handlers) Bar() {}
`

	fileContentBar := `
package rest

type Handlers struct{}

func (h *Handlers) Baz() {}
func (h *Handlers) Qux() {}
`

	t.Run("swap handlers", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()
		apiFilePathFoo := filepath.Join(tmpDir, "api_foo.go")
		apiFilePathBar := filepath.Join(tmpDir, "api_bar.go")

		file, err := os.Create(apiFilePathFoo)
		require.NoError(t, err, "Failed to create test file")
		defer file.Close()

		_, err = file.WriteString(fileContentFoo)
		require.NoError(t, err, "Failed to write test file")

		file, err = os.Create(apiFilePathBar)
		require.NoError(t, err, "Failed to create test file")
		defer file.Close()

		_, err = file.WriteString(fileContentBar)
		require.NoError(t, err, "Failed to write test file")

		o := &CodeGen{
			stderr:       &bytes.Buffer{},
			specPath:     "",
			operations:   map[string][]string{"foo": {"Foo", "Bar"}, "bar": {"Baz", "Qux"}},
			handlersPath: tmpDir,
		}

		err = o.ensureFunctionMethods()
		require.NoError(t, err)

		// Now an extra operation from another tag is misplaced
		tag := "bar"
		newMethod := "ExtraFoo"
		o.operations[tag] = append(o.operations[tag], newMethod)
		fileContentFooExtra := fileContentFoo + fmt.Sprintf("\nfunc (h *Handlers) %s() {}", newMethod)

		file, err = os.Create(apiFilePathFoo)
		require.NoError(t, err, "Failed to create test file")
		defer file.Close()

		_, err = file.WriteString(fileContentFooExtra)
		require.NoError(t, err, "Failed to write test file")

		err = o.ensureFunctionMethods()
		require.NoError(t, err)

		// TODO: check both files have swapped handlers
		// not testing  anything right now
		// var found bool
		// for _, decl := range fileContentFooExtra.Decls {
		// 	if fd, ok := decl.(*dst.FuncDecl); ok {
		// 		if fd.Name.Name == newMethod {
		// 			found = true
		// 			break
		// 		}
		// 	}
		// }
		// if !found {
		// 	t.Errorf("Handlers method '%s' not found in the second file after appending.", methodNameToRemove)
		// }
	})

	t.Run("missing api file", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()
		apiFilePathFoo := filepath.Join(tmpDir, "api_foo.go")
		apiFilePathBar := filepath.Join(tmpDir, "api_bar.go")

		file, err := os.Create(apiFilePathFoo)
		require.NoError(t, err, "Failed to create test file")
		defer file.Close()

		_, err = file.WriteString(fileContentFoo)
		require.NoError(t, err, "Failed to write test file")

		fileBar, err := os.Create(apiFilePathBar)
		require.NoError(t, err, "Failed to create test file")
		defer fileBar.Close()

		_, err = fileBar.WriteString(fileContentBar)
		require.NoError(t, err, "Failed to write test file")

		o := &CodeGen{
			stderr:       &bytes.Buffer{},
			specPath:     "",
			operations:   map[string][]string{"foo": {"Foo", "Bar"}, "bar": {"Baz", "Qux"}},
			handlersPath: tmpDir,
		}

		err = o.ensureFunctionMethods()
		require.NoError(t, err)

		// Now an extra operation from another tag is misplaced
		tag := "newtag"
		newMethod := "ExtraFoo"
		o.operations[tag] = append(o.operations[tag], newMethod)
		fileContentFooExtra := fileContentFoo + fmt.Sprintf("\nfunc (h *Handlers) %s() {}", newMethod)

		file, err = os.Create(apiFilePathFoo)
		require.NoError(t, err, "Failed to create test file")
		defer file.Close()

		_, err = file.WriteString(fileContentFooExtra)
		require.NoError(t, err, "Failed to write test file")

		err = o.ensureFunctionMethods()
		require.Error(t, err)
		assert.Contains(t, err.Error(), fmt.Sprintf("misplaced method for operation ID %q - should be in api_%s.go (file does not exist)", newMethod, tag))
	})
}

func TestEnsureFunctionMethods_MissingMethod(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	fileContent := `
package rest

type Handlers struct{}

func (h *Handlers) Foo() {}
func (h *Handlers) Bar() {}
`

	apiFilePath := tmpDir + "/api_foo.go"
	file, err := os.Create(apiFilePath)
	require.NoError(t, err, "Failed to create test file")
	defer file.Close()

	_, err = file.WriteString(fileContent)
	require.NoError(t, err, "Failed to write test file")

	cg := &CodeGen{
		stderr:       &bytes.Buffer{},
		specPath:     "",
		operations:   map[string][]string{"foo": {"Foo", "Bar"}},
		handlersPath: tmpDir,
	}

	err = cg.ensureFunctionMethods()
	require.NoError(t, err)

	cg.operations["foo"] = []string{"Foo", "Bar", "Baz"}

	err = cg.ensureFunctionMethods()
	require.Error(t, err, "expected error due to missing Baz, not misplaced in another handler")

	assert.Contains(t, err.Error(), "missing function method for operation ID \"Baz\"")
}

func TestRemoveAndAppendHandlersMethod(t *testing.T) {
	t.Parallel()

	sourceCode := `
package main

type Handlers struct{}

func (h *Handlers) GetSomething() {
	return "GetSomething"
}

func (h *Handlers) UpdateSomething() {
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
				t.Errorf("Handlers method '%s' still found in the first file after removal.", methodNameToRemove)
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
	if err := printer.Fprint(&buf, &token.FileSet{}, secondFile); err != nil {
		fmt.Println("Error printing AST:", err)
		return
	}

	fmt.Printf("buf.String(): %v\n", buf.String())

	if !found {
		t.Errorf("Handlers method '%s' not found in the second file after appending.", methodNameToRemove)
	}
}
