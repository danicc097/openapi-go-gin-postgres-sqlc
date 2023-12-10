package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateUser(c *gin.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller := getUserFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	body := &models.UpdateUserRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return nil, nil
	}

	user, err := h.svc.User.Update(c, tx, db.UserID{UUID: request.Id}, caller, body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return nil, nil
	}

	role, ok := h.svc.Authorization.RoleByRank(user.RoleRank)
	if !ok {
		renderErrorResponse(c, fmt.Sprintf("Role with rank %d not found", user.RoleRank), nil)

		return nil, nil
	}

	res := User{User: *user, Role: Role(role.Name)}

	renderResponse(c, res, http.StatusOK)
	return nil, nil
}

func (h *StrictHandlers) DeleteUser(c *gin.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	defer newOTelSpanWithUser(c).End()

	tx := GetTxFromCtx(c)

	_, err := h.svc.User.Delete(c, tx, db.NewUserID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)

		return nil, nil
	}

	c.Status(http.StatusNoContent)
	return nil, nil
}

func (h *StrictHandlers) GetCurrentUser(c *gin.Context, request GetCurrentUserRequestObject) (GetCurrentUserResponseObject, error) {
	defer newOTelSpanWithUser(c).End()

	caller := getUserFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("get-current-user") // filterable with event="update-user"

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return nil, nil
	}

	res := User{User: *caller, Role: Role(role.Name)}

	c.JSON(http.StatusOK, res)
	return nil, nil
}

func (h *StrictHandlers) UpdateUserAuthorization(c *gin.Context, request UpdateUserAuthorizationRequestObject) (UpdateUserAuthorizationResponseObject, error) {
	caller := getUserFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	body := &models.UpdateUserAuthRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return nil, nil
	}

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, db.UserID{UUID: request.Id}, caller, body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return nil, nil
	}

	c.Status(http.StatusNoContent)
	return nil, nil
}
