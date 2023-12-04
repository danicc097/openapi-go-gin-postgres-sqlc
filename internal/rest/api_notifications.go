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

	nn, err := h.svc.Notification.PaginatedNotifications(c.Request.Context(), h.pool, caller.UserID, params)
	if err != nil {
		renderErrorResponse(c, "Could not fetch notifications", err)

		return
	}

	items := make([]Notification, len(nn))
	for i, un := range nn {
		items[i] = Notification{
			UserNotification: un,
			Notification:     *un.NotificationJoin,
		}
	}
	res := PaginatedNotificationsResponse{
		Page: PaginationPage{
			NextCursor: fmt.Sprint(nn[len(nn)-1].UserNotificationID),
		},
		Items: items,
	}

	renderResponse(c, res, http.StatusOK)
}
