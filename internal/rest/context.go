package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

const (
	userCtxKey             = "user"
	userInfoCtxKey         = "user-info"
	responseWriteCtxKey    = "response-writer"
	ginContextKey          = "middleware.openapi/gin-context"
	userDataKey            = "middleware.openapi/user-data"
	skipResponseValidation = "skip-response-validation"
	skipRequestValidation  = "skip-request-validation"
)

func getSkipRequestValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipRequestValidation).(bool)
	if !ok {
		return false
	}

	return skip
}

func getSkipResponseValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipResponseValidation).(bool)
	if !ok {
		return false
	}

	return skip
}

func getUserFromCtx(c *gin.Context) *db.User {
	user, ok := c.Value(userCtxKey).(*db.User)
	if !ok {
		return nil
	}

	return user
}

func ctxWithUser(c *gin.Context, user *db.User) {
	c.Set(userCtxKey, user)
}

func getUserInfoFromCtx(c *gin.Context) []byte {
	user, ok := c.Value(userInfoCtxKey).([]byte)
	if !ok {
		return nil
	}

	return user
}

func ctxWithUserInfo(c *gin.Context, userinfo []byte) {
	c.Set(userInfoCtxKey, userinfo)
}

// Helper function to get the gin context from within requests. It returns
// nil if not found or wrong type.
// TODO why would we need this?
func getGinContextFromCtx(c context.Context) *gin.Context {
	ginCtx, ok := c.Value(ginContextKey).(*gin.Context)
	if !ok {
		return nil
	}
	return ginCtx
}

func getUserDataFromCtx(c context.Context) any {
	return c.Value(userDataKey)
}
