package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) GetPaginatedNotifications(c *gin.Context, request GetPaginatedNotificationsRequestObject) (GetPaginatedNotificationsResponseObject, error) {
	defer newOTelSpanWithUser(c).End()
	caller := getUserFromCtx(c)

	nn, err := h.svc.Notification.PaginatedNotifications(c.Request.Context(), h.pool, caller.UserID, request.Params)
	if err != nil {
		renderErrorResponse(c, "Could not fetch notifications", err)

		return nil, nil
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

	return GetPaginatedNotifications200JSONResponse(res), nil
}
