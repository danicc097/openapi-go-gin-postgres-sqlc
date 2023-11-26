package rest

import (
	"fmt"
	"time"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"go.uber.org/zap"
)

const (
	apiKeyHeaderKey = "x-api-key"
)

// Handlers implements ServerInterface.
type Handlers struct {
	svc *services.Services

	logger         *zap.SugaredLogger
	pool           *pgxpool.Pool
	moviesvcclient v1.MovieGenreClient
	specPath       string
	authmw         *authMiddleware
	event          *Event
	provider       rp.RelyingParty
}

// NewHandlers returns a server implementation of an openapi specification.
func NewHandlers(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	moviesvcclient v1.MovieGenreClient,
	specPath string,
	svcs *services.Services,
	authmw *authMiddleware, // middleware needed here since it's generated code
	provider rp.RelyingParty,
) *Handlers {
	event := newSSEServer()

	// we can have as many of these but need to delay call
	go func() {
		for {
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			event.Message <- currentTime
			time.Sleep(time.Second * 2)
		}
	}()

	// we can have as many of these but need to delay call
	// we probably won't have an infinite running goroutine like this,
	// will send messages to channels on specific events.
	// but will be useful if we need to check something external
	// every X timeframe (e.g. wiki documents alert, new documents loaded for an active workitem, etc.)

	go func() {
		for {
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("user notifications - The Current Time Is %v", now)

			event.Message2 <- currentTime
			time.Sleep(time.Second * 2)
		}
	}()

	return &Handlers{
		logger:         logger,
		pool:           pool,
		moviesvcclient: moviesvcclient,
		svc:            svcs,
		authmw:         authmw,
		event:          event,
		provider:       provider,
		specPath:       specPath,
	}
}

// middlewares to be applied after authMiddlewares, based on operation IDs.
func (h *Handlers) middlewares(opID OperationID) []gin.HandlerFunc {
	defaultMws := []gin.HandlerFunc{}

	dbMw := newDBMiddleware(h.logger, h.pool)

	if opID != MyProviderLogin {
		defaultMws = append(defaultMws, dbMw.BeginTransaction())
	}

	switch opID {
	case Events:
		return append(
			defaultMws,
			SSEHeadersMiddleware(),
			h.event.serveHTTP(),
		)
	case MyProviderCallback:
		return append(
			defaultMws,
			h.codeExchange(),
		)
	default:
		return defaultMws
	}
}
