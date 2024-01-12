package internal

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	configLock = &sync.Mutex{}

	// Config returns the app global config initialized from environment variables.
	// [Read] locks not needed if there are no writes involved. Config is only populated at startup so there won't be any more writes.
	Config *AppConfig
)

type OIDCConfig struct {
	ClientID     string `env:"OIDC_CLIENT_ID"`
	ClientSecret string `env:"OIDC_CLIENT_SECRET"`
	Issuer       string `env:"OIDC_ISSUER"`
	Scopes       string `env:"OIDC_SCOPES"`
	Domain       string `env:"OIDC_DOMAIN"`
}

type PostgresConfig struct {
	// Port represents the db port to use in the application, depending on setup (dockerized or not).
	Port         int    `env:"DB_PORT"`
	User         string `env:"POSTGRES_USER"`
	Password     string `env:"POSTGRES_PASSWORD"`
	Server       string `env:"POSTGRES_SERVER"`
	DB           string `env:"POSTGRES_DB"`
	TraceEnabled bool   `env:"POSTGRES_TRACE,false"`
}

type RedisConfig struct {
	DB   int    `env:"REDIS_DB"`
	Host string `env:"REDIS_HOST"`
}

type SuperAdminConfig struct {
	DefaultEmail string `env:"DEFAULT_SUPERADMIN_EMAIL"`
}

// TODO: handle actual custom types in config.
type AppEnv = string

const (
	AppEnvDev  AppEnv = "dev"
	AppEnvProd AppEnv = "prod"
	AppEnvCI   AppEnv = "ci"
	AppEnvE2E  AppEnv = "e2e"
)

// AppConfig contains app settings.
type AppConfig struct {
	Postgres   PostgresConfig
	Redis      RedisConfig
	OIDC       OIDCConfig
	SuperAdmin SuperAdminConfig

	Domain         string `env:"DOMAIN"`
	APIPort        string `env:"API_PORT"`
	APIVersion     string `env:"API_VERSION"`
	APIPrefix      string `env:"API_PREFIX"`
	AppEnv         AppEnv `env:"APP_ENV"`
	SigningKey     string `env:"SIGNING_KEY"`
	BuildVersion   string `env:"BUILD_VERSION"`
	CookieDomain   string `env:"COOKIE_DOMAIN"`
	LoginCookieKey string `env:"LOGIN_COOKIE_KEY"`

	ScopePolicyPath string `env:"SCOPE_POLICY_PATH"`
	RolePolicyPath  string `env:"ROLE_POLICY_PATH"`
}

// NewAppConfig initializes app config from current environment variables.
// config can be replaced with subsequent calls.
func NewAppConfig() error {
	configLock.Lock()
	defer configLock.Unlock()

	cfg := &AppConfig{}

	if err := loadEnvToConfig(cfg); err != nil {
		return fmt.Errorf("loadEnvToConfig: %w", err)
	}

	Config = cfg

	return nil
}

// loadEnvToConfig loads env vars to a given struct based on an `env` tag.
func loadEnvToConfig(config any) error {
	cfg := reflect.ValueOf(config)

	if cfg.Kind() == reflect.Pointer {
		cfg = cfg.Elem()
	}

	for idx := 0; idx < cfg.NumField(); idx++ {
		fType := cfg.Type().Field(idx)
		fld := cfg.Field(idx)

		if fld.Kind() == reflect.Struct {
			if !fld.CanInterface() { // unexported
				continue
			}
			if err := loadEnvToConfig(fld.Addr().Interface()); err != nil {
				return fmt.Errorf("nested struct %q env loading: %w", cfg.Type().Field(idx).Name, err)
			}
		}

		if !fld.CanSet() {
			continue
		}

		if env, ok := fType.Tag.Lookup("env"); ok && len(env) > 0 {
			err := setEnvToField(env, fld)
			if err != nil {
				return fmt.Errorf("could not set %q to %q: %w", env, cfg.Type().Field(idx).Name, err)
			}
		}
	}

	return nil
}

func splitEnvTag(s string) (string, string) {
	x := strings.Split(s, ",")
	if len(x) == 1 {
		return x[0], ""
	}
	return x[0], x[1]
}

func setEnvToField(envTag string, field reflect.Value) error {
	envvar, defaultVal := splitEnvTag(envTag)
	val, present := os.LookupEnv(envvar)

	if !present && field.Kind() != reflect.Pointer && defaultVal == "" {
		return fmt.Errorf("%s is not set but required", envvar)
	}

	if !present && field.Kind() != reflect.Pointer && defaultVal != "" {
		val = defaultVal
	}

	var isPtr bool

	kind := field.Kind()
	if kind == reflect.Pointer {
		kind = field.Type().Elem().Kind()
		isPtr = true
	}

	if val == "" && isPtr && kind != reflect.String {
		return nil
	}

	switch kind {
	case reflect.String:
		if !present && isPtr {
			setVal[*string](false, field, nil) // since default val is always ""

			return nil
		}
		setVal(isPtr, field, val)
	case reflect.Int:
		v, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("could not convert %s to int: %w", envvar, err)
		}
		setVal(isPtr, field, v)
	case reflect.Bool:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return fmt.Errorf("could not convert %s to bool: %w", envvar, err)
		}
		setVal(isPtr, field, v)
	default:
		return fmt.Errorf("unsupported type for env tag %q: %T", envvar, field.Interface())
	}

	return nil
}

func setVal[T any](isPtr bool, field reflect.Value, v T) {
	if isPtr {
		field.Set(reflect.ValueOf(&v))
	} else {
		field.Set(reflect.ValueOf(v))
	}
}

// Returns the directory of the file this function lives in.
func getFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(b))
}
