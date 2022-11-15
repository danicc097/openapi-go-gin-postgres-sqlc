package rest

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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
	// TODO return user from gin context (set when reading token in auth middleware)
	// getUserFromCtx(c)
	u, err := db.UserByUsername(c.Request.Context(), h.pool, "user_2", db.UserWithJoin(db.UserJoins{TimeEntries: true, Teams: true}))
	if err != nil {
		renderErrorResponse(c, "could not find user", err)
	}
	c.JSON(http.StatusOK, u)
}

// UpdateUser updates the user by id.
func (h *Handlers) UpdateUser(c *gin.Context, id string) {
	/*
		curl -X 'POST'   'https://localhost:8090/v2/user/{}'   -H 'accept: application/json'   -H 'Authorization: Bearer fsefse'  -d '{"username":"user","email":"email","role":"admin"}'
	*/
	// https://github.com/xo/xo/blob/master/_examples/booktest/sql/postgres_schema.sql
	// https://github.com/xo/xo/blob/master/_examples/booktest/postgres.go
	// we can call functions directly: presumably should also work for update on mat views, vacuum etc.
	// it can also generate custom queries like sqlc:
	// https://github.com/xo/xo/blob/master/_examples/booktest/sql/postgres_query.sql
	// is AuthorBookResultsByTags
	ctx := c.Request.Context()

	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		renderErrorResponse(c, "err::", err)
		return
	}
	defer tx.Rollback(ctx)

	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	s := newOTELSpan(ctx, "User.updateUser", trace.WithAttributes(userIDAttribute(c)))
	s.AddEvent("update-user") // filterable with event="update-user"
	defer s.End()

	// TODO back to OAS schema.
	// only role can be updated, username and email come from idp.
	// lets add first_name last_name.
	// var body models.UpdateUserRequest

	// if err := c.BindJSON(&body); err != nil {
	// 	renderErrorResponse(c, "err::", err)

	// 	return
	// }

	// // TODO extract to helper
	// var role db.UserRole

	// err = role.Scan([]byte(body.Role))
	// if err != nil {
	// 	renderErrorResponse(c, "err::", err)

	// 	return
	// }

	// h.logger.Sugar().Infof("body is :%#v", body)

	// user, err := h.usvc.UserByEmail(c, tx, body.Email)
	// if err != nil {
	// 	fmt.Printf("failed h.usvc.UserByEmail: %s\n", err)
	// }

	// h.logger.Sugar().Infof("user by email: %v", user)

	// if user == nil {
	// 	err = h.usvc.Register(c, tx, &db.User{
	// 		Username:  body.Username,
	// 		Email:     body.Email,
	// 		Role:      role,
	// 		FirstName: sql.NullString{String: "firstname", Valid: true},
	// 	})
	// 	if err != nil {
	// 		fmt.Printf("failed h.usvc.UserByEmail: %s\n", err)
	// 		renderErrorResponse(c, "user could not be created", err)

	// 		return
	// 	}
	// 	renderResponse(c, "user created", http.StatusOK)

	// 	return
	// }
	// user.Username = body.Username
	// user.Email = body.Email
	// user.Role = role

	// if err = h.usvc.Upsert(c, tx, user); err != nil {
	// 	renderErrorResponse(c, "err: ", err)

	// 	return
	// }

	tx.Commit(ctx)
}
