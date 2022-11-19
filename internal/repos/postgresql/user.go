package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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

func (u *User) Create(ctx context.Context, d db.DBTX, user *db.User) error {
	return user.Save(ctx, d)
}

func (u *User) UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	return db.UserByEmail(ctx, d, email)
}

func (u *User) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	fmt.Printf("apiKey: %v\n", apiKey)
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", parseErrorDetail(err))
	}

	if uak.User == nil {
		return nil, fmt.Errorf("could not join user by api key")
	}

	return uak.User, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error) {
	uak := &db.UserAPIKey{
		APIKey:    uuid.NewString(),
		ExpiresOn: time.Now().AddDate(1, 0, 0),
		UserID:    user.UserID,
	}
	if err := uak.Save(ctx, d); err != nil {
		return nil, fmt.Errorf("could not save api key: %w", parseErrorDetail(err))
	}

	user.APIKeyID = pointers.Int(uak.UserAPIKeyID)
	if err := user.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
	}

	return uak, nil
}
