package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminPing ping pongs.
func (h *dummyStrictHandlers) AdminPing(c *gin.Context, request AdminPingRequestObject) (AdminPingResponseObject, error) {
	return AdminPing200TextResponse("pong"), nil
}
