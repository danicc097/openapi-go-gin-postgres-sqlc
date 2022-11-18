package rest

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
)

// how could we use custom types with gin context?
const (
	userCtxKey    = "user"
	ginContextKey = "middleware.openapi/gin-context"
	userDataKey   = "middleware.openapi/user-data"
)

func getUserFromCtx(c *gin.Context) *db.User {
	iuser, ok := c.Get(userCtxKey)
	fmt.Printf("getUserFromCtx iuser: %v\n", iuser)
	if !ok {
		return nil
	}

	user, ok := iuser.(*db.User)
	fmt.Printf("getUserFromCtx user: %v\n", user)
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
