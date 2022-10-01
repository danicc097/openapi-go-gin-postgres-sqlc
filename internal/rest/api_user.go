package rest

import (
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// User handles routes with the 'user' tag.
type User struct {
	logger   *zap.Logger
	userSvc  UserService
	authnSvc AuthenticationService
	authzSvc AuthorizationService
}

// NewUser returns a new handler for the 'user' route group.
func NewUser(
	logger *zap.Logger,
	userSvc UserService,
	authnSvc AuthenticationService,
	authzSvc AuthorizationService,
) *User {
	return &User{
		logger:   logger,
		userSvc:  userSvc,
		authnSvc: authnSvc,
		authzSvc: authzSvc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *User) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(createUser),
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.createUser,
			Middlewares: h.middlewares(createUser),
		},
		{
			Name:        string(createUsersWithArrayInput),
			Method:      http.MethodPost,
			Pattern:     "/user/createWithArray",
			HandlerFunc: h.createUsersWithArrayInput,
			Middlewares: h.middlewares(createUsersWithArrayInput),
		},
		{
			Name:        string(deleteUser),
			Method:      http.MethodDelete,
			Pattern:     "/user/:username",
			HandlerFunc: h.deleteUser,
			Middlewares: h.middlewares(deleteUser),
		},
		{
			Name:        string(getUserByName),
			Method:      http.MethodGet,
			Pattern:     "/user/:username",
			HandlerFunc: h.getUserByName,
			Middlewares: h.middlewares(getUserByName),
		},
		{
			Name:        string(loginUser),
			Method:      http.MethodGet,
			Pattern:     "/user/login",
			HandlerFunc: h.loginUser,
			Middlewares: h.middlewares(loginUser),
		},
		{
			Name:        string(logoutUser),
			Method:      http.MethodGet,
			Pattern:     "/user/logout",
			HandlerFunc: h.logoutUser,
			Middlewares: h.middlewares(logoutUser),
		},
		{
			Name:        string(updateUser),
			Method:      http.MethodPut,
			Pattern:     "/user/:username",
			HandlerFunc: h.updateUser,
			Middlewares: h.middlewares(updateUser),
		},
	}

	registerRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *User) middlewares(opID userOpID) []gin.HandlerFunc {
	authMw := newAuthMiddleware(h.logger, h.authnSvc, h.authzSvc, h.userSvc)

	switch opID {
	case createUser:
		return []gin.HandlerFunc{authMw.EnsureAuthenticated()}
	default:
		return []gin.HandlerFunc{}
	}
}

// CreateUser creates a new user.
func (h *User) createUser(c *gin.Context) {
	ctx := c.Request.Context()

	uid := ""
	if u := getUserFromCtx(c); u != nil {
		uid = fmt.Sprintf("%d", u.UserID)
	}
	uida := tracing.UserIDAttribute.String(uid)
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	s := newOTELSpan(ctx, "User.CreateUser", trace.WithAttributes(uida))
	s.AddEvent("create-user") // filterable with event="create-user"
	defer s.End()

	var user models.CreateUserRequest

	if err := c.BindJSON(&user); err != nil {
		renderErrorResponse(c, "error creating user", err)
		return
	}

	res, err := h.userSvc.Create(ctx, user)
	if err != nil {
		renderErrorResponse(c, "error creating user", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// createUsersWithArrayInput creates list of users with given input array.
func (h *User) createUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// deleteUser delete user.
func (h *User) deleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getUserByName get user by user name.
func (h *User) getUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// loginUser logs user into the system.
func (h *User) loginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// logoutUser logs out current logged in user session.
func (h *User) logoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updateUser updated user.
func (h *User) updateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// updatePet update an existing pet.
func (h *User) updatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
