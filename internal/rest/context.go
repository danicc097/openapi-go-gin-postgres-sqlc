package rest

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"
)

const (
	userCtxKey                  = "user"
	userInfoCtxKey              = "user-info"
	ginContextCtxKey            = "middleware.openapi/gin-context"
	userDataCtxKey              = "middleware.openapi/user-data"
	validateResponseCtxKey      = "skip-response-validation"
	skipRequestValidationCtxKey = "skip-request-validation"
	transactionCtxKey           = "transaction"
	spanCtxKey                  = "span"
	errorCtxKey                 = "error"
)

func getSkipRequestValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipRequestValidationCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

func getValidateResponseFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(validateResponseCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

// getUserFromCtx returns basic information from the current user.
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

func getTxFromCtx(c *gin.Context) pgx.Tx {
	tx, ok := c.Value(transactionCtxKey).(pgx.Tx)
	if !ok {
		return nil
	}

	return tx
}

func ctxWithTx(c *gin.Context, tx pgx.Tx) {
	c.Set(transactionCtxKey, tx)
}

func getSpanFromCtx(c *gin.Context) trace.Span {
	span, ok := c.Value(spanCtxKey).(trace.Span)
	if !ok {
		return nil
	}

	return span
}

func ctxWithSpan(c *gin.Context, span trace.Span) {
	c.Set(spanCtxKey, span)
}

// Helper function to get the gin context from within requests. It returns
// nil if not found or wrong type.
// Useful for kin-openapi functions which only accept context.
func getGinContextFromCtx(c context.Context) *gin.Context {
	ginCtx, ok := c.Value(ginContextCtxKey).(*gin.Context)
	if !ok {
		return nil
	}
	return ginCtx
}

func getUserDataFromCtx(c context.Context) any {
	return c.Value(userDataCtxKey)
}

func ctxHasErrorResponse(c *gin.Context) bool {
	_, ok := c.Value(errorCtxKey).(struct{})

	return ok
}

// ctxWithErrorResponse signals current request will receive an error.
func ctxWithErrorResponse(c *gin.Context) {
	c.Set(errorCtxKey, struct{}{})
}
