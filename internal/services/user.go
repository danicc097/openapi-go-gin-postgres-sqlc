package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type User struct {
	urepo UserRepo
	// services can call other services. In this case
	// the interface for this service would defined in the same package where the implementation is called,
	// which makes little sense since in this package we will always want to pass this package's
	// service and not mix-n-match with other service implementation in our handlers.
	// so let's use the struct directly instead to make sure that doesn't happen.
	// this is not the case for repos, where we need to mock them, or pass different ones (e.g. Write-Through Caching Pattern)
	// also note repos should have single responsibility. Do NOT call a repo from another repo, that's business logic
	// that goes into a service.
	someService *SomeService
	logger      *zap.Logger
	pool        *pgxpool.Pool
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
func (u *User) Upsert(ctx context.Context, user *crud.User) error {
	defer newOTELSpan(ctx, "User.Upsert").End()

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

// Register registers a user record.
func (u *User) Register(ctx context.Context, user *crud.User) error {
	defer newOTELSpan(ctx, "User.Register").End()

	err := u.urepo.Create(ctx, user)
	if err != nil {
		return errors.Wrap(err, "urepo.Create")
	}

	return nil
}

// UserByEmail gets a user by email.
func (u *User) UserByEmail(ctx context.Context, email string) (*crud.User, error) {
	defer newOTELSpan(ctx, "User.UserByEmail").End()

	// TODO remove once traces tested
	// TODO counterfeiter on MovieGenreClient, package name <dir>testing with generated pb
	DummyMoviePrediction(ctx)

	user, err := u.urepo.UserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByEmail")
	}

	return user, nil
}
