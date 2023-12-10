package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminPing ping pongs.
func (h *dummyStrictHandlers) AdminPing(c *gin.Context, request AdminPingRequestObject) (AdminPingResponseObject, error) {
	return AdminPing200TextResponse("pong"), nil
}

func (h *StrictHandlers) AdminPing(c *gin.Context, request AdminPingRequestObject) (AdminPingResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
