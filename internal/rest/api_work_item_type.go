package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *dummyStrictHandlers) DeleteWorkItemType(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *dummyStrictHandlers) GetWorkItemType(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *dummyStrictHandlers) CreateWorkItemType(c *gin.Context, projectName models.ProjectName) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *dummyStrictHandlers) UpdateWorkItemType(c *gin.Context, projectName models.ProjectName, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) GetWorkItemType(c *gin.Context, request GetWorkItemTypeRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) CreateWorkItemType(c *gin.Context, request CreateWorkItemTypeRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) UpdateWorkItemType(c *gin.Context, request UpdateWorkItemTypeRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *StrictHandlers) DeleteWorkItemType(c *gin.Context, request DeleteWorkItemTypeRequestObject) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}
