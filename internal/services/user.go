package services

import (
	"context"

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
	urepo  UserRepo
}

// NewUser returns a new User service.
func NewUser(urepo UserRepo, logger *zap.Logger) *User {
	return &User{
		logger: logger,
		urepo:  urepo,
	}
}

// Upsert upserts a user record.
func (u *User) Upsert(ctx context.Context, d db.DBTX, user *db.User) error {
	defer newOTELSpan(ctx, "User.Upsert").End()

	if err := u.urepo.Upsert(ctx, d, user); err != nil {
		// TODO database info is leaked if its inaccessible
		return errors.Wrap(err, "urepo.Upsert")
	}

	return nil
}

// Register registers a user record.
func (u *User) Register(ctx context.Context, d db.DBTX, user *db.User) error {
	defer newOTELSpan(ctx, "User.Register").End()

	if err := u.urepo.Create(ctx, d, user); err != nil {
		return errors.Wrap(err, "urepo.Create")
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
