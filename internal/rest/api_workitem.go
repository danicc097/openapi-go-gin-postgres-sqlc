package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// create workitem
func (h *Handlers) CreateWorkitem(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "CreateWorkitem", trace.WithAttributes(userIDAttribute(c))).End()

	// caller := getUserFromCtx(c)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// delete workitem
func (h *Handlers) DeleteWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "DeleteWorkitem", trace.WithAttributes(userIDAttribute(c))).End()

	// caller := getUserFromCtx(c)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// get workitem
func (h *Handlers) GetWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "GetWorkitem", trace.WithAttributes(userIDAttribute(c))).End()

	// caller := getUserFromCtx(c)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// update workitem
func (h *Handlers) UpdateWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "UpdateWorkitem", trace.WithAttributes(userIDAttribute(c))).End()

	// caller := getUserFromCtx(c)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// create workitem comment
func (h *Handlers) CreateWorkitemComment(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "CreateWorkitemComment", trace.WithAttributes(userIDAttribute(c))).End()

	// caller := getUserFromCtx(c)

	c.JSON(http.StatusNotImplemented, "not implemented")
}
