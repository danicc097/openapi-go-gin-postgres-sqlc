package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zitadel/oidc/pkg/client/rp"
	"github.com/zitadel/oidc/pkg/oidc"
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
	provider       rp.RelyingParty
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
	provider rp.RelyingParty,
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
		provider:       provider,
	}
}

// middlewares to be applied after authMiddlewares.
func (h *Handlers) middlewares(opID OperationID) []gin.HandlerFunc {
	state := func() string {
		return uuid.New().String()
	}

	marshalToken := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty) {
		data, err := json.Marshal(tokens)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}

	// TODO /auth/logout

	switch opID {
	case Events:
		return []gin.HandlerFunc{SSEHeadersMiddleware(), h.event.serveHTTP()}
	case MyProviderCallback:
		// return []gin.HandlerFunc{gin.WrapH(rp.CodeExchangeHandler(rp.UserinfoCallback(marshalUserinfo), h.provider))}
		// without calling userinfo
		// register the CodeExchangeHandler at the conf.MyProviderCallbackPath
		// the CodeExchangeHandler handles the auth response, creates the token request and calls the callback function
		// with the returned tokens from the token endpoint
		// in this example the callback function itself is wrapped by the UserinfoCallback which
		// will call the Userinfo endpoint, check the sub and pass the info into the callback function
		// TODO in reality we would redirect back to our app frontend instead
		// this callback must also be extended to create an access token (or any other login method) for our own
		// app (and create user if not found) once the external auth request is authorized
		return []gin.HandlerFunc{gin.WrapH(rp.CodeExchangeHandler(marshalToken, h.provider))}
	case MyProviderLogin:
		// the AuthURLHandler creates the auth request and redirects the user to the auth server
		// including state handling with secure cookie and the possibility to use PKCE
		return []gin.HandlerFunc{gin.WrapH(rp.AuthURLHandler(state, h.provider))}
	default:
		return []gin.HandlerFunc{}
	}
}
