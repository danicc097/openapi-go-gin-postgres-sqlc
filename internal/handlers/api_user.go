package handlers

import (
	"context"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/gin-gonic/gin"
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
	routes := []rest.Route{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "/user",
			HandlerFunc: h.CreateUser,
			Middlewares: h.middlewares("CreateUser"),
		},
		{
			Name:        "CreateUsersWithArrayInput",
			Method:      http.MethodPost,
			Pattern:     "/user/createWithArray",
			HandlerFunc: h.CreateUsersWithArrayInput,
			Middlewares: h.middlewares("CreateUsersWithArrayInput"),
		},
		{
			Name:        "DeleteUser",
			Method:      http.MethodDelete,
			Pattern:     "/user/:username",
			HandlerFunc: h.DeleteUser,
			Middlewares: h.middlewares("DeleteUser"),
		},
		{
			Name:        "GetUserByName",
			Method:      http.MethodGet,
			Pattern:     "/user/:username",
			HandlerFunc: h.GetUserByName,
			Middlewares: h.middlewares("GetUserByName"),
		},
		{
			Name:        "LoginUser",
			Method:      http.MethodGet,
			Pattern:     "/user/login",
			HandlerFunc: h.LoginUser,
			Middlewares: h.middlewares("LoginUser"),
		},
		{
			Name:        "LogoutUser",
			Method:      http.MethodGet,
			Pattern:     "/user/logout",
			HandlerFunc: h.LogoutUser,
			Middlewares: h.middlewares("LogoutUser"),
		},
		{
			Name:        "UpdateUser",
			Method:      http.MethodPut,
			Pattern:     "/user/:username",
			HandlerFunc: h.UpdateUser,
			Middlewares: h.middlewares("UpdateUser"),
		},
	}

	rest.RegisterRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *User) middlewares(opID string) []gin.HandlerFunc {
	authMw := NewAuthMw(h.logger, h.authnSvc, h.authzSvc, h.userSvc)

	switch opID {
	case "CreateUser":
		return []gin.HandlerFunc{authMw.EnsureAuthenticated()}
	default:
		return []gin.HandlerFunc{}
	}
}

// CreateUser creates a new user.
func (h *User) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	if err := c.BindJSON(&user); err != nil {
		h.logger.Sugar().Debugf("CreateUser.user: %v", user)
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	res, err := h.userSvc.Create(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateUsersWithArrayInput creates list of users with given input array.
func (h *User) CreateUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeleteUser delete user.
func (h *User) DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetUserByName get user by user name.
func (h *User) GetUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LoginUser logs user into the system.
func (h *User) LoginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LogoutUser logs out current logged in user session.
func (h *User) LogoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdateUser updated user.
func (h *User) UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}
