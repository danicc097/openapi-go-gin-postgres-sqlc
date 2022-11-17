package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestAdminPingRoute(t *testing.T) {
	t.Parallel()

	// TODO helper method to create unique users for tests with specific roles, ranks, scopes, api keys...
	// we ensure they're all unique to each test and everything can therefore be parallelized.

	logger := zaptest.NewLogger(t)
	usvc := services.NewUser(repos.NewUserWrapped(postgresql.NewUser(), otelName, repos.UserWrappedConfig{}, nil), logger)
	authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("services.NewAuthorization: %v", err)
	}
	authnsvc := services.NewAuthentication(logger, usvc, pool)

	ff := testutil.NewFixtureFactory(usvc, pool, authnsvc, authzsvc)
	ufixture, err := ff.CreateUser(context.Background(), testutil.CreateUserParams{
		Role:       models.RoleAdmin,
		WithAPIKey: true,
	})
	if err != nil {
		t.Fatalf("ff.CreateUser: %s", err)
	}

	srv, err := runTestServer(t, pool, []gin.HandlerFunc{func(c *gin.Context) {
		ctxWithUser(c, ufixture.User)
	}})
	if err != nil {
		t.Fatalf("Couldn't run test server: %s\n", err)
	}
	defer srv.Close()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/admin/ping", nil)
	// FIXME see middleware.auth_test for fixture factory
	req.Header.Add("x-api-key", ufixture.APIKey)

	srv.Handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
