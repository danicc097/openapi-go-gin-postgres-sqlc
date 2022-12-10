package internal

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"sync"
)

var (
	lock = &sync.Mutex{}

	config *AppConfig
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
	Port     string `env:"DB_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Server   string `env:"POSTGRES_SERVER"`
	DB       string `env:"POSTGRES_DB"`
}

type RedisConfig struct {
	DB   string `env:"REDIS_DB"`
	Host string `env:"REDIS_HOST"`
}

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

// NewAppConfig returns a new AppConfig singleton from current environment variables.
func NewAppConfig() error {
	cfg := &AppConfig{}

	if config != nil {
		return errors.New("AppConfig has already been created")
	}

	lock.Lock()
	defer lock.Unlock()

	if err := LoadEnvToConfig(cfg); err != nil {
		return fmt.Errorf("LoadEnvToConfig: %w", err)
	}
	config = cfg

	return nil
}

// Config returns the current app config and panics if not initialized via NewAppConfig.
func Config() *AppConfig {
	if config == nil {
		panic("app configuration has not yet been initialized")
	}

	return config
}

// LoadEnvToConfig loads env vars to a given struct based on an `env` tag.
func LoadEnvToConfig(cfg any) error {
	st := reflect.ValueOf(cfg)

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	for idx := 0; idx < st.NumField(); idx++ {
		fType := st.Type().Field(idx)

		if f := st.Field(idx); f.Kind() == reflect.Struct {
			if !f.CanInterface() { // unexported
				continue
			}
			if err := LoadEnvToConfig(f.Addr().Interface()); err != nil {
				return fmt.Errorf("nested struct %s env loading: %w", f.Addr().String(), err)
			}
		}

		if !st.Field(idx).CanSet() {
			continue
		}

		if env, ok := fType.Tag.Lookup("env"); ok && len(env) > 0 {
			switch st.Field(idx).Kind() {
			// expand as needed
			case reflect.String:
				st.Field(idx).SetString(os.Getenv(env))
			case reflect.Int:
				v, err := strconv.Atoi(os.Getenv(env))
				if err != nil {
					return fmt.Errorf("could not convert %s to int: %w", env, err)
				}
				st.Field(idx).SetInt(int64(v))
			}
		}
	}

	return nil
}

// Returns the directory of the file this function lives in.
func getFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(b))
}

// var localConfigPath = filepath.Join(getFileRuntimeDirectory(), "config/%s.json")

// // GenerateConfigTemplate creates a template.json config file for reference.
// func GenerateConfigTemplate() error {
// 	cfg, err := json.MarshalIndent(&AppConfig{}, "", "  ")
// 	if err != nil {
// 		return errors.Wrap(err, "could not marshal template config json")
// 	}
// 	if err := os.WriteFile(fmt.Sprintf(localConfigPath, "template"), cfg, 0o777); err != nil {
// 		return errors.Wrap(err, "could not save template config json")
// 	}

// 	return nil
// }
