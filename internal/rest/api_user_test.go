package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, pool, []gin.HandlerFunc{})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	type params struct {
		user any
	}

	type want struct {
		status int
	}

	testCases := []struct {
		name   string
		params params
		want   want
	}{
		{
			"ValidParams",
			params{
				user: models.CreateUserRequest{
					Email:    testutil.RandomEmail(),
					Password: "password",
					Username: testutil.RandomName(),
				},
			},
			want{status: http.StatusOK},
		},
		{
			"UsernameValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    testutil.RandomEmail(),
					Password: "password",
					Username: "[]]]",
				},
			},
			want{status: http.StatusBadRequest},
		},
		{
			"EmailValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    "bad",
					Password: "password",
					Username: testutil.RandomName(),
				},
			},
			want{status: http.StatusBadRequest},
		},
		{
			"PasswordValidationFailed",
			params{
				user: models.CreateUserRequest{
					Email:    testutil.RandomEmail(),
					Password: "short",
					Username: testutil.RandomName(),
				},
			},
			want{status: http.StatusBadRequest},
		},
		{
			"BadParams",
			params{
				user: struct {
					Bad string `json:"bad,omitempty"`
				}{"bad"},
			},
			want{status: http.StatusBadRequest},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer

			if err := json.NewEncoder(&buf).Encode(tc.params.user); err != nil {
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
			assert.Equal(t, tc.want.status, resp.Code)
		})
	}
}
