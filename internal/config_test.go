package internal_test

import (
	"os"
	"strings"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewAppConfig(t *testing.T) {
	type nestedCfg struct {
		Name string `env:"TEST_CFG_NAME"`
	}
	// NOTE: zero need and counterproductive to allow pointer nested structs for config
	type cfg struct {
		NestedCfg      nestedCfg
		Length         int     `env:"TEST_CFG_LEN"`
		OptionalLength *int    `env:"TEST_CFG_OPT_LEN"`
		OptionalString *string `env:"TEST_CFG_STRING_PTR"`
	}

	type params struct {
		name        string
		want        *cfg
		errContains string
		environ     map[string]string
	}

	tests := []params{
		{
			name:    "correct env",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalLength: pointers.New(40)},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_OPT_LEN": "40"},
		},
		{
			name:    "correct env with missing pointer fields",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10"},
		},
		{
			name:        "non pointer field corresponding envvar not set",
			environ:     map[string]string{"TEST_CFG_LEN": "10"},
			errContains: `could not set "TEST_CFG_NAME" to "Name": TEST_CFG_NAME is not set but required`,
		},
		{
			name:    "empty but set envvar string pointer field is not nil",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalString: pointers.New("")},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_STRING_PTR": ""},
		},
		{
			name:    "unset envvar string pointer field is nil",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10"},
		},
		{
			name:        "bad env conversion",
			environ:     map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "aaa"},
			errContains: `could not set "TEST_CFG_LEN" to "Length": could not convert TEST_CFG_LEN to int`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			curEnviron := os.Environ()
			for _, entry := range curEnviron {
				parts := strings.SplitN(entry, "=", 2)
				os.Unsetenv(parts[0])
			}
			for k, v := range tt.environ {
				os.Setenv(k, v)
			}
			t.Cleanup(func() {
				for k := range tt.environ {
					os.Unsetenv(k)
				}
				for _, envvar := range curEnviron {
					ss := strings.SplitN(envvar, "=", 2)
					os.Setenv(ss[0], ss[1])
				}
			})

			c := &cfg{}
			err := internal.LoadEnvToConfig(c)
			if err != nil && tt.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tt.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.Contains(t, err.Error(), tt.errContains)

				return
			}

			assert.Equal(t, tt.want, c)
		})
	}
}
