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

// create workitem comment.
func (h *Handlers) CreateWorkitemComment(c *gin.Context, id models.SerialID) {
	defer newOTelSpanWithUser(c).End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	_ = tx

	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) CreateWorkitem(c *gin.Context) {
	ctx := c.Request.Context()

	span := newOTelSpanWithUser(c)
	defer span.End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)

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

		res = DemoWorkItems{
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

		res = DemoTwoWorkItems{
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

func (h *Handlers) DeleteWorkitem(c *gin.Context, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) UpdateWorkitem(c *gin.Context, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func (h *Handlers) GetWorkItem(c *gin.Context, id models.SerialID) {
	c.JSON(http.StatusNotImplemented, "not implemented")
}
