package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create workitem
func (h *Handlers) CreateWorkitem(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "CreateWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// delete workitem
func (h *Handlers) DeleteWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "DeleteWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// get workitem
func (h *Handlers) GetWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "GetWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// update workitem
func (h *Handlers) UpdateWorkitem(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "UpdateWorkitem").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}

// create workitem comment
func (h *Handlers) CreateWorkitemComment(c *gin.Context, id int) {
	ctx := c.Request.Context()

	defer newOTELSpanWithUser(c, "CreateWorkitemComment").End()

	// caller := getUserFromCtx(c)
	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	c.JSON(http.StatusNotImplemented, "not implemented")
}
