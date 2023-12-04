package rest

import (
	"fmt"
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
	fmt.Printf("notifications: %v\n", notifications)
	res := PaginatedNotificationsResponse{
		Page: PaginationPage{
			NextCursor: fmt.Sprint(notifications[len(notifications)-1].UserNotificationID),
		},
		Items: notifications,
	}

	renderResponse(c, res, http.StatusOK)
}
