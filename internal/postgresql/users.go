package postgresql

import (
	"context"
	"database/sql"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql/gen"
)

// User represents the repository used for interacting with User records.
type User struct {
	q *gen.Queries
}

// NewUser instantiates the User repository.
func NewUser(d gen.DBTX) *User {
	return &User{
		q: gen.New(d),
	}
}

// TODO methods change role, make admin, etc.

// Create inserts a new user record.
func (t *User) Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error) {
	// TODO creating salt, etc. delegated to jwt.go service
	// https://github.com/appleboy/gin-jwt
	newID, err := t.q.RegisterNewUser(ctx, gen.RegisterNewUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
	}

	return models.CreateUserResponse{
		UserId:      newID,
		AccessToken: "",
	}, nil
}

// Update inserts a new user record.
func (t *User) Update(ctx context.Context, params models.UpdateUserRequest) error {
	err := t.q.UpdateUserById(ctx, gen.UpdateUserByIdParams{
		Username: sql.NullString{String: params.Username, Valid: true},
		Email:    sql.NullString{String: params.Email, Valid: true},
		Password: sql.NullString{String: params.Password, Valid: true},
	})
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
	}

	return nil
}
