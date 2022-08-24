package handlers

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Auth handles authentication and authorization middleware.
type Auth struct {
	Logger   *zap.Logger
	authnSvc AuthenticationService
	authzSvc AuthorizationService
	userSvc  UserService
}

// AuthConf represents the required configuration for auth middleware.
type AuthConf struct {
}

func NewAuthMw(
	logger *zap.Logger,
	authnSvc AuthenticationService,
	authzSvc AuthorizationService,
	userSvc UserService,
) *Auth {
	return &Auth{
		Logger:   logger,
		authnSvc: authnSvc,
		authzSvc: authzSvc,
		userSvc:  userSvc,
	}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check jwt.
func (t *Auth) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO use authorization service, which in turn uses the user service to check role
// based on token -> email -> GetUserByEmail
func (t *Auth) EnsureAuthorized(requiredRole db.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
		// u := userSvc.getUserByToken...
		// t.authzSvc.IsAuthorized(u.Role, requiredRole)
	}
}

// EnsureVerified checks whether the client is verified.
func (t *Auth) EnsureVerified() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
		// u := userSvc.getUserByToken...
		// ... u.isVerified
	}
}
