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
	apiKeyHeaderKey        = "x-api-key"
	authorizationHeaderKey = "x-api-key"
)

type StrictHandlers struct {
	svc *services.Services

	logger         *zap.SugaredLogger
	pool           *pgxpool.Pool
	moviesvcclient v1.MovieGenreClient
	specPath       string
	authmw         *authMiddleware
	event          *Event
	provider       rp.RelyingParty
}

// IMPORTANT: oapi codegen uses its own types for responses and request bodies,
// and we absolutely do not want this, since we would need to manually convert
// to or from oapi's Rest<..> and Db<..> structs.
// is it worth it to rewrite strict server gen to use rest package structs for request/response bodies?
// we could check in templates with a simple if stmt
// that if a type in rest package exists with the same name we don't prepend `externalRef0.`
// We already have the rest pkg struct list from ast-parser gen
// var _ StrictServerInterface = (*StrictHandlers)(nil)

// NewStrictHandlers returns a server implementation of an openapi specification.
func NewStrictHandlers(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	moviesvcclient v1.MovieGenreClient,
	specPath string,
	svcs *services.Services,
	authmw *authMiddleware, // middleware needed here since it's generated code
	provider rp.RelyingParty,
) *StrictHandlers {
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

	return &StrictHandlers{
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
func (h *StrictHandlers) middlewares(opID OperationID) []gin.HandlerFunc {
	defaultMws := []gin.HandlerFunc{}

	dbMw := newDBMiddleware(h.logger, h.pool)
	tracingMw := newTracingMiddleware()

	ignoredOperationID := opID == MyProviderLogin

	if !ignoredOperationID {
		defaultMws = append(defaultMws, tracingMw.WithSpan(), dbMw.BeginTransaction())
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

	// TODO: last mw should be event dispatcher middleware, that will dispatch pending ones
	// if renderErrorResponse was not called, ie !ctxHasErrorResponse()
}
