package rest

import (
	"encoding/json"
	"log"
	"time"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
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

// StrictHandlers implements StrictServerInterface.
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

var _ StrictServerInterface = (*StrictHandlers)(nil)

// NewStrictHandlers returns a server implementation of an openapi specification.
func NewStrictHandlers(
	logger *zap.SugaredLogger, pool *pgxpool.Pool,
	event *EventServer,
	moviesvcclient v1.MovieGenreClient,
	specPath string,
	svcs *services.Services,
	authmw *authMiddleware, // middleware needed here since it's generated code
	provider rp.RelyingParty,
) StrictServerInterface {
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

			event.Publish(string(msgData), models.TopicAppDebug)

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

	// event cleanup will be the last to run
	defaultMws := []gin.HandlerFunc{h.event.EventDispatcher(), tracingMw.RequestIDMiddleware("my-app")}

	ignoredOperationID := opID == MyProviderLogin || opID == Events

	if !ignoredOperationID {
		defaultMws = append(defaultMws, tracingMw.WithSpan(), dbMw.BeginTransaction())
	}

	switch opID {
	case Events:
		// TODO: cors allow origin cannot be "*" since we use credentials to get user from cookie
		return append(
			defaultMws,
			skipResponseValidationMw, // stream not supported
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
