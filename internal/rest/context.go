package rest

import (
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/gin-gonic/gin"
)

const userCtxKey = "user"

type authenticatedCtxKey struct{}

func GetUserFromCtx(c *gin.Context) *db.Users {
	user, ok := c.Value(userCtxKey).(*db.Users)
	if !ok {
		return nil
	}

	return user
}

func CtxWithUser(c *gin.Context, user *db.Users) {
	c.Set(userCtxKey, user)
}
