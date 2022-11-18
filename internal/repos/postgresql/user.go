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

// TODO use xo instead. need triggers
// // Create inserts a new user record.
// func (u *User) Create(ctx context.Context, d db.DBTX, params models.CreateUserRequest) (models.CreateUserResponse, error) {

// 	// TODO logger needs to be passed down to repo as well
// 	// environment.Logger.Sugar().Infof("users.Create.params: %v", params)
// 	_, err := u.q.GetUser(ctx, db.GetUserParams{
// 		Username: sql.NullString{String: params.Username},
// 	})
// 	if err == nil {
// 		return models.CreateUserResponse{}, internal.WrapErrorf(err, internal.ErrorCodeAlreadyExists, fmt.Sprintf("username %s already exists", params.Username))
// 	}

// 	_, err = u.q.GetUser(ctx, db.GetUserParams{
// 		Email: sql.NullString{String: params.Email},
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
func (u *User) Create(ctx context.Context, d db.DBTX, user *db.User) error {
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
	// IMPORTANT: the above is useless for unique constraints (if code is unique violation 23505, we need to know where,
	// a conflict
	// status code is not enough). we can use a simple regex if we get access to postgres error:
	// 	ERROR:  23505: duplicate key value violates unique constraint "users_external_id_key"
	// DETAIL:  Key (external_id)=(provider_external_id1) already exists.

	// save is almost the same as insert but detects if the user struct was created before
	// in which case it updated instead
	err := user.Save(ctx, d)
	if err != nil {
		// TODO return internal error with appropiate code (that gets converted to status code) and ui shows error.Root()
		// e.g. username "superadmin" already exists
		return fmt.Errorf("could not create user: %v", parseErrorDetail(err))
	}

	return nil
}

func (u *User) UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	user, err := db.UserByEmail(ctx, d, email)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %v", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %v", parseErrorDetail(err))
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
		return nil, fmt.Errorf("could not save api key: %v", parseErrorDetail(err))
	}

	user.APIKeyID = pointers.Int(uak.UserAPIKeyID)
	if err := user.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("could not update user: %v", parseErrorDetail(err))
	}

	return uak, nil
}

// TODO use xo
// // Update inserts a new user record.
// func (u *User) Update(ctx context.Context, d db.DBTX, params models.UpdateUserRequest) error {
// 	err := u.q.UpdateUserById(ctx, db.UpdateUserByIdParams{
// 		Username: sql.NullString{String: params.Username},
// 		Email:    sql.NullString{String: params.Email},
// 		Password: sql.NullString{String: params.Password},
// 	})
// 	if err != nil {
// 		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
// 	}

// 	return nil
// }
