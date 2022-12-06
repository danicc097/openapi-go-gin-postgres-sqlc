// Code generated by gowrap. DO NOT EDIT.
// template: ../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package repos

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserWithRetry implements User interface instrumented with retries
type UserWithRetry struct {
	User
	_retryCount    int
	_retryInterval time.Duration
}

// NewUserWithRetry returns UserWithRetry
func NewUserWithRetry(base User, retryCount int, retryInterval time.Duration) UserWithRetry {
	return UserWithRetry{
		User:           base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// Create implements User
func (_d UserWithRetry) Create(ctx context.Context, d db.DBTX, params UserCreateParams) (up1 *db.User, err error) {
	up1, err = _d.User.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.Create(ctx, d, params)
	}
	return
}

// CreateAPIKey implements User
func (_d UserWithRetry) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (up1 *db.UserAPIKey, err error) {
	up1, err = _d.User.CreateAPIKey(ctx, d, user)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.CreateAPIKey(ctx, d, user)
	}
	return
}

// Update implements User
func (_d UserWithRetry) Update(ctx context.Context, d db.DBTX, id string, params UserUpdateParams) (up1 *db.User, err error) {
	up1, err = _d.User.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.Update(ctx, d, id, params)
	}
	return
}

// UserByAPIKey implements User
func (_d UserWithRetry) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.User, err error) {
	up1, err = _d.User.UserByAPIKey(ctx, d, apiKey)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.UserByAPIKey(ctx, d, apiKey)
	}
	return
}

// UserByEmail implements User
func (_d UserWithRetry) UserByEmail(ctx context.Context, d db.DBTX, email string) (up1 *db.User, err error) {
	up1, err = _d.User.UserByEmail(ctx, d, email)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.UserByEmail(ctx, d, email)
	}
	return
}

// UserByExternalID implements User
func (_d UserWithRetry) UserByExternalID(ctx context.Context, d db.DBTX, extID string) (up1 *db.User, err error) {
	up1, err = _d.User.UserByExternalID(ctx, d, extID)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.UserByExternalID(ctx, d, extID)
	}
	return
}

// UserByID implements User
func (_d UserWithRetry) UserByID(ctx context.Context, d db.DBTX, id string) (up1 *db.User, err error) {
	up1, err = _d.User.UserByID(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.UserByID(ctx, d, id)
	}
	return
}

// UserByUsername implements User
func (_d UserWithRetry) UserByUsername(ctx context.Context, d db.DBTX, username string) (up1 *db.User, err error) {
	up1, err = _d.User.UserByUsername(ctx, d, username)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		up1, err = _d.User.UserByUsername(ctx, d, username)
	}
	return
}
