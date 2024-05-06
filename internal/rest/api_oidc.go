package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

const authRedirectCookieKey = "auth_redirect_uri"

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
		CtxWithUserInfo(c, rbw.body.Bytes())
		rbw.body = &bytes.Buffer{}
		c.Next()
		rbw.ResponseWriter.Write(rbw.body.Bytes())
	}
}

func (h *StrictHandlers) MyProviderCallback(c *gin.Context, request MyProviderCallbackRequestObject) (MyProviderCallbackResponseObject, error) {
	c.Set(skipRequestValidationCtxKey, true)

	userinfo, err := GetUserInfoFromCtx(c)
	if err != nil {
		renderErrorResponse(c, "OIDC authentication error", internal.WrapErrorf(err, models.ErrorCodeOIDC, "user info not found"))
		return nil, nil
	}

	ctx := c.Request.Context()

	u, err := h.svc.Authentication.GetOrRegisterUserFromUserInfo(ctx, *userinfo)
	if err != nil {
		renderErrorResponse(c, "OIDC authentication error", internal.WrapErrorf(err, models.ErrorCodeOIDC, "could not get or register user"))
		return nil, nil
	}

	accessToken, err := h.svc.Authentication.CreateAccessTokenForUser(ctx, u)
	if err != nil {
		renderErrorResponse(c, "OIDC authentication error", internal.WrapErrorf(err, models.ErrorCodeOIDC, "could not create access token"))
		return nil, nil
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     internal.Config.LoginCookieKey,
		Value:    accessToken,
		Path:     "/",
		MaxAge:   3600 * 24 * 7,
		Domain:   internal.Config.CookieDomain,
		Secure:   true,
		HttpOnly: false, // must access via JS
		SameSite: http.SameSiteNoneMode,
	})

	redirectURI, err := c.Cookie(authRedirectCookieKey)
	if err != nil {
		redirectURI = internal.BuildAPIURL("docs")
	}

	return MyProviderCallback302Response{
		Headers: MyProviderCallback302ResponseHeaders{
			Location: redirectURI,
		},
	}, nil
}

func (h *StrictHandlers) MyProviderLogin(c *gin.Context, request MyProviderLoginRequestObject) (MyProviderLoginResponseObject, error) {
	c.Set(skipRequestValidationCtxKey, true)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     authRedirectCookieKey,
		Value:    request.Params.AuthRedirectUri,
		Path:     "/",
		MaxAge:   3600 * 24 * 7,
		Domain:   internal.Config.CookieDomain,
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})

	gin.WrapH(rp.AuthURLHandler(state, h.provider))(c)

	return nil, nil // redirect handled by middleware above
}
