package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

type OIDCConfig struct {
	ClientID     string  `env:"OIDC_CLIENT_ID"`
	ClientSecret string  `env:"OIDC_CLIENT_SECRET"`
	Issuer       string  `env:"OIDC_ISSUER"`
	Scopes       string  `env:"OIDC_SCOPES"`
	Domain       string  `env:"OIDC_DOMAIN"`
	ServerPort   *string `env:"OIDC_SERVER_PORT"`
}

type PostgresConfig struct {
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Server   string `env:"POSTGRES_SERVER"`
	DB       string `env:"POSTGRES_DB"`
}

type RedisConfig struct {
	DB   string `env:"REDIS_DB"`
	Host string `env:"REDIS_HOST"`
}

// UPDATE:
// config is used from values in environment and is just a convenience struct instead of
// an error prone os.getenv. this way everything is loaded at once on startup and is validated.

// sharing config between packages:
// os.getenv is very convenient, but we can get close: internal.Config() returns the already built and validated *internal.AppConfig{}
// Returning a copy every single time would be too bad. its an internal package so
// just need to ensure nothing silly is being done like overwriting it... therefore a pointer is ok

// AppConfig contains app settings which are read from a config file. Excluded fields from JSON are read from environment variables.
type AppConfig struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	OIDC     OIDCConfig

	Domain     string `env:"DOMAIN"`
	APIPort    string `env:"API_PORT"`
	APIVersion string `env:"API_VERSION"`
	APIPrefix  string `env:"API_PREFIX"`
	AppEnv     string `env:"APP_ENV"`
	SigningKey string `env:"SIGNING_KEY"`
}

// don't. instead do it in NewConfig, so calling it without an error is enough to know its valid.
func (ac *AppConfig) Validate() error {
	// see https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go and other examples
	return nil
}

// NewAppConfig returns a new AppConfig.
func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

// Returns the directory of the file this function lives in.
func getFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(b))
}

var localConfigPath = filepath.Join(getFileRuntimeDirectory(), "config/%s.json")

// GenerateConfigTemplate creates a template.json config file for reference.
func GenerateConfigTemplate() error {
	cfg, err := json.MarshalIndent(&AppConfig{}, "", "  ")
	if err != nil {
		return errors.Wrap(err, "could not marshal template config json")
	}
	if err := os.WriteFile(fmt.Sprintf(localConfigPath, "template"), cfg, 0o777); err != nil {
		return errors.Wrap(err, "could not save template config json")
	}

	return nil
}
