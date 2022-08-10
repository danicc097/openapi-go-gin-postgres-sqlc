package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestPingRoute(t *testing.T) {
	router := gen.NewRouter()

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}

func TestCreateUserRoute(t *testing.T) {
	var buf bytes.Buffer

	provider, err := vault.New()
	if err != nil {
		t.Errorf("%v", err)
	}

	conf := envvar.New(provider)

	// TODO set postgres database to postgres_test
	// and run migrations up programatically

	pool, err := postgresql.New(conf)
	if err != nil {
		t.Errorf("%v", err)
	}

	environment.Pool = pool
	environment.Logger = zaptest.NewLogger(t)

	router := gen.NewRouter()
	f := models.CreateUserRequest{
		Email:    "email",
		Password: "password",
		Username: "username",
	}

	if err := json.NewEncoder(&buf).Encode(f); err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("%v", &buf)

	req, err := http.NewRequest(http.MethodPost, os.Getenv("API_VERSION")+"/user", &buf)
	if err != nil {
		t.Errorf("%v", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	// assert.Equal(t, http.StatusOK, resp.Code) // FIXME
}
