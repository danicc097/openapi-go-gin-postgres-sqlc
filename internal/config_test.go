package internal_test

import (
	"os"
	"strings"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewAppConfig(t *testing.T) {
	type nestedCfg struct {
		Name string `env:"TEST_CFG_NAME"`
	}
	type cfg struct {
		NestedCfg nestedCfg
		Length    int `env:"TEST_CFG_LENGTH"`
	}

	type params struct {
		name        string
		want        *cfg
		errContains string
		env         map[string]string
	}

	tests := []params{
		{
			name: "correct env",
			want: &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10},
			env:  map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LENGTH": "10"},
		},
		{
			name:        "non pointer field corresponding envvar not set",
			env:         map[string]string{"TEST_CFG_LENGTH": "10"},
			errContains: `could not set "TEST_CFG_NAME" to "Name": TEST_CFG_NAME is not set but required`,
		},
		{
			name: "empty but set env vars dont raise an error",
			env:  map[string]string{"TEST_CFG_NAME": "", "TEST_CFG_LENGTH": "10"},
			want: &cfg{NestedCfg: nestedCfg{Name: ""}, Length: 10},
		},
		{
			name:        "bad env conversion",
			env:         map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LENGTH": "aaa"},
			errContains: `could not set "TEST_CFG_LENGTH" to "Length": could not convert TEST_CFG_LENGTH to int`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			curEnviron := os.Environ()
			for _, entry := range curEnviron {
				parts := strings.SplitN(entry, "=", 2)
				os.Unsetenv(parts[0])
			}
			for k, v := range tt.env {
				os.Setenv(k, v)
			}
			t.Cleanup(func() {
				for k := range tt.env {
					os.Unsetenv(k)
				}
				for _, entry := range curEnviron {
					parts := strings.SplitN(entry, "=", 2)
					os.Setenv(parts[0], parts[1])
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
