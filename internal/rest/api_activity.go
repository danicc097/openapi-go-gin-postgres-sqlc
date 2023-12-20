package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateActivity(c *gin.Context, request CreateActivityRequestObject) (CreateActivityResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteActivity(c *gin.Context, request DeleteActivityRequestObject) (DeleteActivityResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetActivity(c *gin.Context, request GetActivityRequestObject) (GetActivityResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateActivity(c *gin.Context, request UpdateActivityRequestObject) (UpdateActivityResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
