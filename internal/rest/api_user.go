package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

// DeleteUser deletes the user by id.
func (h *Handlers) DeleteUser(c *gin.Context, id uuid.UUID) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "DeleteUser", trace.WithAttributes(userIDAttribute(c))).End()

	tx := getTxFromCtx(c)

	_, err := h.usvc.Delete(c, tx, id)
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)

		return
	}

	c.Status(http.StatusNoContent)
}

// GetCurrentUser returns the logged in user.
func (h *Handlers) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "GetCurrentUser", trace.WithAttributes(userIDAttribute(c))).End()

	caller := getUserFromCtx(c)

	role, ok := h.authzsvc.RoleByRank(caller.RoleRank)
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
	ctx := c.Request.Context()
	caller := getUserFromCtx(c)

	s := newOTELSpan(ctx, "UpdateUser", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx := getTxFromCtx(c)

	body := &models.UpdateUserRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	user, err := h.usvc.Update(c, tx, id.String(), caller, body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "Error saving changes", err)

		return
	}

	role, ok := h.authzsvc.RoleByRank(user.RoleRank)
	if !ok {
		renderErrorResponse(c, fmt.Sprintf("Role with rank %d not found", user.RoleRank), nil)

		return
	}

	res := User{User: *user, Role: role.Name}

	renderResponse(c, res, http.StatusOK)
}

// UpdateUserAuthorization updates authorization information, e.g. roles, scopes.
func (h *Handlers) UpdateUserAuthorization(c *gin.Context, id uuid.UUID) {
	ctx := c.Request.Context()
	caller := getUserFromCtx(c)

	s := newOTELSpan(ctx, "UpdateUserAuthorization", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx := getTxFromCtx(c)

	body := &models.UpdateUserAuthRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	if _, err := h.usvc.UpdateUserAuthorization(c, tx, id.String(), caller, body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return
	}

	if err := tx.Commit(ctx); err != nil {
		renderErrorResponse(c, "Error saving changes", err)

		return
	}

	c.Status(http.StatusNoContent)
}
