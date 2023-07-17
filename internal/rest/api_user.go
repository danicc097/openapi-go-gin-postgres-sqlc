package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"
)

// DeleteUser deletes the user by id.
func (h *Handlers) DeleteUser(c *gin.Context, id string) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetCurrentUser returns the logged in user.
func (h *Handlers) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "GetCurrentUser", trace.WithAttributes(userIDAttribute(c))).End()

	//  user from context isntead has the appropriate joins already (teams, etc.)
	user := getUserFromCtx(c)

	role, ok := h.authzsvc.RoleByRank(user.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", user.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return
	}

	res := User{User: *user, Role: role.Name}

	c.JSON(http.StatusOK, res)
}

// UpdateUser updates the user by id.
func (h *Handlers) UpdateUser(c *gin.Context, id string) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	ctx := c.Request.Context()

	s := newOTELSpan(ctx, "UpdateUser", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "database error", internal.WrapErrorf(err, models.ErrorCodePrivate, "could not being tx"))

		return
	}
	defer tx.Rollback(ctx)

	body := &models.UpdateUserRequest{}

	if err := c.BindJSON(body); err != nil {
		renderErrorResponse(c, "invalid data", internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "invalid data"))

		return
	}

	caller := getUserFromCtx(c)
	if caller == nil {
		renderErrorResponse(c, "Could not get current user", nil)

		return
	}

	user, err := h.usvc.Update(c, tx, id, caller, body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "Could not save changes", err)

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

// UpdateUserAuthorization updates authorizastion information, e.g. roles and scopes.
func (h *Handlers) UpdateUserAuthorization(c *gin.Context, id string) {
	ctx := c.Request.Context()

	s := newOTELSpan(ctx, "UpdateUserAuthorization", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "database error", internal.WrapErrorf(err, models.ErrorCodePrivate, "could not being tx"))

		return
	}
	defer tx.Rollback(ctx)

	body := &models.UpdateUserAuthRequest{}

	if err := c.BindJSON(body); err != nil {
		renderErrorResponse(c, "invalid data", internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "invalid data"))

		return
	}

	caller := getUserFromCtx(c)
	if caller == nil {
		renderErrorResponse(c, "Could not get current user", nil)

		return
	}

	if _, err := h.usvc.UpdateUserAuthorization(c, tx, id, caller, body); err != nil {
		renderErrorResponse(c, "could not update user auth", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "could not save changes", err)

		return
	}

	c.Status(http.StatusNoContent)
}
