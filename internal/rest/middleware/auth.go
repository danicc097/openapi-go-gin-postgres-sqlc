package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Auth handles authentication and authorization middleware.
type Auth struct {
	Logger *zap.Logger
}

func NewAuth(logger *zap.Logger) *Auth {
	return &Auth{Logger: logger}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check jwt.
func (t *Auth) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO grab user role.
func (t *Auth) EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Logger.Sugar().Info("Would have run EnsureAuthorized")
	}
}
