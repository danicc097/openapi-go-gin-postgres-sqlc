package rest

import (
	"fmt"
	"time"

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
	event          *Event
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
	event := newSSEServer()

	// we can have as many of these but need to delay call
	go func() {
		for {
			// We are streaming current time to clients in the interval 10 seconds
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			event.Message <- currentTime
		}
	}()

	return &Handlers{
		logger:         logger,
		pool:           pool,
		movieSvcClient: movieSvcClient,
		usvc:           usvc,
		authzsvc:       authzsvc,
		authnsvc:       authnsvc,
		authmw:         authmw,
		event:          event,
	}
}

// middlewares to be applied after authMiddlewares.
func (h *Handlers) middlewares(opID OperationID) []gin.HandlerFunc {
	switch opID {
	case Events:
		return []gin.HandlerFunc{SSEHeadersMiddleware(), h.event.serveHTTP()}
	default:
		return []gin.HandlerFunc{}
	}
}
