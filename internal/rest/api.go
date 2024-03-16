package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"go.uber.org/zap"
)

const (
	ApiKeyHeaderKey        = "x-api-key"
	AuthorizationHeaderKey = "authorization"
)

type StrictHandlers struct {
	svc *services.Services

	logger         *zap.SugaredLogger
	pool           *pgxpool.Pool
	moviesvcclient v1.MovieGenreClient
	specPath       string
	authmw         *authMiddleware
	event          *EventServer
	provider       rp.RelyingParty
}

// IMPORTANT: oapi codegen uses its own types for responses and request bodies,
// and we absolutely do not want this, since we would need to manually convert
// to or from oapi's Rest<..> and Db<..> structs.
// is it worth it to rewrite strict server gen to use rest package structs for request/response bodies?
// we could check in templates with a simple if stmt
// that if a type in rest package exists with the same name we don't prepend `externalRef0.`
// We already have the rest pkg struct list from ast-parser gen.
var _ StrictServerInterface = (*StrictHandlers)(nil)

// NewStrictHandlers returns a server implementation of an openapi specification.
func NewStrictHandlers(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	moviesvcclient v1.MovieGenreClient,
	specPath string,
	svcs *services.Services,
	authmw *authMiddleware, // middleware needed here since it's generated code
	provider rp.RelyingParty,
) StrictServerInterface {
	event := newSSEServer()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)
			event.Publish(currentTime, TopicsGlobalAlerts)

		}
	}()

	go func() {
		for {
			data := struct {
				Field1 string `json:"field1"`
				Field2 int    `json:"field2"`
			}{
				Field1: "value1",
				Field2: 42,
			}

			msgData, err := json.Marshal(data)
			if err != nil {
				log.Printf("Error marshaling JSON: %v", err)
				continue
			}

			event.Publish(string(msgData), TopicsGlobalAlerts)

			time.Sleep(time.Second * 1)
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

func skipRequestValidationMw(c *gin.Context) { c.Set(skipRequestValidationCtxKey, true) }

func skipResponseValidationMw(c *gin.Context) { c.Set(skipResponseValidationCtxKey, true) }

// middlewares to be applied after authMiddlewares, based on operation IDs.
func (h *StrictHandlers) middlewares(opID OperationID) []gin.HandlerFunc {
	dbMw := newDBMiddleware(h.logger, h.pool)
	tracingMw := newTracingMiddleware()

	defaultMws := []gin.HandlerFunc{tracingMw.RequestIDMiddleware("my-app")}

	ignoredOperationID := opID == MyProviderLogin || opID == Events

	if !ignoredOperationID {
		defaultMws = append(defaultMws, tracingMw.WithSpan(), dbMw.BeginTransaction())
	}

	switch opID {
	case Events:
		// TODO: last mw should be event dispatcher middleware, that will dispatch pending ones
		// if renderErrorResponse was not called, ie !CtxHasErrorResponse()
		return append(
			defaultMws,
			skipRequestValidationMw,
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
