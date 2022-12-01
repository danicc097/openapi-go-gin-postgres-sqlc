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
	ClientID     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	Issuer       string `json:"issuer" validate:"required"`
	Scopes       string `json:"scopes" validate:"required"`

	Domain     string  `json:"-" validate:"required"` // env var
	ServerPort *string `json:"-"`                     // optional,env var
}

type PostgresConfig struct {
	Port     string `json:"-" validate:"required"` // env var
	User     string `json:"-" validate:"required"` // env var
	Password string `json:"-" validate:"required"` // env var
	Server   string `json:"-" validate:"required"` // env var
	DB       string `json:"-" validate:"required"` // env var
}

type RedisConfig struct {
	DB   string `json:"-" validate:"required"` // env var
	Host string `json:"-" validate:"required"` // env var
}

// TODO frontend still needs .env.<env> for dynamic config.json
// without rebuilding.
// Postgres and backend also need .env shared.
// We could have DBConfig struct with json:"-"
// and on startup fill it in with whatever is in the current env,
// then validate it as usual.
// this way postgres can reuse the .env.* as usual
// and get typed config
// ^ do the same for .env values that are shared with frontend:
// TDLR whatever appears in .env.template must have json:"-" and env is loaded to the struct after unmarshalling config/*.json

// UPDATE:
// IMPORTANT: bin/project sources .env.* and is vital to work with. docker-compose also uses it...
// so in the end it the config struct will for the most part be filled with env vars

// sharing config between packages:
// os.getenv is very convenient.
// the closest we can get: internal.Config() returns the already built and validated *config.
// returning a copy every single time is too bad. its an internal package so
// just need to ensure nothing silly is being done like overwriting it...

// AppConfig contains app settings which are read from a config file. Excluded fields from JSON are read from environment variables.
type AppConfig struct {
	Postgres PostgresConfig `json:"-" validate:"required"`
	Redis    RedisConfig    `json:"-" validate:"required"`
	OIDC     OIDCConfig     `json:"oidc" validate:"required"`

	Domain     string `json:"-" validate:"required"`
	APIPort    string `json:"-" validate:"required"`
	APIVersion string `json:"-" validate:"required"`
	APIPrefix  string `json:"-" validate:"required"`
	AppEnv     string `json:"-" validate:"required"`
	SigningKey string `json:"signingKey" validate:"required"`
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
