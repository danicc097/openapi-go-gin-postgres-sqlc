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

// createUser creates a new user.
// TODO remove handler once oidc imp., but will use the service in /login.
// we can use upsert on every new login with xoxo to ensure email and username are always up to date
// or registered the first time
// func (h *User) createUser(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	// span attribute not inheritable:
// 	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
// 	s := newOTELSpan(ctx, "User.CreateUser", trace.WithAttributes(userIDAttribute(c)))
// 	s.AddEvent("create-user") // filterable with event="create-user"
// 	defer s.End()

// 	var user models.CreateUserRequest

// 	if err := c.BindJSON(&user); err != nil {
// 		renderErrorResponse(c, "error creating user", err)

// 		return
// 	}

// 	res, err := h.usvc.Create(ctx, user)
// 	if err != nil {
// 		renderErrorResponse(c, "error creating user", err)

// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// DeleteUser deletes the user by id.
func (h *Handlers) DeleteUser(c *gin.Context, id string) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetCurrentUser returns the logged in user.
func (h *Handlers) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()

	defer newOTELSpan(ctx, "User.UpdateUser", trace.WithAttributes(userIDAttribute(c))).End()

	//  user from context isntead has the appropiate joins already (teams, etc.)
	user := getUserFromCtx(c)
	if user == nil {
		renderErrorResponse(c, "user not found", errors.New("user not found"))

		return
	}

	role, ok := h.authzsvc.RoleByRank(user.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", user.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return
	}

	res := UserResponse{UserPublic: user.ToPublic(), Role: role.Role, Scopes: user.Scopes}

	c.JSON(http.StatusOK, res)
}

// UpdateUser updates the user by id.
func (h *Handlers) UpdateUser(c *gin.Context, id string) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	ctx := c.Request.Context()

	s := newOTELSpan(ctx, "User.UpdateUser", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "database error", err)

		return
	}
	defer tx.Rollback(ctx)

	body := &models.UpdateUserRequest{}

	if err := c.BindJSON(body); err != nil {
		renderErrorResponse(c, "invalid data", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid data"))

		return
	}

	caller := getUserFromCtx(c)
	if caller == nil {
		renderErrorResponse(c, "Could not get user from context.", nil)

		return
	}

	user, err := h.usvc.Update(c, tx, id, caller, body)
	if err != nil {
		renderErrorResponse(c, "err: ", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "could not save changes", err)

		return
	}

	renderResponse(c, user, http.StatusOK)
}

// UpdateUserAuthorization updates authorizastion information, e.g. roles and scopes.
func (h *Handlers) UpdateUserAuthorization(c *gin.Context, id string) {
	ctx := c.Request.Context()

	s := newOTELSpan(ctx, "User.UpdateUserAuthorization", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		renderErrorResponse(c, "database error", err)

		return
	}
	defer tx.Rollback(ctx)

	body := &models.UpdateUserAuthRequest{}

	if err := c.BindJSON(body); err != nil {
		renderErrorResponse(c, "invalid data", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid data"))

		return
	}

	caller := getUserFromCtx(c)
	if caller == nil {
		renderErrorResponse(c, "Could not get user from context.", nil)

		return
	}

	user, err := h.usvc.UpdateUserAuthorization(c, tx, id, caller, body)
	if err != nil {
		renderErrorResponse(c, "err: ", err)

		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		renderErrorResponse(c, "could not save changes", err)

		return
	}

	renderResponse(c, user, http.StatusOK)
}
