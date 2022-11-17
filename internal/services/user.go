package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const synopsis = `
Asian horror cinema often depicts stomach-churning scenes of gore and zombie outbreaks quite vividly and The Sadness ticks all the right boxes.
Chaos and anarchy descend on the city of Taipei as residents turn into mass killers. In the wake of such a deadly viral pandemic, Jim and Kat are a young couple who seek to find each other. Violence, killing and massacre only seem to rise while the government and authorities remain complacent.
Among the most gruesome horror movies of 2022, The Sadness lives up to its name and is not for the faint-hearted. In fact, a trigger warning is also issued at the beginning for those who may not be able to endure watching all the slashing and blood.
`

type User struct {
	logger *zap.Logger
	urepo  repos.User
}

// NewUser returns a new User service.
func NewUser(urepo repos.User, logger *zap.Logger) *User {
	return &User{
		logger: logger,
		urepo:  urepo,
	}
}

// Register registers a user record.
// TODO accepts basic parameters and everything else is default, returns a *db.User. must not pass a db.User here
// IMPORTANT: no endpoint for user creation. Only when coming from auth server.
// we will not support password auth.
func (u *User) Register(ctx context.Context, d db.DBTX, user *db.User) error {
	defer newOTELSpan(ctx, "User.Register").End()

	// TODO construct db.User and fill missing fields with default roles, etc.

	if err := u.urepo.Create(ctx, d, user); err != nil {
		return errors.Wrap(err, "urepo.Create")
	}

	return nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) error {
	defer newOTELSpan(ctx, "User.CreateAPIKey").End()

	if _, err := u.urepo.CreateAPIKey(ctx, d, user); err != nil {
		return errors.Wrap(err, "urepo.CreateAPIKey")
	}

	return nil
}

// UserByEmail gets a user by email.
func (u *User) UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByEmail").End()

	user, err := u.urepo.UserByEmail(ctx, d, email)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByEmail")
	}

	return user, nil
}

// UserByAPIKey gets a user by apiKey.
func (u *User) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByAPIKey").End()

	user, err := u.urepo.UserByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByEmail")
	}

	return user, nil
}
