package rest

import (
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// Handlers implements ServerInterface.
type Handlers struct {
	usvc           *services.User
	logger         *zap.Logger
	pool           *pgxpool.Pool
	movieSvcClient v1.MovieGenreClient
	authmw         *authMiddleware
	authzsvc       *services.Authorization
	authnsvc       *services.Authentication
}

// NewHandlers returns an server implementation of an openapi specification.
func NewHandlers(
	logger *zap.Logger,
	pool *pgxpool.Pool,
	movieSvcClient v1.MovieGenreClient,
	usvc *services.User,
	authzsvc *services.Authorization,
	authnsvc *services.Authentication,
	authmw *authMiddleware,
) *Handlers {
	return &Handlers{
		logger:         logger,
		pool:           pool,
		movieSvcClient: movieSvcClient,
		usvc:           usvc,
		authzsvc:       authzsvc,
		authnsvc:       authnsvc,
		authmw:         authmw,
	}
}

func (h *Handlers) middlewares(opID OperationID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}
