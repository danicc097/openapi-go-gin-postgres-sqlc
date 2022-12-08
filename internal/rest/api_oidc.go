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
