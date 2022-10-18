package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
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

// Upsert upserts a user record.
func (u *User) Upsert(ctx context.Context, user crud.User) error {
	u.logger.Sugar().Debugf("CreateUser.user: %v", user)

	// TODO remove once traces tested
	// TODO counterfeiter on MovieGenreClient, package name <dir>testing with generated pb
	DummyMoviePrediction(ctx)

	err := u.urepo.Upsert(ctx, user)
	if err != nil {
		// TODO database info is leaked if its inaccessible
		return errors.Wrap(err, "urepo.Upsert")
	}

	return nil
}
