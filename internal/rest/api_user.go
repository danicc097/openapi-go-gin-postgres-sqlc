package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteUser deletes the user by id.
func (h *Handlers) DeleteUser(c *gin.Context, id uuid.UUID) {
	ctx := c.Request.Context()

	defer newOTelSpanWithUser(c).End()

	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	_, err := h.svc.user.Delete(c, tx, db.NewUserID(id))
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

	role, ok := h.svc.authz.RoleByRank(caller.RoleRank)
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

	s := newOTelSpanWithUser(c)
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	body := &models.UpdateUserRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	user, err := h.svc.user.Update(c, tx, db.UserID{UUID: id}, caller, body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "Database error", err)

		return
	}

	role, ok := h.svc.authz.RoleByRank(user.RoleRank)
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

	s := newOTelSpanWithUser(c)
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx := getTxFromCtx(c)
	defer tx.Rollback(ctx)

	body := &models.UpdateUserAuthRequest{}
	if shouldReturn := parseBody(c, body); shouldReturn {
		return
	}

	if _, err := h.svc.user.UpdateUserAuthorization(c, tx, db.UserID{UUID: id}, caller, body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return
	}

	if err := tx.Commit(ctx); err != nil {
		renderErrorResponse(c, "Database error", internal.WrapErrorf(err, models.ErrorCodePrivate, "could not commit transaction"))

		return
	}

	c.Status(http.StatusNoContent)
}
