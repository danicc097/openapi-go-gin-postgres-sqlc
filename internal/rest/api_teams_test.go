package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO
func TestCreateTeamRoute(t *testing.T) {
	t.Parallel()

	srv, err := runTestServer(t, testPool, []gin.HandlerFunc{})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}

	t.Cleanup(func() {
		srv.Close()
	})

	ff := newTestFixtureFactory(t)

	t.Run("authenticated user", func(t *testing.T) {
		t.Parallel()

		role := models.RoleAdvancedUser
		scopes := []models.Scope{models.ScopeProjectSettingsWrite}

		ufixture, err := ff.CreateUser(context.Background(), resttestutil.CreateUserParams{
			Role:       role,
			WithAPIKey: true,
			Scopes:     scopes,
		})
		if err != nil {
			t.Fatalf("ff.CreateUser: %s", err)
		}

		req, err := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/user/me", &bytes.Buffer{})
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Add("x-api-key", ufixture.APIKey.APIKey)

		resp := httptest.NewRecorder()

		srv.Handler.ServeHTTP(resp, req)

		ures := UserResponse{User: ufixture.User, Role: role, Scopes: ufixture.User.Scopes}

		res, err := json.Marshal(ures)
		if err != nil {
			t.Fatalf("could not marshal user fixture")
		}

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, string(res), resp.Body.String())
	})
}

// TODO test PATCH update user and update user authorixation routes
// func TestUpdateUserRoute(t *testing.T) {
// 	t.Parallel()

// 	srv, err := runTestServer(t, pool, []gin.HandlerFunc{})
// 	if err != nil {
// 		t.Fatalf("Couldn't run test server: %s\n", err)
// 	}
// 	defer srv.Close()

// 	type params struct {
// 		user any
// 	}

// 	type want struct {
// 		status int
// 	}

// 	testCases := []struct {
// 		name   string
// 		params params
// 		want   want
// 	}{
// 		{
// 			"ValidParams",
// 			params{
// 				user: models.CreateUserRequest{
// 					Email:    testutil.RandomEmail(),
// 					Password: "password",
// 					Username: testutil.RandomName(),
// 				},
// 			},
// 			want{status: http.StatusOK},
// 		},
// 		{
// 			"UsernameValidationFailed",
// 			params{
// 				user: models.CreateUserRequest{
// 					Email:    testutil.RandomEmail(),
// 					Password: "password",
// 					Username: "[]]]",
// 				},
// 			},
// 			want{status: http.StatusBadRequest},
// 		},
// 		{
// 			"EmailValidationFailed",
// 			params{
// 				user: models.CreateUserRequest{
// 					Email:    "bad",
// 					Password: "password",
// 					Username: testutil.RandomName(),
// 				},
// 			},
// 			want{status: http.StatusBadRequest},
// 		},
// 		{
// 			"PasswordValidationFailed",
// 			params{
// 				user: models.CreateUserRequest{
// 					Email:    testutil.RandomEmail(),
// 					Password: "short",
// 					Username: testutil.RandomName(),
// 				},
// 			},
// 			want{status: http.StatusBadRequest},
// 		},
// 		{
// 			"BadParams",
// 			params{
// 				user: struct {
// 					Bad string `json:"bad,omitempty"`
// 				}{"bad"},
// 			},
// 			want{status: http.StatusBadRequest},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()
// 			var buf bytes.Buffer

// 			if err := json.NewEncoder(&buf).Encode(tc.params.user); err != nil {
// 				t.Errorf("%v", err)
// 			}

// 			t.Logf("%v", &buf)

// 			req, err := http.NewRequest(http.MethodPost, os.Getenv("API_VERSION")+"/user", &buf)
// 			req.Header.Add("Content-Type", "application/json")
// 			req.Header.Add("x-api-key", "dummy-key")
// 			if err != nil {
// 				t.Errorf("%v", err)
// 			}

// 			resp := httptest.NewRecorder()

// 			srv.Handler.ServeHTTP(resp, req)
// 			t.Logf("%v", resp)
// 			assert.Equal(t, tc.want.status, resp.Code)
// 		})
// 	}
// }
