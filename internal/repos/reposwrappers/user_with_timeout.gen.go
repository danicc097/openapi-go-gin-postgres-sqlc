// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	_sourceRepos "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// UserWithTimeout implements _sourceRepos.User interface instrumented with timeouts
type UserWithTimeout struct {
	_sourceRepos.User
	config UserWithTimeoutConfig
}

type UserWithTimeoutConfig struct {
	ByAPIKeyTimeout time.Duration

	ByEmailTimeout time.Duration

	ByExternalIDTimeout time.Duration

	ByIDTimeout time.Duration

	ByProjectTimeout time.Duration

	ByTeamTimeout time.Duration

	ByUsernameTimeout time.Duration

	CreateTimeout time.Duration

	CreateAPIKeyTimeout time.Duration

	DeleteTimeout time.Duration

	DeleteAPIKeyTimeout time.Duration

	PaginatedTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewUserWithTimeout returns UserWithTimeout
func NewUserWithTimeout(base _sourceRepos.User, config UserWithTimeoutConfig) UserWithTimeout {
	return UserWithTimeout{
		User:   base,
		config: config,
	}
}

// ByAPIKey implements _sourceRepos.User
func (_d UserWithTimeout) ByAPIKey(ctx context.Context, d models.DBTX, apiKey string) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.ByAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.ByAPIKey(ctx, d, apiKey)
}

// ByEmail implements _sourceRepos.User
func (_d UserWithTimeout) ByEmail(ctx context.Context, d models.DBTX, email string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.ByEmailTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByEmailTimeout)
		defer cancelFunc()
	}
	return _d.User.ByEmail(ctx, d, email, opts...)
}

// ByExternalID implements _sourceRepos.User
func (_d UserWithTimeout) ByExternalID(ctx context.Context, d models.DBTX, extID string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.ByExternalIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByExternalIDTimeout)
		defer cancelFunc()
	}
	return _d.User.ByExternalID(ctx, d, extID, opts...)
}

// ByID implements _sourceRepos.User
func (_d UserWithTimeout) ByID(ctx context.Context, d models.DBTX, id models.UserID, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.User.ByID(ctx, d, id, opts...)
}

// ByProject implements _sourceRepos.User
func (_d UserWithTimeout) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID) (ua1 []models.User, err error) {
	var cancelFunc func()
	if _d.config.ByProjectTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByProjectTimeout)
		defer cancelFunc()
	}
	return _d.User.ByProject(ctx, d, projectID)
}

// ByTeam implements _sourceRepos.User
func (_d UserWithTimeout) ByTeam(ctx context.Context, d models.DBTX, teamID models.TeamID) (ua1 []models.User, err error) {
	var cancelFunc func()
	if _d.config.ByTeamTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByTeamTimeout)
		defer cancelFunc()
	}
	return _d.User.ByTeam(ctx, d, teamID)
}

// ByUsername implements _sourceRepos.User
func (_d UserWithTimeout) ByUsername(ctx context.Context, d models.DBTX, username string, opts ...models.UserSelectConfigOption) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.ByUsernameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByUsernameTimeout)
		defer cancelFunc()
	}
	return _d.User.ByUsername(ctx, d, username, opts...)
}

// Create implements _sourceRepos.User
func (_d UserWithTimeout) Create(ctx context.Context, d models.DBTX, params *models.UserCreateParams) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.User.Create(ctx, d, params)
}

// CreateAPIKey implements _sourceRepos.User
func (_d UserWithTimeout) CreateAPIKey(ctx context.Context, d models.DBTX, user *models.User) (up1 *models.UserAPIKey, err error) {
	var cancelFunc func()
	if _d.config.CreateAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.CreateAPIKey(ctx, d, user)
}

// Delete implements _sourceRepos.User
func (_d UserWithTimeout) Delete(ctx context.Context, d models.DBTX, id models.UserID) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.User.Delete(ctx, d, id)
}

// DeleteAPIKey implements _sourceRepos.User
func (_d UserWithTimeout) DeleteAPIKey(ctx context.Context, d models.DBTX, apiKey string) (up1 *models.UserAPIKey, err error) {
	var cancelFunc func()
	if _d.config.DeleteAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.DeleteAPIKey(ctx, d, apiKey)
}

// Paginated implements _sourceRepos.User
func (_d UserWithTimeout) Paginated(ctx context.Context, d models.DBTX, params _sourceRepos.GetPaginatedUsersParams) (ua1 []models.User, err error) {
	var cancelFunc func()
	if _d.config.PaginatedTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.PaginatedTimeout)
		defer cancelFunc()
	}
	return _d.User.Paginated(ctx, d, params)
}

// Update implements _sourceRepos.User
func (_d UserWithTimeout) Update(ctx context.Context, d models.DBTX, id models.UserID, params *models.UserUpdateParams) (up1 *models.User, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.User.Update(ctx, d, id, params)
}
