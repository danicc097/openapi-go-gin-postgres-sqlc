package pregen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

			og := New(&stderr, path.Join(baseDir, tc.File), "")

			err := og.analyzeSpec()
			if err != nil && tc.ErrContains != "" {
				assert.ErrorContains(t, err, tc.ErrContains)
			} else if err != nil {
				t.Fatalf("err: %s\nstderr: %s\n", err, &stderr)
			}
		})
	}
}

func TestEnsureFunctionsForOperationIDs(t *testing.T) {
	t.Parallel()

	tempDir, err := ioutil.TempDir("", "test_pregen")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	defaultAPISource := `
package rest

type Handlers struct{}

func (h *Handlers) MyProviderLogin(c *gin.Context) {}

func someOtherFunction() {}
`
	defaultAPIFilePath := filepath.Join(tempDir, "api_default.go")
	assert.NoError(t, ioutil.WriteFile(defaultAPIFilePath, []byte(defaultAPISource), 0o666))

	customAPISource := `
package rest

type Handlers struct{}

func (h *Handlers) MyCustomFunction(c *gin.Context) {}

func (h *Handlers) UnRelatedFunction(c *gin.Context) {}

func anotherFunction() {}
`
	customAPIFilePath := filepath.Join(tempDir, "api_custom.go")
	assert.NoError(t, ioutil.WriteFile(customAPIFilePath, []byte(customAPISource), 0o666))

	specSource := `
{
  "openapi": "3.0.0",
	"info": {"description": "a", "version":"1", "title": "1"},
  "paths": {
    "/my-provider/login": {
      "post": {
        "operationId": "MyProviderLogin",
        "tags": ["default"],
        "responses": { "200": { "description": "Successful response" } }
      }
    },
    "/my-custom/function": {
      "get": {
        "operationId": "MyCustomFunction",
        "tags": ["custom"],
        "responses": { "200": { "description": "Successful response" } }
      }
    }
  }
}
`
	fmt.Printf("specSource: %v\n", specSource)
	// err = EnsureFunctionsForOperationIDs(specFilePath, opIDAuthInfoFilePath)
	assert.NoError(t, err)
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
