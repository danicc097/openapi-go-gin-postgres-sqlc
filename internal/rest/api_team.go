package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateTeam(c *gin.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	ctx := c.Request.Context()
	tx := GetTxFromCtx(c)

	params := request.Body.TeamCreateParams
	params.ProjectID = internal.ProjectIDByName[request.ProjectName]

	team, err := h.svc.Team.Create(ctx, tx, &params)
	if err != nil {
		renderErrorResponse(c, "Could not create team", err)

		return nil, nil
	}

	return CreateTeam201JSONResponse{Team: *team}, nil
}

func (h *StrictHandlers) UpdateTeam(c *gin.Context, request UpdateTeamRequestObject) (UpdateTeamResponseObject, error) {
	ctx := c.Request.Context()
	tx := GetTxFromCtx(c)

	params := request.Body.TeamUpdateParams

	team, err := h.svc.Team.Update(ctx, tx, db.TeamID(request.Id), &params)
	if err != nil {
		renderErrorResponse(c, "Could not update team", err)

		return nil, nil
	}

	return UpdateTeam200JSONResponse{Team: *team}, nil
}

func (h *StrictHandlers) GetTeam(c *gin.Context, request GetTeamRequestObject) (GetTeamResponseObject, error) {
	ctx := c.Request.Context()
	tx := GetTxFromCtx(c)

	team, err := h.svc.Team.ByID(ctx, tx, db.TeamID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not get team", err)

		return nil, nil
	}

	return GetTeam200JSONResponse{Team: *team}, nil
}

func (h *StrictHandlers) DeleteTeam(c *gin.Context, request DeleteTeamRequestObject) (DeleteTeamResponseObject, error) {
	ctx := c.Request.Context()
	tx := GetTxFromCtx(c)

	_, err := h.svc.Team.Delete(ctx, tx, db.TeamID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete team", err)

		return nil, nil
	}

	return DeleteTeam204Response{}, nil
}
