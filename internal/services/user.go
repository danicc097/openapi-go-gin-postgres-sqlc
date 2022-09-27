package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

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

	// TODO remove once traces tested
	DummyMoviePrediction()

	res, err := u.urepo.Create(ctx, params)
	if err != nil {
		// TODO database info is leaked if its inaccessible
		return models.CreateUserResponse{}, errors.Wrap(err, "urepo.Create")
	}

	return res, nil
}
