package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) GetTimeEntry(c *gin.Context, request GetTimeEntryRequestObject) (GetTimeEntryResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateTimeEntry(c *gin.Context, request UpdateTimeEntryRequestObject) (UpdateTimeEntryResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) CreateTimeEntry(c *gin.Context, request CreateTimeEntryRequestObject) (CreateTimeEntryResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteTimeEntry(c *gin.Context, request DeleteTimeEntryRequestObject) (DeleteTimeEntryResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
