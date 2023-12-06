package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteUser deletes the user by id.
func (h *Handlers) DeleteUser(c *gin.Context, id uuid.UUID) {
	defer newOTelSpanWithUser(c).End()

	tx := getTxFromCtx(c)

	_, err := h.svc.User.Delete(c, tx, db.NewUserID(id))
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)

		return
	}

	c.Status(http.StatusNoContent)
}

// GetCurrentUser returns the logged in user.
func (h *Handlers) GetCurrentUser(c *gin.Context) {
	defer newOTelSpanWithUser(c).End()

	caller := getUserFromCtx(c)

	span := getSpanFromCtx(c)
	span.AddEvent("get-current-user") // filterable with event="update-user"

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return
	}

	res := User{User: *caller, Role: role.Name}

	c.JSON(http.StatusOK, res)
}

// UpdateUser updates the user by id.
func (h *Handlers) UpdateUser(c *gin.Context, id uuid.UUID) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller := getUserFromCtx(c)

	span := getSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := getTxFromCtx(c)

	body := &models.UpdateUserRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	user, err := h.svc.User.Update(c, tx, db.UserID{UUID: id}, caller, body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return
	}

	role, ok := h.svc.Authorization.RoleByRank(user.RoleRank)
	if !ok {
		renderErrorResponse(c, fmt.Sprintf("Role with rank %d not found", user.RoleRank), nil)

		return
	}

	res := User{User: *user, Role: role.Name}

	renderResponse(c, res, http.StatusOK)
}

// UpdateUserAuthorization updates authorization information, e.g. roles, scopes.
func (h *Handlers) UpdateUserAuthorization(c *gin.Context, id uuid.UUID) {
	caller := getUserFromCtx(c)

	span := getSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := getTxFromCtx(c)

	body := &models.UpdateUserAuthRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, db.UserID{UUID: id}, caller, body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return
	}

	c.Status(http.StatusNoContent)
}
