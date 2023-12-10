package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateTeam(c *gin.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateTeam(c *gin.Context, request UpdateTeamRequestObject) (UpdateTeamResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) GetTeam(c *gin.Context, request GetTeamRequestObject) (GetTeamResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteTeam(c *gin.Context, request DeleteTeamRequestObject) (DeleteTeamResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
