package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/tests"
	"github.com/stretchr/testify/assert"
)

func setupSuite(t *testing.T) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		t.Helper()
	}
}

func TestPingRoute(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	srv := tests.NewServer(t)
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}

func TestCreateUserRoute(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	srv := tests.NewServer(t)
	defer srv.Close()

	var buf bytes.Buffer

	type Params struct {
		User interface{}
	}

	type Want struct {
		Status int
	}

	cases := []struct {
		Name   string
		Params Params
		Want   Want
	}{
		{
			"Valid params",
			Params{
				User: models.CreateUserRequest{
					Email:    "email",
					Password: "password",
					Username: "username",
				}},
			Want{Status: http.StatusOK},
		},
		{
			"Bad params",
			Params{
				User: struct {
					Bad string `json:"bad,omitempty"`
				}{"bad"}},
			Want{Status: http.StatusBadRequest},
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			if err := json.NewEncoder(&buf).Encode(test.Params.User); err != nil {
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
			assert.Equal(t, test.Want.Status, resp.Code)
		})
	}
}
