package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetPaginatedNotifications(c *gin.Context, params models.GetPaginatedNotificationsParams) {
	defer newOTelSpanWithUser(c).End()
	caller := getUserFromCtx(c)

	notifications, err := h.svc.Notification.PaginatedNotifications(c.Request.Context(), h.pool, caller.UserID, params)
	if err != nil {
		renderErrorResponse(c, "Could not fetch notifications", err)

		return
	}

	// TODO: pagination responses must have special format {_page: {nextCursor: <..>, ...}, items: <response>}
	// can have generics for PaginationBaseResponse struct, and its implementations in rest models
	// type MyPaginationResponse = PaginationBaseResponse[[]SomeThings] get converted to openapi schema
	// and validated as usual
	renderResponse(c, notifications, http.StatusOK)
}
