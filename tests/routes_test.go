package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestPingRoute(t *testing.T) {
	// TODO need helper func to create handlers and run all Register()
	router := gen.NewRouter()

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}

func TestCreateUserRoute(t *testing.T) {
	var buf bytes.Buffer

	// TODO move to helpers
	// then have teardown and setup as needed per domain
	// teardownSuite := setupSuite(t)
	// defer teardownSuite(t)
	// cleanest: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
	pool := NewDB(t)
	environment.Pool = pool
	environment.Logger = zaptest.NewLogger(t)

	// TODO log responses: add to helpers
	// https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin

	router := gen.NewRouter()

	type Input struct {
		User interface{}
	}

	type Output struct {
		Status int
	}

	cases := []struct {
		Input  Input
		Output Output
	}{
		{
			Input{
				User: models.CreateUserRequest{
					Email:    "email",
					Password: "password",
					Username: "username",
				}},
			Output{Status: http.StatusOK},
		},
		{
			Input{
				User: struct {
					Bad string `json:"bad,omitempty"`
				}{"bad"}},
			Output{Status: http.StatusBadRequest},
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if err := json.NewEncoder(&buf).Encode(tt.Input.User); err != nil {
				t.Errorf("%v", err)
			}

			t.Logf("%v", &buf)

			req, err := http.NewRequest(http.MethodPost, os.Getenv("API_VERSION")+"/user", &buf)
			if err != nil {
				t.Errorf("%v", err)
			}

			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)
			t.Logf("%v", resp)
			assert.Equal(t, tt.Output.Status, resp.Code)
		})
	}
}
