package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkItemTag(c *gin.Context, request CreateWorkItemTagRequestObject) (CreateWorkItemTagResponseObject, error) {
	ctx := c.Request.Context()
	tx := GetTxFromCtx(c)
	caller, _ := GetUserCallerFromCtx(c)

	body := request.Body
	body.WorkItemTagCreateParams.ProjectID = internal.ProjectIDByName[request.ProjectName]

	wit, err := h.svc.WorkItemTag.Create(ctx, tx, caller, &body.WorkItemTagCreateParams)
	if err != nil {
		renderErrorResponse(c, "Could not create work item tag", err)

		return nil, nil
	}

	return CreateWorkItemTag201JSONResponse{WorkItemTag: *wit}, nil
}

func (h *StrictHandlers) GetWorkItemTag(c *gin.Context, request GetWorkItemTagRequestObject) (GetWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteWorkItemTag(c *gin.Context, request DeleteWorkItemTagRequestObject) (DeleteWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) UpdateWorkItemTag(c *gin.Context, request UpdateWorkItemTagRequestObject) (UpdateWorkItemTagResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
