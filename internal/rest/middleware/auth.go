package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Auth handles authentication and authorization middleware.
type Auth struct {
	conf *AuthConf
}

// AuthConf represents the required configuration for auth middleware.
type AuthConf struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}

func NewAuth(conf *AuthConf) *Auth {
	return &Auth{conf: conf}
}

// EnsureAuthenticated checks whether the client is authenticated.
// TODO check jwt.
func (t *Auth) EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.conf.Logger.Sugar().Info("Would have run EnsureAuthenticated")
	}
}

// EnsureAuthorized checks whether the client is authorized.
// TODO use authorization service, which in turn uses the user service to check role
// based on token -> email -> GetUserByEmail
func (t *Auth) EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.conf.Logger.Sugar().Info("Would have run EnsureAuthorized")
	}
}
