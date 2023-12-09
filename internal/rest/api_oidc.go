package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

func (h *StrictHandlers) MyProviderLogin(c *gin.Context) {
	c.Set(skipRequestValidationCtxKey, true)

	gin.WrapH(rp.AuthURLHandler(state, h.provider))(c)

	// use adaptation of https://github.com/zitadel/oidc/blob/main/example/client/app/app.go

	// X TODO if env is dev should have helper func or just use zitadel oidcserver but with dummy data?
	// X real server, though fake, will be hard to keep in sync. much easier to use a map of tokens that get returned
	// X and key is set in .env.dev -> "DEV_USER": <map key> so that backend route on login gets that inmemory token,
	// X and sets to cookie instead of redirecting to auth server and then doing all that in MyProviderCallback
	// X (user already exists, all users got created via project db.initial-data )
	// X if the above is done, we still need to test provider callback logic anyway (we can use mocked returned tokens in the same style as the above ones)
	// X for testing, the app_env switch is not what will happen in prod. should int tests run in app_env=ci or app_env=prod?
	// X to test auth server calls: https://deliveroo.engineering/2018/09/11/testing-with-third-parties-in-go.html

	// IMPORTANT: easiest if import.meta is dev then use headers.set('x-api-key', `...`) for UI.
	// that needs to have remove Authorization header removed. (or could fallthrough if auth header check failed so both can be used at the same time)
	// its the same we do to test out swagger ui.
	// initial-data for dev can create api keys for every user.
}

func (h *StrictHandlers) MyProviderCallback(c *gin.Context) {
	c.Set(skipRequestValidationCtxKey, true)

	userinfo := getUserInfoFromCtx(c)
	if userinfo == nil {
		renderErrorResponse(c, "OIDC authentication error", internal.NewErrorf(models.ErrorCodeOIDC, "could not get OIDC userinfo from context"))
	}
	fmt.Printf("userinfo in MyProviderCallback: %v\n", string(userinfo))
	// {"email":"admin@admin.com","email_verified":true,"family_name":"Admin","given_name":"Mr","locale":"de","name":"Mr Admin","preferred_username":"admin","sub":"id1"}

	// GetOrRegisterUserFromUserInfo
	// CreateAccessTokenForUser
	// get "auth:redirect-uri" cookie

	c.String(200, "would have been redirected to app frontend with actual user and logged in with JWT")
}

func state() string {
	return uuid.New().String()
}

func (h *StrictHandlers) marshalUserinfo(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info *oidc.UserInfo) {
	data, err := json.Marshal(info)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not marshal userinfo: %s", err.Error()), http.StatusInternalServerError)

		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not write userinfo: %s", err.Error()), http.StatusInternalServerError)

		return
	}
}

func (h *StrictHandlers) codeExchange() gin.HandlerFunc {
	return func(c *gin.Context) {
		rbw := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = rbw
		rp.CodeExchangeHandler(rp.UserinfoCallback(h.marshalUserinfo), h.provider).ServeHTTP(rbw, c.Request)
		ctxWithUserInfo(c, rbw.body.Bytes())
		rbw.body = &bytes.Buffer{}
		c.Next()
		rbw.ResponseWriter.Write(rbw.body.Bytes())
	}
}
