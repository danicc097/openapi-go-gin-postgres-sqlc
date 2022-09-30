package rest

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
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
// TODO check jwt.
func (t *authMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO use authorization service, which in turn uses the user service to check role
// based on token -> email -> GetUserByEmail
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
