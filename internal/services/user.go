package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type User struct {
	urepo  repos.User
	logger *zap.Logger
	pool   *pgxpool.Pool
}

// use case: mock for handler unit test
type UserService interface {
	Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error)
}

// NewUser returns a new User service.
func NewUser(urepo repos.User, logger *zap.Logger, pool *pgxpool.Pool) *User {
	return &User{
		urepo:  urepo,
		logger: logger,
		pool:   pool,
	}
}

// Create inserts a new user record.
func (u *User) Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error) {
	u.logger.Sugar().Debugf("CreateUser.user: %v", params)

	res, err := u.urepo.Create(ctx, params)
	if err != nil {
		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "urepo.Create")
	}

	return res, nil
}
