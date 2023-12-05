package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateTeam(c *gin.Context, projectName models.ProjectName) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) DeleteTeam(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) GetTeam(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) UpdateTeam(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}
