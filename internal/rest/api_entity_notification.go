package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) CreateEntityNotification(c *gin.Context, request CreateEntityNotificationRequestObject) (CreateEntityNotificationResponseObject, error) {
	tx := GetTxFromCtx(c)

	params := request.Body.EntityNotificationCreateParams

	entityNotification, err := h.svc.EntityNotification.Create(c, tx, &params)
	if err != nil {
		renderErrorResponse(c, "Could not create entity notification", err)

		return nil, nil
	}

	res := EntityNotification{
		EntityNotification: *entityNotification,
		// joins, if any
	}

	return CreateEntityNotification201JSONResponse(res), nil
}

func (h *StrictHandlers) GetEntityNotification(c *gin.Context, request GetEntityNotificationRequestObject) (GetEntityNotificationResponseObject, error) {
	tx := GetTxFromCtx(c)

	entityNotification, err := h.svc.EntityNotification.ByID(c, tx, db.EntityNotificationID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not create entity notification", err)

		return nil, nil
	}

	res := EntityNotification{
		EntityNotification: *entityNotification,
		// joins, if any
	}

	return GetEntityNotification200JSONResponse(res), nil
}

func (h *StrictHandlers) UpdateEntityNotification(c *gin.Context, request UpdateEntityNotificationRequestObject) (UpdateEntityNotificationResponseObject, error) {
	tx := GetTxFromCtx(c)

	params := request.Body.EntityNotificationUpdateParams

	entityNotification, err := h.svc.EntityNotification.Update(c, tx, db.EntityNotificationID(request.Id), &params)
	if err != nil {
		renderErrorResponse(c, "Could not update entity notification", err)

		return nil, nil
	}

	res := EntityNotification{
		EntityNotification: *entityNotification,
		// joins, if any
	}

	return UpdateEntityNotification200JSONResponse(res), nil
}

func (h *StrictHandlers) DeleteEntityNotification(c *gin.Context, request DeleteEntityNotificationRequestObject) (DeleteEntityNotificationResponseObject, error) {
	tx := GetTxFromCtx(c)

	_, err := h.svc.EntityNotification.Delete(c, tx, db.EntityNotificationID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete entity notification", err)

		return nil, nil
	}

	return DeleteEntityNotification204Response{}, nil
}
