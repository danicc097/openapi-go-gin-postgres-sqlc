package rest

import (
	"bytes"
	"io"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateWorkitem(c *gin.Context, request UpdateWorkitemRequestObject) (UpdateWorkitemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) DeleteWorkitem(c *gin.Context, request DeleteWorkitemRequestObject) (DeleteWorkitemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) CreateWorkitem(c *gin.Context, request CreateWorkitemRequestObject) (CreateWorkitemResponseObject, error) {
	ctx := c.Request.Context()

	span := GetSpanFromCtx(c)

	caller, _ := GetUserCallerFromCtx(c)
	tx := GetTxFromCtx(c)

	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		renderErrorResponse(c, "Failed to read request body", err)

		return nil, nil
	}
	span.SetAttributes(tracing.MetadataAttribute(jsonBody))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBody))

	// body := &models.CreateWorkItemRequest{}
	// if err := json.Unmarshal(jsonBody, body); err != nil {
	// 	return nil, nil
	// }

	var res any // depends on project
	b, err := request.Body.ValueByDiscriminator()
	if err != nil {
		renderErrorResponse(c, "Failed to read discriminator", err)

		return nil, nil
	}

	switch body := b.(type) {
	case CreateDemoWorkItemRequest:
		workItem, err := h.svc.DemoWorkItem.Create(ctx, tx, caller, body.DemoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return nil, nil
		}

		res = DemoWorkItem{
			WorkItemBase: fillBaseWorkItemResponse(workItem),
			DemoWorkItem: *workItem.DemoWorkItemJoin,
		}
	case CreateDemoTwoWorkItemRequest:
		workItem, err := h.svc.DemoTwoWorkItem.Create(ctx, tx, caller, body.DemoTwoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return nil, nil
		}

		res = DemoTwoWorkItem{
			WorkItemBase:    fillBaseWorkItemResponse(workItem),
			DemoTwoWorkItem: *workItem.DemoTwoWorkItemJoin,
		}
	default:
		renderErrorResponse(c, "Unknown body", internal.NewErrorf(models.ErrorCodeUnknown, "%+v", b))

		return nil, nil
	}

	return CreateWorkitem201JSONResponse{union: rawMessage(res)}, nil
}

func (h *StrictHandlers) GetWorkItem(c *gin.Context, request GetWorkItemRequestObject) (GetWorkItemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func fillBaseWorkItemResponse(workItem *db.WorkItem) WorkItemBase {
	return WorkItemBase{
		WorkItem: *workItem,
		SharedWorkItemJoins: SharedWorkItemJoins{
			Members:          workItem.AssigneesJoin,
			WorkItemTags:     workItem.WorkItemTagsJoin,
			TimeEntries:      workItem.TimeEntriesJoin,
			WorkItemComments: workItem.WorkItemCommentsJoin,
			WorkItemType:     workItem.WorkItemTypeJoin,
		},
	}
}
