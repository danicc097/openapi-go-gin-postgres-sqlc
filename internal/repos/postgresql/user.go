package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
)

// User represents the repository used for interacting with User records.
type User struct {
	q *db.Queries
}

// NewUser instantiates the User repository.
func NewUser(d db.DBTX) *User {
	return &User{
		q: db.New(d),
	}
}

// TODO use xo instead. need triggers
// Create inserts a new user record.
func (u *User) Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error) {
	defer newOTELSpan(ctx, "User.Create").End()

	// TODO logger needs to be passed down to repo as well
	// environment.Logger.Sugar().Infof("users.Create.params: %v", params)
	// TODO creating salt, etc. delegated to jwt.go service
	// https://github.com/appleboy/gin-jwt
	_, err := u.q.GetUser(ctx, db.GetUserParams{
		Username: sql.NullString{String: params.Username, Valid: true},
	})
	if err == nil {
		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("username %s already exists", params.Username))
	}

	_, err = u.q.GetUser(ctx, db.GetUserParams{
		Email: sql.NullString{String: params.Email, Valid: true},
	})
	if err == nil {
		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("email %s already exists", params.Email))
	}

	newID, err := u.q.RegisterNewUser(ctx, db.RegisterNewUserParams{
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
func (u *User) Update(ctx context.Context, params models.UpdateUserRequest) error {
	err := u.q.UpdateUserById(ctx, db.UpdateUserByIdParams{
		Username: sql.NullString{String: params.Username, Valid: true},
		Email:    sql.NullString{String: params.Email, Valid: true},
		Password: sql.NullString{String: params.Password, Valid: true},
	})
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
	}

	return nil
}
