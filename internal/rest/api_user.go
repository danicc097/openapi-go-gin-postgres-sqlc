package rest

import (
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateUser(c *gin.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller, _ := GetUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	user, err := h.svc.User.Update(c, tx, models.UserID{UUID: request.Id}, caller, request.Body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)
	}

	role, ok := h.svc.Authorization.RoleByRank(user.RoleRank)
	if !ok {
		renderErrorResponse(c, fmt.Sprintf("Role with rank %d not found", user.RoleRank), nil)
	}

	res := UserResponse{User: user, Role: role.Name}

	return UpdateUser200JSONResponse(res), nil
}

func (h *StrictHandlers) DeleteUser(c *gin.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	tx := GetTxFromCtx(c)

	_, err := h.svc.User.Delete(c, tx, models.NewUserID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)
	}

	return DeleteUser204Response{}, nil
}

func (h *StrictHandlers) GetCurrentUser(c *gin.Context, request GetCurrentUserRequestObject) (GetCurrentUserResponseObject, error) {
	caller, _ := GetUserCallerFromCtx(c)

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))
	}

	res := UserResponse{
		User:     caller.User,
		Role:     role.Name,
		Teams:    &caller.Teams,
		Projects: &caller.Projects,
		APIKey:   caller.APIKey,
	}

	return GetCurrentUser200JSONResponse(res), nil
}

func (h *StrictHandlers) UpdateUserAuthorization(c *gin.Context, request UpdateUserAuthorizationRequestObject) (UpdateUserAuthorizationResponseObject, error) {
	caller, _ := GetUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, models.UserID{UUID: request.Id}, caller, request.Body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)
	}

	return UpdateUserAuthorization204Response{}, nil
}

func (h *StrictHandlers) GetPaginatedUsers(c *gin.Context, request GetPaginatedUsersRequestObject) (GetPaginatedUsersResponseObject, error) {
	users, err := h.svc.User.Paginated(c, h.pool, request.Params)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)
	}

	nextCursor := ""
	if len(users) > 0 {
		lastUser := users[len(users)-1]
		nextCursor, err = getNextCursor(lastUser, request.Params.Column, models.TableEntityUser)
		if err != nil {
			renderErrorResponse(c, "Could not define next cursor", err)
		}
	}
	items := make([]UserResponse, len(users))
	for i, u := range users {
		u := u
		role, _ := h.svc.Authorization.RoleByRank(u.RoleRank)
		items[i] = UserResponse{
			User:     &u,
			Role:     role.Name,
			Teams:    u.MemberTeamsJoin,
			Projects: u.MemberProjectsJoin,
		}
	}
	res := PaginatedUsersResponse{
		Page: PaginationPage{
			NextCursor: nextCursor,
		},
		Items: items,
	}

	return GetPaginatedUsers200JSONResponse(res), nil
}
