package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/pkg/client/rp"
	"github.com/zitadel/oidc/pkg/oidc"
)

func (h *Handlers) MyProviderLogin(c *gin.Context) {
	c.Set(skipRequestValidation, true)

	// TODO if env is dev should have helper func or just use zitadel oidcserver but with dummy data?
	// real server, though fake, will be hard to keep in sync. much easier to use a map of tokens that get returned
	// and key is set in .env.dev -> "DEV_USER": <map key> so that backend route on login gets that inmemory token,
	// and sets to cookie instead of redirecting to auth server and then doing all that in MyProviderCallback
	// (user already exists, all users got created via project db.initial-data )
	// if the above is done, we still need to test provider callback logic anyway (we can use mocked returned tokens in the same style as the above ones)
	// for testing, the app_env switch is not what will happen in prod. should int tests run in app_env=ci or app_env=prod?
	// to test auth server calls: https://deliveroo.engineering/2018/09/11/testing-with-third-parties-in-go.html
}

func (h *Handlers) MyProviderCallback(c *gin.Context) {
	c.Set(skipRequestValidation, true)
}

func state() string {
	return uuid.New().String()
}

// TODO we would create our own authentication here (another jwt), create user if not exists and then redirect back to our app frontend
func (h *Handlers) marshalToken(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty) {
	data, err := json.Marshal(tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
