package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkItemComment(c *gin.Context, request CreateWorkItemCommentRequestObject) (CreateWorkItemCommentResponseObject, error) {
	tx := GetTxFromCtx(c)

	params := request.Body.WorkItemCommentCreateParams

	workItemComment, err := h.svc.WorkItemComment.Create(c, tx, &params)
	if err != nil {
		renderErrorResponse(c, "Could not create work item comment", err)

		return nil, nil
	}

	res := WorkItemComment{
		WorkItemComment: *workItemComment,
		// joins, if any
	}

	return CreateWorkItemComment201JSONResponse(res), nil
}

func (h *StrictHandlers) GetWorkItemComment(c *gin.Context, request GetWorkItemCommentRequestObject) (GetWorkItemCommentResponseObject, error) {
	tx := GetTxFromCtx(c)

	workItemComment, err := h.svc.WorkItemComment.ByID(c, tx, request.WorkItemCommentID)
	if err != nil {
		renderErrorResponse(c, "Could not create work item comment", err)

		return nil, nil
	}

	res := WorkItemComment{
		WorkItemComment: *workItemComment,
		// joins, if any
	}

	return GetWorkItemComment200JSONResponse(res), nil
}

func (h *StrictHandlers) UpdateWorkItemComment(c *gin.Context, request UpdateWorkItemCommentRequestObject) (UpdateWorkItemCommentResponseObject, error) {
	tx := GetTxFromCtx(c)
	caller, _ := GetUserCallerFromCtx(c)

	params := request.Body.WorkItemCommentUpdateParams

	workItemComment, err := h.svc.WorkItemComment.Update(c, tx, caller, db.WorkItemCommentID(request.WorkItemCommentID), &params)
	if err != nil {
		renderErrorResponse(c, "Could not update work item comment", err)

		return nil, nil
	}

	res := WorkItemComment{
		WorkItemComment: *workItemComment,
		// joins, if any
	}

	return UpdateWorkItemComment200JSONResponse(res), nil
}

func (h *StrictHandlers) DeleteWorkItemComment(c *gin.Context, request DeleteWorkItemCommentRequestObject) (DeleteWorkItemCommentResponseObject, error) {
	tx := GetTxFromCtx(c)
	caller, _ := GetUserCallerFromCtx(c)

	_, err := h.svc.WorkItemComment.Delete(c, tx, caller, db.WorkItemCommentID(request.WorkItemCommentID))
	if err != nil {
		renderErrorResponse(c, "Could not delete work item comment", err)

		return nil, nil
	}

	return DeleteWorkItemComment204Response{}, nil
}
