package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// User defines the datastore/repository handling persisting User records.
// TODO just crud (for impl see if xo for repo and sqlc for services can be used alongside easily
// or need to have some postgen)
type UserRepo interface {
	Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error)
}

type User struct {
	urepo  UserRepo
	logger *zap.Logger
	pool   *pgxpool.Pool
}

// NewUser returns a new User service.
func NewUser(urepo UserRepo, logger *zap.Logger, pool *pgxpool.Pool) *User {
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
