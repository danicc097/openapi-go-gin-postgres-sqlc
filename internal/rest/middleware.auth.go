package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// authMiddleware handles authentication and authorization middleware.
type authMiddleware struct {
	Logger   *zap.Logger
	authnSvc AuthenticationService
	authzSvc AuthorizationService
	userSvc  UserService
}

func newAuthMiddleware(
	logger *zap.Logger,
	authnSvc AuthenticationService,
	authzSvc AuthorizationService,
	userSvc UserService,
) *authMiddleware {
	return &authMiddleware{
		Logger:   logger,
		authnSvc: authnSvc,
		authzSvc: authzSvc,
		userSvc:  userSvc,
	}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check app-specific jwt or api_key
func (t *authMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO use authorization service, which in turn uses the user service to check role
// based on token -> email -> GetUserByEmail -> role
func (t *authMiddleware) EnsureAuthorized(requiredRole db.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := getUserFromCtx(c)
		if user == nil {
			renderErrorResponse(c, "Could not get user from context.", nil)

			return
		}
		err := t.authzSvc.IsAuthorized(user.Role, requiredRole)
		if err != nil {
			renderErrorResponse(c, "Unauthorized.", err)

			return
		}
	}
}

// EnsureVerified checks whether the client is verified.
func (t *authMiddleware) EnsureVerified() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
		// u := userSvc.getUserByToken...
		// ... u.isVerified
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
