package rest

import (
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateActivity(c *gin.Context, request CreateActivityRequestObject) (CreateActivityResponseObject, error) {
	activity, err := h.svc.Activity.Create(c.Request.Context(), h.pool, request.ProjectName, &request.Body.ActivityCreateParams)
	if err != nil {
		renderErrorResponse(c, "could not create activity", err)
	}

	return CreateActivity201JSONResponse{Activity: *activity}, nil
}

func (h *StrictHandlers) DeleteActivity(c *gin.Context, request DeleteActivityRequestObject) (DeleteActivityResponseObject, error) {
	if _, err := h.svc.Activity.Delete(c.Request.Context(), h.pool, request.ActivityID); err != nil {
		renderErrorResponse(c, "could not delete activity", err)
	}

	return DeleteActivity204Response{}, nil
}

func (h *StrictHandlers) GetActivity(c *gin.Context, request GetActivityRequestObject) (GetActivityResponseObject, error) {
	activity, err := h.svc.Activity.ByID(c.Request.Context(), h.pool, request.ActivityID)
	if err != nil {
		renderErrorResponse(c, "could not get activity", err)
	}

	return GetActivity200JSONResponse{Activity: *activity}, nil
}

func (h *StrictHandlers) UpdateActivity(c *gin.Context, request UpdateActivityRequestObject) (UpdateActivityResponseObject, error) {
	activity, err := h.svc.Activity.Update(c.Request.Context(), h.pool, request.ActivityID, &request.Body.ActivityUpdateParams)
	if err != nil {
		renderErrorResponse(c, "could not update activity", err)
	}

	return UpdateActivity200JSONResponse{Activity: *activity}, nil
}
