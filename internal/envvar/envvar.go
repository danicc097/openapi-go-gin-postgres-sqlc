// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package envvar

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// Provider ...
type Provider interface {
	Get(key string) (string, error)
}

// Configuration ...
type Configuration struct{}

// Load reads the env filename and loads it into ENV for the current process.
// It also initializes/replaces app configuration.
func Load(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return internal.NewErrorf(models.ErrorCodeUnknown, "%s", fmt.Sprintf("loading %s env var file: %s", filename, err))
	}

	if err := internal.NewAppConfig(); err != nil {
		return internal.WrapErrorf(err, models.ErrorCodeUnknown, "internal.NewAppConfig")
	}

	return nil
}

// New ...
func New() *Configuration {
	return &Configuration{}
}

// Get returns the value from environment variable `<key>`. When an environment variable `<key>_SECURE` exists
// the provider is used for getting the value.
func (c *Configuration) Get(key string) (string, error) {
	res := os.Getenv(key)

	return res, nil
}

var errEnvVarEmpty = errors.New("env var empty")

// GetEnv returns an environment variable's value or a default if empty.
func GetEnv(key, dft string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return dft
	}

	return v
}
