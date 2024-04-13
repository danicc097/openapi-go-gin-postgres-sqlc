package rest

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateUser(c *gin.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller, _ := GetUserCallerFromCtx(c)

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

	res := User{User: user, Role: Role(role.Name)}

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
	caller, _ := GetUserCallerFromCtx(c)

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return nil, nil
	}

	res := User{
		User:     caller.User,
		Role:     Role(role.Name),
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

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, db.UserID{UUID: request.Id}, caller, request.Body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return nil, nil
	}

	return UpdateUserAuthorization204Response{}, nil
}

func formatCursorValue(value interface{}) (string, error) {
	switch v := value.(type) {
	case time.Time:
		return v.Format(time.RFC3339Nano), nil
	case string:
		return v, nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	default:
		return "", fmt.Errorf("unhandled cursor type: %v", v)
	}
}

func getNextCursor(entity interface{}, jsonFieldName string, tableEntity db.TableEntity) (string, error) {
	if entity == nil {
		return "", fmt.Errorf("no entity given")
	}

	if _, ok := db.EntityFields[tableEntity]; !ok {
		return "", fmt.Errorf("no entity found")
	}

	entityType := reflect.TypeOf(entity)
	entityValue := reflect.ValueOf(entity)

	for i := 0; i < entityType.NumField(); i++ {
		structField := entityType.Field(i)
		jsonTag := structField.Tag.Get("json")
		if jsonTag == jsonFieldName {
			fieldValue := entityValue.Field(i).Interface()
			return formatCursorValue(fieldValue)
		}
	}

	return "", fmt.Errorf("no json tag with value: %v", jsonFieldName)
}

func (h *StrictHandlers) GetPaginatedUsers(c *gin.Context, request GetPaginatedUsersRequestObject) (GetPaginatedUsersResponseObject, error) {
	users, err := h.svc.User.Paginated(c, h.pool, request.Params)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return nil, nil
	}

	nextCursor := ""
	if len(users) > 0 {
		lastUser := users[len(users)-1]
		nextCursor, err = getNextCursor(lastUser, request.Params.Column, db.TableEntityUser)
		if err != nil {
			renderErrorResponse(c, "Could not define next cursor", err)

			return nil, nil
		}
	}
	items := make([]User, len(users))
	for i, u := range users {
		u := u
		role, _ := h.svc.Authorization.RoleByRank(u.RoleRank)
		items[i] = User{
			User:     &u,
			Role:     Role(role.Name),
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
