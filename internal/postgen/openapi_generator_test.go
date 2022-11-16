package postgen

import (
	"bytes"
	"path"
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
			`path "/pet/ConflictEndpointPet": method "GET": operationId is required for postgen`,
		},
		{
			"more_than_one_tag",
			"more_than_one_tag.yaml",
			`path "/pet/ConflictEndpointPet": method "GET": at most one tag is permitted for postgen`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var stderr bytes.Buffer

			og := NewOpenapiGenerator(&Conf{}, &stderr, path.Join(baseDir, tc.File))

			err := og.analyzeSpec()
			if err != nil && tc.ErrContains != "" {
				assert.ErrorContains(t, err, tc.ErrContains)
			} else if err != nil {
				t.Fatalf("err: %s\nstderr: %s\n", err, &stderr)
			}
		})
	}
}
