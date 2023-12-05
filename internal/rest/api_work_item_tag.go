package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) UpdateWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) CreateWorkItemTag(c *gin.Context, projectName models.ProjectName) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) GetWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) DeleteWorkItemTag(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}
