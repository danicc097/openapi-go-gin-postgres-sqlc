package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthMiddleware handles authentication and authorization middleware.
type AuthMiddleware struct {
	Logger *zap.Logger
}

func NewAuthMiddleware(logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{Logger: logger}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check jwt.
func (t *AuthMiddleware) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO grab user role.
func (t *AuthMiddleware) EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
	}
}
