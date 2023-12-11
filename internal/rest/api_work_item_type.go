package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkItemType(c *gin.Context, request CreateWorkItemTypeRequestObject) (CreateWorkItemTypeResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteWorkItemType(c *gin.Context, request DeleteWorkItemTypeRequestObject) (DeleteWorkItemTypeResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateWorkItemType(c *gin.Context, request UpdateWorkItemTypeRequestObject) (UpdateWorkItemTypeResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetWorkItemType(c *gin.Context, request GetWorkItemTypeRequestObject) (GetWorkItemTypeResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
