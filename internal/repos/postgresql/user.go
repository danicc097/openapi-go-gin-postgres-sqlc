package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/google/uuid"
)

// User represents the repository used for interacting with User records.
type User struct {
	q *db.Queries
}

// NewUser instantiates the User repository.
func NewUser() *User {
	return &User{
		q: db.New(),
	}
}

var _ repos.User = (*User)(nil)

func (u *User) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (*db.User, error) {
	params.Scopes = slices.Unique(params.Scopes)
	user, err := db.CreateUser(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) Update(ctx context.Context, d db.DBTX, id uuid.UUID, params *db.UserUpdateParams) (*db.User, error) {
	user, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user by id: %w", parseErrorDetail(err))
	}

	if params.Scopes != nil {
		*params.Scopes = slices.Unique(*params.Scopes)
	}

	user, err = user.Update(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
	}

	return user, err
}

func (u *User) Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error) {
	user, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user by id %w", parseErrorDetail(err))
	}

	if err := user.SoftDelete(ctx, d); err != nil {
		return nil, fmt.Errorf("could not mark user as deleted: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.User, error) {
	user, err := db.UserByExternalID(ctx, d, extID)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	user, err := db.UserByEmail(ctx, d, email)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error) {
	user, err := db.UserByUsername(ctx, d, username)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByID(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error) {
	user, err := db.UserByUserID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", parseErrorDetail(err))
	}

	if uak.UserJoin == nil {
		return nil, fmt.Errorf("could not join user by api key")
	}

	return uak.UserJoin, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error) {
	uak := &db.UserAPIKey{
		APIKey:    uuid.NewString(),
		ExpiresOn: time.Now().AddDate(1, 0, 0),
		UserID:    user.UserID,
	}
	if _, err := uak.Insert(ctx, d); err != nil {
		return nil, fmt.Errorf("could not save api key: %w", parseErrorDetail(err))
	}

	if _, err := user.Update(ctx, d, &db.UserUpdateParams{APIKeyID: pointers.New(pointers.New(uak.UserAPIKeyID))}); err != nil {
		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
	}

	return uak, nil
}
