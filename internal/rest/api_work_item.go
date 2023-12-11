package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateWorkitemComment(c *gin.Context, request CreateWorkitemCommentRequestObject) (CreateWorkitemCommentResponseObject, error) {
	// caller := getUserFromCtx(c)
	tx := GetTxFromCtx(c)
	_ = tx

	c.JSON(http.StatusNotImplemented, "not implemented")
	return nil, nil
}

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

	// caller := getUserFromCtx(c)
	tx := GetTxFromCtx(c)

	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		renderErrorResponse(c, "Failed to read request body", err)

		return nil, nil
	}
	span.SetAttributes(tracing.MetadataAttribute(jsonBody))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBody))

	body := &models.CreateWorkItemRequest{}
	if err := json.Unmarshal(jsonBody, body); err != nil {
		return nil, nil
	}

	var res any // depends on project

	switch disc, _ := body.Discriminator(); models.Project(disc) {
	case models.ProjectDemo:
		body := &CreateDemoWorkItemRequest{}
		if shouldReturn := parseBody(c, body); shouldReturn {
			return nil, nil
		}

		workItem, err := h.svc.DemoWorkItem.Create(ctx, tx, body.DemoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return nil, nil
		}

		res = DemoWorkItems{
			WorkItem:             *workItem,
			SharedWorkItemFields: SharedWorkItemFields{},
			DemoWorkItem:         *workItem.DemoWorkItemJoin,
		}
	case models.ProjectDemoTwo:
		body := &CreateDemoTwoWorkItemRequest{}
		if shouldReturn := parseBody(c, body); shouldReturn {
			return nil, nil
		}

		workItem, err := h.svc.DemoTwoWorkItem.Create(ctx, tx, body.DemoTwoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return nil, nil
		}

		res = DemoTwoWorkItems{
			WorkItem:             *workItem,
			SharedWorkItemFields: SharedWorkItemFields{},
			DemoTwoWorkItem:      *workItem.DemoTwoWorkItemJoin,
		}
	default:
		renderErrorResponse(c, fmt.Sprintf("Unknown project %q", disc), nil)

		return nil, nil
	}

	c.JSON(http.StatusCreated, res)

	return nil, nil
}

func (h *StrictHandlers) GetWorkItem(c *gin.Context, request GetWorkItemRequestObject) (GetWorkItemResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
