package rest

import (
	"net/http"

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
	}

	registerRoutes(r, routes, "/user", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *User) middlewares(opID userOpID) []gin.HandlerFunc {
	authMw := newAuthMiddleware(h.logger, h.authnSvc, h.authzSvc, h.userSvc)

	switch opID {
	case getCurrentUser:
		return []gin.HandlerFunc{authMw.EnsureAuthenticated()}
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

// 	res, err := h.userSvc.Create(ctx, user)
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
