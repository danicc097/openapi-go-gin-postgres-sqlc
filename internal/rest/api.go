package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Handlers implements ServerInterface.
type Handlers struct {
	usvc           *services.User
	logger         *zap.Logger
	pool           *pgxpool.Pool
	movieSvcClient v1.MovieGenreClient
	authmw         *authMiddleware
}

// NewHandlers returns an server implementation of an openapi specification.
// TODO retry wrapper
func NewHandlers(
	logger *zap.Logger,
	pool *pgxpool.Pool,
	movieSvcClient v1.MovieGenreClient,
	usvc *services.User,
	authmw *authMiddleware,
) *Handlers {
	return &Handlers{
		logger:         logger,
		pool:           pool,
		movieSvcClient: movieSvcClient,
		usvc:           usvc,
		authmw:         authmw,
	}
}

func (h *Handlers) middlewares(opID operationID) []gin.HandlerFunc {
	switch opID {
	case DeleteUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(AuthRestriction{MinimumRole: models.RoleAdmin}),
		}
	case UpdateUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case UpdateUserAuthorization:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case AdminPing:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(AuthRestriction{MinimumRole: models.RoleAdmin}),
		}
	case GetCurrentUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	default:
		return []gin.HandlerFunc{}
	}
}
