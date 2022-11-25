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
	ClientID     string  `json:"clientId" validate:"required"`
	ClientSecret string  `json:"clientSecret" validate:"required"`
	Domain       string  `json:"domain" validate:"required"`
	Issuer       string  `json:"issuer" validate:"required"`
	Scopes       string  `json:"scopes" validate:"required"`
	ServerPort   *string `json:"serverPort"`
}

type PostgresConfig struct {
	Port     string `json:"-" validate:"required"`
	User     string `json:"-" validate:"required"`
	Password string `json:"-" validate:"required"`
	Server   string `json:"-" validate:"required"`
	Database string `json:"-" validate:"required"`
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

// sharing config between packages:
// os.getenv is very convenient.
// the closest we can get: internal.Config() returns the already built and validated *config.
// returning a copy every single time is too bad. its an internal package so
// just need to ensure nothing silly is being done like overwriting it...

// AppConfig contains app settings which are read from a config file. Excluded fields from JSON are read from environment variables.
type AppConfig struct {
	Postgres   PostgresConfig `json:"-" validate:"required"`
	OIDC       OIDCConfig     `json:"oidc" validate:"required"`
	SigningKey string         `json:"signingKey" validate:"required"`
}

func (ac *AppConfig) Validate() error {
	// see https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go and other examples
	return nil
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
