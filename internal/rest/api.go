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

// Handlers implements ServerInterface.
type Handlers struct {
	usvc            *services.User
	demoworkitemsvc *services.DemoWorkItem
	workitemtagsvc  *services.WorkItemTag
	logger          *zap.SugaredLogger
	pool            *pgxpool.Pool
	movieSvcClient  v1.MovieGenreClient
	authmw          *authMiddleware
	authzsvc        *services.Authorization
	authnsvc        *services.Authentication
	event           *Event
	provider        rp.RelyingParty
}

// NewHandlers returns an server implementation of an openapi specification.
func NewHandlers(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	movieSvcClient v1.MovieGenreClient,
	usvc *services.User,
	demoworkitemsvc *services.DemoWorkItem,
	workitemtagsvc *services.WorkItemTag,
	authzsvc *services.Authorization,
	authnsvc *services.Authentication,
	authmw *authMiddleware,
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
		logger:          logger,
		pool:            pool,
		movieSvcClient:  movieSvcClient,
		usvc:            usvc,
		authzsvc:        authzsvc,
		authnsvc:        authnsvc,
		demoworkitemsvc: demoworkitemsvc,
		workitemtagsvc:  workitemtagsvc,
		authmw:          authmw,
		event:           event,
		provider:        provider,
	}
}

// middlewares to be applied after authMiddlewares.
func (h *Handlers) middlewares(opID OperationID) []gin.HandlerFunc {
	switch opID {
	case Events:
		return []gin.HandlerFunc{
			SSEHeadersMiddleware(),
			h.event.serveHTTP(),
		}
	case MyProviderCallback:
		return []gin.HandlerFunc{
			h.codeExchange(),
		}
	case MyProviderLogin:
		// redirects the user to the auth server including state handling with secure cookie and the possibility to use PKCE
		return []gin.HandlerFunc{
			gin.WrapH(rp.AuthURLHandler(state, h.provider)),
		}
	default:
		return []gin.HandlerFunc{}
	}
}
