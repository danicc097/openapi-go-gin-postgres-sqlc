package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
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

// GetProjectWorkitems
func (h *Handlers) GetProjectWorkitems(c *gin.Context, id int, params models.GetProjectWorkitemsParams) {
	c.String(http.StatusNotImplemented, "not implemented")
}
