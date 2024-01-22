package internal

import (
	"fmt"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

// nolint: paralleltest // cannot set env in parallel tests
func TestNewAppConfig(t *testing.T) {
	type nestedCfg struct {
		Name string `env:"TEST_CFG_NAME"`
	}
	// NOTE: zero need and counterproductive to allow pointer nested structs for config.
	type cfg struct {
		NestedCfg       nestedCfg
		Length          int     `env:"TEST_CFG_LEN"`
		OptionalLength  *int    `env:"TEST_CFG_OPT_LEN"`
		OptionalString  *string `env:"TEST_CFG_STRING_PTR"`
		OptionalBool    *bool   `env:"TEST_CFG_BOOL_PTR"`
		BoolWithDefault bool    `env:"TEST_CFG_BOOL_DEFAULT,false"`
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
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalLength: pointers.New(40), BoolWithDefault: true},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_OPT_LEN": "40", "TEST_CFG_BOOL_DEFAULT": "true"},
		},
		{
			name:    "correct env with missing pointer fields",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, BoolWithDefault: false},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10"},
		},
		{
			name:    "empty but set envvar string pointer field is not nil",
			want:    &cfg{NestedCfg: nestedCfg{Name: "name"}, Length: 10, OptionalString: pointers.New("2")},
			environ: map[string]string{"TEST_CFG_NAME": "name", "TEST_CFG_LEN": "10", "TEST_CFG_STRING_PTR": "2"},
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
		{
			name:        "non pointer field corresponding envvar not set",
			environ:     map[string]string{"TEST_CFG_LEN": "10"},
			errContains: `could not set "TEST_CFG_NAME" to "Name": TEST_CFG_NAME is not set but required`,
		},
	}
	// nolint: paralleltest // cannot set env in parallel tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.environ {
				t.Setenv(k, v)
			}

			c := &cfg{}
			err := loadEnvToConfig(c)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.ErrorContains(t, err, tc.errContains)

				return
			}

			assert.Equal(t, tc.want, c)
		})
	}
}

type MyEnum string

const (
	Value1 MyEnum = "Value1"
	Value2 MyEnum = "Value2"
)

// Decode decodes an env var value.
func (e *MyEnum) Decode(value string) error {
	switch value {
	case "1":
		*e = Value1
	case "2":
		*e = Value2
	default:
		return fmt.Errorf("invalid value for MyEnum: %v", value)
	}
	return nil
}

// nolint: paralleltest // cannot set env in parallel tests
func TestEnumDecoderConfig(t *testing.T) {
	type cfg struct {
		EnumWithDefault MyEnum  `env:"TEST_CFG_ENUM,1"`
		OptionalEnum    *MyEnum `env:"TEST_CFG_ENUM_OPT"`
	}

	type params struct {
		name        string
		want        *cfg
		errContains string
		environ     map[string]string
	}

	tests := []params{
		{
			name:    "correct enum field value",
			want:    &cfg{OptionalEnum: pointers.New(Value2), EnumWithDefault: Value1},
			environ: map[string]string{"TEST_CFG_ENUM_OPT": "2"},
		},
		{
			name:    "correct default override",
			want:    &cfg{EnumWithDefault: Value2},
			environ: map[string]string{"TEST_CFG_ENUM": "2"},
		},
		{
			name:        "invalid enum field value",
			errContains: `invalid value for MyEnum: badvalue`,
			environ:     map[string]string{"TEST_CFG_ENUM": "badvalue"},
		},
	}
	// nolint: paralleltest // cannot set env in parallel tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.environ {
				t.Setenv(k, v)
			}

			c := &cfg{}
			err := loadEnvToConfig(c)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.ErrorContains(t, err, tc.errContains)

				return
			}

			assert.Equal(t, tc.want, c)
		})
	}
}

func TestBadAppConfig(t *testing.T) {
	type nestedCfg struct {
		Name string `env:"TEST_CFG_NAME"`
	}

	t.Run("struct_has_env_tag", func(t *testing.T) {
		errContains := "unsupported type for env tag"
		environ := map[string]string{"ENV_ON_STRUCT": "10", "TEST_CFG_NAME": "name"}
		for k, v := range environ {
			t.Setenv(k, v)
		}

		type cfg struct {
			NestedCfg nestedCfg `env:"ENV_ON_STRUCT"`
		}
		c := &cfg{}
		err := loadEnvToConfig(c)
		if err != nil && errContains == "" {
			t.Errorf("unexpected error: %v", err)

			return
		}
		if errContains != "" {
			if err == nil {
				t.Errorf("expected error but got nothing")

				return
			}
			assert.ErrorContains(t, err, errContains)

			return
		}
	})
}
