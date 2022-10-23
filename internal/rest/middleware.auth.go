package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// authMiddleware handles authentication and authorization middleware.
type authMiddleware struct {
	logger         *zap.Logger
	pool           *pgxpool.Pool
	movieSvcClient v1.MovieGenreClient
}

func newAuthMiddleware(
	logger *zap.Logger,
	pool *pgxpool.Pool,
	movieSvcClient v1.MovieGenreClient,
) *authMiddleware {
	return &authMiddleware{
		logger:         logger,
		pool:           pool,
		movieSvcClient: movieSvcClient,
	}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check app-specific jwt or api_key
// else redirect to /auth/{provider}/login (no auth middleware here or in */callback).
func (t *authMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.logger.Sugar().Info("Would have run EnsureAuthenticated and set user in ctx")
		authsvc := services.NewAuthentication(t.pool, t.logger)
		// if x-api-key header found
		authsvc.GetUserFromApiKey(c.Request.Context())
		// if auth header with bearer scheme found
		authsvc.GetUserFromToken(c.Request.Context())

		// set user to context
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO use authorization service, which in turn uses the user service to check role
// based on token -> email -> GetUserByEmail -> role
func (t *authMiddleware) EnsureAuthorized(requiredRole db.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authzSvc := services.NewAuthorization(t.logger)
		user := getUserFromCtx(c)
		if user == nil {
			renderErrorResponse(c, "Could not get user from context.", nil)

			return
		}
		err := authzSvc.IsAuthorized(user.Role, requiredRole)
		if err != nil {
			renderErrorResponse(c, "Unauthorized.", err)

			return
		}
	}
}

func verifyAuthentication(c context.Context, input *openapi3filter.AuthenticationInput) error {
	// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#securityRequirementObject
	switch input.SecurityScheme.Type {
	case "apiKey":
		if input.SecurityScheme.In != "header" {
			return fmt.Errorf("api key authentication only supported in header")
		}

		_, found := input.RequestValidationInput.Request.Header[http.CanonicalHeaderKey(input.SecurityScheme.Name)]

		if !found {
			return fmt.Errorf("%v not found in header", input.SecurityScheme.Name)
		}
	case "http":
		if input.SecurityScheme.Scheme != "bearer" {
			return fmt.Errorf("http security scheme only supports 'bearer' scheme")
		}

		authHeader, found := input.RequestValidationInput.Request.Header[http.CanonicalHeaderKey("Authorization")]
		if !found {
			return fmt.Errorf("authorization header missing")
		}

		if !strings.HasPrefix(authHeader[0], "Bearer ") {
			return fmt.Errorf("mismatching scheme in %s - expected Bearer", authHeader[0])
		}
	}

	return nil
}
