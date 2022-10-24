package rest

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// User handles routes with the 'user' tag.
type User struct {
	usvc           *services.User
	logger         *zap.Logger
	pool           *pgxpool.Pool
	movieSvcClient v1.MovieGenreClient
	authmw         *authMiddleware
}

// NewUser returns a new handler for the 'user' route group.
func NewUser(
	logger *zap.Logger,
	pool *pgxpool.Pool,
	movieSvcClient v1.MovieGenreClient,
	usvc *services.User,
	authmw *authMiddleware,
) *User {
	return &User{
		logger:         logger,
		pool:           pool,
		movieSvcClient: movieSvcClient,
		usvc:           usvc,
		authmw:         authmw,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(deleteUser),
			Method:      http.MethodDelete,
			Pattern:     "/user/:id",
			HandlerFunc: h.deleteUser,
			Middlewares: h.middlewares(deleteUser),
		},
		{
			Name:        string(getCurrentUser),
			Method:      http.MethodGet,
			Pattern:     "/user/me",
			HandlerFunc: h.getCurrentUser,
			Middlewares: h.middlewares(getCurrentUser),
		},
		{
			Name:        string(updateUser),
			Method:      http.MethodPut,
			Pattern:     "/user/:id",
			HandlerFunc: h.updateUser,
			Middlewares: h.middlewares(updateUser),
		},
	}

	registerRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *User) middlewares(opID userOpID) []gin.HandlerFunc {
	switch opID {
	case deleteUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthorized(db.RoleAdmin),
		}
	case updateUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthorized(db.RoleAdmin),
		}
	default:
		return []gin.HandlerFunc{}
	}
}

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

// deleteUser deletes the user by id.
func (h *User) deleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getCurrentUser returns the logged in user.
func (h *User) getCurrentUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updateUser updates the user by id.
func (h *User) updateUser(c *gin.Context) {
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
	var body models.UpdateUserRequest

	if err := c.BindJSON(&body); err != nil {
		renderErrorResponse(c, "err::", err)

		return
	}

	// TODO extract to helper
	var role db.Role

	err = role.Scan([]byte(body.Role))
	if err != nil {
		renderErrorResponse(c, "err::", err)

		return
	}

	h.logger.Sugar().Infof("body is :%#v", body)

	user, err := h.usvc.UserByEmail(c, tx, body.Email)
	if err != nil {
		fmt.Printf("failed h.usvc.UserByEmail: %s\n", err)
	}

	h.logger.Sugar().Infof("user by email: %v", user)

	if user == nil {
		err = h.usvc.Register(c, tx, &db.User{
			Username:  body.Username,
			Email:     body.Email,
			Role:      role,
			FirstName: sql.NullString{String: "firstname", Valid: true},
		})
		if err != nil {
			fmt.Printf("failed h.usvc.UserByEmail: %s\n", err)
			renderErrorResponse(c, "user could not be created", err)

			return
		}
		renderResponse(c, "user created", http.StatusOK)

		return
	}
	user.Username = body.Username
	user.Email = body.Email
	user.Role = role

	if err = h.usvc.Upsert(c, tx, user); err != nil {
		renderErrorResponse(c, "err: ", err)

		return
	}

	tx.Commit(ctx)
}
