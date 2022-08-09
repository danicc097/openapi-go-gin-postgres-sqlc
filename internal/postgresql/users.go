package postgresql

// import (
// 	"context"
// 	"errors"

// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql/gen"
// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v4"
// )

// // User represents the repository used for interacting with User records.
// type User struct {
// 	q *gen.Queries
// }

// // NewUser instantiates the User repository.
// func NewUser(d db.DBTX) *User {
// 	return &User{
// 		q: db.New(d),
// 	}
// }

// // Create inserts a new user record.
// func (t *User) Create(ctx context.Context, params internal.CreateParams) (internal.User, error) {
// 	defer newOTELSpan(ctx, "User.Create").End()

// 	newID, err := t.q.InsertUser(ctx, db.InsertUserParams{
// 		Description: params.Description,
// 		Priority:    newPriority(params.Priority),
// 		StartDate:   newNullTime(params.Dates.Start),
// 		DueDate:     newNullTime(params.Dates.Due),
// 	})
// 	if err != nil {
// 		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
// 	}

// 	return internal.User{
// 		ID:          newID.String(),
// 		Description: params.Description,
// 		Priority:    params.Priority,
// 		Dates:       params.Dates,
// 	}, nil
// }

// // Delete deletes the existing record matching the id.
// func (t *User) Delete(ctx context.Context, id string) error {
// 	defer newOTELSpan(ctx, "User.Delete").End()

// 	//-

// 	val, err := uuid.Parse(id)
// 	if err != nil {
// 		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid uuid")
// 	}

// 	_, err = t.q.DeleteUser(ctx, val)
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
// 		}

// 		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete user")
// 	}

// 	return nil
// }
