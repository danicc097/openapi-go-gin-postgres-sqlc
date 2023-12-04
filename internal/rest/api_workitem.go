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

// create workitem.
func (h *Handlers) CreateWorkitem(c *gin.Context) {
	ctx := c.Request.Context()

	span := newOTelSpanWithUser(c)
	defer span.End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		renderErrorResponse(c, "Failed to read request body", err)

		return
	}
	span.SetAttributes(tracing.MetadataAttribute(jsonBody))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBody))

	body := &models.WorkItemCreateRequest{}
	if err := json.Unmarshal(jsonBody, body); err != nil {
		return
	}

	var res any // depends on project

	switch disc, _ := body.Discriminator(); models.Project(disc) {
	case models.ProjectDemo:
		body := &DemoWorkItemCreateRequest{}
		if shouldReturn := parseBody(c, body); shouldReturn {
			return
		}

		workItem, err := h.svc.DemoWorkItem.Create(ctx, tx, body.DemoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return
		}

		res = DemoWorkItemsResponse{
			WorkItem:             *workItem,
			SharedWorkItemFields: SharedWorkItemFields{},
			DemoWorkItem:         *workItem.DemoWorkItemJoin,
		}
	case models.ProjectDemoTwo:
		body := &DemoTwoWorkItemCreateRequest{}
		if shouldReturn := parseBody(c, body); shouldReturn {
			return
		}

		workItem, err := h.svc.DemoTwoWorkItem.Create(ctx, tx, body.DemoTwoWorkItemCreateParams)
		if err != nil {
			renderErrorResponse(c, "Could not create work item", err)

			return
		}

		res = DemoTwoWorkItemsResponse{
			WorkItem:             *workItem,
			SharedWorkItemFields: SharedWorkItemFields{},
			DemoTwoWorkItem:      *workItem.DemoTwoWorkItemJoin,
		}
	default:
		renderErrorResponse(c, fmt.Sprintf("Unknown project %q", disc), nil)

		return
	}

	c.JSON(http.StatusCreated, res)
}

// delete workitem.
func (h *Handlers) DeleteWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTelSpanWithUser(c).End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// get workitem.
func (h *Handlers) GetWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTelSpanWithUser(c).End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// update workitem.
func (h *Handlers) UpdateWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTelSpanWithUser(c).End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// create workitem comment.
func (h *Handlers) CreateWorkitemComment(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTelSpanWithUser(c).End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}
