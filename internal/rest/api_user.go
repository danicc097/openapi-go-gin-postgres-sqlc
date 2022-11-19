package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
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
	//  user from context isntead has the appropiate joins already (teams, etc.)
	u := getUserFromCtx(c)
	if u == nil {
		renderErrorResponse(c, "user not found", errors.New("user not found"))

		return
	}
	c.JSON(http.StatusOK, u)
}

// UpdateUser updates the user by id.
func (h *Handlers) UpdateUser(c *gin.Context, id string) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	ctx := c.Request.Context()

	s := newOTELSpan(ctx, "User.UpdateUser", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		renderErrorResponse(c, "database error", err)
		return
	}
	defer tx.Rollback(ctx)

	body := &models.UpdateUserRequest{}

	if err := c.BindJSON(body); err != nil {
		renderErrorResponse(c, "invalid data", internal.NewErrorf(internal.ErrorCodeInvalidArgument, "invalid data", err))

		return
	}

	user := getUserFromCtx(c)
	if user == nil {
		renderErrorResponse(c, "Could not get user from context.", nil)

		return
	}

	var rank *int16
	if body.Role != nil {
		role, err := h.authmw.authzsvc.RoleByName(string(*body.Role))
		if err != nil {
			renderErrorResponse(c, fmt.Sprintf("Could not find role %q", *body.Role), err)
		}
		rank = &role.Rank
	}

	var scopes *[]string
	if body.Scopes != nil {
		ss := make([]string, 0, len(*body.Scopes))
		for _, r := range *body.Scopes {
			ss = append(ss, string(r))
		}
		scopes = &ss
	}
	// get target user by path param id. if target != current user and user rank <admin
	// not allowed to modify personal data.
	// NOTE: case for role and scopes update: any user can ADD (never delete unless rank >admin) another user's role and scope
	// as long as the current user has that scope.
	// but this belong in another operation: PATCH /user/:id/auth  op id UpdateUserAuth. INternally the service is the same: usvc.Update()
	u, err := h.usvc.Update(c, tx, repos.UserUpdateParams{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		ID:        id,
		Rank:      rank,
		Scopes:    scopes,
	})
	if err != nil {
		renderErrorResponse(c, "err: ", err)

		return
	}

	renderResponse(c, u, http.StatusOK)

	tx.Commit(ctx)
}
