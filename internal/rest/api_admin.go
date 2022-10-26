package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminPing ping pongs.
func (h *Handlers) AdminPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
