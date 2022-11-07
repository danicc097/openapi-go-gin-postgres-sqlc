package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// authMiddleware handles authentication and authorization middleware.
type authMiddleware struct {
	logger   *zap.Logger
	pool     *pgxpool.Pool
	authnsvc *services.Authentication
	authzsvc *services.Authorization
	usersvc  *services.User
}

func newAuthMiddleware(
	logger *zap.Logger,
	pool *pgxpool.Pool,
	authnsvc *services.Authentication,
	authzsvc *services.Authorization,
	usersvc *services.User,
) *authMiddleware {
	return &authMiddleware{
		logger:   logger,
		pool:     pool,
		authnsvc: authnsvc,
		authzsvc: authzsvc,
		usersvc:  usersvc,
	}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check app-specific jwt or api_key
// else redirect to /auth/{provider}/login (no auth middleware here or in */callback).
func (a *authMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		a.logger.Sugar().Info("Would have run EnsureAuthenticated and set user in ctx")
		// if x-api-key header found
		a.authnsvc.GetUserFromApiKey(c.Request.Context())
		// if auth header with bearer scheme found
		a.authnsvc.GetUserFromToken(c.Request.Context())

		// set user to context
		// ctxWithUser(c)
	}
}

// TODO EnsureAuthorizedRole and EnsureAuthorizedScopes(scopes ...Scopes)
// 1. x-required-scopes read by yq in spec
// 2. generate a JSON file for frontend and backend to use: {<operationID>: [<...scopes>], ...}.
// 3.  new method authMiddleware.EnsureAuthorizedScopes(opID operationID, user *db.User), which
// 4. uses the loaded JSON to check if operationIDScopes[opID] exists, in which case
// checks if user.scopes contains the required scopes as per spec
// it belongs here, not in a service since this is specific to rest.
type operationIDScopes = map[operationID][]string

type AuthRestriction struct {
	MinimumRole    models.Role
	RequiredScopes []models.Scope
}

// EnsureAuthorized checks whether the client is authorized.
func (a *authMiddleware) EnsureAuthorized(config AuthRestriction) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := getUserFromCtx(c)
		if user == nil {
			renderErrorResponse(c, "Could not get user from context.", nil)

			return
		}

		if err := a.authzsvc.HasRequiredRole(user.RoleRank, config.MinimumRole); err != nil {
			renderErrorResponse(c, "Unauthorized.", err)

			return
		}

		if err := a.authzsvc.HasRequiredScopes(user.Scopes, config.RequiredScopes); err != nil {
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
