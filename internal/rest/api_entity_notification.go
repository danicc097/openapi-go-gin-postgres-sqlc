package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateEntityNotification(c *gin.Context, request CreateEntityNotificationRequestObject) (CreateEntityNotificationResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteEntityNotification(c *gin.Context, request DeleteEntityNotificationRequestObject) (DeleteEntityNotificationResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetEntityNotification(c *gin.Context, request GetEntityNotificationRequestObject) (GetEntityNotificationResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateEntityNotification(c *gin.Context, request UpdateEntityNotificationRequestObject) (UpdateEntityNotificationResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
