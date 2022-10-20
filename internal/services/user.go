package services

import (
	"context"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const synopsis = `
Asian horror cinema often depicts stomach-churning scenes of gore and zombie outbreaks quite vividly and The Sadness ticks all the right boxes.
Chaos and anarchy descend on the city of Taipei as residents turn into mass killers. In the wake of such a deadly viral pandemic, Jim and Kat are a young couple who seek to find each other. Violence, killing and massacre only seem to rise while the government and authorities remain complacent.
Among the most gruesome horror movies of 2022, The Sadness lives up to its name and is not for the faint-hearted. In fact, a trigger warning is also issued at the beginning for those who may not be able to endure watching all the slashing and blood.
`

type User struct {
	urepo UserRepo
	// services can call other services. In this case
	// the interface for this service would defined in the same package where the implementation is called,
	// which makes little sense since in this package we will always want to pass this package's
	// service and not mix-n-match with other service implementation in our handlers.
	// so let's use the struct directly instead to make sure that doesn't happen.
	// this is not the case for repos, where we need to mock them, or pass different ones (e.g. Write-Through Caching Pattern)
	// IMPORTANT: repos should have single responsibility. Do NOT call a repo from another repo, that's business logic that goes into a service.
	// regarding testing, service package testing need not mock the service or do they?
	// if we define it as in interface we are forcing consumers to use this package's interface..
	someService *SomeService
	logger      *zap.Logger
	pool        *pgxpool.Pool
	movieSvc    *moviePrediction
}

// NewUser returns a new User service.
func NewUser(urepo UserRepo, logger *zap.Logger, pool *pgxpool.Pool, movieSvcClient v1.MovieGenreClient) *User {
	return &User{
		urepo:  urepo,
		logger: logger,
		pool:   pool,
		// NewMoviePrediction would receive repo interfaces, etc via args to NewUser,
		//  but we dont want to use an interface for moviePrediction.
		// in this package impl we will always want the actual moviePrediction impl. of the package.
		// Note that User service can be mocked the same way if need be since we dont pass any concrete service
		// in args, just the building blocks, which should ALWAYS be interfaces for things we control.
		movieSvc: NewMoviePrediction(movieSvcClient),
	}
}

// Upsert upserts a user record.
func (u *User) Upsert(ctx context.Context, user *crud.User) error {
	defer newOTELSpan(ctx, "User.Upsert").End()

	predictions, _ := u.movieSvc.PredictMovieGenre(ctx, synopsis)
	u.logger.Sugar().Infof("Movie predictions: %v", predictions)

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

	predictions, _ := u.movieSvc.PredictMovieGenre(ctx, synopsis)
	u.logger.Sugar().Infof("Movie predictions: %v", predictions)

	user, err := u.urepo.UserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByEmail")
	}

	return user, nil
}
