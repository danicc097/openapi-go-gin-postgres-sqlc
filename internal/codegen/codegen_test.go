package codegen

import (
	"bytes"
	"os"
	"path"
	"strings"
	"testing"

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
			`path "/pet/ConflictEndpointPet": method "GET": operationId "Conflict-Endpoint-Pet" does not match pattern "^[a-zA-Z0-9]*$"`,
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
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			var stderr bytes.Buffer

			og := New(&stderr, path.Join(baseDir, tc.File), "", "")

			err := og.analyzeSpec()
			if err != nil && tc.ErrContains != "" {
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

func TestParseAndCheckFunctions(t *testing.T) {
	t.Parallel()

	fileContent := `
package rest

type Handlers struct{}

func (h *Handlers) Foo() {}
func (h *Handlers) Bar() {}
`

	opIDs := []string{"Foo", "Bar", "Baz"}

	err := parseAndCheckFunctions(strings.NewReader(fileContent), opIDs)
	require.Error(t, err, "expected error due to missing Baz")

	assert.Contains(t, err.Error(), "missing function method for operation ID 'Baz'")

	opIDs = []string{"Foo", "Bar"}
	err = parseAndCheckFunctions(strings.NewReader(fileContent), opIDs)
	require.NoError(t, err)
}

func TestEnsureFunctionMethods(t *testing.T) {
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
	require.Error(t, err, "expected error due to missing Baz")

	assert.Contains(t, err.Error(), "missing function method for operation ID 'Baz'")
}
