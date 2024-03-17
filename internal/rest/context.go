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
	userCtxKey                   = ctxKeyPrefix + "user"
	userInfoCtxKey               = ctxKeyPrefix + "user-info"
	ginContextCtxKey             = ctxKeyPrefix + "middleware.openapi/gin-context"
	userDataCtxKey               = ctxKeyPrefix + "middleware.openapi/user-data"
	skipResponseValidationCtxKey = ctxKeyPrefix + "skip-response-validation"
	skipRequestValidationCtxKey  = ctxKeyPrefix + "skip-request-validation"
	transactionCtxKey            = ctxKeyPrefix + "transaction"
	spanCtxKey                   = ctxKeyPrefix + "span"
	errorCtxKey                  = ctxKeyPrefix + "error"
)

type requestIDCtxKey struct{}

// NOTE: request ID is set on Request's context since it may be used by services.
func GetRequestIDFromCtx(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDCtxKey{}).(string)
	return requestID
}

func GetSkipRequestValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipRequestValidationCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

func GetSkipResponseValidationFromCtx(c *gin.Context) bool {
	skip, ok := c.Value(skipResponseValidationCtxKey).(bool)
	if !ok {
		return false
	}

	return skip
}

// GetUserCallerFromCtx returns basic information from the current user.
func GetUserCallerFromCtx(c *gin.Context) (services.CtxUser, error) {
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
		APIKey:   user.UserAPIKeyJoin,
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

func CtxWithTx(c *gin.Context, tx pgx.Tx) {
	c.Set(transactionCtxKey, tx)
}

func GetSpanFromCtx(c *gin.Context) trace.Span {
	span, ok := c.Value(spanCtxKey).(trace.Span)
	if !ok {
		return nil
	}

	return span
}

func CtxWithSpan(c *gin.Context, span trace.Span) {
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

// GetRequestHasErrorFromCtx returns whether the current request has an error.
func GetRequestHasErrorFromCtx(c *gin.Context) bool {
	_, ok := c.Value(errorCtxKey).(struct{})

	return ok
}

// CtxWithRequestError signals that the current request has an error.
func CtxWithRequestError(c *gin.Context) {
	c.Set(errorCtxKey, struct{}{})
}
