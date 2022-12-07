package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeProject
func (h *Handlers) InitializeProject(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "not implemented")
}

// GetProjectBoard
func (h *Handlers) GetProjectBoard(c *gin.Context, id int) {
	c.String(http.StatusNotImplemented, "not implemented")
}
