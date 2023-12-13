package rest

import (
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) AdminPing(c *gin.Context, request AdminPingRequestObject) (AdminPingResponseObject, error) {
	return AdminPing200TextResponse("pong"), nil
}
