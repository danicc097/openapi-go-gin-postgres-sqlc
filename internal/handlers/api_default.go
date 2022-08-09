package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ping pongs.
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
