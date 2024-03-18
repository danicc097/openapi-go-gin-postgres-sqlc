package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// authMiddleware handles authentication and authorization middleware.
type authMiddleware struct {
	logger *zap.SugaredLogger
	pool   *pgxpool.Pool
	svc    *services.Services
}

func NewAuthMiddleware(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	svcs *services.Services,
) *authMiddleware {
	return &authMiddleware{
		logger: logger,
		pool:   pool,
		svc:    svcs,
	}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check app-specific jwt or api_key
// else redirect to /auth/{provider}/login (no auth middleware here or in */callback).
func (m *authMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get(ApiKeyHeaderKey)
		auth := c.Request.Header.Get(AuthorizationHeaderKey)
		if apiKey != "" {
			u, err := m.svc.Authentication.GetUserFromAPIKey(c.Request.Context(), apiKey) // includes caller joins
			if err != nil || u == nil {
				renderErrorResponse(c, "Unauthenticated", internal.NewErrorf(models.ErrorCodeUnauthenticated, "could not get user from api key"))
				c.Abort()

				return
			}

			CtxWithUserCaller(c, u)

			c.Next() // executes the pending handlers. What goes below is cleanup after the complete request.

			return
		}
		if strings.HasPrefix(auth, "Bearer ") {
			u, err := m.svc.Authentication.GetUserFromAccessToken(c.Request.Context(), strings.Split(auth, "Bearer ")[1]) // includes caller joins
			if err != nil || u == nil {
				renderErrorResponse(c, "Unauthenticated", internal.NewErrorf(models.ErrorCodeUnauthenticated, "could not get user from token: %s", err))
				c.Abort()

				return
			}

			CtxWithUserCaller(c, u)

			c.Next() // executes the pending handlers. What goes below is cleanup after the complete request.

			return
		}

		renderErrorResponse(c, "Unauthenticated", internal.NewErrorf(models.ErrorCodeUnauthenticated, "could not get user from token"))
		c.Abort()
	}
}

type AuthRestriction struct {
	MinimumRole    models.Role
	RequiredScopes models.Scopes
}

// EnsureAuthorized checks whether the client is authorized with either a
// minimum role or has all required scopes.
func (m *authMiddleware) EnsureAuthorized(config AuthRestriction) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMsg := ""
		errs := []string{}
		user, err := GetUserCallerFromCtx(c)
		if err != nil {
			renderErrorResponse(c, "Could not get current user.", nil)
			c.Abort()

			return
		}

		if config.MinimumRole != "" {
			userRole, ok := m.svc.Authorization.RoleByRank(user.RoleRank)
			if !ok {
				renderErrorResponse(c, fmt.Sprintf("Unknown rank value: %d", user.RoleRank), errors.New("unknown rank"))
				c.Abort()

				return
			}

			if err := m.svc.Authorization.HasRequiredRole(userRole, config.MinimumRole); err == nil {
				c.Next()

				return
			}
			errs = append(errs, fmt.Sprintf("role %s", config.MinimumRole))
		}

		if len(config.RequiredScopes) > 0 {
			if err := m.svc.Authorization.HasRequiredScopes(user.Scopes, config.RequiredScopes); err == nil {
				c.Next()

				return
			}
			errs = append(errs, fmt.Sprintf("scope(s) %s", slices.JoinWithAnd(config.RequiredScopes)))
		}

		if len(errs) == 1 {
			errorMsg = errs[0] + " is required"
		} else if len(errs) > 1 {
			errorMsg = fmt.Sprintf("either %s are required", slices.Join(errs, " or "))
		}

		renderErrorResponse(c, "Unauthorized", internal.NewErrorf(models.ErrorCodeUnauthorized, "unauthorized: "+errorMsg))
		c.Abort()
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

		authHeader, found := input.RequestValidationInput.Request.Header[http.CanonicalHeaderKey(AuthorizationHeaderKey)]
		if !found {
			return fmt.Errorf("authorization header missing")
		}

		if !strings.HasPrefix(authHeader[0], "Bearer ") {
			return fmt.Errorf("mismatching scheme in %s - expected Bearer", authHeader[0])
		}
	}

	return nil
}
