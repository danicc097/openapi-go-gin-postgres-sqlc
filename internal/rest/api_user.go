package rest

import (
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateUser(c *gin.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller, _ := getUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	user, err := h.svc.User.Update(c, tx, db.UserID{UUID: request.Id}, caller, request.Body)
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

	return UpdateUser200JSONResponse(res), nil
}

func (h *StrictHandlers) DeleteUser(c *gin.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	tx := GetTxFromCtx(c)

	_, err := h.svc.User.Delete(c, tx, db.NewUserID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)

		return nil, nil
	}

	return DeleteUser204Response{}, nil
}

func (h *StrictHandlers) GetCurrentUser(c *gin.Context, request GetCurrentUserRequestObject) (GetCurrentUserResponseObject, error) {
	caller, _ := getUserCallerFromCtx(c)

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return nil, nil
	}

	res := User{
		User:     *caller.User,
		Role:     Role(role.Name),
		Teams:    &caller.Teams,
		Projects: &caller.Projects,
		APIKey:   caller.APIKey,
	}

	return GetCurrentUser200JSONResponse(res), nil
}

func (h *StrictHandlers) UpdateUserAuthorization(c *gin.Context, request UpdateUserAuthorizationRequestObject) (UpdateUserAuthorizationResponseObject, error) {
	caller, _ := getUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, db.UserID{UUID: request.Id}, caller, request.Body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return nil, nil
	}

	return UpdateUserAuthorization204Response{}, nil
}
