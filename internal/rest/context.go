package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

// how could we use custom types with gin context?
const (
	userCtxKey             = "user"
	ginContextKey          = "middleware.openapi/gin-context"
	userDataKey            = "middleware.openapi/user-data"
	skipResponseValidation = "skip-response-validation"
)

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
