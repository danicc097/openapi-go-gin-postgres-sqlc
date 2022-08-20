package tests_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/tests"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	srv, err := tests.Run(t, "../.env", ":8099")
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}

func TestCreateUserRoute(t *testing.T) {
	// TODO move to helpers
	// then have teardown and setup as needed per domain
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	// cleanest: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
	// TODO log responses: add to helpers
	// https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin

	// FIXME missing programmatic migrations, etc. from removed NewDB, see old commits

	var buf bytes.Buffer

	srv, err := tests.Run(t, "../.env", ":8099")
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	type Input struct {
		User interface{}
	}

	type Output struct {
		Status int
	}

	cases := []struct {
		Name   string
		Input  Input
		Output Output
	}{
		{
			"Valid params",
			Input{
				User: models.CreateUserRequest{
					Email:    "email",
					Password: "password",
					Username: "username",
				}},
			Output{Status: http.StatusOK},
		},
		{
			"Bad params",
			Input{
				User: struct {
					Bad string `json:"bad,omitempty"`
				}{"bad"}},
			Output{Status: http.StatusBadRequest},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			if err := json.NewEncoder(&buf).Encode(test.Input.User); err != nil {
				t.Errorf("%v", err)
			}

			t.Logf("%v", &buf)

			req, err := http.NewRequest(http.MethodPost, os.Getenv("API_VERSION")+"/user", &buf)
			if err != nil {
				t.Errorf("%v", err)
			}

			resp := httptest.NewRecorder()

			srv.Handler.ServeHTTP(resp, req)
			t.Logf("%v", resp)
			assert.Equal(t, test.Output.Status, resp.Code)
		})
	}
}

func setupSuite(t *testing.T) func(t *testing.T) {

	return func(t *testing.T) {
	}
}
