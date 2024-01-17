package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"go.opentelemetry.io/otel/trace"
)

const ctxKeyPrefix = "rest-"

const (
	userCtxKey                  = ctxKeyPrefix + "user"
	userInfoCtxKey              = ctxKeyPrefix + "user-info"
	ginContextCtxKey            = ctxKeyPrefix + "middleware.openapi/gin-context"
	userDataCtxKey              = ctxKeyPrefix + "middleware.openapi/user-data"
	validateResponseCtxKey      = ctxKeyPrefix + "skip-response-validation"
	skipRequestValidationCtxKey = ctxKeyPrefix + "skip-request-validation"
	transactionCtxKey           = ctxKeyPrefix + "transaction"
	spanCtxKey                  = ctxKeyPrefix + "span"
	errorCtxKey                 = ctxKeyPrefix + "error"
)

func GetSkipRequestValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipRequestValidationCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

func GetValidateResponseFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(validateResponseCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

// getUserCallerFromCtx returns basic information from the current user.
func getUserCallerFromCtx(c *gin.Context) (services.CtxUser, error) {
	user, ok := c.Value(userCtxKey).(services.CtxUser)
	if !ok {
		return services.CtxUser{}, errors.New("user not found in ctx")
	}

	return user, nil
}

func CtxWithUserCaller(c *gin.Context, user *db.User) {
	c.Set(userCtxKey, services.CtxUser{
		User:     user,
		Teams:    *user.MemberTeamsJoin,
		Projects: *user.MemberProjectsJoin,
		APIKey:   user.APIKeyJoin,
	})
}

func GetUserInfoFromCtx(c *gin.Context) (*oidc.UserInfo, error) {
	userInfoBlob, ok := c.Value(userInfoCtxKey).([]byte)
	if !ok {
		return nil, errors.New("empty value")
	}
	var userInfo oidc.UserInfo
	err := json.Unmarshal(userInfoBlob, &userInfo)
	if err != nil {
		return nil, fmt.Errorf("could not load user info: %w", err)
	}

	return &userInfo, nil
}

func CtxWithUserInfo(c *gin.Context, userinfo []byte) {
	c.Set(userInfoCtxKey, userinfo)
}

// GetTxFromCtx returns the ongoing db transaction.
// Automatic commit and rollback is handled in db middleware.
func GetTxFromCtx(c *gin.Context) pgx.Tx {
	tx, ok := c.Value(transactionCtxKey).(pgx.Tx)
	if !ok {
		return nil
	}

	return tx
}

func ctxWithTx(c *gin.Context, tx pgx.Tx) {
	c.Set(transactionCtxKey, tx)
}

func GetSpanFromCtx(c *gin.Context) trace.Span {
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
func GetGinContextFromCtx(c context.Context) *gin.Context {
	ginCtx, ok := c.Value(ginContextCtxKey).(*gin.Context)
	if !ok {
		return nil
	}
	return ginCtx
}

func GetUserDataFromCtx(c context.Context) any {
	return c.Value(userDataCtxKey)
}

func CtxHasErrorResponse(c *gin.Context) bool {
	_, ok := c.Value(errorCtxKey).(struct{})

	return ok
}

// ctxWithErrorResponse signals current request will receive an error.
func ctxWithErrorResponse(c *gin.Context) {
	c.Set(errorCtxKey, struct{}{})
}
