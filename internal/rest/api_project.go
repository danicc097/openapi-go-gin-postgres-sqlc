package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateProjectConfig(c *gin.Context, request UpdateProjectConfigRequestObject) (UpdateProjectConfigResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetProject(c *gin.Context, request GetProjectRequestObject) (GetProjectResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetProjectBoard(c *gin.Context, request GetProjectBoardRequestObject) (GetProjectBoardResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetProjectWorkitems(c *gin.Context, request GetProjectWorkitemsRequestObject) (GetProjectWorkitemsResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetProjectConfig(c *gin.Context, request GetProjectConfigRequestObject) (GetProjectConfigResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) InitializeProject(c *gin.Context, request InitializeProjectRequestObject) (InitializeProjectResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
