package rest

import (
	"context"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

type authenticatedCtxKey struct{}

// TODO move elsewhere
func GetUser(ctx *gin.Context) *db.Users {
	user, ok := ctx.Value(string(userCtxKey)).(*db.Users)
	if !ok {
		// Log this issue
		return nil
	}

	return user
}

func isAuthenticated(ctx context.Context) bool {
	authenticated, ok := ctx.Value(authenticatedCtxKey{}).(bool)
	if !ok {
		// Log this issue
		return false
	}
	return authenticated
}

// Auth handles authentication and authorization middleware.
type Auth struct {
	Logger   *zap.Logger
	authnSvc AuthenticationService
	authzSvc AuthorizationService
	userSvc  UserService
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
		user := GetUser(c)
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
func (t *Auth) EnsureVerified() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
		// u := userSvc.getUserByToken...
		// ... u.isVerified
	}
}
