package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// User represents the repository used for interacting with User records.
type User struct {
	q  *db.Queries
	db db.DBTX
}

// NewUser instantiates the User repository.
func NewUser(d db.DBTX) *User {
	// IMPORTANT: d must either be a regular conn or a transaction. repo should not be concerned.
	return &User{
		q:  db.New(), // pgx.Pool also implements db.DBTX
		db: d,
	}
}

// TODO use xo instead. need triggers
// // Create inserts a new user record.
// func (u *User) Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error) {
// 	defer newOTELSpan(ctx, "User.Create").End()

// 	// TODO logger needs to be passed down to repo as well
// 	// environment.Logger.Sugar().Infof("users.Create.params: %v", params)
// 	_, err := u.q.GetUser(ctx, db.GetUserParams{
// 		Username: sql.NullString{String: params.Username, Valid: true},
// 	})
// 	if err == nil {
// 		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("username %s already exists", params.Username))
// 	}

// 	_, err = u.q.GetUser(ctx, db.GetUserParams{
// 		Email: sql.NullString{String: params.Email, Valid: true},
// 	})
// 	if err == nil {
// 		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("email %s already exists", params.Email))
// 	}

// 	newID, err := u.q.RegisterNewUser(ctx, db.RegisterNewUserParams{
// 		Username: params.Username,
// 		Email:    params.Email,
// 		Password: params.Password,
// 	})
// 	if err != nil {
// 		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
// 	}

// 	return models.CreateUserResponse{
// 		UserId:      newID,
// 		AccessToken: "",
// 	}, nil
// }

// Create inserts a new user record.
func (u *User) Create(ctx context.Context, user *db.User) error {
	defer newOTELSpan(ctx, "User.Create").End()
	// https://github.com/xo/xo/blob/master/_examples/booktest/postgres.go
	// != Save, where pks are provided.
	// TODO use pgconn.PgError to handle conflicts (unique key violation) and return
	// internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("??? %s already exists
	// or a generic conflict if no known constraint name matched
	// see
	// https://github.com/jackc/pgx/issues/1334
	// (^ replace hardcoded errors with constants in https://github.com/jackc/pgerrcode/blob/master/errcode.go)
	// https://github.com/jackc/pgx/issues/474
	// (^ latest comments - see https://github.com/jackc/pgerrcode/)

	// save inserts or updates if already exists
	return user.Save(ctx, u.db)
}

// Upsert upserts a new user record.
func (u *User) Upsert(ctx context.Context, user *db.User) error {
	defer newOTELSpan(ctx, "User.Upsert").End()
	// https://github.com/xo/xo/blob/master/_examples/booktest/postgres.go
	return user.Upsert(ctx, u.db)
}

func (u *User) UserByEmail(ctx context.Context, email string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByEmail").End()

	user, err := db.UserByEmail(ctx, u.db, email)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %v", err)
	}
	fmt.Printf("user by email in repo is : %#v\n", user)
	return user, nil
}

// TODO use xo
// // Update inserts a new user record.
// func (u *User) Update(ctx context.Context, params models.UpdateUserRequest) error {
// 	err := u.q.UpdateUserById(ctx, db.UpdateUserByIdParams{
// 		Username: sql.NullString{String: params.Username, Valid: true},
// 		Email:    sql.NullString{String: params.Email, Valid: true},
// 		Password: sql.NullString{String: params.Password, Valid: true},
// 	})
// 	if err != nil {
// 		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
// 	}

// 	return nil
// }
