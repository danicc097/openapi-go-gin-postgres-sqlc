package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/stretchr/testify/assert"
)

func setupSuite(t *testing.T) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		t.Helper()
	}
}

func TestPingRoute(t *testing.T) {
	t.Parallel()

	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/ping", nil)
	resp := httptest.NewRecorder()
	t.Logf("rqt: %s", req.URL)
	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}

func TestCreateUserRoute(t *testing.T) {
	t.Parallel()

	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	type params struct {
		user any
	}

	type want struct {
		status int
	}

	cases := []struct {
		name   string
		params params
		want   want
	}{
		{
			"ValidParams",
			params{
				user: models.CreateUserRequest{
					Email:    format.RandomEmail(),
					Password: "password",
					Username: format.RandomName(),
				}},
			want{status: http.StatusOK},
		},
		{
			"UsernameValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    format.RandomEmail(),
					Password: "password",
					Username: "[]]]",
				}},
			want{status: http.StatusBadRequest},
		},
		{
			"EmailValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    "bad",
					Password: "password",
					Username: format.RandomName(),
				}},
			want{status: http.StatusBadRequest},
		},
		{
			"PasswordValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    format.RandomEmail(),
					Password: "short",
					Username: format.RandomName(),
				}},
			want{status: http.StatusBadRequest},
		},
		{
			"BadParams",
			params{
				user: struct {
					Bad string `json:"bad,omitempty"`
				}{"bad"}},
			want{status: http.StatusBadRequest},
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer

			if err := json.NewEncoder(&buf).Encode(test.params.user); err != nil {
				t.Errorf("%v", err)
			}

			t.Logf("%v", &buf)

			req, err := http.NewRequest(http.MethodPost, os.Getenv("API_VERSION")+"/user", &buf)
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				t.Errorf("%v", err)
			}

			resp := httptest.NewRecorder()

			srv.Handler.ServeHTTP(resp, req)
			t.Logf("%v", resp)
			assert.Equal(t, test.want.status, resp.Code)
		})
	}
}
